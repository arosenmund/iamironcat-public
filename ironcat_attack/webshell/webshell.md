# Webshell


Standard web shell that will run on linux or windows, listening for connections.

1. First you need to create a new version of the file with binary obfuscation. In the **Attack Desktop** open a terminal.
2. Change to the workshop directory hosting the ironcat_attack modules.
`cd /home/pslearner/workshop/ironcat_attack`
3. From here run the generate_payload.sh script agains the /webshell/webshell.go file.
``
4. This will generate the new .exe in the folder and host it for you on a simple web server.
5. Now, open the **Windows Target** device.
6. Open Firefox.
7. Browse to the Attack Desktop at **http://172.31.24.10:8000**
8. Navigate to the webshell.exe file, and click it to download.
9. Open the command prompt by right clicking the dektop icon and choosing "Run as Administrator"
10. Change directory to the newly downloaded file. `cd c:\Users\





