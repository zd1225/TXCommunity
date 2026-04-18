package main
import "fmt"

type Counter struct {
    count int
}

func (c *Counter) Inc() *Counter {
    c.count++
    return c  // 返回指针
}

func main() {
    // ✅ 可以：单次调用，Go 自动取地址
    c1 := Counter{count: 0}
    c1.Inc()  // 自动转为 (&c1).Inc()
    fmt.Println(c1.count)  // 1
    
    // ✅️ 链式调用也可以，有没有&符号都可以
    c2 := Counter{count: 0}//或者c2 := &Counter{count: 0}
    c2.Inc().Inc()  
	fmt.Println(c2.count)

	//❌️Cirle{pi : 3.14}这是个匿名变量，没有变量指向它，不可寻址
	//Counter{count: 100}.count = 99
}

