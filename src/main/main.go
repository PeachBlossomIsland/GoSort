package main

import (
	"fmt"
	"pipeline"
)

func main() {
	out := pipeline.MergeNodes(
		pipeline.InmemSort(pipeline.ArraySource(2, 5, 3, 4)),
		pipeline.InmemSort(pipeline.ArraySource(0, 4, 7, 6)))
	for {
		for i := range out {
			fmt.Println("在main中")
			fmt.Println(i)
		}
	}
}
