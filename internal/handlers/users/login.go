package users

import (
	"fmt"
	"net/http"

	"github.com/RightHandSide/CVWO-App/internal/api"
	"github.com/RightHandSide/CVWO-App/internal/dataaccess"
	"github.com/RightHandSide/CVWO-App/internal/database"
	"github.com/pkg/errors"
)

func HandleLogin(w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	msg := SuccessfulLoginUsersMessage
	code := 0

	db, err := database.GetDB()
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrRetrieveDatabase, LoginUsers))
	}

	user, err := dataaccess.LoginUser(db, r)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrLoginUser, LoginUsers))
	}

	if user == nil {
		msg = PleaseRegister
		code = 1
	}

	return &api.Response{
		Messages:  []string{msg},
		ErrorCode: code,
	}, nil
}
