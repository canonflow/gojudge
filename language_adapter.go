package gojudge

type LanguageAdapter struct {
	Name           string
	Compiled       bool
	CompileCommand string
	RunCommand     string
	Extension      string
	SanitizeStderr func(err string) string
}

func NewLanguageAdapter(name string, isCompiled bool, compileCommand string, runCommand string, extension string, sanitizeFunc func(err string) string) *LanguageAdapter {
	return &LanguageAdapter{
		Name:           name,
		Compiled:       isCompiled,
		CompileCommand: compileCommand,
		RunCommand:     runCommand,
		Extension:      extension,
		SanitizeStderr: sanitizeFunc,
	}
}

func (l *LanguageAdapter) GetName() string {
	return l.Name
}

func (l *LanguageAdapter) IsCompiled() bool {
	return l.Compiled
}

func (l *LanguageAdapter) GetCompileCommand() string {
	// clang++ {base_path}/{file} -o {base_path}/{compiled}
	// javac {base_path}/{file} -d {base_path}

	// finalCompileCommand := l.CompileCommand

	// if !strings.Contains(finalCompileCommand, "{base_path}") {
	// 	return "", errors.New("Invalid Compile Command! Missing template {base_path}")
	// }

	// finalCompileCommand = strings.ReplaceAll(finalCompileCommand, "{base_path}", basePath)

	// if !strings.Contains(finalCompileCommand, "{file}") {
	// 	return "", errors.New("Invalid Compile Command! Missing template {file}")
	// }

	// finalCompileCommand = strings.ReplaceAll(finalCompileCommand, "{file}", filename)

	// if l.CompileCommand != "" {
	// 	if !strings.Contains(finalCompileCommand, "{compiled}") {
	// 		return "", errors.New("Invalid Compile Command! Missing template {compile} for CompiledFilename")
	// 	}
	// 	finalCompileCommand = strings.ReplaceAll(finalCompileCommand, "{compiled}", l.CompiledFilename)
	// }

	// return finalCompileCommand, nil
	return l.CompileCommand
}

func (l *LanguageAdapter) GetRunCommand() string {
	return l.RunCommand
}

func (l *LanguageAdapter) GetExtension() string {
	return l.Extension
}

func (l *LanguageAdapter) GetSanitizeFunction() func(err string) string {
	return l.SanitizeStderr
}
