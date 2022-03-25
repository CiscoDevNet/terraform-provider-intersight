# FabricServerRole

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "fabric.ServerRole"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "fabric.ServerRole"]
**AutoNegotiationDisabled** | Pointer to **bool** | Auto negotiation configuration for server port. This configuration is required only for FEX Model N9K-C93180YC-FX3 connected with 100G speed port on UCS-FI-6536 and should be set as True. | [optional] [default to false]
**Fec** | Pointer to **string** | Forward error correction configuration for server port. This configuration is required only for FEX Model N9K-C93180YC-FX3 connected with 25G speed ports on UCS-FI-6454/UCS-FI-64108 and should be set as Cl74. * &#x60;Auto&#x60; - Forward error correction option &#39;Auto&#39;. * &#x60;Cl91&#x60; - Forward error correction option &#39;cl91&#39;. * &#x60;Cl74&#x60; - Forward error correction option &#39;cl74&#39;. | [optional] [default to "Auto"]

## Methods

### NewFabricServerRole

`func NewFabricServerRole(classId string, objectType string, ) *FabricServerRole`

NewFabricServerRole instantiates a new FabricServerRole object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFabricServerRoleWithDefaults

`func NewFabricServerRoleWithDefaults() *FabricServerRole`

NewFabricServerRoleWithDefaults instantiates a new FabricServerRole object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *FabricServerRole) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *FabricServerRole) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *FabricServerRole) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *FabricServerRole) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *FabricServerRole) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *FabricServerRole) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAutoNegotiationDisabled

`func (o *FabricServerRole) GetAutoNegotiationDisabled() bool`

GetAutoNegotiationDisabled returns the AutoNegotiationDisabled field if non-nil, zero value otherwise.

### GetAutoNegotiationDisabledOk

`func (o *FabricServerRole) GetAutoNegotiationDisabledOk() (*bool, bool)`

GetAutoNegotiationDisabledOk returns a tuple with the AutoNegotiationDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoNegotiationDisabled

`func (o *FabricServerRole) SetAutoNegotiationDisabled(v bool)`

SetAutoNegotiationDisabled sets AutoNegotiationDisabled field to given value.

### HasAutoNegotiationDisabled

`func (o *FabricServerRole) HasAutoNegotiationDisabled() bool`

HasAutoNegotiationDisabled returns a boolean if a field has been set.

### GetFec

`func (o *FabricServerRole) GetFec() string`

GetFec returns the Fec field if non-nil, zero value otherwise.

### GetFecOk

`func (o *FabricServerRole) GetFecOk() (*string, bool)`

GetFecOk returns a tuple with the Fec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFec

`func (o *FabricServerRole) SetFec(v string)`

SetFec sets Fec field to given value.

### HasFec

`func (o *FabricServerRole) HasFec() bool`

HasFec returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


