<!-- .slide: class="title" -->

# Configuring Clients

<img src="taming_velociraptors.png" class="fixed" style="right: 0px;  z-index: -10;"/>

---

<!-- .slide: class="content" -->

## Deploying clients - Windows

* We typically distribute signed MSI packages which include the
  client’s config file inside them.  This makes it easier to deploy as
  there is only one package to install.

* Velociraptor can create the MSI to target the correct Org using the
  `Server.Utils.CreateMSI` artifact.

---

<!-- .slide: class="content small-font" -->
## Domain deployment

We can deploy the MSI to the entire domain using group policy.

Two Methods:
1. Via scheduled tasks.
2. Via assigned software.

---

<!-- .slide: class="full_screen_diagram small-font" -->

Create a share to serve the MSI from.

![](making_share.png)

---

<!-- .slide: class="full_screen_diagram small-font" -->

Ensure everyone has read access from this share - and only administrators have write access!

![](setting_share_permissions.png)

---

<!-- .slide: class="full_screen_diagram small-font" -->

Use the group policy management tool create a new Group Policy Object in the domain (or OU)

![](creating_gpo.png)

---

<!-- .slide: class="full_screen_diagram small-font" -->
Edit the new GPO

![](editing_gpo.png)

---
<!-- .slide: class="full_screen_diagram small-font" -->

![](editing_gpo_1.png)

---

<!-- .slide: class="full_screen_diagram small-font" -->
Ensure the new scheduled task is run as system

![](editing_gpo_2.png)

---

<!-- .slide: class="full_screen_diagram small-font" -->

Using scheduled tasks you can run any binary - use this method to run
interactive collection if you do not have a dedicated Velociraptor
server

![](editing_gpo_new_action.png)

---

<!-- .slide: class="full_screen_diagram small-font" -->
Ensure the new scheduled task is run only once

![](editing_gpo_run_once.png)

---

<!-- .slide: class="full_screen_diagram small-font" -->

Method 2: install via assigned software packages in GPO. The main
advantage here is that it is possible to upgrade or uninstall
Velociraptor easily

![](editing_gpo_assigned_software.png)

---

<!-- .slide: class="full_screen_diagram small-font" -->

You will need to wait until group policy is updated on the endpoint or
until the next reboot. The endpoint must be on the AD LAN

![](editing_gpo_assigned_software_2.png)
