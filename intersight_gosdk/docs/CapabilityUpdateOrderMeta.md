# CapabilityUpdateOrderMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "capability.UpdateOrderMeta"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "capability.UpdateOrderMeta"]
**Category** | Pointer to **string** | The category of the model series. | [optional] [readonly] 
**Description** | Pointer to **string** | Verbose description regarding this group. | [optional] [readonly] 
**PlatformTag** | Pointer to **string** | The platform tag value of the model series. | [optional] [readonly] 
**UpdateOrder** | Pointer to [**[]CapabilityUpdateOrderListType**](CapabilityUpdateOrderListType.md) |  | [optional] 

## Methods

### NewCapabilityUpdateOrderMeta

`func NewCapabilityUpdateOrderMeta(classId string, objectType string, ) *CapabilityUpdateOrderMeta`

NewCapabilityUpdateOrderMeta instantiates a new CapabilityUpdateOrderMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCapabilityUpdateOrderMetaWithDefaults

`func NewCapabilityUpdateOrderMetaWithDefaults() *CapabilityUpdateOrderMeta`

NewCapabilityUpdateOrderMetaWithDefaults instantiates a new CapabilityUpdateOrderMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CapabilityUpdateOrderMeta) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CapabilityUpdateOrderMeta) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CapabilityUpdateOrderMeta) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CapabilityUpdateOrderMeta) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CapabilityUpdateOrderMeta) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CapabilityUpdateOrderMeta) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCategory

`func (o *CapabilityUpdateOrderMeta) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *CapabilityUpdateOrderMeta) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *CapabilityUpdateOrderMeta) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *CapabilityUpdateOrderMeta) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetDescription

`func (o *CapabilityUpdateOrderMeta) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CapabilityUpdateOrderMeta) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CapabilityUpdateOrderMeta) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CapabilityUpdateOrderMeta) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPlatformTag

`func (o *CapabilityUpdateOrderMeta) GetPlatformTag() string`

GetPlatformTag returns the PlatformTag field if non-nil, zero value otherwise.

### GetPlatformTagOk

`func (o *CapabilityUpdateOrderMeta) GetPlatformTagOk() (*string, bool)`

GetPlatformTagOk returns a tuple with the PlatformTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformTag

`func (o *CapabilityUpdateOrderMeta) SetPlatformTag(v string)`

SetPlatformTag sets PlatformTag field to given value.

### HasPlatformTag

`func (o *CapabilityUpdateOrderMeta) HasPlatformTag() bool`

HasPlatformTag returns a boolean if a field has been set.

### GetUpdateOrder

`func (o *CapabilityUpdateOrderMeta) GetUpdateOrder() []CapabilityUpdateOrderListType`

GetUpdateOrder returns the UpdateOrder field if non-nil, zero value otherwise.

### GetUpdateOrderOk

`func (o *CapabilityUpdateOrderMeta) GetUpdateOrderOk() (*[]CapabilityUpdateOrderListType, bool)`

GetUpdateOrderOk returns a tuple with the UpdateOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateOrder

`func (o *CapabilityUpdateOrderMeta) SetUpdateOrder(v []CapabilityUpdateOrderListType)`

SetUpdateOrder sets UpdateOrder field to given value.

### HasUpdateOrder

`func (o *CapabilityUpdateOrderMeta) HasUpdateOrder() bool`

HasUpdateOrder returns a boolean if a field has been set.

### SetUpdateOrderNil

`func (o *CapabilityUpdateOrderMeta) SetUpdateOrderNil(b bool)`

 SetUpdateOrderNil sets the value for UpdateOrder to be an explicit nil

### UnsetUpdateOrder
`func (o *CapabilityUpdateOrderMeta) UnsetUpdateOrder()`

UnsetUpdateOrder ensures that no value is present for UpdateOrder, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


