package httpserver

import (
	"io"
	"jdshow/internal/demo"
	"jdshow/internal/source/shanghai"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func testHandler() http.Handler {
	companies, jobs := demo.Seed(time.Now())
	companies, jobs = demo.Catalog(companies, jobs, time.Now())
	companies, jobs = shanghai.Load("../../data/company_jobs/private/shanghai_internet", companies, jobs, slog.New(slog.NewTextHandler(io.Discard, nil)))
	return New(companies, jobs, slog.New(slog.NewTextHandler(io.Discard, nil))).Routes()
}
func TestPages(t *testing.T) {
	for _, path := range []string{"/", "/companies", "/companies/baidu", "/jobs", "/jobs/job-001"} {
		r := httptest.NewRequest(http.MethodGet, path, nil)
		w := httptest.NewRecorder()
		testHandler().ServeHTTP(w, r)
		if w.Code != http.StatusOK {
			t.Fatalf("%s: got %d", path, w.Code)
		}
	}
}
func TestCompanyFilterLoadsDewu(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "/companies?industry=上海互联网与科技&type=互联网%20%2F%20AI&city=上海", nil)
	w := httptest.NewRecorder()
	testHandler().ServeHTTP(w, r)
	if !strings.Contains(w.Body.String(), "得物") {
		t.Fatal("Dewu is missing from filtered company page")
	}
}
func TestVerifiedSourceLink(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "/jobs/job-001", nil)
	w := httptest.NewRecorder()
	testHandler().ServeHTTP(w, r)
	if !strings.Contains(w.Body.String(), "https://talent.baidu.com/jobs/social-list") {
		t.Fatal("official source link missing")
	}
}
func TestHealth(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "/healthz", nil)
	w := httptest.NewRecorder()
	health(w, r)
	if w.Code != http.StatusOK || !strings.Contains(w.Body.String(), "ok") {
		t.Fatal("health check failed")
	}
}
