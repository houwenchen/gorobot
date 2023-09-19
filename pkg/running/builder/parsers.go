package builder

import (
	"fmt"
	"houwenchen/gorobot/pkg/parsing/model"
	"houwenchen/gorobot/pkg/running"
	"reflect"
)

type Parser interface {
	Name() string
	ParseSuiteFile(Path, TestDefaults) running.TestSuite
	ParseInitFile(Path, TestDefaults) running.TestSuite
	ParseResourceFile(Path) running.ResourceFile
}

type RobotParser interface {
	Parser

	ParseModel(model.File, TestDefaults) running.TestSuite
	ParseResourceModel(model.File) running.ResourceFile
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
	extensions []interface{}
}

type restParser struct {
}

type jsonParser struct {
}

type customParser struct {
}

type Path struct {
}

func NewParser() *parser {
	return &parser{}
}

func (p *parser) Name() string {
	rType := reflect.TypeOf(p)
	return rType.String()
}

func (p *parser) ParseSuiteFile(path Path, testDefaults TestDefaults) running.TestSuite {
	panic(fmt.Sprintf("%s does not support parsing suite files.", p.Name()))
}

func (p *parser) ParseInitFile(path Path, testDefaults TestDefaults) running.TestSuite {
	panic(fmt.Sprintf("%s does not support parsing suite files.", p.Name()))
}

func (p *parser) ParseResourceFile(path Path) running.ResourceFile {
	panic(fmt.Sprintf("%s does not support parsing suite files.", p.Name()))
}

func NewRobotParser() *robotParser {
	return &robotParser{}
}

func (rp *robotParser) Name() string {
	rType := reflect.TypeOf(rp)
	return rType.String()
}

func NewRestParser() *RestParser {
	return nil
}

func NewJsonParser() *JsonParser {
	return nil
}
