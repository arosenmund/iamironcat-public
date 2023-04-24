# Webshell IOCs

Activity:

- Runs as a process by the name of the executable in the command line.
- Creates a firewall rules to listen on port 59679 called Chrome-Sync
- Listens on 443 on any interface.
- Will connect with a web page allowing external entities to execute commands.