package main

// import (
// 	"fmt"
// )

// func main() {
// 	boxes := ""
// 	width := 0
// 	n, _ := fmt.Scan(&boxes, &width)
// 	if n == 0 {
// 		fmt.Println("error")
// 	}

// 	grids := make([][]rune, width)
// 	for i := range grids {
// 		// 只有 1 行，放下所有箱子最大长度只有 1000
// 		grids[i] = make([]rune, 1000)
// 	}

// 	odd := 1
// 	colmn := 0
// 	index := 0
// 	rboxes := []rune(boxes)
// 	for index < len(rboxes) {
// 		if odd%2 == 1 {
// 			for j := 0; j < width && index < len(rboxes); j++ {
// 				grids[j][colmn] = rboxes[index]
// 				index++
// 			}
// 		} else {
// 			for j := width - 1; j >= 0 && index < len(rboxes); j-- {
// 				grids[j][colmn] = rboxes[index]
// 				index++
// 			}
// 		}
// 		odd++
// 		colmn++
// 	}
// 	for i := range grids {
// 		row := grids[i]
// 		for j := range row {
// 			if row[j] == 0 {
// 				break
// 			} else {
// 				fmt.Print(string(row[j]))
// 			}
// 		}
// 		if i < width-1 {
// 			fmt.Print("\n")
// 		}
// 	}
// }
