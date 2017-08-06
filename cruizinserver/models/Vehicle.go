package models

type Vehicle struct {
	Vid   int    `json:"vid"`
	Year  int    `json:"year"`
	Make  string `json:"make"`
	Model string `json:"model"`
}
