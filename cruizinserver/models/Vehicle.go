package models

type Vehicle struct {
	Id    int    `json:"id"`
	Year  int    `json:"year"`
	Make  string `json:"make"`
	Model string `json:"model"`
	Trim  string `json:"trim"`
}
