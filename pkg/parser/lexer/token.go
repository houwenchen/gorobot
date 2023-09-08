package lexer

type StatementTokens []Token

type Token struct {
	SETTING_HEADER       string
	VARIABLE_HEADER      string
	TESTCASE_HEADER      string
	TASK_HEADER          string
	KEYWORD_HEADER       string
	COMMENT_HEADER       string
	INVALID_HEADER       string
	FATAL_INVALID_HEADER string

	TESTCASE_NAME  string
	KEYWORD_NAME   string
	SUITE_NAME     string
	DOCUMENTATION  string
	SUITE_SETUP    string
	SUITE_TEARDOWN string
	METADATA       string
	TEST_SETUP     string
	TEST_TEARDOWN  string
	TEST_TEMPLATE  string
	TEST_TIMEOUT   string
	TEST_TAGS      string
	DEFAULT_TAGS   string
	KEYWORD_TAGS   string
	LIBRARY        string
	RESOURCE       string
	VARIABLES      string
	SETUP          string
	TEARDOWN       string
	TEMPLATE       string
	TIMEOUT        string
	TAGS           string
	ARGUMENTS      string

	RETURN         string
	RETURN_SETTING string

	WITH_NAME string
	AS        string

	NAME             string
	VARIABLE         string
	ARGUMENT         string
	ASSIGN           string
	KEYWORD          string
	FOR              string
	FOR_SEPARATOR    string
	END              string
	IF               string
	INLINE_IF        string
	ELSE_IF          string
	ELSE             string
	TRY              string
	EXCEPT           string
	FINALLY          string
	WHILE            string
	RETURN_STATEMENT string
	CONTINUE         string
	BREAK            string
	OPTION           string

	SEPARATOR    string
	COMMENT      string
	CONTINUATION string
	CONFIG       string
	EOL          string
	EOS          string

	ERROR string

	FATAL_ERROR string

	NON_DATA_TOKENS []string
	SETTING_TOKENS  []string
	HEADER_TOKENS   []string
	ALLOW_VARIABLES []string
}

func NewToken() *Token {
	t := &Token{
		SETTING_HEADER:       "SETTING HEADER",
		VARIABLE_HEADER:      "VARIABLE HEADER",
		TESTCASE_HEADER:      "TESTCASE HEADER",
		TASK_HEADER:          "TASK HEADER",
		KEYWORD_HEADER:       "KEYWORD HEADER",
		COMMENT_HEADER:       "COMMENT HEADER",
		INVALID_HEADER:       "INVALID HEADER",
		FATAL_INVALID_HEADER: "FATAL INVALID HEADER",

		TESTCASE_NAME:  "TESTCASE NAME",
		KEYWORD_NAME:   "KEYWORD NAME",
		SUITE_NAME:     "SUITE NAME",
		DOCUMENTATION:  "DOCUMENTATION",
		SUITE_SETUP:    "SUITE SETUP",
		SUITE_TEARDOWN: "SUITE TEARDOWN",
		METADATA:       "METADATA",
		TEST_SETUP:     "TEST SETUP",
		TEST_TEARDOWN:  "TEST TEARDOWN",
		TEST_TEMPLATE:  "TEST TEMPLATE",
		TEST_TIMEOUT:   "TEST TIMEOUT",
		TEST_TAGS:      "TEST TAGS",
		DEFAULT_TAGS:   "DEFAULT TAGS",
		KEYWORD_TAGS:   "KEYWORD TAGS",
		LIBRARY:        "LIBRARY",
		RESOURCE:       "RESOURCE",
		VARIABLES:      "VARIABLES",
		SETUP:          "SETUP",
		TEARDOWN:       "TEARDOWN",
		TEMPLATE:       "TEMPLATE",
		TIMEOUT:        "TIMEOUT",
		TAGS:           "TAGS",
		ARGUMENTS:      "ARGUMENTS",

		RETURN:         "RETURN",
		RETURN_SETTING: "RETURN",

		WITH_NAME: "AS",
		AS:        "AS",

		NAME:             "NAME",
		VARIABLE:         "VARIABLE",
		ARGUMENT:         "ARGUMENT",
		ASSIGN:           "ASSIGN",
		KEYWORD:          "KEYWORD",
		FOR:              "FOR",
		FOR_SEPARATOR:    "FOR SEPARATOR",
		END:              "END",
		IF:               "IF",
		INLINE_IF:        "INLINE IF",
		ELSE_IF:          "ELSE IF",
		ELSE:             "ELSE",
		TRY:              "TRY",
		EXCEPT:           "EXCEPT",
		FINALLY:          "FINALLY",
		WHILE:            "WHILE",
		RETURN_STATEMENT: "RETURN STATEMENT",
		CONTINUE:         "CONTINUE",
		BREAK:            "BREAK",
		OPTION:           "OPTION",

		SEPARATOR:    "SEPARATOR",
		COMMENT:      "COMMENT",
		CONTINUATION: "CONTINUATION",
		CONFIG:       "CONFIG",
		EOL:          "EOL",
		EOS:          "EOS",

		ERROR: "ERROR",

		FATAL_ERROR: "FATAL ERROR",

		NON_DATA_TOKENS: []string{
			"SEPARATOR",
			"COMMENT",
			"CONTINUATION",
			"EOL",
			"EOS",
		},
		SETTING_TOKENS: []string{
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
		},
		HEADER_TOKENS: []string{
			"SETTING HEADER",
			"VARIABLE HEADER",
			"TESTCASE HEADER",
			"TASK HEADER",
			"KEYWORD HEADER",
			"COMMENT HEADER",
			"INVALID HEADER",
		},
		ALLOW_VARIABLES: []string{
			"NAME",
			"ARGUMENT",
			"TESTCASE NAME",
			"KEYWORD NAME",
		},
	}

	return t
}
