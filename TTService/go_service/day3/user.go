package main

import(
	"fmt"
	"errors"
	"os"
	"time"
)

var ErrUserNotDefound = errors.New("该用户不存在")

func getUser(id int)(string, error) {
	if id <= 0 {
		return "", fmt.Errorf("参数错误:%w",ErrUserNotDefound)
	}
	return "查询成功", nil
}

func testPanic(){

	defer func(){
		//这里是捕获异常
		if err := recover(); err != nil{
			fmt.Println("恢复了：",err)
		}
	}()

	//可能发生崩溃的代码
	panic("这里发生了数据越界崩溃")
}

func main() {


	testPanic()
	fmt.Println("程序继续运行")

	_, err := getUser(0)
	if errors.Is(err, ErrUserNotDefound) {
		fmt.Println("捕获到错误:",err)
	}

	//打开文件
	file, err := os.Open("../question.txt")
	if err != nil {
		fmt.Println("打开文件失败：", err)
		return
	}
	defer file.Close()//确保文件被关闭

	data := make([]byte, 1024)
	n, err := file.Read(data)
	if err != nil {
		fmt.Println("读取文件失败：", err)
		return
	}
	fmt.Println("读取了 %d  字节:\n%s", n, string(data[:n]))


	createChan()

}


func createChan(){
	//无缓冲的channel，同步阻塞，需要发送端和接收端同时准备好
	ch1 := make(chan int)
	// ch2 := make(chan string, 5)
	
	// 发送方 放进 子 goroutine,否则会卡死，互相等待
	go func(){
		//发送数据
		ch1 <- 42//这里会阻塞，直到接收方接受解除阻塞
		fmt.Println("发送完成")
	}()

	// //发送数据
	// ch1 <- 42

	time.Sleep(5 * time.Second)
	//接受数据
	value := <-ch1//准备接受，解除阻塞
	time.Sleep(5 * time.Second)

	fmt.Println("接收到：", value)

	//打印结果：
	发送完成
	......
	5秒后
    接收到： 42

	// vallue, ok := <-ch1
	
	//关闭channel
	close(ch1)

	// //遍历channel
	// for v := range ch {
	// 	fmt.Println(v)
	// }

}