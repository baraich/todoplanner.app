package tui

import "github.com/charmbracelet/lipgloss"

var TaskNormalStyle = lipgloss.NewStyle().
	Padding(0, 2).
	Foreground(lipgloss.Color("245"))

var SelectedTaskStyle = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("203")).
	Padding(0, 2)

var TaskDoneStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("160")).
	Strikethrough(true)

var TaskUnDoneStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("244"))
