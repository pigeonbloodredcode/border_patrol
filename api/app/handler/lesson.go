package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"border_patrol/api/app/model"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func GetAllLessons(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	lessons := []model.Lesson{}
	db.Find(&lessons)
	respondJSON(w, http.StatusOK, lessons)
}

func CreateLesson(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	lesson := model.Lesson{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&lesson); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&lesson).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusCreated, lesson)
}

func GetLesson(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	title := vars["title"]
	lesson := getLessonError404(db, title, w, r)
	if lesson == nil {
		return
	}
	respondJSON(w, http.StatusOK, lesson)
}

func UpdateLesson(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	title := vars["title"]
	log.Println("UpdateLesson",title)
	lesson := getLessonError404(db, title, w, r)

	if lesson == nil {
		return
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&lesson); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&lesson).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, lesson)
}

func DeleteLesson(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	title := vars["title"]
	log.Println("DeleteLesson",title)
	lesson := getLessonError404(db, title, w, r)

	if lesson == nil {
		return
	}

	if err := db.Unscoped().Delete(&lesson).Error; err != nil {			
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	log.Println("Delete ", title, "Finish")

	respondJSON(w, http.StatusNoContent, nil)
	
}

// getLessonerror404 gets a lesson instance if exists, or respond the 404 error otherwise
func getLessonError404(db *gorm.DB, content string, w http.ResponseWriter, r *http.Request) *model.Lesson {
	lesson := model.Lesson{}
	if err := db.First(&lesson, model.Lesson{Content: content}).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return nil
	}
	return &lesson
}

