package pipeline

import (
	"fmt"
	"sort"
)

func ArraySource(a ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, v := range a {
			fmt.Println("在Arraysort中")
			out <- v
		}
		close(out)
	}()
	return out
}

func InmemSort(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		a := []int{}
		for v := range in {
			a = append(a, v)
		}
		//排序
		sort.Ints(a)
		//输出
		for _, v := range a {
			fmt.Println("inmemsort")
			out <- v
		}
		close(out)
	}()
	return out
}

func MergeNodes(in1, in2 <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		v1, ok1 := <-in1
		v2, ok2 := <-in2
		for ok1 || ok2 {
			if !ok2 || (ok1 && v1 <= v2) {
				out <- v1
				v1, ok1 = <-in1
			} else {
				out <- v2
				v2, ok2 = <-in2
			}
		}
		close(out)
	}()
	return out
}
