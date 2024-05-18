package axiston

import (
	"fmt"
	"os"
)

// ClientVersion is the current version of this client.
var ClientVersion string = "0.1.0"

// GoVersion is the required version of the Go runtime.
var GoVersion string = "1.22.2"

var envApiKey string = getEnv("AXISTON_API_KEY", "")
var userAgent = fmt.Sprintf("Axiston/%s (Go; Ver. %s)", ClientVersion, GoVersion)
var envUserAgent string = getEnv("AXISTON_USER_AGENT", userAgent)
var envBaseUrl string = getEnv("AXISTON_BASE_URL", "https://api.axiston.com")

func getEnv(key, df string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return df
}
