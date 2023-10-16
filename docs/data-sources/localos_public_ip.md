---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "localos_public_ip Data Source - local-os"
subcategory: ""
description: |-
  public_ip data source gets the public IP of the machine that reads the data source as a CIDR.
---

# localos_public_ip (Data Source)

`public_ip` data source gets the public IP of the machine that reads the data source as a CIDR.



<!-- schema generated by tfplugindocs -->
## Schema

### Read-Only

- `cidr` (String, Sensitive) /32 public IP CIDR of machine running terraform,


    This is useful when deploying e.g. test infrastructure for which you want to only grant access to your own workstation.
    You can use this to set up firewalls, cloud security groups etc.

    Will be empty string if the internet is not accessible from the caller.
- `id` (String) Resource identifier