package repository

import (
	"database/sql"
	"encoding/json"
	"jdshow/internal/company"
	"jdshow/internal/job"
	"log/slog"
	_ "modernc.org/sqlite"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type Store struct{ db *sql.DB }

func Open(path string) (*Store, error) {
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return nil, err
	}
	db, err := sql.Open("sqlite", path)
	if err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(1)
	return &Store{db: db}, nil
}
func (s *Store) Close() error { return s.db.Close() }
func (s *Store) Migrate() error {
	paths := []string{"db/migrations/001_init.sql", "../../db/migrations/001_init.sql"}
	var b []byte
	var err error
	for _, path := range paths {
		b, err = os.ReadFile(path)
		if err == nil {
			break
		}
	}
	if err != nil {
		return err
	}
	if _, err = s.db.Exec(string(b)); err != nil {
		return err
	}
	for _, statement := range []string{
		`ALTER TABLE companies ADD COLUMN career_name TEXT NOT NULL DEFAULT ''`,
		`ALTER TABLE companies ADD COLUMN career_url TEXT NOT NULL DEFAULT ''`,
	} {
		if _, err = s.db.Exec(statement); err != nil && !strings.Contains(err.Error(), "duplicate column name") {
			return err
		}
	}
	return nil
}
func (s *Store) DB() *sql.DB { return s.db }
func (s *Store) SeedCompany(c company.Company) error {
	typ, err := s.dict("company_types", "private", c.Type)
	if err != nil {
		return err
	}
	_, err = s.db.Exec(`INSERT INTO companies(id,name,english_name,company_type_id,business_summary,employee_size,background,website_url,career_name,career_url,logo_url,data_status) VALUES(?,?,?,?,?,?,?,?,?,?,?,?) ON CONFLICT(id) DO UPDATE SET name=excluded.name,english_name=excluded.english_name,company_type_id=excluded.company_type_id,business_summary=excluded.business_summary,employee_size=excluded.employee_size,background=excluded.background,website_url=excluded.website_url,career_name=excluded.career_name,career_url=excluded.career_url,logo_url=excluded.logo_url,data_status=excluded.data_status,updated_at=CURRENT_TIMESTAMP`, c.ID, c.Name, c.EnglishName, typ, c.Description, coalesce(c.EmployeeSize, "待核验"), coalesce(c.Background, "待核验"), c.Website, c.CareerName, c.CareerURL, c.LogoURL, coalesce(c.DataStatus, "candidate_pending_verification"))
	if err != nil {
		return err
	}
	if _, err = s.db.Exec(`DELETE FROM company_cities WHERE company_id=?`, c.ID); err != nil {
		return err
	}
	city, err := s.dict("cities", "city", coalesce(c.City, "待核验"))
	if err == nil {
		_, err = s.db.Exec(`INSERT OR IGNORE INTO company_cities(company_id,city_id) VALUES(?,?)`, c.ID, city)
	}
	if err != nil {
		return err
	}
	if _, err = s.db.Exec(`DELETE FROM company_sources WHERE company_id=?`, c.ID); err != nil {
		return err
	}
	for _, source := range c.Sources {
		if _, err = s.db.Exec(`INSERT INTO company_sources(company_id,source_type,source_url,source_title,supports_fields,verified_at) VALUES(?,?,?,?,?,?)`, c.ID, source.Type, source.URL, source.Title, source.SupportsFields, source.VerifiedAt); err != nil {
			return err
		}
	}
	if _, err = s.db.Exec(`DELETE FROM company_industries WHERE company_id=?`, c.ID); err != nil {
		return err
	}
	ind, err := s.dict("industries", "industry", coalesce(c.Industry, "待核验"))
	if err == nil {
		_, err = s.db.Exec(`INSERT OR IGNORE INTO company_industries(company_id,industry_id) VALUES(?,?)`, c.ID, ind)
	}
	return err
}
func (s *Store) SeedJob(j job.Job) error {
	jt, err := s.dict("job_types", "job", j.Category)
	if err != nil {
		return err
	}
	req, _ := json.Marshal(j.Requirements)
	skills, _ := json.Marshal(j.Skills)
	_, err = s.db.Exec(`INSERT INTO jobs(id,company_id,job_type_id,title,city,experience,education,salary,status,published_at,updated_at,source_name,source_url,description,requirements_json,skills_json) VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?) ON CONFLICT(id) DO UPDATE SET company_id=excluded.company_id,title=excluded.title,city=excluded.city,experience=excluded.experience,education=excluded.education,salary=excluded.salary,status=excluded.status,source_url=excluded.source_url,description=excluded.description,requirements_json=excluded.requirements_json,skills_json=excluded.skills_json`, j.ID, j.CompanyID, jt, j.Title, j.City, j.Experience, j.Education, j.Salary, j.Status, dateValue(j.PublishedAt), dateValue(j.UpdatedAt), j.SourceName, j.SourceURL, j.Description, string(req), string(skills))
	return err
}
func (s *Store) dict(table, code, name string) (int64, error) {
	var id int64
	err := s.db.QueryRow(`SELECT id FROM `+table+` WHERE name=?`, name).Scan(&id)
	if err == nil {
		return id, nil
	}
	if err != sql.ErrNoRows {
		return 0, err
	}
	res, err := s.db.Exec(`INSERT INTO `+table+`(code,name) VALUES(?,?)`, code+"-"+strings.ReplaceAll(strings.ReplaceAll(name, " ", "_"), "/", "_"), name)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}
func (s *Store) Companies(typ, city, industry string) ([]company.Company, error) {
	sources := map[string][]company.Source{}
	sourceRows, err := s.db.Query(`SELECT company_id,source_type,source_url,source_title,supports_fields,verified_at FROM company_sources ORDER BY id`)
	if err != nil {
		return nil, err
	}
	for sourceRows.Next() {
		var id string
		var source company.Source
		if err := sourceRows.Scan(&id, &source.Type, &source.URL, &source.Title, &source.SupportsFields, &source.VerifiedAt); err != nil {
			sourceRows.Close()
			return nil, err
		}
		sources[id] = append(sources[id], source)
	}
	sourceRows.Close()
	q := `SELECT c.id,c.name,c.english_name,ct.name,c.business_summary,c.employee_size,c.background,c.website_url,c.career_name,c.career_url,c.logo_url,c.data_status,COALESCE((SELECT ci.name FROM company_cities cc JOIN cities ci ON ci.id=cc.city_id WHERE cc.company_id=c.id LIMIT 1),''),COALESCE((SELECT i.name FROM company_industries co JOIN industries i ON i.id=co.industry_id WHERE co.company_id=c.id LIMIT 1),'') FROM companies c JOIN company_types ct ON ct.id=c.company_type_id WHERE 1=1`
	args := []any{}
	if typ != "" {
		q += " AND ct.name=?"
		args = append(args, typ)
	}
	if city != "" {
		q += " AND EXISTS(SELECT 1 FROM company_cities cc JOIN cities ci ON ci.id=cc.city_id WHERE cc.company_id=c.id AND ci.name=?)"
		args = append(args, city)
	}
	if industry != "" {
		q += " AND EXISTS(SELECT 1 FROM company_industries co JOIN industries i ON i.id=co.industry_id WHERE co.company_id=c.id AND i.name=?)"
		args = append(args, industry)
	}
	q += " ORDER BY c.name"
	rows, err := s.db.Query(q, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	out := []company.Company{}
	for rows.Next() {
		var c company.Company
		if err := rows.Scan(&c.ID, &c.Name, &c.EnglishName, &c.Type, &c.Description, &c.EmployeeSize, &c.Background, &c.Website, &c.CareerName, &c.CareerURL, &c.LogoURL, &c.DataStatus, &c.City, &c.Industry); err != nil {
			return nil, err
		}
		c.Sources = sources[c.ID]
		out = append(out, c)
	}
	return out, rows.Err()
}
func (s *Store) Jobs() ([]job.Job, error) {
	rows, err := s.db.Query(`SELECT j.id,j.company_id,j.title,j.city,COALESCE(jt.name,''),j.experience,education,salary,status,published_at,updated_at,source_name,source_url,description,requirements_json,skills_json FROM jobs j LEFT JOIN job_types jt ON jt.id=j.job_type_id ORDER BY published_at DESC`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	out := []job.Job{}
	for rows.Next() {
		var j job.Job
		var typ string
		var pub, upd sql.NullString
		var req, skills string
		if err := rows.Scan(&j.ID, &j.CompanyID, &j.Title, &j.City, &typ, &j.Experience, &j.Education, &j.Salary, &j.Status, &pub, &upd, &j.SourceName, &j.SourceURL, &j.Description, &req, &skills); err != nil {
			return nil, err
		}
		j.Category = typ
		if pub.Valid {
			j.PublishedAt, _ = time.Parse(time.RFC3339Nano, pub.String)
		}
		if upd.Valid {
			j.UpdatedAt, _ = time.Parse(time.RFC3339Nano, upd.String)
		}
		_ = json.Unmarshal([]byte(req), &j.Requirements)
		_ = json.Unmarshal([]byte(skills), &j.Skills)
		out = append(out, j)
	}
	return out, rows.Err()
}
func dateValue(t time.Time) any {
	if t.IsZero() {
		return nil
	}
	return t.Format(time.RFC3339Nano)
}
func coalesce(v, fallback string) string {
	if v == "" {
		return fallback
	}
	return v
}
func LoadLog(logger *slog.Logger, err error) {
	if err != nil {
		logger.Error("database operation failed", "error", err)
	}
}
