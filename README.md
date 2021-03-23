# elogger

Simple Structured logger

Currently only JSON output format

v1.0.3 What's new

* EventTag func to log with event name only. mostly for the event dashboard
* Error func to log error with the event name in pointer receiver fashion
* Fatal func to log and kill if err != nill. mostly for critical env missing or in case of resource connection fail

Example:
```var logData elogger.StructuredLog
logData.ID = sessionID
logData.Account = userID
logData.RawInterface = anyInterface

elogger.Event(logData, "error", "eventName", err.Error())

logData.EventTag("AwesomeThingDone")

logData.Error("testEventLog", "test event logger")

logData.Fatal("dbConFail", "Can't connect to database")
```
