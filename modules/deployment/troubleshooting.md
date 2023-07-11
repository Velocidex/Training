<!-- .slide: class="title" -->

# Troubleshooting deployments

## When things do not go to plan!

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

## Velociraptor Network communications

* Velociraptor clients connect to the server over HTTPS POST
  messages.
   * URLs used are set in the config file `Client.server_urls` list.

* The communication is encrypted/signed using the Velociraptor
  internal PKI. You can not change this!

<img src="post_message_format.png" class="mid-height">


---

<!-- .slide: class="content small-font" -->

## HTTP based TLS security

* Communication occurs over stadard TLS POST
* Two supported modes:
   1. Self signed mode (`Client.use_self_signed_ssl: true`)
      * Server presents a certificate issued by the internal Velociraptor CA
      * Client verifies cert directly
      * Client **does not** trust public PKI or on host root store!

   2. PKI mode
      * Client uses root CA chains to verify connections (public CAs,
        Host root store or embedded root certs).
      * This is suitable for MITM proxies.
      * Add trusted certs to `Client.Crypto.root_certs` in PEM format

---

<!-- .slide: class="content small-font" -->

## Debugging client communications

* When things go wrong.... Start client manually

```
velociraptor --config client.config.yaml client -v
```

* Sometimes network filtering can prevent a connection.
    1. Test by connecting to the server using `curl` to fetch the
       server's internal certificate.

```
curl.exe -v https://example.com/server/pem
```
