package repository

import (
	"context"
	"database/sql"
	"learningGoCleanArchitecture/domain/model"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *model.User)
}

type userRepository struct {
	db sql.DB
}

func NewUserReposoty(db *sql.DB) UserRepository {
	return &userRepository{*db}
}

func (ur &userRepository) CreateUser(ctx context.Context, user *model.User) {
	ur.db.Query("INSERT INTO テーブル名(列名１, 列名2.....)")
}
