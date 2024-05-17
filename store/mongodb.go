package store

import (
	"context"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
	"main.go/model"
)

type StudentMongoDB struct {
	collection *mongo.Collection
	logger     *zap.Logger
}

const StudentCollection = "students"

func NewStudentMongoDB(db *mongo.Database, logger *zap.Logger) *StudentMongoDB {
	return &StudentMongoDB{
		collection: db.Collection(StudentCollection),
		logger:     logger,
	}

}

func (store *StudentMongoDB) Save(ctx context.Context, m model.Student) error {

	if _, err := store.collection.InsertOne(ctx, m); err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return DuplicateStudentError{
				ID: m.ID,
			}
		}
		return fmt.Errorf("document creation on mongodb failed: %w", err)
	}
	return nil
}
func (store *StudentMongoDB) Get(ctx context.Context, id int64) (model.Student, error) {
	var student model.Student
	res := store.collection.FindOne(ctx, bson.M{
		"id": id,
	})

	if err := res.Err(); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return student, StudentNotFundError{
				ID: id,
			}
		}
		return student, fmt.Errorf("can not read from collections")

	}

	if err := res.Decode(student); err != nil {
		return student, fmt.Errorf("cannot decode result into student %w", err)
	}
	return student, nil

}
func (store *StudentMongoDB) GetAll(ctx context.Context) ([]model.Student, error) {
	cursor, err := store.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, fmt.Errorf("cannot read from collection %w", err)
	}
	students := make([]model.Student, 0)

	for cursor.Next(ctx) {
		var student model.Student
		if err := cursor.Err(); err != nil {
			return nil, fmt.Errorf("can not read from current cursor from collections")

		}

		if err := cursor.Decode(&student); err != nil {
			return nil, fmt.Errorf("cannot decode current cursor into student %w", err)

		}
		students = append(students, student)

	}
	return students, nil

}
