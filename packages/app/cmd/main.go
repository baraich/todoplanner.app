package main

import (
	"io"
	"log/slog"
	"net"

	"github.com/baraich/todoplanner.app/base/resource"
	"github.com/charmbracelet/log"
	"github.com/charmbracelet/ssh"
	"github.com/charmbracelet/wish"
)

func main() {
	s, err := wish.NewServer(
		wish.WithAddress(net.JoinHostPort("0.0.0.0", "2222")),
		wish.WithHostKeyPEM([]byte(resource.Resource.TodoPlannerKey.Private)),
		wish.WithMiddleware(
			func(next ssh.Handler) ssh.Handler {
				return func(s ssh.Session) {
					io.WriteString(s, "Hello, World!\n")
					next(s)
				}
			},
		),
	)

	if err != nil {
		log.Fatal("Error: Failed to start the server: ", err)
	}

	log.Info("Starting the SSH server...")
	if err := s.ListenAndServe(); err != nil {
		log.Fatal("Error: Failed to listen and serve: ", err)
	}
	slog.Info("Shutting down the SSH server...")
}
