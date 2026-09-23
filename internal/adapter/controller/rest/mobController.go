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
	Spawn(name string, lvl string) (*string, error)
}

type ISlayMobUsecase interface {
	Slay(id string) error
}

type IHitMobUsecase interface {
	Hit(id string, damage int) (*int, error)
}

type ISeeMobUsecase interface {
	See(id string) (*entity.Mob, error)
}

type MobController struct {
	SeeUsecase   ISeeMobUsecase
	SlayUsecase  ISlayMobUsecase
	SpawnUsecase ISpawnMobUsecase
	HitUsecase   IHitMobUsecase
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
		fmt.Println(err)
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
		fmt.Printf("error while slaying mob: %s", err.Error())
		respondWithJSON(w, http.StatusInternalServerError, map[string]string{"error:": "error while slaying mob"})

		return
	}

	respondWithJSON(w, http.StatusOK, nil)
}

func (c *MobController) Spawn(w http.ResponseWriter, r *http.Request) {
	var input model.MobSpawnRequest
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		fmt.Printf("error while spawning mob: %s\n", err.Error())
		respondWithJSON(w, http.StatusBadRequest, map[string]string{"error": "wrong request body"})

		return
	}

	name := input.Name
	if name == nil {
		fmt.Println("name cannot be nil")
		respondWithJSON(w, http.StatusBadRequest, map[string]string{"error": "name is required"})

		return
	}

	lvl := input.Lvl
	if lvl == nil {
		fmt.Println("lvl cannot be nil")
		respondWithJSON(w, http.StatusBadRequest, map[string]string{"error": "lvl is required"})

		return
	}

	id, err := c.SpawnUsecase.Spawn(*name, *lvl)

	if err != nil {
		fmt.Printf("error while spawning mob: %s", err.Error())
		respondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": "error while spawning mob"})

		return
	}

	spawnResponse := model.MobSpawnResponse{Id: *id}
	respondWithJSON(w, http.StatusOK, spawnResponse)
}

func (c *MobController) Hit(w http.ResponseWriter, r *http.Request) {
	var input model.MobHitRequest

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		fmt.Printf("error docoding MobHitRequest: %s", err.Error())
		respondWithJSON(w, http.StatusBadRequest, map[string]string{"error": "wrong request body"})

		return
	}

	if input.Damage == nil {
		fmt.Println("damage cannot be nil")
		respondWithJSON(w, http.StatusBadRequest, map[string]string{"error": "damage is required"})

		return
	}

	hpLeft, err := c.HitUsecase.Hit(input.Id, *input.Damage)
	if err != nil {
		fmt.Printf("error while hitting mob: %s", err.Error())
		respondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": "error while hitting mob"})

		return
	}

	encErr := json.NewEncoder(w).Encode(model.MobHitResponse{Id: input.Id, Hp: *hpLeft})
	if encErr != nil {
		fmt.Printf("error while encoding MobHitResponse: %s", err.Error())
		respondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": "error while hitting mob"})

		return
	}

}
