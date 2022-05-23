package authentication

// Request ...
type Request struct {
}

// CheckPermission ...
func (Request) CheckPermission(payload StaffCheckPermissionBody) (*Response, error) {
	return requestNats(SubjectRequestNatsCheckPermission, toBytes(payload))
}

// GetPermission ...
func (Request) GetPermission(payload StaffGetPermissions) (*Response, error) {
	return requestNats(SubjectRequestNatsGetPermission, toBytes(payload))
}
