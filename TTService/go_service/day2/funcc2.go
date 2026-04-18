package main
import "fmt"

func main() {

	// var users []User = 
	// {
	// 	User(ID: 1001, Age: 22, Name: "jo"),
	// 	User(ID: 1002, Age: 24, Name: "tom"),
	// 	User(ID: 1003, Age: 25, Name: "lily"),
	// }

	var users []User = []User{
		{ID: 1001, Age: 22, Name: "jo"},
		{ID: 1002, Age: 24, Name: "tom"},
		{ID: 1003, Age: 25, Name: "lily"},
	}
	//查找指定人
	var account int = 1001
	customer, isSuc := Find(users, account, func(user User, acc int) bool {
		if user.ID == acc {
			return true
		}
		return false
	})

	if isSuc {
		fmt.Println(customer)
	} else {
		fmt.Println("id 不存在")
	}

	names := Map(users, func(user User) string {
		return user.Name
	})
	fmt.Println(names)

	ages := Reduce(users, 0, func(sum int, user User) int {
		return sum + user.Age
	})
	fmt.Println(ages)

	adults := Filter(users, func(user User) bool {
		return user.Age > 22
	})
	fmt.Println(adults)

}

type User struct {
	ID int
	Age int
	Name string
}

//

//聚合函数
func Reduce[T any, R any](slice []T, init R, f func(R, T) R) R {
	var result R = init
	for _, item := range slice {
		result = f(result, item)
	}
	return result
}

//Map函数
func Map[T any, R any](slice []T, f func(T) R) []R {
	var result []R
	for _, item := range slice {
		var s R = f(item)
		result = append(result, s)
	}
	return result
}

//Filter
func Filter[T any](slice []T, f func(T) bool) []T{
	var result []T
	for _, item := range slice {
		if f(item) {
			result = append(result, item)
		}
	}
	return result
}

//find函数
func Find[T any](slice []T, account int, f func(s T, acc int) bool) (T, bool) {
	for _, item := range slice {
		if f(item, account) {
			return item, true
		}
	}
	var zero T
	return zero, false
} 