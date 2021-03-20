// models/website.go

package models

// Website model
type Website struct {
	ID   uint   `json:"id" gorm:"primary_key"`
	Name string `json:"name"`
	URL  string `json:"url"`
}

// CreateWebsiteInput model
type CreateWebsiteInput struct {
	Name string `json:"name" binding:"required"`
	URL  string `json:"url" binding:"required"`
}

// UpdateWebsiteInput model
type UpdateWebsiteInput struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
