package main

import (
	"github.com/yaoapp/yao/cmd"
	"github.com/yaoapp/yao/utils"

	_ "github.com/yaoapp/gou/encoding"
	_ "github.com/yaoapp/yao/aigc"
	_ "github.com/yaoapp/yao/crypto"
	_ "github.com/yaoapp/yao/helper"
	_ "github.com/yaoapp/yao/myapi"
	_ "github.com/yaoapp/yao/openai"
	_ "github.com/yaoapp/yao/wework"
)

// 主程序
func main() {
	utils.Init()
	cmd.Execute()
}
