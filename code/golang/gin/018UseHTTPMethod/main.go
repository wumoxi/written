package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/wumoxi/toolkit"
)

// RESTFullAPI为RESTFullAPI接口类型
type RESTFullAPI interface {
	GetAll() interface{}                          // 获取所有资源(对于RESTFull中的GET)
	GetOne(int) (interface{}, error)              // 获取一条资源(对于RESTFull中的GET)
	Add(interface{}) (interface{}, error)         // 添加一条资源(对于RESTFull中的POST)
	Change(int, interface{}) (interface{}, error) // 修改一条资源(提供资源的完整信息对于RESTFull中的PUT)
	Modify(int, interface{}) (interface{}, error) // 修改一条资源(提供资源的完整信息对于RESTFull中的PATCH)
	Delete(int) (bool, error)                     // 删除一条资源
	Options() (map[string]interface{}, error)     // 获取信息，关于资源的哪些属性是客户端可以改变的
	Head() map[string]string                      // 获取资源元信息
}

// User用户类型声明
type User struct {
	Id      int    `json:"id" canChangeMethod:"-"`
	Name    string `json:"name" canChangeMethod:"change"`
	Email   string `json:"email" canChangeMethod:"change"`
	Phone   string `json:"phone" canChangeMethod:"change"`
	Sex     string `json:"sex" canChangeMethod:"change"`
	Age     int    `json:"age" canChangeMethod:"change,modify"`
	Address string `json:"address"  canChangeMethod:"change,modify"`
}

// userToInterface将用户类型转换为接口类型
func (u *User) userToInterface() interface{} {
	var o interface{} = u
	return o
}

// Users所有用户类型声明
type Users map[int]*User

// usersToInterface将所有用户类型转换为接口类型
func (u *Users) usersToInterface() interface{} {
	var o interface{} = u
	return o
}

// nextUserId获取一下个将要添加用户的ID
func (u *Users) nextUserId() int {
	if len(*u) == 0 {
		return 1
	}
	return len(*u) + 1
}

// checkUserExists检测用户是否已经存在
func (u *Users) checkUserExists(name string) bool {
	for _, user := range *u {
		if user.Name == name {
			return true
		}
	}
	return false
}

// GetAll获取所有用户
func (u *Users) GetAll() interface{} {
	return u.usersToInterface()
}

// GetOne获取一条用户
func (u *Users) GetOne(id int) (interface{}, error) {
	if user, ok := (*u)[id]; ok {
		return user.userToInterface(), nil
	}
	return nil, errors.New("user does not exist")
}

// Add添加一条用户
func (u *Users) Add(params interface{}) (interface{}, error) {
	user, ok := params.(*User)
	if !ok {
		return nil, errors.New("invalid parameter type")
	}

	// 检测是否存在
	exists := u.checkUserExists(user.Name)
	if exists {
		return nil, errors.New(fmt.Sprintf("user exists for name, name is %s", user.Name))
	}

	user.Id = u.nextUserId()
	(*u)[user.Id] = user
	return user.userToInterface(), nil
}

// Change修改一条用户
func (u *Users) Change(id int, params interface{}) (interface{}, error) {
	user, ok := (*u)[id]
	if !ok {
		return nil, errors.New("user does not exist")
	}
	cu := params.(*User)
	user.Name = cu.Name
	user.Email = cu.Email
	user.Phone = cu.Phone
	user.Sex = cu.Sex
	user.Age = cu.Age
	user.Address = cu.Address
	return user.userToInterface(), nil
}

// Modify修改一条用户
func (u *Users) Modify(id int, params interface{}) (interface{}, error) {
	user, ok := (*u)[id]
	if !ok {
		return nil, errors.New("user does not exist")
	}
	cu := params.(*User)
	user.Age = cu.Age
	user.Address = cu.Address
	return user.userToInterface(), nil
}

// Delete删除一条用户
func (u *Users) Delete(id int) (bool, error) {
	_, ok := (*u)[id]
	if !ok {
		return false, errors.New("user does not exist")
	}
	users := *u
	delete(users, id)
	return true, nil
}

// Options获取信息，关于用户资源的哪些属性是客户端可以改变的
func (u *Users) Options() (map[string]interface{}, error) {
	// 通过标签值获取字段名数组
	fieldsByChange, err := toolkit.GetFieldNameByTagValue(User{}, "change")
	if err != nil {
		return nil, err
	}

	// 通过标签值获取字段名数组
	fieldsByModify, err := toolkit.GetFieldNameByTagValue(User{}, "modify")
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"Endpoint /users/:id for HTTP PUT method":   fieldsByChange,
		"Endpoint /users/:id for HTTP PATCH method": fieldsByModify,
	}, nil
}

// Head返回用户资源元信息
func (u *Users) Head() map[string]string {
	marshal, _ := json.Marshal(u.GetAll())
	l := len(marshal)
	h := make(map[string]string)
	h["Content-Length"] = strconv.Itoa(l)
	h["Content-Type"] = "application/json"
	return h
}

// 存储用户列表
var users = make(Users)

// 获取所有用户API接口处理函数
func GetUserAll() gin.HandlerFunc {
	return func(context *gin.Context) {
		returnSuccessfully(context, http.StatusOK, gin.H{
			"describe": "User List data",
			"users":    users.GetAll(),
		})
	}
}

// 获取单个用户API接口处理函数
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
		user, ok := user.(*User)
		if !ok {
			returnError(context, http.StatusInternalServerError, err)
		}

		returnSuccessfully(context, http.StatusOK, gin.H{
			"describe": "User detail data",
			"user":     user,
		})
	}
}

// 添加用户API接口处理函数
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
		user, err := users.Add(&User{Name: name, Phone: phone, Email: email, Sex: sex, Address: address, Age: ageInt})
		if err != nil {
			returnError(context, http.StatusInternalServerError, err)
			return
		}
		user, ok := user.(*User)
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

// 修改用户API接口处理函数
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
		user, err := users.Change(idInt, &User{Name: name, Phone: phone, Email: email, Sex: sex, Address: address, Age: ageInt})
		if err != nil {
			returnError(context, http.StatusInternalServerError, err)
			return
		}
		user, ok := user.(*User)
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

// 修改用户API接口处理函数
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
		user, err := users.Modify(idInt, &User{Address: address, Age: ageInt})
		if err != nil {
			returnError(context, http.StatusInternalServerError, err)
			return
		}
		user, ok := user.(*User)
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

// 获取信息，关于用户资源的哪些属性是客户端可以改变的
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

func main() {
	// 用默认的中间件创建一个杜松子酒路由器(记录器logger和恢复recovery（无崩溃）中间件)
	router := gin.Default()

	// 获取所有用户信息
	router.GET("/users", GetUserAll())
	// 获取一个用户信息
	router.GET("/users/:id", GetOne())
	// 添加一个用户
	router.POST("/users", AddUser())
	// 修改一个用户完整信息
	router.PUT("/users/:id", ChangeUser())
	// 修改一个用户部分信息
	router.PATCH("/users/:id", ModifyUser())
	// 删除一个用户
	router.DELETE("/users/:id", DeleteUser())
	// 获取用户资源的元数据
	router.HEAD("/users", HeadUser())
	// 获取信息，关于用户资源的哪些属性是客户端可以改变的。
	router.OPTIONS("/users", OptionsUser())

	// 监听并在0.0.0.0:8080上启动服务
	log.Fatal(router.Run(":8080"))
}
