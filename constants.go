package authentication

const (
	SubjectRequestNatsCheckPermission = "authentication.request.staff.check_permission"
	SubjectRequestNatsGetPermission   = "authentication.request.staff.list_permission"
	SubjectRequestNatsSyncStaff       = "authentication.request.staff.sync_data"

	SubjectRequestNatsStaffCreate  = "authentication.request.staff.create"
	SubjectRequestNatsStaffUpdate  = "authentication.request.staff.update"
	SubjectRequestNatsStaffGetInfo = "authentication.request.staff.get_info"
	SubjectRequestNatsStaffGetList = "authentication.request.staff.list"

	SubjectRequestNatsAddLogs     = "authentication.request.logs.create"
	SubjectRequestNatsLogsGetList = "authentication.request.logs.list"

	SubjectRequestNatsGetTokenWithPhone = "authentication.request.staff.get_token"
)
