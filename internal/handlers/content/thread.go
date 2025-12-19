package content

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/RightHandSide/CVWO-App/internal/api"
	"github.com/RightHandSide/CVWO-App/internal/dataaccess"
	"github.com/RightHandSide/CVWO-App/internal/database"
	"github.com/pkg/errors"
)

func HandleThread(w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	db, err := database.GetDB()
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrRetrieveDatabase, ListThreads))
	}

	threads, err := dataaccess.ListThread(db)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrRetrieveThreads, ListThreads))
	}

	data, err := json.Marshal(threads)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrEncodeView, ListThreads))
	}

	return &api.Response{
		Payload: api.Payload{
			Data: data,
		},
		Messages: []string{SuccessfulListThreadsMessage},
	}, nil
}

// Print out all threads
// Post would be /thread-id/
// Singular Post would be /thread-id/post-id
// Remember to do something about ErrorCode
