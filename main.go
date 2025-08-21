package main

import (
	"bytes"
	"fmt"
	"math"
	"net/http"
	"os"
	"sort"
	"strconv"
	"sync"
	"sync/atomic"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/zxNewBee/init_project/dataLib"

	"github.com/urfave/cli/v2"

	"golang.org/x/sync/singleflight"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var gormDb *gorm.DB
var sqlxDb *sqlx.DB

func main() {
	fmt.Println("Hello, World!")
	gormDb = connectSql()
	//sqlxDb = connectSqlx()
	//////////////task1//////////////
	//taskOne()
	//////////////task2///////////////
	//taskTwo()
	//////////////task3///////////////
	taskThree()
	/////////////////测试///////////////
	//test()

}
func taskOne() {
	//查找只出现一次的元素
	elements := []int{3, 2, 2, 5, 3, 3, 12, 5}
	outValue := findElement(elements)
	fmt.Println(outValue)
	//有效的括
	success := isStrValid("{()}[]{()}")
	fmt.Println(success)
	//
	arr := []uint{9, 9, 9}
	outArr := plusOne(arr)
	fmt.Println(outArr)
	//删除数组重复元素
	arrEles := []int{1, 1, 2, 3, 3, 4, 5, 5, 5, 6}
	arrNew := deleteRepeatElement(arrEles)
	fmt.Println(arrNew)
	//合并区间
	intervals := [][]int{{2, 3}, {3, 1}, {4, 5}, {4, 10}, {8, 9}}
	interfavalsNew := deleteAndMergeInbtervals(intervals)
	fmt.Println(interfavalsNew)
	//两数之和
	numArr := []int{1, 2, 3, 4, 5, 6, 8}
	//index := findAddNumIndex(numArr, 7)
	index := findAddNumIndexByMap(numArr, 7)
	fmt.Println(index)
}
func taskTwo() {
	//编写一个Go程序，定义一个函数，该函数接收一个整数指针作为参数，
	//在函数内部将该指针指向的值增0，然后在主函数中调用该函数并输出修改后的值
	intPtr := 10
	modifyNum := modifyPtrValue(&intPtr)
	fmt.Println(modifyNum)
	//实现一个函数，接收一个整数切片的指针，将切片中的每个元素乘以2
	values := []int{10, 20, 30}
	ptrSlice := make([]*int, len(values))

	for i := range ptrSlice {
		ptrSlice[i] = &values[i]
	}
	modifyNums := multfy2ByPtr(ptrSlice)
	fmt.Println(modifyNums)
	//编写一个程序，使用 go 关键字启动两个协程，一个协程打印从10的奇数，
	// 另一个协程打印从20的偶数
	go printEvenNum()
	go printOddNum()
	time.Sleep(3 * time.Second)

	//定义一Shape 接口，包Area() Perimeter() 两个方法
	// 然后创建 Rectangle Circle 结构体，实现 Shape 接口
	// 在主函数中，创建这两个结构体的实例，并调用它们的 Area() Perimeter() 方法
	cicle := Circle{Radius: 10}
	fmt.Printf("Circle area is:%f\n", cicle.Area())
	fmt.Printf("Circle perimeter is:%f\n", cicle.Perimeter())
	rectangle := Rectangle{Width: 10, Height: 30}
	fmt.Printf("Rectangle area is:%f\n", rectangle.Area())
	fmt.Printf("Rectangle perimeter is:%f\n", rectangle.Perimeter())

	//使用组合的方式创建一Person 结构体，包含 Name Age 字段，再创建一Employee 结构体，
	// 组合 Person 结构体并添加 EmployeeID 字段。为 Employee 结构体实现一PrintInfo() 方法，输出员工的信息
	employee := Employee{
		Person:     Person{Name: "abc", Age: 18},
		EmployeeId: "0001",
	}
	employee.PrintInfo()
	ch := make(chan int)
	go getNumber(ch)
	go printNumber(ch)
	// 编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。
	// 启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
	saveCounter()
	saveCounterAtomic()
}
func taskThree() {
	//db := connectSql()

	//dataLib.CreateStudentInfo(db)
	//dataLib.QueryStudentsOver18(db)

	//queryBySqlx()
	//creteBookTable()
	createBlogTable()
	// // 假设有两个表： accounts 表（包含字段 id 主键， balance 账户余额）和 transactions 表
	// // （包含字段 id 主键， from_account_id 转出账户ID， to_account_id 转入账户ID， amount 转账金额）
	// // 要求 ：
	// // 编写一个事务，实现从账户 A 向账户 B 转账 100 元的操作。在事务中，需要先检查账户 A 的余额是否足够，
	// // 如果足够则从账户 A 扣除 100 元，向账户 B 增加 100 元，
	// // 并在 transactions 表中记录该笔转账信息。如果余额不足，则回滚事务

	// // 账户/交易表结构迁移与示例转账
	// if err := dataLib.AutoMigrateAccountAndTransactions(db); err != nil {
	// 	panic(err)
	// }

	// // 确保两个账户存在（示例）
	// db.FirstOrCreate(&dataLib.Account{}, dataLib.Account{ID: 1})
	// db.FirstOrCreate(&dataLib.Account{}, dataLib.Account{ID: 2})
	// // 为演示方便，给账户1充点
	// db.Model(&dataLib.Account{}).Where("id = ?", 1).Update("balance", gorm.Expr("GREATEST(balance, ?)", 200))

	// if err := dataLib.TransferFunds(db, 1, 2, 100); err != nil {
	// 	fmt.Println("转账失败:", err)
	// } else {
	// 	fmt.Println("转账成功: 1 -> 2 金额 100")
	// }
	// //gorm.G[dataLib.Account](db).Create(context.Background(), &dataLib.Account{ID: 3, Balance: 999})

}
func test() {

	//helloGin()

	//calculateCap()
	// newSlice()

	// var wg sync.WaitGroup

	// // 启动5个worker
	// for i := 1; i <= 5; i++ {
	// 	wg.Add(1) // 增加计数
	// 	go workerForWaitGroup(i, &wg)
	// }

	// // 等待所有worker完成
	// wg.Wait()
	// fmt.Println("All workers completed")

	// instance := GetInstance()
	// instance.DoSomething("johnson", 11)
	// instance2 := GetInstance()
	// instance2.DoSomething("johnson2", 12)
	//poolTest()
	//singleflightTest()
	//TestSingleFlight()
	//testCli()
}

// 给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素
// 可以使用 for 循环遍历数组，结if 条件判断map 数据结构来解决，例如通过 map 记录每个元素出现的次数，然后再遍map 找到出现次数的元素
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

//给定一个只包括 '(')'{'}'[']' 的字符串，判断字符串是否有效
//有效字符串需满足
//左括号必须用相同类型的右括号闭合
//左括号必须以正确的顺序闭合
//每个右括号都有一个对应的相同类型的左括号

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

// 给定一个表大整的整数数digits，其digits[i] 是整数的i 位数字。这些数字按从左到右，从最高位到最低位排列。这个大整数不包含任何前0
// 将大整数1，并返回结果的数字数组
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
		// 初始化代..
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

	// 清空并放回池
	buffer1.Reset()
	bufferPool.Put(buffer1)

	// 获取一个Buffer（可能是刚才放回的那个）
	buffer2 := bufferPool.Get().(*bytes.Buffer)
	buffer2.WriteString("World")
	fmt.Println("Buffer2:", buffer2.String())

	// 清空并放回池
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

			// 清空并放
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
	fmt.Printf("总请求数=%d,请求成功%d,请求失败%d", n, n-failCnt, failCnt)
}

func get(index int, key string) (interface{}, error) {
	var err error
	if atomic.AddInt32(&offset, 1) == 3 { // 假设偏移offset == 3 执行耗时长，超时失败
		time.Sleep(time.Microsecond * 500)
		err = fmt.Errorf("")
		fmt.Printf("faild index %d\n", index)
	}
	fmt.Printf("success index %d\n", index)
	return key, err
}

func testCli() {
	app := &cli.App{
		Name:  "calc",
		Usage: "命令行计算器",
		Commands: []*cli.Command{
			{
				Name:    "add",
				Aliases: []string{"a"},
				Usage:   "加法: calc add <num1> <num2>",
				Action: func(c *cli.Context) error {
					return calculate(c, "+")
				},
			},
			{
				Name:    "sub",
				Aliases: []string{"s"},
				Usage:   "减法: calc sub <num1> <num2>",
				Action: func(c *cli.Context) error {
					return calculate(c, "-")
				},
			},
			{
				Name:    "mul",
				Aliases: []string{"m"},
				Usage:   "乘法: calc mul <num1> <num2>",
				Action: func(c *cli.Context) error {
					return calculate(c, "*")
				},
			},
			{
				Name:    "div",
				Aliases: []string{"d"},
				Usage:   "除法: calc div <num1> <num2>",
				Action: func(c *cli.Context) error {
					return calculate(c, "/")
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "错误: %v\n", err)
		os.Exit(1)
	}
}

// 公共计算逻辑
func calculate(c *cli.Context, op string) error {
	if c.NArg() != 2 {
		return fmt.Errorf("需要两个数字参数\n示例: calc %s 10 5", c.Command.Name)
	}

	// 解析参数
	a, err := strconv.ParseFloat(c.Args().Get(0), 64)
	if err != nil {
		return fmt.Errorf("无效数字: %s", c.Args().Get(0))
	}

	b, err := strconv.ParseFloat(c.Args().Get(1), 64)
	if err != nil {
		return fmt.Errorf("无效数字: %s", c.Args().Get(1))
	}

	// 执行计算
	var result float64
	switch op {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		if b == 0 {
			return fmt.Errorf("除数不能为零")
		}
		result = a / b
	default:
		return fmt.Errorf("")
	}

	// 输出结果
	fmt.Printf("结果: %.2f %s %.2f = %.2f\n", a, op, b, result)
	return nil
}

// 复杂度保证o(1) 可使用快慢双指针
func deleteRepeatElement(arr []int) []int {
	slowIndex := 0
	for fastIndex := 1; fastIndex < len(arr); fastIndex++ {
		if arr[slowIndex] != arr[fastIndex] {
			slowIndex++
			arr[slowIndex] = arr[fastIndex]
		} else {

		}
	}
	return arr[:slowIndex+1]
}

func deleteAndMergeInbtervals(intervals [][]int) [][]int {
	//对二维排
	for index := 0; index < len(intervals); index++ {
		sort.Slice(intervals[index], func(i, j int) bool {
			return intervals[index][i] < intervals[index][j]
		})
	}
	//对一维排
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	fmt.Println(intervals)
	outIntervals := [][]int{intervals[0]}
	for i := 0; i < len(intervals)-1; i++ {
		lastMerge := outIntervals[len(outIntervals)-1]
		if lastMerge[1] > intervals[i+1][0] {
			arr := []int{lastMerge[0], max(lastMerge[1], intervals[i+1][1])}
			outIntervals[len(outIntervals)-1] = arr
		} else {
			outIntervals = append(outIntervals, intervals[i+1])
		}
	}
	return outIntervals
}

// 返回下标
func findAddNumIndex(arr []int, target int) []int {
	//addResult := make(map[int]int, 0)
	outIndex := []int{}

	for i := 0; i < len(arr)-2; i++ {
		for j := 1; j < len(arr)-1; j++ {
			result := arr[i] + arr[j]
			if result == target {

				outIndex = append(outIndex, i)
				outIndex = append(outIndex, j)
				return outIndex
			}
		}

	}
	return outIndex
}

// 返回下标
func findAddNumIndexByMap(arr []int, target int) []int {
	subResult := make(map[int]int, 0)
	for i, num := range arr {
		subNum := target - num
		if index, found := subResult[subNum]; found {
			return []int{index, i}
		}
		subResult[num] = subNum
	}
	return nil
}

func modifyPtrValue(ptr *int) int {
	*ptr += 10
	return *ptr
}

func multfy2ByPtr(arr []*int) []int {
	outValue := make([]int, 0)
	for _, num := range arr {
		outNum := (*num) * 2
		outValue = append(outValue, outNum)
	}
	return outValue
}

func printOddNum() {
	for i := 1; i < 10; i++ {
		if i%2 == 1 {
			fmt.Println(i)

		}
	}
}
func printEvenNum() {
	for i := 1; i < 10; i++ {
		if i%2 == 0 {
			fmt.Println(i)

		}
	}
}

type Shape interface {
	Area()
	Perimeter()
}
type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	fmt.Println("Rectangle 实现 Area方法")
	return r.Height * r.Width
}

func (r Rectangle) Perimeter() float64 {
	fmt.Println("Rectangle 实现 perimeter")
	return 2 * (r.Height + r.Width)
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	fmt.Println("Circle 实现 Area方法")
	return c.Radius * c.Radius * math.Pi
}
func (c Circle) Perimeter() float64 {
	fmt.Println("Circle 实现 Area方法")
	return 2 * c.Radius * math.Pi
}

type Person struct {
	Name string
	Age  int
}

type Employee struct {
	Person
	EmployeeId string
}

func (e Employee) PrintInfo() {
	fmt.Println(e.EmployeeId, e.Name, e.Age)
}
func getNumber(ch chan<- int) {
	for i := 1; i < 10; i++ {
		ch <- i
	}
	close(ch)
}

func printNumber(ch <-chan int) {
	for num := range ch {
		fmt.Println(num)
	}
}
func saveCounter() {
	var mu sync.Mutex
	var wg sync.WaitGroup
	count := 0
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				mu.Lock()
				count++
				mu.Unlock()
			}

		}()
	}
	wg.Wait()
	fmt.Println("count:", count)
}
func saveCounterAtomic() {
	var wg sync.WaitGroup
	var count int64
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				atomic.AddInt64(&count, 1)
			}

		}()
	}
	wg.Wait()
	fmt.Println("atomic count:", count)
}

func connectSql() *gorm.DB {
	db, err := gorm.Open(mysql.Open("root:123@tcp(127.0.0.1:3306)/new_schema?charset=utf8mb4&parseTime=True&loc=Local"))
	if err != nil {
		panic(err)
	}
	defer db.DB()
	return db

}

//	func connectSql() {
//		db, err := gorm.Open(mysql.Open("root:123@tcp(127.0.0.1:3306)/new_schema?charset=utf8mb4&parseTime=True&loc=Local"))
//		if err != nil {
//			panic(err)
//		}
//		defer db.DB()
//		dataLib.Run(db)
//	}
func helloGin() {
	r := gin.Default()
	r.GET("/user", func(c *gin.Context) {
		c.String(200, "hello world")
		c.Redirect(http.StatusMovedPermanently, "www.baidu.com")

	})
	err := r.Run(":8080")
	if err != nil {
		panic(err)
	}
}
func connectSqlx() *sqlx.DB {
	db, err := sqlx.Connect("mysql", "root:123@tcp(127.0.0.1:3306)/new_schema?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	return db
}
func queryBySqlx() {
	db, err := sqlx.Connect("mysql", "root:123@tcp(127.0.0.1:3306)/new_schema?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	row, err := db.Queryx("SELECT id FROM accounts ORDER BY balance DESC LIMIT ?", 1)
	if err != nil {
		fmt.Println("查询失败:", err)
		panic(err)
	}
	employeeMap := make(map[string]interface{})
	if row.Next() {
		err = row.MapScan(employeeMap)
		if err != nil {
			panic(err)
		}
		fmt.Println("id is:", employeeMap["id"])
	} else {
		fmt.Println("no data")
	}

}
func creteBookTable() {

	book := dataLib.CreateBookTable(sqlxDb)

	fmt.Println("book table created:", book.BookId, book.Title, book.BookName)

}

func createBlogTable() {
	dataLib.CreateBlogTable(gormDb)
}
