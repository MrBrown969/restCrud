package usecase

import (
	"errors"
	"uuid"
)

type ICreateMobRepository interface {
	Create(name string, lvl string, id string, hp int) error
}
type SpawnMobUsecase struct {
	Repo ICreateMobRepository
}

func (u *SpawnMobUsecase) Spawn(name string, lvl string) (*string, error) {
	id := uuid.New().String()

	var hp int
	switch lvl {
	case "raid boss":
		hp = 1000000
	case "1":
		hp = 5
	case "chmo":
		hp = 6
	case "gym boss":
		hp = 50000
	default:
		hp = 1
	}

	err := u.Repo.Create(name, lvl, id, hp)
	if err != nil {
		return nil, errors.New("failed to create mob: " + err.Error())
	}

	return &id, nil
}
