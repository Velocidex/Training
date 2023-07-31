<!-- .slide: class="title" -->

# Troubleshooting

### When things do not go to plan!

---

<!-- .slide: class="content" -->

## Server fails to start

* By default the server is running as a systemd service
* Error message is not the most useful:

```
● velociraptor_server.service - Velociraptor linux amd64
    Loaded: loaded (/etc/systemd/system/velociraptor_server.service; enabled; vendor preset: enabled)
    Active: activating (auto-restart) (Result: exit-code) since Fri 2021-12-31 15:32:58 AEST; 1min 1s ago
    Process: 3561364 ExecStart=/usr/local/bin/velociraptor --config /etc/velociraptor/server.config.yaml frontend (code=exited, status=1/FAILURE)
    Main PID: 3561364 (code=exited, status=1/FAILURE)

```

---

<!-- .slide: class="content small-font" -->

## Starting the server manually

* You can start the server manually to see any error messages from
  Velociraptor
  * Be sure to stop the service first `service velociraptor_server stop`
  * Make sure to switch to the `velociraptor` user first (otherwise
    filesystem permissions can be messed up)

```bash wrap
# sudo -u velociraptor bash
$ velociraptor frontend -v
Dec 31 15:47:18 devbox velociraptor[3572509]: velociraptor.bin: error: frontend: loading config file: failed to acquire target io.Writer: failed to create a new file /mnt/data/logs/Velociraptor_debug.log.202112270000: failed to open file /mnt/data/logs/Velociraptor_debug.log.202112270000: open /mnt/data/logs/Velociraptor_debug.log.202112270000: permission denied
```

* Fix with `chown -R velociraptor:velociraptor /path/to/filestore/`

---

<!-- .slide: class="content small-font" -->

## Debugging client communications

* What could go wrong?

   1. No connectivity between client and server
   2. Unable to establish secure comms.

* To see client logs run the client manually with the `-v` flag.

```
velociraptor.exe --config client.config.yaml client -v
```

---

<!-- .slide: class="full_screen_diagram" -->

## Network Connectivity problems

![](./network_comms_problems.png)

---

<!-- .slide: class="content small-font" -->

## Network Connectivity problems

You can verify network connectivity and TLS configuration by using
curl to fetch the server certificate:

```
curl.exe -k https://server:8889/server.pem
```

* For self signed deployments curl needs the `-k` flag to ignore
  untrusted certificates.

---

<!-- .slide: class="content small-font" -->

## Network Connectivity problems

* Captive portals may interfere with the communication

![](./captive_portal.png)

---

<!-- .slide: class="content small-font" -->
## Network Connectivity problems

* You can view the certificate details by using openssl.
   * Check for expiry times

```sh
curl https://test.velocidex-training.com/server.pem | openssl x509 -text
```
