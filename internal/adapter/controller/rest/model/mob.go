package model

type MobSpawnRequest struct {
	Id   string
	Name string
	Lvl  string
	Hp   int
}
type MobSpawnResponse struct {
	Id string
}

type MobSlayRequest struct {
	Id string
}
type MobSlayResponse struct {
	Id    string
	Slain bool
}

type MobSeeRequest struct {
	Id string
}
type MobSeeResponse struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Lvl  string `json:"lvl"`
	Hp   int    `json:"hp"`
}

type MobHitRequest struct {
	Id     string
	Damage *int
}
type MobHitResponse struct {
	Id string
	Hp int
}
