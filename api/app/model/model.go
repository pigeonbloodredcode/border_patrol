package model

import (
	"gorm.io/gorm"
	// "github.com/jinzhu/gorm"
	//_"gorm.io/driver/mysql"
	//"gorm.io/driver/sqlite"
	//_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Employee struct {
	gorm.Model
	//EmployeeID uint   `gorm:"primaryKey" json:"employeeID"`
	Name   string `gorm:"unique" json:"name"`
	City   string `json:"city"`
	Email  string `gorm:"unique" json:"email"`
	Status bool   `json:"status"`

	Admin  []Admin  `gorm:"foreignKey:EmployeeID"`
	Lesson []Lesson `gorm:"foreignKey:EmployeeID"`
	Score  []Score  `gorm:"foreignKey:EmployeeID"`
}
type Admin struct {
	gorm.Model
	Name       string `json:"name"`
	EmployeeID uint
	Role       []Role `gorm:"foreignKey:AdminID"`
}
type Role struct {
	gorm.Model
	Role        string `json:"role"`
	Access_poin string `json:"access_poin"`
	AdminID     uint
}

type Lesson struct {
	gorm.Model
	Content string `json:"content"`
	Point   int32  `json:"point"`
	Comment string `json:"comment"`

	EmployeeID uint

	Lesson_Question []Lesson_Question `gorm:"foreignKey:LessonID"`
	Lesson_Vdo      []Lesson_Vdo      `gorm:"foreignKey:LessonID"`
}

type Lesson_Question struct {
	LessonID   uint
	QuestionID uint
}

type Question struct {
	SetQuestion  int16  `json:"setquestion"`
	NameQuestion string `json:"namequestion"`

	Question_ChooseID uint

	//Score           []Score           `gorm:"foreignKey:QuestionID"`//
	//Lesson_Question []Lesson_Question `gorm:"foreignKey:QuestionID"`//
}

type Question_Choose struct {
	QuestionNumber int16
	QuestionName   string `json:"questionName"`
	Answer1        string `json:"answer1"`
	Answer2        string `json:"answer2"`
	Answer3        string `json:"answer3"`
	Answer4        string `json:"answer4"`
	TrueAnswer     string `json:"trueanswer"`

	//Question []Question  `gorm:"foreignKey:Question_ChooseID"`


}

type Score struct {
	
	QuestionID uint
	EmployeeID uint
}

type Lesson_Vdo struct {
	LessonID uint
	VdoID    int
}

type Vdo struct {
	gorm.Model
	Title    string `gorm:"unique" json:"title"`
	Detail   string `json:"detail"`
	Director int    `json:"director"`
	Status   bool   `json:"status"`

	Lesson_Vdo []Lesson_Vdo `gorm:"foreignKey:VdoID"`
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
	db.AutoMigrate(&Employee{})
	db.AutoMigrate(&Vdo{})
	db.AutoMigrate(&Admin{})
	db.AutoMigrate(&Role{})
	db.AutoMigrate(&Lesson_Vdo{})
	db.AutoMigrate(&Vdo{})
	db.AutoMigrate(&Score{})
	db.AutoMigrate(&Lesson{})

	db.AutoMigrate(&Question{})
	db.AutoMigrate(&Lesson_Question{})
	db.AutoMigrate(&Question_Choose{})

	return db
}
