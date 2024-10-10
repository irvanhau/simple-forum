package posts

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"situs-forum/internal/model/posts"
	"strconv"
)

func (h *Handler) CreateComment(c *gin.Context) {
	ctx := c.Request.Context()

	var req posts.CreateCommentRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	postIDStr := c.Param("postID")
	postID, err := strconv.ParseInt(postIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errors.New("Post ID Pada Param Tidak Valid").Error()})
		return
	}

	userID := c.GetInt64("userID")

	err = h.postSvc.CreateComment(ctx, postID, userID, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusCreated)
}
