# RDP Brute Force

1 to 1 Brute Force of RDP Credentials from 10.0.0.126 to 10.0.0.111

RDP Brute force can be difficult to detect if the auditing for failures is not turned on and those subsequent logs are not monitored.

Through traffic analysis there are two ways to understand what is happening here.

The first is by the number of attempted connections.

The second is by the size of a failed connections vs a successful one.

On key bit of data is the mstshahs=administrator field which will show the acount that is attempting to login. Remember RDP is encrytped by default once the full session is esablished.

References:

[Security Event Triage: Detecting Malicious Traffic with Signature and Session Analysis](https://www.pluralsight.com/courses/set-detecting-malicious-traffic-signature-session-analysis)

[Security Event Triage: Detecting Network Anomalies with Behavioral Analysis](www.pluralsight.com/courses/set-detecting-network-anomalies-behavioral-analysis)

