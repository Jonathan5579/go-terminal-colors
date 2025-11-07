package main

/** This comes mainly from https://www.lihaoyi.com/post/BuildyourownCommandLinewithANSIescapecodes.html*/
import (
	"fmt"
	"github.com/charmbracelet/log"

	"bufio"
	"time"
	"os"
	"strings"
)

/**Constants*/
const TERMINAL_BLACK = "\033[30m"
const	TERMINAL_RED = "\033[31m"
const TERMINAL_GREEN = "\033[32m"
const TERMINAL_YELLOW = "\033[33m"
const TERMINAL_BLUE = "\033[34m"
const TERMINAL_MAGENTA = "\033[35m"
const TERMINAL_CYAN = "\033[36m"
const TERMINAL_GRAY = "\033[37m"
const TERMINAL_WHITE = "\033[97m"
const TERMINAL_RESET_COLOR = "\033[0m"
const TERMINAL_MOVE_LEFT = "\u001b[1000D" //\u001b[1000D, which means "move cursor left by 1000 characters). 
const PROGRESS_BAR_ELEMENT = "━"

func main() {

	log.Info("Surprisingly, colors in terminal output is just text surrounded by special characters:\n")
	fmt.Println("\033[31m Red \033[31m \033[32m Green \033[32m \033[33m Yellow \033[33m \033[34m Blue \033[34m \033[35m Magenta \033[35m \033[36m Cyan \033[36m \033[37m Gray \033[37m")
	fmt.Println("")
	time.Sleep(4000 * time.Millisecond)

	log.Info("Inline progress bar works by: move cursor to left, write output. Repeat:\n")
	time.Sleep(2000 * time.Millisecond)
	PausedProgress()
	fmt.Println("")

	log.Info("With this and colors you can make beautiful progress bars:\n")
	time.Sleep(3000 * time.Millisecond)
	BeautyProgressBar(TERMINAL_RED)
	BeautyProgressBar(TERMINAL_GREEN)
	BeautyProgressBar(TERMINAL_YELLOW)
	BeautyProgressBar(TERMINAL_BLUE)
	BeautyProgressBar(TERMINAL_MAGENTA)
	BeautyProgressBar(TERMINAL_CYAN)
	BeautyProgressBar(TERMINAL_GRAY)
	fmt.Println(TERMINAL_RESET_COLOR)

	log.Info("Artificially go red -> yellow -> green:\n")
	time.Sleep(3000 * time.Millisecond)
	AutomaticColorProgressBar()
	fmt.Println(TERMINAL_RESET_COLOR)

	log.Info("Or a rainbow:\n")
	time.Sleep(2000 * time.Millisecond)
	RainbowProgressBar()
}

func PausedProgress(){
	writer := bufio.NewWriter(os.Stdout)
	progressOutput := ""

	for i := 1; i<=10; i++{
		time.Sleep(300 * time.Millisecond)

		err := writeToOutput(writer, TERMINAL_MOVE_LEFT)
		if err != nil { return }
		time.Sleep(500 * time.Millisecond)
		progressOutput := fmt.Sprintf("    %d %%", i)
		err = writeToOutput(writer, progressOutput)
		if err != nil { return }
	}

	writeToOutput(writer, fmt.Sprintf("%s\n",progressOutput))
}

func BeautyProgressBar(color string){
	writer := bufio.NewWriter(os.Stdout)
	//progressBarElement := fmt.Sprintf("%[1]s━%[1]s",color)
	progressBarElement := fmt.Sprintf("%[1]s%[2]s%[1]s",color, PROGRESS_BAR_ELEMENT)
	progressOutput := ""

	for i := 1; i<=100; i++{
		time.Sleep(10 * time.Millisecond)


		repeatedString := strings.Repeat(progressBarElement, i/2)
		progressOutput = fmt.Sprintf("%s %d%% %s", TERMINAL_MOVE_LEFT, i, repeatedString)

		err := writeToOutput(writer, progressOutput)
		if err != nil { return }
	}

	writeToOutput(writer, fmt.Sprintf("%s  Done!\n",progressOutput))
}

func AutomaticColorProgressBar(){
	writer := bufio.NewWriter(os.Stdout)
	progressOutput := ""

	for i := 1; i<=100; i++{
		time.Sleep(25 * time.Millisecond)

		color:= TERMINAL_YELLOW
		if i <= 30{
			color = TERMINAL_RED
		}
		if i == 100{
			color = TERMINAL_GREEN
		}

		progressBarElement := fmt.Sprintf("%[1]s%[2]s%[1]s",color, PROGRESS_BAR_ELEMENT)
		repeatedString := strings.Repeat(progressBarElement, i/2)
		progressOutput = fmt.Sprintf("%s %d%% %s", TERMINAL_MOVE_LEFT, i, repeatedString)

		err := writeToOutput(writer, progressOutput)
		if err != nil { return }
	}

	writeToOutput(writer, fmt.Sprintf("%s  Done!\n",progressOutput))
}

func RainbowProgressBar(){
	writer := bufio.NewWriter(os.Stdout)
	progressOutput := ""

	for i := 1; i<=100; i++{
		time.Sleep(25 * time.Millisecond)

		color := auxRainbowColors(i)

		progressBarElement := fmt.Sprintf("%[1]s%[2]s%[1]s",color, PROGRESS_BAR_ELEMENT)

		repeatedString := strings.Repeat(progressBarElement, i/2)
		progressOutput = fmt.Sprintf("%s %d%% %s", TERMINAL_MOVE_LEFT, i, repeatedString)

		err := writeToOutput(writer, progressOutput)
		if err != nil { return }
	}

	writeToOutput(writer, fmt.Sprintf("%s  Done!\n",progressOutput))
}

func auxRainbowColors(value int) (result string) {
	switch {
		case value >= 0 && value <= 16:
			return TERMINAL_RED
		case value >= 17 && value <= 33:
			return TERMINAL_GREEN
		case value >= 34 && value <= 50:
			return TERMINAL_YELLOW
		case value >= 51 && value <= 67:
			return TERMINAL_CYAN
		case value >= 68 && value <= 84:
			return TERMINAL_BLUE
		case value >= 85 && value <= 99:
			return TERMINAL_MAGENTA
		case value == 100:
			return TERMINAL_GREEN
		default:
			return TERMINAL_WHITE
	}
}

func writeToOutput(writer *bufio.Writer, text string) (err error){
	_, err = writer.WriteString(text)
	if err != nil {
		fmt.Println("Error writing:", err)
		return
	}
	err = writer.Flush()
	if err != nil {
		fmt.Println("Error flushing:", err)
		return
	}
	return nil
}