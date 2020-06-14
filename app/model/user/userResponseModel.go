package user

type UserResponse struct {
	Role string   `json:"role,omitempty"`
	Data UserData `json:"user_data,omitempty"`
}
