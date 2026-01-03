# SSL Labs Grade Checker

Hello there.

To use the simple implementation of the challenge, you just have to open the root folder and execute:

```bash
go run .
```


The program will analyze the SSL/TLS configuration of a given host and return a grade.
It can sometimes take quite a while to complete depending of the status of the servers of SSL Labs

# Grade meaning

The returned grade can be interpreted as follows, based on the official SSL Labs documentation:
https://success.qualys.com/support/s/article/000005828

A+ – Exceptional configuration

A – Strong commercial security

B – Adequate security with modern clients, but uses older and potentially obsolete cryptography with older clients; may contain minor configuration issues

C – Obsolete configuration; uses obsolete cryptography even with modern clients; may contain more significant configuration problems

D – Configuration with security issues that are typically difficult or unlikely to be exploited, but can and should be addressed

E – Unused

F – Exploitable and/or patchable problems, misconfigured server, insecure protocols, etc.