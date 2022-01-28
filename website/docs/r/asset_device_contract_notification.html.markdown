---
subcategory: "asset"
layout: "intersight"
page_title: "Intersight: intersight_asset_device_contract_notification"
description: |-
  Appliances use this object to send request to cloud to compute contract status. Cloud uses this to send the contract status response back to appliance.
---

# Resource: intersight_asset_device_contract_notification
Appliances use this object to send request to cloud to compute contract status. Cloud uses this to send the contract status response back to appliance.
## Argument Reference
The following arguments are supported:
* `account_moid`:(string)(ReadOnly) The Account ID for this managed object. 
* `ancestors`:(Array)(ReadOnly) An array of relationships to moBaseMo resources. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `contract`:(HashMap) - Contract information for the Cisco support contract purchased for the Cisco device. 
This complex property has following sub-properties:
  + `bill_to`:(HashMap) -(ReadOnly) BillTo address of listed for the contract. 
This complex property has following sub-properties:
    + `address1`:(string)(ReadOnly) Address Line one of the address information. example \ PO BOX 641570\ . 
    + `address2`:(string)(ReadOnly) Address Line two of the address information. example \ Cisco Systems\ . 
    + `address3`:(string)(ReadOnly) Address Line three of the address information. example \ Cisco Systems\ . 
    + `city`:(string)(ReadOnly) City in which the address resides. example \ San Jose\ . 
    + `country`:(string)(ReadOnly) Country in which the address resides. example \ US\ . 
    + `county`:(string)(ReadOnly) County in which the address resides. example \ Washington County\ . 
    + `location`:(string)(ReadOnly) Location in which the address resides. example \ 14852\ . 
    + `name`:(string)(ReadOnly) Name of the user whose address is being populated. 
    + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
    + `postal_code`:(string)(ReadOnly) Postal Code in which the address resides. example \ 95164-1570\ . 
    + `province`:(string)(ReadOnly) Province in which the address resides. example \ AB\ . 
    + `state`:(string)(ReadOnly) State in which the address resides. example \ CA\ . 
  + `bill_to_global_ultimate`:(HashMap) -(ReadOnly) BillToGlobalUltimate information listed in the contract. 
This complex property has following sub-properties:
    + `id`:(string)(ReadOnly) ID of the user in BillToGlobal. 
    + `name`:(string)(ReadOnly) Name of the user in BillToGlobal. 
    + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `contract_number`:(string)(ReadOnly) Contract number for the Cisco support contract purchased for the Cisco device. 
  + `line_status`:(string)(ReadOnly) Contract status as per the Cisco Contract APIx. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
* `contract_status`:(string) Calculated contract status that is derived based on the service line status and contract end date. It is different from serviceLineStatus property. serviceLineStatus gives us ACTIVE, OVERDUE, EXPIRED. These are transformed into Active, Expiring Soon and Not Covered.* `Unknown` - The device's contract status cannot be determined.* `Not Covered` - The Cisco device does not have a valid support contract.* `Active` - The Cisco device is covered under a active support contract.* `Expiring Soon` - The contract for this Cisco device is going to expire in the next 30 days. 
* `contract_status_reason`:(string) Reason for contract status. In case of Not Covered, reason is either Terminated or Expired.* `` - There is no reason for the specified contract status.* `Line Item Expired` - The Cisco device does not have a valid support contract, it has expired.* `Line Item Terminated` - The Cisco device does not have a valid support contract, it has been terminated. 
* `contract_updated_time`:(string) Date and time indicating when the contract data is last refreshed. 
* `covered_product_line_end_date`:(string) End date of the covered product line. The coverage end date is fetched from Cisco SN2INFO API. 
* `create_time`:(string)(ReadOnly) The time when this managed object was created. 
* `device_id`:(string) Unique identifier of the Cisco device. 
* `domain_group_moid`:(string)(ReadOnly) The DomainGroup ID for this managed object. 
* `end_customer`:(HashMap) - End customer information for the Cisco support contract purchased for the Cisco device. 
This complex property has following sub-properties:
  + `address`:(HashMap) -(ReadOnly) Address as per the information provided by the user. 
This complex property has following sub-properties:
    + `address1`:(string)(ReadOnly) Address Line one of the address information. example \ PO BOX 641570\ . 
    + `address2`:(string)(ReadOnly) Address Line two of the address information. example \ Cisco Systems\ . 
    + `address3`:(string)(ReadOnly) Address Line three of the address information. example \ Cisco Systems\ . 
    + `city`:(string)(ReadOnly) City in which the address resides. example \ San Jose\ . 
    + `country`:(string)(ReadOnly) Country in which the address resides. example \ US\ . 
    + `county`:(string)(ReadOnly) County in which the address resides. example \ Washington County\ . 
    + `location`:(string)(ReadOnly) Location in which the address resides. example \ 14852\ . 
    + `name`:(string)(ReadOnly) Name of the user whose address is being populated. 
    + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
    + `postal_code`:(string)(ReadOnly) Postal Code in which the address resides. example \ 95164-1570\ . 
    + `province`:(string)(ReadOnly) Province in which the address resides. example \ AB\ . 
    + `state`:(string)(ReadOnly) State in which the address resides. example \ CA\ . 
  + `id`:(string)(ReadOnly) Unique identifier for an end customer. This identifier is allocated by Cisco. 
  + `name`:(string)(ReadOnly) Name as per the information provided by the user. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
* `end_user_global_ultimate`:(HashMap) - EndUserGlobalUltimate information listed in the contract. 
This complex property has following sub-properties:
  + `id`:(string)(ReadOnly) ID of the user in BillToGlobal. 
  + `name`:(string)(ReadOnly) Name of the user in BillToGlobal. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
* `is_valid`:(bool) Validates if the device is a genuine Cisco device. Validated is done using the Cisco SN2INFO APIs. 
* `item_type`:(string) Item type of this specific Cisco device. example \ Chassis\ . 
* `maintenance_purchase_order_number`:(string) Maintenance purchase order number for the Cisco device. 
* `maintenance_sales_order_number`:(string) Maintenance sales order number for the Cisco device. 
* `mod_time`:(string)(ReadOnly) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
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
* `product`:(HashMap) - Product information of the offering under a contract. 
This complex property has following sub-properties:
  + `bill_to`:(HashMap) -(ReadOnly) Billing address provided by customer while buying this Cisco product. 
This complex property has following sub-properties:
    + `address1`:(string)(ReadOnly) Address Line one of the address information. example \ PO BOX 641570\ . 
    + `address2`:(string)(ReadOnly) Address Line two of the address information. example \ Cisco Systems\ . 
    + `address3`:(string)(ReadOnly) Address Line three of the address information. example \ Cisco Systems\ . 
    + `city`:(string)(ReadOnly) City in which the address resides. example \ San Jose\ . 
    + `country`:(string)(ReadOnly) Country in which the address resides. example \ US\ . 
    + `county`:(string)(ReadOnly) County in which the address resides. example \ Washington County\ . 
    + `location`:(string)(ReadOnly) Location in which the address resides. example \ 14852\ . 
    + `name`:(string)(ReadOnly) Name of the user whose address is being populated. 
    + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
    + `postal_code`:(string)(ReadOnly) Postal Code in which the address resides. example \ 95164-1570\ . 
    + `province`:(string)(ReadOnly) Province in which the address resides. example \ AB\ . 
    + `state`:(string)(ReadOnly) State in which the address resides. example \ CA\ . 
  + `description`:(string)(ReadOnly) Short description of the Cisco product that helps identify the product easily. example \ DISTI:UCS 6248UP 1RU Fabric Int/No PSU/32 UP/ 12p LIC\ . 
  + `family`:(string)(ReadOnly) Family that the product belongs to. Example \ UCSB\ . 
  + `group`:(string)(ReadOnly) Group that the product belongs to. It is one higher level categorization above family. example \ Switch\ . 
  + `number`:(string)(ReadOnly) Product number that identifies the product. example PID. example \ UCS-FI-6248UP-CH2\ . 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `ship_to`:(HashMap) -(ReadOnly) Shipping address provided by customer while buying this Cisco product. 
This complex property has following sub-properties:
    + `address1`:(string)(ReadOnly) Address Line one of the address information. example \ PO BOX 641570\ . 
    + `address2`:(string)(ReadOnly) Address Line two of the address information. example \ Cisco Systems\ . 
    + `address3`:(string)(ReadOnly) Address Line three of the address information. example \ Cisco Systems\ . 
    + `city`:(string)(ReadOnly) City in which the address resides. example \ San Jose\ . 
    + `country`:(string)(ReadOnly) Country in which the address resides. example \ US\ . 
    + `county`:(string)(ReadOnly) County in which the address resides. example \ Washington County\ . 
    + `location`:(string)(ReadOnly) Location in which the address resides. example \ 14852\ . 
    + `name`:(string)(ReadOnly) Name of the user whose address is being populated. 
    + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
    + `postal_code`:(string)(ReadOnly) Postal Code in which the address resides. example \ 95164-1570\ . 
    + `province`:(string)(ReadOnly) Province in which the address resides. example \ AB\ . 
    + `state`:(string)(ReadOnly) State in which the address resides. example \ CA\ . 
  + `sub_type`:(string)(ReadOnly) Sub type of the product being specified. example \ UCS 6200 SER\ . 
* `purchase_order_number`:(string) Purchase order number for the Cisco device. It is a unique number assigned for every purchase. 
* `registered_device`:(HashMap) -(ReadOnly) A reference to a assetDeviceRegistration resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `reseller_global_ultimate`:(HashMap) - ResellerGlobalUltimate information listed in the contract. 
This complex property has following sub-properties:
  + `id`:(string)(ReadOnly) ID of the user in BillToGlobal. 
  + `name`:(string)(ReadOnly) Name of the user in BillToGlobal. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
* `sales_order_number`:(string) Sales order number for the Cisco device. It is a unique number assigned for every sale. 
* `service_description`:(string) The type of service contract that covers the Cisco device. 
* `service_end_date`:(string) End date for the Cisco service contract that covers this Cisco device. 
* `service_level`:(string) The type of service contract that covers the Cisco device. 
* `service_sku`:(string) The SKU of the service contract that covers the Cisco device. 
* `service_start_date`:(string) Start date for the Cisco service contract that covers this Cisco device. 
* `shared_scope`:(string)(ReadOnly) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `state_contract`:(string) Internal property used for triggering and tracking actions for contract information.* `Update` - Sn2Info/Contract information needs to be updated.* `OK` - Sn2Info/Contract information was fetched succcessfuly and updated.* `Failed` - Sn2Info/Contract information was not available  or failed while fetching.* `Retry` - Sn2Info/Contract information update failed and will be retried later. 
* `state_sn2_info`:(string) Internal property used for triggering and tracking actions for sn2info information.* `Update` - Sn2Info/Contract information needs to be updated.* `OK` - Sn2Info/Contract information was fetched succcessfuly and updated.* `Failed` - Sn2Info/Contract information was not available  or failed while fetching.* `Retry` - Sn2Info/Contract information update failed and will be retried later. 
* `tags`:(Array)
This complex property has following sub-properties:
  + `key`:(string) The string representation of a tag key. 
  + `value`:(string) The string representation of a tag value. 
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
* `warranty_end_date`:(string) End date for the warranty that covers the Cisco device. 
* `warranty_type`:(string) Type of warranty that covers the Cisco device. 


## Import
`intersight_asset_device_contract_notification` can be imported using the Moid of the object, e.g.
```
$ terraform import intersight_asset_device_contract_notification.example 1234567890987654321abcde
``` 