package posts

import (
	"context"
	"situs-forum/internal/configs"
	"situs-forum/internal/model/posts"
)

type postRepository interface {
	GetAllPost(ctx context.Context, limit, offset int) (posts.GetAllResponse, error)
	CreatePost(ctx context.Context, mdl posts.PostModel) error
	GetPostByID(ctx context.Context, id int64) (*posts.Post, error)

	CreateComment(ctx context.Context, mdl posts.CommentModel) error
	GetCommentByPostID(ctx context.Context, postID int64) ([]posts.Comment, error)

	CreateUserActivity(ctx context.Context, model posts.UserActivityModel) error
	GetUserActivity(ctx context.Context, model posts.UserActivityModel) (*posts.UserActivityModel, error)
	UpdateUserActivity(ctx context.Context, model posts.UserActivityModel) error
	CountLikeByPostID(ctx context.Context, postID int64) (int, error)
}

type Service struct {
	repo postRepository
	cfg  *configs.Config
}

func NewService(cfg *configs.Config, repo postRepository) *Service {
	return &Service{
		repo: repo,
		cfg:  cfg,
	}
}
