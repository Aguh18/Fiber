package request



type UserCreateRequest  struct {
	
	Name      string `json:"name" validate:"required"`
	Email     string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
	Address    string `json:"address" validate:"required"`
	Phone     string `json:"phone" validate:"required"`

}

type UserUpdateRequest  struct {
	
	Name      string `json:"name"`
	Email     string `json:"email" `
	Address    string `json:"address" `
	Phone     string `json:"phone" `

}


type UserEmailUpdate  struct {
	
	Email     string `json:"email" `

}


