package app

import (
	"jdshow/internal/company"
	"jdshow/internal/httpserver"
	"jdshow/internal/repository"
	"log/slog"
)

func New(logger *slog.Logger) *httpserver.Server {
	store, err := repository.Open("data/db/jdshow.sqlite")
	if err != nil {
		logger.Error("database open failed", "error", err)
		return httpserver.New(map[string]company.Company{}, nil, logger)
	}
	if err := store.Migrate(); err != nil {
		logger.Error("database migration failed", "error", err)
		return httpserver.New(map[string]company.Company{}, nil, logger)
	}

	dbCompanies, err := store.Companies("", "", "")
	if err != nil {
		logger.Error("company query failed", "error", err)
		return httpserver.New(map[string]company.Company{}, nil, logger)
	}
	dbJobs, err := store.Jobs()
	if err != nil {
		logger.Error("job query failed", "error", err)
		return httpserver.New(map[string]company.Company{}, nil, logger)
	}
	companyMap := make(map[string]company.Company, len(dbCompanies))
	for _, c := range dbCompanies {
		companyMap[c.ID] = c
	}
	return httpserver.New(companyMap, dbJobs, logger)
}
