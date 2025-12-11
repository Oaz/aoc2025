package day10

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var debug = false

func Log(format string, a ...any) {
	if debug {
		fmt.Printf(format, a...)
	}
}

type Machine struct {
	LightsTarget int
	LightButtons []int
	Buttons      [][]int
	Joltages     []uint
}

func parseLightsTarget(input string) int {
	content := regexp.MustCompile(`\[([#\.]+)\]`).FindStringSubmatch(input)[1]
	target := 0
	for i, c := range content {
		if c == '#' {
			target += 1 << i
		}
	}
	return target
}

func parseButtons(input string) ([][]int, []int) {
	content := regexp.MustCompile(`\((\d+(?:,\d+)*)\)`).FindAllStringSubmatch(input, -1)
	buttons := make([][]int, 0)
	lightButtons := make([]int, 0)
	for _, buttonInput := range content {
		button := make([]int, 0)
		lightButton := 0
		for _, c := range strings.Split(buttonInput[1], ",") {
			n, _ := strconv.Atoi(c)
			button = append(button, n)
			lightButton += 1 << n
		}
		buttons = append(buttons, button)
		lightButtons = append(lightButtons, lightButton)
	}
	return buttons, lightButtons
}

func parseJoltage(input string) []uint {
	content := regexp.MustCompile(`\{(\d+(?:,\d+)*)\}`).FindStringSubmatch(input)
	joltages := make([]uint, 0)
	for _, c := range strings.Split(content[1], ",") {
		x, _ := strconv.Atoi(c)
		joltages = append(joltages, uint(x))
	}
	return joltages
}

func ParseInput(input string) []Machine {
	rows := strings.Split(input, "\n")
	machines := make([]Machine, len(rows))
	for row, rowContent := range rows {
		buttons, lightsButtons := parseButtons(rowContent)
		machines[row] = Machine{
			LightsTarget: parseLightsTarget(rowContent),
			LightButtons: lightsButtons,
			Buttons:      buttons,
			Joltages:     parseJoltage(rowContent),
		}
	}
	return machines
}

func (machine Machine) LightsAfterButtonPress(lights int, press int) int {
	return lights ^ machine.LightButtons[press]
}

func (machine Machine) LightsAfterButtonPresses(presses []int) int {
	lights := 0
	for _, buttonIndex := range presses {
		lights = machine.LightsAfterButtonPress(lights, buttonIndex)
	}
	return lights
}

func (machine Machine) FindFewestButtonPressForLights() int {
	lights := []int{0}
	count := 0
	for {
		count++
		newPotentialLights := make([]int, 0)
		for _, light := range lights {
			for buttonIndex := range machine.LightButtons {
				newLight := machine.LightsAfterButtonPress(light, buttonIndex)
				if newLight == machine.LightsTarget {
					return count
				}
				newPotentialLights = append(newPotentialLights, newLight)
			}
		}
		lights = newPotentialLights
	}
}

func Part1(machines []Machine) int {
	result := 0
	for _, machine := range machines {
		result += machine.FindFewestButtonPressForLights()
	}
	return result
}

func executeGlpk(input string) (string, error) {
	tmpFile, err := os.CreateTemp("", "glpk_input*.mod")
	if err != nil {
		return "", err
	}
	defer os.Remove(tmpFile.Name())
	_, err = tmpFile.WriteString(input)
	if err != nil {
		return "", err
	}
	tmpFile.Close()
	cmd := exec.Command("glpsol", "--lp", tmpFile.Name(), "-o", "/dev/stdout")
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err = cmd.Run()
	if err != nil {
		return "", fmt.Errorf("glpsol error: %v, stderr: %s", err, stderr.String())
	}
	return out.String(), nil
}

func parseGlpkResult(output string) (float64, string, bool) {
	re := regexp.MustCompile(`Objective:\s*obj\s*=\s*([\d\.eE+-]+)\s*\((MINimum|MAXimum)\)`)

	for _, line := range strings.Split(output, "\n") {
		matches := re.FindStringSubmatch(line)
		if matches != nil && len(matches) == 3 {
			var value float64
			_, err := fmt.Sscanf(matches[1], "%f", &value)
			if err == nil {
				return value, matches[2], true
			}
		}
	}
	return 0, "", false
}

func (machine Machine) FindFewestButtonPressForJoltages() int {
	var lp strings.Builder
	lp.WriteString("Minimize\n obj:")
	for iButton := range machine.Buttons {
		lp.WriteString(fmt.Sprintf(" + p%d", iButton))
	}
	lp.WriteString("\n\nSubject To\n")
	for iJoltage, joltage := range machine.Joltages {
		lp.WriteString(fmt.Sprintf(" j%d:", iJoltage))
		for iButton, button := range machine.Buttons {
			for _, joltageButton := range button {
				if joltageButton == iJoltage {
					lp.WriteString(fmt.Sprintf(" + p%d", iButton))
				}
			}
		}
		lp.WriteString(fmt.Sprintf(" = %d\n", joltage))
	}
	lp.WriteString("\n\nBounds\n")
	for iButton := range machine.Buttons {
		lp.WriteString(fmt.Sprintf(" p%d >= 0\n", iButton))
	}
	lp.WriteString("\n\nIntegers\n")
	for iButton := range machine.Buttons {
		lp.WriteString(fmt.Sprintf(" p%d\n", iButton))
	}
	lp.WriteString("\n\nEnd\n")
	result, err := executeGlpk(lp.String())
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return 0
	}
	value, _, ok := parseGlpkResult(result)
	if !ok {
		fmt.Printf("Error: could not parse result: %s\n", result)
		return 0
	}
	return int(value)
}

func Part2(machines []Machine) int {
	result := 0
	for i, machine := range machines {
		Log("Machine %v: %v\n", i, machine)
		start := time.Now()
		fewest := machine.FindFewestButtonPressForJoltages()
		elapsed := time.Since(start)
		Log("Fewest button presses: %v in %v\n", fewest, elapsed)
		result += fewest
	}
	return result
}
