package user

func ToResponse(user *User) UserResponse {
	role := user.Role
	if role == "" {
		role = "user"
	}

	return UserResponse{
		ID:        user.ID.String(),
		Username:  user.Username,
		Email:     user.Email,
		Role:      role,
		IsActive:  user.IsActive,
		CreatedAt: user.CreatedAt.Unix(),
	}
}

func ToResponses(users []User) []UserResponse {
	responses := make([]UserResponse, 0, len(users))

	for _, user := range users {
		responses = append(responses, ToResponse(&user))
	}

	return responses
}
