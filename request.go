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

// SyncData ...
func (Request) SyncData(payload ListStaff) (*Response, error) {
	return requestNats(SubjectRequestNatsSyncStaff, toBytes(payload))
}

// CreateStaff ...
func (Request) CreateStaff(payload Staff) (*Response, error) {
	return requestNats(SubjectRequestNatsStaffCreate, toBytes(payload))
}

// UpdateStaff ...
func (Request) UpdateStaff(payload Staff) (*Response, error) {
	return requestNats(SubjectRequestNatsStaffUpdate, toBytes(payload))
}

// GetInfoStaff ...
func (Request) GetInfoStaff(payload GetInfoStaff) (*Response, error) {
	return requestNats(SubjectRequestNatsStaffGetInfo, toBytes(payload))
}

// GetListStaff ...
func (Request) GetListStaff(payload GetInfoStaff) (*Response, error) {
	return requestNats(SubjectRequestNatsStaffGetList, toBytes(payload))
}

// SaveLog ...
func (Request) SaveLog(payload Log) (*Response, error) {
	return requestNats(SubjectRequestNatsAddLogs, toBytes(payload))
}

// GetListLogs ...
func (Request) GetListLogs(payload *CommonQuery) (*Response, error) {
	return requestNats(SubjectRequestNatsLogsGetList, toBytes(payload))
}

// GetTokenWithPhoneNumber ...
func (Request) GetTokenWithPhoneNumber(payload *GetTokenByPhoneBody) (*Response, error) {
	return requestNats(SubjectRequestNatsGetTokenWithPhone, toBytes(payload))
}
