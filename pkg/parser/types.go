package parser

import "time"

const (
	// 文件扩展名
	Extention string = ".gorobot"
)

type Settings struct {
	Document string
	Timeout  time.Duration
}

type Packages struct {
	Library []string
}

type Variables struct {
	Variable map[string]interface{}
}

type TestCases struct {
}

type TestCase struct {
}

type Keywords struct {
}

type Keyword struct {
}
