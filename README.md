# iamironcat-public

Public assets for the  blueteam ironcat threat emulation suite.
B.I.T.E.S

Youtube Channel:

[Ironcat Malware Youtube Series](https://youtube.com/playlist?list=PLle9BGvw9HqKWkUfmB9K-jcI_2UGCIVfB)


## Ironcat Attack App

This has now been upgraded as an application to be run on the target machine. A binary called "ironcat_attack" for both linux and windows has been provided. 

On windows, it will run on port 8080, local host, and you can launch the TTPs included. The TTPs are wrapped in executables, and the information for each is included in the application.

On linux, as long as tcpreplay is installed, launching the network attack will replay a pcap to the main interface. If you have a network analysis tool listening, it will catch this traffic. If you are in your own environment, this will include your environment traffic.

This is just V1, upgrades will be implemented to make this even easier.

If you would like to contribute, let me know!

## Public Release V1.1 Ransowmare in Attack Scenarios

Scenario Attacks currently include the Ironcat v1 Malware.
[Ironcat_Public_V1](./scenario_attacks/)


## Attack Modules

These are now found in 2 locations:

[Network Attacks](./network_attacks/)

[Endpoint Attacks](./endpoint_attacks/)

### What is dis?
These are small code wrappers for windows and linux that will create an executable that writes a .bat or .sh file, with your batch or bash script, and then executes it.

To test the detection of TTP's that leverage system commands, simply modify the strings that are written to the script file with code to emulate a specific TTP activity.  Then compile and run on your test machine.  This should enabled low code blue teams to rapidly create tests for TTPs identified through incident response or threat intelligence, and then test their tooling, detections, or mitigations.

As time goes on, various samples will be provided. Updated with iaiprvv1.1 for defender evasion as of 1-25-23.

## Catamorph

[Catamorph](./catamorph/)

This is an intentionally basica capability that will allow you modify any executable program you create with the Ironcat suite to give it a new hash value. Using this you can rapidly test for detection of the activity within binary rather than the detection of a hash that is easily modified, but you don't have to actually recopile the program.
