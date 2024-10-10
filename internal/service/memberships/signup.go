package memberships

import (
	"context"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"situs-forum/internal/model/memberships"
	"time"
)

func (s *Service) SignUp(ctx context.Context, req memberships.SignUpRequest) error {
	user, err := s.repo.GetUser(ctx, req.Email, req.Username, 0)
	if err != nil {
		return err
	}

	if user != nil {
		return errors.New("Username or Email already exists")
	}

	pass, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	now := time.Now()
	model := memberships.UserModel{
		Email:     req.Email,
		Password:  string(pass),
		Username:  req.Username,
		CreatedAt: now,
		UpdatedAt: now,
		CreatedBy: req.Email,
		UpdatedBy: req.Email,
	}

	err = s.repo.CreateUser(ctx, &model)

	if err != nil {
		return err
	}

	return nil
}
