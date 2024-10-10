package memberships

import "time"

type (
	SignUpRequest struct {
		Email    string `json:"email"`
		Username string `json:"username"`
		Password string `json:"password"`
	}

	LoginRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	RefreshTokenRequest struct {
		Token string `json:"token"`
	}
)

type (
	LoginResponse struct {
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
	}

	RefreshResponse struct {
		AccessToken string `json:"access_token"`
	}
)

type (
	UserModel struct {
		ID        int64     `db:"id"`
		Email     string    `db:"email"`
		Password  string    `db:"password"`
		Username  string    `db:"username"`
		CreatedAt time.Time `db:"created_at"`
		UpdatedAt time.Time `db:"updated_at"`
		CreatedBy string    `db:"created_by"`
		UpdatedBy string    `db:"updated_by"`
	}

	RefreshTokenModel struct {
		ID           int64     `db:"id"`
		UserID       int64     `db:"user_id"`
		RefreshToken string    `db:"refresh_token"`
		ExpiredAt    time.Time `db:"expired_at"`
		CreatedAt    time.Time `db:"created_at"`
		UpdatedAt    time.Time `db:"updated_at"`
		CreatedBy    string    `db:"created_by"`
		UpdatedBy    string    `db:"updated_by"`
	}
)
