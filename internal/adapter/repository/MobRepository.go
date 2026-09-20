package repository

import (
	"fmt"

	"github.com/MrBrown969/restCrud/internal/adapter/repository/model"
	"github.com/MrBrown969/restCrud/internal/domain/entity"
	"github.com/jmoiron/sqlx"
)

type MobRepository struct {
	Db *sqlx.DB
}

func (m *MobRepository) Get(id string) (*entity.Mob, error) {
	var mobRepo model.MobRepoModel
	query := `SELECT * FROM mobs WHERE id = $1;`
	err := m.Db.Get(&mobRepo, query, id)

	if err != nil {
		return nil, fmt.Errorf("failed to get mob by id: %w", err)
	}

	builder := &entity.MobBuilder{}
	mob, err := builder.Id(mobRepo.Id).Name(mobRepo.Name).Hp(mobRepo.Hp).Lvl(mobRepo.Lvl).Build()
	if err != nil {
		return nil, fmt.Errorf("failed to build mob by id: %w", err)
	}

	return mob, nil
}
