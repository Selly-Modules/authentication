package authentication

import (
	"github.com/Selly-Modules/natsio"
	"go.mongodb.org/mongo-driver/bson"
)

// StaffVerifyCodeBody ...
type StaffVerifyCodeBody struct {
	Code    string `json:"code"`
	StaffID string `json:"staffId"`
	Source  string `json:"source"`
}

// StaffVerifyCodeResponse ...
type StaffVerifyCodeResponse struct {
	Success bool `json:"success"`
}

// CommonQuery ...
type CommonQuery struct {
	Page    int64  `json:"page"`
	Limit   int64  `json:"limit"`
	Keyword string `json:"keyword"`
	Sort    bson.D `json:"sort"`
	StaffID string `json:"staffID"`
}

// GetTokenByPhoneBody ...
type GetTokenByPhoneBody struct {
	Phone string `json:"phone"`
}

// GetTokenByPhoneResponse ...
type GetTokenByPhoneResponse struct {
	Token string `json:"token"`
}

// GetInfoStaff
// Convert to bson bytes
// When partner use. Bson unmarshal convert to struct
type GetInfoStaff struct {
	Condition interface{}
	Page      int64
	Limit     int64
	Sort      bson.D
}

// Log ...
type Log struct {
	Reference  string                 `json:"reference"`
	Payload    PayloadCheckPermission `json:"payload"`
	Source     string                 `json:"source"`
	Permission []string               `json:"permission"`
	Action     string                 `json:"action"`
}

// LogsResponse ...
type LogsResponse struct {
	ID        string                 `json:"_id"`
	Staff     string                 `json:"staff"`
	Payload   PayloadCheckPermission `json:"payload"`
	Action    string                 `json:"action"`
	CreatedAt string                 `json:"createdAt"`
}

// ListPageResponse ...
type ListPageResponse struct {
	Data  []interface{} `json:"data"`
	Total int64         `json:"total"`
	Page  int64         `json:"page"`
	Limit int64         `json:"limit"`
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

// Agent ...
type Agent struct {
	Source   string `json:"source"`
	IP       string `json:"ip"`
	Platform string `json:"platform"`
	Version  string `json:"version"`
}

// StaffCheckPermissionBody ...
type StaffCheckPermissionBody struct {
	StaffID    string                 `json:"staffId"`
	Token      string                 `json:"token"`
	Permission []string               `json:"permission"`
	Source     string                 `json:"source"`
	Code       string                 `json:"code,omitempty"`
	IsRoot     string                 `json:"isRoot"`
	Agent      Agent                  `json:"agent"`
	DeviceId   string                 `json:"deviceId"`
	Payload    PayloadCheckPermission `json:"payload"`
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
	Code    int    `json:"code"`
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
