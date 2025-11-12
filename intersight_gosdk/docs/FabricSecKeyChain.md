# FabricSecKeyChain

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "fabric.SecKeyChain"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "fabric.SecKeyChain"]
**Name** | Pointer to **string** | Must be a maximum of 63 characters, without spacing. | [optional] 
**SecKeys** | Pointer to [**[]FabricSecKey**](FabricSecKey.md) |  | [optional] 

## Methods

### NewFabricSecKeyChain

`func NewFabricSecKeyChain(classId string, objectType string, ) *FabricSecKeyChain`

NewFabricSecKeyChain instantiates a new FabricSecKeyChain object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFabricSecKeyChainWithDefaults

`func NewFabricSecKeyChainWithDefaults() *FabricSecKeyChain`

NewFabricSecKeyChainWithDefaults instantiates a new FabricSecKeyChain object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *FabricSecKeyChain) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *FabricSecKeyChain) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *FabricSecKeyChain) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *FabricSecKeyChain) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *FabricSecKeyChain) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *FabricSecKeyChain) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetName

`func (o *FabricSecKeyChain) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FabricSecKeyChain) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FabricSecKeyChain) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FabricSecKeyChain) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSecKeys

`func (o *FabricSecKeyChain) GetSecKeys() []FabricSecKey`

GetSecKeys returns the SecKeys field if non-nil, zero value otherwise.

### GetSecKeysOk

`func (o *FabricSecKeyChain) GetSecKeysOk() (*[]FabricSecKey, bool)`

GetSecKeysOk returns a tuple with the SecKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecKeys

`func (o *FabricSecKeyChain) SetSecKeys(v []FabricSecKey)`

SetSecKeys sets SecKeys field to given value.

### HasSecKeys

`func (o *FabricSecKeyChain) HasSecKeys() bool`

HasSecKeys returns a boolean if a field has been set.

### SetSecKeysNil

`func (o *FabricSecKeyChain) SetSecKeysNil(b bool)`

 SetSecKeysNil sets the value for SecKeys to be an explicit nil

### UnsetSecKeys
`func (o *FabricSecKeyChain) UnsetSecKeys()`

UnsetSecKeys ensures that no value is present for SecKeys, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


