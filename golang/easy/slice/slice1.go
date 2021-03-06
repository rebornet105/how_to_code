/*
* @Title:   切片
* @Author:  pzqu
* @Date:    2020-04-05 23:54
* @url:     https://github.com/pzqu/how_to_code
*/
package main

import (
	"fmt"
	"how_to_code/golang/utils"
)

func sliceDemo1() {
	// 声明空切片，注意用var声明空切片，在使用前必须赋值。不然值为 nil
	var sliceTmp []int

	if sliceTmp == nil{
		fmt.Println("var []int is nil")
	}

	// 声明的同时赋值
	slice1 := []int{1, 2}
	fmt.Printf("输出切片：%v \n", slice1)

	for _,v := range slice1{
		fmt.Println(v)
	}

	// 用make声明，并指定长度
	sliceLen := make([]float64, 5)

	// 把数组切片
	arr := [5]int{1, 2, 3, 4, 5}
	sliceTmp = arr[3:] //从index=3截取到最后
	fmt.Printf("arr := [5]int{1,2,3,4,5} 切片 arr[3:] 得到 %v \n", sliceTmp)
	fmt.Printf("arr[:3] ： %v  \n", arr[:3])
	fmt.Printf("arr[2:3] : %v  \n", arr[2:3])

	// 多维数组
	slice2 := [][]int{{1, 2}, {1, 2, 3}, {1, 2, 3, 4}}
	fmt.Printf("输出多维切片: %v \n", slice2)

	// 忽略未使用切片
	utils.IgnoreUnused(sliceLen)
}

func main() {
	sliceDemo1()
}
