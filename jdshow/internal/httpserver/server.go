package httpserver

import (
	"html/template"
	"jdshow/internal/company"
	"jdshow/internal/job"
	"log/slog"
	"net/http"
	"sort"
	"strings"
	"time"
)

type Server struct {
	companies map[string]company.Company
	jobs      []job.Job
	templates *template.Template
	logger    *slog.Logger
}
type Data struct {
	Title, Active, Page, Query, City, Category, CompanyType, CompanyCity, Industry string
	Jobs                                                                           []job.Job
	Companies                                                                      []company.Company
	Job                                                                            job.Job
	Company                                                                        company.Company
	Stats                                                                          Stats
}
type Stats struct{ ActiveJobs, NewToday, NewThisWeek, CompanyCount, SourceCount int }

func New(companies map[string]company.Company, jobs []job.Job, logger *slog.Logger) *Server {
	f := template.FuncMap{"lower": strings.ToLower, "companyName": func(id string) string { return companies[id].Name }, "formatDate": func(t time.Time) string { return t.Format("01月02日") }, "formatDateTime": func(t time.Time) string { return t.Format("2006-01-02 15:04") }, "cities": func() []string { return []string{"北京", "上海", "深圳", "杭州", "广州", "成都"} }, "categories": func() []string {
		return []string{"后端开发", "算法 / AI", "云计算 / 基础设施", "数据工程", "产品技术"}
	}, "companyTypes": func() []string {
		return []string{"互联网 / AI", "国企 / 央企科技", "外企科技", "AI 公司"}
	}, "companyCities": func() []string { return []string{"上海", "北京", "深圳", "杭州", "广州", "成都", "全国"} }, "industries": func() []string {
		return []string{"上海互联网与科技", "互联网与 AI", "云计算与数字基础设施", "软件、云计算与工业科技", "搜索、AI、云计算", "大模型、智能应用", "云计算、通信、算力", "软件、云计算、AI"}
	}}
	return &Server{companies: companies, jobs: jobs, templates: template.Must(template.New("pages").Funcs(f).ParseFS(templateFiles, "templates/*.html")), logger: logger}
}
func (s *Server) Routes() http.Handler {
	m := http.NewServeMux()
	m.HandleFunc("GET /", s.home)
	m.HandleFunc("GET /jobs", s.jobsPage)
	m.HandleFunc("GET /jobs/{id}", s.jobDetail)
	m.HandleFunc("GET /companies", s.companiesPage)
	m.HandleFunc("GET /companies/{id}", s.companyDetail)
	m.Handle("GET /static/", http.StripPrefix("/static/", http.FileServer(http.Dir("web/static"))))
	m.HandleFunc("GET /healthz", health)
	m.HandleFunc("GET /readyz", health)
	return headers(logging(m, s.logger))
}
func (s *Server) home(w http.ResponseWriter, r *http.Request) {
	j := append([]job.Job{}, s.jobs...)
	sort.Slice(j, func(i, k int) bool { return j[i].PublishedAt.After(j[k].PublishedAt) })
	s.render(w, "home", Data{Title: "招聘雷达", Active: "home", Jobs: j[:min(4, len(j))], Companies: s.companyList()[:min(4, len(s.companies))], Stats: s.stats()})
}
func (s *Server) jobsPage(w http.ResponseWriter, r *http.Request) {
	q, c, cat := r.URL.Query().Get("q"), r.URL.Query().Get("city"), r.URL.Query().Get("category")
	out := []job.Job{}
	for _, j := range s.jobs {
		x := s.companies[j.CompanyID]
		if q != "" && !strings.Contains(strings.ToLower(j.Title+" "+j.Description+" "+strings.Join(j.Skills, " ")+x.Name), strings.ToLower(q)) {
			continue
		}
		if c != "" && j.City != c {
			continue
		}
		if cat != "" && j.Category != cat {
			continue
		}
		out = append(out, j)
	}
	sort.Slice(out, func(i, k int) bool { return out[i].PublishedAt.After(out[k].PublishedAt) })
	s.render(w, "jobs", Data{Title: "搜索岗位", Active: "jobs", Jobs: out, Query: q, City: c, Category: cat})
}
func (s *Server) jobDetail(w http.ResponseWriter, r *http.Request) {
	for _, j := range s.jobs {
		if j.ID == r.PathValue("id") {
			s.render(w, "job", Data{Title: j.Title, Active: "detail", Job: j, Company: s.companies[j.CompanyID]})
			return
		}
	}
	http.NotFound(w, r)
}
func (s *Server) companiesPage(w http.ResponseWriter, r *http.Request) {
	typ, city, ind := r.URL.Query().Get("type"), r.URL.Query().Get("city"), r.URL.Query().Get("industry")
	out := []company.Company{}
	for _, c := range s.companyList() {
		if typ != "" && c.Type != typ {
			continue
		}
		if city != "" && c.City != city {
			continue
		}
		if ind != "" && c.Industry != ind {
			continue
		}
		out = append(out, c)
	}
	s.render(w, "companies", Data{Title: "公司库", Active: "companies", Page: "company-list", Companies: out, CompanyType: typ, CompanyCity: city, Industry: ind})
}
func (s *Server) companyDetail(w http.ResponseWriter, r *http.Request) {
	c, ok := s.companies[r.PathValue("id")]
	if !ok {
		http.NotFound(w, r)
		return
	}
	out := []job.Job{}
	for _, j := range s.jobs {
		if j.CompanyID == c.ID {
			out = append(out, j)
		}
	}
	s.render(w, "company", Data{Title: c.Name, Active: "companies", Page: "company-detail", Company: c, Jobs: out})
}
func (s *Server) companyList() []company.Company {
	out := []company.Company{}
	for _, c := range s.companies {
		out = append(out, c)
	}
	sort.Slice(out, func(i, j int) bool { return out[i].Name < out[j].Name })
	return out
}
func (s *Server) stats() Stats {
	out := Stats{CompanyCount: len(s.companies)}
	src := map[string]bool{}
	for _, j := range s.jobs {
		if j.Status != "REMOVED" {
			out.ActiveJobs++
		}
		if j.PublishedAt.After(time.Now().Add(-24 * time.Hour)) {
			out.NewToday++
		}
		if j.PublishedAt.After(time.Now().Add(-7 * 24 * time.Hour)) {
			out.NewThisWeek++
		}
		src[j.SourceName] = true
	}
	out.SourceCount = len(src)
	return out
}
func (s *Server) render(w http.ResponseWriter, name string, d Data) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := s.templates.ExecuteTemplate(w, name, d); err != nil {
		s.logger.Error("template render failed", "error", err)
	}
}
func health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"status":"ok"}`))
}
func logging(next http.Handler, l *slog.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		st := time.Now()
		next.ServeHTTP(w, r)
		l.Info("http request", "method", r.Method, "path", r.URL.Path, "duration_ms", time.Since(st).Milliseconds())
	})
}
func headers(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.Header().Set("X-Frame-Options", "DENY")
		w.Header().Set("Referrer-Policy", "strict-origin-when-cross-origin")
		next.ServeHTTP(w, r)
	})
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
