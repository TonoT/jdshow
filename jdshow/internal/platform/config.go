package platform

import "os"

func Address() string {
	if value := os.Getenv("ADDRESS"); value != "" {
		return value
	}
	return ":8080"
}
