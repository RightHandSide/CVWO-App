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

func HandleUser(w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	// Get Database
	db, err := database.GetDB()
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrRetrieveDatabase, ListUsers))
	}

	// List User
	users, err := userdata.ListUser(db)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrRetrieveUsers, ListUsers))
	}

	// Encode users
	data, err := json.Marshal(users)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrEncodeView, ListUsers))
	}

	// Return api.Response: {Payload: User List}, {Message: Successful}
	return &api.Response{
		Payload: api.Payload{
			Data: data,
		},
		Messages: []string{SuccessfulListUsersMessage},
	}, nil
}
