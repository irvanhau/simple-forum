package memberships

import (
	"context"
	"errors"
	"github.com/rs/zerolog/log"
	"situs-forum/internal/model/memberships"
	"situs-forum/pkg/jwt"
	"time"
)

func (s *Service) ValidateRefreshToken(ctx context.Context, userID int64, req memberships.RefreshTokenRequest) (string, error) {
	existingRefreshToken, err := s.repo.GetRefreshToken(ctx, userID, time.Now())
	if err != nil {
		log.Error().Err(err).Msg("Error get refresh token from database")
		return "", err
	}

	if existingRefreshToken == nil {
		return "", errors.New("refresh token has expired")
	}

	// means token in database is not matched with request token, throw error invalid refresh token
	if existingRefreshToken.RefreshToken != req.Token {
		return "", errors.New("refresh token is invalid")
	}

	user, err := s.repo.GetUser(ctx, "", "", userID)
	if err != nil {
		log.Error().Err(err).Msg("Error get user from database")
		return "", err
	}
	if user == nil {
		return "", errors.New("user not found")
	}

	token, err := jwt.CreateToken(user.ID, user.Username, s.cfg.Service.SecretJWT)
	if err != nil {
		log.Error().Err(err).Msg("Error create token")
		return "", err
	}

	return token, nil
}
