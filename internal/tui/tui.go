package tui

import (
	"log"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/evertras/bubble-table/table"

	"crypto/internal/gate"
	"crypto/internal/stex"
)

// magenta code #FF00FF

var styleBase = lipgloss.NewStyle().Foreground(lipgloss.Color("#a7a")).BorderForeground(lipgloss.Color("#a38")).Align(lipgloss.Right)

type Model struct {
	gateTable table.Model
	stexTable table.Model

	gateChan chan *gate.GateInfo
	stexChan chan *stex.StexInfo

	gatePriceInfo gate.GateInfo
	stexPriceInfo stex.StexInfo

	gateData []*gate.GateInfo
	stexData []*stex.StexInfo
}

func NewModel(gateChan chan *gate.GateInfo,
	stexChan chan *stex.StexInfo,
	gatePriceInfo gate.GateInfo,
	stexPriceInfo stex.StexInfo) Model {

	return Model{
		gateTable:     table.New(generateGateColumns()),
		stexTable:     table.New(generateStexColumns()),
		gateChan:      gateChan,
		stexChan:      stexChan,
		gatePriceInfo: gatePriceInfo,
		stexPriceInfo: stexPriceInfo,
	}
}

// Generate columns based on how many are critical to show some summary
func generateGateColumns() []table.Column {

	return []table.Column{
		table.NewColumn("Ask Price", "Gate Ask Price", 15),
		table.NewColumn("Ask Amount", "Gate Ask Amount", 15),
		table.NewColumn("Bid Price", "Gate Bid Price", 15),
		table.NewColumn("Bid Amount", "Gate Bid Amount", 15),
	}
}

func generateStexColumns() []table.Column {

	return []table.Column{
		table.NewColumn("Ask Price", "Stex Ask Price", 15),
		table.NewColumn("Ask Amount", "Stex Ask Amount", 15),
		table.NewColumn("Bid Price", "Stex Bid Price", 15),
		table.NewColumn("Bid Amount", "Stex Bid Amount", 15),
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	// m.table, cmd = m.table.Update(msg)
	cmds = append(cmds, cmd)

	select {
	case gateInfo := <-m.gateChan:

		m.gateData = []*gate.GateInfo{gateInfo}

		m.gateTable = m.gateTable.WithRows(generateGateRowsFromData(m.gateData)).WithColumns(generateGateColumns()).WithStaticFooter("Gate").
			BorderRounded().
			WithBaseStyle(styleBase).
			Focused(true)

		cmds = append(cmds, func() tea.Msg {
			return m.gatePriceInfo
		})

	case stexInfo := <-m.stexChan:
		// fmt.Println(stexOrder.Ask[0])
		m.stexData = []*stex.StexInfo{stexInfo}

		m.stexTable = m.stexTable.WithRows(generateStexRowsFromData(m.stexData)).WithColumns(generateStexColumns()).WithStaticFooter("Stex").
			BorderRounded().
			WithBaseStyle(styleBase).
			Focused(true)
		cmds = append(cmds, func() tea.Msg {
			return m.stexPriceInfo
		})
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc", "q":
			cmds = append(cmds, tea.Quit)
		}
	}

	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	body := strings.Builder{}

	pad := lipgloss.NewStyle().Padding(1)

	body.WriteString(pad.Render(m.gateTable.View()))
	body.WriteString(pad.Render(m.stexTable.View()))

	return body.String()
}

func generateGateRowsFromData(data []*gate.GateInfo) []table.Row {
	rows := []table.Row{}

	for _, entry := range data {
		row := table.NewRow(table.RowData{
			"Ask Price":  entry.AskPrice,
			"Ask Amount": entry.AskAmount,
			"Bid Price":  entry.BidPrice,
			"Bid Amount": entry.BidAmount,
		})

		rows = append(rows, row)
	}

	return rows
}

func generateStexRowsFromData(data []*stex.StexInfo) []table.Row {
	rows := []table.Row{}

	for _, entry := range data {
		row := table.NewRow(table.RowData{
			"Ask Price":  entry.AskPrice,
			"Ask Amount": entry.AskAmount,
			"Bid Price":  entry.BidPrice,
			"Bid Amount": entry.BidAmount,
		})

		rows = append(rows, row)
	}

	return rows
}

func RunTUI(gateChan chan *gate.GateInfo,
	stexChan chan *stex.StexInfo,
	gatePriceInfo gate.GateInfo,
	stexPriceInfo stex.StexInfo) {

	p := tea.NewProgram(NewModel(gateChan, stexChan, gatePriceInfo, stexPriceInfo))

	if err := p.Start(); err != nil {
		log.Fatal(err)
	}
}
