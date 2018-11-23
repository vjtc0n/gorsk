package query

import (
	"github.com/labstack/echo"
	"github.com/veep-provider/pkg/utl/model"
)

// List prepares data for list queries
func List(u *veep.AuthUser) (*veep.ListQuery, error) {
	switch true {
	case u.Role <= veep.AdminRole: // user is SuperAdmin or Admin
		return nil, nil
	case u.Role == veep.CompanyAdminRole:
		return &veep.ListQuery{Query: "company_id = ?", ID: u.CompanyID}, nil
	case u.Role == veep.LocationAdminRole:
		return &veep.ListQuery{Query: "location_id = ?", ID: u.LocationID}, nil
	default:
		return nil, echo.ErrForbidden
	}
}
