# CapabilityUpdateOrderListType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "capability.UpdateOrderListType"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "capability.UpdateOrderListType"]
**InterimVersion** | Pointer to **string** | The interim version to be updated to in the update order. | [optional] [readonly] 
**SourceVersion** | Pointer to **string** | The source version of the update order. | [optional] [readonly] 
**SupportedModels** | Pointer to **[]string** |  | [optional] 
**TargetVersion** | Pointer to **string** | The target version of the update order. | [optional] [readonly] 

## Methods

### NewCapabilityUpdateOrderListType

`func NewCapabilityUpdateOrderListType(classId string, objectType string, ) *CapabilityUpdateOrderListType`

NewCapabilityUpdateOrderListType instantiates a new CapabilityUpdateOrderListType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCapabilityUpdateOrderListTypeWithDefaults

`func NewCapabilityUpdateOrderListTypeWithDefaults() *CapabilityUpdateOrderListType`

NewCapabilityUpdateOrderListTypeWithDefaults instantiates a new CapabilityUpdateOrderListType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CapabilityUpdateOrderListType) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CapabilityUpdateOrderListType) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CapabilityUpdateOrderListType) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CapabilityUpdateOrderListType) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CapabilityUpdateOrderListType) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CapabilityUpdateOrderListType) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetInterimVersion

`func (o *CapabilityUpdateOrderListType) GetInterimVersion() string`

GetInterimVersion returns the InterimVersion field if non-nil, zero value otherwise.

### GetInterimVersionOk

`func (o *CapabilityUpdateOrderListType) GetInterimVersionOk() (*string, bool)`

GetInterimVersionOk returns a tuple with the InterimVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterimVersion

`func (o *CapabilityUpdateOrderListType) SetInterimVersion(v string)`

SetInterimVersion sets InterimVersion field to given value.

### HasInterimVersion

`func (o *CapabilityUpdateOrderListType) HasInterimVersion() bool`

HasInterimVersion returns a boolean if a field has been set.

### GetSourceVersion

`func (o *CapabilityUpdateOrderListType) GetSourceVersion() string`

GetSourceVersion returns the SourceVersion field if non-nil, zero value otherwise.

### GetSourceVersionOk

`func (o *CapabilityUpdateOrderListType) GetSourceVersionOk() (*string, bool)`

GetSourceVersionOk returns a tuple with the SourceVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceVersion

`func (o *CapabilityUpdateOrderListType) SetSourceVersion(v string)`

SetSourceVersion sets SourceVersion field to given value.

### HasSourceVersion

`func (o *CapabilityUpdateOrderListType) HasSourceVersion() bool`

HasSourceVersion returns a boolean if a field has been set.

### GetSupportedModels

`func (o *CapabilityUpdateOrderListType) GetSupportedModels() []string`

GetSupportedModels returns the SupportedModels field if non-nil, zero value otherwise.

### GetSupportedModelsOk

`func (o *CapabilityUpdateOrderListType) GetSupportedModelsOk() (*[]string, bool)`

GetSupportedModelsOk returns a tuple with the SupportedModels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedModels

`func (o *CapabilityUpdateOrderListType) SetSupportedModels(v []string)`

SetSupportedModels sets SupportedModels field to given value.

### HasSupportedModels

`func (o *CapabilityUpdateOrderListType) HasSupportedModels() bool`

HasSupportedModels returns a boolean if a field has been set.

### SetSupportedModelsNil

`func (o *CapabilityUpdateOrderListType) SetSupportedModelsNil(b bool)`

 SetSupportedModelsNil sets the value for SupportedModels to be an explicit nil

### UnsetSupportedModels
`func (o *CapabilityUpdateOrderListType) UnsetSupportedModels()`

UnsetSupportedModels ensures that no value is present for SupportedModels, not even an explicit nil
### GetTargetVersion

`func (o *CapabilityUpdateOrderListType) GetTargetVersion() string`

GetTargetVersion returns the TargetVersion field if non-nil, zero value otherwise.

### GetTargetVersionOk

`func (o *CapabilityUpdateOrderListType) GetTargetVersionOk() (*string, bool)`

GetTargetVersionOk returns a tuple with the TargetVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetVersion

`func (o *CapabilityUpdateOrderListType) SetTargetVersion(v string)`

SetTargetVersion sets TargetVersion field to given value.

### HasTargetVersion

`func (o *CapabilityUpdateOrderListType) HasTargetVersion() bool`

HasTargetVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


