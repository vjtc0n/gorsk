package user

import (
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	"github.com/labstack/echo"
	"github.com/veep-provider/pkg/api/user/platform/pgsql"
	"github.com/veep-provider/pkg/utl/model"
)

// Service represents user application interface
type Service interface {
	Create(echo.Context, veep.User) (*veep.User, error)
	List(echo.Context, *veep.Pagination) ([]veep.User, error)
	View(echo.Context, int) (*veep.User, error)
	Delete(echo.Context, int) error
	Update(echo.Context, *Update) (*veep.User, error)
}

// New creates new user application service
func New(db *pg.DB, udb UDB, rbac RBAC, sec Securer) *User {
	return &User{db: db, udb: udb, rbac: rbac, sec: sec}
}

// Initialize initalizes User application service with defaults
func Initialize(db *pg.DB, rbac RBAC, sec Securer) *User {
	return New(db, pgsql.NewUser(), rbac, sec)
}

// User represents user application service
type User struct {
	db   *pg.DB
	udb  UDB
	rbac RBAC
	sec  Securer
}

// Securer represents security interface
type Securer interface {
	Hash(string) string
}

// UDB represents user repository interface
type UDB interface {
	Create(orm.DB, veep.User) (*veep.User, error)
	View(orm.DB, int) (*veep.User, error)
	List(orm.DB, *veep.ListQuery, *veep.Pagination) ([]veep.User, error)
	Update(orm.DB, *veep.User) error
	Delete(orm.DB, *veep.User) error
}

// RBAC represents role-based-access-control interface
type RBAC interface {
	User(echo.Context) *veep.AuthUser
	EnforceUser(echo.Context, int) error
	AccountCreate(echo.Context, veep.AccessRole, int, int) error
	IsLowerRole(echo.Context, veep.AccessRole) error
}
