package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type HelloReq struct {
	g.Meta `path:"/hello" tags:"Hello" method:"get" summary:"Greetings" description:"Hello, I'm 65050437 Noppakorn" consumes:"application/json" produces:"text/html"`
}
type HelloRes struct {
	g.Meta `mime:"text/html" example:"string"`
}
