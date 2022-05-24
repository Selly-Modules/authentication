package authentication

import (
	"github.com/Selly-Modules/natsio"
)

// Staff ...
type Staff struct {
	ID                   string   `json:"_id"` // ObjectID
	Name                 string   `json:"name"`
	Phone                string   `json:"phone"`
	Active               bool     `json:"active"`
	Role                 string   `json:"role,omitempty"`   // ObjectID
	Avatar               []byte   `json:"avatar,omitempty"` // FilePhoto
	CreatedAt            string   `json:"createdAt"`        // ISOString
	UpdatedAt            string   `json:"updatedAt"`        // ISOString
	IsRoot               bool     `json:"isRoot"`
	Permissions          []string `json:"permissions"`
	NotAllowedLoginAdmin bool     `json:"notAllowedLoginAdmin"`
	Source               []string `json:"source"`
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
	StaffID    string `json:"staffId"`
	Permission string `json:"permission"`
	Source     string `json:"source"`
	Code       string `json:"code,omitempty"`
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
