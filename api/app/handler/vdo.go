package handler

import (
	"encoding/json"
	"net/http"

	"border_patrol/api/app/model"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func GetAllVdos(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vdos := []model.Vdo{}
	db.Find(&vdos)
	respondJSON(w, http.StatusOK, vdos)
}

func CreateVdo(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vdo := model.Vdo{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&vdo); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&vdo).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusCreated, vdo)
}

func GetVdo(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	name := vars["name"]
	vdo := getVdoError404(db, name, w, r)
	if vdo == nil {
		return
	}
	respondJSON(w, http.StatusOK, vdo)
}

func UpdateVdo(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	name := vars["name"]
	vdo := getVdoError404(db, name, w, r)
	if vdo == nil {
		return
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&vdo); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&vdo).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, vdo)
}

func DeleteVdo(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	name := vars["name"]
	vdo := getVdoError404(db, name, w, r)
	if vdo == nil {
		return
	}
	if err := db.Delete(&vdo).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusNoContent, nil)
}

func DisableVdo(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	name := vars["name"]
	vdo := getVdoError404(db, name, w, r)
	if vdo == nil {
		return
	}
	vdo.Disable()
	if err := db.Save(&vdo).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, vdo)
}

func EnableVdo(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	name := vars["name"]
	vdo := getVdoError404(db, name, w, r)
	if vdo == nil {
		return
	}
	vdo.Enable()
	if err := db.Save(&vdo).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, vdo)
}

// getVdoerror404 gets a vdo instance if exists, or respond the 404 error otherwise
func getVdoError404(db *gorm.DB, title string, w http.ResponseWriter, r *http.Request) *model.Vdo {
	vdo := model.Vdo{}
	if err := db.First(&vdo, model.Vdo{Title: title}).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return nil
	}
	return &vdo
}

