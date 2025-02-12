package main

import (
	"log/slog"
	"net"

	"github.com/baraich/todoplanner.app/base/resource"
	"github.com/baraich/todoplanner.app/base/tui"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/log"
	"github.com/charmbracelet/ssh"
	"github.com/charmbracelet/wish"
	"github.com/charmbracelet/wish/activeterm"
	"github.com/charmbracelet/wish/bubbletea"
)

func main() {
	s, err := wish.NewServer(
		wish.WithAddress(net.JoinHostPort("0.0.0.0", "2222")),
		wish.WithHostKeyPEM([]byte(resource.Resource.TodoPlannerKey.Private)),
		wish.WithMiddleware(
			bubbletea.Middleware(teaHandler),
			activeterm.Middleware(),
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

func teaHandler(s ssh.Session) (tea.Model, []tea.ProgramOption) {
	model := tui.InitModel()
	return model, []tea.ProgramOption{tea.WithAltScreen()}
}
