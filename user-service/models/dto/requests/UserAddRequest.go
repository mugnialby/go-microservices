package requests

type UsersAddRequest struct {
	Username  string  `json:"username" validate:"required,min=1,max=32"`
	Password  string  `json:"password" validate:"required,min=8,max=32"`
	FirstName string  `json:"firstName" validate:"required,min=1,max=32"`
	LastName  *string `json:"lastName"`
	Email     *string `json:"email" validate:"required,min=8,max=32,email"`
	ManagerId uint64  `json:"managerId"`
	Status    *string `json:"status" validate:"required,min=1,max=1"`
	CreatedBy string  `json:"createdBy" validate:"required,min=1,max=32"`
}
