---
subcategory: "partnerintegration"
layout: "intersight"
page_title: "Intersight: intersight_partnerintegration_file"
description: |-
        A partner integration artifact which will be files containing development code uploaded by our partners to get the build for dc/microservice.

---

# Resource: intersight_partnerintegration_file
A partner integration artifact which will be files containing development code uploaded by our partners to get the build for dc/microservice.
## Argument Reference
The following arguments are supported:
* `account_moid`:(string)(ReadOnly) The Account ID for this managed object. 
* `ancestors`:(Array)(ReadOnly) An array of relationships to moBaseMo resources. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `catalog`:(HashMap) - A reference to a softwarerepositoryCatalog resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `create_time`:(string)(ReadOnly) The time when this managed object was created. 
* `description`:(string) User provided description about the file. Cisco provided description for image inventoried from a Cisco repository. 
* `domain_group_moid`:(string)(ReadOnly) The DomainGroup ID for this managed object. 
* `download_count`:(int)(ReadOnly) The number of times this file has been downloaded from the local repository. It is used by the repository monitoring process to determine the files that are to be evicted from the cache. 
* `feature_source`:(string)(ReadOnly) The name of the feature to which the uploaded file belongs.* `System` - This indicates system initiated file uploads.* `OpenAPIImport` - This indicates an OpenAPI file upload.* `PartnerIntegrationImport` - This indicates a Partner-Integration Appliance user file uploads. 
* `file_path`:(string)(ReadOnly) Path of the file being uploaded. 
* `file_type`:(string) Type of the file being uploaded.* `None` - Invalid file type for partnerIntegration appliance.* `Model` - Model file of Generic Device.* `Etl` - ETL file of Generic Device.* `Ui` - UI file of Generic Device.* `DeviceConnector` - Generic Device Connector file. 
* `import_action`:(string) The action to be performed on the imported file. If 'PreCache' is set, the image will be cached in Appliance. Applicable in Intersight appliance deployment. If 'Evict' is set, the cached file will be removed. Applicable in Intersight appliance deployment. If 'GeneratePreSignedUploadUrl' is set, generates pre signed URL (s) for the file to be imported into the repository. Applicable for local machine source. The URL (s) will be populated under LocalMachine file server. If 'CompleteImportProcess' is set, the ImportState is marked as 'Imported'. Applicable for local machine source. If 'Cancel' is set, the ImportState is marked as 'Failed'. Applicable for local machine source.* `None` - No action should be taken on the imported file.* `GeneratePreSignedUploadUrl` - Generate pre signed URL of file for importing into the repository.* `GeneratePreSignedDownloadUrl` - Generate pre signed URL of file in the repository to download.* `CompleteImportProcess` - Mark that the import process of the file into the repository is complete.* `MarkImportFailed` - Mark to indicate that the import process of the file into the repository failed.* `PreCache` - Cache the file into the Intersight Appliance.* `Cancel` - The cancel import process for the file into the repository.* `Extract` - The action to extract the file in the external repository.* `Evict` - Evict the cached file from the Intersight Appliance. 
* `import_state`:(string)(ReadOnly) The state  of this file in the repository or Appliance. The importState is updated during the import operation and as part of the repository monitoring process.* `ReadyForImport` - The image is ready to be imported into the repository.* `Importing` - The image is being imported into the repository.* `Imported` - The image has been extracted and imported into the repository.* `PendingExtraction` - Indicates that the image has been imported but not extracted in the repository.* `Extracting` - Indicates that the image is being extracted into the repository.* `Extracted` - Indicates that the image has been extracted into the repository.* `Failed` - The image import from an external source to the repository has failed.* `MetaOnly` - The image is present in an external repository.* `ReadyForCache` - The image is ready to be cached into the Intersight Appliance.* `Caching` - Indicates that the image is being cached into the Intersight Appliance or endpoint cache.* `Cached` - Indicates that the image has been cached into the Intersight Appliance or endpoint cache.* `CachingFailed` - Indicates that the image caching into the Intersight Appliance failed or endpoint cache.* `Corrupted` - Indicates that the image in the local repository (or endpoint cache) has been corrupted after it was cached.* `Evicted` - Indicates that the image has been evicted from the Intersight Appliance (or endpoint cache) to reclaim storage space.* `Invalid` - Indicates that the corresponding distributable MO has been removed from the backend. This can be due to unpublishing of an image. 
* `imported_time`:(string)(ReadOnly) The time at which this image or file was imported/cached into the repositry. if the 'ImportState' is 'Imported', the time at which this image or file was imported. if the 'ImportState' is 'Cached', the time at which this image or file was cached. 
* `last_access_time`:(string)(ReadOnly) The time at which this file was last downloaded from the local repository. It is used by the repository monitoring process to determine the files that are to be evicted from the cache. 
* `md5e_tag`:(string) The MD5 ETag for a file that is stored in Intersight repository or in the appliance cache. Warning - MD5 is currently broken and this will be migrated to SHA shortly. 
* `md5sum`:(string) The md5sum checksum of the file. This information is available for all Cisco distributed images and files imported to the local repository. 
* `mod_time`:(string)(ReadOnly) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) The name of the file. It is populated as part of the image import operation. 
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
* `release_date`:(string)(ReadOnly) The date on which the file was released or distributed by its vendor. 
* `sha512sum`:(string) The sha512sum of the file. This information is available for all Cisco distributed images and files imported to the local repository. 
* `shared_scope`:(string)(ReadOnly) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `size`:(int) The size (in bytes) of the file. This information is available for all Cisco distributed images and files imported to the local repository. 
* `software_advisory_url`:(string) The software advisory, if any, provided by the vendor for this file. 
* `nr_source`:(HashMap) - Location of the file in an external repository. 
This complex property has following sub-properties:
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
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
* `nr_version`:(string) Vendor provided version for the file. 
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
* `workspace_name`:(string)(ReadOnly) The partner integration workspace to use to upload the File. 


## Import
`intersight_partnerintegration_file` can be imported using the Moid of the object, e.g.
```
$ terraform import intersight_partnerintegration_file.example 1234567890987654321abcde
``` 
