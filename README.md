logs
----

A golang logger


Install
=======

```
	go get github.com/heiing/logs
```


Config
======

see config.json


Usage
=====

Step 1: import

```
	import github.com/heiing/logs
```

### using DefaultLogger ###

Step 2: Log Messages

```
	logs.Error("your message")
	logs.Info("your notice message")
```

### using Customize Logger ###

Step 2: Parse a config file to LogsConfig or create a LogsConfig.

Step 3: Create a Logger and log messages

```
	logger := logs.NewLogger(&LogsConfig)
	logger.Error("your message")
	
	// or set it to the logs
	logs.SetDefaultLogger(logger)
	logs.Error("your message")
```