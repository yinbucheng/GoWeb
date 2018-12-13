package main

import (
	"bucheng.first/utils"
	"fmt"
)

func main() {
	//PostJsonTest()
	heads := make(map[string]string)
	heads["nice"] = "haha"
	heads["good"] = "byte"
	result := utils.PosJsonWithHead("http://127.0.0.1/user/testList", heads, "{\"name\":\"yinchong\",\"age\":20}")
	fmt.Println(result)
}

func PostJsonTest() {
	result := utils.PostJson("http://127.0.0.1/user/testList", "{\"name\":\"yinchong\",\"age\":20}")
	fmt.Println(result)
}
