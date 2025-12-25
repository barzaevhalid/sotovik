package user

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/barzaevhalid/sotovik/internal/domain"
	"github.com/barzaevhalid/sotovik/utils"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo *UserRepository
}

func NewUserService(repo *UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) Register(ctx context.Context, username, email, password, phone string) (string, string, error) {

	// existing, err := s.repo.GetByEmail(ctx, email)

	// if err != nil && errors.Is(err, pgx.ErrNoRows) {
	// 	return "", "", fmt.Errorf("get user by email: %w", err)
	// }
	// if existing != nil {
	// 	return "", "", domain.ErrUserAlreadyExists
	// }

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return "", "", err
	}

	user := &User{
		Username:     username,
		Email:        email,
		PasswordHash: string(hash),
		Phone:        phone,
	}

	if err := s.repo.Create(ctx, user); err != nil {
		return "", "", err
	}

	token, err := utils.GenerateJWT(user.ID, user.Role)
	if err != nil {
		return "", "", err
	}
	refreshToken, err := utils.GenerateRefreshToken(user.ID)

	if err != nil {
		return "", "", err
	}
	return token, refreshToken, nil
}

func (s *UserService) VerifyRefreshToken(tokenStr string) (int64, error) {

	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("REFRESH_SECRET")), nil
	})

	if err != nil || !token.Valid {
		return 0, fmt.Errorf("invalid refresh token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, fmt.Errorf("invalid refresh token claims")
	}

	sub, ok := claims["sub"].(string)
	if !ok {
		return 0, fmt.Errorf("invalid refresh token claims")
	}

	userId, err := strconv.ParseInt(sub, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid user id in token")
	}
	return userId, nil
}

func (s *UserService) GetById(ctx context.Context, id int64) (*User, error) {
	user, err := s.repo.GetById(ctx, id)

	if err != nil {
		return nil, err
	}

	return user, err

}

func (s *UserService) Login(ctx context.Context, email, password string) (string, error) {
	user, err := s.repo.GetByEmail(ctx, email)

	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			return "", ErrInvalidCredentials
		}
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return "", ErrInvalidCredentials
	}

	token, err := utils.GenerateJWT(user.ID, user.Role)
	if err != nil {
		return "", fmt.Errorf("generate jwt: %w", err)
	}
	return token, nil
}
