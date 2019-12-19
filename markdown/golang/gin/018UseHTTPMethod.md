# 使用HTTP方法

Gin框架对RESTfulAPI设计支持比较友好，下面简单做一个用户管理API服务，这个服务没有将用户写入在存储系统如MYSQL或都是一些其它的存储服务，而是直接使用Slice进行存储到内存中，服务重启数据将会丢失，不过在这里这是一个Demo真的是项目中不会这么干，绝对允许这么做，对于这个 Demo 演示RESTFulAPI是够用了！

由于这个 Demo 实现篇幅比较长，所有会分段讲解，请接着向下看，如果你都能理解了，肯定会有不小的收获！Go!

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
```