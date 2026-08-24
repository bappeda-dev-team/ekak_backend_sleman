package user

type UserCreateRequest struct {
	Nip      string                  `json:"nip"`
	Email    string                  `json:"email"`
	Password string                  `json:"password"`
	IsActive bool                    `json:"is_active"`
	Role     []UserRoleCreateRequest `json:"role"`
}

type UserRoleCreateRequest struct {
	RoleId int `json:"role_id"`
}
