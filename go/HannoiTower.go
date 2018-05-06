// created by Cans, 20180506

package main

import (
	"fmt"
)

func HannoiTower(n int, from, helper, to byte) {
	if n == 1 {
		fmt.Printf("%c ---> %c\n", from, to)
	} else {
		// 待搬动盘子上面的盘子移到辅助柱子上
		HannoiTower(n-1, from, to, helper)
		fmt.Printf("%c ---> %c\n", from, to)
		// from 与 helper 角色互换
		from, helper = helper, from
		// 递归处理辅助柱子上的盘子
		HannoiTower(n-1, from, helper, to)
	}
}

func main() {
	level := 5
	HannoiTower(level, 'A', 'B', 'C')
}
