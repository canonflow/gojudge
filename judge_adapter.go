package gojudge

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"os/exec"
	"strconv"
	"strings"
	"syscall"
)

type JudgeResult struct {
	Stdout  string `json:"stdout"`
	Stderr  string `json:"stderr"`
	Verdict string `json:"verdict"`
	Runtime string `json:"runtime"`
}

type JudgeAdapter struct {
	Languages []Language
}

func NewJudgeAdapter() *JudgeAdapter {
	return &JudgeAdapter{
		Languages: []Language{},
	}
}

func (judge *JudgeAdapter) RegisterNewLanguage(lang Language) {
	judge.Languages = append(judge.Languages, lang)
}

func (judge *JudgeAdapter) GetLanguages() map[string]Language {
	result := make(map[string]Language)

	for _, lang := range judge.Languages {
		result[lang.GetName()] = lang
	}

	return result
}

func (judge *JudgeAdapter) Compile(context context.Context, lang Language) error {
	compileCommand := lang.GetCompileCommand()

	cmd := exec.CommandContext(context, "bash", "-c", escapeShellArg(compileCommand))

	var stdout, stderr bytes.Buffer

	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	// Check the exit code
	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			if _, ok := exitErr.Sys().(syscall.WaitStatus); ok {
				errStr := stderr.String()
				// Sanitize the stderr
				errStr = lang.GetSanitizeFunction()(errStr)
				return errors.New(errStr)
			}
		}
	}

	return nil
}

func (judge *JudgeAdapter) Judge(context context.Context, lang Language, memoryLimitMB int, timeLimit int) JudgeResult {
	judgeResult := JudgeResult{
		Stdout:  "",
		Stderr:  "",
		Runtime: fmt.Sprintf("%.3f", float32(memoryLimitMB)),
		Verdict: VERDICT_ACCEPTED,
	}

	runCommand := lang.GetRunCommand()
	memoryLimit := memoryLimitMB * 1024
	cmd := exec.CommandContext(
		context,
		"bash",
		"-c",
		escapeShellArg(fmt.Sprintf("ulimit -St %d -Sm %d ; %s", timeLimit, memoryLimit, runCommand)),
	)

	var stdout, stderr bytes.Buffer

	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		judgeResult.Stderr = VERDICT_TIME_LIMIT_EXCEEDED
		judgeResult.Verdict = VERDICT_TIME_LIMIT_EXCEEDED
	}

	stdoutStr := stdout.String()
	stderrStr := stderr.String()

	if strings.Contains(stderrStr, ULIMIT_MEMORY_LIMIT) {
		// Check memory limit
		judgeResult.Verdict = VERDICT_MEMORY_LIMIT_EXCEEDED
		judgeResult.Stderr = VERDICT_MEMORY_LIMIT_EXCEEDED
	} else if strings.Contains(stderrStr, ULIMIT_TIME_LIMIT) {
		// Check time limit
		judgeResult.Verdict = VERDICT_TIME_LIMIT_EXCEEDED
		judgeResult.Stderr = VERDICT_TIME_LIMIT_EXCEEDED
	}

	// Check run time error
	if judgeResult.Verdict == VERDICT_ACCEPTED && !hasRealAtPos1(stderrStr) {
		judgeResult.Verdict = VERDICT_RUNTIME_ERROR
		// Sanitize the stderr
		judgeResult.Stderr = lang.GetSanitizeFunction()(stderrStr)
	}

	// Accepted
	if judgeResult.Verdict == VERDICT_ACCEPTED {
		judgeResult.Stdout = stdoutStr
		judgeResult.Stderr = stderrStr
		runtime, err := parseRuntime(stderr.String())
		if err != nil {
			judgeResult.Runtime = fmt.Sprintf("%.3f", float32(timeLimit))
		} else {
			judgeResult.Runtime = strconv.FormatFloat(runtime, 'f', -1, 64)
		}
	}

	return judgeResult
}

func hasRealAtPos1(s string) bool {
	if len(s) < 5 {
		return false
	}
	// ASCII-safe; if you expect non-ASCII, convert to rune slice.
	return s[1:5] == "real"
}
