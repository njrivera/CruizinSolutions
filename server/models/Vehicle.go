package models

type Vehicle struct {
	Vid   int    `json:"vid"`
	Year  string `json:"year"`
	Make  string `json:"make"`
	Model string `json:"model"`
}
