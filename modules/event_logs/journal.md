<!-- .slide: class="content small-font" -->

## The Linux Journal Logs

* Recent Linux systems use `systemd`
* Logs are structured and incorporate indexes for fast searching.

* You can view the logs with `journalctl`

```
journalctl --file /run/log/journal/*/*.journal
```

---

<!-- .slide: class="content small-font" -->

## Exercise: Parsing Journal Logs

* It is useful to be able to parse the log file directly.

* Write a VQL parser for the Journal log file.
* The format is documented
  https://www.freedesktop.org/wiki/Software/systemd/journal-files/

* Get a sample file from https://github.com/Velocidex/velociraptor/raw/master/artifacts/testdata/files/system.journal
