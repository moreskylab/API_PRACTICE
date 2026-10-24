package dto

// ObjectCreate strictly requires a Name parameter
type ObjectCreate struct {
	Name string                 `json:"name" binding:"required"`
	Data map[string]interface{} `json:"data"`
}
