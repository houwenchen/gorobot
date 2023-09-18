package parsing

// import (
// 	"os"
// 	"path"
// )

// // 解析 .gorobot 文件
// type Parser interface {
// 	BuildTestSuite()
// }

// type CustomParser struct {
// }

// type parser struct {
// 	caseFilePath string
// 	caseName     string
// 	caseFileExt  string
// 	caseInfo     []byte
// }

// func newParser(path string) *parser {
// 	p := &parser{}
// 	p.caseFilePath = path
// 	err := p.handleTestFile()
// 	if err != nil {
// 		panic(err)
// 	}

// 	return p
// }

// // 获取 case 文件的名字以及文件扩展名
// func (p *parser) handleTestFile() error {
// 	p.caseName = path.Base(p.caseFilePath)
// 	p.caseFileExt = path.Ext(p.caseName)
// 	caseInfo, err := os.ReadFile(p.caseFilePath)
// 	if err != nil {
// 		return err
// 	}

// 	p.caseInfo = caseInfo
// 	return nil
// }
