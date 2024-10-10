package posts

import (
	"context"
	"github.com/rs/zerolog/log"
	"situs-forum/internal/model/posts"
)

func (s *Service) GetPostByID(ctx context.Context, postID int64) (*posts.GetPostResponse, error) {
	postDetail, err := s.repo.GetPostByID(ctx, postID)
	if err != nil {
		log.Error().Err(err).Msg("error get post by id to database")
		return nil, err
	}

	likeCount, err := s.repo.CountLikeByPostID(ctx, postID)
	if err != nil {
		log.Error().Err(err).Msg("error count like by post id to database")
		return nil, err
	}

	comments, err := s.repo.GetCommentByPostID(ctx, postID)
	if err != nil {
		log.Error().Err(err).Msg("error get all comment by post id to database")
		return nil, err
	}

	return &posts.GetPostResponse{
		PostDetail: posts.Post{
			ID:           postDetail.ID,
			UserID:       postDetail.UserID,
			UserName:     postDetail.UserName,
			PostTitle:    postDetail.PostTitle,
			PostContent:  postDetail.PostContent,
			PostHashtags: postDetail.PostHashtags,
			IsLiked:      postDetail.IsLiked,
		},
		LikeCount: likeCount,
		Comments:  comments,
	}, nil
}
