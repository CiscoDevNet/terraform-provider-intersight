# MemoryPolicyInventory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "memory.PolicyInventory"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "memory.PolicyInventory"]
**EnableDimmBlocklisting** | Pointer to **bool** | Enable DIMM Blocklisting on the server. | [optional] [readonly] [default to false]
**TargetMo** | Pointer to [**NullableMoBaseMoRelationship**](MoBaseMoRelationship.md) |  | [optional] 

## Methods

### NewMemoryPolicyInventory

`func NewMemoryPolicyInventory(classId string, objectType string, ) *MemoryPolicyInventory`

NewMemoryPolicyInventory instantiates a new MemoryPolicyInventory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMemoryPolicyInventoryWithDefaults

`func NewMemoryPolicyInventoryWithDefaults() *MemoryPolicyInventory`

NewMemoryPolicyInventoryWithDefaults instantiates a new MemoryPolicyInventory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *MemoryPolicyInventory) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *MemoryPolicyInventory) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *MemoryPolicyInventory) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *MemoryPolicyInventory) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *MemoryPolicyInventory) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *MemoryPolicyInventory) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetEnableDimmBlocklisting

`func (o *MemoryPolicyInventory) GetEnableDimmBlocklisting() bool`

GetEnableDimmBlocklisting returns the EnableDimmBlocklisting field if non-nil, zero value otherwise.

### GetEnableDimmBlocklistingOk

`func (o *MemoryPolicyInventory) GetEnableDimmBlocklistingOk() (*bool, bool)`

GetEnableDimmBlocklistingOk returns a tuple with the EnableDimmBlocklisting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDimmBlocklisting

`func (o *MemoryPolicyInventory) SetEnableDimmBlocklisting(v bool)`

SetEnableDimmBlocklisting sets EnableDimmBlocklisting field to given value.

### HasEnableDimmBlocklisting

`func (o *MemoryPolicyInventory) HasEnableDimmBlocklisting() bool`

HasEnableDimmBlocklisting returns a boolean if a field has been set.

### GetTargetMo

`func (o *MemoryPolicyInventory) GetTargetMo() MoBaseMoRelationship`

GetTargetMo returns the TargetMo field if non-nil, zero value otherwise.

### GetTargetMoOk

`func (o *MemoryPolicyInventory) GetTargetMoOk() (*MoBaseMoRelationship, bool)`

GetTargetMoOk returns a tuple with the TargetMo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetMo

`func (o *MemoryPolicyInventory) SetTargetMo(v MoBaseMoRelationship)`

SetTargetMo sets TargetMo field to given value.

### HasTargetMo

`func (o *MemoryPolicyInventory) HasTargetMo() bool`

HasTargetMo returns a boolean if a field has been set.

### SetTargetMoNil

`func (o *MemoryPolicyInventory) SetTargetMoNil(b bool)`

 SetTargetMoNil sets the value for TargetMo to be an explicit nil

### UnsetTargetMo
`func (o *MemoryPolicyInventory) UnsetTargetMo()`

UnsetTargetMo ensures that no value is present for TargetMo, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


