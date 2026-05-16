package main

import (
	"apiserver/internal/core/apiserver"
	"apiserver/internal/core/repository/postgres"
	sessions_repository "apiserver/internal/features/sessions/repository"
	sessions_service "apiserver/internal/features/sessions/service"
	sessions_transport "apiserver/internal/features/sessions/transport"
	users_repository "apiserver/internal/features/users/repository"
	users_service "apiserver/internal/features/users/service"
	users_transport "apiserver/internal/features/users/transport"
	"log"
	"net/http"
)

func main() {
	store := postgres.NewStore()
	if err := store.Open(); err != nil {
		log.Fatal(err)
	}

	defer store.Close()

	usersRepository := users_repository.NewRepository(store)
	usersService := users_service.NewService(usersRepository)
	usersTransport := users_transport.NewTransport(usersService)

	sessionsRepository := sessions_repository.NewRepository(store)
	sessionsService := sessions_service.NewService(usersRepository, sessionsRepository)
	sessionsTransport := sessions_transport.NewTransport(sessionsService)

	router := http.NewServeMux()
	router.HandleFunc("POST /users", usersTransport.CreateHandler())
	router.HandleFunc("GET /users/{email}", usersTransport.FindByEmailHandler())
	router.HandleFunc("POST /sessions", sessionsTransport.CreateSessionHandler())

	server := apiserver.NewServer(router)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
