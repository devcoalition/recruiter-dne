package rep

import "time"

type Rep struct {
	ID          int
	UserID      int
	RecruiterID int
	Score       int
	EmailBlob   string // * Figure out right data structure for this
	Created     time.Time
	Updated     time.Time
	// * Could have this enum limited to several values, such as "spam email", "cold call", "got me a job", "great negotiator"
	// Reason      string
}

type Storage interface {
	Creator
	Retriever
	Upserter
	Deleter
}

type Creator interface {
	CreateRep(Rep) (Rep, error)
}

type Retriever interface {
	RetrieverRep(Rep) (Rep, error)
}

type Upserter interface {
	UpsertRep(Rep) (Rep, error)
}

type Deleter interface {
	DeleteRep(Rep) error
}
