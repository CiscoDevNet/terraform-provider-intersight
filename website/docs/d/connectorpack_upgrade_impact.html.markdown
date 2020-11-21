
---
layout: "intersight"
page_title: "Intersight: intersight_connectorpack_upgrade_impact"
sidebar_current: "docs-intersight-data-source-connectorpack-upgrade-impact"
description: |-
Used to determine the list of connector packs to be installed on a target UCS Director in its next upgrade cycle. Accepts the moid of the target UcsdInfo as part of the filter query. Given below is a sample url :- https://{{target}}/api/v1/connectorpack/UpgradeImpacts?$filter= ( UcsdInfo.Moid eq <<MoId>> ).
---

# Data Source: intersight_connectorpack._upgrade_impact
Used to determine the list of connector packs to be installed on a target UCS Director in its next upgrade cycle. Accepts the moid of the target UcsdInfo as part of the filter query. Given below is a sample url :- https://{{target}}/api/v1/connectorpack/UpgradeImpacts?$filter= ( UcsdInfo.Moid eq <<MoId>> ).
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `is_eligible_for_upgrade`:(bool) States whether the UCS Director is eligible for an upgrade. Set to true if connector packs are available for upgrade, else set to false. 
* `is_update_downloaded`:(bool) States whether all the requisite updates have been downloaded to the target UCS Director. Set to true if all connector packs required to upgrade UCS Director to the next iteration have been downloaded, else set to false. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
