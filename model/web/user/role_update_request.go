package user

type RoleUpdateRequest struct {
	Id   int    `json:"id"`
	Role string `json:"role" validate:"required"`
}
