package parsing

const (
	// 文件扩展名
	Extention string = ".gorobot"
)

type TokenType string

const (
	SETTING_HEADER       TokenType = "SETTING HEADER"
	VARIABLE_HEADER      TokenType = "VARIABLE HEADER"
	TESTCASE_HEADER      TokenType = "TESTCASE HEADER"
	TASK_HEADER          TokenType = "TASK HEADER"
	KEYWORD_HEADER       TokenType = "KEYWORD HEADER"
	COMMENT_HEADER       TokenType = "COMMENT HEADER"
	INVALID_HEADER       TokenType = "INVALID HEADER"
	FATAL_INVALID_HEADER TokenType = "FATAL INVALID HEADER"

	TESTCASE_NAME  TokenType = "TESTCASE NAME"
	KEYWORD_NAME   TokenType = "KEYWORD NAME"
	SUITE_NAME     TokenType = "SUITE NAME"
	DOCUMENTATION  TokenType = "DOCUMENTATION"
	SUITE_SETUP    TokenType = "SUITE SETUP"
	SUITE_TEARDOWN TokenType = "SUITE TEARDOWN"
	METADATA       TokenType = "METADATA"
	TEST_SETUP     TokenType = "TEST SETUP"
	TEST_TEARDOWN  TokenType = "TEST TEARDOWN"
	TEST_TEMPLATE  TokenType = "TEST TEMPLATE"
	TEST_TIMEOUT   TokenType = "TEST TIMEOUT"
	TEST_TAGS      TokenType = "TEST TAGS"
	DEFAULT_TAGS   TokenType = "DEFAULT TAGS"
	KEYWORD_TAGS   TokenType = "KEYWORD TAGS"
	LIBRARY        TokenType = "LIBRARY"
	RESOURCE       TokenType = "RESOURCE"
	VARIABLES      TokenType = "VARIABLES"
	SETUP          TokenType = "SETUP"
	TEARDOWN       TokenType = "TEARDOWN"
	TEMPLATE       TokenType = "TEMPLATE"
	TIMEOUT        TokenType = "TIMEOUT"
	TAGS           TokenType = "TAGS"
	ARGUMENTS      TokenType = "ARGUMENTS"

	RETURN         TokenType = "RETURN"
	RETURN_SETTING TokenType = "RETURN"

	WITH_NAME TokenType = "AS"
	AS        TokenType = "AS"

	NAME             TokenType = "NAME"
	VARIABLE         TokenType = "VARIABLE"
	ARGUMENT         TokenType = "ARGUMENT"
	ASSIGN           TokenType = "ASSIGN"
	KEYWORD          TokenType = "KEYWORD"
	FOR              TokenType = "FOR"
	FOR_SEPARATOR    TokenType = "FOR SEPARATOR"
	END              TokenType = "END"
	IF               TokenType = "IF"
	INLINE_IF        TokenType = "INLINE IF"
	ELSE_IF          TokenType = "ELSE IF"
	ELSE             TokenType = "ELSE"
	TRY              TokenType = "TRY"
	EXCEPT           TokenType = "EXCEPT"
	FINALLY          TokenType = "FINALLY"
	WHILE            TokenType = "WHILE"
	RETURN_STATEMENT TokenType = "RETURN STATEMENT"
	CONTINUE         TokenType = "CONTINUE"
	BREAK            TokenType = "BREAK"
	OPTION           TokenType = "OPTION"

	SEPARATOR    TokenType = "SEPARATOR"
	COMMENT      TokenType = "COMMENT"
	CONTINUATION TokenType = "CONTINUATION"
	CONFIG       TokenType = "CONFIG"
	EOL          TokenType = "EOL"
	EOS          TokenType = "EOS"

	ERROR TokenType = "ERROR"

	FATAL_ERROR TokenType = "FATAL ERROR"
)

var (
	NON_DATA_TOKENS = []string{
		"SEPARATOR",
		"COMMENT",
		"CONTINUATION",
		"EOL",
		"EOS",
	}

	SETTING_TOKENS = []string{
		"DOCUMENTATION",
		"SUITE NAME",
		"SUITE SETUP",
		"SUITE TEARDOWN",
		"METADATA",
		"TEST SETUP",
		"TEST TEARDOWN",
		"TEST TEMPLATE",
		"TEST TIMEOUT",
		"TEST TAGS",
		"DEFAULT TAGS",
		"KEYWORD TAGS",
		"LIBRARY",
		"RESOURCE",
		"VARIABLES",
		"SETUP",
		"TEARDOWN",
		"TEMPLATE",
		"TIMEOUT",
		"TAGS",
		"ARGUMENTS",
		"RETURN",
	}

	HEADER_TOKENS = []string{
		"SETTING HEADER",
		"VARIABLE HEADER",
		"TESTCASE HEADER",
		"TASK HEADER",
		"KEYWORD HEADER",
		"COMMENT HEADER",
		"INVALID HEADER",
	}
	ALLOW_VARIABLES = []string{
		"NAME",
		"ARGUMENT",
		"TESTCASE NAME",
		"KEYWORD NAME",
	}
)
