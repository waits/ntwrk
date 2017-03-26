package main

import (
	"fmt"
	"math"
	"net"
	"strings"
	"time"
)

const unitBase = 1000

var suffixes = [...]string{"b", "kb", "Mb", "Gb", "Tb", "Pb", "Eb"}

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

// formatProgress constructs a progress bar string for `n / d`.
func formatProgress(n, d time.Duration) string {
	ratio := float64(n) / float64(d)
	bar := strings.Repeat("=", int(ratio*20+0.5))
	return fmt.Sprintf("[%-20s] %3.0f%%", bar, ratio*100)
}
