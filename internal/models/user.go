package models

import "fmt"

// Structure
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Thread struct {
	ID          int    `json:"id"`
	User_ID     int    `json:"user_id"`
	Title       string `json:"title"`
	Description string `json:"desc"`
}

type Post struct {
	ID        int    `json:"id"`
	Thread_ID int    `json:"thread_id"`
	Title     string `json:"title"`
	Body      string `json:"body"`
}

type Comment struct {
	ID      int    `json:"id"`
	Post_ID int    `json:"post_id"`
	User_ID int    `json:"user_id"`
	Body    string `json:"body"`
}

func (user *User) Greet() string {
	return fmt.Sprintf("Hello, I am %s", user.Name)
}
