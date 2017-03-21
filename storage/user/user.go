package user

import (
	"errors"
	"strings"
	"time"
)

// User is the domain type for a recipient-of-spam's e-mail address.
type User struct {
	ID      int
	Email   string
	Status  Status
	Created time.Time
	Updated time.Time
}

// Storage represents the suite of interfaces by which a User can be manipulated.
// A new storage implementation must support all of these interfaces.
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

// Status defines a named type for a User's Status, ie: their interest in receiving
// e-mails from recruiters.
type Status int

const (
	NoEmail Status = iota
	SomeEmail
	AllEmail
)

var ErrInvalidStatus = errors.New("Invalid user.Status")

// String returns a string representation of a Status.
func (s Status) String() string {
	switch s {
	case NoEmail:
		return "no_email"
	case SomeEmail:
		return "some_email"
	case AllEmail:
		return "all_email"
	}
	return ""
}

// StatusType converts a string to a user.Status.
func StatusType(s string) (Status, error) {
	switch strings.ToLower(s) {
	case "no_email":
		return NoEmail, nil
	case "some_email":
		return SomeEmail, nil
	case "all_email":
		return AllEmail, nil
	default:
		return "", ErrInvalidStatus
	}
}
