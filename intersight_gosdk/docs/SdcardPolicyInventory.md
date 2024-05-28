# SdcardPolicyInventory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "sdcard.PolicyInventory"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "sdcard.PolicyInventory"]
**Partitions** | Pointer to [**[]SdcardPartition**](SdcardPartition.md) |  | [optional] 
**TargetMo** | Pointer to [**NullableMoBaseMoRelationship**](MoBaseMoRelationship.md) |  | [optional] 

## Methods

### NewSdcardPolicyInventory

`func NewSdcardPolicyInventory(classId string, objectType string, ) *SdcardPolicyInventory`

NewSdcardPolicyInventory instantiates a new SdcardPolicyInventory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSdcardPolicyInventoryWithDefaults

`func NewSdcardPolicyInventoryWithDefaults() *SdcardPolicyInventory`

NewSdcardPolicyInventoryWithDefaults instantiates a new SdcardPolicyInventory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *SdcardPolicyInventory) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *SdcardPolicyInventory) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *SdcardPolicyInventory) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *SdcardPolicyInventory) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *SdcardPolicyInventory) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *SdcardPolicyInventory) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetPartitions

`func (o *SdcardPolicyInventory) GetPartitions() []SdcardPartition`

GetPartitions returns the Partitions field if non-nil, zero value otherwise.

### GetPartitionsOk

`func (o *SdcardPolicyInventory) GetPartitionsOk() (*[]SdcardPartition, bool)`

GetPartitionsOk returns a tuple with the Partitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitions

`func (o *SdcardPolicyInventory) SetPartitions(v []SdcardPartition)`

SetPartitions sets Partitions field to given value.

### HasPartitions

`func (o *SdcardPolicyInventory) HasPartitions() bool`

HasPartitions returns a boolean if a field has been set.

### SetPartitionsNil

`func (o *SdcardPolicyInventory) SetPartitionsNil(b bool)`

 SetPartitionsNil sets the value for Partitions to be an explicit nil

### UnsetPartitions
`func (o *SdcardPolicyInventory) UnsetPartitions()`

UnsetPartitions ensures that no value is present for Partitions, not even an explicit nil
### GetTargetMo

`func (o *SdcardPolicyInventory) GetTargetMo() MoBaseMoRelationship`

GetTargetMo returns the TargetMo field if non-nil, zero value otherwise.

### GetTargetMoOk

`func (o *SdcardPolicyInventory) GetTargetMoOk() (*MoBaseMoRelationship, bool)`

GetTargetMoOk returns a tuple with the TargetMo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetMo

`func (o *SdcardPolicyInventory) SetTargetMo(v MoBaseMoRelationship)`

SetTargetMo sets TargetMo field to given value.

### HasTargetMo

`func (o *SdcardPolicyInventory) HasTargetMo() bool`

HasTargetMo returns a boolean if a field has been set.

### SetTargetMoNil

`func (o *SdcardPolicyInventory) SetTargetMoNil(b bool)`

 SetTargetMoNil sets the value for TargetMo to be an explicit nil

### UnsetTargetMo
`func (o *SdcardPolicyInventory) UnsetTargetMo()`

UnsetTargetMo ensures that no value is present for TargetMo, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


