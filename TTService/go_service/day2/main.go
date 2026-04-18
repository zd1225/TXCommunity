package main
import (
	"fmt"
	//"errors"
)

func main() {
	//创建 类型、长度、容量
	s:= make([]int, 0, 10)
	//添加元素
	s = append(s, 1, 2, 3)
	// 截取
	sub := s[0:2]
	// 赋值，必须用copy，不能直接=（因为用=会共享内存，修改一处，另一处也会变）
	dst := make([]int, len(s))
	copy(dst, s)
	fmt.Println(s, sub, dst)


	//arr := make([]int{12,12,12}, 0, 10) 不可以添加{12,12,12}，不属于一个类型
	arr := make([]int, 0, 10)
	//返回的是新的切片
	arr = append(arr, 1, 2, 3)
	subArr := arr[0:2]
	desArr := make([]int, len(arr))
	copy(desArr, arr)
	fmt.Println(arr, subArr, desArr)

	
	m := map[string]int{
		"apple": 10,
	}

	//判断key是否存在
	_, ok := m["pear"]

	if !ok {
		fmt.Println("key不存在")
	}
	val, key := m["apple"]
	fmt.Println("value的值:", val,"key是否存在：", key)

	//嵌套
	userTag := map[string][]string{
		"app": {"start", "pause", "stop"},
	}
	fmt.Println(userTag["app"])

	mm := map[string]int{
		"apple": 10,
		"pitch": 20,
	}

	_, ook := mm["pear"]
	if !ook {
		fmt.Println("key值不存在")
	}
	users := map[string][]string{
		"users": {"jo", "tom", "lucy"},
	}
	fmt.Println(users["users"])


	// 指针
	a := 3
	b := 4

	var p1 *int = &a
	p2 := &b
	// swap(&a, &b)
	swap(p1, p2)
	fmt.Println("下面是指针交换：")
	fmt.Println(a, b)

	// 结构体(下面不可以，不可对不可寻址的变量赋值)
	//User{ID: 1001, Name: "jodan", Age: 28}.Name = "Lisa"
	
}

type User struct{
	ID int
	Name string
	Age int
}

//这里实际上生成了两个临时变量temp1=*b, temp2=*a;
func swap(a, b *int){
	*a, *b = *b, *a
}