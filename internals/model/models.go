package model

type Category struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type CreateCategory struct {
	Name string `json:"name"`
}
