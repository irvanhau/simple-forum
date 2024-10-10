package posts

import (
	"context"
	"github.com/rs/zerolog/log"
	"situs-forum/internal/model/posts"
	"strconv"
	"time"
)

func (s *Service) CreateComment(ctx context.Context, postID, userID int64, req posts.CreateCommentRequest) error {
	now := time.Now()
	model := posts.CommentModel{
		PostID:         postID,
		UserID:         userID,
		CommentContent: req.CommentContent,
		CreatedAt:      now,
		UpdatedAt:      now,
		CreatedBy:      strconv.FormatInt(userID, 10),
		UpdatedBy:      strconv.FormatInt(userID, 10),
	}

	err := s.repo.CreateComment(ctx, model)
	if err != nil {
		log.Error().Err(err).Msg("Failed to create comment to repository")
		return err
	}
	return nil
}
