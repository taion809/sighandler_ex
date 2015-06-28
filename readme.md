Sighandler
==========

Simple signal handler example in Go.  Reloads config on SIGHUP.

Example
=======

#### Pre hup config
```toml
bind_address = "0.0.0.0:5555"
always_respond = "Before config change"
```

#### Launch server
```
☁  sighandler  ./sighandler
2015/06/28 19:09:26 -->> Sighandler Engaged!
2015/06/28 19:09:26 -->> Send me your kill commands.
```

#### Send request
```bash
while
do
  curl localhost:5555/greet
  sleep 1
done
```

#### Examine Server
```
☁  sighandler  ./sighandler
2015/06/28 19:09:26 -->> Sighandler Engaged!
2015/06/28 19:09:26 -->> Send me your kill commands.
2015/06/28 19:11:21 -->> Responding:  Before config change
2015/06/28 19:11:22 -->> Responding:  Before config change
2015/06/28 19:11:23 -->> Responding:  Before config change
2015/06/28 19:11:24 -->> Responding:  Before config change
```

#### Edit config
```toml
bind_address = "0.0.0.0:5555"
always_respond = "After config change"
```

#### Send SIGHUP
```bash
kill -1 $SIGSERV_PID
```

#### Examine
```
☁  sighandler  ./sighandler
2015/06/28 19:09:26 -->> Sighandler Engaged!
2015/06/28 19:09:26 -->> Send me your kill commands.
2015/06/28 19:11:21 -->> Responding:  Before config change
2015/06/28 19:11:22 -->> Responding:  Before config change
2015/06/28 19:11:23 -->> Responding:  Before config change
2015/06/28 19:11:24 -->> Responding:  Before config change
2015/06/28 19:12:26 !! Got  hangup
2015/06/28 19:12:33 -->> Responding:  After config change
2015/06/28 19:12:34 -->> Responding:  After config change
2015/06/28 19:12:35 -->> Responding:  After config change
2015/06/28 19:12:36 -->> Responding:  After config change
```
