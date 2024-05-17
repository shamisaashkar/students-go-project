package store

import (
	"context"
	"fmt"
	"main.go/model"
)

// var (
//
//	ErrStudentNotFound  = errors.New("Student not found")
//	ErrDuplicateStudent = errors.New("duplicate student")
//
// )
type StudentNotFundError struct {
	ID int64
}

func (err StudentNotFundError) Error() string {
	return fmt.Sprintf("student %d  not found", err.ID)
}

type DuplicateStudentError struct {
	ID int64
}

func (err DuplicateStudentError) Error() string {
	return fmt.Sprintf("student %d  already exists", err.ID)
}

type Student interface {
	Save(context.Context, model.Student) error
	Get(context.Context, int64) (model.Student, error)
	GetAll(context.Context) ([]model.Student, error)
}
