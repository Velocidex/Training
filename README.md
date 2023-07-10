# Velociraptor: Digging Deeper
## The official training workshop

This repository contains the official Velociraptor training workshop
material.

## How to use this repository.

First build the HTML for this repository into a new directory.

```
go run ./src/ generate /tmp/output
```

You can serve the directory over the network:

```
go run ./src/ serve --directory /tmp/output --port 1313
```

Alternative you can simply open the HTML files in your browser.
