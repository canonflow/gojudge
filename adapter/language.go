package gojudge

import (
	"errors"
	"strings"
)

type LanguageAdapter struct {
	Name           string
	Compiled       bool
	CompileCommand string
	JudgeCommand   string
	Extension      string

	// METADATA
	CompiledFilename string
	ResultFilename   string
}

func NewLanguageAdapter(name string, isCompiled bool, compileCommand string, judgeCommand string, extension string) *LanguageAdapter {
	return &LanguageAdapter{
		Name:           name,
		Compiled:       isCompiled,
		CompileCommand: compileCommand,
		JudgeCommand:   judgeCommand,
		Extension:      extension,
	}
}

func (l *LanguageAdapter) GetName() string {
	return l.Name
}

func (l *LanguageAdapter) IsCompiled() bool {
	return l.Compiled
}

func (l *LanguageAdapter) GetCompileCommand(basePath string, filename string) (string, error) {
	// clang++ {base_path}/{file} -o {base_path}/{compiled}
	// javac {base_path}/{file} -d {base_path}

	finalCompileCommand := l.CompileCommand

	if !strings.Contains(finalCompileCommand, "{base_path}") {
		return "", errors.New("Invalid Compile Command! Missing template {base_path}")
	}

	finalCompileCommand = strings.ReplaceAll(finalCompileCommand, "{base_path}", basePath)

	if !strings.Contains(finalCompileCommand, "{file}") {
		return "", errors.New("Invalid Compile Command! Missing template {file}")
	}

	finalCompileCommand = strings.ReplaceAll(finalCompileCommand, "{file}", filename)

	if l.CompileCommand != "" {
		if !strings.Contains(finalCompileCommand, "{compiled}") {
			return "", errors.New("Invalid Compile Command! Missing template {compile} for CompiledFilename")
		}
		finalCompileCommand = strings.ReplaceAll(finalCompileCommand, "{compiled}", l.CompiledFilename)
	}

	return finalCompileCommand, nil
}

func (l *LanguageAdapter) GetJudgeCommand(filename string, testcase string) string {
	return l.JudgeCommand
}

func (l *LanguageAdapter) GetExtension() string {
	return l.Extension
}
