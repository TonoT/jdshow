package company

type Source struct {
	Type           string
	URL            string
	Title          string
	SupportsFields string
	VerifiedAt     string
}

type Company struct {
	ID           string
	Name         string
	EnglishName  string
	Type         string
	Industry     string
	City         string
	Description  string
	EmployeeSize string
	Background   string
	Tags         []string
	Website      string
	CareerName   string
	CareerURL    string
	LogoURL      string
	DataStatus   string
	Sources      []Source
}
