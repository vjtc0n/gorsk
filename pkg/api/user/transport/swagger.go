package transport

import (
	"github.com/veep-provider/pkg/utl/model"
)

// User model response
// swagger:response userResp
type swaggUserResponse struct {
	// in:body
	Body struct {
		*veep.User
	}
}

// Users model response
// swagger:response userListResp
type swaggUserListResponse struct {
	// in:body
	Body struct {
		Users []veep.User `json:"users"`
		Page  int          `json:"page"`
	}
}
