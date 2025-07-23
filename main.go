package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

var (
	passIcon    = lipgloss.NewStyle().SetString("✔").Foreground(lipgloss.Color("42"))
	failIcon    = lipgloss.NewStyle().SetString("✖").Foreground(lipgloss.Color("9"))
	runLabel    = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("7"))
	testStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("14"))
	timeStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("10"))
	assertStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("8"))
	errorStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("1")).Bold(true)
)

func main() {
	maxRuns := flag.Int("max", 0, "Maximum number of runs before stopping (0 for unlimited)")
	help := flag.Bool("help", false, "Show usage information")
	flag.Parse()

	if *help {
		fmt.Println("Usage: flaky [options] <test_filter>")
		fmt.Println("Options:")
		fmt.Println("  --max <n>   Maximum number of runs before stopping (0 for unlimited)")
		fmt.Println("  --help      Show this help message")
		os.Exit(0)
	}

	if flag.NArg() == 0 {
		fmt.Println(errorStyle.Render("No flaky test specified"))
		return
	}

	filter := flag.Arg(0)
	runCommandUntilFail(*maxRuns, "php", "artisan", "test", "-d memory_limit=-6144M", "--filter="+filter)
}

func runCommandUntilFail(maxRuns int, command string, args ...string) {
	i := 1

	for {
		cmd := exec.Command(command, args...)

		var out bytes.Buffer
		var stderr bytes.Buffer
		cmd.Stdout = &out
		cmd.Stderr = &stderr

		cmd.Run()
		output := out.String() + stderr.String()

		testName := extractTestName(output)
		duration := extractDuration(output)
		assertions := extractAssertions(output)

		if testName == "" {
			testName = "UnknownTest"
		}
		if duration == "" {
			duration = "?.??s"
		}
		if assertions == "" {
			assertions = "unknown assertions"
		}

		label := runLabel.Render(fmt.Sprintf("Run %d:", i))
		test := testStyle.Render(testName)
		time := timeStyle.Render(duration)
		assert := assertStyle.Render(fmt.Sprintf("(%s)", assertions))

		if strings.Contains(output, "FAIL") {
			fmt.Printf("%s %s %s %s %s\n", failIcon, label, test, time, assert)
			break
		} else {
			fmt.Printf("%s %s %s %s %s\n", passIcon, label, test, time, assert)
			i++
		}

		if maxRuns > 0 && i > maxRuns {
			fmt.Println(runLabel.Render("Reached max run limit, stopping."))
			break
		}
	}
}

func extractTestName(output string) string {
	re := regexp.MustCompile(`(?m)^\s*PASS\s+(Tests\\[^\r\n]+)`)
	match := re.FindStringSubmatch(output)
	if len(match) > 1 {
		return strings.ReplaceAll(match[1], `\`, `/`)
	}
	return ""
}

func extractDuration(output string) string {
	re := regexp.MustCompile(`(?m)^\s*Duration:\s+([\d.]+s)`)
	match := re.FindStringSubmatch(output)
	if len(match) > 1 {
		return match[1]
	}
	return ""
}

func extractAssertions(output string) string {
	re := regexp.MustCompile(`(?m)\((\d+)\s+assertions\)`)
	match := re.FindStringSubmatch(output)
	if len(match) > 1 {
		return match[1] + " assertions"
	}
	return ""
}
