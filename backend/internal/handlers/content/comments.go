package content

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/RightHandSide/CVWO-App/internal/api"
	contentdata "github.com/RightHandSide/CVWO-App/internal/dataaccess/content"
	userdata "github.com/RightHandSide/CVWO-App/internal/dataaccess/users"
	"github.com/RightHandSide/CVWO-App/internal/database"
	"github.com/go-chi/chi/v5"
	"github.com/pkg/errors"
)

func HandleComment(w http.ResponseWriter, r *http.Request) (*api.Response, error) {
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
		return nil, errors.Wrap(err, fmt.Sprintf(ErrRetrieveDatabase, ListComments))
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

	// Post ID
	strid := chi.URLParam(r, "id")
	intid, err := strconv.Atoi(strid)
	if err != nil {
		// If Not ID, Return "Invalid URL"
		w.WriteHeader(http.StatusUnauthorized)
		return &api.Response{
			Messages:  []string{"Invalid URL"},
			ErrorCode: 1,
		}, nil
	}

	// List Comment
	post, comments, err := contentdata.ListCommentByID(db, intid)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrRetrieveComments, ListComments))
	}

	data := map[string]interface{}{
		"head": post,
		"tail": comments,
	}

	// Encode
	encoded_data, err := json.Marshal(data)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrEncodeView, ListComments))
	}

	// Return api.Response: {Payload: [Post, Comment List]}, {Message: Successful}
	return &api.Response{
		Payload: api.Payload{
			Data: encoded_data,
		},
		Messages: []string{SuccessfulListCommentsMessage},
	}, nil
}

func CreateComment(w http.ResponseWriter, r *http.Request) (*api.Response, error) {
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
		return nil, errors.Wrap(err, fmt.Sprintf(ErrRetrieveDatabase, CreateComments))
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

	// Post ID
	strid := chi.URLParam(r, "id")
	intid, err := strconv.Atoi(strid)
	if err != nil {
		// If Not ID, Return "Invalid URL"
		w.WriteHeader(http.StatusUnauthorized)
		return &api.Response{
			Messages:  []string{"Invalid URL"},
			ErrorCode: 1,
		}, nil
	}

	// Create Comment
	if err := contentdata.MakeComment(db, user, r, intid); err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrCreateComments, CreateComments))
	}

	// Return api.Response: {Message: Paste If Error}, {ErrorCode: Error If Not 0}
	return &api.Response{
		Messages:  []string{SuccessfulCreateCommentsMessage},
		ErrorCode: 0,
	}, nil
}
