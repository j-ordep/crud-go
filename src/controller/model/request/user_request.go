package request

// validator = validate:"required" ("required" = "campo obrigatorio", não pode ser vazio)
// como usamos Gin a validação é com -> binding:""
type UserRequest struct {
	Email 	 	string `json:"email" binding:"required,email"`
	Password	string `json:"password" binding:"required,min=6,containsany=!@#$%*_"`
	Name 		string `json:"name" binding:"required,min=3,max=100"` // min e max de caracteres
	Age 		int8   `json:"age" binding:"required,min=1,max=120"` // min e max de valores numericos
}