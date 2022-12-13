# HyperflexLunAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hyperflex.Lun"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hyperflex.Lun"]
**Description** | Pointer to **string** | The description of iSCSI LUN. | [optional] [readonly] 
**DsCapacityInBytes** | Pointer to **int64** | The datastore capacity in bytes. | [optional] [readonly] 
**DsName** | Pointer to **string** | The HyperFlex datastore name. | [optional] [readonly] 
**DsUuid** | Pointer to **string** | The HyperFlex datastore UUID. | [optional] [readonly] 
**InventorySource** | Pointer to **string** | Source of the lun inventory. * &#x60;NOT_APPLICABLE&#x60; - The source of the HyperFlex inventory is not applicable. * &#x60;ONLINE&#x60; - The source of the HyperFlex inventory is online. * &#x60;OFFLINE&#x60; - The source of the HyperFlex inventory is offline. | [optional] [readonly] [default to "NOT_APPLICABLE"]
**IsEncrypted** | Pointer to **bool** | Indicates if the datastore on which LUN resides is encrypted or un-encrypted. | [optional] [readonly] 
**LunId** | Pointer to **string** | The identity of HyperFlex iSCSI LUN. | [optional] [readonly] 
**Name** | Pointer to **string** | Name of the HyperFlex iSCSI LUN. | [optional] [readonly] 
**SerialNo** | Pointer to **string** | Serial number of HyperFlex iSCSI LUN. | [optional] [readonly] 
**TotalCapacityInBytes** | Pointer to **int64** | Total capacity of iSCSI LUN in bytes. | [optional] [readonly] 
**Tuuid** | Pointer to **string** | HyperFlex iSCSI Target associated with the HyperFlex iSCSI LUN. | [optional] [readonly] 
**UsedCapacityInBytes** | Pointer to **int64** | Used capacity of iSCSI LUN in bytes. | [optional] [readonly] 
**Uuid** | Pointer to **string** | UUID of the HyperFlex iSCSI LUN. | [optional] [readonly] 
**Version** | Pointer to **int64** | Version of the HyperFlex iSCSI lun. | [optional] [readonly] 
**Target** | Pointer to [**HyperflexTargetRelationship**](HyperflexTargetRelationship.md) |  | [optional] 

## Methods

### NewHyperflexLunAllOf

`func NewHyperflexLunAllOf(classId string, objectType string, ) *HyperflexLunAllOf`

NewHyperflexLunAllOf instantiates a new HyperflexLunAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexLunAllOfWithDefaults

`func NewHyperflexLunAllOfWithDefaults() *HyperflexLunAllOf`

NewHyperflexLunAllOfWithDefaults instantiates a new HyperflexLunAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HyperflexLunAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexLunAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexLunAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HyperflexLunAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexLunAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexLunAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDescription

`func (o *HyperflexLunAllOf) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *HyperflexLunAllOf) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *HyperflexLunAllOf) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *HyperflexLunAllOf) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDsCapacityInBytes

`func (o *HyperflexLunAllOf) GetDsCapacityInBytes() int64`

GetDsCapacityInBytes returns the DsCapacityInBytes field if non-nil, zero value otherwise.

### GetDsCapacityInBytesOk

`func (o *HyperflexLunAllOf) GetDsCapacityInBytesOk() (*int64, bool)`

GetDsCapacityInBytesOk returns a tuple with the DsCapacityInBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDsCapacityInBytes

`func (o *HyperflexLunAllOf) SetDsCapacityInBytes(v int64)`

SetDsCapacityInBytes sets DsCapacityInBytes field to given value.

### HasDsCapacityInBytes

`func (o *HyperflexLunAllOf) HasDsCapacityInBytes() bool`

HasDsCapacityInBytes returns a boolean if a field has been set.

### GetDsName

`func (o *HyperflexLunAllOf) GetDsName() string`

GetDsName returns the DsName field if non-nil, zero value otherwise.

### GetDsNameOk

`func (o *HyperflexLunAllOf) GetDsNameOk() (*string, bool)`

GetDsNameOk returns a tuple with the DsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDsName

`func (o *HyperflexLunAllOf) SetDsName(v string)`

SetDsName sets DsName field to given value.

### HasDsName

`func (o *HyperflexLunAllOf) HasDsName() bool`

HasDsName returns a boolean if a field has been set.

### GetDsUuid

`func (o *HyperflexLunAllOf) GetDsUuid() string`

GetDsUuid returns the DsUuid field if non-nil, zero value otherwise.

### GetDsUuidOk

`func (o *HyperflexLunAllOf) GetDsUuidOk() (*string, bool)`

GetDsUuidOk returns a tuple with the DsUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDsUuid

`func (o *HyperflexLunAllOf) SetDsUuid(v string)`

SetDsUuid sets DsUuid field to given value.

### HasDsUuid

`func (o *HyperflexLunAllOf) HasDsUuid() bool`

HasDsUuid returns a boolean if a field has been set.

### GetInventorySource

`func (o *HyperflexLunAllOf) GetInventorySource() string`

GetInventorySource returns the InventorySource field if non-nil, zero value otherwise.

### GetInventorySourceOk

`func (o *HyperflexLunAllOf) GetInventorySourceOk() (*string, bool)`

GetInventorySourceOk returns a tuple with the InventorySource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventorySource

`func (o *HyperflexLunAllOf) SetInventorySource(v string)`

SetInventorySource sets InventorySource field to given value.

### HasInventorySource

`func (o *HyperflexLunAllOf) HasInventorySource() bool`

HasInventorySource returns a boolean if a field has been set.

### GetIsEncrypted

`func (o *HyperflexLunAllOf) GetIsEncrypted() bool`

GetIsEncrypted returns the IsEncrypted field if non-nil, zero value otherwise.

### GetIsEncryptedOk

`func (o *HyperflexLunAllOf) GetIsEncryptedOk() (*bool, bool)`

GetIsEncryptedOk returns a tuple with the IsEncrypted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEncrypted

`func (o *HyperflexLunAllOf) SetIsEncrypted(v bool)`

SetIsEncrypted sets IsEncrypted field to given value.

### HasIsEncrypted

`func (o *HyperflexLunAllOf) HasIsEncrypted() bool`

HasIsEncrypted returns a boolean if a field has been set.

### GetLunId

`func (o *HyperflexLunAllOf) GetLunId() string`

GetLunId returns the LunId field if non-nil, zero value otherwise.

### GetLunIdOk

`func (o *HyperflexLunAllOf) GetLunIdOk() (*string, bool)`

GetLunIdOk returns a tuple with the LunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLunId

`func (o *HyperflexLunAllOf) SetLunId(v string)`

SetLunId sets LunId field to given value.

### HasLunId

`func (o *HyperflexLunAllOf) HasLunId() bool`

HasLunId returns a boolean if a field has been set.

### GetName

`func (o *HyperflexLunAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HyperflexLunAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HyperflexLunAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *HyperflexLunAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSerialNo

`func (o *HyperflexLunAllOf) GetSerialNo() string`

GetSerialNo returns the SerialNo field if non-nil, zero value otherwise.

### GetSerialNoOk

`func (o *HyperflexLunAllOf) GetSerialNoOk() (*string, bool)`

GetSerialNoOk returns a tuple with the SerialNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNo

`func (o *HyperflexLunAllOf) SetSerialNo(v string)`

SetSerialNo sets SerialNo field to given value.

### HasSerialNo

`func (o *HyperflexLunAllOf) HasSerialNo() bool`

HasSerialNo returns a boolean if a field has been set.

### GetTotalCapacityInBytes

`func (o *HyperflexLunAllOf) GetTotalCapacityInBytes() int64`

GetTotalCapacityInBytes returns the TotalCapacityInBytes field if non-nil, zero value otherwise.

### GetTotalCapacityInBytesOk

`func (o *HyperflexLunAllOf) GetTotalCapacityInBytesOk() (*int64, bool)`

GetTotalCapacityInBytesOk returns a tuple with the TotalCapacityInBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCapacityInBytes

`func (o *HyperflexLunAllOf) SetTotalCapacityInBytes(v int64)`

SetTotalCapacityInBytes sets TotalCapacityInBytes field to given value.

### HasTotalCapacityInBytes

`func (o *HyperflexLunAllOf) HasTotalCapacityInBytes() bool`

HasTotalCapacityInBytes returns a boolean if a field has been set.

### GetTuuid

`func (o *HyperflexLunAllOf) GetTuuid() string`

GetTuuid returns the Tuuid field if non-nil, zero value otherwise.

### GetTuuidOk

`func (o *HyperflexLunAllOf) GetTuuidOk() (*string, bool)`

GetTuuidOk returns a tuple with the Tuuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTuuid

`func (o *HyperflexLunAllOf) SetTuuid(v string)`

SetTuuid sets Tuuid field to given value.

### HasTuuid

`func (o *HyperflexLunAllOf) HasTuuid() bool`

HasTuuid returns a boolean if a field has been set.

### GetUsedCapacityInBytes

`func (o *HyperflexLunAllOf) GetUsedCapacityInBytes() int64`

GetUsedCapacityInBytes returns the UsedCapacityInBytes field if non-nil, zero value otherwise.

### GetUsedCapacityInBytesOk

`func (o *HyperflexLunAllOf) GetUsedCapacityInBytesOk() (*int64, bool)`

GetUsedCapacityInBytesOk returns a tuple with the UsedCapacityInBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedCapacityInBytes

`func (o *HyperflexLunAllOf) SetUsedCapacityInBytes(v int64)`

SetUsedCapacityInBytes sets UsedCapacityInBytes field to given value.

### HasUsedCapacityInBytes

`func (o *HyperflexLunAllOf) HasUsedCapacityInBytes() bool`

HasUsedCapacityInBytes returns a boolean if a field has been set.

### GetUuid

`func (o *HyperflexLunAllOf) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *HyperflexLunAllOf) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *HyperflexLunAllOf) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *HyperflexLunAllOf) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetVersion

`func (o *HyperflexLunAllOf) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *HyperflexLunAllOf) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *HyperflexLunAllOf) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *HyperflexLunAllOf) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetTarget

`func (o *HyperflexLunAllOf) GetTarget() HyperflexTargetRelationship`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *HyperflexLunAllOf) GetTargetOk() (*HyperflexTargetRelationship, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *HyperflexLunAllOf) SetTarget(v HyperflexTargetRelationship)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *HyperflexLunAllOf) HasTarget() bool`

HasTarget returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


