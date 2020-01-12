package server

import (
	"classifieds-api/internal/app/delivery"
	"classifieds-api/internal/app/repository"
	"classifieds-api/internal/app/usecase"
	dbInit "classifieds-api/internal/db"
	"database/sql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
	"time"
)

type Params struct {
	Port        int64
	Url         string
	DatabaseURL string
}

type Server struct {
	Mux    *mux.Router
	Params *Params
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.Mux.ServeHTTP(w, r)
}

func NewServer(params *Params) (*Server, error) {
	server := &Server{
		Mux:    mux.NewRouter().PathPrefix(params.Url).Subrouter(),
		Params: params,
	}

	return server, nil
}

func StartApp(params Params) error {
	server, err := NewServer(&params)
	if err != nil {
		return err
	}

	db, err := newDB(params.DatabaseURL)
	if err != nil {
		return err
	}

	defer func() {
		if err := db.Close(); err != nil {
			log.Println(err)
		}
	}()

	server.ConfigureServer(db)

	portStr := strconv.FormatInt(params.Port, 10)

	return http.ListenAndServe(":"+portStr, server)
}

func newDB(connStr string) (*sql.DB, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(time.Hour)
	if err := dbInit.InitTables(db); err != nil {
		return nil, err
	}

	return db, nil
}

func (s *Server) ConfigureServer(db *sql.DB) {
	rep := repository.NewAdRepository(db)

	use := usecase.NewAdUsecase(rep)

	delivery.NewAdHandler(s.Mux, use)
}
