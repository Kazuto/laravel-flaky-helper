package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

func main() {
	// Parse command-line arguments
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
		fmt.Println("No flaky test specified")
		return
	}

	filter := flag.Arg(0)

	runCommandUntilFail(*maxRuns, "php", "artisan", "test", "-d memory_limit=-6144M", "--filter="+filter)
}

func runCommandUntilFail(maxRuns int, command string, args ...string) {
	red := "\033[31m"
	green := "\033[32m"
	gray := "\033[37m"
	reset := "\033[0m"

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

		if strings.Contains(output, "FAIL") {
			fmt.Printf("%s%s Run %d: %s %s (%s, failed)%s\n", red, gray, i, testName, duration, assertions, reset)
			break
		} else {
			fmt.Printf("%s%s Run %d: %s %s (%s)%s\n", green, gray, i, testName, duration, assertions, reset)
			i++
		}

		if maxRuns > 0 && i > maxRuns {
			fmt.Println("Reached max run limit, stopping.")
			break
		}
	}
}

func extractTestName(output string) string {
	// Match: PASS  Tests\Support\Something\TestName
	re := regexp.MustCompile(`(?m)^\s*PASS\s+(Tests\\[^\r\n]+)`)
	match := re.FindStringSubmatch(output)
	if len(match) > 1 {
		// Replace backslashes with slashes
		return strings.ReplaceAll(match[1], `\`, `/`)
	}
	return ""
}

func extractDuration(output string) string {
	// Match: Duration: 3.51s
	re := regexp.MustCompile(`(?m)^\s*Duration:\s+([\d.]+s)`)
	match := re.FindStringSubmatch(output)
	if len(match) > 1 {
		return match[1]
	}
	return ""
}

func extractAssertions(output string) string {
	// Match line like: Tests:    2 passed (299 assertions)
	re := regexp.MustCompile(`(?m)\((\d+)\s+assertions\)`)
	match := re.FindStringSubmatch(output)
	if len(match) > 1 {
		return match[1] + " assertions"
	}
	return ""
}
