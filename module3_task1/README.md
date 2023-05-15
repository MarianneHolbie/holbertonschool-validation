<h1 align="center"> Module 2: Testing in the Software Development Methodology

# Learning Objectives
This project aims at practicing with automated tests. The goal is to understand the pros and cons of different testing methods to be able to understand the value of doing, or not doing, a kind of test.

After this project, you should be able to:

Understand what linting is the extent of its usages (which kind of file can be linted, and the impact of running it often)
Understand the difference between unit tests and integration tests
Use code coverage as a helper to write tests
Understand that not only “classical” code is to be tested, but also a lot of the artifacts we can generate
Understand how “component”-based testing for acceptance and end to end validation is to be used

## Prerequisites

The following elements are required in addition to the previous module (“Module 1: Introduction to DevOps: Automate Everything to Focus on What Really Matters”) prerequisites.:

### Concepts

You should have a basic knowledge of the following concepts:

- What a compiled language is (C/C#/Golang/Rust/etc.)

	- Generation process from source to executable binary
	- Basic types: string, integer, boolean, maps, arrays
	- Basic algorithmic: loops, conditional, functions
	
- Installing command line tools with NPM (in addition to package managers)

- Understand the basics of the HTTP protocol (client/server, verbs, headers)

### Tooling

This project needs the following tools / services:

- An HTML5-compliant web browser (Firefox, Chrome, Opera, Safari, Edge, etc.)
- A free account on [GitHub](https://github.com/), referenced as `GitHub Handle`
- A shell terminal with bash, zsh or ksh, including the standard Unix toolset (ls, cd, etc.)
- [GNU](https://www.gnu.org/software/make/) Make in version 3.81+
- [Git](https://git-scm.com/book/en/v2/Getting-Started-The-Command-Line) (command line) in version 2+
- [Go Hugo](https://gohugo.io/) v0.84+
- [Golang](https://intranet.hbtn.io/rltoken/5ypbIenKj6LiymRm619--A) v1.15.*
- [NPM](https://intranet.hbtn.io/rltoken/RcU82lwHHO4xEQCtWEv1sg)v7 + [NodeJS](https://intranet.hbtn.io/rltoken/XWIqoQhjv16uVWfGbCdInw) v14.*
- [markdownlint-cli](https://intranet.hbtn.io/rltoken/hplwMW8M8BKVQyhDso0pOw) v0.26.0
- [markdown-link-check](https://intranet.hbtn.io/rltoken/BRJGBHXvkAUKt50KrFOm0A) v3.8.6
- [Holberton's W3C Validator](https://intranet.hbtn.io/rltoken/ll8gJ8CPoI9tfn1OTDE8rA)
- [yq](https://intranet.hbtn.io/rltoken/9wlxJjlqCE6XyPa6TQ0RsQ) v4.5.0
- [shellcheck](https://intranet.hbtn.io/rltoken/7e95a2wDfOHFQGKJqRlHgg) v0.*
- [yamllint](https://intranet.hbtn.io/rltoken/B1BZ_C_5ANyq005Vd0LWNw) v1.*
- [jq](https://intranet.hbtn.io/rltoken/pVjsOvuSQavip_1Y4u--4Q) v1.*


<h1 align="center"> How to use the make file:

## Lifecycle

In the DevOps methodology, the development lifecycle is generally staying the same. Use the following steps :

To execute the Makefile use the following syntax:
 ```make <command>```

 command are availaible :
* `help`:
    - show all command description

* `hugo-build`:
    - Builds a new version of the website to folder `/dist/` 
 
* `go-build`:
	- compile the source code of the application to a binary named ```awesome-api``` (the name ```awesome-api``` comes from the command ```go mod init github.com/<your github handle>/awesome-api```) with the command go build. The first build may takes some times. Build run only if lint is not failed.

* `build`:
	- execute `hugo-build` AND `go-build`

* `hugo-clean`:
    - Removes the contents the folder  `/dist/`

* `go-clean`:
	- Stop the application. Delete the binary ```awesome-api``` and the log file ```awesome-api.log```

* `clean`:
	- execute `hugo-clean` AND `go-clean`

* `post`:
    - Creates a new post in the contents/post folder with POST_TITLE and POST_NAME set from the ENV variables.

* `check`:
	- Lint of the Markdown source files using command line AND analysis of the links with command line. If one test fails, the command failed.
 
* `validate`:
	- validate the file ./dist/index.html by using command line. But non-blocking quality indicator

* `test`: Test to ensure that it behaves as expected. 

* `lint`: Test lint in the files

* `unit-tests`: Run files with the _test.go suffix 

* `integration-tests`: Run Golang integration tests 

* `run`: Run the application in background by executing the binary ```awesome-api```, and write logs into a file named ```awesome-api.log```

* `stop`: Stop the application with the command kill XXXXX where XXXXX is the Process ID of the application.

## Workflow

* Workflow module3_task 0 : 
-> Clone the repository,
-> Use ubuntu-22.04 as the virtual machine.
-> Position yourself in the correct directory and execute the command make help to validate that the Makefile is present and implements the help target.
-> This workflow must be triggered:
		Each time there is new code pushed on your repository,
		And once per day (whatever time you want).
* Workflow module3_task 1 : 
-> Ensure that the workflow is executed into an Ubuntu 22.04
-> install with script setup.sh
-> 1 job with at least 7 steps (checkout, running setup.sh and then the 5 make commands)

# Story
Continuing your journey as a Software Engineer at Awesome Inc., you want to provide early visibility on your work to your colleagues to allow you to iterate on the most important issues or improvements for the company.

By defining a Software Delivery Pipeline with automated tasks, you’ll ensure that the collaboration between teams is improved, and that your team grows in maturity while providing an efficient process to ensure that you can deliver the application often.