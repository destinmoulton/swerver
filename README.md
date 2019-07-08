### Swerver

Swerver is a basic server management utility that provides a web interface for managing a server on a LAN. I use it to manage my Raspberry Pi.

**Features:**

-   Run scripts
-   Monitor, start, and stop services
-   System information including memory usage, uptime, and average load

Do **not** use swerver outside of a LAN protected by a secure firewall. Swerver provides no user authentication and it would be dumb to use it on anything that contains sensitive information.

### Installation and Configuration

```
$ git clone https://github.com/destinmoulton/swerver.git
$ cd swerver
$ go get github.com/joho/godotenv
$ go get github.com/gin-gonic/gin
```

#### Configuration

```
$ cd /swerver-install-path
$ cp .env.template .env
```

Configure the values in `.env`.

```
$ go build swerver
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
