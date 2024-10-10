package memberships

import (
	"context"
	"situs-forum/internal/configs"
	"situs-forum/internal/model/memberships"
	"time"
)

type membershipRepository interface {
	GetUser(ctx context.Context, email, username string, userID int64) (*memberships.UserModel, error)
	CreateUser(ctx context.Context, model *memberships.UserModel) error

	InsertRefreshToken(ctx context.Context, model memberships.RefreshTokenModel) error
	GetRefreshToken(ctx context.Context, userID int64, now time.Time) (*memberships.RefreshTokenModel, error)
}

type Service struct {
	repo membershipRepository
	cfg  *configs.Config
}

func NewService(repo membershipRepository, config *configs.Config) *Service {
	return &Service{repo: repo, cfg: config}
}
