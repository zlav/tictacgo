package tictactoe

func (d *Tictactoegame) checkWinner(r, c int) {
	player := d.players[d.currPlayer]
	for i := 0; i < columns; i++ {
		if (i != c) && d.board.GetCellValue(r, i) != player.GetIcon() {
			break
		}
		if i == columns-1 {
			d.winner = player
			return
		}
	}

	for i := 0; i < rows; i++ {
		if (i != r) && d.board.GetCellValue(i, c) != player.GetIcon() {
			break
		}
		if i == rows-1 {
			d.winner = player
			return
		}
	}

	if r == c {
		for i := 0; i < rows; i++ {
			if (i != r) && d.board.GetCellValue(i, i) != player.GetIcon() {
				break
			}
			if i == rows-1 {
				d.winner = player
				return
			}
		}
	}

	if r+c == rows-1 {
		for i := 0; i < rows; i++ {
			if (i != r) && d.board.GetCellValue(i, (rows-i)-1) != player.GetIcon() {
				break
			}
			if i == rows-1 {
				d.winner = player
				return
			}
		}
	}
}
