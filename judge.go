package gojudge

var (
	// ===== VERDICT DATA
	VERDICT_ACCEPTED              = "accepted"
	VERDICT_WRONG_ANSWER          = "wrong answer"
	VERDICT_COMPILE_ERROR         = "compile error"
	VERDICT_RUNTIME_ERROR         = "runtime answer"
	VERDICT_TIME_LIMIT_EXCEEDED   = "time limit exceeded"
	VERDICT_MEMORY_LIMIT_EXCEEDED = "memory limit answer"

	// ===== ULIMIT STRING
	ULIMIT_TIME_LIMIT   = "CPU time limit exceeded"
	ULIMIT_MEMORY_LIMIT = "Memory size limit exceeded"

	// COMPILE AND JUDGE META
	BASE_COMPILE_FILENAME = "compiled_program"
	BASE_OUTPUT_FILENAME  = "result.out"
)
