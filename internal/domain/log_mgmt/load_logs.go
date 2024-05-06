package log_mgmt

import (
	"bufio"
	"log/slog"
	"os"
)

/*
 *  Function reads logs one at a time and sends them through a log channel for concurrent processing
 */
func LoadLogs(filePath string, logChan chan<- string) {
	file, err := os.Open(filePath)
	if err != nil {
		slog.Error("file", "error", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		logChan <- scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		slog.Error("scanner", "error", err)
	}

	close(logChan)
}
