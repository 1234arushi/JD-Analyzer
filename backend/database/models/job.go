package models

type Job struct {
	JobID    uint64 `gorm:"primaryKey;column:job_id"`
	Company  string `gorm:"column:company"`
	Role     string `gorm:"column:role"`
	Location string `gorm:"column:location"`
	RawJD    string `gorm:"type:text;column:raw_jd"` //'text' means explicitly telling gorm that don't take default space as varchar(255)
}

func (Job) TableName() string {
	return "jobs"
}
