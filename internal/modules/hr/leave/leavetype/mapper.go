package leavetype

func ToLeaveTypeResponse(lt *LeaveType) LeaveTypeResponse {
	return LeaveTypeResponse{
		ID:             lt.ID.String(),
		Name:           lt.Name,
		Code:           lt.Code,
		MaxDaysPerYear: lt.MaxDaysPerYear,
		IsPaid:         lt.IsPaid,
		Description:    lt.Description,
		CreatedAt:      lt.CreatedAt.Unix(),
	}
}

func ToLeaveTypeResponseList(list []LeaveType) []LeaveTypeResponse {
	resp := make([]LeaveTypeResponse, len(list))
	for i, item := range list {
		resp[i] = ToLeaveTypeResponse(&item)
	}
	return resp
}
