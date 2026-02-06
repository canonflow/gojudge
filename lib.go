package gojudge

import (
	"regexp"
	"strings"
)

func filterCompilerError(errorMessage string) string {
	patterns := []*regexp.Regexp{
		// Windows paths
		regexp.MustCompile(`[A-Za-z]:\\(?:[^\\/:*?"<>|\r\n]+\\)*[^\\/:*?"<>|\r\n]*`),
		// Unix paths
		regexp.MustCompile(`(?:/[^/\s]+)+`),
		// Lines containing file paths
		regexp.MustCompile(`^.*[\\/].*$`),
	}

	filtered := errorMessage
	for _, p := range patterns {
		filtered = p.ReplaceAllString(filtered, "")
	}

	// Remove empty lines
	empty := regexp.MustCompile(`(?m)^\s*[\r\n]`)
	filtered = empty.ReplaceAllString(filtered, "")

	return strings.TrimSpace(filtered)
}

func before(s, sep string) string {
	idx := strings.Index(s, sep)
	if idx == -1 {
		return s
	}
	return s[:idx]
}

func CPPBaseSanitizer(stderr string) string {
	if stderr == "" {
		return stderr
	}

	err := before(stderr, "real")

	tempErr := before(err, "(core dumped)")

	if tempErr != "" {
		filtered := filterCompilerError(tempErr)
		if filtered != "" {
			return filtered
		}

		return tempErr
	}

	filtered := filterCompilerError(err)

	if filtered != "" {
		return filtered
	}

	return err
}

func JavaBaseSanitizer(stderr string) string {
	if stderr == "" {
		return stderr
	}

	err := before(stderr, "real")

	tempErr := before(err, "(core dumped)")

	if tempErr != "" {
		filtered := filterCompilerError(tempErr)
		if filtered != "" {
			return filtered
		}

		return tempErr
	}

	filtered := filterCompilerError(err)

	if filtered != "" {
		return filtered
	}

	return err
}

func PascalBaseSanitizer(stderr string) string {
	if stderr == "" {
		return stderr
	}

	re := regexp.MustCompile(`(?m)\w+\.pas\(\d+,\d+\) (?:Error|Fatal|Note): (.+)$`)
	matches := re.FindAllStringSubmatch(stderr, -1)

	if matches == nil {
		return "Run Time Error"
	}

	var out []string
	for _, m := range matches {
		out = append(out, m[1])
	}

	return strings.Join(out, "\n")
}
