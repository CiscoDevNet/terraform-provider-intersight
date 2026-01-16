---
subcategory: "server"
layout: "intersight"
page_title: "Intersight: intersight_server_profile"
description: |-
        The Profile object represents a server profile, specifying configuration settings for a physical server. It encapsulates all necessary information to deploy, manage, and monitor a server within the infrastructure.
        #### Purpose
        A Profile is used to apply detailed configuration and policy settings to an individual server, supporting lifecycle management, compliance, and automation.
        #### Key Concepts
        - **Configuration Management:** Centralizes all relevant configuration states, changes, and compliance results for the server.
        - **Policy Attachment:** Associates various policies and pools (such as UUID pools and resource pools) to define operational parameters.
        - **Assignment and Deployment:** Supports both static and dynamic assignment of physical servers, as well as monitoring deployment status and results.
        - **Lifecycle Tracking:** Integrates with workflow engines and maintains audit trails for configuration changes and applied policies.

---

# Resource: intersight_server_profile
The Profile object represents a server profile, specifying configuration settings for a physical server. It encapsulates all necessary information to deploy, manage, and monitor a server within the infrastructure.
#### Purpose
A Profile is used to apply detailed configuration and policy settings to an individual server, supporting lifecycle management, compliance, and automation.
#### Key Concepts
- **Configuration Management:** Centralizes all relevant configuration states, changes, and compliance results for the server.
- **Policy Attachment:** Associates various policies and pools (such as UUID pools and resource pools) to define operational parameters.
- **Assignment and Deployment:** Supports both static and dynamic assignment of physical servers, as well as monitoring deployment status and results.
- **Lifecycle Tracking:** Integrates with workflow engines and maintains audit trails for configuration changes and applied policies.
## Usage Example
### Resource Creation

```hcl
resource "intersight_server_profile" "server1" {
  name   = "server1"
  action = "No-op"
  tags {
    key   = "server"
    value = "demo"
  }
  organization {
    object_type = "organization.Organization"
    moid        = var.organization
  }
}

variable "organization" {
   type = string
   description = "<value for organization>"
 }
```
## Argument Reference
The following arguments are supported:
* `account_moid`:(string)(ReadOnly) The Account ID for this managed object. 
* `action`:(string) User initiated action. Each profile type has its own supported actions. For HyperFlex cluster profile, the supported actions are -- Validate, Deploy, Continue, Retry, Abort, Unassign For server profile, the support actions are -- Deploy, Unassign. 
* `action_params`:(Array)
This complex property has following sub-properties:
  + `name`:(string) The action parameter identifier. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `value`:(string) The action parameter value. 
* `ancestors`:(Array)(ReadOnly) An array of relationships to moBaseMo resources. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `assigned_server`:(HashMap) - A reference to a computePhysical resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `associated_server`:(HashMap) -(ReadOnly) A reference to a computePhysical resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `associated_server_pool`:(HashMap) - A reference to a resourcepoolPool resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `config_change_context`:(HashMap) -(ReadOnly) The configuration change state and results of the last change operation. 
This complex property has following sub-properties:
  + `config_change_error`:(string)(ReadOnly) Indicates reason for failure state of configChangeState. 
  + `config_change_state`:(string)(ReadOnly) Indicates a profile's configuration change state. Used for tracking pending-changes and out-of-synch states.* `Ok` - Config change state represents Validation for change/drift is successful or is not applicable.* `Validating-change` - Config change state represents policy changes are being validated. This state starts when policy is changed and becomes different from deployed changes (Pending-changes).* `Validating-drift` - Config change state represents endpoint configuration changes are being validated. This state starts when endpoint is changed and endpoint configuration becomes different from policy configured changes (Out-of-sync).* `Change-failed` - Config change state represents there is internal error in calculating policy change.* `Drift-failed` - Config change state represents there is internal error in calculating endpoint configuraion drift. 
  + `initial_config_context`:(HashMap) -(ReadOnly) Stores initial Configuration state. Used for reverting back to initial state of ConfigContext in case of validation failure. 
This complex property has following sub-properties:
    + `config_state`:(string)(ReadOnly) Indicates a profile's configuration deploying state. Values -- Assigned, Not-assigned, Associated, Pending-changes, Out-of-sync, Validating, Configuring, Failed. 
    + `config_state_summary`:(string)(ReadOnly) Indicates a profile's configuration deploying state. Values -- Assigned, Not-assigned, Associated, InConsistent, Validating, Configuring, Failed, Activating, UnConfiguring.* `None` - The default state is none.* `Not-assigned` - Server is not assigned to the profile.* `Assigned` - Server is assigned to the profile and the configurations are not yet deployed.* `Preparing` - Preparing to deploy the configuration.* `Validating` - Profile validation in progress.* `Evaluating` - Policy edit configuration change evaluation in progress.* `Configuring` - Profile deploy operation is in progress.* `UnConfiguring` - Server is unassigned and config cleanup is in progress.* `Analyzing` - Profile changes are being analyzed.* `Activating` - Configuration is being activated at the endpoint.* `Inconsistent` - Profile is inconsistent with the endpoint configuration.* `Associated` - The profile configuration has been applied to the endpoint and no inconsistencies have been detected.* `Failed` - The last action on the profile has failed.* `Not-complete` - Config import operation on the profile is not complete.* `Waiting-for-resource` - Waiting for the resource to be allocated for the profile.* `Partially-deployed` - The profile configuration has been applied on a subset of endpoints. 
    + `config_type`:(string)(ReadOnly) The type of configuration running on the profile. Since profile deployments can configure multiple different settings, configType indicates which type of configuration is currently in progress. 
    + `control_action`:(string) System action to trigger the appropriate workflow. Values -- No_op, ConfigChange, Deploy, Unbind. 
    + `error_state`:(string) Indicates a profile's error state. Values -- Validation-error (Static validation error), Pre-config-error (Runtime validation error), Config-error (Runtime configuration error). 
    + `inconsistency_reason`:
                (Array of schema.TypeString) -
    + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
    + `oper_state`:(string)(ReadOnly) Combined state (configState, and operational state of the associated physical resource) to indicate the current state of the profile. Values -- n/a, Power-off, Pending-changes, Configuring, Ok, Failed. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
* `config_change_details`:(Array)(ReadOnly) An array of relationships to serverConfigChangeDetail resources. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `config_changes`:(HashMap) -(ReadOnly) Pending configuration changes at the summary level. Detail changes are saved in configChangeDetails as a separate object. 
This complex property has following sub-properties:
  + `changes`:
                (Array of schema.TypeString) -
  + `disruptions`:
                (Array of schema.TypeString) -
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `policy_disruptions`:(Array)
This complex property has following sub-properties:
    + `disruptions`:
                (Array of schema.TypeString) -
    + `is_only_required_by_other_policies`:(bool) The current policy has to be redeployed only because there are other policy changes that require this. 
    + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
    + `policy_name`:(string) Name of the policy that, when modified, causes the disruption. 
    + `policy_pending_action`:(string) Name of the action which is pending on this policy. Example, if policy is not yet activated we mark this field as not-activated. Currently we support two actions, not-deployed and not-activated. 
    + `required_by_policies`:
                (Array of schema.TypeString) -
* `config_context`:(HashMap) - The configuration state and results of the last configuration operation. 
This complex property has following sub-properties:
  + `config_state`:(string)(ReadOnly) Indicates a profile's configuration deploying state. Values -- Assigned, Not-assigned, Associated, Pending-changes, Out-of-sync, Validating, Configuring, Failed. 
  + `config_state_summary`:(string)(ReadOnly) Indicates a profile's configuration deploying state. Values -- Assigned, Not-assigned, Associated, InConsistent, Validating, Configuring, Failed, Activating, UnConfiguring.* `None` - The default state is none.* `Not-assigned` - Server is not assigned to the profile.* `Assigned` - Server is assigned to the profile and the configurations are not yet deployed.* `Preparing` - Preparing to deploy the configuration.* `Validating` - Profile validation in progress.* `Evaluating` - Policy edit configuration change evaluation in progress.* `Configuring` - Profile deploy operation is in progress.* `UnConfiguring` - Server is unassigned and config cleanup is in progress.* `Analyzing` - Profile changes are being analyzed.* `Activating` - Configuration is being activated at the endpoint.* `Inconsistent` - Profile is inconsistent with the endpoint configuration.* `Associated` - The profile configuration has been applied to the endpoint and no inconsistencies have been detected.* `Failed` - The last action on the profile has failed.* `Not-complete` - Config import operation on the profile is not complete.* `Waiting-for-resource` - Waiting for the resource to be allocated for the profile.* `Partially-deployed` - The profile configuration has been applied on a subset of endpoints. 
  + `config_type`:(string)(ReadOnly) The type of configuration running on the profile. Since profile deployments can configure multiple different settings, configType indicates which type of configuration is currently in progress. 
  + `control_action`:(string) System action to trigger the appropriate workflow. Values -- No_op, ConfigChange, Deploy, Unbind. 
  + `error_state`:(string) Indicates a profile's error state. Values -- Validation-error (Static validation error), Pre-config-error (Runtime validation error), Config-error (Runtime configuration error). 
  + `inconsistency_reason`:
                (Array of schema.TypeString) -
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `oper_state`:(string)(ReadOnly) Combined state (configState, and operational state of the associated physical resource) to indicate the current state of the profile. Values -- n/a, Power-off, Pending-changes, Configuring, Ok, Failed. 
* `config_result`:(HashMap) -(ReadOnly) A reference to a serverConfigResult resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `create_time`:(string)(ReadOnly) The time when this managed object was created. 
* `deploy_status`:(string)(ReadOnly) The status of the server profile indicating if deployment has been initiated on both fabric interconnects or not.* `None` - Switch profiles not deployed on either of the switches.* `Complete` - Both switch profiles of the cluster profile are deployed.* `Partial` - Only one of the switch profiles of the cluster profile is deployed. 
* `deployed_policies`:
                (Array of schema.TypeString) -
* `deployed_switches`:(string)(ReadOnly) The property which determines if the deployment should be skipped on any of the Fabric Interconnects. It is set based on the state of a fabric interconnect to Intersight before the deployment of the server proile begins.* `None` - Server profile configuration not deployed on either of the fabric interconnects.* `AB` - Server profile configuration deployed on both fabric interconnects.* `A` - Server profile configuration deployed on fabric interconnect A only.* `B` - Server profile configuration deployed on fabric interconnect B only. 
* `description`:(string) Description of the profile. 
* `domain_group_moid`:(string)(ReadOnly) The DomainGroup ID for this managed object. 
* `internal_reservation_references`:(Array)
This complex property has following sub-properties:
  + `additional_properties`:(JSON as string) - Additional Properties as per object type, can be added as JSON using `jsonencode()`. Allowed Types are: [fcpool.ReservationReference](#fcpoolReservationReference)
[ippool.ReservationReference](#ippoolReservationReference)
[iqnpool.ReservationReference](#iqnpoolReservationReference)
[macpool.ReservationReference](#macpoolReservationReference)
[uuidpool.ReservationReference](#uuidpoolReservationReference)
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
  + `reservation_moid`:(string) The moid of the reservation object. 
* `is_pmc_deployed_secure_passphrase_set`:(bool)(ReadOnly) Indicates whether the value of the 'pmcDeployedSecurePassphrase' property has been set. 
* `leased_server`:(HashMap) - A reference to a computePhysical resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `location_details`:(HashMap) -(ReadOnly) Location details associated with the managed object. When a top level resource is assigned to a location, the details of the associated location are propagated to related inventory and configuration objects like servers, profiles, alarms and metrics to facilitate location based filtering. 
This complex property has following sub-properties:
  + `address`:(HashMap) -(ReadOnly) The location's street address. 
This complex property has following sub-properties:
    + `address1`:(string) The primary street address. 
    + `address2`:(string) Additional address information, such as suite number or floor. 
    + `city`:(string) The city where the address is located. 
    + `country`:(string) The country code in ISO 3166-1 alpha-2 format.* `Unknown` - The value Unknown is used when the country code is not known or applicable.* `AD` - The country code for Andorra.* `AE` - The country code for United Arab Emirates.* `AF` - The country code for Afghanistan.* `AG` - The country code for Antigua and Barbuda.* `AI` - The country code for Anguilla.* `AL` - The country code for Albania.* `AM` - The country code for Armenia.* `AO` - The country code for Angola.* `AQ` - The country code for Antarctica.* `AR` - The country code for Argentina.* `AS` - The country code for American Samoa.* `AT` - The country code for Austria.* `AU` - The country code for Australia.* `AW` - The country code for Aruba.* `AX` - The country code for Åland Islands.* `AZ` - The country code for Azerbaijan.* `BA` - The country code for Bosnia and Herzegovina.* `BB` - The country code for Barbados.* `BD` - The country code for Bangladesh.* `BE` - The country code for Belgium.* `BF` - The country code for Burkina Faso.* `BG` - The country code for Bulgaria.* `BH` - The country code for Bahrain.* `BI` - The country code for Burundi.* `BJ` - The country code for Benin.* `BL` - The country code for Saint Barthélemy.* `BM` - The country code for Bermuda.* `BN` - The country code for Brunei Darussalam.* `BO` - The country code for Plurinational State of Bolivia.* `BQ` - The country code for Sint Eustatius and Saba Bonaire.* `BR` - The country code for Brazil.* `BS` - The country code for Bahamas.* `BT` - The country code for Bhutan.* `BV` - The country code for Bouvet Island.* `BW` - The country code for Botswana.* `BZ` - The country code for Belize.* `CA` - The country code for Canada.* `CC` - The country code for Cocos (Keeling) Islands.* `CD` - The country code for Democratic Republic of the Congo.* `CF` - The country code for Central African Republic.* `CG` - The country code for Congo.* `CH` - The country code for Switzerland.* `CI` - The country code for Côte d'Ivoire.* `CK` - The country code for Cook Islands.* `CL` - The country code for Chile.* `CM` - The country code for Cameroon.* `CN` - The country code for China.* `CO` - The country code for Colombia.* `CR` - The country code for Costa Rica.* `CV` - The country code for Cabo Verde.* `CW` - The country code for Curaçao.* `CX` - The country code for Christmas Island.* `CY` - The country code for Cyprus.* `CZ` - The country code for Czechia.* `DE` - The country code for Germany.* `DJ` - The country code for Djibouti.* `DK` - The country code for Denmark.* `DM` - The country code for Dominica.* `DO` - The country code for Dominican Republic.* `DZ` - The country code for Algeria.* `EC` - The country code for Ecuador.* `EE` - The country code for Estonia.* `EG` - The country code for Egypt.* `EH` - The country code for Western Sahara.* `ER` - The country code for Eritrea.* `ES` - The country code for Spain.* `ET` - The country code for Ethiopia.* `FI` - The country code for Finland.* `FJ` - The country code for Fiji.* `FK` - The country code for Falkland Islands (Malvinas).* `FM` - The country code for Federated States of Micronesia.* `FO` - The country code for Faroe Islands.* `FR` - The country code for France.* `GA` - The country code for Gabon.* `GB` - The country code for United Kingdom of Great Britain and Northern Ireland.* `GD` - The country code for Grenada.* `GE` - The country code for Georgia.* `GF` - The country code for French Guiana.* `GG` - The country code for Guernsey.* `GH` - The country code for Ghana.* `GI` - The country code for Gibraltar.* `GL` - The country code for Greenland.* `GM` - The country code for Gambia.* `GN` - The country code for Guinea.* `GP` - The country code for Guadeloupe.* `GQ` - The country code for Equatorial Guinea.* `GR` - The country code for Greece.* `GS` - The country code for South Georgia and the South Sandwich Islands.* `GT` - The country code for Guatemala.* `GU` - The country code for Guam.* `GW` - The country code for Guinea-Bissau.* `GY` - The country code for Guyana.* `HK` - The country code for Hong Kong.* `HM` - The country code for Heard Island and McDonald Islands.* `HN` - The country code for Honduras.* `HR` - The country code for Croatia.* `HT` - The country code for Haiti.* `HU` - The country code for Hungary.* `ID` - The country code for Indonesia.* `IE` - The country code for Ireland.* `IL` - The country code for Israel.* `IM` - The country code for Isle of Man.* `IN` - The country code for India.* `IO` - The country code for British Indian Ocean Territory.* `IQ` - The country code for Iraq.* `IS` - The country code for Iceland.* `IT` - The country code for Italy.* `JE` - The country code for Jersey.* `JM` - The country code for Jamaica.* `JO` - The country code for Jordan.* `JP` - The country code for Japan.* `KE` - The country code for Kenya.* `KG` - The country code for Kyrgyzstan.* `KH` - The country code for Cambodia.* `KI` - The country code for Kiribati.* `KM` - The country code for Comoros.* `KN` - The country code for Saint Kitts and Nevis.* `KR` - The country code for Republic of Korea.* `KW` - The country code for Kuwait.* `KY` - The country code for Cayman Islands.* `KZ` - The country code for Kazakhstan.* `LA` - The country code for Lao People's Democratic Republic.* `LB` - The country code for Lebanon.* `LC` - The country code for Saint Lucia.* `LI` - The country code for Liechtenstein.* `LK` - The country code for Sri Lanka.* `LR` - The country code for Liberia.* `LS` - The country code for Lesotho.* `LT` - The country code for Lithuania.* `LU` - The country code for Luxembourg.* `LV` - The country code for Latvia.* `LY` - The country code for Libya.* `MA` - The country code for Morocco.* `MC` - The country code for Monaco.* `MD` - The country code for Republic of Moldova.* `ME` - The country code for Montenegro.* `MF` - The country code for Saint Martin (French part).* `MG` - The country code for Madagascar.* `MH` - The country code for Marshall Islands.* `MK` - The country code for North Macedonia.* `ML` - The country code for Mali.* `MM` - The country code for Myanmar.* `MN` - The country code for Mongolia.* `MO` - The country code for Macao.* `MP` - The country code for Northern Mariana Islands.* `MQ` - The country code for Martinique.* `MR` - The country code for Mauritania.* `MS` - The country code for Montserrat.* `MT` - The country code for Malta.* `MU` - The country code for Mauritius.* `MV` - The country code for Maldives.* `MW` - The country code for Malawi.* `MX` - The country code for Mexico.* `MY` - The country code for Malaysia.* `MZ` - The country code for Mozambique.* `NA` - The country code for Namibia.* `NC` - The country code for New Caledonia.* `NE` - The country code for Niger.* `NF` - The country code for Norfolk Island.* `NG` - The country code for Nigeria.* `NI` - The country code for Nicaragua.* `NL` - The country code for Kingdom of the Netherlands.* `NO` - The country code for Norway.* `NP` - The country code for Nepal.* `NR` - The country code for Nauru.* `NU` - The country code for Niue.* `NZ` - The country code for New Zealand.* `OM` - The country code for Oman.* `PA` - The country code for Panama.* `PE` - The country code for Peru.* `PF` - The country code for French Polynesia.* `PG` - The country code for Papua New Guinea.* `PH` - The country code for Philippines.* `PK` - The country code for Pakistan.* `PL` - The country code for Poland.* `PM` - The country code for Saint Pierre and Miquelon.* `PN` - The country code for Pitcairn.* `PR` - The country code for Puerto Rico.* `PS` - The country code for State of Palestine.* `PT` - The country code for Portugal.* `PW` - The country code for Palau.* `PY` - The country code for Paraguay.* `QA` - The country code for Qatar.* `RE` - The country code for Réunion.* `RO` - The country code for Romania.* `RS` - The country code for Serbia.* `RW` - The country code for Rwanda.* `SA` - The country code for Saudi Arabia.* `SB` - The country code for Solomon Islands.* `SC` - The country code for Seychelles.* `SD` - The country code for Sudan.* `SE` - The country code for Sweden.* `SG` - The country code for Singapore.* `SH` - The country code for Ascension and Tristan da Cunha Saint Helena.* `SI` - The country code for Slovenia.* `SJ` - The country code for Svalbard and Jan Mayen.* `SK` - The country code for Slovakia.* `SL` - The country code for Sierra Leone.* `SM` - The country code for San Marino.* `SN` - The country code for Senegal.* `SO` - The country code for Somalia.* `SR` - The country code for Suriname.* `SS` - The country code for South Sudan.* `ST` - The country code for Sao Tome and Principe.* `SV` - The country code for El Salvador.* `SX` - The country code for Sint Maarten (Dutch part).* `SZ` - The country code for Eswatini.* `TC` - The country code for Turks and Caicos Islands.* `TD` - The country code for Chad.* `TF` - The country code for French Southern Territories.* `TG` - The country code for Togo.* `TH` - The country code for Thailand.* `TJ` - The country code for Tajikistan.* `TK` - The country code for Tokelau.* `TL` - The country code for Timor-Leste.* `TM` - The country code for Turkmenistan.* `TN` - The country code for Tunisia.* `TO` - The country code for Tonga.* `TR` - The country code for Türkiye.* `TT` - The country code for Trinidad and Tobago.* `TV` - The country code for Tuvalu.* `TW` - The country code for Province of China Taiwan.* `TZ` - The country code for United Republic of Tanzania.* `UA` - The country code for Ukraine.* `UG` - The country code for Uganda.* `UM` - The country code for United States Minor Outlying Islands.* `US` - The country code for United States of America.* `UY` - The country code for Uruguay.* `UZ` - The country code for Uzbekistan.* `VA` - The country code for Holy See.* `VC` - The country code for Saint Vincent and the Grenadines.* `VE` - The country code for Bolivarian Republic of Venezuela.* `VG` - The country code for Virgin Islands (British).* `VI` - The country code for Virgin Islands (U.S.).* `VN` - The country code for Viet Nam.* `VU` - The country code for Vanuatu.* `WF` - The country code for Wallis and Futuna.* `WS` - The country code for Samoa.* `YE` - The country code for Yemen.* `YT` - The country code for Mayotte.* `ZA` - The country code for South Africa.* `ZM` - The country code for Zambia.* `ZW` - The country code for Zimbabwe. 
    + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
    + `postal_code`:(string) The postal or ZIP code for the address. 
    + `state_province`:(string) The state or province where the address is located. 
  + `coordinates`:(HashMap) -(ReadOnly) The location's longitude and latitude coordinates. 
This complex property has following sub-properties:
    + `latitude`:(float) The latitude coordinate value. 
    + `longitude`:(float) The longitude coordinate value. 
    + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `name`:(string)(ReadOnly) A user provided name for the location. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
* `management_mode`:(string)(ReadOnly) The management mode of the server.* `IntersightStandalone` - Intersight Standalone mode of operation.* `Intersight` - Intersight managed mode of operation. 
* `mod_time`:(string)(ReadOnly) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the profile instance or profile template. 
* `organization`:(HashMap) - A reference to a organizationOrganization resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `overridden_list`:
                (Array of schema.TypeString) -
* `owners`:
                (Array of schema.TypeString) -(ReadOnly)
* `parent`:(HashMap) -(ReadOnly) A reference to a moBaseMo resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `partially_deployed_policies`:(Array)
This complex property has following sub-properties:
  + `end_point_context`:(string)(ReadOnly) Information about the endpoint to which it is applied.* `Server` - Configuration is applied to a server context.* `FI` - Configuration is applied to a Fabric Identifier (FI) context.* `IOM` - Configuration is applied to an Input/Output Module (IOM) context. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `policy`:(string)(ReadOnly) The name of the policy for which entry is created. 
* `permission_resources`:(Array)(ReadOnly) An array of relationships to moBaseMo resources. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `pmc_deployed_secure_passphrase`:(string) Secure passphrase that is already deployed on all the Persistent Memory Modules on the server. This deployed passphrase is required during deploy of server profile if secure passphrase is changed or security is disabled in the attached persistent memory policy. 
* `policy_bucket`:(Array) An array of relationships to policyAbstractPolicy resources. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `policy_change_details`:(Array)
This complex property has following sub-properties:
  + `changes`:
                (Array of schema.TypeString) -
  + `config_change_context`:(HashMap) - Context information on the change. 
This complex property has following sub-properties:
    + `dependent_policy_list`:
                (Array of schema.TypeString) -
    + `entity_data`:(JSON as string)(ReadOnly) The data of the object present in config result context. 
    + `entity_moid`:(string) The Moid of the object present in config result context. 
    + `entity_name`:(string) The name of the object present in config result context. 
    + `entity_type`:(string) The type of the object present in config result context. 
    + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
    + `parent_moid`:(string) The Moid of the parent object present in config result context. 
    + `parent_policy_object_type`:(string) The type of the policy object associated with the profile. 
    + `parent_type`:(string) The type of the parent object present in config result context. 
  + `config_change_flag`:(string) Config change flag to differentiate Pending-changes and Config-drift.* `Pending-changes` - Config change flag represents changes are due to not deployed changes from Intersight.* `Drift-changes` - Config change flag represents changes are due to endpoint configuration changes. 
  + `disruptions`:
                (Array of schema.TypeString) -
  + `message`:(string) Detailed description of the config change. 
  + `mod_status`:(string) Modification status of the mo in this config change.* `None` - The 'none' operation/state.Indicates a configuration profile has been deployed, and the desired configuration matches the actual device configuration.* `Created` - The 'create' operation/state.Indicates a configuration profile has been created and associated with a device, but the configuration specified in the profilehas not been applied yet. For example, this could happen when the user creates a server profile and has not deployed the profile yet.* `Modified` - The 'update' operation/state.Indicates some of the desired configuration changes specified in a profile have not been been applied to the associated device.This happens when the user has made changes to a profile and has not deployed the changes yet, or when the workflow to applythe configuration changes has not completed successfully.* `Deleted` - The 'delete' operation/state.Indicates a configuration profile has been been disassociated from a device and the user has not undeployed these changes yet. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
* `post_deploy_action`:
                (Array of schema.TypeString) -
* `removed_policies`:
                (Array of schema.TypeString) -
* `reported_policy_changes`:(Array)
This complex property has following sub-properties:
  + `change_id`:(string)(ReadOnly) The change evaluation identifier for which the change is reported. 
  + `change_status`:(string)(ReadOnly) The status of policy change evaluation which has been reported.* `Initiated` - The status when policy change evaluation is triggered for a policy.* `Reported` - The status when policy change evaluation is reported for a policy. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `policy_type`:(string)(ReadOnly) The type of policy for which the change has been reported. 
* `reservation_references`:(Array)
This complex property has following sub-properties:
  + `additional_properties`:(JSON as string) - Additional Properties as per object type, can be added as JSON using `jsonencode()`. Allowed Types are: [fcpool.ReservationReference](#fcpoolReservationReference)
[ippool.ReservationReference](#ippoolReservationReference)
[iqnpool.ReservationReference](#iqnpoolReservationReference)
[macpool.ReservationReference](#macpoolReservationReference)
[uuidpool.ReservationReference](#uuidpoolReservationReference)
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
  + `reservation_moid`:(string) The moid of the reservation object. 
* `resource_lease`:(HashMap) -(ReadOnly) A reference to a resourcepoolLease resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `running_workflows`:(Array)(ReadOnly) An array of relationships to workflowWorkflowInfo resources. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `scheduled_actions`:(Array)
This complex property has following sub-properties:
  + `action`:(string) Name of the action to be performed on the profile. 
  + `action_qualifier`:(HashMap) - Qualifiers to control the action being triggered. Action qualifiers are to be specified based on the type of disruptions that an action is to be restricted to. For example, trigger the 'Deploy' action to only perform network and management plane configurations. 
This complex property has following sub-properties:
    + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `proceed_on_reboot`:(bool) ProceedOnReboot can be used to acknowledge server reboot while triggering deploy/activate. 
* `server_assignment_mode`:(string) Source of the server assigned to the Server Profile. Values can be Static, Pool or None. Static is used if a server is attached directly to a Server Profile. Pool is used if a resource pool is attached to a Server Profile. None is used if no server or resource pool is attached to a Server Profile. Slot or Serial pre-assignment is also considered to be None as it is different form of Assign Later.* `None` - No server is assigned to the server profile.* `Static` - Server is directly assigned to server profile using assign server.* `Pool` - Server is assigned from a resource pool. 
* `server_pool`:(HashMap) - A reference to a resourcepoolPool resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `server_pre_assign_by_serial`:(string) Serial number of the server that would be assigned to this pre-assigned Server Profile. It can be any string that adheres to the following constraints:It should start and end with an alphanumeric character.It cannot be more than 20 characters. 
* `server_pre_assign_by_slot`:(HashMap) - Server profile is pre-assigned to a server using slot. 
This complex property has following sub-properties:
  + `chassis_id`:(int) Chassis-id of the slot that would be assigned to this pre-assigned server profile. 
  + `domain_name`:(string) The domain name of the Fabric Interconnect or the chassis name prefix.The name must start and end with an alphanumeric character, can include underscores and hyphens, and has a maximum length of 30 characters. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `slot_id`:(int) Slot-id of the server that would be assigned to this pre-assigned server profile. 
* `shared_scope`:(string)(ReadOnly) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `src_template`:(HashMap) - A reference to a policyAbstractProfile resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `static_uuid_address`:(string) The UUID address for the server must include UUID prefix xxxxxxxx-xxxx-xxxx along with the UUID suffix of format xxxx-xxxxxxxxxxxx. 
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
* `target_platform`:(string) The platform for which the server profile is applicable. It can either be a server that is operating in standalone mode or which is attached to a Fabric Interconnect managed by Intersight.* `Standalone` - Servers which are operating in standalone mode i.e. not connected to a Fabric Interconnected.* `FIAttached` - Servers which are connected to a Fabric Interconnect that is managed by Intersight.* `UnifiedEdgeServer` - Unified Edge sleds that is managed by Intersight. 
* `template_actions`:(Array)
This complex property has following sub-properties:
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `params`:(Array)
This complex property has following sub-properties:
    + `name`:(string) The action parameter identifier. The supported values are SyncType and SyncTimer for the template sync action.* `None` - The default parameter that implies that no action parameter is required for the template action.* `SyncType` - The parameter that describes the type of sync action such as SyncOne or SyncFailed supported on any template or derived object.* `SyncTimer` - The parameter for the initial delay in seconds after which the sync action must be executed. The supported range is from 0 to 60 seconds.* `OverriddenList` - The parameter applicable in attach operation indicating the configurations that must override the template configurations. 
    + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
    + `value`:(string) The action parameter value is based on the action parameter type. Supported action parameters and their values are-a) Name - SyncType, Supported Values - SyncFailed, SyncOne.b) Name - SyncTimer, Supported Values - 0 to 60 seconds.c) Name - OverriddenList, Supported Values - Comma Separated list of overridable configurations. 
  + `type`:(string) The action type to be executed.* `Sync` - The action to merge values from the template to its derived objects.* `Deploy` - The action to execute deploy action on all the objects derived from the template that is mainly applicable for the various profile types.* `Detach` - The action to detach the current derived object from its attached template.* `Attach` - The action to attach the current object to the specified template. 
* `template_sync_errors`:(Array)
This complex property has following sub-properties:
  + `message`:(string)(ReadOnly) The localized message based on the locale setting of the user's context providing the error description. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `type`:(string)(ReadOnly) The error type that indicates the point of failure.* `Transient` - Any error which is a runtime error due to some other action in progress on the derived object that is blocking the sync action. This error type is retriable.For example, when vNIC Template is updated, but the derived vNIC or vNICs are part of a LAN Connectivity policy associated with a profile being deployed to endpoint. In this scenario, the derived vNIC update displays this error.* `Validation` - When the template sync on the derived object fails due to an incorrect configuration, it displays a validation error. This error type is considered fatal and not retried.For example, when a new policy is attached to a server profile template, the sync to a derived server profile fails due to the policy attach errors.* `User` - Any configuration error due to incorrect or invalid input and that requires user intervention for correction, is displayed under this category. This error type is considered fatal and not retried.For example, when a new policy is attached to a server profile template, the sync to a derived server profile fails. This happens when the policyis not applicable to the server assigned to the server profile, such as the Power policy that is not applicable for UCS Rack servers.* `Internal` - Any application internal errors are displayed under this category. This error type is considered fatal and not retried. 
* `template_sync_status`:(string)(ReadOnly) The sync status of the current MO wrt the attached Template MO.* `None` - The Enum value represents that the object is not attached to any template.* `OK` - The Enum value represents that the object values are in sync with attached template.* `Scheduled` - The Enum value represents that the object sync from attached template is scheduled from template.* `InProgress` - The Enum value represents that the object sync with the attached template is in progress.* `OutOfSync` - The Enum value represents that the object values are not in sync with attached template. 
* `type`:(string) Defines the type of the profile. Accepted values are instance or template.* `instance` - The profile defines the configuration for a specific instance of a target. 
* `user_label`:(string) User label assigned to the server profile. 
* `uuid`:(string)(ReadOnly) The UUID address that is assigned to the server based on the UUID pool. 
* `uuid_address_type`:(string) UUID address allocation type selected to assign an UUID address for the server.* `NONE` - The user did not assign any UUID address.* `STATIC` - The user assigns a static UUID address.* `POOL` - The user selects a pool from which the address will be leased. 
* `uuid_lease`:(HashMap) -(ReadOnly) A reference to a uuidpoolUuidLease resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `uuid_pool`:(HashMap) - A reference to a uuidpoolPool resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
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

### Custom keywords
These are
* `wait_for_completion`:(bool) This model object can trigger workflows. Use this option to wait for all running workflows to reach a complete state. Default value is True i.e. wait.

## Import
`intersight_server_profile` can be imported using the Moid of the object, e.g.
```
$ terraform import intersight_server_profile.example 1234567890987654321abcde
```
## Allowed Types in `AdditionalProperties`
 
### [fcpool.ReservationReference](#argument-reference)
The reference to the reservation object.
* `consumer_name`:(string) The consumer name for which the reserved fc pool would be used. 
* `consumer_type`:(string) The consumer type for which the reserved fc pool would be used.* `Vhba` - FC reservation would be used by Vhba.* `WWNN` - FC reservation would be used by WWNN. 

### [ippool.ReservationReference](#argument-reference)
The reference to the reservation object.
* `consumer_name`:(string) The consumer name for which the reserved IP would be used. 
* `consumer_type`:(string) The consumer type for which the reserved IP would be used.* `OutofbandIpv4-Access` - IP reservation would be used for out of band management.* `InbandIpv4-Access` - IP reservation would be used for inband management.* `InbandIpv6-Access` - IP reservation would be used for inband management.* `ISCSI` - IP reservation would be used for ISCSI management. 

### [iqnpool.ReservationReference](#argument-reference)
The reference to the reservation object.

### [macpool.ReservationReference](#argument-reference)
The reference to the reservation object.
* `consumer_name`:(string) The consumer name for which the reserved MAC would be used. 
* `consumer_type`:(string) The consumer type for which the reserved MAC would be used.* `Vnic` - MAC reservation would be used by VNIC. 

### [uuidpool.ReservationReference](#argument-reference)
The reference to the reservation object.
  
### [fcpool.ReservationReference](#argument-reference)
The reference to the reservation object.
* `consumer_name`:(string) The consumer name for which the reserved fc pool would be used. 
* `consumer_type`:(string) The consumer type for which the reserved fc pool would be used.* `Vhba` - FC reservation would be used by Vhba.* `WWNN` - FC reservation would be used by WWNN. 

### [ippool.ReservationReference](#argument-reference)
The reference to the reservation object.
* `consumer_name`:(string) The consumer name for which the reserved IP would be used. 
* `consumer_type`:(string) The consumer type for which the reserved IP would be used.* `OutofbandIpv4-Access` - IP reservation would be used for out of band management.* `InbandIpv4-Access` - IP reservation would be used for inband management.* `InbandIpv6-Access` - IP reservation would be used for inband management.* `ISCSI` - IP reservation would be used for ISCSI management. 

### [iqnpool.ReservationReference](#argument-reference)
The reference to the reservation object.

### [macpool.ReservationReference](#argument-reference)
The reference to the reservation object.
* `consumer_name`:(string) The consumer name for which the reserved MAC would be used. 
* `consumer_type`:(string) The consumer type for which the reserved MAC would be used.* `Vnic` - MAC reservation would be used by VNIC. 

### [uuidpool.ReservationReference](#argument-reference)
The reference to the reservation object.
  
