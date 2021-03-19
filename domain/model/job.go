package model

// Job struct
type Job struct {
	ID              int    `gorm:"primary_key" json:"ID"`
	Title           string `json:"Title"`
	NormalizedTitle string `json:"NormalizedTitle"`
}

// ExtJob struct
type ExtJob struct {
	UUID               string `json:"uuid"`
	Title              string `json:"title"`
	NormalizedJobTitle string `json:"normalized_job_title"`
	ParentUUID         string `json:"parent_uuid"`
}

// APIResult struct
type APIResult struct {
	Results []ExtJob
}

// TableName function
func (Job) TableName() string { return "jobs" }
