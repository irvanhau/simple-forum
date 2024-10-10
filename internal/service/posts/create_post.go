package posts

import (
	"context"
	"github.com/rs/zerolog/log"
	"situs-forum/internal/model/posts"
	"strconv"
	"strings"
	"time"
)

func (s *Service) CreatePost(ctx context.Context, userID int64, req posts.CreatePostRequest) error {
	postHastags := strings.Join(req.PostHashtags, `,`)
	// Contoh Join ['hastag1','hastag2'] => 'hastag1','hastag2'

	now := time.Now()

	model := posts.PostModel{
		UserID:       userID,
		PostTitle:    req.PostTitle,
		PostContent:  req.PostContent,
		PostHashtags: postHastags,
		CreatedAt:    now,
		UpdatedAt:    now,
		CreatedBy:    strconv.FormatInt(userID, 10),
		UpdatedBy:    strconv.FormatInt(userID, 10),
	}

	err := s.repo.CreatePost(ctx, model)
	if err != nil {
		log.Error().Err(err).Msg("error create post to repository")
		return err
	}
	return nil
}
