package requests

type UsersUpdateRequest struct {
	FirstName  string  `json:"firstName" validate:"required,min=1,max=32"`
	LastName   *string `json:"lastName"`
	ModifiedBy *string `json:"modifiedBy" validate:"required,min=1,max=32"`
}
