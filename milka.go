package main

import (
	"fmt"
	"strings"
	"time"
)

func center(s string, w int) string {
	length := len([]rune(s))
	if length >= w {
		return s
	}
	pad := w - length
	left := pad / 2
	return strings.Repeat(" ", left) + s + strings.Repeat(" ", pad-left)
}

// drawPackageBody prints the common elements: mountains, Milka text, cow, and nutrition info.
func drawPackageBody(w int, topPaddingLines int) {
	line := func(content string) {
		fmt.Println("  " + bgPurple + fgWhite + center(content, w) + reset)
	}

	leftRight := func(left, right string) {
		space := w - len([]rune(left)) - len([]rune(right))
		if space < 0 {
			space = 0
		}
		fmt.Println("  " + bgPurple + fgWhite + left + strings.Repeat(" ", space) + right + reset)
	}

	for i := 0; i < topPaddingLines; i++ {
		line("")
	}

	line("      _.-/`.              .-/`.    ")
	line("     /::::\\ \\        _.-/::::\\ \\   ")
	line("  __/::::::\\ \\______/:::::::::\\ \\_ ")
	line(" /:::::::::::::::::::::::::::::\\ \\ ")
	line("")
	fmt.Println("  " + bgPurple + boldWhite + center("M  I  L  K  A", w) + reset)
	line("")
	line("")
	line("A l p e n m i l c h")
	line("Alpine Milk")
	line("")
	line("")
	line("          ( )_( )          ")
	line("         ( o   o )         ")
	line("        ---(_)---          ")
	line("       _ / /|\\ \\ _         ")
	line("      (           )        ")
	line("")
	line("")
	leftRight("  100 g / 3.5 oz e", "|kcal|  ")
	leftRight("                  ", "| 530|  ")
	line("")
}

func drawFullPackage(w int) {
	fmt.Println()

	// Top zigzag edge (purple text, default bg)
	topEdge := strings.Repeat("▲", w)
	fmt.Println("  " + fgPurple + topEdge + reset)

	// Draw mountains, text, and cow with 5 initial blank lines
	drawPackageBody(w, 5)

	// Bottom zigzag
	bottomEdge := strings.Repeat("▼", w)
	fmt.Println("  " + fgPurple + bottomEdge + reset)
	fmt.Println("\n   Enjoy your Milka! 🐄⛰️💜")
}

func drawOpenedPackage(w int) {
	tornLines := []string{
		"  " + fgPurple + "▲▲" + reset + "  " + bgBrown + fgBrown + "___________" + reset + " " + fgPurple + strings.Repeat("▲", w-16) + reset,
		"  " + bgPurple + fgWhite + `   \` + reset + bgBrown + fgBrown + "|  _  _  _  |" + reset + bgPurple + fgWhite + `/` + strings.Repeat(" ", w-18) + reset,
		"  " + bgPurple + fgWhite + `    \` + reset + bgBrown + fgBrown + "| | || |(  |" + reset + bgPurple + fgWhite + `/` + strings.Repeat(" ", w-18) + reset,
		"  " + bgPurple + fgWhite + `     \` + reset + bgBrown + fgBrown + "||_||_| \\_|" + reset + bgPurple + fgWhite + `/` + strings.Repeat(" ", w-18) + reset,
		"  " + bgPurple + fgWhite + `      \_________` + strings.Repeat(" ", w-16) + reset,
	}

	for i := 1; i <= 5; i++ {
		fmt.Print("\033[H\033[2J") // Clear screen for next frame
		fmt.Println()

		// Print the tore foil lines depending on the animation step
		for j := 0; j < i; j++ {
			fmt.Println(tornLines[j])
		}

		// Draw mountains, text, and cow adjusting the variable padding
		drawPackageBody(w, 6-i)

		// Bottom zigzag
		bottomEdge := strings.Repeat("▼", w)
		fmt.Println("  " + fgPurple + bottomEdge + reset)

		if i == 5 {
			fmt.Println("\n   " + green + "[OK] " + reset + "You took a piece! Mmm... Delicious! 🍫😋")
		} else {
			fmt.Println("\n   Unwrapping... 🎁")
		}

		if i < 5 {
			time.Sleep(300 * time.Millisecond) // Animation speed
		}
	}
}
