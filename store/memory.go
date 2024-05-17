package store

import (
	// "github.com/labstack/gommon/log"
	"context"
	"go.uber.org/zap"
	"main.go/model"
)

type StudentINMemory struct {
	students map[int64]model.Student
	logger   *zap.Logger
}

func NewStudentINMemory(logger *zap.Logger) *StudentINMemory {
	return &StudentINMemory{
		students: make(map[int64]model.Student),
		logger:   logger,
	}

}
func (m *StudentINMemory) Save(_ context.Context, s model.Student) error {
	if _, ok := m.students[s.ID]; ok {
		return DuplicateStudentError{
			ID: s.ID,
		}
	}
	m.students[s.ID] = s

	m.logger.Debug("current students", zap.Any("students", m.students))
	return nil

}
func (m *StudentINMemory) Get(_ context.Context, id int64) (model.Student, error) {
	s, ok := m.students[id]
	if ok {
		return s, nil
	}
	return s, StudentNotFundError{
		ID: id,
	}

}
func (m *StudentINMemory) GetAll(_ context.Context) ([]model.Student, error) {
	ss := make([]model.Student, 0)
	for _, s := range m.students {
		ss = append(ss, s)
	}
	return ss, nil

}
