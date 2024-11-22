package main

// Define a new custom type (Explicit Type Declaration), type name, struct is the type.
type Article struct {
	ID        int    `db:"id" json:"id"`
	Title     string `db:"title" json:"title"`
	Content   string `db:"content" json:"content"`
	Author    string `db:"author" json:"author"`
	CreatedAt string `db:"created_at" json:"created_at"`
}
