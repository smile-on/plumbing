# Plumbing
The set of tools that helps instantiates HTTP services built from pipeline of command line scripts.  This is meant to be a dirty quick way to have bunch of small services up and running in no time.

The key difference from other minimalistic "frameworks" - **no dependencies** on 3rd party or assumed IT infrastructure. Just have it up and running for proof of concept demo. Security is not in a scope of concerns.  Prove that the demo works - this is begin and end of our scope. Be aware!

**Rationality**

I am tired fixing Python scripts that has tons of ever changing dependencies.

I do not have luxury of time to spend on IT services every time I need to kick small demo.

My main prototype platform is Linux headless servers. No promises for Windows.
The exit_status is string; "ok" for return code 0 or error description otherwise.


## Use Cases
### 1. a call triggers comand on the server 
```
synchronous flow
+--------+ 1. GET /cmd1/p1  +---------------+
|        +----------------> |               |
| Client |                  | Server        |
|        |  3. exit_status  |  ~/piped    |
|        | <----------------+   2. $cmd1 p1 |
+---------                  +---------------+
```
### 2. comand started by scheduler waits for trigger file before proceed.
Example of waiting for trigger file with 60 seconds timeout.
```
$ wait4file -t 10 -n 6 //share/trigger.file
return code 0 as soon the trigger file detected
otherwise returns code 1
```

## Installation
Simply build project with Go tool on your platform.

```bash
mkdir prj; cd prj
export  GOPATH=`pwd`
go get github.com/smile-on/plumbing/runner
cd src
go install github.com/smile-on/plumbing/cmd/piped

```


### How To Run 

Create ini file for your service. Below is an example of serviceTouch.ini:
```ini
[/touch/{p1}]
touch {{.p1}}
```
Now you can use piped executable with your configuration.
```bash
piped -ini serviceTouch.ini -listen :8080
```

Note, you may use piped binary on any computer of that platform just by copying one binary file and creating configuration file. Pipe daemon does not have any dependencies beside standard GO runtime requirements.

Enjoy!
