package osCommands

import (
	"fmt"
	"os"
)

func OsFileCommands() {
	defer (func() {
		fmt.Println("defered func - close files here!")
	})()

	args := os.Args

	hostName, err := os.Hostname()
	if err != nil {
		fmt.Println("Failed to get hostname")
		return
	}

	exec, err := os.Executable()
	if err != nil {
		fmt.Println("Failed to get executable")
		return
	}

	fmt.Println("Hostname: ", hostName)
	fmt.Println("Executable: ", exec)

	fmt.Println("Effective Group Id: ", os.Getegid()) // returns -1 on windows
	fmt.Printf("Args: %#v \n", args)

	// current program to exit with the given status code. Conventionally, code zero indicates success,
	//  non-zero an error. The program terminates immediately; deferred functions are not run even with exit 0(success).
	os.Exit(0)
}
