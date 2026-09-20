package usecase

import "github.com/MrBrown969/restCrud/internal/domain/entity"

type IGetRepository interface {
	Get(id string) (*entity.Mob, error)
}
type SeeMobUsecase struct {
	Repo IGetRepository
}

func (uc *SeeMobUsecase) See(id string) (*entity.Mob, error) {
	mob, err := uc.Repo.Get(id)
	if err != nil {
		return nil, err
	}

	return mob, nil
}
