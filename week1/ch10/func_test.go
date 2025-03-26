package ch10

import (
	"fmt"
	"testing"
)

func sum(a int ,b int) int{
	return a+b
}

func sum2(a ...int) int{
	if len(a) == 0{
		return 0
	}
	return sum2(a[1:]...) + a[0]
}

func sum3(a ...int) int{
	sum :=0
	for _,i := range a{
		sum += i
	}
	return sum
}

func TestSum(t *testing.T){
	fmt.Println(sum(1,2))
	fmt.Println(sum2(1,2,3,4,5))
	fmt.Println(sum3(1,2,3,4,5))
}

func TestDefer(t *testing.T){
	//函数返回前，类似于java的finally
	defer func(){
		fmt.Println("defer调用")
	}()
	fmt.Print("start")
}

func TestDefer2(t *testing.T) {
	var f = func() {
		defer fmt.Println("D")
		fmt.Println("F")
	}

	f()
	fmt.Println("M")
	// output: F D M
}
func TestDefer3(t *testing.T) {
	var f = func(i int) (r int) {
		defer func() {
			r += i
		}()

		/*
			流程：
			先将返回值result设为2
			执行defer语句，将result更新
			真正返回给调用方
		*/
		return 2
	}

	fmt.Println(f(10))
}
