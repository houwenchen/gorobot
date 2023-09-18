package builders

import "houwenchen/gorobot/pkg/running"

type Parser interface {
	ParseSuiteFile(Path, TestDefaults) running.TestSuite
	ParseInitFile(Path, TestDefaults) running.TestSuite
	ParseResourceFile(Path) running.ResourceFile
}

type RobotParser interface {
	Parser
}

type RestParser interface {
	RobotParser
}

type JsonParser interface {
	Parser
}

type CustomParser interface {
	Parser
}

type parser struct {
}

type robotParser struct {
}

type restParser struct {
}

type jsonParser struct {
}

type customParser struct {
}

type Path struct {
}

func NewRobotParser() *RobotParser {
	return &RobotParser{}
}

func NewRestParser() *RestParser {
	return &RestParser{}
}

func NewJsonParser() *JsonParser {
	return &JsonParser{}
}
