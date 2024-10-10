package memberships

import (
	"context"
	"github.com/gin-gonic/gin"
	"situs-forum/internal/middleware"
	"situs-forum/internal/model/memberships"
)

type MembershipService interface {
	SignUp(ctx context.Context, req memberships.SignUpRequest) error
	Login(ctx context.Context, req memberships.LoginRequest) (string, string, error)
	ValidateRefreshToken(ctx context.Context, userID int64, req memberships.RefreshTokenRequest) (string, error)
}

type Handler struct {
	*gin.Engine
	membershipSvc MembershipService
}

func NewHandler(api *gin.Engine, service MembershipService) *Handler {
	return &Handler{
		Engine:        api,
		membershipSvc: service,
	}
}

func (h *Handler) RegisterRoute() {
	route := h.Group("memberships")
	route.POST("/sign-up", h.SignUp)
	route.POST("/login", h.Login)
	route.GET("/ping", h.Ping)

	routeRefresh := h.Group(`memberships`)
	routeRefresh.Use(middleware.AuthRefreshMiddleware())
	routeRefresh.POST(`/refresh`, h.Refresh)
}
