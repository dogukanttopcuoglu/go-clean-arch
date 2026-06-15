package memory

import (
	"context"

	"github.com/dogukanttopcuoglu/clean-lab/internal/entity"
)

type UserRepo struct {
	usersByID    map[string]entity.User
	usersByEmail map[string]entity.User
}

func NewUserRepo() *UserRepo {
	return &UserRepo{
		usersByID:    make(map[string]entity.User),
		usersByEmail: make(map[string]entity.User),
	}
}

func (r *UserRepo) Store(ctx context.Context, user *entity.User) error {
	_ = ctx
	if _, ok := r.usersByEmail[user.Email]; ok {
		return entity.ErrUserAlreadyExists
	}

	if _, ok := r.usersByID[user.ID]; ok {
		return entity.ErrUserAlreadyExists
	}

	r.usersByID[user.ID] = *user

	r.usersByEmail[user.Email] = *user

	return nil
}

func (r *UserRepo) GetByID(ctx context.Context, id string) (entity.User, error) {
	_ = ctx
	user, ok := r.usersByID[id]
	if !ok {
		return entity.User{}, entity.ErrUserNotFound
	}

	return user, nil
}

func (r *UserRepo) GetByEmail(ctx context.Context, email string) (entity.User, error) {
	_ = ctx
	user, ok := r.usersByEmail[email]
	if !ok {
		return entity.User{}, entity.ErrUserNotFound
	}

	return user, nil
}
