// Copyright 2024 Daytona Platforms Inc.
// SPDX-License-Identifier: Apache-2.0

package target

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/daytonaio/daytona/pkg/serverapiclient"
	view_util "github.com/daytonaio/daytona/pkg/views/util"
)

type item struct {
	target serverapiclient.ProviderTarget
}

func (i item) Title() string { return *i.target.Name }
func (i item) Description() string {
	if i.target.ProviderInfo != nil {
		return *i.target.ProviderInfo.Name
	}
	return ""
}
func (i item) FilterValue() string { return *i.target.Name }

type model struct {
	list   list.Model
	choice *serverapiclient.ProviderTarget
	footer string
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "q", "ctrl+c":
			return m, tea.Quit

		case "enter":
			i, ok := m.list.SelectedItem().(item)
			if ok {
				m.choice = &i.target
			}
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		h, v := view_util.DocStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return view_util.DocStyle.Render(m.list.View() + m.footer)
}
