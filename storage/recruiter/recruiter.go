package recruiter

import (
	"time"

	"github.com/devcoalition/recruiter-dne/storage/rep"
)

// Recruiter is the domain type for a recruiter object, which contains a
// list of their reputation.
type Recruiter struct {
	ID       int       `json:"id"`
	Email    string    `json:"email"`
	Name     string    `json:"name"`
	Company  string    `json:"company,omitempty"`
	Reps     []rep.Rep `json:"reps,omitempty"`
	Score    int       `json:"score"`
	Created  time.Time `json:"created,omitempty"` // * Curious what happens if we just have omitempty here
	Updated  time.Time `json:"updated,omitempty"`
	Disabled bool      `json:"disabled,omitempty"`
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
	RetrieveRecruiter(Recruiter) (Recruiter, error)
	// * RetrieveRecruiters, limit 500 -- random sample?
}

type Upserter interface {
	UpsertRecruiter(Recruiter) (Recruiter, error)
}

type Deleter interface {
	DeleteRecruiter(Recruiter) error
}
