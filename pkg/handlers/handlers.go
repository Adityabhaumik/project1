package handlers

import (
	"log"

	"net/http"

	"github.com/Adityabhaumik/project1/pkg/config"
	"github.com/Adityabhaumik/project1/pkg/models"
	"github.com/Adityabhaumik/project1/pkg/render"
)

// template Data holds data sent from handlers to templates

var Repo *Repository

type Repository struct {
	app *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		app: a,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	// n, err := fmt.Fprintf(w, "hello world in browser from home Page")
	// if err != nil {
	// 	log.Println("error occoured", err)
	// }
	// fmt.Println(fmt.Sprintf("Number of bytes %d", n))
	log.Println("home accessed")
	remoteIp := r.RemoteAddr
	m.app.Session.Put(r.Context(), "remote_ip", remoteIp)
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	StringMap := make(map[string]string)
	StringMap["test"] = "hello world"
	remoteIp := m.app.Session.GetString(r.Context(), "remote_ip")
	StringMap["remote_ip"] = remoteIp
	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: StringMap,
	})
}

// func addValues(x int, y int) int {
// 	return x + y
// }

// func Divide(w http.ResponseWriter, r *http.Request) {
// 	f, err := divideValues(100.0, 0.0)
// 	if err != nil {
// 		fmt.Fprintf(w, "Cant divide by 0")
// 		return
// 	}
// 	fmt.Fprint(w, fmt.Sprintf("The divided value is %f", f))
// }

// func divideValues(x, y float32) (float32, error) {
// 	if y <= 0 {
// 		err := errors.New("Cant divide by 0")
// 		return 0.0, err
// 	}
// 	result := x / y
// 	return result, nil
// }
