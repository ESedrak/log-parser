package ip_mgmt

/*
 * Returns the number of unique IP addresses stored in the IPCounter instance
 * Potentially do not need - but its a nice touch
 */
func (i *IPCounter) UniqueIPs() int {
	return len(i.IPCounts)
}
