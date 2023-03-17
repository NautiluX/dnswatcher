# DNS Watcher

TMUX based utility for observing multiple DNS names and records

## Usage

Create a config file like the following:

```
---
targets:
- nameserver: 1.1.1.1
  name: google.com
  recordtype: NS
- nameserver: 1.1.1.1
  name: google.com
  recordtype: A
- nameserver: 8.8.8.8
  name: google.com
  recordtype: NS
- nameserver: 8.8.8.8
  name: google.com
  recordtype: A
```

Run dnswatcher:


```
./dnswatcher example.yaml
```

Observe DNS responses.

To stop, press CTRL+B D (detach from tmux session).
