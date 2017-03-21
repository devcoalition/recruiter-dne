package recruiter

import (
	"time"

	"github.com/devcoalition/recruiter-dne/storage/rep"
)

// Recruiter is the domain type for a recruiter object, which contains a
// list of their reputation.
type Recruiter struct {
	ID      int
	Email   string
	Name    string
	Reps    []rep.Rep
	Created time.Time
	Updated time.Time
}

// Storage represents the suite of interfaces by which a Recruiter can be
// manipulated. A new storage implementation must support all of these
// interfaces.
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
