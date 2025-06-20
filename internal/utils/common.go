package utils

import (
	"encoding/json"
	"net/http"
)

// WriteJSONResponse writes a standard JSON response with a single resource.
//
//	{
//	 "status": "success",
//	 "data": { ... },       // single resource or object
//	 "message": "Optional success message"
//	}

func WriteJSONResponse(w http.ResponseWriter, status int, data interface{}, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	resp := map[string]interface{}{
		"status": status,
		"data":   data,
	}

	if message != "" {
		resp["message"] = message
	}

	err := json.NewEncoder(w).Encode(resp)
	if err != nil {
		return
	}
}


// WriteJSONResponseList writes a JSON response with a list of resources.
//
//	{
//	 "status": "success",
//	 "data": [
//	   { ... },             // resource item 1
//	   { ... }              // resource item 2
//	 ],
//	 "message": "Optional success message"
//	}

func WriteJSONResponseList(w http.ResponseWriter, status int, data []interface{}, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	resp := map[string]interface{}{
		"status" : "success",
		"data" : data,
	}
	if message != "" {
		resp["message"] = message
	}

	err := json.NewEncoder(w).Encode(resp)
	if err != nil {
		return
	}
}

// Pagination metadata
type Pagination struct{
	Page int `json:"page"`
	PageSize int `json:"page_size"`
	TotalPages int `json:"total_pages"`
	TotalItems int `json:"total_items"`
}

// WriteJSONPaginatedResponse writes a paginated JSON response.
//
//	{
//	 "status": "success",
//	 "data": [
//	   { ... },             // resource item 1
//	   { ... }              // resource item 2
//	 ],
//	 "pagination": {
//	   "page": 1,
//	   "page_size": 20,
//	   "total_pages": 5,
//	   "total_items": 100
//	 },
//	 "message": "Optional success message"
//	}

func WriteJSONPaginatedResponse(w http.ResponseWriter, status int, data []interface{}, pagination Pagination, message string){
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	resp := map[string]interface{}{
		"status":     "success",
		"data":       data,
		"pagination": pagination,
	}
	if message != "" {
		resp["message"] = message
	}

	err := json.NewEncoder(w).Encode(resp)
	if err != nil {
		return
	}
}

// WriteErrorResponse writes a standard JSON error response.
//
//	{
//	 "status": "error",
//	 "error": {
//	   "code": 400,
//	   "message": "Invalid request parameters",
//	   "details": "Optional detailed error info"
//	 }
//	}
func WriteErrorResponse(w http.ResponseWriter, status int, errMsg string, details string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	resp := map[string]interface{}{
		"status": "error",
		"error": map[string]interface{}{
			"code":    status,
			"message": errMsg,
		},
	}
	if details != "" {
		resp["error"].(map[string]interface{})["details"] = details
	}

	err := json.NewEncoder(w).Encode(resp)
	if err != nil {
		return
	}
}

// WriteErrorResponseList writes a JSON error response with a list of validation errors.
//
//	{
//	 "status": "error",
//	 "message": "Validation failed",
//	 "code": 400,
//	 "errors": [
//	   {
//	     "field": "title",
//	     "message": "Title is required"
//	   },
//	   {
//	     "field": "email",
//	     "message": "Email must be a valid email"
//	   }
//	 ]
//	}
func WriteErrorResponseList(w http.ResponseWriter, status int, message string, errors interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	resp := map[string]interface{}{
		"status":  "error",
		"message": message,
		"code":    status,
		"errors":  errors,
	}

	err := json.NewEncoder(w).Encode(resp)
	if err != nil {
		return
	}
}

