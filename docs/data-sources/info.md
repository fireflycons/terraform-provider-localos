---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "localos_info Data Source - terraform-provider-localos"
subcategory: ""
description: |-
  The info data source gets information about the operating system and environment of the machine that is running terraform.
---

# localos_info (Data Source)

The `info` data source gets information about the operating system and environment of the machine that is running terraform.

## Example Usage

```terraform
data "localos_info" "os_info" {}

output "os_name" {
  value = data.localos_info.os_info.name
}

output "os_arch" {
  value = data.localos_info.os_info.arch
}

output "os_is_windows" {
  value = data.localos_info.os_info.is_windows
}

output "os_path_var" {
  value = nonsensitive(data.localos_info.os_info.environment.PATH)
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Read-Only

- `arch` (String) OS Architecture, e.g. "amd64"
- `environment` (Map of String, Sensitive) Map of all environment variables, with key=variable name (case sensitive), value=variable value.
- `id` (String) Resource identifier
- `is_windows` (Boolean) Utility attribute to quickly determine windows/not windows. Other supported OS are assumed to follow POSIX semantics.
- `name` (String) OS Name, e.g. "linux"
