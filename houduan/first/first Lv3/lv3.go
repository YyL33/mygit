package main

import "fmt"

func absoultemin(a [5]int) int {

	for i := 0; i < len(a); i++ {
		if a[i] < 0 {
			a[i] = -1 * a[i]
		}
	}

	for j := 0; j < len(a); j++ { //遍历数组并按从小到大排序
		for k := j + 1; k < len(a); k++ {
			if a[j] > a[k] {
				temp := a[j]
				a[j] = a[k]
				a[k] = temp
			}
		}
	}
	fmt.Println(a)
	var min = a[0]

	fmt.Printf("其中最小绝对值为:%d", min)

	return min

}
func main() {
	var a [5]int
	var i int
	fmt.Println("请输入5个数：")
	//fmt.Scanf("%d", &num)
	for i = 0; i < 5; i++ {
		fmt.Printf("请输入数组a[%d]的值", i)
		fmt.Scanf("%d", &a[i]) //接受数组的收值，一个一个输入
	}
	absoultemin(a)
}
