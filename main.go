package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"text/tabwriter"
)

type Process struct {
	Command string
	PID     string
	User    string
	Address string
}

func main() {
	cmd := exec.Command("lsof", "-iTCP", "-sTCP:LISTEN", "-P", "-n")
	output, err := cmd.Output()

	if err != nil {
		fmt.Println("✅ No open ports found or an error occurred while running lsof.", err)
		return
	}

	lines := strings.Split(string(output), "\n")
	var processes []Process

	for i, line := range lines {
		if i == 0 || len(strings.TrimSpace(line)) == 0 {
			continue
		}

		fields := strings.Fields(line)
		if len(fields) >= 9 {
			pid := fields[1]
			address := fields[len(fields)-1]
			if address == "(LISTEN)" && len(fields) > 1 {
				address = fields[len(fields)-2]
			}

			// Get full command using ps
			cmdArgs := exec.Command("ps", "-p", pid, "-o", "command=")
			fullCmdOutput, err := cmdArgs.Output()
			command := fields[0]
			if err == nil {
				fullCmd := strings.TrimSpace(string(fullCmdOutput))
				if fullCmd != "" {
					command = fullCmd
				}
			}

			proc := Process{
				Command: command,
				PID:     pid,
				User:    fields[2],
				Address: address,
			}
			processes = append(processes, proc)
		}
	}

	if len(processes) == 0 {
		fmt.Println("No processes found to terminate.")
		return
	}

	fmt.Println("\n List of currently running ports")
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)

	// Print header
	_, _ = fmt.Fprintln(w, "No\tCommand\tPID\tUser\tAddress")
	_, _ = fmt.Fprintln(w, "--\t-------\t---\t----\t-------")

	for i, p := range processes {
		cmdDisplay := p.Command
		if len(cmdDisplay) > 60 {
			cmdDisplay = cmdDisplay[:30] + "..." + cmdDisplay[len(cmdDisplay)-27:]
		}
		_, _ = fmt.Fprintf(w, "%d\t%s\t%s\t%s\t%s\n", i+1, cmdDisplay, p.PID, p.User, p.Address)
	}
	_ = w.Flush()

	fmt.Print("\n Select a number to terminate (Cancel: 0): ")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	choice, err := strconv.Atoi(input)
	if err != nil || choice < 1 || choice > len(processes) {
		if choice == 0 {
			fmt.Println("Operation cancelled.")
		} else {
			fmt.Println("Invalid selection.")
		}
		return
	}

	target := processes[choice-1]
	fmt.Printf("\nTerminating process [%s] (PID: %s)...\n", target.Command, target.PID)
	killCmd := exec.Command("kill", "-9", target.PID)
	if err := killCmd.Run(); err != nil {
		fmt.Printf("❌ Failed: %v\n", err)
		fmt.Println("💡 Tip: If it's a permission issue, try running with 'sudo'.")
	} else {
		fmt.Println("💀 Successfully terminated!")
	}
}
