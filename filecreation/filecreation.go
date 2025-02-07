package filecreation

import (
	"fmt"
)

func destinationChoose() {

	var input string

	fmt.Println("Choose the directory in which the Anime should be downloaded:")
	fmt.Scan(&input)
	fmt.Println(&input)

}
