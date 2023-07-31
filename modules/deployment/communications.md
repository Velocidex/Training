<!-- .slide: class="title" -->

# Network Communications

## Understanding client/server communications

---

<!-- .slide: class="content small-font" -->

## Network Communications overview

* Communication through HTTPS Post
  * Full duplex communication allows bydirectional data flow.
  * URLs used are set in the config file `Client.server_urls` list.
* Payload is encrypted and signed using the internal Velociraptor PKI
  * Clients generate their own private/public key set
  * Client ID is tied to their private key (Hash)
  * Therefore it is not possible to impersonate a client id without
    compromising the private key.

---

<!-- .slide: class="full_screen_diagram" -->

## Network Communications overview

![](./network_comms.png)

---

<!-- .slide: class="content small-font" -->

## Velociraptor’s internal PKI

* The configuration wizard create an internal CA with an X509
  certificate and a private key. This CA is used to

   1. Create initial server certificates and any additional
      certificates for key rotation.

   1. CA public certificate is embedded in the client’s configuration
      and is used to verify server communications.

   1. The CA is used to create API keys for programmatic access. The
      server is then able to verify API clients.

* The configuration file contains the CA’s X509 certificate in the
  **Client.ca_certificate** parameter
  * It is embedded in the  client configuration.
  * The private key is contained in the **CA.private_key** parameter.

---

<!-- .slide: class="content small-font" -->

## TLS verification

* Velociraptor currently supports 2 modes for deployment via the config wizard:

    1. Self signed mode uses internal CAs for the TLS
       certificates. The client knows it is in self signed mode if the
       **Client.use_self_signed_ssl** flag is true.
      * Server presents a certificate issued by the internal Velociraptor CA
      * Client verifies cert directly
      * Client **does not** trust public PKI or on host root store!
        Only trust included CA certificate.
      * This pins the server’s certificate inside the client

* For MITM configuration make sure to add the CA's root to the config file.

    2. PKI Mode: Proper certificates minted by Let’s encrypt.
      * Client uses root CA chains to verify connections (public CAs,
        Host root store or embedded root certs).
      * This is suitable for MITM proxies.
      * Add trusted certs to `Client.Crypto.root_certs` in PEM format

---

<!-- .slide: class="content" -->

## Adding Proxy and MITM settings

* Some networks require communication via a proxy.
  * Add proxy to config file at `Client.proxy`
  * PAC or Windows Auth not supported.
* Some networks have an SSL MITM preoxy
  * Add root certs to `Client.Crypto.root_certs`
  * Note that this only affects the outer TLS layer - the proxy still
    has no visibility of the comms.

---

<!-- .slide: class="content small-font" -->

## Configuration reference

* Many settings available in the config file.
* The config Wizard produces a reasonable start configuration.
  * You can customize it to deal with different needs.
* Full configuration reference at https://docs.velociraptor.app/docs/deployment/references/

* The knowledge base contains tips for many tasks:
   * https://docs.velociraptor.app/knowledge_base/
   * [How Do I Use My Own SSL Certificates?](https://docs.velociraptor.app/knowledge_base/tips/ssl/)
   * [How To Fix “Certificate Has Expired Or Not Yet Valid Error”?   ](https://docs.velociraptor.app/knowledge_base/tips/rolling_certificates/)
