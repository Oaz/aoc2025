package helpers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

type AocSettings struct {
	sessionCookie string
	baseURL       string
}

func Aoc() *AocSettings {
	session := os.Getenv("AOC_SESSION")
	year := os.Getenv("AOC_YEAR")

	return &AocSettings{
		sessionCookie: session,
		baseURL:       fmt.Sprintf("https://adventofcode.com/%s/day", year),
	}
}

func (settings *AocSettings) ReadPuzzleInputForTest(t *testing.T, day int) string {
	result, err := settings.ReadPuzzleInput(day)
	if err != nil {
		t.Fatalf("Failed to read puzzle input: %v", err)
	}
	return strings.TrimSpace(result)
}

func (settings *AocSettings) ReadPuzzleInput(day int) (string, error) {
	filename, err := DefineInputFileName(day)
	if err != nil {
		return "", err
	}
	if _, err := os.Stat(filename); err == nil {
		return ReadTextFileFullContent(filename)
	}
	url := fmt.Sprintf("%s/%d/input", settings.baseURL, day)

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}
	req.AddCookie(&http.Cookie{
		Name:  "session",
		Value: settings.sessionCookie,
	})
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("HTTP %d: %s", resp.StatusCode, string(body))
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	content := string(body)
	return WriteTextFileFullContent(filename, content)
}

func DefineInputFileName(day int) (string, error) {
	_, currentFile, _, ok := OuterCaller()
	if !ok {
		return "", fmt.Errorf("could not determine current file location")
	}
	currentDir := filepath.Dir(currentFile)
	filename := filepath.Join(currentDir, fmt.Sprintf("day%02d.txt", day))
	return filename, nil
}

func ReadTextFileFullContent(filename string) (string, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func WriteTextFileFullContent(filename string, content string) (string, error) {
	dir := filepath.Dir(filename)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return "", err
	}
	err := os.WriteFile(filename, []byte(content), 0644)
	if err != nil {
		return "", err
	}
	return content, nil
}
