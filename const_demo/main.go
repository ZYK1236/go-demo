package main

const (
	// const () 中可以写关键字 iota，每行累加1，第一行默认0
	zyk = iota
	mxw
	zkw
)

const (
	// 如果第一行修改默认值，则后续的值都会和第一行进行相同的处理
	// 例如 zyk2 = 0 * 10 = 0, mxw2 = 1 * 10 = 10
	zyk2 = iota * 10
	mxw2
	zkw2
)

func main() {
	const length int = 10
	var param int = 10
	param -= 1
	print(zkw2)
}