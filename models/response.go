package models

type Response struct {
	Data  interface{} `json:"data,omitempty"`
	Error interface{} `json:"error,omitempty"`
}
