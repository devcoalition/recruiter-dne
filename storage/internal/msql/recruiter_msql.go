package msql

import (
	"database/sql"
	"errors"
	"math/rand"

	"github.com/devcoalition/recruiter-dne/storage/recruiter"
)

// recruiterSQL defines a MySQL implementation of recruiter.Storage.
type recruiterSQL struct {
	master *sql.DB
	slaves []*sql.DB
}

// NewRecruiterSQL constructs a new recruiterSQL struct.
func NewRecruiterSQL(master *sql.DB, slaves ...*sql.DB) recruiter.Storage {
	return recruiterSQL{
		master: master,
		slaves: slaves,
	}
}

// getRecruiterSQLSlaveDB is used to get a Slave Database for running read only
// queries. This function returns the master when no slaves have been declared.
func (rsql recruiterSQL) getRecruiterSQLSlaveDB() *sql.DB {
	if len(rsql.slaves) > 0 {
		return rsql.slaves[rand.Intn(len(rsql.slaves))]
	}
	return rsql.master
}

// CreateRecruiter creates a new Recruiter in the database.
func (rsql recruiterSQL) CreateRecruiter(r recruiter.Recruiter) (recruiter.Recruiter, error) {
	db := rsql.master

	query := `
		INSERT INTO recruiters (email, name)
		VALUES (?, ?)`
	res, err := db.Exec(query, r.Email, r.Name)
	if err != nil {
		return recruiter.Recruiter{}, err
	}

	// Note: r.Created / r.Updated are not being populated here.
	r.ID, err = res.LastInsertId()
	if err != nil {
		return recruiter.Recruiter{}, err
	}
	return r, nil
}

// RetrieveRecruiter retrieves a Recruiter from the database.
func (rsql recruiterSQL) RetrieveRecruiter(r recruiter.Recruiter) (recruiter.Recruiter, error) {
	db := rsql.getRecruiterSQLSlaveDB()

	var err error
	var rr recruiter.Recruiter

	// * Needs to be updated to pull reps
	query := `
		SELECT id, email, name, created, updated
		FROM recruiters`
	if r.ID != 0 {
		query += `
		WHERE id = ?`
		err = db.QueryRow(query, r.ID).Scan(&rr.ID, &rr.Email, &rr.Name, &rr.Created, &rr.Updated)
	} else if r.Email != "" {
		query += `
		WHERE email = ?`
		err = db.QueryRow(query, r.Email).Scan(&rr.ID, &rr.Email, &rr.Name, &rr.Created, &rr.Updated)
	} else {
		return recruiter.Recruiter{}, errors.New("Neither r.ID nor r.Email was provided")
	}
	if err != nil {
		return recruiter.Recruiter{}, err
	}

	return rr, nil
}

// UpsertRecruiters upserts a Recruiter into the database.
func (rsql recruiterSQL) UpsertRecruiter(r recruiter.Recruiter) (recruiter.Recruiter, error) {
	db := rsql.master

	query := `
		INSERT INTO recruiters (id, email, name)
		VALUES (?, ?, ?)
		ON DUPLICATE KEY
		UPDATE email = VALUES(email), name = VALUES(name)`

	res, err := db.Exec(query, r.ID, r.Email, r.Name)
	if err != nil {
		return recruiter.Recruiter{}, err
	}

	// Note: r.Created / r.Updated are not being populated here.
	if res.RowsAffected() == 1 {
		r.ID = res.LastInsertId()
	}
	return r, nil
}

// DeleteRecruiter deletes a Recruiter from the database.
func (rsql recruiterSQL) DeleteRecruiter(r recruiter.Recruiter) error {
	db := rsql.master

	query := `
		DELETE FROM recruiters
		WHERE id = ?`

	_, err := db.Exec(query, r.ID)
	if err != nil {
		return recruiter.Recruiter{}, err
	}
	return nil
}
