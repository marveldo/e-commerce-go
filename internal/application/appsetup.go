package application

import (
	"fmt"

	"github.com/marveldo/gogin/internal/application/handlers"
	"github.com/marveldo/gogin/internal/application/repository"
	"github.com/marveldo/gogin/internal/application/routes"
	"github.com/marveldo/gogin/internal/application/services"
	"github.com/marveldo/gogin/internal/application/validator"
	"github.com/marveldo/gogin/internal/config"
	"github.com/marveldo/gogin/internal/db"
)

func Setup() {
	cfg := config.LoadConfig()
	dab, err := db.Setup(cfg)
	if err != nil {panic(fmt.Sprintf("Failed with error : %v", err))}
	err = dab.AutoMigrate(db.Get_db_models()...)
	if err != nil {panic(fmt.Sprintf("Migration failed with error : %v", err))}
	validator.RegisterAllValidators()
	t_r := repository.TesterRepository{DB : dab}
	t_s := services.TesterService{R : &t_r}
	u_r := repository.Userrespository{DB: dab}
	u_s := services.Userservice{R: &u_r}
	a_r := repository.AuthorRepository{DB: dab}
	a_s := services.AuthorService{R : &a_r , U : &u_r}
    
	e := routes.GetEngine()
	handlers.NewTestHandler(e, &t_s)
	handlers.NewUserHandler(e, &u_s )
	handlers.NewAuthorHandler(e, &a_s)
    e.Run()
}