package leaverequest

func ToLeaveRequestResponse(lr *LeaveRequest) LeaveRequestResponse {
	resp := LeaveRequestResponse{
		ID:          lr.ID.String(),
		EmployeeID:  lr.EmployeeID.String(),
		LeaveTypeID: lr.LeaveTypeID.String(),
		StartDate:   lr.StartDate.Format("2006-01-02"),
		EndDate:     lr.EndDate.Format("2006-01-02"),
		TotalDays:   lr.TotalDays,
		Reason:      lr.Reason,
		Status:      lr.Status,
		Remarks:     lr.Remarks,
		CreatedAt:   lr.CreatedAt.Unix(),
	}

	if lr.Employee.FirstName != "" {
		resp.EmployeeName = lr.Employee.FirstName + " " + lr.Employee.LastName
	}
	if lr.LeaveType.Name != "" {
		resp.LeaveTypeName = lr.LeaveType.Name
	}
	if lr.ApprovedBy != nil {
		resp.ApprovedBy = lr.ApprovedBy.String()
	}

	return resp
}

func ToLeaveRequestResponseList(list []LeaveRequest) []LeaveRequestResponse {
	resp := make([]LeaveRequestResponse, len(list))
	for i, item := range list {
		resp[i] = ToLeaveRequestResponse(&item)
	}
	return resp
}
