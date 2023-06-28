package main

func minimizeSet(divisor1 int, divisor2 int, uniqueCnt1 int, uniqueCnt2 int) int {
	n := 1

	//arr1 := make([]int, 0, uniqueCnt1)
	//arr2 := make([]int, 0, uniqueCnt2)
	//arr12 := make([]int, 0)
	k1 := 0
	k2 := 0
	k12 := 0

	for {
		if n%divisor1 != 0 {
			//arr1 = append(arr1, n)
			k1 += 1
		}

		if n%divisor2 != 0 {
			//arr2 = append(arr2, n)
			k2 += 1
		}

		if n%divisor1 != 0 && n%divisor2 != 0 {
			//arr12 = append(arr12, n)
			k12 += 1
		}

		//if len(arr1) >= uniqueCnt1 && len(arr2) >= uniqueCnt2 && len{
		//	break
		//}
		if k1 >= uniqueCnt1 && k2 >= uniqueCnt2 && k1+k2-k12 >= uniqueCnt1+uniqueCnt2 {
			return n
		}
		n += 1
	}
}
