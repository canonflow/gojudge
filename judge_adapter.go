package gojudge

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"os/exec"
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

func (judge *JudgeAdapter) Compile(context context.Context, lang Language, filename string) error {
	compileCommand := lang.GetCompileCommand()

	cmd := exec.CommandContext(context, compileCommand)

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

func (judge *JudgeAdapter) Judge(lang Language, testcase string, memoriLimit int, timeLimit int) JudgeResult {
	judgeResult := JudgeResult{
		Stdout:  "",
		Stderr:  "",
		Runtime: fmt.Sprintf("%.3f", float32(memoriLimit)),
		Verdict: VERDICT_ACCEPTED,
	}

	return judgeResult
}
