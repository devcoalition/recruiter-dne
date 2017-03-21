package msql

import (
	"database/sql"
	"errors"
	"math/rand"

	"github.com/devcoalition/recruiter-dne/storage/user"
)

// userSQL defines a MySQL implementation of user.Storage.
type userSQL struct {
	master *sql.DB
	slaves []*sql.DB
}

// NewUserSQL constructs a new userSQL struct.
func NewUserSQL(master *sql.DB, slaves ...*sql.DB) user.Storage {
	return userSQL{
		master: master,
		slaves: slaves,
	}
}

// getUserSQLSlaveDB is used to get a Slave Database for running read only queries.
// This function will return a Master DB only if there are no SlaveDB's declared
// in the userSQL struct.
func (usql userSQL) getUserSQLSlaveDB() *sql.DB {
	if len(usql.slaves) > 0 {
		return usql.slaves[rand.Intn(len(usql.slaves))]
	}
	return usql.master
}

// CreateUser creates a new User in the database.
func (usql userSQL) CreateUser(u user.User) (user.User, error) {
	db := usql.master

	query := `
		INSERT INTO users (email, status)
		VALUES (?, ?)`
	res, err := db.Exec(query, u.Email, u.Status.String())
	if err != nil {
		return user.User{}, err
	}

	// Note: u.Created / u.Updated are not being populated at this time.
	u.ID, err = res.LastInsertId()
	if err != nil {
		return user.User{}, err
	}
	return u, nil
}

// RetrieveUser retrieves a User from the database.
func (usql userSQL) RetrieveUser(u user.User) (user.User, error) {
	db := usql.getUserSQLSlaveDB()

	var err error
	var r user.User
	var status string

	query := `
		SELECT id, email, status, created, updated
		FROM users`
	if u.ID != 0 {
		query += `
		WHERE id = ?`
		err = db.QueryRow(query, u.ID).Scan(&r.ID, &r.Email, &status, &r.Created, &r.Updated)
	} else if u.Email != "" {
		query += `
		WHERE email = ?`
		err = db.QueryRow(query, u.Email).Scan(&r.ID, &r.Email, &status, &r.Created, &r.Updated)
	} else {
		return user.User{}, errors.New("Neither u.ID nor u.Email was provided")
	}
	if err != nil {
		return user.User{}, err
	}

	r.Status, err = user.StatusType(status)
	if err != nil {
		return user.User{}, err
	}

	return r, nil
}

// UpsertUser upserts a User into the database.
func (usql userSQL) UpsertUser(u user.User) (user.User, error) {
	db := usql.master

	query := `
		INSERT INTO users (id, email, status)
		VALUES (?, ?, ?)
		ON DUPLICATE KEY
		UPDATE email = VALUES(email), status = VALUES(status)`

	res, err := db.Exec(query, u.ID, u.Email, u.Status.String())
	if err != nil {
		return user.User{}, err
	}

	if res.RowsAffected() == 1 {
		u.ID = res.LastInsertId()
	}
	return u, nil
}

// DeleteUser deletes a User from the database.
func (usql userSQL) DeleteUser(u user.User) error {
	db := usql.master

	query := `
		DELETE FROM users
		WHERE id = ?`

	_, err := db.Exec(query, u.ID, u.Email, u.Status.String())
	if err != nil {
		return user.User{}, err
	}
	return nil
}
