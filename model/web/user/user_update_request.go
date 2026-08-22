package user

type UserUpdateRequest struct {
	Id       int                     `json:"id"`
	Nip      string                  `json:"nip"`
	Email    string                  `json:"email"`
	Password string                  `json:"password"`
	IsActive bool                    `json:"is_active"`
	Role     []UserRoleUpdateRequest `json:"role"`
}

type UserRoleUpdateRequest struct {
	RoleId int `json:"role_id"`
}
