package main

/*
Дан массив n неотрицательных целых чисел, представляющий карту высот.
Ширина каждой отметки равна 1. Вычислите, сколько уровней может быть заполнено водой после дождя.
//
//       #
//   #---##-#
//_#-##-######
Input: height = [0,1,0,2,1,0,1,3,2,1,2,1] //0 1 _ 2 _ _ _ 3 2 _ 2 _
//
Ответ: 6
*/

func NonEmptyLevels(height []int) int {
	stack := []int{height[0]}
	ans := 0
	for _, v := range height {
		for stack[len(stack)-1] > v {
			stack = stack[:len(stack)-1]
			ans += 1
		}
		stack = append(stack, v)
	}

	return ans
}
