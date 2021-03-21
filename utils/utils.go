// utils/utils.go

package utils

import (
	"is-my-website-down/models"
	"net/http"
	"net/url"
)

func IsWebsiteDown(website models.Website, c chan string) {
	_, err := http.Get(website.URL)
	if err != nil {
		c <- website.URL + " - " + website.Name + " - DOWN!"
		return
	}

	c <- ""
}

func IsURL(websiteUrl string) bool {
	_, err := url.ParseRequestURI(websiteUrl)
	if err != nil {
		return false
	}
	return true
}
