# HTTP Workshop Lab


# Attack Desktop
##
1. In the Attack Desktop, open a terminal, navigate to the  ironcat_attack/c2/ichttpc2 folder.
`cd /home/pslearner/workshop/ironcat_attack/c2/ichttpc2/client`

2. Notice the already compiled client.exe file.  Check it's md5 hash using md5sum.
`md5sum client.exe`

3. Now check the 

4. Time to move the payload over to the Windows Target. Host the file using python http simple server.
`python3 -m http.server`

5. Use the drop down in the top left to connect to the Windows Target device.

6. Open a firefox browser window and browse to http:172.31.24.10:8000

7. Click **client.exe** to download it to the device.

8. If there are warnings from smart screen, acknoledge them and choose "keep anyway"
Note: this is not the same service as windows defender, it is an artifact of us using the browser to move it over, not the payload being caught by anti-virus.


Before you run the client, you need to start the server, that makes sense right??

## Run the Server

1. Swap back to the **Attack Desktop**.

2. In the terminal window with the ironcat_attack folder open, stop the python server with `ctrl+c`, and change to the server directory.
`cd /home/pslearner/workshop/ironcat_attack/c2/ichttpc2/server`

3. This side does not require obfuscation, so you can simply run the pre-compiled binary.
`sudo ./server`

4. When it says "Enter Mode" type 0 and press enter.  There will then start up and is ready to recieve connections.

Now let's create a happy connection.

## Run the client

1. Swap back to the Windows Target.

2. Use the desktop icon for CMD. Right click, and choose 'Run As Administrator'

3. Change directory to the downloads folder. (This one is a little tricky.)
`cd C:\Users\pslearner.PS-WIN-1\Downloads`

4. Run the client using 2 or more arguments, they can be anything.
ex. `client.exe 1 1`
Note: Unless something has changed significantly for the worse for my workshop, you should have ZERO notifications from anti-virus.

5. Swap back to your Attack Desktop, inspect the terminal where you ran the server executable. You will see connections from the client.  It will checkin ever 10 seconds. You will then be queried for input for "mode" 
    - if you input 0. It will continue to checkin
    - Input 1, and it will perform os enumeration
    - input 2 and it will return asking you for a command, full shell
For the purpose of this part of the workshop, you don't need to interact, time to see if you can identify artifiacts of this communication.

## Identify IOCs

You have many tools as your disoposal. Velociraptor is most powerful, but sysmon as well as the full sysinternals sweet is available.  The point of this workshop is emulation but I will give a bit of time for you to be able to use the tools available and see what ways people come up with to identify the activity.

REMEMBER! You know what you have done, but how would you specifically identify this activity as malicious in a enterprise environment?  The point of this, is for you to take this home, and truly test your technology, people, and processes.

## Sysmon
- Symon is already installed, but you can find the Event Log in the Windows Event Logs.
- Sysinternals is in the PATH, so you can call the tools from the command line
    - procexp
    - tcpview
- Wireshark is on the desktop
- To start velociraptor and run hunts follow the following steps.
    - Open CMD as administrator
    - Chand directory: `cd c:\Users\Public\Desktop\LAB_FILES\`
    - Run velociraptor: `velociraptor.exe gui`
    - Open Firefox and browse to localhost:8889
    - Username: admin Password: passowrd
    - Use the tool bar on the left to run hunts.
    - View the results of your hunts


## Questions to Ask

What are the ports involved, and why do stand out?

Is there anything in the http traffic that would stand out?

If it is being actively used for c2, what kind of signature could you create?


- [HTTP C2 IOCs](./HTTPc2_IOCs.md)

Once you are done. 

1. Stop the client.exe by pressing 'ctrl+c' in the terminal window on the Windows Target.
2. In the Attack Desktop, stop the server by doing the same thing.

Now lets discuss some of the IOCs and detections.









