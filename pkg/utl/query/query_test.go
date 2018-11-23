package query_test

import (
	"testing"

	"github.com/veep-provider/pkg/utl/model"

	"github.com/labstack/echo"

	"github.com/veep-provider/pkg/utl/query"
	"github.com/stretchr/testify/assert"
)

func TestList(t *testing.T) {
	type args struct {
		user *veep.AuthUser
	}
	cases := []struct {
		name     string
		args     args
		wantData *veep.ListQuery
		wantErr  error
	}{
		{
			name: "Super admin user",
			args: args{user: &veep.AuthUser{
				Role: veep.SuperAdminRole,
			}},
		},
		{
			name: "Company admin user",
			args: args{user: &veep.AuthUser{
				Role:      veep.CompanyAdminRole,
				CompanyID: 1,
			}},
			wantData: &veep.ListQuery{
				Query: "company_id = ?",
				ID:    1},
		},
		{
			name: "Location admin user",
			args: args{user: &veep.AuthUser{
				Role:       veep.LocationAdminRole,
				CompanyID:  1,
				LocationID: 2,
			}},
			wantData: &veep.ListQuery{
				Query: "location_id = ?",
				ID:    2},
		},
		{
			name: "Normal user",
			args: args{user: &veep.AuthUser{
				Role: veep.UserRole,
			}},
			wantErr: echo.ErrForbidden,
		},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			q, err := query.List(tt.args.user)
			assert.Equal(t, tt.wantData, q)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}
