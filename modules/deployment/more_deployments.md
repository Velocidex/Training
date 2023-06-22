<!-- .slide: class="title" -->

# Additional deployment scenarios

---

<!-- .slide: class="content" -->

## Using multiple OAuth providers

* It is possible to use multiple oauth providers at the same time.
* This is very useful when you need to provide access to people
  outside your org!

TODO

---

<!-- .slide: class="content" -->

## Exercise: Additional OAuth provider: Azure

* Add Azure as an alternate authentication provider!
* The instructor will illustrate the Azure process for obtaining credentials
* The instructor will provide the Azure credentials for your own VM.

---

<!-- .slide: class="content" -->

## Client certificate based authentication.

* You do not need to use any authentication provider at all!
* It is possible to rely solely on client side certificates

TODO

---

<!-- .slide: class="content" -->

## Multi-Frontend deployments

* We recommend single server deployment for networks < 20k endpoints.
* For larger networks, we recommend multi-frontend deployments.
* These require a shared network filesystem.

TODO

---

<!-- .slide: class="content" -->

## Exercise: Create a Master/Minion deployment

* Create a multi-frontend deployment running on the same host VM
* This allows the GUI to run on a separate process and allows unloaded
  post processing/notebook operations.

---

<!-- .slide: class="content" -->

## Exercise: Customizing the dashboard

* When running multiple deployments or multiple orgs it is convenient
  to customize the dashboard.

* Customize the dash board to add your name to the main page. This
  helps identify your deployment.

---

<!-- .slide: class="content" -->

## Server Lockdown Mode

* Velociraptor is an extremely powerful tool.
* A Velociraptor Server Admin account takeover can be very dangerous!
* But we still want to have it available so we can respond quickly.
* `Server Lockdown Mode` prevents Velociraptor from performing any
  destructive actions while in lockdown!

Add the following to the `server.config.yaml` and restart the server

```
lockdown: true
```
