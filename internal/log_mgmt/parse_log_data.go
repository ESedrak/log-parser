package log_mgmt

import (
	"errors"
	"regexp"
)

/*
 * Question: Will the logs always be in the same format?
 *
 * This function assumes logs coming through will in the same logsFormat
 * Takes raw log data as input, applies a regular expression pattern to parse each log entry based on a given format template, and returns a slice containing the parsed log entries.
 *
 * Regex format will be (?P<client_ip>.*) - - \[(?P<time_stamp>.*)\] "(?P<http_method>.*) (?P<request_path>.*) (?P<_>.*)" (?P<response_code>.*) (?P<_>.*) "-" "(?P<user_agent>.*)"
 */
func ParseLogData(logsData []byte) ([][]string, error) {
	logsFormat := `$client_ip - - \[$time_stamp\] "$http_method $request_path $_" $response_code $_ "-" "$user_agent"`
	regexFormat := regexp.MustCompile(`\$([\w_]*)`).ReplaceAllString(logsFormat, `(?P<$1>.*)`)

	re := regexp.MustCompile(regexFormat)

	matches := re.FindAllStringSubmatch(string(logsData), -1)

	if len(matches) == 0 {
		return nil, errors.New("no log data matches")
	}

	// exclude index 0 of each array - as this is returning the full log
	result := make([][]string, 0, len(matches))
	for _, match := range matches {
		result = append(result, match[1:])
	}

	return result, nil
}
