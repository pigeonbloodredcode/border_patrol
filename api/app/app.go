package app

import (
	_ "fmt"
	"log"
	"net/http"

	"border_patrol/api/app/handler"
	"border_patrol/api/app/model"
	"border_patrol/api/config"

	"github.com/gorilla/mux"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
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
	a.Get("/employees/{name}", a.GetEmployee)
	a.Post("/employees", a.CreateEmployee)
	a.Put("/employees/{name}", a.UpdateEmployee)
	a.Delete("/employees/{name}", a.DeleteEmployee)
	a.Put("/employees/{name}/disable", a.DisableEmployee)
	a.Put("/employees/{name}/enable", a.EnableEmployee)



	// Routing for handling the projects VDOS 
	a.Get("/vdos", a.GetAllVdos)
	a.Get("/vdos/{title}", a.GetVdo)
	a.Post("/vdos", a.CreateVdo)
	a.Put("/vdos/{title}", a.UpdateVdo)
	a.Delete("/vdos/{title}", a.DeleteVdo)
	a.Put("/vdos/{title}/disable", a.DisableVdo)
	a.Put("/vdos/{title}/enable", a.EnableVdo)


	// Routing for handling the projects question_chooses
	a.Get("/question_chooses", a.GetAllQuestion_Chooses)
	a.Get("/question_chooses/{question_number}", a.GetQuestion_Choose)
	a.Post("/question_chooses", a.CreateQuestion_Choose)
	a.Put("/question_chooses/{question_number}", a.UpdateQuestion_Choose)
	a.Delete("/question_chooses/{question_number}", a.DeleteQuestion_Choose)




	// Routing for handling the projects Lessons 
	a.Get("/lessons", a.GetAllLessons)
	a.Get("/lessons/{title}", a.GetLesson)
	a.Post("/lessons", a.CreateLesson)
	a.Put("/lessons/{title}", a.UpdateLesson)
	a.Delete("/lessons/{title}", a.DeleteLesson)



}






//////////////////////////////////////////////////////////////////////////////////
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


// Handlers to manage Question_Choose Data
func (a *App) GetAllQuestion_Chooses(w http.ResponseWriter, r *http.Request) {
	handler.GetAllQuestion_Chooses(a.DB, w, r)
}
func (a *App) CreateQuestion_Choose(w http.ResponseWriter, r *http.Request) {
	handler.CreateQuestion_Choose(a.DB, w, r)
}
func (a *App) GetQuestion_Choose(w http.ResponseWriter, r *http.Request) {
	handler.GetQuestion_Choose(a.DB, w, r)
}
func (a *App) UpdateQuestion_Choose(w http.ResponseWriter, r *http.Request) {
	handler.UpdateQuestion_Choose(a.DB, w, r)
}
func (a *App) DeleteQuestion_Choose(w http.ResponseWriter, r *http.Request) {
	handler.DeleteQuestion_Choose(a.DB, w, r)
}

// Handlers to manage Lesson Data
func (a *App) GetAllLessons(w http.ResponseWriter, r *http.Request) {
	handler.GetAllLessons(a.DB, w, r)
}
func (a *App) CreateLesson(w http.ResponseWriter, r *http.Request) {
	handler.CreateLesson(a.DB, w, r)
}
func (a *App) GetLesson(w http.ResponseWriter, r *http.Request) {
	handler.GetLesson(a.DB, w, r)
}
func (a *App) UpdateLesson(w http.ResponseWriter, r *http.Request) {
	handler.UpdateLesson(a.DB, w, r)
}
func (a *App) DeleteLesson(w http.ResponseWriter, r *http.Request) {
	handler.DeleteLesson(a.DB, w, r)
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