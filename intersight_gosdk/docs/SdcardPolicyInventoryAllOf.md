# SdcardPolicyInventoryAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "sdcard.PolicyInventory"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "sdcard.PolicyInventory"]
**Partitions** | Pointer to [**[]SdcardPartition**](SdcardPartition.md) |  | [optional] 
**TargetMo** | Pointer to [**MoBaseMoRelationship**](MoBaseMoRelationship.md) |  | [optional] 

## Methods

### NewSdcardPolicyInventoryAllOf

`func NewSdcardPolicyInventoryAllOf(classId string, objectType string, ) *SdcardPolicyInventoryAllOf`

NewSdcardPolicyInventoryAllOf instantiates a new SdcardPolicyInventoryAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSdcardPolicyInventoryAllOfWithDefaults

`func NewSdcardPolicyInventoryAllOfWithDefaults() *SdcardPolicyInventoryAllOf`

NewSdcardPolicyInventoryAllOfWithDefaults instantiates a new SdcardPolicyInventoryAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *SdcardPolicyInventoryAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *SdcardPolicyInventoryAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *SdcardPolicyInventoryAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *SdcardPolicyInventoryAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *SdcardPolicyInventoryAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *SdcardPolicyInventoryAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetPartitions

`func (o *SdcardPolicyInventoryAllOf) GetPartitions() []SdcardPartition`

GetPartitions returns the Partitions field if non-nil, zero value otherwise.

### GetPartitionsOk

`func (o *SdcardPolicyInventoryAllOf) GetPartitionsOk() (*[]SdcardPartition, bool)`

GetPartitionsOk returns a tuple with the Partitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitions

`func (o *SdcardPolicyInventoryAllOf) SetPartitions(v []SdcardPartition)`

SetPartitions sets Partitions field to given value.

### HasPartitions

`func (o *SdcardPolicyInventoryAllOf) HasPartitions() bool`

HasPartitions returns a boolean if a field has been set.

### SetPartitionsNil

`func (o *SdcardPolicyInventoryAllOf) SetPartitionsNil(b bool)`

 SetPartitionsNil sets the value for Partitions to be an explicit nil

### UnsetPartitions
`func (o *SdcardPolicyInventoryAllOf) UnsetPartitions()`

UnsetPartitions ensures that no value is present for Partitions, not even an explicit nil
### GetTargetMo

`func (o *SdcardPolicyInventoryAllOf) GetTargetMo() MoBaseMoRelationship`

GetTargetMo returns the TargetMo field if non-nil, zero value otherwise.

### GetTargetMoOk

`func (o *SdcardPolicyInventoryAllOf) GetTargetMoOk() (*MoBaseMoRelationship, bool)`

GetTargetMoOk returns a tuple with the TargetMo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetMo

`func (o *SdcardPolicyInventoryAllOf) SetTargetMo(v MoBaseMoRelationship)`

SetTargetMo sets TargetMo field to given value.

### HasTargetMo

`func (o *SdcardPolicyInventoryAllOf) HasTargetMo() bool`

HasTargetMo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


