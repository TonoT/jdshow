package repository

import (
	"jdshow/internal/company"
	"jdshow/internal/job"
	"testing"
	"time"
)

func TestStoreSeedsAndQueriesRelatedData(t *testing.T) {
	store, err := Open(t.TempDir() + "/test.sqlite")
	if err != nil {
		t.Fatal(err)
	}
	defer store.Close()
	if err := store.Migrate(); err != nil {
		t.Fatal(err)
	}
	c := company.Company{ID: "company-1", Name: "测试公司", Type: "互联网 / AI", Industry: "上海互联网与科技", City: "上海", Description: "公司介绍", EmployeeSize: "100-499", Background: "民营科技公司", DataStatus: "verified"}
	if err := store.SeedCompany(c); err != nil {
		t.Fatal(err)
	}
	j := job.Job{ID: "job-1", CompanyID: c.ID, Title: "Go 工程师", City: "上海", Category: "后端开发", PublishedAt: time.Now(), Requirements: []string{"Go"}, Skills: []string{"Go", "SQLite"}}
	if err := store.SeedJob(j); err != nil {
		t.Fatal(err)
	}
	if err := store.SeedJob(j); err != nil {
		t.Fatal(err)
	}
	companies, err := store.Companies("互联网 / AI", "上海", "上海互联网与科技")
	if err != nil || len(companies) != 1 || companies[0].EmployeeSize != "100-499" {
		t.Fatalf("company query failed: err=%v companies=%+v", err, companies)
	}
	jobs, err := store.Jobs()
	if err != nil || len(jobs) != 1 || jobs[0].CompanyID != c.ID || len(jobs[0].Skills) != 2 {
		t.Fatalf("job query failed: err=%v jobs=%+v", err, jobs)
	}
}
