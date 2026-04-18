package main
import "fmt"
import "errors"

func main(){
	fmt.Println("Hello GO 我是iOS转后端第一天。")
	fmt.Println("加油呀！")

	//声明变量
	var name string
	name = "iOS 转后端"
	fmt.Println(name)

	//自动推到类型（最常用）所以这里age是个变量
	age := 25
	fmt.Println(age)
	age = 12
	fmt.Println(age)

	a,b := 10, 20
	fmt.Println(a, b)

	//常量
	const pi = 3.1415926
	fmt.Println(pi)
	/***
	 *不可寻址变量不能更改
	pi = 3.4
	fmt.Println(pi)
	*/
	
//解Go语言中的结构体字面量和结构体方法,需要可寻址的变量

	/*方式一：
	这种方式不可以，因为要修改字面量，也就是临时变量，或者匿名变量
	go 语言不允许
	*
	Cirle{pi : 3.14} = 3.1415 
	*/

	/*方式二
	这种方式允许，因为没有更改字面量
	* 这个操作并没有修改原始的字面量
	（因为字面量是临时值，且方法接收者是值传递，所以修改的是副本），
	而是返回了一个新的Circle实例。
	*/

	c := Cirle{pi : 3.14}.WithPi(3.1415926)
	fmt.Println(c.pi)

	fmt.Println(add(3, 5))
	fmt.Println(swap(3, 5))
	fmt.Println(devide(3, 3))
	fmt.Println(devide(3, 0))


	ios_name := "jodan"
	ios_age := 24
	ios_job := "golang 开发者"
	fmt.Println(ios_name, ios_age, ios_job)
	
	var max_value int  = max(6, 100)
	// max_value := max(5, 7)
	fmt.Println(max_value)
	fmt.Println(add2(5, 6))
	fmt.Println(sub(7, 4))
	fmt.Println(mul(4, 5))

	lst := []int{5, 7, 8, 10}
	result, err := query(lst, 5)
	fmt.Println(result, err)

}

func query(slice []int, a int) (int, error) {
	if contains(slice, a) {
		return 1, nil
	}
	return 0, errors.New("不包含此元素")
}

func contains[T comparable](slice []T, target T) bool {
	for _, item := range slice {
		if item == target {
			return true
		}
	}
	return false
}

func add2(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func mul(a, b int) int {
	return a * b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}




// 普通函数
func add(a, b int) int {
	return a + b
}

// 交换
func swap(a, b int) (int, int) {
	return b, a
}

//除法
func devide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("除数不能为0")
	}
	return a / b, nil
}

type Cirle struct{
	pi float64
}
func (c Cirle) WithPi(newPi float64) Cirle{
	c.pi = newPi
	return c
}