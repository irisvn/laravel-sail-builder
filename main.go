package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	r := gin.New()

	// Logger middleware
	r.Use(gin.Logger())

	// Recovery middleware
	r.Use(gin.Recovery())

	r.GET("/:folder", func(c *gin.Context) {
		// Get parameters from request
		folder := strings.TrimSpace(c.Param("folder"))
		version := strings.TrimSpace(c.Query("version"))
		with := strings.TrimSpace(c.Query("with"))

		if folder == "" {
			c.String(http.StatusBadRequest, "Invalid URL: folder parameter is required")
			return
		}

		if !isValidFolder(folder) {
			c.String(http.StatusBadRequest, "Invalid URL: folder string contains only alpha-numeric characters, dashes, and underscores")
			return
		}

		if !isValidVersion(version) {
			c.String(http.StatusBadRequest, "Invalid URL: version string contains only digits and \".\"")
			return
		}

		if !isValidWith(with) {
			c.String(http.StatusBadRequest, "Invalid URL: Please provide one or more of the supported services (mysql, pgsql, mariadb, redis, memcached, meilisearch, typesense, minio, mailpit, selenium, soketi) or none.")
			return
		}

		// Read template file
		templateData, err := ioutil.ReadFile("laravel.temp")
		if err != nil {
			c.String(http.StatusInternalServerError, "Error reading template file: %v", err)
			return
		}

		// Replace placeholders with actual values
		output := strings.ReplaceAll(string(templateData), "{{folder}}", folder)

		// Check if version is not empty, then replace
		if version != "" {
			output = strings.ReplaceAll(output, "{{version}}", ":^"+version)
		} else {
			// If version is empty, remove it from the template
			output = strings.ReplaceAll(output, "{{version}}", "")
		}

		// Check if with is not empty, then replace
		if with != "" {
			output = strings.ReplaceAll(output, "{{with}}", with)
		} else {
			// If with is empty, remove it from the template
			output = strings.ReplaceAll(output, "{{with}}", "mysql,redis,meilisearch,mailpit,selenium")
		}

		// Return response
		c.String(http.StatusOK, output)
	})

	fmt.Println("Server is running at http://localhost:8080")
	err := r.Run(":8080")
	if err != nil {
		panic(err)
	}
}

// isValidFolder checks if the folder string contains only alpha-numeric characters, dashes, and underscores
func isValidFolder(folder string) bool {
	// Define regular expression pattern for folder validation
	pattern := "^[a-zA-Z0-9_-]+$"
	match, _ := regexp.MatchString(pattern, folder)
	return match
}

// isValidVersion checks if the version string contains only digits and "."
func isValidVersion(version string) bool {
	// Define regular expression pattern for version validation
	pattern := "^[0-9.]*$"
	match, _ := regexp.MatchString(pattern, version)
	return match
}

// isValidWith checks if the with string contains supported services or "none"
func isValidWith(with string) bool {

	// Define regular expression pattern for with validation
	pattern := "^[a-z,]*$"
	match, _ := regexp.MatchString(pattern, with)
	if !match {
		return false
	}

	// List of supported services
	supportedServices := []string{
		"mysql", "pgsql", "mariadb", "redis", "memcached",
		"meilisearch", "typesense", "minio", "mailpit",
		"selenium", "soketi", "none", "",
	}

	// Split the with string by commas
	services := strings.Split(with, ",")

	// Check if each service is in supported services
	for _, service := range services {
		// Trim the service name
		service = strings.TrimSpace(service)

		// Check if the service is in supported services
		found := false
		for _, supportedService := range supportedServices {
			if service == supportedService {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}

	return true
}
