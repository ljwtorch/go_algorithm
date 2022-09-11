package main

import (
	"errors"
	"fmt"
	"strconv"
)

// Stack 使用数组来模拟一个栈的使用
type Stack struct {
	Max int      //栈最大存放个数
	Top int      //表示栈顶
	arr [100]int //数组模拟栈
}

// Push 入栈
func (s *Stack) Push(val int) (err error) {
	//先判断栈是否已满
	if s.Top == s.Max-1 {
		fmt.Println("stack full")
		return errors.New("stack full")
	}
	//栈顶自增后加入数据
	s.Top++
	s.arr[s.Top] = val
	return
}

// Pop 出栈
func (s *Stack) Pop() (val int, err error) {
	//判空
	if s.Top == -1 {
		fmt.Println("stack empty")
		return 0, errors.New("stack empty")
	}
	//先取值 栈顶指针再--
	val = s.arr[s.Top]
	s.Top--
	return val, nil
}

// List 栈的遍历
func (s *Stack) List() {
	//判空
	if s.Top == -1 {
		fmt.Println("stack empty")
		return
	}
	fmt.Println("栈中数据为：")
	for i := s.Top; i >= 0; i-- {
		fmt.Printf("arr[%d]=%d\n", i, s.arr[i])
	}
}

// IsOper 判断运算符
func (s *Stack) IsOper(val int) bool {
	if val == 42 || val == 43 || val == 45 || val == 47 {
		return true
	} else {
		return false
	}
}

// Cal 判断运算方法
func (s *Stack) Cal(num1 int, num2 int, oper int) int {
	res := 0
	switch oper {
	case 42:
		res = num2 * num1
	case 43:
		res = num2 + num1
	case 45:
		res = num2 - num1
	case 47:
		res = num2 / num1
	default:
		fmt.Println("运算符错误")
	}
	return res
}

// Priority 判断运算符优先级
func (s *Stack) Priority(oper int) int {
	res := 0
	if oper == 42 || oper == 47 {
		res = 1
	} else if oper == 43 || oper == 45 {
		res = 0
	}
	return res
}

func main() {
	numstack := Stack{
		Max: 20,
		Top: -1,
	}

	operstack := Stack{
		Max: 20,
		Top: -1,
	}

	exp := "30+30*6-4-6"
	//index辅助扫描exp
	index := 0
	//配合运算，定义需要的变量
	num1 := 0
	num2 := 0
	oper := 0
	result := 0
	keepNum := ""

	for {
		ch := exp[index : index+1]  //字符串
		temp := int([]byte(ch)[0])  //对应的ASCII值
		if operstack.IsOper(temp) { //说明是符号
			//如果operStack是一个空栈，直接入栈
			if operstack.Top == -1 { //空栈
				operstack.Push(temp)
			} else {
				/*
					如果发现opertstack栈顶的运算符的优先级大于等于当前准备入栈的运算符的优先级，
					就从符号栈pop出，并从数栈也pop两个数，进行运算，运算后的结果再重新入栈到数栈，
					符号再入符号栈
				*/
				if operstack.Priority(operstack.arr[operstack.Top]) >= operstack.Priority(temp) {
					num1, _ = numstack.Pop()
					num2, _ = numstack.Pop()
					oper, _ = operstack.Pop()
					result = operstack.Cal(num1, num2, oper) //注意弹出的顺序
					//将计算结果重新入栈、将运算符入栈
					numstack.Push(result)
					operstack.Push(temp)
				} else { //否则直接入栈
					operstack.Push(temp)
				}
			}
		} else { //说明是数
			//处理多位数的思路
			//1.定义一个keepNum string 用来拼接
			keepNum += ch
			//如果已经到最后，直接keepNum转换加进去
			if index == len(exp)-1 {
				val, _ := strconv.ParseInt(keepNum, 10, 64)
				numstack.Push(int(val))
			} else {
				//2.每次向index的后面字符测试，判断是不是运算符，然后处理
				if operstack.IsOper(int([]byte(exp[index+1 : index+2])[0])) {
					val, _ := strconv.ParseInt(keepNum, 10, 64)
					numstack.Push(int(val))
					keepNum = ""
				}
			}
		}
		//继续扫描，先判断index是否已经扫描到计算表达式的最后
		if index+1 == len(exp) {
			break
		}
		index++
	}
	/*
		如果扫描表达式完毕，依次从符号栈取出符号，然后从数栈取出两个数，
		运算后的结果，入数栈，直到符号栈为空
	*/
	for {
		if operstack.Top == -1 {
			break
		} else {
			num1, _ = numstack.Pop()
			num2, _ = numstack.Pop()
			oper, _ = operstack.Pop()
			result = operstack.Cal(num1, num2, oper) //注意弹出的顺序
			//将计算结果重新入栈、将运算符入栈
			numstack.Push(result)
		}
	}
	//如果运行正确，则结果就是numStack最后一个数
	res, _ := numstack.Pop()
	fmt.Printf("%s = %v", exp, res)
}
