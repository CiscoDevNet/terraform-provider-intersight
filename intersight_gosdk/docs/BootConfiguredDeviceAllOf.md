# BootConfiguredDeviceAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**Name** | Pointer to **string** | The name of the boot device configured in the boot policy. | [optional] 
**Order** | Pointer to **int64** | The order of the boot device configured in the boot policy. | [optional] 
**State** | Pointer to **string** | The state of the boot device configured in the boot policy. | [optional] 
**Type** | Pointer to **string** | The type of the boot device configured in the boot policy. | [optional] 

## Methods

### NewBootConfiguredDeviceAllOf

`func NewBootConfiguredDeviceAllOf(classId string, objectType string, ) *BootConfiguredDeviceAllOf`

NewBootConfiguredDeviceAllOf instantiates a new BootConfiguredDeviceAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBootConfiguredDeviceAllOfWithDefaults

`func NewBootConfiguredDeviceAllOfWithDefaults() *BootConfiguredDeviceAllOf`

NewBootConfiguredDeviceAllOfWithDefaults instantiates a new BootConfiguredDeviceAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *BootConfiguredDeviceAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *BootConfiguredDeviceAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *BootConfiguredDeviceAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *BootConfiguredDeviceAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *BootConfiguredDeviceAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *BootConfiguredDeviceAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetName

`func (o *BootConfiguredDeviceAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BootConfiguredDeviceAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BootConfiguredDeviceAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BootConfiguredDeviceAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrder

`func (o *BootConfiguredDeviceAllOf) GetOrder() int64`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *BootConfiguredDeviceAllOf) GetOrderOk() (*int64, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *BootConfiguredDeviceAllOf) SetOrder(v int64)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *BootConfiguredDeviceAllOf) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetState

`func (o *BootConfiguredDeviceAllOf) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *BootConfiguredDeviceAllOf) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *BootConfiguredDeviceAllOf) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *BootConfiguredDeviceAllOf) HasState() bool`

HasState returns a boolean if a field has been set.

### GetType

`func (o *BootConfiguredDeviceAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BootConfiguredDeviceAllOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BootConfiguredDeviceAllOf) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *BootConfiguredDeviceAllOf) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


