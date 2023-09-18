package running

// import (
// 	"context"
// 	"time"

// 	"github.com/google/cel-go/parser"
// )

// // 整个工具的对外接口
// // 将各个小接口集成进来
// type Runner interface {
// 	parser.Parser

// 	Run() error
// }

// type runner struct {
// }

// func (r *runner) Run() error {
// 	return nil
// }

// func (r *runner) setup() (context.Context, context.CancelFunc) {
// 	ctx := context.TODO()
// 	timeout := 100 * time.Minute

// 	return context.WithTimeout(ctx, timeout)
// }
