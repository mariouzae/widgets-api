package route

import (
	"net/http"
	"widgets-api/app/auth"
	"widgets-api/app/controller"
	"widgets-api/app/util"

	"github.com/julienschmidt/httprouter"
)

var APIVersion = "/api/v1"

type Server struct {
	r *httprouter.Router
}

// Load all routes
func Load() http.Handler {
	route := httprouter.New()

	// handler for serving static files
	route.ServeFiles("/static/*filepath", http.Dir(util.GetRootPath()+"/../src/widgets-api/template"))

	// Index
	route.GET("/", controller.Index)

	// Users
	route.GET("/users", controller.Users)

	// Widgets
	route.GET("/widgets", controller.Widgets)

	// User authentication
	route.POST(APIVersion+"/login", controller.Login)

	// Retrieve all users
	route.GET(APIVersion+"/users", auth.AuthHandler(controller.GetUsers))

	// Retrieve user by id
	route.GET(APIVersion+"/users/:id", auth.AuthHandler(controller.GetUserById))

	// Retrieve all widgets
	route.GET(APIVersion+"/widgets", auth.AuthHandler(controller.GetWidgets))

	// Retrieve widget by id
	route.GET(APIVersion+"/widgets/:id", auth.AuthHandler(controller.GetWidgetById))

	// Create a new widget
	route.POST(APIVersion+"/widgets", auth.AuthHandler(controller.CreateWidget))

	// Update an existent widget by id
	route.PUT(APIVersion+"/widgets/:id", auth.AuthHandler(controller.UpdateWidget))
	return &Server{route}
}

// To avoid CORS issues always allow any origin
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	s.r.ServeHTTP(w, r)
}
