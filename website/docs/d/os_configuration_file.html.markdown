---
subcategory: "os"
layout: "intersight"
page_title: "Intersight: intersight_os_configuration_file"
description: |-
        ConfigurationFiles are operating system “answer files” used to enable unattended OS installation (for example, kickstart, unattended XML, or cloud-init/seed-style files). A ConfigurationFile can be either a static file or a template containing placeholders that are resolved at install time.
        #### Purpose
        Provide reusable, centrally managed OS installation configuration content—either as fixed answers or parameterized templates—so OS installs can be automated consistently across servers.
        #### Key Concepts
        - **Unattended install enablement**: Stores OS-specific answer content used by installation workflows.
        - **Template support with placeholders**: `fileContent` may use Go `text/template` syntax; `placeholders` exposes detected placeholder names so UIs/workflows can collect required inputs dynamically.
        - **Pre-canned vs user-defined entries**: `supported` distinguishes Intersight-provided, QA-validated templates from user-created ones; `internal` can mark wizard-uploaded files that should not appear in general management views.
        - **OS applicability mapping**: `distributions` links the configuration file to one or more supported OS distributions (`hcl.OperatingSystem`).
        - **Licensed capability**: CRUD operations require the Advantage entitlement.
        - **Stable identity**: Identified by `name` (and `catalog` as part of identity in the model), enabling consistent referencing.

---

# Data Source: intersight_os_configuration_file
ConfigurationFiles are operating system “answer files” used to enable unattended OS installation (for example, kickstart, unattended XML, or cloud-init/seed-style files). A ConfigurationFile can be either a static file or a template containing placeholders that are resolved at install time.
#### Purpose
Provide reusable, centrally managed OS installation configuration content—either as fixed answers or parameterized templates—so OS installs can be automated consistently across servers.
#### Key Concepts
- **Unattended install enablement**: Stores OS-specific answer content used by installation workflows.
- **Template support with placeholders**: `fileContent` may use Go `text/template` syntax; `placeholders` exposes detected placeholder names so UIs/workflows can collect required inputs dynamically.
- **Pre-canned vs user-defined entries**: `supported` distinguishes Intersight-provided, QA-validated templates from user-created ones; `internal` can mark wizard-uploaded files that should not appear in general management views.
- **OS applicability mapping**: `distributions` links the configuration file to one or more supported OS distributions (`hcl.OperatingSystem`).
- **Licensed capability**: CRUD operations require the Advantage entitlement.
- **Stable identity**: Identified by `name` (and `catalog` as part of identity in the model), enabling consistent referencing.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_os_configuration_file.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `description`:(string) Description of the OS ConfigurationFile. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `file_content`:(string) The content of the entire configuration file is stored as value. The contentcan either be a static file content or a template content.The template is expected to conform to the golang template syntax. The valuesfrom os.Answers properties will be used to populate this template. 
* `internal`:(bool) The internal flag is set to true when configuration file is uploaded from OS Install wizard. Internal Configuration files will not be displayed in Answer Management Page. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) The name of the OS ConfigurationFile that uniquely identifies the configuration file. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `supported`:(bool) An internal property that is used to distinguish between the pre-canned OSconfiguration file entries and user provided entries. 
 
