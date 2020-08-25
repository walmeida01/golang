package model

//Usuario representa um usuário no sistema
type Usuario struct {
	ID   int    `json:"id"`
	Nome string `json:"nome"`
}
