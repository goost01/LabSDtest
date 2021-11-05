package main

import (
	"fmt"
	"io/ioutil" // Handling Files
	"log"
	"os"
	"strconv" // Converting Number to String
)

func main() {

	pdata := []byte("All my data i want to write to a file")

	err := ioutil.WriteFile("player.txt", pdata, 0777)

	if err != nil {
		fmt.Println(err)
	}

	// Removing Files Players
	for r := 0; r < 10; r++ {
		for i := 0; i < 10; i++ {
			err = os.Remove("jugador_" + strconv.Itoa(i) + "__ronda_" + strconv.Itoa(r) + ".txt")

			if err != nil {
				log.Fatal(err)
			}

		}
	}

}
