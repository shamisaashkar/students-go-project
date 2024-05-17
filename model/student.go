package model

type Student struct {
	ID        int64   `bson:"id"`
	FirstName string  `bson:"first_name"`
	LastName  string  `bson:"last_name"`
	Average   float64 `bson:"average"`
}
