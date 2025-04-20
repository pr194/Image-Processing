package server

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/pr194/Collaborative-tool/routes"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Server struct {
	DB     *gorm.DB
	Router *http.ServeMux
}

func (s *Server) ConnectDatabase() error {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	dsn := os.Getenv("SUPABASE_DSN")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("Failed to connect to Supabase: %w", err)
	}

	fmt.Println("Connected to Supabase!")

	s.DB = db

	return nil

}

func NewServer() *Server {
	router := http.NewServeMux()

	server := &Server{

		Router: router,
	}

	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Server is running"))
	})
	routes.RegisterRoutes(router)
	return server
}

func (s *Server) Start(port string) {
	httpServer := &http.Server{
		Addr:    ":" + port,
		Handler: s.Router,
	}

	log.Println("🚀 Server is running on port " + port)
	if err := httpServer.ListenAndServe(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
