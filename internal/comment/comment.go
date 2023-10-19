package comment

import (
	"context"
	"errors"
	"fmt"
)

var (
	ErrFetchingCmt   = errors.New("error fetching comment")
	ErrNotImplmented = errors.New("error not implemented")
)

type Comment struct {
	Id     string
	Slug   string
	Body   string
	Author string
}

type Store interface {
	GetComment(ctx context.Context, id string) (Comment, error)
}

type Service struct {
	Store Store
}

func NewService(store Store) *Service {
	return &Service{
		Store: store,
	}
}

func (s *Service) GetComment(ctx context.Context, id string) (Comment, error) {
	c, err := s.Store.GetComment(ctx, id)
	if err != nil {
		fmt.Println(err)
		return Comment{}, ErrFetchingCmt
	}
	return c, nil
}

func (s *Service) UpdateComment(ctx context.Context, cmt Comment) error {
	return ErrNotImplmented
}

func (s *Service) DeleteComment(ctx context.Context, id string) error {
	return ErrNotImplmented
}

func (s *Service) CreateComment(ctx context.Context, cmt Comment) (Comment, error) {
	return Comment{}, ErrNotImplmented
}
