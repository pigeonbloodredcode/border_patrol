package model

import (
	_ "gorm.io/driver/sqlite"
	"gorm.io/gorm"
	// "github.com/jinzhu/gorm"
	//_"gorm.io/driver/mysql"
	//"gorm.io/driver/sqlite"
	//_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Employee struct {
	////EmployeeID        uint           `gorm:"primaryKey"`
	gorm.Model
	Name   string `gorm:"unique" json:"name"`
	City   string `json:"city"`
	Email  string `gorm:"unique" json:"email"`
	Status bool   `json:"status"`

	Admin []Admin `gorm:"foreignKey:EmployeeID"`
	Score []Score `gorm:"foreignKey:EmployeeID"`
	////Lesson []Lesson `gorm:"foreignKey:EmployeeID"`
}
type Admin struct {
	////AdminID        uint           `gorm:"primaryKey"`
	gorm.Model
	Name       string `json:"name"`
	EmployeeID uint
	Role       []Role `gorm:"foreignKey:AdminID"`
}
type Role struct {
	////RoleID        uint           `gorm:"primaryKey"`
	gorm.Model
	Role        string `json:"role"`
	Access_poin string `json:"access_poin"`
	AdminID     uint
}

type Lesson struct {
	////LessonID        uint           `gorm:"primaryKey"`
	gorm.Model
	Header  string `gorm:"unique" json:"header"`
	Content string `json:"content"`

	Lesson_Question []Lesson_Question `gorm:"foreignKey:LessonID"`
	Vdo             []Vdo             `gorm:"foreignKey:LessonID"`
}

type Lesson_Question struct {
	////Lesson_QuestionID        uint           `gorm:"primaryKey"`
	gorm.Model
	LessonID   uint
	QuestionID uint
}

type Question struct {
	gorm.Model
	QuestionID            uint    `gorm:"primaryKey"`
	QuestionSet           uint    //สำหรับบอกกลุ่มข้อสอบ
	NameQuestion_question string  `json:"namequestion"` //ชื่อกลุ่มคำถาม
	Score                 []Score `gorm:"foreignKey:QuestionSet_score;references:NameQuestion_question;"`
}
type Question_Choose struct {
	/////Question_ChooseID        uint           `gorm:"primaryKey"`
	gorm.Model
	QuestionNumber int16  `json:"questionNumber"`
	QuestionName   string `json:"questionName"`
	Answer1        string `json:"answer1"`
	Answer2        string `json:"answer2"`
	Answer3        string `json:"answer3"`
	Answer4        string `json:"answer4"`
	TrueAnswer     string `json:"trueanswer"`

	Question []Question `gorm:"foreignKey:QuestionID;references:ID;"`
	//Question      []Question      `gorm:"foreignKey:QuestionID; references:ID;"`
	//Question []Question  `gorm:"foreignKey:QuestionID;references:ID;"`

}

type Score struct {
	////ScoreID        uint           `gorm:"primaryKey"`
	gorm.Model
	QuestionSet_score uint //มีความสัมพันกับ  Question ดึงกลุ่มคำถามแล้วแสดงผลคะแนน ของแต่ละคน
	EmployeeID        uint //มีความสัมพันกับ เจ้าหน้าที่ ดึงว่าคนในที่นี้ ทำมาคะแนนเท่าไร
}

type Vdo struct {
	////VdoID    uint           `gorm:"primaryKey"`
	gorm.Model
	LessonID uint

	Title    string `json:"title"`   //
	Src_dir  string `json:"src_dir"` //folder ID
	Director string `json:"director"`
	Status   bool   `json:"status"`
}

func (e *Employee) Disable() {
	e.Status = false
}
func (p *Employee) Enable() {
	p.Status = true
}
func (e *Vdo) Disable() {
	e.Status = false
}
func (p *Vdo) Enable() {
	p.Status = true
}

// DBMigrate will create and migrate the tables, and then make the some relationships if necessary
func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(
		&Employee{}, &Vdo{}, &Admin{}, &Role{}, &Vdo{}, &Score{}, &Lesson{},
		&Lesson_Question{}, &Question_Choose{}, &Question{},
	)

	return db
}
