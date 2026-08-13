PRAGMA foreign_keys = ON;

CREATE TABLE IF NOT EXISTS company_types (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  code TEXT NOT NULL UNIQUE,
  name TEXT NOT NULL UNIQUE
);
CREATE TABLE IF NOT EXISTS industries (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  code TEXT NOT NULL UNIQUE,
  name TEXT NOT NULL UNIQUE
);
CREATE TABLE IF NOT EXISTS cities (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  code TEXT NOT NULL UNIQUE,
  name TEXT NOT NULL UNIQUE
);
CREATE TABLE IF NOT EXISTS companies (
  id TEXT PRIMARY KEY,
  name TEXT NOT NULL UNIQUE,
  english_name TEXT NOT NULL DEFAULT '',
  company_type_id INTEGER NOT NULL REFERENCES company_types(id),
  business_summary TEXT NOT NULL DEFAULT '',
  employee_size TEXT NOT NULL DEFAULT '待核验',
  background TEXT NOT NULL DEFAULT '待核验',
  website_url TEXT NOT NULL DEFAULT '',
  career_name TEXT NOT NULL DEFAULT '',
  career_url TEXT NOT NULL DEFAULT '',
  logo_url TEXT NOT NULL DEFAULT '',
  data_status TEXT NOT NULL DEFAULT 'candidate_pending_verification',
  created_at TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP
);
CREATE TABLE IF NOT EXISTS company_sources (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  company_id TEXT NOT NULL REFERENCES companies(id) ON DELETE CASCADE,
  source_type TEXT NOT NULL,
  source_url TEXT NOT NULL,
  source_title TEXT NOT NULL DEFAULT '',
  supports_fields TEXT NOT NULL DEFAULT '',
  verified_at TEXT NOT NULL,
  UNIQUE(company_id, source_url)
);
CREATE TABLE IF NOT EXISTS company_cities (
  company_id TEXT NOT NULL REFERENCES companies(id) ON DELETE CASCADE,
  city_id INTEGER NOT NULL REFERENCES cities(id),
  PRIMARY KEY (company_id, city_id)
);
CREATE TABLE IF NOT EXISTS company_industries (
  company_id TEXT NOT NULL REFERENCES companies(id) ON DELETE CASCADE,
  industry_id INTEGER NOT NULL REFERENCES industries(id),
  PRIMARY KEY (company_id, industry_id)
);
CREATE TABLE IF NOT EXISTS job_types (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  code TEXT NOT NULL UNIQUE,
  name TEXT NOT NULL UNIQUE
);
CREATE TABLE IF NOT EXISTS jobs (
  id TEXT PRIMARY KEY,
  company_id TEXT NOT NULL REFERENCES companies(id) ON DELETE CASCADE,
  job_type_id INTEGER REFERENCES job_types(id),
  title TEXT NOT NULL,
  city TEXT NOT NULL DEFAULT '',
  experience TEXT NOT NULL DEFAULT '',
  education TEXT NOT NULL DEFAULT '',
  salary TEXT NOT NULL DEFAULT '',
  status TEXT NOT NULL DEFAULT 'ACTIVE',
  published_at TEXT,
  updated_at TEXT,
  source_name TEXT NOT NULL DEFAULT '',
  source_url TEXT NOT NULL DEFAULT '',
  description TEXT NOT NULL DEFAULT '',
  requirements_json TEXT NOT NULL DEFAULT '[]',
  skills_json TEXT NOT NULL DEFAULT '[]'
);
CREATE INDEX IF NOT EXISTS idx_companies_type ON companies(company_type_id);
CREATE INDEX IF NOT EXISTS idx_jobs_company ON jobs(company_id);
CREATE INDEX IF NOT EXISTS idx_jobs_status_date ON jobs(status, published_at DESC);
