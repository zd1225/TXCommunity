package main
import (
	"fmt"
	"strings"
	"sort"
)

// range 是语言内置的语法，不是标准库类型
// 编译器直接生成遍历代码

type User struct {
	ID int
	Name string
	Age int
}
func main(){
	users := []User{
		{ID: 1001, Name: "jo", Age: 4},
		{ID: 1002, Name: "tom", Age: 55},
		{ID: 1003, Name: "lily", Age: 43},
		{ID: 1004, Name: "lucy", Age: 45},
		{ID: 1005, Name: "lulu", Age: 45},
	}

	// 统计每个年龄出现的次数
	seen := make(map[int]int)
	for _, item := range users {
		//方法一：
		// seen[item.Age]++
		//方法二:
		if _, ok := seen[item.Age]; ok {
			seen[item.Age]++
		}else{
			seen[item.Age] = 1
		}
	}
	fmt.Println("用户年龄重复数：", seen)

	limitID := 1004
	res, user := Find(users, limitID, func(sta int, user User) bool{
		if user.ID == sta {
			return true
		}
		return false
	})

	if res {
		fmt.Println("查找用户：",user)
	}else{
		fmt.Println("没有该用户")
	}

	limitAge := 66
	adults := Filter(users, limitAge, func(sta int, user User) bool{
		if user.Age >= sta {
			return true
		}
		return false
	})
	fmt.Println("筛选出的用户：", adults)

	customers := Map(users, func(user User) string{
		return strings.ToUpper(user.Name)
	})
	fmt.Println("映射出的用户：",customers)

	//按照年龄从小到大排序
	sort.Slice(users, func(i,j int)bool {
		return users[i].Age < users[j].Age
	})
	fmt.Println("按照年龄排序后所有用户：",users)

	ages := Reduce(users, 0, func(sum int, user User) int {
		return sum + user.Age
	})
	fmt.Println("聚合出的用户：",ages)

}

func Find[T any, R any](slice []T, state R, f func(sta R, content T) bool) (bool, T) {
	for _, item := range slice {
		if f(state, item) {
			return true, item
		}
	}
	var zero T
	return false, zero
}

func Filter[T any, R any](slice []T, state R,  f func(sta R, content T) bool) []T {
	var results []T
	//res = make([]T, 0, 10)
	fmt.Println("验证空数组：",results)
	fmt.Println(results == nil)
	fmt.Printf("%#v\n",results)
	for _, item := range slice {
		if(f(state, item)){
			results = append(results, item)
		}
	}
	return results
}

func Map[T any, R any](slice []T, f func(content T) R) []R {
	var results []R
	for _, item := range slice {
		// var res R = f(item)
		res := f(item)
		results = append(results, res)
	}
	return results
}

func Reduce[T any, R any](slice []T, initvalue R, f func(sum R, content T) R) R {
	var result R = initvalue
	for _, item := range slice {
		result = f(result, item)
	}
	return result
}