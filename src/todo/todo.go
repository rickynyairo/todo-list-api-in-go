package todo

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	configs "github.com/rickynyairo/todo-list-api-in-go/src/config" 
)

type Config struct {
	*configs.Config
}

func New(configuration *configs.Config) *Config {
	return &Config{configuration}
}

func (config *Config) Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/{todoID}", config.GetATodo)
	router.Delete("/{todoID}", config.DeleteTodo)
	router.Post("/", config.CreateTodo)
	router.Get("/", config.GetAllTodos)
	return router
}

type Todo struct {
	Slug  string `json:"slug"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

func (config *Config) GetATodo(w http.ResponseWriter, r *http.Request) {
	todoID := chi.URLParam(r, "todoID")
	todos := Todo{
		Slug:  todoID,
		Title: "My fist Todo on Go",
		Body:  "Hello world from planet Neptune",
	}
	render.JSON(w, r, todos) // chi router helper for serializing JSON
}

func (config *Config) DeleteTodo(w http.ResponseWriter, r *http.Request) {
	response := make(map[string]string)
	response["message"] = "TODO deleted succesfully"
	render.JSON(w, r, response)
}

func (config *Config) CreateTodo(w http.ResponseWriter, r *http.Request) {
	response := make(map[string]string)
	response["message"] = "TODO created succesfully"
	render.JSON(w, r, response)
}

func (config *Config) GetAllTodos(w http.ResponseWriter, r *http.Request) {
	todos := []Todo{
		{
			Slug:  "slug",
			Title: "My fist Todo on Go",
			Body:  "Hello world from planet Neptune",
		},
	}
	render.JSON(w, r, todos)
}
