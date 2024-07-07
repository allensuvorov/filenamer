package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"unicode"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter LeetCode problem name to produce a nicely formatted file name: ")
	if text, err := reader.ReadString('\n'); err == nil {
		formattedFileName := formatFileName(text)
		fmt.Println(formattedFileName)
		err := copyToClipboard(formattedFileName)
		if err != nil {
			fmt.Println("Failed to copy to clipboard:", err)
		} else {
			fmt.Println("File name copied to clipboard.")
		}
	}
}

func formatFileName(s string) string {
	s = strings.ToLower(s)
	s = strings.TrimSpace(s)
	s = underScore(s)
	s = s + ".go"
	return s
}

func underScore(s string) string {
	b := []rune{}
	inSymbols := false

	for _, v := range s {

		if !inSymbols {
			if isSymbol(v) { // enter symbols
				b = append(b, '_')
				inSymbols = true
			} else {
				b = append(b, v)
			}
		} else {
			if !isSymbol(v) { // exit symbols
				b = append(b, v)
				inSymbols = false
			}
		}
	}
	return string(b)
}

func isSymbol(r rune) bool {
	return !unicode.IsDigit(r) && !unicode.IsLetter(r)
}

func copyToClipboard(text string) error {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "darwin":
		cmd = exec.Command("pbcopy")
	case "linux":
		cmd = exec.Command("xclip", "-selection", "clipboard")
	case "windows":
		cmd = exec.Command("cmd", "/c", "clip")
	default:
		return fmt.Errorf("unsupported platform")
	}

	in, err := cmd.StdinPipe()
	if err != nil {
		return err
	}

	if err := cmd.Start(); err != nil {
		return err
	}

	_, err = in.Write([]byte(text))
	if err != nil {
		return err
	}

	if err := in.Close(); err != nil {
		return err
	}

	return cmd.Wait()
}
