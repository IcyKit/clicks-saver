package main

import (
	"fmt"
	"http-server/configs"
	"http-server/internal/auth"
	"http-server/internal/link"
	"http-server/internal/stat"
	"http-server/internal/user"
	"http-server/pkg/db"
	"http-server/pkg/event"
	"http-server/pkg/middleware"
	"net/http"
)

func App() http.Handler {
	cfg := configs.LoadConfig()
	db := db.NewDB(cfg)
	router := http.NewServeMux()
	eventBus := event.NewEventBus()

	linkRepository := link.NewLinkRepository(db)
	userRepository := user.NewUserRepository(db)
	statRepository := stat.NewStatRepository(db)

	authService := auth.NewAuthService(userRepository)
	statService := stat.NewStatService(&stat.StatServiceDeps{
		EventBus:       eventBus,
		StatRepository: statRepository,
	})

	link.NewLinkHandler(router, link.LinkHandlerDeps{
		LinkRepository: linkRepository,
		EventBus:       eventBus,
		Config:         cfg,
	})
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Config:      cfg,
		AuthService: authService,
	})
	stat.NewStatHandler(router, stat.StatHandlerDeps{
		StatRepository: statRepository,
		Config:         cfg,
	})

	go statService.AddClick()

	stack := middleware.Chain(middleware.CORS, middleware.Logging)

	return stack(router)
}

func main() {
	app := App()
	server := http.Server{
		Addr:    ":8081",
		Handler: app,
	}

	fmt.Println("Server is listening on port 8081")
	server.ListenAndServe()
}
