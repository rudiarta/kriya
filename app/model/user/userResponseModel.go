package user

type UserResponse struct {
	Role string   `json:"role,omitempty"`
	Data UserData `json:"user_data,omitempty"`
}

type UserGetListResponse struct {
	Username string      `json:"username,omitempty"`
	Email    string      `json:"email,omitempty"`
	Status   interface{} `json:"status,omitempty"`
}

type UserGetResponse struct {
	UserID   string `json:"user_id,omitempty"`
	Username string `json:"username,omitempty"`
	Email    string `json:"email,omitempty"`
	RoleName string `json:"role_name,omitempty"`
}
