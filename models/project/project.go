package project

import (
	"time"
)
func (Project) TableName() string {
    return "Project"
}
// Project represents the structure of your table data in Go.
type Project struct {
    ID            string    `gorm:"column:id;primaryKey" json:"id"`
    Name          string
    Description   string
    ClientId      string    `gorm:"uniqueKey;column:clientId" json:"clientId"`
    ClientSecret  string    `gorm:"column:clientSecret" json:"clientSecret"`
    Slug          string
    IsActive      bool
    CreatedAt     time.Time
    UpdatedAt     time.Time
}




