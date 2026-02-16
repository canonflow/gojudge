package test

import (
	"context"
	"fmt"
	"testing"

	"github.com/canonflow/gojudge"
	"github.com/canonflow/gojudge/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

/*
=======================================================
Integration tests for the Base Sanitizer Function
=======================================================
*/
// func TestCPPBaseSanitizer(t *testing.T)
// func TestJavaBaseSanitizer(t *testing.T)
// func TestPascalBaseSanitizer(t *testing.T)

/*
=====================================================================
Integration tests for the Register New Language & GetLanguages
=====================================================================
*/
// func TestRegisterNewLanguage(t *testing.T)

func TestEscapeShellArgFunction(t *testing.T) {
	timeLimit := 1
	memoryLimit := 64 * 1024
	command := "clang++ program/index.cpp -o program/compiled_cpp"
	fullCommand := fmt.Sprintf("ulimit -St %d -Sm %d ; %s", timeLimit, memoryLimit, command)

	finalCommand := gojudge.EscapeShellArg(fullCommand)

	assert.Equal(t, "'"+fullCommand+"'", finalCommand)
}

/*
=======================================================
Integration tests for the Compile Function
=======================================================
*/
func TestCompileSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLang := mocks.NewMockLanguage(ctrl)
	mockLang.EXPECT().GetCompileCommand().Return(":")
	mockLang.EXPECT().GetSanitizeFunction().Return(gojudge.CPPBaseSanitizer).AnyTimes()

	judge := gojudge.NewJudgeAdapter()
	err := judge.Compile(context.Background(), mockLang)

	assert.NoError(t, err)
}

func TestCompileFailed(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLang := mocks.NewMockLanguage(ctrl)
	mockLang.EXPECT().GetCompileCommand().Return("false")
	mockLang.EXPECT().GetSanitizeFunction().Return(gojudge.CPPBaseSanitizer)

	judge := gojudge.NewJudgeAdapter()
	err := judge.Compile(context.Background(), mockLang)

	assert.Error(t, err)
}

/*
=======================================================
Integration tests for the Judge Function
=======================================================
*/

// func TestRuntimeError(t *testing.T)
// func TestMemoryLimitError(t *testing.T)
// func TestTimeLimitError(t *testing.T)
// func TestAccepted(t *testing.T)
