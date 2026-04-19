package main

import (
	"fmt"
	"strings"
)

const (
	// ANSI color codes
	reset      = "\033[0m"
	boldBrown  = "\033[1;33m"
	darkBrown  = "\033[38;5;94m"
	lightBrown = "\033[38;5;130m"
	bgBrown    = "\033[48;5;94m"
	bgDark     = "\033[48;5;52m"
)

// Chocolate bar dimensions
const (
	cols   = 6 // number of chocolate pieces per row
	rows   = 4 // number of chocolate piece rows
	pieceW = 6 // width of each piece in chars
	pieceH = 3 // height of each piece in lines
)

func drawChocolate() {
	fmt.Println()
	fmt.Printf("%s  🍫  Welcome to get-choco! Here's your chocolate bar:  🍫%s\n", boldBrown, reset)
	fmt.Println()

	totalWidth := cols*pieceW + cols + 1

	// Top border
	fmt.Printf("%s%s%s\n", darkBrown, "▄"+strings.Repeat("▄", totalWidth-1), reset)

	for row := 0; row < rows; row++ {
		// Top of piece row (divider line or top)
		printDividerLine(row, totalWidth)

		// Middle lines of each piece
		for line := 0; line < pieceH-1; line++ {
			printPieceLine()
		}
	}

	// Bottom row closure
	printDividerLine(rows, totalWidth)

	// Bottom border
	fmt.Printf("%s%s%s\n", darkBrown, "▀"+strings.Repeat("▀", totalWidth-1), reset)

	fmt.Println()
	fmt.Printf("%s  Enjoy your chocolate! 🍫%s\n", lightBrown, reset)
	fmt.Println()
}

// printDividerLine prints the horizontal divider between chocolate rows.
func printDividerLine(row, totalWidth int) {
	if row == 0 || row == rows {
		// outer horizontal border
		fmt.Printf("%s|%s%s|%s\n", darkBrown, bgBrown+strings.Repeat(" ", totalWidth-2)+reset, darkBrown, reset)
		return
	}
	// inner divider with piece notches
	line := darkBrown + "|" + reset
	for c := 0; c < cols; c++ {
		line += bgDark + strings.Repeat("_", pieceW) + reset
		if c < cols-1 {
			line += darkBrown + "|" + reset
		}
	}
	line += darkBrown + "|" + reset
	fmt.Println(line)
}

// printPieceLine prints a single content line across all pieces in a row.
func printPieceLine() {
	line := darkBrown + "|" + reset
	for c := 0; c < cols; c++ {
		line += bgBrown + strings.Repeat(" ", pieceW) + reset
		if c < cols-1 {
			line += darkBrown + "|" + reset
		}
	}
	line += darkBrown + "|" + reset
	fmt.Println(line)
}

func main() {
	drawChocolate()
}
