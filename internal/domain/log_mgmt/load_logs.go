package log_mgmt

import (
	"bufio"
	"os"
)

/*
 *  Function reads logs one at a time and sends them through a log channel for concurrent processing
 */
func LoadLogs(filePath string, logChan chan<- string, errChan chan<- error) {
	defer close(logChan)
	defer close(errChan)

	file, err := os.Open(filePath)
	if err != nil {
		errChan <- err
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		logChan <- scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		errChan <- err
		return
	}
}
