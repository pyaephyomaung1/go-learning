package dto

import "time"

type CreatePersonRequest struct {
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	Email          string `json:"email"`
	Gender         string `json:"gender"`
	DateOfBirth    string `json:"date_of_birth"`
	CountryOfBirth string `json:"country_of_birth"`
}

type UpdatePersonRequest struct {
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	Email          string `json:"email"`
	Gender         string `json:"gender"`
	DateOfBirth    string `json:"date_of_birth"`
	CountryOfBirth string `json:"country_of_birth"`
}

type PersonResponse struct {
	ID             uint      `json:"id"`
	FirstName      string    `json:"first_name"`
	LastName       string    `json:"last_name"`
	Email          string    `json:"email"`
	Gender         string    `json:"gender"`
	DateOfBirth    time.Time `json:"date_of_birth"`
	CountryOfBirth string    `json:"country_of_birth"`
}
