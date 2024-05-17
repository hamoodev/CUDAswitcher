package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	cudaBasePath := "/usr/local"

	availableVersions, err := getAvailableCUDAVersions(cudaBasePath)
	if err != nil {
		fmt.Printf("Error getting available CUDA versions: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Available CUDA versions:")
	for i, version := range availableVersions {
		fmt.Printf("%d. %s\n", i+1, version)
	}

	fmt.Print("Select a CUDA version: ")
	var choice int
	_, err = fmt.Scanf("%d", &choice)
	if err != nil || choice < 1 || choice > len(availableVersions) {
		fmt.Println("Invalid choice")
		os.Exit(1)
	}

	// Get the selected version
	version := availableVersions[choice-1]

	// Define the path to the Fish config file
	fishConfigPath := filepath.Join(os.Getenv("HOME"), ".config", "fish", "config.fish")

	// Open the Fish config file for reading
	file, err := os.Open(fishConfigPath)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	// Read the contents of the file line by line
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "/usr/local/cuda-") {
			// Replace the old CUDA version with the new version in the relevant lines
			if strings.Contains(line, "set -gx PATH") {
				line = fmt.Sprintf("set -gx PATH /usr/local/cuda-%s/bin $PATH", version)
			} else if strings.Contains(line, "set -gx LD_LIBRARY_PATH") {
				line = fmt.Sprintf("set -gx LD_LIBRARY_PATH /usr/local/cuda-%s/lib64 $LD_LIBRARY_PATH", version)
			}
		}
		lines = append(lines, line)
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		os.Exit(1)
	}

	// Write the updated contents back to the file
	outputFile, err := os.Create(fishConfigPath)
	if err != nil {
		fmt.Printf("Error creating file: %v\n", err)
		os.Exit(1)
	}
	defer outputFile.Close()

	writer := bufio.NewWriter(outputFile)
	for _, line := range lines {
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			fmt.Printf("Error writing to file: %v\n", err)
			os.Exit(1)
		}
	}
	err = writer.Flush()
	if err != nil {
		fmt.Printf("Error flushing writer: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("CUDA version switched to:", version)
}

// getAvailableCUDAVersions returns a list of available CUDA versions found in the given base path.
func getAvailableCUDAVersions(basePath string) ([]string, error) {
	var versions []string

	files, err := os.ReadDir(basePath)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		if file.IsDir() && strings.HasPrefix(file.Name(), "cuda-") {
			version := strings.TrimPrefix(file.Name(), "cuda-")
			versions = append(versions, version)
		}
	}

	return versions, nil
}
