package pgsql

import (
	"github.com/go-pg/pg/orm"
	"github.com/veep-provider/pkg/utl/model"
)

// NewUser returns a new user database instance
func NewUser() *User {
	return &User{}
}

// User represents the client for user table
type User struct{}

// View returns single user by ID
func (u *User) View(db orm.DB, id int) (*veep.User, error) {
	user := &veep.User{Base: veep.Base{ID: id}}
	err := db.Select(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// Update updates user's info
func (u *User) Update(db orm.DB, user *veep.User) error {
	return db.Update(user)
}
