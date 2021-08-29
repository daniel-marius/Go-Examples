package repository

import (
  "go-mux-api/entity"
)

type PostRepository interface {
  Save(post *entity.Post) (*entity.Post, error)
  FindAll() ([]entity.Post, error)
}
