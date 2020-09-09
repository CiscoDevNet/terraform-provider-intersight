# ConnectorpackUpgradeImpactAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectorPack** | Pointer to [**[]ConnectorpackConnectorPackUpdate**](connectorpack.ConnectorPackUpdate.md) |  | [optional] 
**IsEligibleForUpgrade** | Pointer to **bool** | States whether the UCS Director is eligible for an upgrade. Set to true if connector packs are available for upgrade, else set to false. | [optional] [readonly] 
**IsUpdateDownloaded** | Pointer to **bool** | States whether all the requisite updates have been downloaded to the target UCS Director. Set to true if all connector packs required to upgrade UCS Director to the next iteration have been downloaded, else set to false. | [optional] [readonly] 
**UcsdInfo** | Pointer to [**IaasUcsdInfoRelationship**](iaas.UcsdInfo.Relationship.md) |  | [optional] 

## Methods

### NewConnectorpackUpgradeImpactAllOf

`func NewConnectorpackUpgradeImpactAllOf() *ConnectorpackUpgradeImpactAllOf`

NewConnectorpackUpgradeImpactAllOf instantiates a new ConnectorpackUpgradeImpactAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectorpackUpgradeImpactAllOfWithDefaults

`func NewConnectorpackUpgradeImpactAllOfWithDefaults() *ConnectorpackUpgradeImpactAllOf`

NewConnectorpackUpgradeImpactAllOfWithDefaults instantiates a new ConnectorpackUpgradeImpactAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectorPack

`func (o *ConnectorpackUpgradeImpactAllOf) GetConnectorPack() []ConnectorpackConnectorPackUpdate`

GetConnectorPack returns the ConnectorPack field if non-nil, zero value otherwise.

### GetConnectorPackOk

`func (o *ConnectorpackUpgradeImpactAllOf) GetConnectorPackOk() (*[]ConnectorpackConnectorPackUpdate, bool)`

GetConnectorPackOk returns a tuple with the ConnectorPack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectorPack

`func (o *ConnectorpackUpgradeImpactAllOf) SetConnectorPack(v []ConnectorpackConnectorPackUpdate)`

SetConnectorPack sets ConnectorPack field to given value.

### HasConnectorPack

`func (o *ConnectorpackUpgradeImpactAllOf) HasConnectorPack() bool`

HasConnectorPack returns a boolean if a field has been set.

### GetIsEligibleForUpgrade

`func (o *ConnectorpackUpgradeImpactAllOf) GetIsEligibleForUpgrade() bool`

GetIsEligibleForUpgrade returns the IsEligibleForUpgrade field if non-nil, zero value otherwise.

### GetIsEligibleForUpgradeOk

`func (o *ConnectorpackUpgradeImpactAllOf) GetIsEligibleForUpgradeOk() (*bool, bool)`

GetIsEligibleForUpgradeOk returns a tuple with the IsEligibleForUpgrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEligibleForUpgrade

`func (o *ConnectorpackUpgradeImpactAllOf) SetIsEligibleForUpgrade(v bool)`

SetIsEligibleForUpgrade sets IsEligibleForUpgrade field to given value.

### HasIsEligibleForUpgrade

`func (o *ConnectorpackUpgradeImpactAllOf) HasIsEligibleForUpgrade() bool`

HasIsEligibleForUpgrade returns a boolean if a field has been set.

### GetIsUpdateDownloaded

`func (o *ConnectorpackUpgradeImpactAllOf) GetIsUpdateDownloaded() bool`

GetIsUpdateDownloaded returns the IsUpdateDownloaded field if non-nil, zero value otherwise.

### GetIsUpdateDownloadedOk

`func (o *ConnectorpackUpgradeImpactAllOf) GetIsUpdateDownloadedOk() (*bool, bool)`

GetIsUpdateDownloadedOk returns a tuple with the IsUpdateDownloaded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUpdateDownloaded

`func (o *ConnectorpackUpgradeImpactAllOf) SetIsUpdateDownloaded(v bool)`

SetIsUpdateDownloaded sets IsUpdateDownloaded field to given value.

### HasIsUpdateDownloaded

`func (o *ConnectorpackUpgradeImpactAllOf) HasIsUpdateDownloaded() bool`

HasIsUpdateDownloaded returns a boolean if a field has been set.

### GetUcsdInfo

`func (o *ConnectorpackUpgradeImpactAllOf) GetUcsdInfo() IaasUcsdInfoRelationship`

GetUcsdInfo returns the UcsdInfo field if non-nil, zero value otherwise.

### GetUcsdInfoOk

`func (o *ConnectorpackUpgradeImpactAllOf) GetUcsdInfoOk() (*IaasUcsdInfoRelationship, bool)`

GetUcsdInfoOk returns a tuple with the UcsdInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUcsdInfo

`func (o *ConnectorpackUpgradeImpactAllOf) SetUcsdInfo(v IaasUcsdInfoRelationship)`

SetUcsdInfo sets UcsdInfo field to given value.

### HasUcsdInfo

`func (o *ConnectorpackUpgradeImpactAllOf) HasUcsdInfo() bool`

HasUcsdInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


