package model

// ResponseObject maps out your database-like core model
type ResponseObject struct {
	ID   string                 `json:"id"`
	Name string                 `json:"name"`
	Data map[string]interface{} `json:"data"` // Handles Python's dict[str, Any]
}
