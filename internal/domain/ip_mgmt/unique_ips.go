package ip_mgmt

import "errors"

func UniqueIPs(ipCounts map[string]int) (int, error) {
	if len(ipCounts) == 0 {
		return 0, errors.New("no IPs found")
	}
	return len(ipCounts), nil
}
