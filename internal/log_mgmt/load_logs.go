package log_mgmt

import (
	"os"
)

/*
 * Question: Where is the logs coming from?
 *
 * This function is just reading from the logs/log_file.log
 */
func LoadLogs() ([]byte, error) {
	filePath := "logs/log_file.log"

	logs, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	return logs, nil
}
