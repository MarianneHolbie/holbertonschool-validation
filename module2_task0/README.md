# Testing in the Software Development Methodology

## Build an Application using Make

### Prerequisites
* Golang in v1.15.+
* NPM v7+ with NodeJS v14.*
* Python3 with pip module
* golangci-lint

### Project Life-cycle automate with Makefile 
The life-cycle of this project is the following command:

* ```build```  compile the source code of the application to a binary named ```awesome-api``` (the name ```awesome-api``` comes from the command ```go mod init github.com/<your github handle>/awesome-api```) with the command go build. The first build may takes some times.

* ```run```: Run the application in background by executing the binary ```awesome-api```, and write logs into a file named ```awesome-api.log```

* ```stop```: Stop the application with the command kill XXXXX where XXXXX is the Process ID of the application.

* ```clean```: Stop the application. Delete the binary ```awesome-api``` and the log file ```awesome-api.log```

* ```test```: test to ensure that it behaves as expected. 

* ```help```: print a list of all the goals