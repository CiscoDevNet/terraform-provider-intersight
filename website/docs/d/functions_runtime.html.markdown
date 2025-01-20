---
subcategory: "functions"
layout: "intersight"
page_title: "Intersight: intersight_functions_runtime"
description: |-
        The managed object which has info about language runtime.

---

# Data Source: intersight_functions_runtime
The managed object which has info about language runtime.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_functions_runtime.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `action`:(string) Action against the Runtime.* `None` - No action is set, this is the default value for action field.* `Disable` - Disable an instance of a Runtime.* `Deprecate` - Deprecate an instance of a Runtime.* `FlagInsecure` - Flag an instance of a Runtime as insecure. 
* `code_file_name`:(string) Name of file containing function source code. 
* `code_template`:(string) Template to guide on how to compose code. 
* `create_time`:(string) The time when this managed object was created. 
* `create_user`:(string) The user identifier who created the language runtime. 
* `deprecated`:(bool) Indicate if this language runtime is deprecated. 
* `description`:(string) Description of the language runtime. 
* `display_name`:(string) The display name of the language runtime. Display name can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-), period (.), colon (:), space ( ) or an underscore (_). 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `enabled`:(bool) Indicate if this language runtime is enabled. If not, the runtime is not usable at all. 
* `insecure`:(bool) Indicate if this language runtime is insecure due to vulnerabilities. 
* `language_name`:(string) The official name of the programming language.* `Python` - Python programming language. 
* `language_version`:(string) The version of the programming language. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `mod_user`:(string) The user identifier who last updated the language runtime. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) The name of the language runtime. Name can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-), period (.) or an underscore (_). 
* `note`:(string) A note to bring user's attention to the status of this language runtime. 
* `runtime_file_path`:(string) Path to the runtime file. 
* `runtime_upload_moid`:(string) Moid of Upload which represents the uploaded runtime file. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
