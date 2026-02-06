package gojudge

import "context"

type Language interface {
	GetName() string
	IsCompiled() bool
	GetCompileCommand() string
	GetRunCommand() string
	GetExtension() string
	GetSanitizeFunction() func(err string) string
}

type Judge interface {
	Compile(context context.Context, lang Language, filename string) error
	Judge(context context.Context, lang Language, memoriLimit int, timeLimit int) JudgeResult
	RegisterNewLanguage(lang Language)
	GetLanguages() map[string]Language
}
