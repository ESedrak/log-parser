package log_mgmt

import (
	"bufio"
	"log/slog"
	"os"
)

/*
 *  Function reads log entries one line at a time and sends them through a log channel for concurrent processing
 */
func LoadLogs(filename string, logChan chan<- string) {
	file, err := os.Open(filename)
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
