package memberships

import (
	"context"
	"errors"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
	"situs-forum/internal/model/memberships"
	"situs-forum/pkg/jwt"
	tokenUtil "situs-forum/pkg/token"
	"strconv"
	"time"
)

func (s *Service) Login(ctx context.Context, req memberships.LoginRequest) (string, string, error) {
	user, err := s.repo.GetUser(ctx, req.Email, "", 0)
	if err != nil {
		log.Error().Err(err).Msg("Fail to get user")
		return "", "", err
	}
	if user == nil {
		return "", "", errors.New("email not exist")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))

	if err != nil {
		return "", "", errors.New("wrong password")
	}

	token, err := jwt.CreateToken(user.ID, user.Username, s.cfg.Service.SecretJWT)
	if err != nil {
		return "", "", err
	}

	existingRefreshToken, err := s.repo.GetRefreshToken(ctx, user.ID, time.Now())
	if err != nil {
		log.Error().Err(err).Msg("Error get latest refresh token from database")
		return "", "", err
	}

	if existingRefreshToken != nil {
		return token, existingRefreshToken.RefreshToken, nil
	}

	refreshToken := tokenUtil.GenerateRefreshToken()
	if refreshToken == "" {
		return token, "", errors.New("failed to generate refresh token")
	}

	err = s.repo.InsertRefreshToken(ctx, memberships.RefreshTokenModel{
		ID:           0,
		UserID:       user.ID,
		RefreshToken: refreshToken,
		ExpiredAt:    time.Now().Add(10 * 24 * time.Hour),
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
		CreatedBy:    strconv.FormatInt(user.ID, 10),
		UpdatedBy:    strconv.FormatInt(user.ID, 10),
	})

	if err != nil {
		log.Error().Err(err).Msg("Error inserting refresh token to database")
		return token, refreshToken, err
	}

	return token, refreshToken, nil
}
