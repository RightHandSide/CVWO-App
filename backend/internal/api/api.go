package api

import "encoding/json"

type LoginRequest struct {
	Name string `json:"name"`
}

type CreateRequest struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

type Payload struct {
	Meta json.RawMessage `json:"meta,omitempty"`
	Data json.RawMessage `json:"data,omitempty"`
}

type Response struct {
	Payload   Payload  `json:"payload"`
	Messages  []string `json:"messages"`
	ErrorCode int      `json:"errorCode"`
}

// HTTP Data Sent
