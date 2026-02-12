package application

import (
	"fmt"
	"log"

	"github.com/marveldo/gogin/internal/application/handlers"
	"github.com/marveldo/gogin/internal/application/repository"
	"github.com/marveldo/gogin/internal/application/routes"
	"github.com/marveldo/gogin/internal/application/services"
	"github.com/marveldo/gogin/internal/application/tasks"
	"github.com/marveldo/gogin/internal/application/validator"
	"github.com/marveldo/gogin/internal/config"
	"github.com/marveldo/gogin/internal/db"
)

func Setup() {
	cfg := config.LoadConfig()
	dab, err := db.Setup(cfg)
	if err != nil {
		panic(fmt.Sprintf("Failed with error : %v", err))
	}
	err = dab.AutoMigrate(db.Get_db_models()...)
	if err != nil {
		panic(fmt.Sprintf("Migration failed with error : %v", err))
	}
	validator.RegisterAllValidators()
	//Asynq Server

	redis_opt := tasks.GetRedisOptions(cfg)
	client := tasks.CreateAsynqClient(redis_opt)
	server := tasks.GetAsynqServer(redis_opt)
	t_r := repository.TesterRepository{DB: dab}
	t_s := services.TesterService{R: &t_r, C: client}
	u_r := repository.Userrespository{DB: dab}
	u_s := services.Userservice{R: &u_r, C: client}
	a_r := repository.AuthorRepository{DB: dab}
	b_r := repository.BookRepository{DB: dab}
	b_s := services.BookService{B: &b_r, A: &a_r}
	a_s := services.AuthorService{R: &a_r}
	c_r := repository.CartRepository{DB: dab}
	c_s := services.CartService{R: &c_r, U: &u_r}
	w_r := repository.WaitlistRepository{DB: dab}
	w_s := services.WaitlistService{W: &w_r, U: &u_r, B: &b_r}

	e := routes.GetEngine()
	handlers.NewTestHandler(e, &t_s)
	handlers.NewUserHandler(e, &u_s)
	handlers.NewBookHandler(e, &b_s)
	handlers.NewAuthorHandler(e, &a_s)
	handlers.NewCartHandler(e, &c_s)
	handlers.NewWaitlistHandler(e, &w_s)

	asynq_chan := make(chan string)
	gin_chan := make(chan string)
	go func() {
		if err := server.Run(&tasks.Taskhandler{}); err != nil {
			asynq_chan <- fmt.Sprintf("asynq server error: %v", err)
		}
	}()
	go func() {
		if err := e.Run(); err != nil {
			gin_chan <- "Gin server Failed to run"
		}
	}()

	for msg := range asynq_chan {
		log.Fatal(msg)
	}
	for msg := range gin_chan {
		log.Fatal(msg)
	}

	//SHUTDOWN

}
