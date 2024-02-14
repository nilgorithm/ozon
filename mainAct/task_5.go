package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Folder struct {
	Dir     string   `json:"dir"`
	Files   []string `json:"files"`
	Folders []Folder `json:"folders"`
}

func countInfectedFiles(folder Folder, isInfected bool) int {
	count := 0
	for _, file := range folder.Files {
		if isInfected || strings.HasSuffix(file, ".hack") {
			count++
		}
	}
	for _, subfolder := range folder.Folders {
		count += countInfectedFiles(subfolder, isInfected || containsHack(subfolder.Files))
	}
	return count
}

func containsHack(files []string) bool {
	for _, file := range files {
		if strings.HasSuffix(file, ".hack") {
			return true
		}
	}
	return false
}

//func main() {
//	reader := bufio.NewReader(os.Stdin)
//	line, _ := reader.ReadString('\n')
//	t, _ := strconv.Atoi(strings.TrimSpace(line))
//
//	for i := 0; i < t; i++ {
//		line, _ = reader.ReadString('\n')
//		n, _ := strconv.Atoi(strings.TrimSpace(line))
//
//		var jsonLines string
//		for j := 0; j < n; j++ {
//			line, _ = reader.ReadString('\n')
//			jsonLines += line
//		}
//
//		var root Folder
//		json.Unmarshal([]byte(jsonLines), &root)
//
//		fmt.Println(countInfectedFiles(root, containsHack(root.Files)))
//	}
//}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	t, _ := strconv.Atoi(scanner.Text())

	for i := 0; i < t; i++ {
		scanner.Scan()
		n, _ := strconv.Atoi(scanner.Text())

		var jsonLines strings.Builder
		for j := 0; j < n; j++ {
			scanner.Scan()
			jsonLines.WriteString(scanner.Text())
		}

		var root Folder
		json.Unmarshal([]byte(jsonLines.String()), &root)

		fmt.Println(countInfectedFiles(root, containsHack(root.Files)))
	}
}
