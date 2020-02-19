package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var tablines, lines []string
	var blamefile, blameline string
	tabix := 0
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) < 1 {
			continue
		}
		if strings.HasPrefix(line, "\t") {
			line = line[1:]
			if tabix == 0 {
				fields := strings.Split(line, ":")
				wd, _ := os.Getwd()
				ix := strings.Index(fields[0], wd)
				if ix >= 0 {
					line = line[len(wd)+1:]
					fields = strings.Split(line, ":")
					blamefile = fields[0]
					fields = strings.Split(fields[1], " ")
					blameline = fields[0]
				}
			}
			tabix++
			tablines = append(tablines, line)
		} else {
			lines = append(lines, line)
		}
	}
	if err := scanner.Err(); err != nil {
		//fmt.Fprintln(os.Stderr, "reading standard input:", err)
		os.Exit(1)
	}
	tabix = 0
	for _, line := range lines {
		if strings.HasSuffix(line, ")") {
			if tabix < len(tablines) {
				line = fmt.Sprintf(" %-30s %s", line, tablines[tabix])
				tabix++
			}
		}
		fmt.Println(line)
	}
	if blameline != "" && blamefile != "" {
		params := []string{"blame", "-L", blameline + "," + blameline, blamefile}
		out, err := exec.Command("git", params...).CombinedOutput()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("BLAME", string(out))
	}
}
