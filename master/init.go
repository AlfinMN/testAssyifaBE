package master

import (
	"assyifa/controller"
	"assyifa/repositories"
	"assyifa/usecase"
	"database/sql"

	"github.com/gorilla/mux"
)

func Init(r *mux.Router, db *sql.DB) {
	userRepo := repositories.InitUserRepoImpl(db)
	userUsecase := usecase.InitUsecaseImpl(userRepo)
	controller.UserController(r, userUsecase)
}
