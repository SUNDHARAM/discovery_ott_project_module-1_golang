package models

type User struct {
	UserId    int    `gorm:"primaryKey;autoIncrement:true" json:"user_id"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Gender    string `json:"gender"`
	TimeStamp string `json:"time_stamp"`
	IsActive  bool   `json:"is_active"`
	IsAdmin   bool   `json:"is_admin"`
	IsCms     bool   `json:"is_cms"`
	IsSms     bool   `json:"is_sms"`
}
