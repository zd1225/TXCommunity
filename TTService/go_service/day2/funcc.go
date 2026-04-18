package main

import "fmt"

func main() {
	var users []int = []int{23, 35, 55, 66, 88}
	youngs := Filter(users, func(u int) bool{
		return u < 55
	})
	fmt.Println(youngs)

	result := Reduce(users, 0, func(sum int, u int) int {
		return sum + u
	})
	fmt.Println(result)
}

// 1. 筛选
func Filter[T any](slice []T, f func(T) bool) []T {
	var res []T
	for _, item := range slice {
		if f(item) {
			res = append(res, item)
		}
	}
	return res
}

// 2. 映射
func Map[T any, R any](slice []T, f func(T) R) []R {
	var res []R
	for _, item := range slice {
		res = append(res, f(item))
	}
	return res
}

// 3. 聚合
func Reduce[T any, R any](slice []T, init R, f func(R, T) R) R {
	result := init
	for _, item := range slice {
		result = f(result, item)
	}
	return result
}

// 4. 查找
func Find[T any](slice []T, f func(T) bool) (T, bool) {
	for _, item := range slice {
		if f(item) {
			return item, true
		}
	}
	var zero T
	return zero, false
}