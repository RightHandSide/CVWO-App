package users

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/RightHandSide/CVWO-App/internal/api"
	"github.com/RightHandSide/CVWO-App/internal/dataaccess"
	"github.com/RightHandSide/CVWO-App/internal/database"
	"github.com/pkg/errors"
)

func HandleRegister(w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	msg := SuccessfulRegisterUsersMessage
	code := 0

	db, err := database.GetDB()
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrRetrieveDatabase, RegisterUsers))
	}

	name, err := dataaccess.RegisterUser(db, r)
	if err != nil {
		if strings.Contains(err.Error(), "UNIQUE constraint failed") {
			msg = NameRepeated
			code = 1
		} else {
			return nil, errors.Wrap(err, fmt.Sprintf(ErrRegisterUsers, RegisterUsers))
		}
	}
	_ = name

	return &api.Response{
		Messages:  []string{msg},
		ErrorCode: code,
	}, nil
}
