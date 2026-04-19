package main

import (
"bufio"
"fmt"
"os"
"os/signal"
"strings"
"syscall"
"time"
)

const (
reset       = "\033[0m"
bgPurple    = "\033[48;5;135m" // Lighter Lilac
fgPurple    = "\033[38;5;135m"
fgWhite     = "\033[38;5;255m"
boldWhite   = "\033[1;97m"

// Colors for the chocolate inside
bgBrown     = "\033[48;5;94m"
fgBrown     = "\033[38;5;52m"

greenCheck  = "\033[38;5;34m✓\033[0m"
green = "\033[38;5;34m"
red = "\033[38;5;160m"
)

func printLoading(message string) {
	fmt.Print(message)
	for i := 0; i < 3; i++ {
		time.Sleep(500 * time.Millisecond)
		fmt.Print(".")
	}
	fmt.Println(" " + greenCheck + green + " [OK]" + reset)
}

func main() {
	// Ignore SIGINT completely
	signal.Ignore(os.Interrupt, syscall.SIGTERM)
	signal.Ignore(os.Interrupt, syscall.SIGINT)
	signal.Ignore(os.Interrupt, syscall.SIGQUIT)

	// Clear the terminal screen before rendering
	fmt.Print("\033[H\033[2J")
	printLoading("Initializing affection protocols")
	printLoading("Setting up sweetness modules")
	printLoading("Loading Milka package")
	printLoading("Almost there")
	fmt.Print("\033[H\033[2J")

	w := 54
	drawFullPackage(w)

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Do you want to (t)ake a piece or (i)gnore it? [t/i]: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			return
		}

		input = strings.TrimSpace(strings.ToLower(input))
		if input == "t" {
			// Clear again and draw opened output
			fmt.Print("\033[H\033[2J")
			drawOpenedPackage(w)
			break
		} else {
			fmt.Println(red + "[KO] " + reset + "You tried to ignore it, but the Milka is too tempting! 🐄💜")
		}
	}
}
