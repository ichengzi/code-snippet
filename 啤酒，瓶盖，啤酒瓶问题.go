/**************************************************
*　啤酒：2元钱1瓶
*　　　　4个瓶盖换1瓶啤酒
*　　　　2个空瓶换1瓶啤酒
*　　问：10元可以喝几瓶？
**************************************************/

package main

import (
	"fmt"
)

func main() {
	total := 5 + Beer(5, 5)
	fmt.Println(total)
}

func Beer(bottleCap, bottle int) int {
	a, b := bottleCap/4, bottle/2
	if a > 0 || b > 0 {
		bottleCap = a + b + bottleCap%4
		bottle = b + a + bottle%2
		fmt.Printf("%d 瓶盖， %d 酒瓶 \n", bottleCap, bottle)
		return a + b + Beer(bottleCap, bottle)
	}
	fmt.Printf("%d 瓶盖， %d 酒瓶 \n", bottleCap, bottle)
	return 0
}
