package models

type User struct {
	id   string
	name string
	age  int
	mail string
}

type DB struct {
	data map[string]*User
}
