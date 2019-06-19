# Akamai Client

A golang package which helps facilitate making HTTP requests to [Akamai OPEN APIs](https://developer.akamai.com)


## This Fork
### ntpcheat
The parent uses the local RTC to generate the timestamp for the messages.  If the timestamps differ by > +/- 30s, edgegrid won't auth the request (403 Forbidden is instead returned).

If you cannot maintain your RTC in alignment with Akamai, for example your organizaton's NTP server is out of alignment, or you can only connect to it infrequently due to VPN, then you are not going to have a good time.

This fork replaces the makeEdgeTimeStamp() function in edgegrid/signer.go with a version that retrieves the time from <http://time.akamai.com/?iso> immediately before each test, thus bypassing your RTC and the dependency on a sane NTP server.
