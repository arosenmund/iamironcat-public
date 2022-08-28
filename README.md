# iamironcat-public

Public assets for the  blueteam ironcat threat emulation suite.
B.I.T.E.S

## Public Release V1 Ransowmare
[Ironcat_Public_V1](./ironcat_ransomware_v1/)


## Attack Modules

[Ironcat Attack Modules](./ironcat_attack/)

### What is dis?
These are small code wrappers for windows and linux that will create an executable that writes a .bat or .sh file, with your batch or bash script, and then executes it.

To test the detection of TTP's that leverage system commands, simply modify the strings that are written to the script file with code to emulate a specific TTP activity.  Then compile and run on your test machine.  This should enabled low code blue teams to rapidly create tests for TTPs identified through incident response or threat intelligence, and then test their tooling, detections, or mitigations.

As time goes on, various samples will be provided. 

## Catamorph

[Catamorph](./catamorph/)

This is an intentionally basica capability that will allow you modify any executable program you create with the Ironcat suite to give it a new hash value. Using this you can rapidly test for detection of the activity within binary rather than the detection of a hash that is easily modified, but you don't have to actually recopile the program.
