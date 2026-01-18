package application

import (
	"fmt"
	"github.com/marveldo/gogin/internal/application/handlers"
	"github.com/marveldo/gogin/internal/application/repository"
	"github.com/marveldo/gogin/internal/application/routes"
	"github.com/marveldo/gogin/internal/application/services"
	"github.com/marveldo/gogin/internal/config"
	"github.com/marveldo/gogin/internal/db"

)

func Setup() {
	cfg := config.LoadConfig()
	dab, err := db.Setup(cfg)
	if err != nil {panic(fmt.Sprintf("Failed with error : %v", err))}
	err = dab.AutoMigrate(db.Get_db_models()...)
	if err != nil {panic(fmt.Sprintf("Migration failed with error : %v", err))}
	t_r := repository.TesterRepository{DB : dab}
	t_s := services.TesterService{R : &t_r}
    
	e := routes.GetEngine()
	handlers.NewTestHandler(e, &t_s)
    e.Run()
}