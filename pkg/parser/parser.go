package parser

import "path"

// 解析 .gorobot 文件
type Parser interface {
	BuildTestSuite()
}

type parser struct {
	caseFilePath string
	caseName     string
	caseFileExt  string
}

// 获取 case 文件的名字以及文件扩展名
func (p *parser) handleTestFile() {
	p.caseName = path.Base(p.caseFilePath)
	p.caseFileExt = path.Ext(p.caseName)
}
