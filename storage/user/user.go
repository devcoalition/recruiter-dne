package user

import "github.com/devcoalition/recruiter-dne/storage/rep"

type User struct {
	ID     int
	Email  string
	Status string
	Reps   []rep.Rep
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
