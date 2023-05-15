---
subcategory: "tam"
layout: "intersight"
page_title: "Intersight: intersight_tam_advisory_definition"
description: |-
        An Intersight Advisory. An advisory represents an identification of a potential issue and may also include  a recommendation for resolving the said issue. Advisories may be of different kind and severity. for e.g. It could be a security vulnerability or a performance issue or a hardware issue with different recommendations for resolving them.

---

# Resource: intersight_tam_advisory_definition
An Intersight Advisory. An advisory represents an identification of a potential issue and may also include  a recommendation for resolving the said issue. Advisories may be of different kind and severity. for e.g. It could be a security vulnerability or a performance issue or a hardware issue with different recommendations for resolving them.
## Usage Example
### Resource Creation

```hcl
resource "intersight_tam_advisory_definition" "tam_advisory_definition" {
  name  = "tam_advisory_definition"
  state = "ready"
  severity {
    object_type = "tam.SecurityAdvisoryDetails"
  }
  actions {
    operation_type = "create"
    name           = "tam_security_advisories1"
    type           = "restApi"
    object_type    = "tam.SecurityAdvisoryDetails"
    alert_type     = "psirt"
  }
  type = "securityAdvisory"
  api_data_sources {
    object_type = "tam.ApiDataSource"
    name        = "api_data_source_1"
    type        = "intersightApi"

  }
  advisory_details {
    object_type = "tam.SecurityAdvisoryDetails"
    description = "tam security advisory"
    class_id    = "tam.SecurityAdvisoryDetails"
  }

  organization {
    object_type = "organization.Organization"
    moid        = var.organization
  }
}

variable "organization" {
  type        = string
  description = "<value for organization>"
}
```
## Argument Reference
The following arguments are supported:
* `account_moid`:(string)(ReadOnly) The Account ID for this managed object. 
* `actions`:(Array)
This complex property has following sub-properties:
  + `affected_object_type`:(string) Type of the managed object that should be marked with an instance of the Alert (when operation type is create) or that should have an alert instance removed (when operation type is remove). 
  + `alert_type`:(string) Alert type is used to denote the category of an Intersight alert (FieldNotice, equipment Fault etc.).* `psirt` - Respresents the psirt alert type (https://tools.cisco.com/security/center/publicationListing.x).* `fieldNotice` - Respresents the field notice alert type (https://www.cisco.com/c/en/us/support/web/tsd-products-field-notice-summary.html).* `eolAdvisory` - Represents product End of Life (EOL) type (https://www.cisco.com/c/en/us/products/eos-eol-policy.html). 
  + `identifiers`:(Array)
This complex property has following sub-properties:
    + `name`:(string) Name of the filter paramter. 
    + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
    + `value`:(string) Value of the filter paramter. 
  + `name`:(string) Uniquely identifies a given action among the set of actions corresponding to an advisory. Primarily used to store and compare results of subsequent iterations corresponding to the action queries. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `operation_type`:(string) Operation type for the alert action. An action is used to carry out the process of \ reacting\  to an alert condition. For e.g.in case of a fieldNotice alert, the intention may be to create a new alert (if the condition matches and there is no existing alert) or to remove an existing alert when the alert condition has been remedied.* `create` - Create an instance of AdvisoryInstance.* `remove` - Remove an instance of AdvisoryInstance. 
  + `queries`:(Array)
This complex property has following sub-properties:
    + `name`:(string) Name is used to unique identify and result of the given query which can be used by subsequent queries as input data source. 
    + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
    + `priority`:(int) An integer value depicting the priority of the query among the queries that are part of the same QueryEntry collection. 
    + `query`:(string) A SparkSQL query to be used on a given data source. 
  + `type`:(string) Type of Intersight alert. An alert in Intersight could be one of several kinds (FieldNotice, PSIRT etc.). Primarily used for filtering alerts based on the type.* `restApi` - Repesents the use of REST API for carrying out alert actions. 
* `advisory_details`:(HashMap) - Additional details for the advisory definition. For e.g. if the definition corresponds to a security advisory, the details regarding CVE ids and CVSS score would be available here. If the definition is for an end-of-life milestone, the details about the specific milestone will be included. 
This complex property has following sub-properties:
  + `additional_properties`:(JSON as string) - Additional Properties as per object type, can be added as JSON using `jsonencode()`. Allowed Types are: [tam.EolAdvisoryDetails](#tamEolAdvisoryDetails)
[tam.SecurityAdvisoryDetails](#tamSecurityAdvisoryDetails)
  + `description`:(string) Brief description of details specified for an advisory type. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `advisory_id`:(string) Cisco generated identifier for the published security/field-notice/end-of-life advisory. 
* `ancestors`:(Array)(ReadOnly) An array of relationships to moBaseMo resources. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `api_data_sources`:(Array)
This complex property has following sub-properties:
  + `mo_type`:(string) Type of Intersight managed object used as data source. 
  + `name`:(string) Name is used to unique identify and refer a given data source in an alert definition. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `queries`:(Array)
This complex property has following sub-properties:
    + `name`:(string) Name is used to unique identify and result of the given query which can be used by subsequent queries as input data source. 
    + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
    + `priority`:(int) An integer value depicting the priority of the query among the queries that are part of the same QueryEntry collection. 
    + `query`:(string) A SparkSQL query to be used on a given data source. 
  + `type`:(string) Type of data source (for e.g. TextFsmTempalate based, Intersight API based etc.).* `intersightApi` - Collector type for this data collection is Intersight APIs.* `nxos` - Collector type for this data collection is NXOS.* `s3File` - Collector type for this data collection is a file in a cloud hosted object storage bucket. 
* `create_time`:(string)(ReadOnly) The time when this managed object was created. 
* `date_published`:(string) Date when the security/field-notice/end-of-life advisory was first published by Cisco. 
* `date_updated`:(string) Date when the security/field-notice/end-of-life advisory was last updated by Cisco. 
* `description`:(string) Brief description of the advisory details. 
* `domain_group_moid`:(string)(ReadOnly) The DomainGroup ID for this managed object. 
* `execute_on_pod`:(string) Orion pod on which this advisory should process.* `tier1` - Advisory processing will be taken care in first advisory driver of multinode cluster.* `tier2` - Advisory processing will be taken care in second advisory driver of multinode cluster. 
* `external_url`:(string) A link to an external URL describing security Advisory in more details. 
* `mod_time`:(string)(ReadOnly) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) A user defined name for the Intersight Advisory. 
* `organization`:(HashMap) - A reference to a organizationOrganization resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `other_ref_urls`:
                (Array of schema.TypeString) -
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
* `recommendation`:(string) Recommended action to resolve the security advisory. 
* `s3_data_sources`:(Array)
This complex property has following sub-properties:
  + `name`:(string) Name is used to unique identify and refer a given data source in an alert definition. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `queries`:(Array)
This complex property has following sub-properties:
    + `name`:(string) Name is used to unique identify and result of the given query which can be used by subsequent queries as input data source. 
    + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
    + `priority`:(int) An integer value depicting the priority of the query among the queries that are part of the same QueryEntry collection. 
    + `query`:(string) A SparkSQL query to be used on a given data source. 
  + `s3_path`:(string) Path used to access file in s3 containing data. 
  + `type`:(string) Type of data source (for e.g. TextFsmTempalate based, Intersight API based etc.).* `intersightApi` - Collector type for this data collection is Intersight APIs.* `nxos` - Collector type for this data collection is NXOS.* `s3File` - Collector type for this data collection is a file in a cloud hosted object storage bucket. 
* `severity`:(HashMap) - Severity level of the Intersight Advisory. 
This complex property has following sub-properties:
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `shared_scope`:(string)(ReadOnly) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `state`:(string) Current state of the advisory.* `ready` - Advisory has been evaluated. The affected devices would be analyzed and corresponding advisory instances would be created.* `evaluating` - Advisory is currently under evaluation. The affected devices would be analyzed but no advisory instances wouldbe created. The results of the analysis would be made available to Intersight engineering for evaluation and validation. 
* `tags`:(Array)
This complex property has following sub-properties:
  + `key`:(string) The string representation of a tag key. 
  + `value`:(string) The string representation of a tag value. 
* `type`:(string) The type (field notice, security advisory, end-of-life milestone advisory etc.) of Intersight advisory.* `securityAdvisory` - Respresents the psirt alert type (https://tools.cisco.com/security/center/publicationListing.x).* `fieldNotice` - Respresents the field notice alert type (https://www.cisco.com/c/en/us/support/web/tsd-products-field-notice-summary.html).* `eolAdvisory` - Represents product End of Life (EOL) type (https://www.cisco.com/c/en/us/products/eos-eol-policy.html). 
* `nr_version`:(string) Cisco assigned advisory/field-notice/end-of-life version after latest revision. 
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
* `workaround`:(string) Workarounds available for the advisory. 


## Import
`intersight_tam_advisory_definition` can be imported using the Moid of the object, e.g.
```
$ terraform import intersight_tam_advisory_definition.example 1234567890987654321abcde
```
## Allowed Types in `AdditionalProperties`
 
### [tam.EolAdvisoryDetails](#argument-reference)
Details pertaining to the milestone defined by an end-of-life (EOL) milestone advisory.
* `all_milestones`:(Array)
This complex property has following sub-properties:
  + `date`:(string) Date when the specified end-of-life milestone advisory is reached. 
  + `description`:(string) A description of the milestone defined by an end-of-life advisory. 
  + `end_offset`:(int) Number of days (exclusive) relative to the milestone date when the milestone is considered to be not in effect. A nagative number indicates number of days ahead of the milestone date. The default is 2147483647 (0x7FFFFFFF) which means the milestone date range's upper bound is omitted. 
  + `label_hint`:(string) Extra hint to the type of label to be used in display in addition to severity and effective date. How to use it is at UI's descretion.* `upcoming` - This end-of-life (EOL) milestone is upcoming. The label may be changed to more urgent type such as 'imminent' as time progress approaching effective date.* `imminent` - This end-of-life (EOL) milestone is imminent. There will be no label change before the effective date.* `past` - This end-of-life (EOL) milestone has past the effective date.* `info` - This end-of-life (EOL) milestone for informative purpose. 
  + `milestone_type`:(string) Milestone type as defined in Cisco end-of-life (EOL) policy (https://www.cisco.com/c/en/us/products/eos-eol-policy.html) when the specified end-of-life milestone advisory is reached.* `unknown` - The type of end-of-life milestone is not defined.* `endOfSoftwareMaintenanceDate` - The last date that Cisco Engineering may release any final software maintenance releases or bug fixes. After this date, Cisco Engineering may no longer develop, repair, maintain, or test the product software and only critical security updates will be provided on this release train. * `lastDateOfSupport` - The last date to receive service and support for the software. After this date, all support services for the software are unavailable, and the software becomes obsolete. * `endOfLifeAnnouncementDate` - The date the document that announces the end-of-sale and end-of-life of a product is distributed to the general public. 
  + `name`:(string) A milestone defined by an end-of-life advisory. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `start_offset`:(int) Number of days (inclusive) relative to the milestone date when the milestone is considered to be in effect. A nagative number indicates number of days ahead of the milestone date. The default is 0 which means the milestone take effect exactly on the same date as the specified milestone date. A negative value of -2147483648 (0x80000000) indicates the milestone date range's lower bound is omitted. 
* `milestone`:(HashMap) - A end-of-life advisory milestone as defined in Cisco end-of-life (EOL) policy (https://www.cisco.com/c/en/us/products/eos-eol-policy.html). 
This complex property has following sub-properties:
  + `date`:(string) Date when the specified end-of-life milestone advisory is reached. 
  + `description`:(string) A description of the milestone defined by an end-of-life advisory. 
  + `end_offset`:(int) Number of days (exclusive) relative to the milestone date when the milestone is considered to be not in effect. A nagative number indicates number of days ahead of the milestone date. The default is 2147483647 (0x7FFFFFFF) which means the milestone date range's upper bound is omitted. 
  + `label_hint`:(string) Extra hint to the type of label to be used in display in addition to severity and effective date. How to use it is at UI's descretion.* `upcoming` - This end-of-life (EOL) milestone is upcoming. The label may be changed to more urgent type such as 'imminent' as time progress approaching effective date.* `imminent` - This end-of-life (EOL) milestone is imminent. There will be no label change before the effective date.* `past` - This end-of-life (EOL) milestone has past the effective date.* `info` - This end-of-life (EOL) milestone for informative purpose. 
  + `milestone_type`:(string) Milestone type as defined in Cisco end-of-life (EOL) policy (https://www.cisco.com/c/en/us/products/eos-eol-policy.html) when the specified end-of-life milestone advisory is reached.* `unknown` - The type of end-of-life milestone is not defined.* `endOfSoftwareMaintenanceDate` - The last date that Cisco Engineering may release any final software maintenance releases or bug fixes. After this date, Cisco Engineering may no longer develop, repair, maintain, or test the product software and only critical security updates will be provided on this release train. * `lastDateOfSupport` - The last date to receive service and support for the software. After this date, all support services for the software are unavailable, and the software becomes obsolete. * `endOfLifeAnnouncementDate` - The date the document that announces the end-of-sale and end-of-life of a product is distributed to the general public. 
  + `name`:(string) A milestone defined by an end-of-life advisory. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `start_offset`:(int) Number of days (inclusive) relative to the milestone date when the milestone is considered to be in effect. A nagative number indicates number of days ahead of the milestone date. The default is 0 which means the milestone take effect exactly on the same date as the specified milestone date. A negative value of -2147483648 (0x80000000) indicates the milestone date range's lower bound is omitted. 
* `release`:(string) The name of the impacted release this advisory milestone intends to address, e.g. \ 3.5 (2x)\ . 

### [tam.SecurityAdvisoryDetails](#argument-reference)
Details pertaining to a security advisory defined by a given advisory.
* `base_score`:(float) CVSS version 3 base score for the security Advisory. 
* `cve_ids`:
                (Array of schema.TypeString) -
* `environmental_score`:(float) CVSS version 3 environmental score for the security Advisory. 
* `status`:(string) Cisco assigned status of the published advisory. Depends on whether the investigation is complete or on-going.* `interim` - The Cisco investigation for the advisory is ongoing. Cisco will issue revisions to the advisory when additional information, including fixed software release data, becomes available.* `final` - Cisco has completed its evaluation of the vulnerability described in the advisory. There will be no further updates unless there is a material change in the nature of the vulnerability. 
* `temporal_score`:(float) CVSS version 3 temporal score for the security Advisory. 
  
