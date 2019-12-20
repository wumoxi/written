# 使用HTTP方法构建用户管理RESTFulAPI

Gin框架对[RESTfulAPI](http://www.ruanyifeng.com/blog/2014/05/restful_api.html)设计支持比较友好，下面简单做一个用户管理API服务，这个服务没有将用户写入在存储系统如MYSQL或都是一些其它的存储服务，而是直接使用Map进行存储到内存中，服务重启数据将会丢失，不过在这里这是一个Demo真的是项目中不会这么干，也绝对不允许这么做，对于这个 Demo 演示RESTFulAPI是够用了！

由于这个 Demo 实现篇幅比较长，所有会分段讲解，请接着向下看，如果你都能理解了，肯定会有收获。

## 定义RESTFulAPI接口

定义这么一个接口，不光用户管理可以用，其它模块也可以用，那么样一来，每一个模块就可以快速的来实现其功能，如果功能不够用并且其他几个别的模块都需要这些功能也可以添加其它的接口，一个模块可以可以多个接口，多个接口之间也可以进行组合使用。

```go
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
```

## 定义用户模型及操作方法

在用户模型定义上，Users类型实现了RESTFullAPI接口，并且添加了`usersToInterface`, `nextUserId`, `checkUserExists`三个私有方法它们的功能分别依次是`将Users类型转换为interface{}接口类型`、`获取下一个将要添加用户的ID值`、`检查用户是否存在`，这三个方法为辅助方法。与此同时也为`User`结构添加了`userToInterface`方法，用于将用户类型转换为`interface{}`接口类型。

这一部分是数据操作，如果进一步进行优化的话，可以将其放入一个单独的包如：`model/users/users.go`

```go
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

// GetOne获取一个用户
func (u *Users) GetOne(id int) (interface{}, error) {
	if user, ok := (*u)[id]; ok {
		return user.userToInterface(), nil
	}
	return nil, errors.New("user does not exist")
}

// Add添加一个用户
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

// Change修改一个用户
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

// Modify修改一个用户
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

// Delete删除一个用户
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
```

## 定义RESTFulAPI接口处理函数

这几个处理函数与RESTFulAPI方法关于对应如下:

| 处理函数    | RESTFullAPI接口方法 | 处理函数功能描述                                          |
| ----------- | ------------------- | --------------------------------------------------------- |
| GetUserAll  | GET `/users`        | GetAll获取所有用户                                        |
| GetOne      | GET `/users/:id`    | GetOne获取一个用户                                        |
| AddUser     | POST `/users`       | Add添加一个用户                                           |
| ChangeUser  | PUT `/users/:id`    | Change修改一个用户                                        |
| ModifyUser  | PUT `/users/:id`    | Modify修改一个用户                                        |
| DeleteUser  | DELETE `/users/:id` | Delete删除一个用户                                        |
| HeadUser    | HEAD `/users`       | Options获取信息，关于用户资源的哪些属性是客户端可以改变的 |
| OptionsUser | OPTIONS `/users`    | Head返回用户资源元信息                                    |

并且添加了两个辅助函数，分别是 `returnSuccessfully` 和 `returnError`，它们分别用于统一的`错误处理`和`成功响应`数据的返回。

这一部分是API处理函数定义，如果进一步进行优化的话，可以将其放入一个单独的包如：`api/users/users.go`

```go
// users存储用户列表
var users = make(Users)

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
```

## 注册用户管理路由

```go
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
```

## 运行WEB服务器并通过PostMan进行测试

用户模型有单独的测试`main_test.go`文件，当然了你也可以通过单元测试对 API 接口进行测试，不过这里使用 PostMan 进行测试。

```shell
$ go run main.go
``` 

### 使用PostMan测试添加用户`POST`API接口

![添加用户API接口](https://lucklit.oss-cn-beijing.aliyuncs.com/written/Snip20191220_74.png)

### 使用PostMan测试修改一个用户完整信息`PUT`API接口


![添加用户API接口](https://lucklit.oss-cn-beijing.aliyuncs.com/written/Snip20191220_75.png)


### 使用PostMan测试修改一个用户部分信息`PATCH`API接口

![添加用户API接口](https://lucklit.oss-cn-beijing.aliyuncs.com/written/Snip20191220_76.png)

### 使用PostMan测试获取所有用户`GET`API接口

![获取所有用户`](https://lucklit.oss-cn-beijing.aliyuncs.com/written/Snip20191220_81.png)

### 使用PostMan测试获取一个用户`GET`API接口

![获取所有用户`](https://lucklit.oss-cn-beijing.aliyuncs.com/written/Snip20191220_84.png)

### 使用PostMan测试获取用户资源的元数据`HEAD`API接口

![获取用户资源的元数据](https://lucklit.oss-cn-beijing.aliyuncs.com/written/Snip20191220_89.png)

请注意，这里只有 `Response Size` 大小，而没有Body大小，因为在 HEAD 请求方式中无需返回任何响应数据，这就是它与 GET 请求方式最大的区别！

### 使用PostMan测试获取信息，关于用户资源的哪些属性是客户端可以改变的。`OPTIONS`API接口

![获取信息，关于用户资源的哪些属性是客户端可以改变的](https://lucklit.oss-cn-beijing.aliyuncs.com/written/Snip20191220_88.png)


## 目录

[BACK](../gin-use.md)