package tui

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type TaskDefinition struct {
	title string
	done  bool
}

type model struct {
	cursor int
	tasks  []TaskDefinition
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q":
			return nil, tea.Quit

		case "down", "j":
			if m.cursor < len(m.tasks)-1 {
				m.cursor++
			}

		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		case " ":
			m.tasks[m.cursor].done = !m.tasks[m.cursor].done

		default:
			return m, nil
		}
	}
	return m, nil
}

func (m model) View() string {
	var content strings.Builder
	for i, task := range m.tasks {
		var cursor string
		if i == m.cursor {
			cursor = "> "
		} else {
			cursor = "  "
		}

		var taskActionContent string
		if task.done {
			taskActionContent = cursor + "[x] "
		} else {
			taskActionContent = cursor + "[ ] "
		}

		if i == m.cursor {
			if task.done {
				content.WriteString(
					SelectedTaskStyle.Render(
						TaskDoneStyle.Render(taskActionContent + task.title),
					),
				)
			} else {
				content.WriteString(SelectedTaskStyle.Render(taskActionContent + task.title))
			}
		} else {
			if task.done {
				content.WriteString(TaskDoneStyle.Render(taskActionContent + task.title))
			} else {
				content.WriteString(TaskNormalStyle.Render(taskActionContent + task.title))
			}
		}
		content.WriteString("\n")
	}

	return content.String()
}

func InitModel() model {
	return model{
		cursor: 0,
		tasks: []TaskDefinition{
			{title: "Take printout for Children's day tickets.", done: false},
			{title: "Go for a walk.", done: false},
			{title: "Study for upcomming exams.", done: false},
		},
	}
}
