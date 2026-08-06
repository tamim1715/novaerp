package attendance

func ToAttendanceResponse(a *Attendance) AttendanceResponse {
	resp := AttendanceResponse{
		ID:            a.ID.String(),
		EmployeeID:    a.EmployeeID.String(),
		Date:          a.Date.Format("2006-01-02"),
		WorkHours:     a.WorkHours,
		OvertimeHours: a.OvertimeHours,
		Status:        a.Status,
		Notes:         a.Notes,
		CreatedAt:     a.CreatedAt.Unix(),
	}

	if a.CheckIn != nil {
		resp.CheckIn = a.CheckIn.Format("2006-01-02T15:04:05Z07:00")
	}
	if a.CheckOut != nil {
		resp.CheckOut = a.CheckOut.Format("2006-01-02T15:04:05Z07:00")
	}
	if a.Employee.FirstName != "" {
		resp.EmployeeName = a.Employee.FirstName + " " + a.Employee.LastName
	}

	return resp
}

func ToAttendanceResponseList(list []Attendance) []AttendanceResponse {
	resp := make([]AttendanceResponse, len(list))
	for i, item := range list {
		resp[i] = ToAttendanceResponse(&item)
	}
	return resp
}
