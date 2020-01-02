package modelusers

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"toolkit"
)

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
func (u Users) usersToInterface() interface{} {
	var o interface{} = u
	return o
}

// nextUserId获取一下个将要添加用户的ID
func (u Users) nextUserId() int {
	if len(u) == 0 {
		return 1
	}
	return len(u) + 1
}

// checkUserExists检测用户是否已经存在
func (u Users) checkUserExists(name string) bool {
	for _, user := range u {
		if user.Name == name {
			return true
		}
	}
	return false
}

// GetAll获取所有用户
func (u Users) GetAll() interface{} {
	return u.usersToInterface()
}

// GetOne获取一个用户
func (u Users) GetOne(id int) (interface{}, error) {
	if user, ok := u[id]; ok {
		return user.userToInterface(), nil
	}
	return nil, errors.New("user does not exist")
}

// Add添加一个用户
func (u Users) Add(params interface{}) (interface{}, error) {
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
	u[user.Id] = user
	return user.userToInterface(), nil
}

// Change修改一个用户
func (u Users) Change(id int, params interface{}) (interface{}, error) {
	user, ok := u[id]
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
func (u Users) Modify(id int, params interface{}) (interface{}, error) {
	user, ok := u[id]
	if !ok {
		return nil, errors.New("user does not exist")
	}
	cu := params.(*User)
	user.Age = cu.Age
	user.Address = cu.Address
	return user.userToInterface(), nil
}

// Delete删除一个用户
func (u Users) Delete(id int) (bool, error) {
	_, ok := u[id]
	if !ok {
		return false, errors.New("user does not exist")
	}
	users := u
	delete(users, id)
	return true, nil
}

// Options获取信息，关于用户资源的哪些属性是客户端可以改变的
func (u Users) Options() (map[string]interface{}, error) {
	// 通过标签值获取字段名数组
	fieldsByChange, err := toolkit.GetFieldsNameByTag(User{}, "change")
	if err != nil {
		return nil, err
	}

	// 通过标签值获取字段名数组
	fieldsByModify, err := toolkit.GetFieldsNameByTag(User{}, "modify")
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"Endpoint /users/:id for HTTP PUT method":   fieldsByChange,
		"Endpoint /users/:id for HTTP PATCH method": fieldsByModify,
	}, nil
}

// Head返回用户资源元信息
func (u Users) Head() map[string]string {
	marshal, _ := json.Marshal(u.GetAll())
	l := len(marshal)
	h := make(map[string]string)
	h["Content-Length"] = strconv.Itoa(l)
	h["Content-Type"] = "application/json"
	return h
}
