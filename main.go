package main

import (
	"bytes"
	"fmt"
	"sync"
	"time"

	"sync/atomic"

	"golang.org/x/sync/singleflight"
)

func main() {
	fmt.Println("Hello, World!")
	// elements := []int{3, 2, 2, 5, 3, 3, 12, 5}
	// outValue := findElement(elements)
	// fmt.Println(outValue)

	success := isStrValid("{()}[]{()}")
	fmt.Println(success)
	arr := []uint{9, 9, 9}
	outArr := plusOne(arr)
	fmt.Println(outArr)

	calculateCap()
	newSlice()

	var wg sync.WaitGroup

	// 启动5个worker
	for i := 1; i <= 5; i++ {
		wg.Add(1) // 增加计数器
		go workerForWaitGroup(i, &wg)
	}

	// 等待所有worker完成
	wg.Wait()
	fmt.Println("All workers completed")

	instance := GetInstance()
	instance.DoSomething("johnson", 11)
	instance2 := GetInstance()
	instance2.DoSomething("johnson2", 12)
	poolTest()
	singleflightTest()
	TestSingleFlight()
}

// 给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。
// 可以使用 for 循环遍历数组，结合 if 条件判断和 map 数据结构来解决，例如通过 map 记录每个元素出现的次数，然后再遍历 map 找到出现次数为1的元素。
func findElement(arr []int) int {

	var countMap = make(map[int]int)
	//fmt.Println(len(arr))
	for i := 0; i < len(arr); i++ {
		value, ok := countMap[arr[i]]
		if ok {
			countMap[arr[i]] = value + 1
			//fmt.Println(countMap[arr[i]])
		} else {
			countMap[arr[i]] = 1
		}
	}

	for key := range countMap {
		if countMap[key] == 1 {
			return key
		}
	}
	return -1
}

//给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效
//有效字符串需满足：
//左括号必须用相同类型的右括号闭合。
//左括号必须以正确的顺序闭合。
//每个右括号都有一个对应的相同类型的左括号。

func isStrValid(str string) bool {

	strStack := []string{}

	// pairs := map[string]string{
	// 	"(": ")",
	// 	"{": "}",
	// 	"[": "]",
	// }

	for index, value := range str {
		switch string(value) {
		case "(":
			strStack = append(strStack, "(")
		case "{":
			strStack = append(strStack, "{")
		case "[":
			strStack = append(strStack, "[")

		case ")":
			if len(strStack) > 0 && strStack[len(strStack)-1] == "(" {
				strStack = strStack[:len(strStack)-1]
			} else {
				return false
			}
		case "}":
			if len(strStack) > 0 && strStack[len(strStack)-1] == "{" {
				strStack = strStack[:len(strStack)-1]
			} else {
				return false
			}
		case "]":
			if len(strStack) > 0 && strStack[len(strStack)-1] == "[" {
				strStack = strStack[:len(strStack)-1]
			} else {
				return false
			}
		}

		fmt.Println(index, string(value))
	}

	return len(strStack) == 0
}

// 给定一个表示 大整数 的整数数组 digits，其中 digits[i] 是整数的第 i 位数字。这些数字按从左到右，从最高位到最低位排列。这个大整数不包含任何前导 0。
// 将大整数加 1，并返回结果的数字数组。
func plusOne(arr []uint) []uint {

	for i := len(arr) - 1; i > -1; i-- {

		result := arr[i] + 1
		if result > 9 {
			arr[i] = 0
		} else {
			arr[i] += 1
			return arr
		}
	}
	//arr = make([]uint, len(arr))
	arrHead := []uint{1}
	arrHead = append(arrHead, arr...)
	return arrHead
}

func calculateCap() {
	nums := []int{1, 2}
	nums = append(nums, 2, 3, 4)
	fmt.Printf("len:%d cap:%d", len(nums), cap(nums))
}
func newSlice() []int {
	arr := [3]int{1, 2, 3}
	slice := arr[0:1]
	return slice
}

func workerForWaitGroup(id int, wg *sync.WaitGroup) {
	defer wg.Done() // 工作完成时通知WaitGroup

	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Second) // 模拟工作
	fmt.Printf("Worker %d done\n", id)
}

type Singleton struct {
	// 单例结构字段
	name string
	age  int
}

func (s *Singleton) DoSomething(name string, age int) {
	fmt.Println("Singleton doing something")
	s.name = name
	s.age = age
	fmt.Println(s.name, s.age)
}

var (
	instance *Singleton
	once     sync.Once
)

func GetInstance() *Singleton {
	once.Do(func() {
		instance = &Singleton{}
		// 初始化代码...
	})
	return instance
}

func poolTest() {
	// 创建一个池，用于复用bytes.Buffer
	var bufferPool = sync.Pool{
		New: func() interface{} {
			fmt.Println("Creating a new buffer")
			return new(bytes.Buffer)
		},
	}

	// 获取一个Buffer
	buffer1 := bufferPool.Get().(*bytes.Buffer)
	buffer1.WriteString("Hello")
	fmt.Println("Buffer1:", buffer1.String())

	// 清空并放回池中
	buffer1.Reset()
	bufferPool.Put(buffer1)

	// 获取一个Buffer（可能是刚才放回的那个）
	buffer2 := bufferPool.Get().(*bytes.Buffer)
	buffer2.WriteString("World")
	fmt.Println("Buffer2:", buffer2.String())

	// 清空并放回池中
	buffer2.Reset()
	bufferPool.Put(buffer2)

	// 同时获取多个Buffer
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			// 获取Buffer
			buf := bufferPool.Get().(*bytes.Buffer)

			// 使用Buffer
			buf.WriteString(fmt.Sprintf("Goroutine %d", id))
			fmt.Printf("Goroutine %d: %s\n", id, buf.String())

			// 清空并放回
			buf.Reset()
			bufferPool.Put(buf)
		}(i)
	}

	wg.Wait()
}

func singleflightTest() {
	g := new(singleflight.Group)
	go func() {
		v1, _, shared := g.Do("key", func() (interface{}, error) {
			time.Sleep(time.Second * 3)
			return "msg", nil
		})
		fmt.Printf("first call v1:%v,shared:%v\n", v1, shared)
	}()
	time.Sleep(time.Second * 1)

	v2, _, shared2 := g.Do("key", func() (interface{}, error) {

		return "msg2", nil
	})
	fmt.Printf("second call v2:%v,shared:%v\n", v2, shared2)

}

var (
	offset int32 = 0
)

func TestSingleFlight() {
	var (
		n       int32 = 100
		k             = "12344556"
		wg            = sync.WaitGroup{}
		sf      singleflight.Group
		failCnt int32 = 0
	)

	for i := 0; i < int(n); i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			_, err, _ := sf.Do(k, func() (interface{}, error) {
				return get(i, k)
			})
			if err != nil {
				failCnt++
				//atomic.AddInt32(&failCnt, 1)
				fmt.Printf("count %d,failcnt %d", i, failCnt)
				return
			}
		}()
	}

	wg.Wait()
	fmt.Printf("总请求数=%d,请求成功率=%d,请求失败率=%d", n, n-failCnt, failCnt)
}

func get(index int, key string) (interface{}, error) {
	var err error
	if atomic.AddInt32(&offset, 1) == 3 { // 假设偏移量 offset == 3 执行耗时长，超时失败了
		time.Sleep(time.Microsecond * 500)
		err = fmt.Errorf("耗时长")
		fmt.Printf("faild index %d\n", index)
	}
	fmt.Printf("success index %d\n", index)
	return key, err
}
