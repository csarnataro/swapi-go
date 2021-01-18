package utils

import (
	"net/http"
	"strings"
)

// GetServerName builds the server name to use as a prefix in feeds
func GetServerName(request *http.Request) string {
	protocol := "https"
	// useful when developing in local env
	if strings.Contains(request.Host, "localhost") || strings.Contains(request.Host, "127.0.0.1") {
		protocol = "http"
	}
	return protocol + "://" + request.Host
}
