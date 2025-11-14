---
subcategory: "os"
layout: "intersight"
page_title: "Intersight: intersight_os_bulk_install_info"
description: |-
        This MO models the CSV file content which the user uploaded for OS installation. As part of the handler, necessary filed
        in the model can be populated along with respective validation.

---

# Resource: intersight_os_bulk_install_info
This MO models the CSV file content which the user uploaded for OS installation. As part of the handler, necessary filed
in the model can be populated along with respective validation.
## Argument Reference
The following arguments are supported:
* `account_moid`:(string)(ReadOnly) The Account ID for this managed object. 
* `ancestors`:(Array)(ReadOnly) An array of relationships to moBaseMo resources. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `configuration_file`:(HashMap) - A reference to a osConfigurationFile resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `create_time`:(string)(ReadOnly) The time when this managed object was created. 
* `domain_group_moid`:(string)(ReadOnly) The DomainGroup ID for this managed object. 
* `file_content`:(string) The content of the entire CSV file is stored as value. The content can hold complete OS install parameters in two sections.The first section holds generic information about the OS Install like OS Image, SCU Image etc. The second section holdsparameters which are specific to each server level data. 
* `global_config`:(HashMap) -(ReadOnly) The global parameter of the CSV file. 
This complex property has following sub-properties:
  + `configuration_file_name`:(string)(ReadOnly) Name of the Configuration file. 
  + `configuration_source`:(string)(ReadOnly) Configuration source for the OS Installation. 
  + `install_method`:(string)(ReadOnly) The install method to be used for OS installation - vMedia, iPXE.Only vMedia is supported as of now. 
  + `install_target_type`:(string)(ReadOnly) The Prefill install Target Name. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `operating_system_parameters`:(HashMap) -(ReadOnly) Parameters specific to selected OS. 
This complex property has following sub-properties:
    + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
  + `os_image_name`:(string)(ReadOnly) The Operating System Image name. 
  + `override_secure_boot`:(bool) ESXi Secure Boot installation is currently not supported. As a workaround, Secure Boot will be disabled before installation and restored after installation is complete. Enable to Override Secure Boot Configuration. 
  + `scu_image_name`:(string)(ReadOnly) The name of the Server Configuration Utilities Image. 
  + `windows_edition`:(string)(ReadOnly) The Windows OS edition, this property required only for Windows server. 
* `is_file_content_set`:(bool)(ReadOnly) Indicates whether the value of the 'fileContent' property has been set. 
* `mod_time`:(string)(ReadOnly) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) The name of the CSV file, which holds the OS install parameters. 
* `oper_state`:(string)(ReadOnly) Denotes if the operating is pending, in_progress, completed_ok, completed_error.* `Pending` - The initial value of the OperStatus.* `InProgress` - The OperStatus value will be InProgress during execution.* `CompletedOk` - The API is successful with operation then OperStatus will be marked as CompletedOk.* `CompletedError` - The API is failed with operation then OperStatus will be marked as CompletedError.* `CompletedWarning` - The API is completed with some warning then OperStatus will be CompletedWarning. 
* `organization`:(HashMap) - A reference to a organizationOrganization resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `os_image`:(HashMap) - A reference to a softwarerepositoryOperatingSystemFile resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `owners`:
                (Array of schema.TypeString) -(ReadOnly)
* `parent`:(HashMap) -(ReadOnly) A reference to a moBaseMo resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `permission_resources`:(Array)(ReadOnly) An array of relationships to moBaseMo resources. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `scu_image`:(HashMap) - A reference to a firmwareServerConfigurationUtilityDistributable resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `server_configs`:(Array)
This complex property has following sub-properties:
  + `additional_parameters`:(Array)
This complex property has following sub-properties:
    + `is_value_set`:(bool) Flag to indicate if value is set. Value will be used to check if any edit. 
    + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
    + `type`:(JSON as string) Definition of place holder. 
    + `value`:(JSON as string) Value for placeholder provided by user. 
  + `answers`:(HashMap) -(ReadOnly) Answers provided by user for the unattended OS installation. 
This complex property has following sub-properties:
    + `alternate_name_servers`:
                (Array of schema.TypeString) -
    + `answer_file`:(string) If the source of the answers is a static file, the content of the file is stored as valuein this property.The value is mandatory only when the 'Source' property has been set to 'File'. 
    + `hostname`:(string) Hostname to be configured for the server in the OS. 
    + `ip_config_type`:(string) IP configuration type. Values are Static or Dynamic configuration of IP.In case of static IP configuration, IP address, gateway and other details needto be populated. In case of dynamic the IP configuration is obtained dynamicallyfrom DHCP.* `static` - In case of static IP configuraton, provide the details such as IP address, netmask, and gateway.* `DHCP` - In case of dynamic IP configuration, the IP address, netmask and gateway detailsare obtained from DHCP. 
    + `ip_configuration`:(HashMap) - In case of static IP configuration, IP address, netmask and gateway details areprovided. 
This complex property has following sub-properties:
    + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
  + `is_answer_file_set`:(bool)(ReadOnly) Indicates whether the value of the 'answerFile' property has been set. 
  + `is_root_password_crypted`:(bool) Enable to indicate Root Password provided is encrypted. 
  + `is_root_password_set`:(bool)(ReadOnly) Indicates whether the value of the 'rootPassword' property has been set. 
  + `nameserver`:(string) IP address of the name server to be configured in the OS. 
  + `network_device`:(string) Network Device where the IP address must be configured. Network Interface names and MAC address are supported.For SUSE Linux Enterprise Server, Network Interface name is a required input and if provided as a MAC address,A persistent interface name is binded to the MAC address and the interface name will be used for network configuration.Refer https://documentation.suse.com/sles/15-SP2/html/SLES-all/cha-configuration-installation-options.html#CreateProfile-Network-names. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `product_key`:(string) The product key to be used for a specific version of Windows installation. 
  + `root_password`:(string) Password configured for the root / administrator user in the OS. You can enter a plain text or an encrypted password.Intersight encrypts the plaintext password. Enable the Encrypted Password option to provide an encrypted password.For more details on encrypting passwords, see Help Center. 
  + `nr_source`:(string) Answer values can be provided from three sources - Embedded in OS image, static file,or as placeholder values for an answer file template.Source of the answers is given as value, Embedded/File/Template.'Embedded' option indicates that the answer file is embedded within the OS Image. 'File'option indicates that the answers are provided as a file. 'Template' indicates that theplaceholders in the selected os.ConfigurationFile MO are replaced with values providedas os.Answers MO.* `None` - Indicates that answers is not sent and values must be populated from Install Template.  * `Embedded` - Indicates that the answer file is embedded within OS image.* `File` - Indicates that the answer file is a static content that has all thevalues populated.* `Template` - Indicates that the given answers are used to populate the answer filetemplate. The template allows the users to refer some server specificanswers as fields/placeholders and replace these placeholders with theactual values for each Server during OS installation using 'Answers' and'AdditionalParameters' properties in os.Install MO.The answer file templates can be created by users as os.ConfigurationFile objects. 
  + `error_msgs`:
                (Array of schema.TypeString) -
  + `initiator_wwpn`:(string)(ReadOnly) The WWPN Address of the underlying fibre channel interface at the host side used for SAN accesss. Value must be in hexadecimal format xx:xx:xx:xx:xx:xx:xx:xx.  For example, 20:00:D4:C9:3C:35:02:01. 
  + `install_target`:(string)(ReadOnly) The target in which OS installation triggered, the value represented is StorageControllerID follwed by PhysicalDisk SerialNumber in case of Physcial disk or VirtualDriveId for virtual drive. 
  + `lun_id`:(int)(ReadOnly) The Logical Unit Number (LUN) of the install target. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `operating_system_parameters`:(HashMap) -(ReadOnly) Parameters specific to selected OS. 
This complex property has following sub-properties:
    + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
  + `processed_install_target`:(HashMap) -(ReadOnly) The target in which OS installation triggered, this is populated after processing the given data. 
This complex property has following sub-properties:
    + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
  + `serial_number`:(string)(ReadOnly) The Serial Number of the server. 
  + `target_iqn`:(string)(ReadOnly) IQN (iSCSI qualified name) of Storage iSCSI target. Can be up to 255 characters long and has the following format, iqn.yyyy-mm.naming-authority:unique_name. 
  + `target_wwpn`:(string)(ReadOnly) The WWPN Address of the underlying fibre channel interface at the target used by the storage. Value must be in hexadecimal format xx:xx:xx:xx:xx:xx:xx:xx.  For example, 51:4F:0C:50:14:1F:AF:01. 
  + `vnic_mac`:(string)(ReadOnly) MAC address of the VNIC used as iSCSI initiator interface. 
* `servers`:(Array) An array of relationships to computePhysical resources. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `shared_scope`:(string)(ReadOnly) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `tags`:(Array)
This complex property has following sub-properties:
  + `ancestor_definitions`:(Array)
This complex property has following sub-properties:
    + `moid`:(string) The Moid of the referenced REST resource. 
    + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
    + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
  + `definition`:(HashMap) -(ReadOnly) The definition is a reference to the tag definition object.The tag definition object contains the properties of the tag such as name, type, and description. 
This complex property has following sub-properties:
    + `moid`:(string) The Moid of the referenced REST resource. 
    + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
    + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
  + `key`:(string) The string representation of a tag key. 
  + `propagated`:(bool)(ReadOnly) Propagated is a boolean flag that indicates whether the tag is propagated to the related managed objects. 
  + `sys_tag`:(bool)(ReadOnly) Specifies whether the tag is user-defined or owned by the system. 
  + `type`:(string)(ReadOnly) An enum type that defines the type of tag. Supported values are 'pathtag' and 'keyvalue'.* `KeyValue` - KeyValue type of tag. Key is required for these tags. Value is optional.* `PathTag` - Key contain path information. Value is not present for these tags. The path is created by using the '/' character as a delimiter.For example, if the tag is \ A/B/C\ , then \ A\  is the parent tag, \ B\  is the child tag of \ A\  and \ C\  is the child tag of \ B\ . 
  + `value`:(string) The string representation of a tag value. 
* `validation_infos`:(Array)
This complex property has following sub-properties:
  + `error_msg`:(string)(ReadOnly) Validation error message. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `status`:(string)(ReadOnly) The status of the validation step.* `NotValidated` - The validation not started.* `Valid` - The step status marked as valid when respective step validation condition is successful.* `Invalid` - The step status marked as invalid when respective step validation condition is failed.* `InProgress` - The validation is in progress. 
  + `step_name`:(string)(ReadOnly) The validation step name.* `OS Install Schema Validation` - The step to validate the CSV file schema.* `OS Image Validation` - The Operating System Image parameter validation step.* `SCU Image Validation` - The SCU Image parameter validation step.* `Configuration source and file validation` - The Configuration Source and Configuration file validation step.* `Server level data validation` - The server level parameters validation. 
* `version_context`:(HashMap) -(ReadOnly) The versioning info for this managed object. 
This complex property has following sub-properties:
  + `interested_mos`:(Array)
This complex property has following sub-properties:
    + `moid`:(string) The Moid of the referenced REST resource. 
    + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
    + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
  + `marked_for_deletion`:(bool)(ReadOnly) The flag to indicate if snapshot is marked for deletion or not. If flag is set then snapshot will be removed after the successful deployment of the policy. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `ref_mo`:(HashMap) -(ReadOnly) A reference to the original Managed Object. 
This complex property has following sub-properties:
    + `moid`:(string) The Moid of the referenced REST resource. 
    + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
    + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
  + `timestamp`:(string)(ReadOnly) The time this versioned Managed Object was created. 
  + `nr_version`:(string)(ReadOnly) The version of the Managed Object, e.g. an incrementing number or a hash id. 
  + `version_type`:(string)(ReadOnly) Specifies type of version. Currently the only supported value is \ Configured\ that is used to keep track of snapshots of policies and profiles that are intendedto be configured to target endpoints.* `Modified` - Version created every time an object is modified.* `Configured` - Version created every time an object is configured to the service profile.* `Deployed` - Version created for objects related to a service profile when it is deployed. 


## Import
`intersight_os_bulk_install_info` can be imported using the Moid of the object, e.g.
```
$ terraform import intersight_os_bulk_install_info.example 1234567890987654321abcde
``` 
