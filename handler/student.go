package handler

import (
	"errors"
	"io"
	// "log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"main.go/model"
	"main.go/request"
	"main.go/store"
)

type Student struct {
	Store  store.Student
	Logger *zap.Logger
}

func (s Student) GetAll(c echo.Context) error {
	ctx := c.Request().Context()
	ss, err := s.Store.GetAll(ctx)
	if err != nil {
		return echo.ErrInternalServerError
	}
	return c.JSON(http.StatusOK, ss)

}
func (s Student) Get(c echo.Context) error {
	ctx := c.Request().Context()
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return echo.ErrInternalServerError
	}

	st, err := s.Store.Get(ctx, id)
	if err != nil {
		var errNotFound store.StudentNotFundError
		if ok := errors.As(err, &errNotFound); ok {
			// log.Println(err)
			s.Logger.Error("student not found", zap.Error(err))

			return echo.ErrNotFound

		}
		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, st)

}
func (s Student) Create(c echo.Context) error {
	var req request.Student
	if err := c.Bind(&req); err != nil {
		body, _ := io.ReadAll(c.Request().Body)
		// log.Println(err)
		s.Logger.Error("cannnot bind request to student",
			zap.Error(err),
			zap.ByteString("body", body),
		)
		return echo.ErrBadRequest

	}

	if err := req.Validate(); err != nil {
		// log.Println(err)
		s.Logger.Error("request validation failed", zap.Error(err), zap.Any("request", req))
		return echo.ErrBadRequest
	}

	m := model.Student{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		ID:        req.ID,
		Average:   0,
	}
	s.Logger.Info("student creation succeeded")
	ctx := c.Request().Context()
	if err := s.Store.Save(ctx, m); err != nil {
		var errDuplicate store.DuplicateStudentError
		if ok := errors.As(err, &errDuplicate); ok {
			s.Logger.Error("duplicate student",
				zap.Error(err),
				zap.Int64("id", m.ID),
			)
			// log.Println(err)
			return echo.ErrBadRequest
		}
		return echo.ErrInternalServerError
	}
	return c.JSON(http.StatusCreated, nil)

}

func (s Student) Register(g *echo.Group) {
	g.GET("/", s.GetAll)
	g.GET("/:id", s.Get)
	g.POST("/", s.Create)

}
