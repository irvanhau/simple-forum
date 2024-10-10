package posts

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"situs-forum/internal/model/posts"
	"strconv"
)

func (h *Handler) UpsertUserActivity(c *gin.Context) {
	ctx := c.Request.Context()

	var request posts.UserActivityRequest
	if err := c.Bind(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	postIDStr := c.Param("postID")
	postID, err := strconv.ParseInt(postIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errors.New("PostID pada param tidak valide").Error()})
		return
	}

	userID := c.GetInt64("userID")

	err = h.postSvc.UpsertUserActivity(ctx, postID, userID, request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}
