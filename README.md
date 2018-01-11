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
|        |  3. exit_status  |  ~/runnerd    |
|        | <----------------+   2. $cmd1 p1 |
+---------                  +---------------+
```

