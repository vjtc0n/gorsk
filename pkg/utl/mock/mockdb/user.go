package mockdb

import (
	"github.com/go-pg/pg/orm"
	"github.com/veep-provider/pkg/utl/model"
)

// User database mock
type User struct {
	CreateFn         func(orm.DB, veep.User) (*veep.User, error)
	ViewFn           func(orm.DB, int) (*veep.User, error)
	FindByUsernameFn func(orm.DB, string) (*veep.User, error)
	FindByTokenFn    func(orm.DB, string) (*veep.User, error)
	ListFn           func(orm.DB, *veep.ListQuery, *veep.Pagination) ([]veep.User, error)
	DeleteFn         func(orm.DB, *veep.User) error
	UpdateFn         func(orm.DB, *veep.User) error
}

// Create mock
func (u *User) Create(db orm.DB, usr veep.User) (*veep.User, error) {
	return u.CreateFn(db, usr)
}

// View mock
func (u *User) View(db orm.DB, id int) (*veep.User, error) {
	return u.ViewFn(db, id)
}

// FindByUsername mock
func (u *User) FindByUsername(db orm.DB, uname string) (*veep.User, error) {
	return u.FindByUsernameFn(db, uname)
}

// FindByToken mock
func (u *User) FindByToken(db orm.DB, token string) (*veep.User, error) {
	return u.FindByTokenFn(db, token)
}

// List mock
func (u *User) List(db orm.DB, lq *veep.ListQuery, p *veep.Pagination) ([]veep.User, error) {
	return u.ListFn(db, lq, p)
}

// Delete mock
func (u *User) Delete(db orm.DB, usr *veep.User) error {
	return u.DeleteFn(db, usr)
}

// Update mock
func (u *User) Update(db orm.DB, usr *veep.User) error {
	return u.UpdateFn(db, usr)
}
