package usecase

import (
	"time"

	"github.com/newbie007fx/cinemas/internal/module/users/entities"
	"github.com/newbie007fx/cinemas/internal/module/users/models"
	"github.com/newbie007fx/cinemas/internal/module/users/ports"
)

type Usecase struct {
	repo ports.Repository
}

func New(repo ports.Repository) ports.Usecase {
	return &Usecase{
		repo: repo,
	}
}

func (Usecase) mapEntityToModel(userEntity *entities.User) *models.User {
	return &models.User{
		ID:        userEntity.ID,
		Name:      userEntity.Name,
		Username:  userEntity.Username,
		Email:     userEntity.Email,
		CreatedAt: userEntity.CreatedAt.Format(time.RFC3339),
	}
}
