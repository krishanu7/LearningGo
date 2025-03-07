package types

// Crate a struct of student

type Student struct {
	ID    int64    `json:"id" validate:"required"`
	Name  string `json:"name" validate:"required"`
	Age   int	`json:"age"`
	Email string `json:"email" validate:"required,email"`
}
