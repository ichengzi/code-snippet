/**************************************************
*��ơ�ƣ�2ԪǮ1ƿ
*��������4��ƿ�ǻ�1ƿơ��
*��������2����ƿ��1ƿơ��
*�����ʣ�10Ԫ���Ժȼ�ƿ��
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
		fmt.Printf("%d ƿ�ǣ� %d ��ƿ \n", bottleCap, bottle)
		return a + b + Beer(bottleCap, bottle)
	}
	fmt.Printf("%d ƿ�ǣ� %d ��ƿ \n", bottleCap, bottle)
	return 0
}
