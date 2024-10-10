package posts

import (
	"context"
	"github.com/rs/zerolog/log"
	"situs-forum/internal/model/posts"
)

func (s *Service) GetAllPost(ctx context.Context, pageSize, pageIndex int) (posts.GetAllResponse, error) {
	limit := pageSize
	offset := pageSize * (pageIndex - 1)
	response, err := s.repo.GetAllPost(ctx, limit, offset)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get all posts from database")
		return response, err
	}
	return response, nil
}
