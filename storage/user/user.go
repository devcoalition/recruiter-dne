package user

import "time"

type User struct {
	ID      int
	Email   string
	Status  string
	Created time.Time
	Updated time.Time
}

type Storage interface {
	Creator
	Retriever
	Upserter
	Deleter
}

type Creator interface {
	CreateUser(User) (User, error)
}

type Retriever interface {
	RetrieverUser(User) (User, error)
}

type Upserter interface {
	UpsertUser(User) (User, error)
}

type Deleter interface {
	DeleteUser(User) error
}
