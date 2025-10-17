package main

import (
	"fmt"
	"os"
	"regexp"
	"slices"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	keyPressError  bool
	calculateError bool
	hexNumber      string
	result         string
}

func hexCharToDecimal(hexChar rune) (int, error) {
	switch {
	case hexChar >= '0' && hexChar <= '9':
		return int(hexChar - '0'), nil
	case hexChar >= 'a' && hexChar <= 'f':
		return int(hexChar-'a') + 10, nil
	default:
		return 0, fmt.Errorf("invalid hexadecimal character: %c", hexChar)
	}
}

func decimalToBinary(n int) string {
	if n == 0 {
		return "0000"
	}

	var bitSegments []string
	for n > 0 {
		remainder := n % 2
		bitSegments = append(
			bitSegments,
			fmt.Sprintf("%d", remainder),
		)
		n /= 2
	}

	slices.Reverse(bitSegments)

	binary := strings.Join(bitSegments, "")
	// Pad with leading zeros to 4 digits
	return fmt.Sprintf("%04s", binary)
}

func (m model) handleCalculate() (tea.Model, tea.Cmd) {
	if len(m.hexNumber) == 0 {
		m.calculateError = true
		return m, nil
	}

	var accumulator string = "0."
	for _, char := range m.hexNumber {
		decimalValue, _ := hexCharToDecimal(char)
		binaryValue := decimalToBinary(decimalValue)
		accumulator += binaryValue
	}

	m.result = accumulator

	return m, nil
}

func (m model) handleDelete() (tea.Model, tea.Cmd) {
	if len(m.hexNumber) > 0 {
		m.hexNumber = m.hexNumber[0 : len(m.hexNumber)-1]
	}
	return m, nil
}

func (m model) Init() tea.Cmd {
	return nil
}

var hexRegexp = regexp.MustCompile(`^(?:0x)?[0-9a-fA-F]+$`)

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	hasCalculatedResult := len(m.result) > 0

	switch msg := msg.(type) {

	case tea.KeyMsg:
		// reset app error state on every update
		m.calculateError = false
		m.keyPressError = false

		if msg.Type == tea.KeyCtrlC {
			return m, tea.Quit
		}

		if msg.Type == tea.KeyBackspace || msg.Type == tea.KeyDelete {
			return m.handleDelete()
		}

		// guard condition blocking remaining controls if we have calculated result
		if hasCalculatedResult {
			return m, nil
		}

		if msg.Type == tea.KeyEnter {
			return m.handleCalculate()
		}

		// handle adding a pressed key to a hex number
		key := msg.String()
		if hexRegexp.MatchString(key) {
			m.hexNumber += key
			return m, nil
		}

		// if we pressed an unmatched key, we show an error message
		m.keyPressError = true
		return m, nil

	}

	return m, nil
}

func (m model) View() string {
	if len(m.result) > 0 {
		return fmt.Sprintf("Fractional hex Number 0.%s converted to binary is equal to %s \n\nPress 'ctrl+c' to exit.", m.hexNumber, m.result)
	}

	s := fmt.Sprintf("\nEnter a valid fractional hex number: \n\nHex Number: 0.%s_ \n\n", m.hexNumber)

	if m.keyPressError {
		s += fmt.Sprintf("Please use a valid key! (0 to 9 and A to F)")
	}

	if m.calculateError {
		s += fmt.Sprintf("Please enter a hexadecimal number before calculating.")
	}

	s += "\nPress 'ctrl+c' to quit or 'Enter' to submit.\n"

	return s
}

func main() {
	program := tea.NewProgram(model{})

	if _, err := program.Run(); err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}
}
