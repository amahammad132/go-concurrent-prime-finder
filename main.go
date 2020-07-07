package main

import (
	"fmt"
	"sync"
	"time"
)

const numberOfThreads int = 12
var globalPrimes [numberOfThreads][]int

func isPrime(toCheck int) bool {
	//if toCheck == 2 {
	//	return true
	//}
	//if toCheck < 2 {
	//	return false
	//}

	if toCheck % 2 != 0 {
		for i := 3; i <= (toCheck / 3); i += 2 {
			if toCheck % i == 0 {
				return false
			}
		}
		return true
	} else {
		return false
	}

}

func primeSlice(start, end int) []int { // returns a slice of primes
	var slice []int
	/*
		distributor := func(numberOfThreads int) int {
			return (end - start) / numberOfThreads
		}
		fmt.Println("Interval: ", distributor(12)) */
	for num := start; num < end; num++ {
		if isPrime(num) {
			slice = append(slice, num)
		}
	}
	return slice
}

func goroutineExecutor(start, end int) [numberOfThreads + 1][]int {
	//var output []int
	//var twoDimensionalArray = make([][]int, end-start+1)
	//var twoDimensionalArray = make([][]int, 13)
	// var twoDimensionalArray [numberOfThreads + 1][]int
	var twoDimensionalSlice [numberOfThreads + 1][]int
	interval := (end - start) / numberOfThreads
	endVal := 0
	inputStart := 0
	inputEnd := 0
	for _, i := range globalPrimes {


		// code here
		// use inputStart and inputEnd in primeSlice
		// fmt.Println("start:", inputStart, "end:", inputEnd)

		// var process []int
		go func() {
			//output = append(output, primeSlice(inputStart, inputEnd)...)
			inputStart += endVal + start
			inputEnd += endVal + interval
			i = primeSlice(inputStart, inputEnd)
			//twoDimensionalSlice = append(twoDimensionalSlice, primeSlice(inputStart, inputEnd))
			//fmt.Println(primeSlice(inputStart, inputEnd))
			inputStart = inputEnd
		}()


	}
	fmt.Println(len(twoDimensionalSlice))
	return twoDimensionalSlice
}

func main() {
	start := time.Now()


	//var myWaitGroup sync.WaitGroup
	//myWaitGroup.Add(12)
	//var p1 []int
	//go func(waitgroup *sync.WaitGroup) {
	//	p1 = primeSlice(1, 125000)
	//	waitgroup.Done()
	//}(&myWaitGroup)
	//var p11 []int
	//go func(waitgroup *sync.WaitGroup) {
	//	p11 = primeSlice(125001, 250000)
	//	waitgroup.Done()
	//}(&myWaitGroup)
	//var p2 []int
	//go func(waitgroup *sync.WaitGroup) {
	//	p2 = primeSlice(250001, 375000)
	//	waitgroup.Done()
	//}(&myWaitGroup)
	//var p22 []int
	//go func(waitgroup *sync.WaitGroup) {
	//	p22 = primeSlice(375001, 500000)
	//	waitgroup.Done()
	//}(&myWaitGroup)
	//var p3 []int
	//go func(waitgroup *sync.WaitGroup) {
	//	p3 = primeSlice(500001, 562500)
	//	waitgroup.Done()
	//}(&myWaitGroup)
	//var p4 []int
	//go func(waitgroup *sync.WaitGroup) {
	//	p4 = primeSlice(562501, 625000)
	//	waitgroup.Done()
	//}(&myWaitGroup)
	//var p5 []int
	//go func(waitgroup *sync.WaitGroup) {
	//	p5 = primeSlice(625001, 687500)
	//	waitgroup.Done()
	//}(&myWaitGroup)
	//var p6 []int
	//go func(waitgroup *sync.WaitGroup) {
	//	p6 = primeSlice(687501, 750000)
	//	waitgroup.Done()
	//}(&myWaitGroup)
	//var p7 []int
	//go func(waitgroup *sync.WaitGroup) {
	//	p7 = primeSlice(750001, 812500)
	//	waitgroup.Done()
	//}(&myWaitGroup)
	//var p8 []int
	//go func(waitgroup *sync.WaitGroup) {
	//	p8 = primeSlice(812501, 875000)
	//	waitgroup.Done()
	//}(&myWaitGroup)
	//var p9 []int
	//go func(waitgroup *sync.WaitGroup) {
	//	p9 = primeSlice(875001, 937500)
	//	waitgroup.Done()
	//}(&myWaitGroup)
	//var p10 []int;
	//go func(waitgroup *sync.WaitGroup) {
	//	p10 = primeSlice(937501, 1000000)
	//	waitgroup.Done()
	//}(&myWaitGroup)
	//myWaitGroup.Wait()
	//fmt.Println(p1, p11, p2, p22, p3, p4, p5, p6, p7, p8, p9, p10)



	//var myWaitGroup sync.WaitGroup
	//myWaitGroup.Add(12)

	// option 1

	//var p0 []int; go func(waitgroup *sync.WaitGroup){p0 = primeSlice(1, 83333); waitgroup.Done()}(&myWaitGroup)
	//var p1 []int; go func(waitgroup *sync.WaitGroup){p1 = primeSlice(83334, 166666); waitgroup.Done()}(&myWaitGroup)
	//var p2 []int; go func(waitgroup *sync.WaitGroup){p2 = primeSlice(166667, 249999); waitgroup.Done()}(&myWaitGroup)
	//var p3 []int; go func(waitgroup *sync.WaitGroup){p3 = primeSlice(250000, 333332); waitgroup.Done()}(&myWaitGroup)
	//var p4 []int; go func(waitgroup *sync.WaitGroup){p4 = primeSlice(333333, 416665); waitgroup.Done()}(&myWaitGroup)
	//var p5 []int; go func(waitgroup *sync.WaitGroup){p5 = primeSlice(416666, 499998); waitgroup.Done()}(&myWaitGroup)
	//var p6 []int; go func(waitgroup *sync.WaitGroup){p6 = primeSlice(499999, 583331); waitgroup.Done()}(&myWaitGroup)
	//var p7 []int; go func(waitgroup *sync.WaitGroup){p7 = primeSlice(583332, 666664); waitgroup.Done()}(&myWaitGroup)
	//var p8 []int; go func(waitgroup *sync.WaitGroup){p8 = primeSlice(666665, 749997); waitgroup.Done()}(&myWaitGroup)
	//var p9 []int; go func(waitgroup *sync.WaitGroup){p9 = primeSlice(749998, 833330); waitgroup.Done()}(&myWaitGroup)
	//var p10 []int; go func(waitgroup *sync.WaitGroup){p10 = primeSlice(833331, 916663); waitgroup.Done()}(&myWaitGroup)
	//var p11 []int; go func(waitgroup *sync.WaitGroup){p11 = primeSlice(916664, 999996); waitgroup.Done()}(&myWaitGroup)

	// option 2
/*
	var p0 []int; go func(){p0 = primeSlice(1, 83333); defer myWaitGroup.Done()}()
	var p1 []int; go func(){p1 = primeSlice(83334, 166666); defer myWaitGroup.Done()}()
	var p2 []int; go func(){p2 = primeSlice(166667, 249999); defer myWaitGroup.Done()}()
	var p3 []int; go func(){p3 = primeSlice(250000, 333332); defer myWaitGroup.Done()}()
	var p4 []int; go func(){p4 = primeSlice(333333, 416665); defer myWaitGroup.Done()}()
	var p5 []int; go func(){p5 = primeSlice(416666, 499998); defer myWaitGroup.Done()}()
	var p6 []int; go func(){p6 = primeSlice(499999, 583331); defer myWaitGroup.Done()}()
	var p7 []int; go func(){p7 = primeSlice(583332, 666664); defer myWaitGroup.Done()}()
	var p8 []int; go func(){p8 = primeSlice(666665, 749997); defer myWaitGroup.Done()}()
	var p9 []int; go func(){p9 = primeSlice(749998, 833330); defer myWaitGroup.Done()}()
	var p10 []int; go func(){p10 = primeSlice(833331, 916663); defer myWaitGroup.Done()}()
	var p11 []int; go func(){p11 = primeSlice(916664, 999996); defer myWaitGroup.Done()}()
*/
	// myWaitGroup.Wait()
	starting := 1
	end := 1000000
	interval := (end - starting) / numberOfThreads
	endVal := 0
	inputStart := 0
	inputEnd := 0
	var outslice [12][3]int
	var newWaitGroup sync.WaitGroup
	newWaitGroup.Add(numberOfThreads)
	for i := 0; i < numberOfThreads; i++ {
		inputStart += endVal + starting
		inputEnd += endVal + interval

		// code here
		// use inputStart and inputEnd in primeSlice
		if i == numberOfThreads - 1 {inputEnd = end}
		fmt.Println("start:", inputStart, "end:", inputEnd)
		fmt.Println(i)
		go func() {
			outslice[i] = [3]int {1, 2, 3}
			defer newWaitGroup.Done()
		}()
		inputStart = inputEnd
	}
	// newWaitGroup.Wait()
	fmt.Println(outslice)



//	fmt.Println(p0, p1, p2, p3, p4, p5, p6, p7, p8, p9, p10, p11)







	//time.Sleep(4000 * time.Millisecond)

	// test array of slices
	//var primeslices [threadcount][]int
	//go func() {
	//	primeslices = [threadcount][]int {
	//		primeSlice(1, 125000),
	//		primeSlice(125001, 250000),
	//		primeSlice(250001, 375000),
	//		primeSlice(375001, 500000),
	//		primeSlice(500001, 562500),
	//		primeSlice(562501, 625000),
	//		primeSlice(625001, 687500),
	//		primeSlice(687501, 750000),
	//		primeSlice(750001, 812500),
	//		primeSlice(812501, 875000),
	//		primeSlice(875001, 937500),
	//		primeSlice(937501, 1000000),
	//	}
	//}()
	//time.Sleep(12000 * time.Millisecond)
	//fmt.Println(primeslices)



	//fullprimes := append(append(append(append(append(append(append(append(append(append(append(p1, p11...), p2...), p22...), p3...), p4...), p5...), p6...), p7...), p8...), p9...), p10...)
	//fmt.Println(fullprimes)
	//for _, item := range fullprimes {
	//	fmt.Println(item)
	//}


	fmt.Printf("\n\n\n\n\n\n")



	//old code above




	//new code
	/*
		//fmt.Println(goroutineExecutor(1, 1000000, 12))
		ae := goroutineExecutor(1, 1000000)
		time.Sleep(12000 * time.Millisecond)
		fmt.Println(ae)
	*/
	// goroutineExecutor(1, 1000000)
	// time.Sleep(14 * time.Second)
	// fmt.Println(globalPrimes)
	fmt.Println("Time taken: ", time.Since(start))
}