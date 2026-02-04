package gojudge

type Language interface {
	GetName() string
	IsCompiled() bool
	GetCompileCommand(basePath string, filename string) (string, error)
	GetJudgeCommand(filename string, outputFilename string, testcase string) string
	GetExtension() string
}

type Judge interface {
	Compile(lang Language, filename string)
	Judge(lang Language, filename string, memoriLimit int, timeLimit int, testcases []string)
	RegisterNewLanguage(lang Language)
	GetLanguages() map[string]Language
}
