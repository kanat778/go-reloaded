package reloaded

import "fmt"

func BracketsFix(arr []string) []string {
	//for k := 0; k < len(arr); k++ {
	//	fmt.Println(arr[k])
	//
	//}

	fmt.Println(arr)
	counter := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] == "(" || arr[i] == "[" || arr[i] == "{" {
			for j := i + 1; ; {
				if arr[j] == "\n" {
					counter++
					j++
				} else {
					break
				}
			}
			arr[i] = arr[i] + arr[i+counter+1]
			arr = append(arr[:i+1], arr[i+counter+2:]...)
			//i--
		}

		counter = 0

	}
	//fmt.Println("BracketsFix =", arr)
	//
	fmt.Println("+++++++++++++++++++++++++++++++++++++++++++++++")
	//for k := 0; k < len(arr); k++ {
	//	fmt.Println(arr[k])
	//
	//}
	fmt.Println(arr)

	for l := 0; l < len(arr); l++ {
		if arr[l] == ")" || arr[l] == "]" || arr[l] == "}" {
			for v := l - 1; ; {
				if arr[v] == "\n" || IsCommand(arr[v]) {
					counter++
					v--
				} else {
					break
				}
			}
			arr[l-counter-1] = arr[l-counter-1] + arr[l]
			arr = append(arr[:l], arr[l+1:]...)
			l--
		}
		counter = 0
	}
	//fmt.Println("==================================")
	//for k := 0; k < len(arr); k++ {
	//	fmt.Println(arr[k])
	//}
	return arr
}
