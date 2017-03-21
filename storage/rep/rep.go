package rep

import "time"

type Rep struct {
	ID          int
	UserID      int
	RecruiterID int
	Score       int
	Reason      string // * Probably won't to make this an enum once you've thought of all possible values on the frontend
	EmailBlob   string // * Figure out right data structure for this
	Created     time.Time
	Updated     time.Time
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
