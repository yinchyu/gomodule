package gomodule

import (
	"fmt"
	"math"
	"os"
	"runtime"
	"runtime/pprof"
	"sync/atomic"
	"time"
)

type mod struct {
	name    string
	value   string
	age     int
	id      uint32
	context string
}

func putdata(queue *EsQueue) {
	for i := 0; i < 100; i++ {
		go func(i int) {
			newmod := mod{
				age: i,
			}
			for {
				queue.Put(newmod)
			}
		}(i)
	}
}

type a struct {
	name  int
	value int
}

func (h *a) String() string {
	return fmt.Sprintf("name:%v, value:%v", h.name, h.value)
}
func getdata(queue *EsQueue) {
	for i := 0; i < 105; i++ {
		go func(i int) {
			for {
				val, _, _ := queue.Get()
				fmt.Println("get goroutine num:", i, val)
			}
		}(i)
	}
}
func test() {
	queue := NewQueue(math.MaxUint16)
	//putdata(queue)
	//getdata(queue)
	//for{
	//}
	queue.Put("hello world")
	fmt.Println(queue.Get())
	var v uint32 = 21
	v--
	v |= v >> 1
	v |= v >> 2
	v |= v >> 4
	v |= v >> 8
	v |= v >> 16
	v++
	fmt.Println(v)
	var value uint32
	for i := 0; i > 100; i++ {
		go func() {
			atomic.AddUint32(&value, 1)
		}()
		if i%2 == 0 {
			go func() {
				atomic.AddUint32(&value, 2)
			}()
		}
	}
	for {
		fmt.Println(atomic.LoadUint32(&value))
	}
	pproF, _ := os.Create("pprof") // 创建记录文件
	pprof.StartCPUProfile(pproF)   // 开始cpu profile，结果写到文件f中
	defer pprof.StopCPUProfile()
	fmt.Println(runtime.NumCPU())
	start, _ := time.Parse(time.RFC3339, "1560-01-02T15:04:05Z07:00")
	fmt.Println(start)
	end := time.Now()
	// 最大能记录的范围再100年以内
	fmt.Println(end.Sub(start).Nanoseconds())
	fmt.Println(math.MaxInt64)
	ok := atomic.CompareAndSwapUint32(&v, 32, 34)
	if ok {
		fmt.Println(v)
	}
}

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)\n", p.Name, p.Age)
}

// func main() {

// 	queue := NewQueue(100)
// 	queue.Put("hello world")
// 	fmt.Println(queue.Get())

// }
