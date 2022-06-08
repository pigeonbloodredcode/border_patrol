package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	//"border_patrol/api/app"
	"border_patrol/api/app/model"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func GetAllLessons(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	fmt.Println("GetAllLessons")

	type LV struct { //lesson  and vdo
		Id      	string `json:"id"`
		Header  	string `json:"header"`
		Content 	string `json:"content"`
		Title   	string `json:"title"`
		Src_dir 	string `json:"src_dir"`
		Director 	string `json:"director"`
		Status   	string `json:"status"`
		VdoLession_id 	string `json:"vdos.lesson_id"`
		QuestionLession_id 	string `json:"lesson_questions.lesson_id"`



	}

	showSelect := "lessons.id, header, content,title, src_dir, director, status, vdos.lesson_id, lesson_questions.lesson_id "
	rows, err := db.Table("lessons").Select(showSelect).Joins(
		"JOIN vdos ON lessons.id = vdos.lesson_id "	).Joins(
		"JOIN lesson_questions ON lessons.id = lesson_questions.lesson_id "	).Rows()
	if err != nil {
		fmt.Println(err.Error())
	}
	defer rows.Close()

	var id, header, content, title, src_dir, director, status, vdolession_id, questionlession_id string
	var lvs []LV
	for rows.Next() {
		fmt.Println("rows",rows)
		err := rows.Scan(&id ,&header, &content, &title, &src_dir, &director, &status, &vdolession_id, questionlession_id)
		if err != nil {
			fmt.Println("if err",err)
		}
		fmt.Println(id, "................................................................")
		
		lvs = append(lvs, LV{
			Id: id,  
			Header: header, 
			Content: content, 
			Title: title, 
			Src_dir: src_dir,
			Director: director,
			Status: status,
			VdoLession_id: vdolession_id,
			QuestionLession_id: questionlession_id,
		})

	}

	 response, err := json.Marshal(lvs)

	// //fmt.Println("respondJSON", response, "status 200:", http.StatusOK, lv)
	if err != nil {
		//log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(response))

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

	var lesson = model.Lesson{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&lesson); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())

	}
	defer r.Body.Close()

	//	fmt.Println("","decode", decoder, "rBody",r.Body)

	if err := db.Save(&lesson).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
	}

	fmt.Println("lessonID", lesson.ID)

	tx := db.Begin()
	//คำถามครั้งแรกจะตั้งเป็นค่า ตาม leson ก่อนหากมีการเพิ่มคำถามค่อยมาปรับ
	if err := tx.Create(&model.Lesson_Question{LessonID: uint(lesson.ID)}).Error; err != nil {
		fmt.Println("Rollback1", err.Error())
		tx.Rollback()
	}
	if err := tx.Create(&model.Vdo{LessonID: uint(lesson.ID)}).Error; err != nil {
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
	log.Println("DeleteLesson:", "(get)")

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
	log.Println(var_lesson, "id lesson", lesson.ID)

	lesson_question := model.Lesson_Question{}
	if err := db.First(&lesson_question, model.Lesson_Question{LessonID: lesson.ID}).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
	}
	if err := tx.Unscoped().Delete(&lesson_question).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		tx.Rollback()
	}

	tx.Commit() // Commit user1
	// vdo := model.Vdo{}
	// if err := db.First(&vdo, model.Vdo{LessonID: lesson.ID}).Error; err != nil {
	// 	respondError(w, http.StatusInternalServerError, err.Error())
	// }
	// if err := tx.Unscoped().Delete(&vdo).Error; err != nil {
	// 	respondError(w, http.StatusInternalServerError, err.Error())
	// 	tx.Rollback()
	// }

	// //คำถามครั้งแรกจะตั้งเป็นค่า ตาม leson ก่อนหากมีการเพิ่มคำถามค่อยมาปรับ
	// if err := tx.Create(&model.Lesson_Question{LessonID: uint(lession_id)}).Error; err != nil {
	// 	tx.Rollback()
	// }
	// if err := tx.Create(&model.Vdo{LessonID: uint(lession_id)}).Error; err != nil {
	// 	tx.Rollback()
	// }

	// log.Println("Delete ", title, "Finish")
	// respondJSON(w, http.StatusCreated, "ลบ Lesson lesion_quiz vdo")

}

// getLessonerror404 gets a lesson instance if exists, or respond the 404 error otherwise
func getLessonError404(db *gorm.DB, header string, w http.ResponseWriter, r *http.Request) *model.Lesson {
	lesson := model.Lesson{}
	if err := db.First(&lesson, model.Lesson{Header: header}).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return nil
	}
	return &lesson
}

//fmt.Println(rows)
//t1 := db.Exec("SELECT * FROM vdos,lessons")
//result := map[string]interface{}{}
//db.Model(&model.Lesson{}).Find(&result)//db.Model(&model.Lesson{}).First(&result, "id = ?", 1)

//respondJSON(w, http.StatusOK, lvs)

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
// if err := db.Create(&lesson).Error; err != nil {
// 	fmt.Println(err.Error())
// }
// var lession_id int
// for _, data2 := range lesson {
// 	lession_id = int(data2.ID)
// }
// fmt.Println("lessionID", lession_id)

// db, err := gorm.Open(sqlite.Open("test4.db"), &gorm.Config{})
// if err != nil {
// 	panic("failed to connect database")
// }

// db.AutoMigrate(&User{}, &Project{})

// if err := db.SetupJoinTable(&Project{}, "Users", &ProjectUser{}); err != nil {
// 	println(err.Error())
// 	panic("Failed to setup join table")
// }

// bogdan := User { Name: "Bogdan" }

// db.Create(&bogdan)

// db.Create(&Project{Name: "Test1", Users: []User {{ID: bogdan.ID}}})
//db.Model(&model.Lesson{}).Select(" *  ").Joins("left join vdos on vdos.lesson_id = lessons.id").Scan(&model.Vdo{}).Where(1)
//db.Table("lessons").Select(showSelect).Joins("LEFT JOIN `vdos` ON vdos.lesson_id = lessons.id").Find(&lv)
//vdo := []model.Vdo{}
//lesson := []model.Lesson{}
//db.Table("lessons").Select(showSelect).Joins("LEFT JOIN `vdos` ON vdos.lesson_id = lessons.id").Scan(&lv)
//db.Table("lessons").Select(showSelect).Joins("JOIN vdos ON lessons.id = vdos.lesson_id ").Scan(&lv)
//datas, err := db.Table("lessons").Select(" * ").Joins("left join vdos on vdos.lesson_id = lessons.id").Rows()
//datas := db.Model(&model.Lesson{}).Select(" *  ").Joins("left join vdos on vdos.lesson_id = lessons.id").Scan(&model.Vdo{}).Where(true)
// datas := db.Model(&model.Lesson{}).Select(" lessons.id ").Joins("left join vdos on vdos.lesson_id = lessons.id").Scan(&model.Vdo{}).Where(true)

// if datas.Error == nil {
// 	jsonData, _ := json.Marshal(datas)
// 	fmt.Println(string(jsonData))
// }else {
// 	fmt.Println(datas.Error)
// }

//rows, err := db.Table("users").Select("users.name, emails.email").Joins("left join emails on emails.user_id = users.id").Rows()
// for _, data := range datas {
// 	fmt.Println(data)
// }

//lessons := []model.Lesson{}
//db.Joins("Lesson_Question").Joins("Vdo").Find(&lessons, "lesson_id IN ?", []int{7,8})
// lessons := []model.Lesson{}
//  db.Find(&lessons)
// respondJSON(w, http.StatusOK, lessons)

// var lv LV
// if err := db.Table("lessons").Select(" lessons.idc,lessons.header, vdos.lesson_id ").Joins("left join `vdos` on vdos.lesson_id = lessons.id").Find(&lv).Error; err != nil{
// 	respondJSON(w, http.StatusOK, err.Error())
// }
