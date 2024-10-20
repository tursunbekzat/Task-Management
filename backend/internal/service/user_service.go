package service

import (
    "errors"
    "backend/internal/models"
    "backend/internal/repository"
    "github.com/golang-jwt/jwt"
    "time"
)

type UserService struct {
    repo      *repository.UserRepository
    jwtSecret string
}

func NewUserService(repo *repository.UserRepository, jwtSecret string) *UserService {
    return &UserService{
        repo:      repo,
        jwtSecret: jwtSecret,
    }
}

func (s *UserService) Register(create *models.UserCreate) (*models.User, error) {
    // Check if user already exists
    existingUser, _ := s.repo.GetUserByEmail(create.Email)
    if existingUser != nil {
        return nil, errors.New("email already registered")
    }

    user := &models.User{
        Email:     create.Email,
        Password:  create.Password,
        FirstName: create.FirstName,
        LastName:  create.LastName,
        Role:      "user",
    }

    // Hash password before saving
    if err := user.HashPassword(); err != nil {
        return nil, err
    }

    if err := s.repo.CreateUser(user); err != nil {
        return nil, err
    }

    return user, nil
}

func (s *UserService) Login(login *models.UserLogin) (string, error) {
    user, err := s.repo.GetUserByEmail(login.Email)
    if err != nil {
        return "", errors.New("invalid credentials")
    }

    if !user.CheckPassword(login.Password) {
        return "", errors.New("invalid credentials")
    }

    // Generate JWT token
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "user_id": user.ID,
        "email":   user.Email,
        "role":    user.Role,
        "exp":     time.Now().Add(time.Hour * 24).Unix(),
    })

    tokenString, err := token.SignedString([]byte(s.jwtSecret))
    if err != nil {
        return "", err
    }

    return tokenString, nil
}

func (s *UserService) UpdateUser(id uint, update *models.UserUpdate) (*models.User, error) {
    user, err := s.repo.GetUserByID(id)
    if err != nil {
        return nil, err
    }

    if update.FirstName != nil {
        user.FirstName = *update.FirstName
    }
    if update.LastName != nil {
        user.LastName = *update.LastName
    }
    if update.Password != nil {
        user.Password = *update.Password
        if err := user.HashPassword(); err != nil {
            return nil, err
        }
    }

    if err := s.repo.UpdateUser(user); err != nil {
        return nil, err
    }

    return user, nil
}