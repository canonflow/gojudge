package gojudge

import "github.com/canonflow/gojudge"

type JudgeAdapter struct {
	Languages []gojudge.Language
}

func NewJudgeAdapter() *JudgeAdapter {
	return &JudgeAdapter{
		Languages: getSupportedLanguages(),
	}
}

func getSupportedLanguages() []gojudge.Language {
	return []gojudge.Language{
		LanguageAdapter{
			Name:           "cpp",
			Compiled:       true,
			CompileCommand: "clang++ ",
			JudgeCommand:   "",
			Extension:      ".cpp",
		},
	}
}
