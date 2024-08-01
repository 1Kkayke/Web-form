package controller

import "MyPackages/db"

type Subscription struct {
	name  string
	email string
}

func Create(name string, email string) error {
	s := Subscription{name: name, email: email}

	return db.Insert("NewsLetter", s)
}
