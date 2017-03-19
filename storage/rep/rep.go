package rep

import "time"

type Rep struct {
	UserID      int
	RecruiterID int
	Value       bool
	Date        time.Time
	EmailGob    string // * Figure out right data structure for this
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
