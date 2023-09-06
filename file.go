package main

import (
	"fmt"

	"os"
)

func checkfile(filepath string) bool {

	fmt.Println("Hello")

	if v, err := os.Stat(filepath); err != nil {

		if os.IsNotExist(err) {

			return false
		}

		fmt.Println("File info is :", v)

	} else {

		fmt.Println("File info is :", v)
	}

	return true
}

func main() {

	st, err := os.Stat("demo.txt")

	if err != nil {

		panic(err)

	}

	fmt.Println("Modification time:", st.ModTime())
	fmt.Println("File mode :", st.Mode())

	fmode := st.Mode()

	if fmode.IsRegular() {
		fmt.Println("This is regular file")
	}

	fmt.Println("File size :", st.Size())
	fmt.Println("Is dir :", st.IsDir())
	myfile := checkfile("demo.txt")

	fmt.Println("File exists : ", myfile)

}
