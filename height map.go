package main

/*
Дан массив n неотрицательных целых чисел, представляющий карту высот.
Ширина каждой отметки равна 1. Вычислите, сколько уровней может быть заполнено водой после дождя.
//
//       #----#
//   #---##-#-#
//_#-##-#######
Input: height = [0,1,0,2,1,0,1,3,2,1,2,1,3] // 0 _ 0 _ _ 0 1 _ _ 1 _ 1
//
Ответ: 6
*/

func NonEmptyLevels(height []int) int {
	counter := make(map[int]int, 0)
	indMap := make(map[int]struct{}, 0)
	ans := 0
	for _, v := range height {
		if _, ok := indMap[v]; ok && v > 0 {
			for i := 0; i < v; i++ {
				ans += counter[i]
				counter[i] = 0
			}
		}
		counter[v]++
		indMap[v] = struct{}{}
	}
	return ans
}
