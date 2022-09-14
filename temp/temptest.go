package main

import "fmt"

func rangeCheck() {
	i := 0

	// 印证了三元表达式判断到第一次不符合条件时i++就不会执行了，实际的执行顺序如下j的示例相同
	for ; i <= 10; i++ {
		fmt.Println(i)
	}

	j := 0
	for {
		if !(j <= 10) {
			break
		}
		fmt.Println(j)
		j++
	}
	fmt.Println("l start")
	l := 0
	for l++; l <= 20; l *= 2 {
		fmt.Println(l)
	}

	fmt.Println("i is:", i)
	fmt.Println("j is:", j)

}

func ifELseCkeck() {
	// if 判断分支中的短变量声明只在if以及之后的else elseif中生效
	if t := 10; t > 9 {
		fmt.Println(t)
	}
	fmt.Println("x")
}

func main() {
	rangeCheck()
	// for循环中有两个作用域，一个是for之后的隐式作用域另一个是由{}括起来的显示作用域，i := 0 和 i++ 都属于隐式作用域中，{}是隐式作用域中的作用域
	// 所以i<3是在{}每次执行之前执行，i++是在{}每次执行之后执行，在{}中对i进行操作会影响到隐式作用域的i
	// 但是如果在内部新定义一个i（i:=i）则对新定义的i做操作不会影响到外部的i的值，且只在此次循环中生效。
	/*所以一个for循环的真实域可以理解如下，break直接作用到外层作用。continue只在最内层{}中作用，所以外层的i++在continue之后依然会执行
	for {
		i := 0 // 只执行一次
		i < 3{
			statement
		}i++
	}

	*/
	for i := 0; i < 3; i++ {
		fmt.Println(i)
		i := i // 这里声明的变量i遮挡了上面声明的i。
		// 右边的i为上面声明的循环变量i。
		i = 10 // 新声明的i被更改了。
		_ = i
	}
}
