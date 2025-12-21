package content

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/RightHandSide/CVWO-App/internal/api"
	contentdata "github.com/RightHandSide/CVWO-App/internal/dataaccess/content"
	userdata "github.com/RightHandSide/CVWO-App/internal/dataaccess/users"
	"github.com/RightHandSide/CVWO-App/internal/database"
	"github.com/pkg/errors"
)

func HandleThread(w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	// Check for "username" Cookie
	c, err := r.Cookie("username")
	if err != nil || strings.TrimSpace(c.Value) == "" {
		// If Error, Show User "Not Logged In"
		w.WriteHeader(http.StatusUnauthorized)
		return &api.Response{
			Messages:  []string{"Not Logged In"},
			ErrorCode: 1,
		}, nil
	}

	// Get Database
	db, err := database.GetDB()
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrRetrieveDatabase, ListThreads))
	}

	// Get User By Name
	user, err := userdata.GetUserByName(db, c.Value)
	if err != nil || user == nil {
		// If No Such Name, Return "Invalid User"
		w.WriteHeader(http.StatusUnauthorized)
		return &api.Response{
			Messages:  []string{"Invalid User"},
			ErrorCode: 1,
		}, nil
	}

	// List Thread
	threads, err := contentdata.ListThread(db)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrRetrieveThreads, ListThreads))
	}

	// Encode threads
	data, err := json.Marshal(threads)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrEncodeView, ListThreads))
	}

	// Return api.Response: {Payload: Thread List}, {Message: Successful}
	return &api.Response{
		Payload: api.Payload{
			Data: data,
		},
		Messages: []string{SuccessfulListThreadsMessage},
	}, nil
}
