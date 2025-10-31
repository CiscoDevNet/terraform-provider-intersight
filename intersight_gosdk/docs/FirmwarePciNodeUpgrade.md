# FirmwarePciNodeUpgrade

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "firmware.PciNodeUpgrade"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "firmware.PciNodeUpgrade"]
**Device** | Pointer to [**NullableAssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 
**PciNode** | Pointer to [**NullablePciNodeRelationship**](PciNodeRelationship.md) |  | [optional] 

## Methods

### NewFirmwarePciNodeUpgrade

`func NewFirmwarePciNodeUpgrade(classId string, objectType string, ) *FirmwarePciNodeUpgrade`

NewFirmwarePciNodeUpgrade instantiates a new FirmwarePciNodeUpgrade object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirmwarePciNodeUpgradeWithDefaults

`func NewFirmwarePciNodeUpgradeWithDefaults() *FirmwarePciNodeUpgrade`

NewFirmwarePciNodeUpgradeWithDefaults instantiates a new FirmwarePciNodeUpgrade object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *FirmwarePciNodeUpgrade) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *FirmwarePciNodeUpgrade) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *FirmwarePciNodeUpgrade) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *FirmwarePciNodeUpgrade) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *FirmwarePciNodeUpgrade) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *FirmwarePciNodeUpgrade) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDevice

`func (o *FirmwarePciNodeUpgrade) GetDevice() AssetDeviceRegistrationRelationship`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *FirmwarePciNodeUpgrade) GetDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *FirmwarePciNodeUpgrade) SetDevice(v AssetDeviceRegistrationRelationship)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *FirmwarePciNodeUpgrade) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### SetDeviceNil

`func (o *FirmwarePciNodeUpgrade) SetDeviceNil(b bool)`

 SetDeviceNil sets the value for Device to be an explicit nil

### UnsetDevice
`func (o *FirmwarePciNodeUpgrade) UnsetDevice()`

UnsetDevice ensures that no value is present for Device, not even an explicit nil
### GetPciNode

`func (o *FirmwarePciNodeUpgrade) GetPciNode() PciNodeRelationship`

GetPciNode returns the PciNode field if non-nil, zero value otherwise.

### GetPciNodeOk

`func (o *FirmwarePciNodeUpgrade) GetPciNodeOk() (*PciNodeRelationship, bool)`

GetPciNodeOk returns a tuple with the PciNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPciNode

`func (o *FirmwarePciNodeUpgrade) SetPciNode(v PciNodeRelationship)`

SetPciNode sets PciNode field to given value.

### HasPciNode

`func (o *FirmwarePciNodeUpgrade) HasPciNode() bool`

HasPciNode returns a boolean if a field has been set.

### SetPciNodeNil

`func (o *FirmwarePciNodeUpgrade) SetPciNodeNil(b bool)`

 SetPciNodeNil sets the value for PciNode to be an explicit nil

### UnsetPciNode
`func (o *FirmwarePciNodeUpgrade) UnsetPciNode()`

UnsetPciNode ensures that no value is present for PciNode, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


