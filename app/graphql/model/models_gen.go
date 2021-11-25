// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type NewTodo struct {
	Title   string `json:"title"`
	Comment string `json:"comment"`
	UserID  string `json:"userId"`
}

type NewUser struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Todo struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Comment string `json:"comment"`
	User    *User  `json:"user"`
}

type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}