### Swerver

Swerver is a basic server management utility that provides a web interface for managing a server on a LAN. I use it to manage my Raspberry Pi.

**Features:**

-   Run scripts with terminal output for the command
-   Monitor, start, and stop services
-   System information including memory usage, uptime, and average load
-   Interactive config generation for rapid setup/deployment

Do **not** use swerver outside of a LAN protected by a secure firewall. Swerver provides only HTTP Basic authentication via browser prompt. This is insecure and should not be used in a public setting.

### Installation and Configuration

```
$ git clone https://github.com/destinmoulton/swerver.git
$ cd swerver
$ go get github.com/gin-gonic/gin
$ go get github.com/gin-contrib/sessions
```

#### Configuration

The configuration file is stored in `~/.config/swerver/swerver.config.toml`.

To generate a configuration file, run `swerver`:

```
$ cd swerver
$ ./swerver
```

### Setting Swerver Up as a Service

Create a service file `/lib/systemd/system/swerver.service`:

```
[Unit]
Description=swerver

[Service]
Type=simple
Restart=always
RestartSec=5s
ExecStart=/home/pi/swerver/swerver
WorkingDirectory=/home/pi/swerver
User=pi

[Install]
WantedBy=multi-user.target
```

Modify the `ExecStart` and `WorkingDirectory` fields to match the installation directory where you installed swerver.

Now you can use the standard systemctl commands:

```
$ sudo systemctl start swerver
$ sudo systemctl status swerver
$ sudo systemctl stop swerver
$ sudo systemctl status swerver
```

### License

Swerver is open source under the MIT license.
