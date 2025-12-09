package dependencies

import (
	showtimePorts "github.com/newbie007fx/cinemas/internal/module/showtimes/ports"
	showtimeRepository "github.com/newbie007fx/cinemas/internal/module/showtimes/repositories"
	showtimeUsecase "github.com/newbie007fx/cinemas/internal/module/showtimes/usecase"
	userPorts "github.com/newbie007fx/cinemas/internal/module/users/ports"
	userRepository "github.com/newbie007fx/cinemas/internal/module/users/repositories"
	userUsecase "github.com/newbie007fx/cinemas/internal/module/users/usecase"
	"github.com/newbie007fx/cinemas/internal/transport/http/helpers/authentication"
	"github.com/newbie007fx/cinemas/platform/configuration"
	"github.com/newbie007fx/cinemas/platform/database"
)

type Dependency struct {
	DatabaseService *database.DatabaseService
	ConfigService   *configuration.ConfigService

	UserUsecase     userPorts.Usecase
	ShowtimeUsecase showtimePorts.Usecase

	AuthToken *authentication.TokenAuth
}

func New(db *database.DatabaseService, cs *configuration.ConfigService) *Dependency {
	return &Dependency{
		DatabaseService: db,
		ConfigService:   cs,
	}
}

func (dp *Dependency) Init() error {
	conf := dp.ConfigService.GetConfig()

	userRepo := userRepository.New(dp.DatabaseService)
	dp.UserUsecase = userUsecase.New(userRepo)

	showtimeRepo := showtimeRepository.New(dp.DatabaseService)
	dp.ShowtimeUsecase = showtimeUsecase.New(showtimeRepo)

	dp.AuthToken = authentication.New(conf.JWT.Secret)

	return nil
}
