package model

import "time"

type Person struct {
	ID             uint      `gorm:"primaryKey" json:"id"`
	FirstName      string    `json:"first_name"`
	LastName       string    `json:"last_name"`
	Email          string    `json:"email"`
	Gender         string    `json:"gender"`
	DateOfBirth    time.Time `json:"date_of_birth"`
	CountryOfBirth string    `json:"country_of_birth"`
}

func (Person) TableName() string {
	return "person"
}
