/**************************************************
*¡¡Æ¡¾Æ£º2ÔªÇ®1Æ¿
*¡¡¡¡¡¡¡¡4¸öÆ¿¸Ç»»1Æ¿Æ¡¾Æ
*¡¡¡¡¡¡¡¡2¸ö¿ÕÆ¿»»1Æ¿Æ¡¾Æ
*¡¡¡¡ÎÊ£º10Ôª¿ÉÒÔºÈ¼¸Æ¿£¿
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
		fmt.Printf("%d Æ¿¸Ç£¬ %d ¾ÆÆ¿ \n", bottleCap, bottle)
		return a + b + Beer(bottleCap, bottle)
	}
	fmt.Printf("%d Æ¿¸Ç£¬ %d ¾ÆÆ¿ \n", bottleCap, bottle)
	return 0
}
