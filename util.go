package main

import (
	"fmt"
	"math"
	"net"
	"strings"
)

// formatBytes returns the humanized bandwidth based on `bytes` and `seconds`.
func formatBytes(bytes int64, seconds float64) string {
	raw := float64(bytes*8) / seconds
	if raw <= 10 {
		return fmt.Sprintf("%.2f b/s", raw)
	}

	exp := math.Floor(math.Log(raw) / math.Log(unitBase))
	suffix := suffixes[int(exp)]
	bandwidth := raw / math.Pow(unitBase, exp)
	return fmt.Sprintf("%.2f %s/s", bandwidth, suffix)
}

// formatIP extracts an IP address out of `addr`.
func formatIP(addr net.Addr) string {
	chunks := strings.Split(addr.String(), ":")
	ip := strings.Join(chunks[0:len(chunks)-1], ":")
	return strings.Trim(ip, "[]")
}
