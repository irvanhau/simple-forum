package posts

import (
	"context"
	"github.com/gin-gonic/gin"
	"situs-forum/internal/middleware"
	"situs-forum/internal/model/posts"
)

type postService interface {
	GetAllPost(ctx context.Context, pageSize, pageIndex int) (posts.GetAllResponse, error)
	CreatePost(ctx context.Context, userID int64, req posts.CreatePostRequest) error
	GetPostByID(ctx context.Context, postID int64) (*posts.GetPostResponse, error)

	CreateComment(ctx context.Context, postID, userID int64, req posts.CreateCommentRequest) error

	UpsertUserActivity(ctx context.Context, postID, userID int64, req posts.UserActivityRequest) error
}

type Handler struct {
	*gin.Engine
	postSvc postService
}

func NewHandler(postSvc postService, g *gin.Engine) *Handler {
	return &Handler{
		Engine:  g,
		postSvc: postSvc,
	}
}

func (h *Handler) RegisterRoute() {
	route := h.Group("posts")
	route.Use(middleware.AuthMiddleware())

	route.GET(`/`, h.GetAllPost)
	route.POST(`/create`, h.CreatePost)
	route.POST(`/comment/:postID`, h.CreateComment)
	route.PUT(`/user_activity/:postID`, h.UpsertUserActivity)
	route.GET(`/:postID`, h.GetPostByID)
}
