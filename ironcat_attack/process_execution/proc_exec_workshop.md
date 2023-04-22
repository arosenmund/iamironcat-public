# DLL injection

This one is a bit different, payloads are not written in go, and we will not be generating new ones. The same ideas apply to evastion of changing the literal strings, having minimal extraneous code. However, these are compiled with cpp and nim.  Experiementally, even though this is a very generic technique, stock defender did not catch the files, or the execution at the time of creation. <--That is my disclaimer in case this doesn't work!!!\

## Download Payloads.

1. In the Attack Desktop, open a terminal, navigate to the  ironcat_attack/process_execution folder.
`cd /home/pslearner/workshop/ironcat_attack/process_execution`

2. Notice the already compiled injectdll.exe file, as well as the evilm64.dll payload.

3. Time to move the payload over to the Windows Target. Host the file using python http simple server.
`python3 -m http.server`

4. Use the drop down in the top left to connect to the Windows Target device.

6. Open a firefox browser window and browse to http:172.32.14.10:8000

7. Click the two files mentioned above to download it to the device.

8. If there are warnings from smart screen, acknoledge them and choose "keep anyway"
Note: this is not the same service as windows defender, it is an artifact of us using the browser to move it over, not the payload being caught by anti-virus.

9. Open command prompt and change directory to the downloads folder.
`cd c:\Users\pslearner.PS-WIN-1\Downloads`
10. Move the evilm64.dll file to the c:\Users\Public\Desktop location.
`cp evilm64.dll c:\Users\Public\Desktop`
11. Run an instance of notepad.exe to inject into.
`notepad.exe`
12. Run `tasklist` and look for **notpade.exe**. Identify the PID or 4 digit number associated with the running process.
13. Run injectdll.exe with the PID from notepad.exe as an argument. ex. `injectdll.exe 4564`
Profit.

## Questions To Ask
Did I get any alerts.
How would I know a process had a DLL injected?
What IOCs could you monitor for to look for this activity?

It is up to you and your tools on this one.  We can discuss as a group, but it is likely time to wrap it up! Thank you for coming.

All of this including the lab will be available all week. Feel free to play around with these tools, pull down the project for use back home, and even modifiy some of the IOCs in the code we talked about.

Remember you CAN use catamorph in this same repo to rapidly modify the exe with identifiable IOCs and a new hash without a development enviornment.