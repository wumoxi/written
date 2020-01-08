# Golang『结构体、集合和高阶函数』注意点

通常你在应用中定义了一个结构体，那么你也可能需要这个结构体的（指针）对象集合，比如：

```go
// 任意类型
type Any interface{}

// 车辆类型
type Car struct {
	Model        string // 车型
	Manufacturer string // 制造商
	BuildYear    int    // 制造年份
}

// 车辆集合类型
type Cars []*Car
```

然后我们就可以使用高阶函数，实际上也就是把函数作为定义所需方法（其他函数）的参数，例如：

1）定义一个通用的 `Process()` 函数，它接收一个作用于每一辆 car 的 f 函数作参数：

```go
// 通过给定的函数处理所有车辆
func (cs Cars) Process(f func(car *Car)) {
	for _, c := range cs {
		f(c)
	}
}
```

2）在上面的基础上，实现一个查找函数来获取子集合，并在 `Process()` 中传入一个闭包执行（这样就可以访问局部切片 `cars`）：

```go
// 通过给定的函数查找符合条件的车辆
func (cs Cars) FindAll(f func(car *Car) bool) Cars {
	cars := make([]*Car, 0)
	cs.Process(func(car *Car) {
		if f(car) {
			cars = append(cars, car)
		}
	})
	return cars
}
```

3）实现 Map 功能，产出除 car 对象以外的东西：

```go
// 处理所有车辆并返回新数据
func (cs Cars) Map(f func(car *Car) Any) []Any {
	r := make([]Any, len(cs))
	ix := 0
	cs.Process(func(car *Car) {
		r[ix] = f(car)
		ix++
	})
	return r
}
```

现在我们可以定义下面这样的具体查询：

```go
allNewBMWs := allCars.FindAll(func(car *Car) bool {
	return (car.Manufacturer == "BMW") && (car.BuildYear > 2010)
})
```

4）我们也可以根据参数返回不同的函数。也许我们想根据不同的厂商添加汽车到不同的集合，但是这（这种映射关系）可能会是会改变的。所以我们可以定义一个函数来产生特定的添加函数和 map 集：

```go
func MakeSortedAppender(manufacturers []string) (func(car *Car), map[string]Cars) {
	// 准备分类的汽车Map。
	sortedCars := make(map[string]Cars)

	for _, m := range manufacturers {
		sortedCars[m] = make([]*Car, 0)
	}
	sortedCars["Default"] = make([]*Car, 0)

	// 准备追加器函数
	appender := func(car *Car) {
		if _, ok := sortedCars[car.Manufacturer]; ok {
			sortedCars[car.Manufacturer] = append(sortedCars[car.Manufacturer], car)
		} else {
			sortedCars["Default"] = append(sortedCars["Default"], car)
		}
	}
	return appender, sortedCars
}
```

现在我们可以用它把汽车分类为独立的集合，像这样：

```go
manufacturers := []string{"Ford", "Aston Martin", "Land Rover", "BMW", "Jaguar"}
sortedAppender, sortedCars := MakeSortedAppender(manufacturers)
allUnsortedCars.Process(sortedAppender)
BMWCount := len(sortedCars["BMW"])
```

我们让这些代码在下面的程序 cars.go 中执行：

示例 11.18 [cars.go](examples/chapter_11/cars.go)：

```go
package main

import "fmt"

// 任意类型
type Any interface{}

// 车辆类型
type Car struct {
	Model        string // 车型
	Manufacturer string // 制造商
	BuildYear    int    // 制造年份
}

// 车辆集合类型
type Cars []*Car

// 通过给定的函数处理所有车辆
func (cs Cars) Process(f func(car *Car)) {
	for _, car := range cs {
		f(car)
	}
}

// 通过给定的函数查找符合条件的车辆
func (cs Cars) FindAll(f func(car *Car) bool) Cars {
	cars := make([]*Car, 0)
	cs.Process(func(car *Car) {
		if f(car) {
			cars = append(cars, car)
		}
	})
	return cars
}

// 处理所有车辆并返回新数据
func (cs Cars) Map(f func(car *Car) Any) []Any {
	r := make([]Any, len(cs))
	ix := 0
	cs.Process(func(car *Car) {
		r[ix] = f(car)
		ix++
	})
	return r
}

func MakeSortedAppender(manufacturers []string) (func(car *Car), map[string]Cars) {
	// 准备分类的汽车Map。
	sortedCars := make(map[string]Cars)

	for _, m := range manufacturers {
		sortedCars[m] = make([]*Car, 0)
	}
	sortedCars["Default"] = make([]*Car, 0)

	// 准备追加器函数
	appender := func(car *Car) {
		if _, ok := sortedCars[car.Manufacturer]; ok {
			sortedCars[car.Manufacturer] = append(sortedCars[car.Manufacturer], car)
		} else {
			sortedCars["Default"] = append(sortedCars["Default"], car)
		}
	}
	return appender, sortedCars
}

func main() {
	// 创建一些车辆
	ford := &Car{"Fiesta", "Ford", 2008}
	bmw := &Car{"XL 450", "BMW", 2011}
	merc := &Car{"D600", "Mercedes", 2009}
	bmw2 := &Car{"X 800", "BMW", 2008}

	// 查询
	allCars := Cars([]*Car{ford, bmw, merc, bmw2})
	allNewBMWs := allCars.FindAll(func(car *Car) bool {
		return (car.Manufacturer == "BMW") && (car.BuildYear > 2010)
	})
	fmt.Printf("AllCars: %v\n", allCars)
	fmt.Printf("New BMWs: %v\n", allNewBMWs)

	// 制造商
	manufacturers := []string{"Ford", "Aston Martin", "Land Rover", "BMW", "Jaguar"}
	sortedAppender, sortedCars := MakeSortedAppender(manufacturers)
	allCars.Process(sortedAppender)
	fmt.Printf("Map sortedCars: %v\n", sortedCars)
	BMWCount := len(sortedCars["BMW"])
	fmt.Printf("We have %d BMWs\n", BMWCount)
}
```

输出：

```
AllCars: [0xc000070180 0xc0000701b0 0xc0000701e0 0xc000070210]
New BMWs: [0xc0000701b0]
Map sortedCars: map[Aston Martin:[] BMW:[0xc0000701b0 0xc000070210] Default:[0xc0000701e0] Ford:[0xc000070180] Jaguar:[] Land Rover:[]]
We have 2 BMWs
```


## 目录
[Back](../GolangNotice.md)
