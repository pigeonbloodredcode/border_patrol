package handler

import (
	"border_patrol/api/app/model"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"

	"strconv"
)

func GetAllQuestion_Chooses(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	question_numbers := []model.Question_Choose{}
	db.Find(&question_numbers)
	respondJSON(w, http.StatusOK, question_numbers)
}

func GetQuestion_Choose(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	log.Println("GetQuestion_Choose")
	question_number := vars["question_number"]
	question_choose := getQuestion_ChooseError404(db, question_number, w, r)
	if question_choose == nil {
		return
	}
	log.Println(question_choose)

	respondJSON(w, http.StatusOK, question_choose)
}
func CreateQuestion_Choose(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	//question := model.Question{}
	log.Println("CreateQuestion_Choose")
	question_choose := model.Question_Choose{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&question_choose); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()




	





	db.Create(&model.Question_Choose{QuestionNumber: 3, QuestionName: "xxx"})
	db.Create(&model.Question{QuestionID: 3, QuestionSet: 1})
	db.Create(&model.Score{QuestionSet_score: 3})

	// tx := db.Begin()
	// tx.Create(&user1)
	
	// tx.SavePoint("sp1")
	// tx.Create(&user2)
	// tx.RollbackTo("sp1") // Rollback user2
	
	// tx.Commit() // Commit user1




	//db.Create(&model.Lesson_Question{QuestionID: 3, })

	//result := db.FirstOrCreate(&Question, model.Question_Choose{QuestionNumber: 3, QuestionName:"xxx"})
	//db.Create(&model.Question{QuestionID:3, &model.Question_Choose: []&model.Question_Choose {QuestionNumber: 3, QuestionName:"xxx"}})
	//db.Create(&Dog{Name: "dog1", Toys: []Toy{{Name: "toy1"}, {Name: "toy2"}}})
	// if err := db.Save(&question_number).Error; err != nil {
	// 	respondError(w, http.StatusInternalServerError, err.Error())
	// 	return
	// }

	log.Println("CreateQuestion_Choose NewData")

	respondJSON(w, http.StatusCreated, question_choose)
}

func UpdateQuestion_Choose(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	title := vars["title"]
	log.Println("UpdateQuestion_Choose", title)
	question_number := getQuestion_ChooseError404(db, title, w, r)

	if question_number == nil {
		return
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&question_number); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&question_number).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, question_number)
}

func DeleteQuestion_Choose(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	question_number := vars["question_number"]
	log.Println("DeleteQuestion_Choose", question_number)
	vdo := getQuestion_ChooseError404(db, question_number, w, r)

	if vdo == nil {
		return
	}

	if err := db.Unscoped().Delete(&vdo).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	log.Println("Delete ", question_number, "Finish")

	respondJSON(w, http.StatusNoContent, nil)

}

// getQuestion_Chooseerror404 gets a vdo instance if exists, or respond the 404 error otherwise
func getQuestion_ChooseError404(db *gorm.DB, question_number string, w http.ResponseWriter, r *http.Request) *model.Question_Choose {
	intQN, err := strconv.Atoi(question_number) ////กรองเอาแค่ตัวเลขด้วย
	fmt.Println("convert str to int ", err)
	question_choose := model.Question_Choose{}
	if err := db.First(&question_choose, model.Question_Choose{QuestionNumber: int16(intQN)}).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return nil
	}
	return &question_choose
}
