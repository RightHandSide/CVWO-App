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

func HandlePost(w http.ResponseWriter, r *http.Request) (*api.Response, error) {
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
		return nil, errors.Wrap(err, fmt.Sprintf(ErrRetrieveDatabase, ListPosts))
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

	// Thread ID
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

	// List Post
	thread, posts, err := contentdata.ListPostByID(db, intid)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrRetrievePosts, ListPosts))
	}

	data := map[string]interface{}{
		"head": thread,
		"tail": posts,
	}

	// Encode
	encoded_data, err := json.Marshal(data)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrEncodeView, ListPosts))
	}

	// Return api.Response: {Payload: [Thread, Post List]}, {Message: Successful}
	return &api.Response{
		Payload: api.Payload{
			Data: encoded_data,
		},
		Messages: []string{SuccessfulListPostsMessage},
	}, nil
}

func CreatePost(w http.ResponseWriter, r *http.Request) (*api.Response, error) {
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
		return nil, errors.Wrap(err, fmt.Sprintf(ErrRetrieveDatabase, CreatePosts))
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

	// Thread ID
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

	// Create Post
	if err := contentdata.MakePost(db, user, r, intid); err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrCreatePosts, CreatePosts))
	}

	// Return api.Response: {Message: Paste If Error}, {ErrorCode: Error If Not 0}
	return &api.Response{
		Messages:  []string{SuccessfulCreatePostsMessage},
		ErrorCode: 0,
	}, nil
}
