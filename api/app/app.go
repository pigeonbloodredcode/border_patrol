package app

import (
	_"fmt"
	"log"
	"net/http"

	"border_patrol/api/app/handler"
	"border_patrol/api/app/model"
	"border_patrol/api/config"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"gorm.io/driver/sqlite"

)

// App has router and db instances
type App struct {
	Router *mux.Router
	DB     *gorm.DB
}



// App initialize with predefined configuration
func (a *App) Initialize(config *config.Config) {
	
	db, err := gorm.Open(sqlite.Open(config.DB.Dialect), &gorm.Config{})	
	if err != nil {
		log.Fatal("Could not connect database")
	}

	a.DB = model.DBMigrate(db)
	a.Router = mux.NewRouter()
	a.setRouters()
}

// Set all required routers
func (a *App) setRouters() {
	// Routing for handling the projects  EMPLOYEES
	a.Get("/employees", a.GetAllEmployees)
	a.Post("/employees", a.CreateEmployee)
	a.Get("/employees/{title}", a.GetEmployee)
	a.Put("/employees/{title}", a.UpdateEmployee)
	a.Delete("/employees/{title}", a.DeleteEmployee)
	a.Put("/employees/{title}/disable", a.DisableEmployee)
	a.Put("/employees/{title}/enable", a.EnableEmployee)



	// Routing for handling the projects VDOS 
	a.Get("/vdo", a.GetAllVdos)
	a.Post("/vdo", a.CreateVdo)
	a.Get("/vdo/{title}", a.GetVdo)
	a.Put("/vdo/{title}", a.UpdateVdo)
	a.Delete("/vdo/{title}", a.DeleteVdo)
	a.Put("/vdo/{title}/disable", a.DisableVdo)
	a.Put("/vdo/{title}/enable", a.EnableVdo)



}







// Handlers to manage Employee Data
func (a *App) GetAllEmployees(w http.ResponseWriter, r *http.Request) {
	handler.GetAllEmployees(a.DB, w, r)
}

func (a *App) CreateEmployee(w http.ResponseWriter, r *http.Request) {
	handler.CreateEmployee(a.DB, w, r)
}

func (a *App) GetEmployee(w http.ResponseWriter, r *http.Request) {
	handler.GetEmployee(a.DB, w, r)
}

func (a *App) UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	handler.UpdateEmployee(a.DB, w, r)
}

func (a *App) DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	handler.DeleteEmployee(a.DB, w, r)
}

func (a *App) DisableEmployee(w http.ResponseWriter, r *http.Request) {
	handler.DisableEmployee(a.DB, w, r)
}

func (a *App) EnableEmployee(w http.ResponseWriter, r *http.Request) {
	handler.EnableEmployee(a.DB, w, r)
}



// Handlers to manage Vdo Data
func (a *App) GetAllVdos(w http.ResponseWriter, r *http.Request) {
	handler.GetAllVdos(a.DB, w, r)
}

func (a *App) CreateVdo(w http.ResponseWriter, r *http.Request) {
	handler.CreateVdo(a.DB, w, r)
}

func (a *App) GetVdo(w http.ResponseWriter, r *http.Request) {
	handler.GetVdo(a.DB, w, r)
}

func (a *App) UpdateVdo(w http.ResponseWriter, r *http.Request) {
	handler.UpdateVdo(a.DB, w, r)
}

func (a *App) DeleteVdo(w http.ResponseWriter, r *http.Request) {
	handler.DeleteVdo(a.DB, w, r)
}

func (a *App) DisableVdo(w http.ResponseWriter, r *http.Request) {
	handler.DisableVdo(a.DB, w, r)
}

func (a *App) EnableVdo(w http.ResponseWriter, r *http.Request) {
	handler.EnableVdo(a.DB, w, r)
}





/////////////////////////////////////////////////////////////////////////////////////////
// Wrap the router for GET method
func (a *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("GET")
}

// Wrap the router for POST method
func (a *App) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("POST")
}

// Wrap the router for PUT method
func (a *App) Put(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("PUT")
}

// Wrap the router for DELETE method
func (a *App) Delete(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("DELETE")
}



// Run the app on it's router
func (a *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.Router))
}






































	// dbURI := fmt.Sprintf("%s:%s@/%s?charset=%s&parseTime=True",	// 	config.DB.Username,	// 	config.DB.Password,	// 	config.DB.Name,	// 	config.DB.Charset)		//db, err := gorm.Open(config.DB.Dialect, dbURI) 