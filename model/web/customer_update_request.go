package web

type CustomerUpdateRequest struct {
	Id    int    `validate:"required"`
	Name  string `validate:"required,max=255,min=1"`
	Email string `validate:"required,max=255,email"`
	Phone string `validate:"required,max=255,e164"`
}
