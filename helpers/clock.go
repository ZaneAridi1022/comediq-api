package helpers

import (
	"fmt"
	"strings"
	"time"
)

func ClockAdd(clockStr string, minutes int) (string, error) {
	// Normalize input string (e.g. "PM" -> "pm")
	clockStr = strings.ToLower(strings.TrimSpace(clockStr))

	// Parse using Go's time layout (must include Jan 2, 2006 trick)
	parsed, err := time.Parse("3:04pm", clockStr)
	if err != nil {
		return "", fmt.Errorf("invalid time format: %w", err)
	}

	// Add minutes
	parsed = parsed.Add(time.Duration(minutes) * time.Minute)

	// Format back to 12-hour clock with am/pm
	return parsed.Format("3:04pm"), nil
}
