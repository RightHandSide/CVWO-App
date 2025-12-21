package users

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/RightHandSide/CVWO-App/internal/api"
	userdata "github.com/RightHandSide/CVWO-App/internal/dataaccess/users"
	"github.com/RightHandSide/CVWO-App/internal/database"
	"github.com/pkg/errors"
)

func HandleLogin(w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	// Set Message as Successful
	msg := SuccessfulLoginUsersMessage
	code := 0

	// Decode Request Body as api.LoginRequest
	var req api.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		// If Error, Show User "Bad Request"
		w.WriteHeader(http.StatusBadRequest)
		return &api.Response{
			Messages:  []string{"Bad Request"},
			ErrorCode: 1,
		}, nil
	}
	var name string = req.Name

	// Get Database
	db, err := database.GetDB()
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrRetrieveDatabase, LoginUsers))
	}

	// Get User By Name
	user, err := userdata.GetUserByName(db, name)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrLoginUser, LoginUsers))
	}

	// When No Such User, user == nil
	if user == nil {
		msg = PleaseRegister
		code = 1
	} else {
		// Set Cookie {username = user.Name} for Identification
		http.SetCookie(w, &http.Cookie{
			Name:     "username",
			Value:    user.Name,
			Path:     "/",
			HttpOnly: true,
		})
	}

	// Return api.Response: {Message: Paste If Error}, {ErrorCode: Error If Not 0}
	return &api.Response{
		Messages:  []string{msg},
		ErrorCode: code,
	}, nil
}
