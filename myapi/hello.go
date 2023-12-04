package myapi

import (
	gouProcess "github.com/yaoapp/gou/process"
)

func init() {
	gouProcess.Register("yao.myapi.hello", processHello) // 注册处理器
}

// processHello 处理逻辑
func processHello(process *gouProcess.Process) interface{} {
	args := process.Args
	return map[string]interface{}{"args": args, "msg": "ok"}
}
