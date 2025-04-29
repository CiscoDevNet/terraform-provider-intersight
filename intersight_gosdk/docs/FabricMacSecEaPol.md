# FabricMacSecEaPol

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "fabric.MacSecEaPol"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "fabric.MacSecEaPol"]
**EaPolEthertype** | Pointer to **string** | Ethertype to use in extensible authentication protocol over LAN (EAPoL) frames for MACsec key agreement (MKA) protocol data units (PDUs). The range is between 0x600 - 0xffff. | [optional] [default to "34958"]
**EaPolMacAddress** | Pointer to **string** | MAC address to use in extensible authentication protocol over LAN (EAPoL) for MACsec key agreement (MKA) protocol data units (PDUs). EAPol mac address should not be equal to all-zero (0000.0000.0000). | [optional] [default to "0180.C200.0003"]

## Methods

### NewFabricMacSecEaPol

`func NewFabricMacSecEaPol(classId string, objectType string, ) *FabricMacSecEaPol`

NewFabricMacSecEaPol instantiates a new FabricMacSecEaPol object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFabricMacSecEaPolWithDefaults

`func NewFabricMacSecEaPolWithDefaults() *FabricMacSecEaPol`

NewFabricMacSecEaPolWithDefaults instantiates a new FabricMacSecEaPol object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *FabricMacSecEaPol) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *FabricMacSecEaPol) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *FabricMacSecEaPol) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *FabricMacSecEaPol) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *FabricMacSecEaPol) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *FabricMacSecEaPol) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetEaPolEthertype

`func (o *FabricMacSecEaPol) GetEaPolEthertype() string`

GetEaPolEthertype returns the EaPolEthertype field if non-nil, zero value otherwise.

### GetEaPolEthertypeOk

`func (o *FabricMacSecEaPol) GetEaPolEthertypeOk() (*string, bool)`

GetEaPolEthertypeOk returns a tuple with the EaPolEthertype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEaPolEthertype

`func (o *FabricMacSecEaPol) SetEaPolEthertype(v string)`

SetEaPolEthertype sets EaPolEthertype field to given value.

### HasEaPolEthertype

`func (o *FabricMacSecEaPol) HasEaPolEthertype() bool`

HasEaPolEthertype returns a boolean if a field has been set.

### GetEaPolMacAddress

`func (o *FabricMacSecEaPol) GetEaPolMacAddress() string`

GetEaPolMacAddress returns the EaPolMacAddress field if non-nil, zero value otherwise.

### GetEaPolMacAddressOk

`func (o *FabricMacSecEaPol) GetEaPolMacAddressOk() (*string, bool)`

GetEaPolMacAddressOk returns a tuple with the EaPolMacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEaPolMacAddress

`func (o *FabricMacSecEaPol) SetEaPolMacAddress(v string)`

SetEaPolMacAddress sets EaPolMacAddress field to given value.

### HasEaPolMacAddress

`func (o *FabricMacSecEaPol) HasEaPolMacAddress() bool`

HasEaPolMacAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


