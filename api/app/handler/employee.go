package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"border_patrol/api/app/model"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func GetAllEmployees(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	log.Println("begin GetAllEmployees ")
	employees := []model.Employee{}
	db.Find(&employees)
	respondJSON(w, http.StatusOK, employees)
}

func CreateEmployee(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	log.Println("begin createEmployee ")
	employee := model.Employee{}
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&employee); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&employee).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	log.Println(employee, "... ", decoder)
	log.Println("Saving createEmployee ")

	respondJSON(w, http.StatusCreated, employee)
}

func GetEmployee(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	log.Println("begin GetEmployee ", name)
	employee := model.Employee{}
	if err := db.Where("name = ?", name).First(&employee).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondJSON(w, http.StatusOK, employee)
}

func UpdateEmployee(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	employee := getEmployeeError404( db, name, w, r)

	//email ไม่ต้องบันทึก ห้ามซ้ำ  คือตอนส่งใน json มา ไม่ต้องใส่ กันบันทึกพลาด
	if employee == nil {
		return
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&employee); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&employee).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, employee)
}

func DeleteEmployee(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	name := vars["name"]
	employee := getEmployeeError404( db, name, w, r)
	if employee == nil {
		return
	}
	if err := db.Delete(&employee).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusNoContent, nil)
}

func DisableEmployee(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	name := vars["name"]
	employee := getEmployeeError404( db, name, w, r)
	if employee == nil {
		return
	}
	employee.Disable()
	if err := db.Save(&employee).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, employee)
}

func EnableEmployee(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	name := vars["name"]
	employee := getEmployeeError404( db, name, w, r)
	if employee == nil {
		return
	}
	employee.Enable()
	if err := db.Save(&employee).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, employee)
}

// getEmployeeError404 gets a employee instance if exists, or respond the 404 error otherwise
func getEmployeeError404( db *gorm.DB, name string, w http.ResponseWriter, r *http.Request) *model.Employee {
	employee := model.Employee{}
	if err := db.First(&employee, model.Employee{Name: name}).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return nil
	}

	return &employee
	// if fist_select{
	// 	if err := db.First(&employee, model.Employee{Name: name}).Error; err != nil {
	// 		respondError(w, http.StatusNotFound, err.Error())
	// 		return nil
	// 	}
	// }else{
	// 	if err := db.Where("name = ?", name).Error; err != nil {
	// 		respondError(w, http.StatusNotFound, err.Error())
	// 		return nil
	// 	}
	// }

}
