package posts

import (
	"context"
	"errors"
	"github.com/rs/zerolog/log"
	"situs-forum/internal/model/posts"
	"strconv"
	"time"
)

func (s *Service) UpsertUserActivity(ctx context.Context, postID, userID int64, req posts.UserActivityRequest) error {
	now := time.Now()
	model := posts.UserActivityModel{
		PostID:    postID,
		UserID:    userID,
		IsLiked:   req.IsLiked,
		CreatedAt: now,
		UpdatedAt: now,
		CreatedBy: strconv.FormatInt(userID, 10),
		UpdatedBy: strconv.FormatInt(userID, 10),
	}

	userActivity, err := s.repo.GetUserActivity(ctx, model)
	if err != nil {
		log.Error().Err(err).Msg("Error get user activity from database")
		return err
	}
	if userActivity == nil {
		//	Create User Activity
		if !req.IsLiked {
			return errors.New("anda belum pernah like sebelumnya")
		}
		err = s.repo.CreateUserActivity(ctx, model)
	} else {
		//  Update User Activity
		err = s.repo.UpdateUserActivity(ctx, model)
	}

	if err != nil {
		log.Error().Err(err).Msg("Error create / update user activity to database")
		return err
	}

	return nil
}
