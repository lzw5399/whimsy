// Goto
package Loop

import "fmt"

func main() {
	for x := 0; x < 10; x++ {
		for y := 0; y < 10; y++ {
			if y == 2 {
				goto breakHere
			}
		}
	}
	// 手动返回, 避免执行进入标签
	return
	// 标签
breakHere:
	fmt.Println("done")
	/*xx:
	fmt.Println("123")*/
}
