package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"os/exec"
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
	gray := "\033[37m" // Resets color back to default

	i := 1

	for {
		cmd := exec.Command(command, args...)

		var out bytes.Buffer
		var stderr bytes.Buffer
		cmd.Stdout = &out
		cmd.Stderr = &stderr

		cmd.Run()
		output := out.String() + stderr.String() // Combine stdout and stderr

		if strings.Contains(output, "FAIL") {
			fmt.Printf("%s%s Run %d failed\033[0m\n", red, gray, i)
			break
		} else {
			fmt.Printf("%s%s Run %d passed\033[0m\n", green, gray, i)
			i++
		}

		if maxRuns > 0 && i > maxRuns {
			fmt.Println("Reached max run limit, stopping.")
			break
		}
	}
}
