package routerusers

import (
	"written/code/golang/gin/018UseHTTPMethod/api/users"

	"github.com/gin-gonic/gin"
)

func Users(router *gin.Engine) {
	// 获取所有用户信息
	router.GET("/users", apiusers.GetUserAll())
	// 获取一个用户信息
	router.GET("/users/:id", apiusers.GetOne())
	// 添加一个用户
	router.POST("/users", apiusers.AddUser())
	// 修改一个用户完整信息
	router.PUT("/users/:id", apiusers.ChangeUser())
	// 修改一个用户部分信息
	router.PATCH("/users/:id", apiusers.ModifyUser())
	// 删除一个用户
	router.DELETE("/users/:id", apiusers.DeleteUser())
	// 获取用户资源的元数据
	router.HEAD("/users", apiusers.HeadUser())
	// 获取信息，关于用户资源的哪些属性是客户端可以改变的。
	router.OPTIONS("/users", apiusers.OptionsUser())
}
