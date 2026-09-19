package rest

type MobCreateRequest struct {
	Id   string
	Name string
	Lvl  string
	Hp   int
}
type MobCreateResponse struct {
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
	Id   string
	Name string
	Lvl  string
	Hp   int
}

type MobHitRequest struct {
	Id     string
	Damage *int
}
type MobHitResponse struct {
	Id string
	Hp int
}
