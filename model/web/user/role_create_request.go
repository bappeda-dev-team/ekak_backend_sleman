package user

type RoleCreateRequest struct {
	Role string `json:"role" validate:"required"`
}
