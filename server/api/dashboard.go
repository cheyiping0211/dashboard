package v1

import (
    "github.com/gin-gonic/gin"
    "ping/util"
    "ping/model"
    "ping/service"
    "os/exec"
)

/*
* pwd     获取当前路径
* module  获取当前模块
*/
func GetModules (c *gin.Context) {
    module:= c.Query("module") 
    err, data := service.GetModules(module)
	if err != nil {
        util.FailWithMessage(err.Error(), c)
	} else {
		util.OkWithData(data, c)
	}
}

func UpdateModules (c *gin.Context) {
    data := model.Chart{}
    if err := c.ShouldBindJSON(&data); err != nil {
        util.FailWithMessage(err.Error(), c)
        return
     }

	if err := service.UpdateChart(data); err != nil {
		util.FailWithMessage("更新失败", c)
	} else {
		util.OkWithMessage("更新成功", c)
	}
}

func BuildModules (c *gin.Context){
    cmd := exec.Command("cmd", "-c", "ping 127.0.0.8")
    _, err := cmd.Output()
    if err != nil {
        util.FailWithMessage("导出脚本编译错误", c)
        panic(err.Error())
    }   
}

