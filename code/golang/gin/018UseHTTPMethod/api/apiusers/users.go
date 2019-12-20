package apiusers

import (
	"errors"
	"net/http"
	"strconv"
	"written/code/golang/gin/018UseHTTPMethod/model/modelusers"

	"github.com/gin-gonic/gin"
)

// users存储用户列表
var users = make(modelusers.Users)

// GetUserAll获取所有用户API接口处理函数
func GetUserAll() gin.HandlerFunc {
	return func(context *gin.Context) {
		returnSuccessfully(context, http.StatusOK, gin.H{
			"describe": "User List data",
			"users":    users.GetAll(),
		})
	}
}

// GetOne获取单个用户API接口处理函数
func GetOne() gin.HandlerFunc {
	return func(context *gin.Context) {
		id := context.Param("id")
		idInt, err := strconv.Atoi(id)
		if err != nil {
			returnError(context, http.StatusBadRequest, err)
			return
		}

		// 获取一个用户
		user, err := users.GetOne(idInt)
		if err != nil {
			returnError(context, http.StatusNotFound, err)
			return
		}
		user, ok := user.(*modelusers.User)
		if !ok {
			returnError(context, http.StatusInternalServerError, err)
		}

		returnSuccessfully(context, http.StatusOK, gin.H{
			"describe": "User detail data",
			"user":     user,
		})
	}
}

// AddUser添加用户API接口处理函数
func AddUser() gin.HandlerFunc {
	return func(context *gin.Context) {
		name := context.PostForm("name")
		phone := context.PostForm("phone")
		email := context.PostForm("email")
		sex := context.PostForm("sex")
		address := context.PostForm("address")
		age := context.PostForm("age")
		ageInt, err := strconv.Atoi(age)
		if err != nil {
			returnError(context, http.StatusBadRequest, err)
			return
		}

		// 添加用户
		user, err := users.Add(&modelusers.User{Name: name, Phone: phone, Email: email, Sex: sex, Address: address, Age: ageInt})
		if err != nil {
			returnError(context, http.StatusInternalServerError, err)
			return
		}
		user, ok := user.(*modelusers.User)
		if !ok {
			returnError(context, http.StatusInternalServerError, errors.New("add user failed"))
			return
		}

		// 返回添加后的响应数据
		returnSuccessfully(context, http.StatusCreated, gin.H{
			"describe": "User added data",
			"user":     user,
		})
	}
}

// ChangeUser修改用户API接口处理函数
func ChangeUser() gin.HandlerFunc {
	return func(context *gin.Context) {
		id := context.Param("id")
		name := context.PostForm("name")
		phone := context.PostForm("phone")
		email := context.PostForm("email")
		sex := context.PostForm("sex")
		address := context.PostForm("address")
		age := context.PostForm("age")

		// 转换ID为整型
		idInt, err := strconv.Atoi(id)
		if err != nil {
			returnError(context, http.StatusBadRequest, err)
			return
		}

		// 转换Age为整型
		ageInt, err := strconv.Atoi(age)
		if err != nil {
			returnError(context, http.StatusBadRequest, err)
			return
		}

		// 修改用户
		user, err := users.Change(idInt, &modelusers.User{Name: name, Phone: phone, Email: email, Sex: sex, Address: address, Age: ageInt})
		if err != nil {
			returnError(context, http.StatusInternalServerError, err)
			return
		}
		user, ok := user.(*modelusers.User)
		if !ok {
			returnError(context, http.StatusInternalServerError, errors.New("change user failed"))
			return
		}

		// 返回添加后的响应数据
		returnSuccessfully(context, http.StatusCreated, gin.H{
			"describe": "User changed data",
			"user":     user,
		})
	}
}

// ModifyUser修改用户API接口处理函数
func ModifyUser() gin.HandlerFunc {
	return func(context *gin.Context) {
		id := context.Param("id")
		address := context.PostForm("address")
		age := context.PostForm("age")

		// 转换ID为整型
		idInt, err := strconv.Atoi(id)
		if err != nil {
			returnError(context, http.StatusBadRequest, err)
			return
		}

		// 转换Age为整型
		ageInt, err := strconv.Atoi(age)
		if err != nil {
			returnError(context, http.StatusBadRequest, err)
			return
		}

		// 修改用户
		user, err := users.Modify(idInt, &modelusers.User{Address: address, Age: ageInt})
		if err != nil {
			returnError(context, http.StatusInternalServerError, err)
			return
		}
		user, ok := user.(*modelusers.User)
		if !ok {
			returnError(context, http.StatusInternalServerError, errors.New("modify user failed"))
			return
		}

		// 返回添加后的响应数据
		returnSuccessfully(context, http.StatusCreated, gin.H{
			"describe": "User modified data",
			"user":     user,
		})
	}
}

// DeleteUser删除用户API接口处理函数
func DeleteUser() gin.HandlerFunc {
	return func(context *gin.Context) {
		id := context.Param("id")

		// 转换ID为整型
		idInt, err := strconv.Atoi(id)
		if err != nil {
			returnError(context, http.StatusBadRequest, err)
			return
		}

		// 修改用户
		ok, err := users.Delete(idInt)
		if err != nil {
			returnError(context, http.StatusInternalServerError, err)
			return
		}

		// 获取删除状态描述信息
		status := "failed"
		if ok {
			status = "successfully"
		}

		// 返回添加后的响应数据
		returnSuccessfully(context, http.StatusNoContent, gin.H{
			"describe": "User delete for " + status,
			"user":     nil,
		})
	}
}

// HeadUser获取用户资源的元数据
func HeadUser() gin.HandlerFunc {
	return func(context *gin.Context) {
		heads := users.Head()
		context.Writer.Header().Set("Content-Length", heads["Content-Length"])
		context.Writer.Header().Set("Content-Type", heads["Content-Type"])
		return
	}
}

// OptionsUser获取信息，关于用户资源的哪些属性是客户端可以改变的
func OptionsUser() gin.HandlerFunc {
	return func(context *gin.Context) {
		options, err := users.Options()
		if err != nil {
			returnError(context, http.StatusInternalServerError, err)
			return
		}
		returnSuccessfully(context, http.StatusOK, gin.H{
			"CanChange": options,
		})
	}
}

// returnSuccessfully返回请求成功响应数据
func returnSuccessfully(context *gin.Context, status int, data gin.H) {
	if data != nil {
		data["code"] = status
	}
	context.JSON(http.StatusOK, data)
}

// returnError返回请求错误响应数据
func returnError(context *gin.Context, status int, err error) {
	context.JSON(status, gin.H{"error": err.Error(), "code": status})
}
