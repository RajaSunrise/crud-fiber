package req

type User struct{
	Nama string `json:"nama"  validate:"required,min=10,max=50"`
	Email string `json:"email" validate:"required,email"`
	Umur int8 `json:"umur" validate:"required,gt=0"`
}