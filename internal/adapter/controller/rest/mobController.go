package rest

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/MrBrown969/restCrud/internal/adapter/controller/rest/model"
	"github.com/MrBrown969/restCrud/internal/domain/entity"
)

// /mob-pool/
type ISpawnMobUsecase interface {
	Spawn()
}

type ISlayMobUsecase interface {
	Slay(id string) error
}

type IHitMobUsecase interface {
	Hit()
}

type ISeeMobUsecase interface {
	See(id string) (*entity.Mob, error)
}

type MobController struct {
	SeeUsecase  ISeeMobUsecase
	SlayUsecase ISlayMobUsecase
}

func New(see *ISeeMobUsecase) *MobController {
	return &MobController{}
}

func Spawn(request model.MobSpawnRequest) {

}

func (c *MobController) See(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/mob-pool/")
	mob, err := c.SeeUsecase.See(id)
	if err != nil {
		respondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	response := model.MobSeeResponse{}
	response.Id = mob.Id()
	response.Name = mob.Name()
	response.Hp = mob.Hp()
	response.Lvl = mob.Lvl()
	respondWithJSON(w, http.StatusOK, response)
	fmt.Println(mob)
}

func respondWithJSON(w http.ResponseWriter, code int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	err := json.NewEncoder(w).Encode(payload)
	if err != nil {
		fmt.Println(fmt.Errorf("failed to encode response %s", err))
		return
	}
}

func (c *MobController) Slay(w http.ResponseWriter, r *http.Request) {
	err := c.SlayUsecase.Slay(r.PathValue("id"))
	if err != nil {
		respondWithJSON(w, http.StatusInternalServerError, map[string]string{"error while slaying mob": err.Error()})
		return
	}

	respondWithJSON(w, http.StatusOK, nil)

}
