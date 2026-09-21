package usecase

import "fmt"

type IDeleteRepository interface {
	Delete(id string) error
}
type SlayMobUsecase struct {
	Repo IDeleteRepository
}

func (u *SlayMobUsecase) Slay(id string) error {
	err := u.Repo.Delete(id)
	if err != nil {
		return fmt.Errorf("failed to slay mob by id:%s: %s", id, err.Error())
	}
	return nil
}
