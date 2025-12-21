package users

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/RightHandSide/CVWO-App/internal/api"
	userdata "github.com/RightHandSide/CVWO-App/internal/dataaccess/users"
	"github.com/RightHandSide/CVWO-App/internal/database"
	"github.com/pkg/errors"
)

func HandleRegister(w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	// Set Message as Successful
	msg := SuccessfulRegisterUsersMessage
	code := 0

	// Get Database
	db, err := database.GetDB()
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrRetrieveDatabase, RegisterUsers))
	}

	// Register User
	name, err := userdata.RegisterUser(db, r)
	if err != nil {
		// Name Repeated (Name Column Unique), Set Message as NameRepeated
		if strings.Contains(err.Error(), "UNIQUE constraint failed") {
			msg = NameRepeated
			code = 1
		} else {
			return nil, errors.Wrap(err, fmt.Sprintf(ErrRegisterUsers, RegisterUsers))
		}
	}
	_ = name

	// Return api.Response: {Message: Paste If Error}, {ErrorCode: Error If Not 0}
	return &api.Response{
		Messages:  []string{msg},
		ErrorCode: code,
	}, nil
}
