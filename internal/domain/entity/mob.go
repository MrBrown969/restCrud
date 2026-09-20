package entity

import "errors"

type Mob struct {
	id   string
	name string
	lvl  string
	hp   int
}

func (m *Mob) Id() string {
	return m.id
}

func (m *Mob) Name() string {
	return m.name
}

func (m *Mob) Lvl() string {
	return m.lvl
}

func (m *Mob) Hp() int {
	return m.hp
}

type MobBuilder struct {
	id   string
	name string
	lvl  string
	hp   int
}

func (b *MobBuilder) Id(id string) *MobBuilder {
	b.id = id
	return b
}

func (b *MobBuilder) Name(name string) *MobBuilder {
	b.name = name
	return b
}

func (b *MobBuilder) Lvl(lvl string) *MobBuilder {
	b.lvl = lvl
	return b
}

func (b *MobBuilder) Hp(hp int) *MobBuilder {
	b.hp = hp
	return b
}

func (b *MobBuilder) Build() (*Mob, error) {
	err := validate(b)
	if err != nil {
		return nil, err
	}

	return &Mob{
		id:   b.id,
		name: b.name,
		lvl:  b.lvl,
		hp:   b.hp,
	}, nil
}

func validate(b *MobBuilder) error {
	if b.hp <= 0 {
		return errors.New("mob hp must be positive")
	}
	return nil
}
