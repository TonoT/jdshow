package shanghai

import (
	"encoding/json"
	"jdshow/internal/company"
	"jdshow/internal/job"
	"log/slog"
	"os"
	"path/filepath"
)

type index struct {
	Companies []struct {
		CompanyID   string `json:"company_id"`
		CompanyName string `json:"company_name"`
		File        string `json:"file"`
	} `json:"companies"`
}
type record struct {
	CompanyID       string    `json:"company_id"`
	CompanyName     string    `json:"company_name"`
	CompanyType     string    `json:"company_type"`
	PrimaryIndustry string    `json:"primary_industry"`
	IndustryTags    []string  `json:"industry_tags"`
	CityFocus       string    `json:"city_focus"`
	OfficialSiteURL string    `json:"official_site_url"`
	Jobs            []job.Job `json:"jobs"`
}

func Load(root string, companies map[string]company.Company, jobs []job.Job, logger *slog.Logger) (map[string]company.Company, []job.Job) {
	b, err := os.ReadFile(filepath.Join(root, "index.json"))
	if err != nil {
		return companies, jobs
	}
	var idx index
	if json.Unmarshal(b, &idx) != nil {
		return companies, jobs
	}
	names := map[string]bool{}
	for _, c := range companies {
		names[c.Name] = true
	}
	for _, e := range idx.Companies {
		b, err = os.ReadFile(filepath.Join(root, e.File))
		if err != nil {
			logger.Warn("company file unavailable", "file", e.File)
			continue
		}
		var r record
		if json.Unmarshal(b, &r) != nil || r.CompanyID == "" || r.CompanyName == "" || names[r.CompanyName] {
			continue
		}
		industry := r.PrimaryIndustry
		if industry == "shanghai_internet" {
			industry = "上海互联网与科技"
		}
		typ := r.CompanyType
		if typ == "private" {
			typ = "互联网 / AI"
		}
		companies[r.CompanyID] = company.Company{ID: r.CompanyID, Name: r.CompanyName, EnglishName: r.CompanyName, Type: typ, Industry: industry, City: r.CityFocus, Description: r.CompanyName + " 已纳入上海互联网与科技公司候选池，官网和招聘入口待核验。", Tags: append(r.IndustryTags, "来源待核验"), Website: r.OfficialSiteURL}
		names[r.CompanyName] = true
		jobs = append(jobs, r.Jobs...)
	}
	return companies, jobs
}
