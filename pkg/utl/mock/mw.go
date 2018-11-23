package mock

import (
	"github.com/veep-provider/pkg/utl/model"
)

// JWT mock
type JWT struct {
	GenerateTokenFn func(*veep.User) (string, string, error)
}

// GenerateToken mock
func (j *JWT) GenerateToken(u *veep.User) (string, string, error) {
	return j.GenerateTokenFn(u)
}
