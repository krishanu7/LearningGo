package storage

import "github.com/krishanu7/students-api/internal/types"

type Storage interface {
	CreateStudent(id int64, name string, age int, email string) (int64, error)
	GetStudent(id int64) (types.Student, error)
	GetAllStudents() ([]types.Student, error)
}
