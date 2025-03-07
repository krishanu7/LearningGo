package sql

import (
	"database/sql"
	"fmt"

	"github.com/krishanu7/students-api/internal/config"
	"github.com/krishanu7/students-api/internal/types"
	_ "github.com/lib/pq"
)

type Psql struct {
	Db *sql.DB
}

func NewPsql(cfg *config.Config) (*Psql, error) {
	fmt.Println("Connecting to database...", cfg.StoragePath)
	// Open a connection to postgresql
	db, err := sql.Open("postgres", cfg.StoragePath)
	if err != nil {
		return nil, err
	}
	//Check if the connection works
	if err := db.Ping(); err != nil {
		return nil, err
	}
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS students (
		id SERIAL PRIMARY KEY,
		name TEXT NOT NULL,
		email TEXT NOT NULL,
		age INT
	)`)

	if err != nil {
		return nil, err
	}
	return &Psql{Db: db}, nil
}

func (p *Psql) CreateStudent(id int64, name string, age int, email string) (int64, error) {
	var lastId int64
	err := p.Db.QueryRow("INSERT INTO students (id, name, age, email) VALUES ($1, $2, $3, $4) RETURNING id",
		id, name, age, email).Scan(&lastId)

	if err != nil {
		return 0, err
	}
	return lastId, nil
}
 
func (p *Psql) GetStudent(id int64) (types.Student, error) {
	stmt, err := p.Db.Prepare("SELECT id, name, age, email FROM students WHERE id = $1")
	if err != nil {
		return types.Student{}, err
	}
	defer stmt.Close()
	var student types.Student
	err = stmt.QueryRow(id).Scan(&student.ID, &student.Name, &student.Age, &student.Email) // Order should be same as in the query
	if err != nil {
		return types.Student{}, err
	}
	return student, nil
}

func (p *Psql) GetAllStudents() ([]types.Student, error) {
	stmt, err := p.Db.Prepare("SELECT id, name, age, email FROM students")
	if(err != nil) {
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.Query()

	if(err != nil) {
		return nil, err
	}
	defer rows.Close()

	var students []types.Student
	
	for rows.Next() {
		var student types.Student
		err := rows.Scan(&student.ID, &student.Name, &student.Age, &student.Email)
		if err != nil {
			return nil, err
		}
		students = append(students, student)
	}
	return students, nil
}