package main

import "fmt"

func fil_arr(arr *[10][10]rune) { // заполнение поля стандартными символами
	for y := 9; y >= 0; y-- {
		for x := 0; x <= 9; x++ {
			arr[x][y] = '.'
		}
	}
}

func route(arr *[10][10]rune) { // функция для задавания маршрута
	arr[5][0] = '&' //1
	arr[5][1] = '&' //2
	arr[5][2] = '&' //3
	arr[4][2] = '&' //4
	arr[3][2] = '&' //5
	arr[3][3] = '&' //6
	arr[4][3] = '&' //7
	arr[4][4] = '&' //8
	arr[5][4] = '&' //9
	arr[6][4] = '&' //10
	arr[7][4] = '&' //11
	arr[8][4] = '&' //12
	arr[8][5] = '&' //13
	arr[8][6] = '&' //14
	arr[7][6] = '&' //15
	arr[6][6] = '&' //16
	arr[5][6] = '&' //17
	arr[5][7] = '&' //18
	arr[5][8] = '&' //19
	arr[5][9] = '&' //20

}

/*
func route_check(arr *[10][10]rune) { // функция для задавания маршрута
	if arr[5][0] == '&' && arr[5][1] == '&' && arr[5][2] == '&' && arr[4][2] == '&' && arr[3][2] == '&' && arr[3][3] == '&' && arr[4][3] == '&' && arr[4][4] == '&' arr[5][4] == '&' && arr[6][4] == '&' && arr[7][4] == '&' && arr[8][4] == '&' && arr[8][5] == '&' && arr[8][6] == '&' && arr[7][6] == '&' //15 arr[6][6] == '&' //16
	arr[5][6] == '&' //17
	arr[5][7] == '&' //18
	arr[5][8] == '&' //19
	arr[5][9] == '&' //20


}
*/

func route_check(arr *[10][10]rune) { // функция для задавания маршрута
	var arr_check [20]rune
	arr_check[0] = arr[5][0]
	arr_check[1] = arr[5][1]
	arr_check[2] = arr[5][2]
	arr_check[3] = arr[4][2]
	arr_check[4] = arr[3][2]
	arr_check[5] = arr[3][3]
	arr_check[6] = arr[4][3]
	arr_check[7] = arr[4][4]
	arr_check[8] = arr[5][4]
	arr_check[9] = arr[6][4]
	arr_check[10] = arr[7][4]
	arr_check[11] = arr[8][4]
	arr_check[12] = arr[8][5]
	arr_check[13] = arr[8][6]
	arr_check[14] = arr[7][6]
	arr_check[15] = arr[6][6]
	arr_check[16] = arr[5][6]
	arr_check[17] = arr[5][7]
	arr_check[18] = arr[5][8]
	arr_check[19] = arr[5][9]
	kol := 0
	for i := 0; i < 20; i++ {
		if arr_check[i] == '&' {
			kol += 1
		}
	}
	if kol == 20 {
		fmt.Println("You won!")
	} else {
		fmt.Println("You lost..")
	}

}

func print_array(arr *[10][10]rune) {
	for i := 0; i <= 9; i++ {
		fmt.Printf("---")
	}
	fmt.Println()
	for y := 9; y >= 0; y-- {
		for x := 0; x <= 9; x++ {
			fmt.Printf(" %c ", arr[x][y])
		}
		fmt.Println()
	}
	for i := 0; i <= 9; i++ {
		fmt.Printf("---")
	}
	fmt.Println()
}

func switch_button(button rune, x *int, y *int) {

	switch button {
	case 'w':
		{
			*x = *x
			*y += 1
		}
	case 's':
		{
			*x = *x
			*y -= 1
		}
	case 'd':
		{
			*x += 1
			*y = *y
		}
	case 'a':
		{
			*x -= 1
			*y = *y
		}
	}
}

func main() {
	var arr [10][10]rune
	fil_arr(&arr)
	route(&arr)
	fmt.Printf("\n\nw - forward; a - left; d - right; s - back \n	Here is route\n")
	print_array(&arr)
	var ready int
	i := 0
	for i != 5 {
		fmt.Printf("Ready? (1 - yes, 0 - no)\n")
		fmt.Scanf("%d", &ready)
		if ready == 1 {
			i = 5
		} else {
			i = 0
		}
	}

	fil_arr(&arr)
	arr[5][0] = '&'
	print_array(&arr)

	end := false
	var button rune
	new_x := 5
	new_y := 0

	for end != true {
		fmt.Scanf("%c", &button)
		switch_button(button, &new_x, &new_y)
		arr[new_x][new_y] = '&'
		print_array(&arr)

		for y := 9; y >= 0; y-- {
			for x := 0; x <= 9; x++ {
				if arr[x][9] == '&' {
					end = true
				}
			}
		}

	}
	route_check(&arr)

}
