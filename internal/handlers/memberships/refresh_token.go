package memberships

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"situs-forum/internal/model/memberships"
)

func (h *Handler) Refresh(c *gin.Context) {
	ctx := c.Request.Context()

	var req memberships.RefreshTokenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID := c.GetInt64("userID")

	accessToken, err := h.membershipSvc.ValidateRefreshToken(ctx, userID, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, memberships.RefreshResponse{
		AccessToken: accessToken,
	})
}
