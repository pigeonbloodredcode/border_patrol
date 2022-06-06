package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"border_patrol/api/app/model"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func GetAllLessons(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	lessons := []model.Lesson{}
	lesson_question := []model.Lesson_Question{}
	db.Find(&lessons, &lesson_question)

	respondJSON(w, http.StatusOK, lessons)
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

func CreateLesson(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	fmt.Println("CreateLesson")

	var contentMany = "beginGiraffe GiraffeGiraffeGiraffeGiraffeGiraffeGiraffeGiraffeGiraffeGiraffeGiraffeGiraffeGiraffeGiraffeGiraffeGiraffeGiraffeGiraffeGiraffeGiraffeGiraffeGiraffeGiraffeGiraffeGiraffeGiraffeGiraffeGiraffeGiraffeGiraffeGiraffeGiraffeGiraffeGiraffeGiraffeGiraffeGiraffeGiraffeGiraffeGiraffeGiraffeGiraffeGiraffeGiraffeGiraffeGiraffeGiraffeGiraffeGiraffeGiraffeGiraffeGiraffeGiraffeGiraffeGiraffeGiraffeGiraffeGiraffeGiraffeGiraffeGiraffeGiraffeGiraffeGiraffeGiraffeGiraffeGiraffeGiraffeGiraffeGiraffeGiraffeGiraffeGiraffeGiraffeGiraffeGiraffeGiraffeGiraffeGiraffeGiraffeGiraffeGiraffeGiraffeGiraffeGiraffeGiraffeGiraffeGiraffeEnd"
	var lesson = []model.Lesson{{Header: "4", Content: contentMany}}

	if err := db.Create(&lesson).Error; err != nil {
		fmt.Println(err.Error())
	}
	var lession_id int
	for _, data2 := range lesson {
		lession_id = int(data2.ID)
	}
	fmt.Println(lession_id)

	tx := db.Begin()
	//คำถามครั้งแรกจะตั้งเป็นค่า ตาม leson ก่อนหากมีการเพิ่มคำถามค่อยมาปรับ
	if err := tx.Create(&model.Lesson_Question{LessonID: uint(lession_id)}).Error; err != nil {
		fmt.Println("Rollback1", err.Error())
		tx.Rollback()
	}
	if err := tx.Create(&model.Vdo{LessonID: uint(lession_id)}).Error; err != nil {
		fmt.Println("Rollback2", err.Error())
		tx.Rollback()
	}

	tx.Commit() // Commit user1
	respondJSON(w, http.StatusCreated, "บันทึก Lesson")

}

func UpdateLesson(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	lessons := vars["lessons"]
	log.Println("UpdateLesson", lessons)
	lesson := getLessonError404(db, lessons, w, r)

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
	var_lesson := vars["lessons"]
	lesson := model.Lesson{}

	if err := db.First(&lesson, model.Lesson{Header: var_lesson}).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
	}
	tx := db.Begin()
	if err := tx.Unscoped().Delete(&lesson.ID).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		tx.Rollback()
	}
	log.Println("DeleteLesson:", "(get)", var_lesson, "id lesson", lesson.ID)

	lesson_question := model.Lesson_Question{}
	if err := db.First(&lesson_question, model.Lesson_Question{LessonID: lesson.ID}).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
	}
	if err := tx.Unscoped().Delete(&lesson_question).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		tx.Rollback()
	}

	vdo := model.Vdo{}
	if err := db.First(&vdo, model.Vdo{LessonID: lesson.ID}).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
	}
	if err := tx.Unscoped().Delete(&vdo).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		tx.Rollback()
	}

	// //คำถามครั้งแรกจะตั้งเป็นค่า ตาม leson ก่อนหากมีการเพิ่มคำถามค่อยมาปรับ
	// if err := tx.Create(&model.Lesson_Question{LessonID: uint(lession_id)}).Error; err != nil {
	// 	tx.Rollback()
	// }
	// if err := tx.Create(&model.Vdo{LessonID: uint(lession_id)}).Error; err != nil {
	// 	tx.Rollback()
	// }

	// tx.Commit() // Commit user1
	// log.Println("Delete ", title, "Finish")
	// respondJSON(w, http.StatusCreated, "ลบ Lesson lesion_quiz vdo")

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

// db.Create(&model.Question_Choose{QuestionNumber: 3, QuestionName:"xxx"})
// db.Create(&model.Question{QuestionID: 3, QuestionSet:1})
// db.Create(&model.Score{ QuestionSet_score:3})
// func CreateLesson(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
// 	lesson := model.Lesson{}

// 	decoder := json.NewDecoder(r.Body)
// 	if err := decoder.Decode(&lesson); err != nil {
// 		respondError(w, http.StatusBadRequest, err.Error())
// 		return
// 	}
// 	defer r.Body.Close()

// 	if err := db.Save(&lesson).Error; err != nil {
// 		respondError(w, http.StatusInternalServerError, err.Error())
// 		return
// 	}
// 	respondJSON(w, http.StatusCreated, lesson)
// }

//lesson := model.Lesson{}
// decoder := json.NewDecoder(r.Body)
// if err := decoder.Decode(&lesson); err != nil {
// 	respondError(w, http.StatusBadRequest, err.Error())
// 	return
// }
// defer r.Body.Close()
//tx.RollbackTo("sp1")//tx.SavePoint("sp1")
// tx := db.Begin()
// do some database operations in the transaction (use 'tx' from this point, not 'db')
// if err := tx.Create(&model.Lesson{Content: "Giraffe"}).Error; err != nil {
// return any error will rollback
// fmt.Println(err.Error())
//respondError(w, http.StatusInternalServerError, err.Error())
// }

// if err := tx.Create(&model.Lesson_Question{QuestionID: 1}).Error; err != nil {
// 	fmt.Println(err.Error())
//respondError(w, http.StatusInternalServerError, err.Error())
// }

// return nil will commit the whole transaction
//return nil
// if err := db.Save(&lesson).Error; err != nil {
// 	respondError(w, http.StatusInternalServerError, err.Error())
// 	return
// }

// var users = []User{{Name: "jinzhu1"}, {Name: "jinzhu2"}, {Name: "jinzhu3"}}
// db.Create(&users)

// for _, user := range users {
// user.ID // 1,2,3
// }
