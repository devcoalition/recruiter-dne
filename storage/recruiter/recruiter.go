package recruiter

import (
	"time"

	"github.com/devcoalition/recruiter-dne/storage/rep"
)

type Recruiter struct {
	ID      int
	Email   string
	Name    string
	Reps    []rep.Rep
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
	CreateRecruiter(Recruiter) (Recruiter, error)
}

type Retriever interface {
	RetrieverRecruiter(Recruiter) (Recruiter, error)
}

type Upserter interface {
	UpsertRecruiter(Recruiter) (Recruiter, error)
}

type Deleter interface {
	DeleteRecruiter(Recruiter) error
}
