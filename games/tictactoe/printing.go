package tictactoe

import "fmt"

func (d *Tictactoegame) PrintGame() string {
	response := ""
	for r := 0; r < rows; r++ {
		for c := 0; c < columns; c++ {
			response += fmt.Sprintf("|%s", d.board.GetCellValue(r, c).Get())
		}
		response += fmt.Sprint("|")
		if r < 2 {
			response += fmt.Sprintf("\n-------")
		}
		response += fmt.Sprintf("\n")
	}
	response += fmt.Sprintf("\n")
	return response
}

func (d *Tictactoegame) PrintHelp() string {
	response := fmt.Sprintf("Example: Select the numbered spot you wish to place your piece\n")
	count := 1
	for r := 0; r < rows; r++ {
		for c := 0; c < columns; c++ {
			response += fmt.Sprintf("|%d", count)
			count++
		}
		response += fmt.Sprint("|")
		if r < 2 {
			response += fmt.Sprintf("\n-------")
		}
		response += "\n"
	}
	response += fmt.Sprintf("\n\n")
	return response
}
