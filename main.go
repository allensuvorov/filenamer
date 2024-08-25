package main

import (
	"fmt"
	"os/exec"
	"runtime"
	"strings"
	"unicode"
)

func main() {
	// Read from clipboard
	clipboardText, err := readFromClipboard()
	if err != nil {
		fmt.Println("Error reading from clipboard:", err)
		return
	}

	if clipboardText == "" {
		fmt.Println("Error - empty clipboard. Copy the text to clipboard and run this app again.")
		return
	}

	formattedFileName := formatFileName(clipboardText)
	fmt.Println(formattedFileName)
	err = copyToClipboard(formattedFileName)
	if err != nil {
		fmt.Println("Failed to copy to clipboard:", err)
	} else {
		fmt.Println("File name copied to clipboard.")
	}
}

func readFromClipboard() (string, error) {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "darwin":
		cmd = exec.Command("pbpaste")
	case "linux":
		cmd = exec.Command("xclip", "-selection", "clipboard", "-output")
	case "windows":
		cmd = exec.Command("cmd", "/c", "clip")
	default:
		return "", fmt.Errorf("unsupported platform")
	}

	out, err := cmd.Output()
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(out)), nil
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
