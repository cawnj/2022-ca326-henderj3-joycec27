package models

type User struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	PhoneNumber   string `json:"phone_number"`
	CovidPositive bool   `json:"covid_positive"`
}
