package middleware_demo

type HandleFunc func(*Request)

type Request struct {
	index      int
	middleware []HandleFunc
}

// 生成请求
func NewRequest() (request *Request) {
	request = &Request{
		index:      0,
		middleware: make([]HandleFunc, 0),
	}
	return
}

// 注册中间件
func (request *Request) RegisterMiddleware(middleware ...HandleFunc) {
	for _, mid := range middleware {
		request.middleware = append(request.middleware, mid)
	}
}

// 执行中间件
func (request *Request) Next() {
	index := request.index
	if index >= len(request.middleware) {
		return
	}

	request.index++
	request.middleware[index](request)
}
