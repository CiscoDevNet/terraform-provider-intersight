---
subcategory: "tam"
layout: "intersight"
page_title: "Intersight: intersight_tam_security_advisory"
description: |-
  Intersight representation of a Cisco PSIRT (https://tools.cisco.com/security/center/publicationListing.x) advisory definition. It includes the description of the security advisory and a corresponding reference to the published advisory. It also includes the Intersight data sources needed to evaluate the applicability of this advisory for relevant Intersight managed objects. A PSIRT definition is evaluated against all managed object referenced using the included data sources. Only Cisco TAC and Intersight devops engineers have the ability to create PSIRT definitions in Intersight.
---

# Resource: intersight_tam_security_advisory
Intersight representation of a Cisco PSIRT (https://tools.cisco.com/security/center/publicationListing.x) advisory definition. It includes the description of the security advisory and a corresponding reference to the published advisory. It also includes the Intersight data sources needed to evaluate the applicability of this advisory for relevant Intersight managed objects. A PSIRT definition is evaluated against all managed object referenced using the included data sources. Only Cisco TAC and Intersight devops engineers have the ability to create PSIRT definitions in Intersight.
## Usage Example
### Resource Creation

```hcl
resource "intersight_tam_security_advisory" "tam_security_advisory1" {
  name           = "tam_security_advisories1"
  operation_type = "create"
  state          = "ready"
  severity = {
    object_type = "tam.SecurityAdvisoryDetails"
  }
  actions = {
    object_type = "tam.SecurityAdvisoryDetails"
    alert_type  = "psirt"
  }
  type   = "psirt"
  status = "final"
  organization {
    object_type = "organization.Organization"
    moid        = var.organization
  }
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
* `advisory_id`:(string) Cisco generated identifier for the published security advisory. 
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
* `base_score`:(float) CVSS version 3 base score for the security Advisory. 
* `create_time`:(string)(ReadOnly) The time when this managed object was created. 
* `cve_ids`:
                (Array of schema.TypeString) -
* `date_published`:(string) Date when the security advisory was first published by Cisco. 
* `date_updated`:(string) Date when the security advisory was last updated by Cisco. 
* `description`:(string) Brief description of the advisory details. 
* `domain_group_moid`:(string)(ReadOnly) The DomainGroup ID for this managed object. 
* `environmental_score`:(float) CVSS version 3 environmental score for the security Advisory. 
* `external_url`:(string) A link to an external URL describing security Advisory in more details. 
* `mod_time`:(string)(ReadOnly) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) A user defined name for the Intersight Advisory. 
* `organization`:(HashMap) - A reference to a organizationOrganization resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
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
* `recommendation`:(string) Recommended action to resolve the security advisory. 
* `severity`:(HashMap) - Severity level of the Intersight Advisory. 
This complex property has following sub-properties:
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `shared_scope`:(string)(ReadOnly) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `state`:(string) Current state of the advisory.* `ready` - Advisory has been evaluated. The affected devices would be analyzed and corresponding advisory instances would be created.* `evaluating` - Advisory is currently under evaluation. The affected devices would be analyzed but no advisory instances wouldbe created. The results of the analysis would be made available to Intersight engineering for evaluation and validation. 
* `status`:(string) Cisco assigned status of the published advisory based on whether the investigation is complete or on-going.* `interim` - The Cisco investigation for the advisory is ongoing. Cisco will issue revisions to the advisory when additional information, including fixed software release data, becomes available.* `final` - Cisco has completed its evaluation of the vulnerability described in the advisory. There will be no further updates unless there is a material change in the nature of the vulnerability. 
* `tags`:(Array)
This complex property has following sub-properties:
  + `key`:(string) The string representation of a tag key. 
  + `value`:(string) The string representation of a tag value. 
* `temporal_score`:(float) CVSS version 3 temporal score for the security Advisory. 
* `nr_version`:(string) Cisco assigned advisory version after latest revision. 
* `version_context`:(HashMap) -(ReadOnly) The versioning info for this managed object. 
This complex property has following sub-properties:
  + `interested_mos`:(Array)
This complex property has following sub-properties:
    + `moid`:(string) The Moid of the referenced REST resource. 
    + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
    + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
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
`intersight_tam_security_advisory` can be imported using the Moid of the object, e.g.
```
$ terraform import intersight_tam_security_advisory.example 1234567890987654321abcde
``` 