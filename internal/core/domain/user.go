package domain

import (
	"time"
)

type User struct {
	ID			string	
	Name		string			`bson:"name"`
	Email 		string			`bson:"email"`
	Password	string			`bson:"password"`
	CreatedAt	time.Time		`bson:"created_at"`
}