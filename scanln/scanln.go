package runtest

import "fmt"

func main() {
	for {
		var input string
		fmt.Scanln(&input)
		fmt.Println("vveli: ", input)
		if input == "hvatit" {
			fmt.Print("ENDING")
			break
		}
	}

}
