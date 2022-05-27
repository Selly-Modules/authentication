package authentication

import (
	"github.com/Selly-Modules/natsio"
)

// Log ...
type Log struct {
	Reference  string                 `json:"reference"`
	Payload    PayloadCheckPermission `json:"payload"`
	Source     string                 `json:"source"`
	Permission []string               `json:"permission"`
	Action     string                 `json:"action"`
}

// Staff ...
type Staff struct {
	Reference   string   `json:"reference"` // ObjectID
	Name        string   `json:"name"`
	Email       string   `json:"email"`
	Phone       string   `json:"phone"`
	Active      bool     `json:"active"`
	IsRoot      bool     `json:"isRoot"`
	Permissions []string `json:"permissions"`
	Source      []string `json:"source"`
}

// ListStaff ...
type ListStaff struct {
	Staffs []Staff `json:"staffs"`
}

// Config int client elasticsearch
type Config struct {
	ApiKey string
	Nats   natsio.Config
}

// StaffCheckPermissionBody ...
type StaffCheckPermissionBody struct {
	StaffID    string   `json:"staffId"`
	Token      string   `json:"token"`
	Permission []string `json:"permission"`
	Action     string   `json:"action"`
	Source     string   `json:"source"`
	Code       string   `json:"code,omitempty"`
	IsRoot     string   `json:"isRoot"`
}

// PayloadCheckPermission ...
type PayloadCheckPermission struct {
	URL    string `json:"url"`
	Body   string `json:"body"`
	Method string `json:"method"`
}

// StaffCheckPermissionResponse ...
type StaffCheckPermissionResponse struct {
	Message string `json:"message"`
	IsValid bool   `json:"isValid"`
}

// StaffGetPermissions ...
type StaffGetPermissions struct {
	StaffID string `json:"staffId"`
	Source  string `json:"source"`
	Prefix  string `json:"prefix"`
}

// StaffGetPermissionsResponse ...
type StaffGetPermissionsResponse struct {
	Permissions []string `json:"permissions"`
}

// RequestBody ...
type RequestBody struct {
	Body   []byte `json:"body"`
	ApiKey string `json:"apiKey"`
}

// Response
// response to service es
type Response struct {
	Success bool   `json:"success"`
	Data    []byte `json:"data,omitempty"`
	Message string `json:"message"`
}
