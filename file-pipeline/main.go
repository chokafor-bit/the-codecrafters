package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

// CodeCrafters — Operation Gopher Protocol
// Module: File Pipeline
// Author: [peace oigoga]
// Squad:  [Peace Oigoga, Faith Ejembi, Nzekwe Chinaza, Treasure Gabriel, Owulo Celebrate, Ortse Alphonsus, Chekwube Okafor, Emaikwu Godwin, Endurance Ochefije, Chubiyojo Akoh]

// ═══════════════════════════════════════════
// SQUAD PIPELINE CONTRACT
// Squad: [Bench Markers]
// ───────────────────────────────────────────
// Input line types:
//   [line 1: ALL CAP]
//   [line 2: lower]
//   [line 3: Trimspace]
//   [line 4: TODO ]
//
// Transformation rules (in order):
//   1. [Trim all leading and trailing whitespace ]
//   2. [Replace TODO: with ✦ ACTION]
//   3. [Convert ALL CAPS lines to Title Case ]
//   4. [Convert all lowercase lines to uppercase ]
//   5. [ Remove lines that are only dashes or blanks ]
//
// Output format:
//   [Header: yes/no — SENTINEL FIELD REPORT — PROCESSED]
//   [Line numbering format- 1]
//
// Terminal summary fields:
//   Lines read    : [number]
//    Lines written : [number]
//   Lines removed : [number]
//    Rules applied :
//   1. [Trim all leading and trailing whitespace ]
//   2. [Replace TODO: with ✦ ACTION]
//   3. [Convert ALL CAPS lines to Title Case ]
//   4. [Convert all lowercase lines to uppercase ]
//   5. [Remove lines that are only dashes or blanks ]
// ══════════════════════════════════════════

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run . input.txt output.txt")
		return
	}

	input := os.Args[1]
	output := os.Args[2]

	if input == output {
		fmt.Println("Error: input and output cannot be the same file")
		return
	}

	paperTrail(input, output)
}

func paperTrail(input string, output string) {
	file, err := os.Open(input)
	if err != nil {
		fmt.Println("File not found:", input)
		return
	}
	defer file.Close()
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	var result []string
	removed := 0

	for _, line := range lines {
		line = strings.TrimSpace(line)
		isDashOnly, _ := regexp.MatchString(`^[-]+$`, line)
		if line == "" || isDashOnly {
			removed++
			continue
		}
		if strings.HasPrefix(strings.ToLower(line), "todo:") {
			line = "ACTION" + line[5:]
		}
		if line == strings.ToUpper(line) && line != strings.ToLower(line) {
			line = strings.Title(strings.ToLower(line))
		} else if line == strings.ToLower(line) && line != strings.ToUpper(line) {
			line = strings.ToUpper(line)
		}

		result = append(result, line)
	}
	out, err := os.Create(output)
	if err != nil {
		fmt.Println("Cannot create output file")
		return
	}
	defer out.Close()

	writer := bufio.NewWriter(out)
	writer.WriteString("— SENTINEL FIELD REPORT — PROCESSED\n")

	for i, line := range result {
		writer.WriteString(fmt.Sprintf("%d. %s\n", i+1, line))
	}
	writer.Flush()

	fmt.Println("Done.")
	fmt.Printf("Lines read    : %d\n", len(lines))
	fmt.Printf("Lines written : %d\n", len(result))
	fmt.Printf("Lines removed : %d\n", removed)
	fmt.Println("Rules applied :")
	fmt.Println("1. Trim all leading and trailing whitespace")
	fmt.Println("2. Replace TODO: with  ACTION")
	fmt.Println("3. Convert ALL CAPS lines to Title Case")
	fmt.Println("4. Convert all lowercase lines to uppercase")
	fmt.Println("5. Remove lines that are only dashes or blanks")
}
