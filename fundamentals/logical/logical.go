package main

import "fmt"

func Buy(job1, job2 bool) (bool, bool, bool, bool) {
	buy50tv := job1 && job2
	buy42tv := !buy50tv && (job1 || job2)
	buy32tv := job1 != job2
	buyIceCream := job1 || job2

	return buy50tv, buy42tv, buy32tv, buyIceCream
}

func main() {
	tv50, tv42, tv32, iceCream := Buy(true, true)
	fmt.Printf("tv32: %t, tv42: %t, tv50: %t, iceCream: %t", tv32, tv42, tv50, iceCream)
}
