package usecase

import "fmt"

type IUpdateMobRepository interface {
	Update(id string, damage int) (*int, error)
}
type HitMobUsecase struct {
	Repo IUpdateMobRepository
}

func (u *HitMobUsecase) Hit(id string, damage int) (*int, error) {
	if damage < 1 {
		fmt.Println("no damage dealt")
		zero := 0

		return &zero, nil
	}

	hpLeft, err := u.Repo.Update(id, damage)
	if err != nil {
		return nil, err
	}

	return hpLeft, nil
}
