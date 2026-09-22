package repository

import (
	"errors"
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

func (m *MobRepository) Delete(id string) error {
	query := `DELETE FROM mobs WHERE id = $1;`

	result, err := m.Db.Exec(query, id)
	if err != nil {
		return err
	}

	if rowsNum, _ := result.RowsAffected(); rowsNum == 0 {
		return fmt.Errorf("failed to delete by id: %w", err)
	}
	return nil
}

func (m *MobRepository) Create(name string, lvl string, id string, hp int) error {
	query := "INSERT INTO mobs (id, name, lvl, hp) VALUES ($1, $2, $3, $4);"

	result, err := m.Db.Exec(query, id, name, lvl, hp)

	if err != nil {
		return err
	}

	if count, _ := result.RowsAffected(); count == 0 {
		return errors.New("failed to insert mob")
	}

	return nil
}
