# ConnectorpackUpgradeImpact

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "connectorpack.UpgradeImpact"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "connectorpack.UpgradeImpact"]
**ConnectorPack** | Pointer to [**[]ConnectorpackConnectorPackUpdate**](ConnectorpackConnectorPackUpdate.md) |  | [optional] 
**IsEligibleForUpgrade** | Pointer to **bool** | States whether the UCS Director is eligible for an upgrade. Set to true if connector packs are available for upgrade, else set to false. | [optional] [readonly] 
**IsUpdateDownloaded** | Pointer to **bool** | States whether all the requisite updates have been downloaded to the target UCS Director. Set to true if all connector packs required to upgrade UCS Director to the next iteration have been downloaded, else set to false. | [optional] [readonly] 
**UcsdInfo** | Pointer to [**IaasUcsdInfoRelationship**](IaasUcsdInfoRelationship.md) |  | [optional] 

## Methods

### NewConnectorpackUpgradeImpact

`func NewConnectorpackUpgradeImpact(classId string, objectType string, ) *ConnectorpackUpgradeImpact`

NewConnectorpackUpgradeImpact instantiates a new ConnectorpackUpgradeImpact object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectorpackUpgradeImpactWithDefaults

`func NewConnectorpackUpgradeImpactWithDefaults() *ConnectorpackUpgradeImpact`

NewConnectorpackUpgradeImpactWithDefaults instantiates a new ConnectorpackUpgradeImpact object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ConnectorpackUpgradeImpact) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ConnectorpackUpgradeImpact) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ConnectorpackUpgradeImpact) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ConnectorpackUpgradeImpact) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ConnectorpackUpgradeImpact) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ConnectorpackUpgradeImpact) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetConnectorPack

`func (o *ConnectorpackUpgradeImpact) GetConnectorPack() []ConnectorpackConnectorPackUpdate`

GetConnectorPack returns the ConnectorPack field if non-nil, zero value otherwise.

### GetConnectorPackOk

`func (o *ConnectorpackUpgradeImpact) GetConnectorPackOk() (*[]ConnectorpackConnectorPackUpdate, bool)`

GetConnectorPackOk returns a tuple with the ConnectorPack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectorPack

`func (o *ConnectorpackUpgradeImpact) SetConnectorPack(v []ConnectorpackConnectorPackUpdate)`

SetConnectorPack sets ConnectorPack field to given value.

### HasConnectorPack

`func (o *ConnectorpackUpgradeImpact) HasConnectorPack() bool`

HasConnectorPack returns a boolean if a field has been set.

### SetConnectorPackNil

`func (o *ConnectorpackUpgradeImpact) SetConnectorPackNil(b bool)`

 SetConnectorPackNil sets the value for ConnectorPack to be an explicit nil

### UnsetConnectorPack
`func (o *ConnectorpackUpgradeImpact) UnsetConnectorPack()`

UnsetConnectorPack ensures that no value is present for ConnectorPack, not even an explicit nil
### GetIsEligibleForUpgrade

`func (o *ConnectorpackUpgradeImpact) GetIsEligibleForUpgrade() bool`

GetIsEligibleForUpgrade returns the IsEligibleForUpgrade field if non-nil, zero value otherwise.

### GetIsEligibleForUpgradeOk

`func (o *ConnectorpackUpgradeImpact) GetIsEligibleForUpgradeOk() (*bool, bool)`

GetIsEligibleForUpgradeOk returns a tuple with the IsEligibleForUpgrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEligibleForUpgrade

`func (o *ConnectorpackUpgradeImpact) SetIsEligibleForUpgrade(v bool)`

SetIsEligibleForUpgrade sets IsEligibleForUpgrade field to given value.

### HasIsEligibleForUpgrade

`func (o *ConnectorpackUpgradeImpact) HasIsEligibleForUpgrade() bool`

HasIsEligibleForUpgrade returns a boolean if a field has been set.

### GetIsUpdateDownloaded

`func (o *ConnectorpackUpgradeImpact) GetIsUpdateDownloaded() bool`

GetIsUpdateDownloaded returns the IsUpdateDownloaded field if non-nil, zero value otherwise.

### GetIsUpdateDownloadedOk

`func (o *ConnectorpackUpgradeImpact) GetIsUpdateDownloadedOk() (*bool, bool)`

GetIsUpdateDownloadedOk returns a tuple with the IsUpdateDownloaded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUpdateDownloaded

`func (o *ConnectorpackUpgradeImpact) SetIsUpdateDownloaded(v bool)`

SetIsUpdateDownloaded sets IsUpdateDownloaded field to given value.

### HasIsUpdateDownloaded

`func (o *ConnectorpackUpgradeImpact) HasIsUpdateDownloaded() bool`

HasIsUpdateDownloaded returns a boolean if a field has been set.

### GetUcsdInfo

`func (o *ConnectorpackUpgradeImpact) GetUcsdInfo() IaasUcsdInfoRelationship`

GetUcsdInfo returns the UcsdInfo field if non-nil, zero value otherwise.

### GetUcsdInfoOk

`func (o *ConnectorpackUpgradeImpact) GetUcsdInfoOk() (*IaasUcsdInfoRelationship, bool)`

GetUcsdInfoOk returns a tuple with the UcsdInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUcsdInfo

`func (o *ConnectorpackUpgradeImpact) SetUcsdInfo(v IaasUcsdInfoRelationship)`

SetUcsdInfo sets UcsdInfo field to given value.

### HasUcsdInfo

`func (o *ConnectorpackUpgradeImpact) HasUcsdInfo() bool`

HasUcsdInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


