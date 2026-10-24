package dto

// ObjectUpdate handles partial data changes
type ObjectUpdate struct {
	Name *string                `json:"name,omitempty"`
	Data map[string]interface{} `json:"data,omitempty"`
}
