# elogger

Simple Structured logger

Currently only JSON output format

Example:
```var logData elogger.StructuredLog
logData.ID = sessionID
logData.Account = userID
logData.RawInterface = anyInterface
logData.RawInterface = anyInterface

elogger.Event(logData, "error", "eventName", err.Error())