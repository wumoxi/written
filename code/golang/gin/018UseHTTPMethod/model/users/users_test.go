package modelusers

import (
	"fmt"
	"testing"
)

var users = make(Users)

func TestUser(t *testing.T) {
	add(t)
	getAll(t)
	getOne(t)
	change(t)
	modify(t)
	del(t)
}

// add测试添加用户
func add(t *testing.T) {
	var expected1 = &User{Id: 1, Name: "张三", Email: "zhangsan@foxmail.com", Phone: "13792992888", Sex: "male", Age: 18, Address: "北京市东城区永外大街182号."}
	var expected2 = &User{Id: 2, Name: "李四", Email: "lisi@foxmail.com", Phone: "13792998888", Sex: "male", Age: 25, Address: "北京市东城区永外大街32号."}
	// 添加用户
	user, err := users.Add(&User{Name: "张三", Email: "zhangsan@foxmail.com", Phone: "13792992888", Sex: "male", Age: 18, Address: "北京市东城区永外大街182号."})
	if err != nil {
		t.Errorf("add user erro: %s\n", err.Error())
	}
	actual, ok := user.(*User)
	if !ok {
		t.Errorf("invalid parameter type")
	}
	// 验证添加
	if *actual != *expected1 {
		t.Errorf("expected: %+v, got: %+v\n", expected1, actual)
	}

	// 添加用户
	user, err = users.Add(&User{Name: "李四", Email: "lisi@foxmail.com", Phone: "13792998888", Sex: "male", Age: 25, Address: "北京市东城区永外大街32号."})
	if err != nil {
		t.Errorf("add user erro: %s\n", err.Error())
	}
	actual, ok = user.(*User)
	if !ok {
		t.Errorf("invalid parameter type")
	}
	// 验证添加
	if *actual != *expected2 {
		t.Errorf("expected: %+v, got: %+v\n", expected2, actual)
	}
}

// getAll测试获取所有用户
func getAll(t *testing.T) {
	var expected = make(Users)
	expected[1] = &User{Id: 1, Name: "张三", Email: "zhangsan@foxmail.com", Phone: "13792992888", Sex: "male", Age: 18, Address: "北京市东城区永外大街182号."}
	expected[2] = &User{Id: 2, Name: "李四", Email: "lisi@foxmail.com", Phone: "13792998888", Sex: "male", Age: 25, Address: "北京市东城区永外大街32号."}

	if len(expected) != len(users) {
		t.Errorf("result shuld %d, got: %d\n", len(expected), len(users))
	}

	users, ok := users.GetAll().(Users)
	if !ok {
		t.Errorf("invalid parameter type")
	}
	for id, user := range users {
		if id == 1 {
			if *user != *expected[id] {
				t.Errorf("result shuld %+v, got: %+v\n", *user, *expected[id])
			}
		}
	}
}

// getOne测试获取一个用户
func getOne(t *testing.T) {
	var expected = &User{Id: 1, Name: "张三", Email: "zhangsan@foxmail.com", Phone: "13792992888", Sex: "male", Age: 18, Address: "北京市东城区永外大街182号."}

	user, err := users.GetOne(1)
	if err != nil {
		t.Errorf("get one user error: %s\n", err.Error())
	}

	actual, ok := user.(*User)
	if !ok {
		t.Errorf("invalid parameter type")
	}

	// 验证添加
	if *actual != *expected {
		t.Errorf("expected: %+v, got: %+v\n", expected, actual)
	}
}

// change测试修改一个用户
func change(t *testing.T) {
	var expected = &User{Id: 1, Name: "张三", Email: "zhangsan@126.com", Phone: "13728298283", Sex: "male", Age: 22, Address: "北京市东城区永外大街54号."}

	// 添加用户
	user, err := users.Change(1, &User{Name: "张三", Email: "zhangsan@126.com", Phone: "13728298283", Sex: "male", Age: 22, Address: "北京市东城区永外大街54号."})
	if err != nil {
		t.Errorf("change user erro: %s\n", err.Error())
	}

	actual, ok := user.(*User)
	if !ok {
		t.Errorf("invalid parameter type")
	}

	// 验证修改
	if *actual != *expected {
		t.Errorf("expected: %+v, got: %+v\n", expected, actual)
	}
}

// change测试修改一个用户
func modify(t *testing.T) {
	var expected = &User{Id: 1, Name: "张三", Email: "zhangsan@126.com", Phone: "13728298283", Sex: "male", Age: 12, Address: "北京市东城区永外大街18号."}

	// 添加用户
	user, err := users.Modify(1, &User{Age: 12, Address: "北京市东城区永外大街18号."})
	if err != nil {
		t.Errorf("modify user erro: %s\n", err.Error())
	}

	actual, ok := user.(*User)
	if !ok {
		t.Errorf("invalid parameter type")
	}

	// 验证修改
	if *actual != *expected {
		t.Errorf("expected: %+v, got: %+v\n", expected, actual)
	}
}

// del测试删除一个用户
func del(t *testing.T) {
	// 添加用户
	ok, err := users.Delete(1)
	if err != nil {
		t.Errorf("delete user erro: %s\n", err.Error())
	}

	// 验证修改
	if ok != true {
		t.Errorf("expected: %+v, got: %+v\n", true, ok)
	}

	// 输出当前所有用户
	users, ok := users.GetAll().(Users)
	if !ok {
		fmt.Printf("invalied parameter type")
	}
	for id, user := range users {
		// 肯定只会输出一条记录因为在这个测试中一共添加了两个用户，一直对ID为1的记录进行操作，所以结果一定是：
		// user id: 2, user: &{Id:2 Name:李四 Email:lisi@foxmail.com Phone:13792998888 Sex:male Age:25 Address:北京市东城区永外大街32号.}
		t.Logf("user id: %d, user: %+v\n", id, *user)
	}
}
