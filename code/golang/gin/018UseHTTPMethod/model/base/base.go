package base

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
