# Ironcat Public Release V1


Do not run on production systems that contain business information.  Though there is a decryption function build in! There is no guaruntee that your data will not be corrupted. Ironcat is not liable for misuse resulting in destruction of property.

## Use

1. Open windows defender or anti-virus settings and turn of "cloud" or "real time submission". We don't want to run the fun for anyone else with the free and public version.
> If you are already getting virus detections and blocks...Pat your anti-virus on the back and then turn it off to test other detection mechanisms, or SOC functions. Remember anti-virus is easily evaded.

1. Unzip the iaiprv1.zip file using the password: **ironcat**

1. Move to the desired execution location. Best results with execution NOT in the **C:\Users** folder.

1. Open an administrative command propmt.

1. Change directory to the location with the iaiprv1.exe file.

1. Run the encyprt command. `iaiprv1.exe encrypt <pw>` where <pw> is a password key that you supply.
> It is criticl that you remember this password key for decryption.

1. Watch it work and see if you are able to find the IOC's and Behaviors in your tooling or SOC response.



## Emulated IOCs & Behavior




## Detections 

The yara signature generated for the original non-public release version of ironcat ransomware is included for reference.  **win.ironcat_auto.yar**

Use: `yara -s win.ironcat_auto.yar <./folder with ironcat executable>`







