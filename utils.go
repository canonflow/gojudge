package gojudge

import (
	"fmt"
	"strconv"
	"strings"
)

func escapeShellArg(arg string) string {
	escaped := strings.ReplaceAll(arg, "'", `'\'`)
	return "'" + escaped + "'"
}

func EscapeShellArg(arg string) string {
	return escapeShellArg(arg)
	// escaped := strings.ReplaceAll(arg, "'", `'\'`)
	// return "'" + escaped + "'"
}

func parseRuntime(stream string) (float64, error) {
	// Start from "real"
	idx := strings.Index(stream, "real")
	if idx == -1 {
		return 0, fmt.Errorf(`keyword "real" not found`)
	}
	str := stream[idx:]

	// Replace comma with dot
	str = strings.ReplaceAll(str, ",", ".")

	// Find positions of m and s in the trimmed string
	im := strings.Index(str, "m")
	is := strings.Index(str, "s")

	if im == -1 || is == -1 {
		return 0, fmt.Errorf(`runtime format invalid (missing "m" or "s")`)
	}

	// Extract minute string: starts after "real "
	minStr := strings.TrimSpace(str[5:im])
	// Extract seconds
	secStr := strings.TrimSpace(str[im+1 : is])

	// Convert to numeric
	m, err := strconv.Atoi(minStr)
	if err != nil {
		return 0, fmt.Errorf("invalid minutes: %w", err)
	}

	// seconds can be float
	s, err := strconv.ParseFloat(secStr, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid seconds: %w", err)
	}

	// Total seconds
	total := float64(m)*60 + s

	return total, nil
}

// func filterCompilerError(errorMessage string) string {
// 	patterns := []*regexp.Regexp{
// 		// Windows paths: C:\folder\file.ext
// 		regexp.MustCompile(`[A-Za-z]:\\(?:[^\\/:*?"<>|\r\n]+\\)*[^\\/:*?"<>|\r\n]*`),

// 		// Unix paths: /folder/sub/file.ext
// 		regexp.MustCompile(`(?:/[^/\s]+)+`),

// 		// Any line containing a path
// 		regexp.MustCompile(`(?m)^.*[\\/].*$`),
// 	}

// 	filtered := errorMessage

// 	// Apply patterns
// 	for _, re := range patterns {
// 		filtered = re.ReplaceAllString(filtered, "")
// 	}

// 	// Remove empty lines
// 	emptyLine := regexp.MustCompile(`(?m)^\s*[\r\n]`)
// 	filtered = emptyLine.ReplaceAllString(filtered, "")

// 	return strings.TrimSpace(filtered)
// }
