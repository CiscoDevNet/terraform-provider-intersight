# VnicPlacementSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "vnic.PlacementSettings"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "vnic.PlacementSettings"]
**Id** | Pointer to **string** | PCIe Slot where the VIC adapter is installed. Supported values are (1-15) and MLOM. | [optional] 
**PciLink** | Pointer to **int64** | The PCI Link used as transport for the virtual interface. All VIC adapters have a single PCI link except VIC 1385 which has two. | [optional] [default to 0]
**SwitchId** | Pointer to **string** | The fabric port to which the vNICs will be associated. * &#x60;None&#x60; - Fabric Id is not set to either A or B for the standalone case where the server is not connected to Fabric Interconnects. The value &#39;None&#39; should be used. * &#x60;A&#x60; - Fabric A of the FI cluster. * &#x60;B&#x60; - Fabric B of the FI cluster. | [optional] [default to "None"]
**Uplink** | Pointer to **int64** | Adapter port on which the virtual interface will be created. | [optional] 

## Methods

### NewVnicPlacementSettings

`func NewVnicPlacementSettings(classId string, objectType string, ) *VnicPlacementSettings`

NewVnicPlacementSettings instantiates a new VnicPlacementSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVnicPlacementSettingsWithDefaults

`func NewVnicPlacementSettingsWithDefaults() *VnicPlacementSettings`

NewVnicPlacementSettingsWithDefaults instantiates a new VnicPlacementSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VnicPlacementSettings) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VnicPlacementSettings) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VnicPlacementSettings) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VnicPlacementSettings) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VnicPlacementSettings) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VnicPlacementSettings) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetId

`func (o *VnicPlacementSettings) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VnicPlacementSettings) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VnicPlacementSettings) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *VnicPlacementSettings) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPciLink

`func (o *VnicPlacementSettings) GetPciLink() int64`

GetPciLink returns the PciLink field if non-nil, zero value otherwise.

### GetPciLinkOk

`func (o *VnicPlacementSettings) GetPciLinkOk() (*int64, bool)`

GetPciLinkOk returns a tuple with the PciLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPciLink

`func (o *VnicPlacementSettings) SetPciLink(v int64)`

SetPciLink sets PciLink field to given value.

### HasPciLink

`func (o *VnicPlacementSettings) HasPciLink() bool`

HasPciLink returns a boolean if a field has been set.

### GetSwitchId

`func (o *VnicPlacementSettings) GetSwitchId() string`

GetSwitchId returns the SwitchId field if non-nil, zero value otherwise.

### GetSwitchIdOk

`func (o *VnicPlacementSettings) GetSwitchIdOk() (*string, bool)`

GetSwitchIdOk returns a tuple with the SwitchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchId

`func (o *VnicPlacementSettings) SetSwitchId(v string)`

SetSwitchId sets SwitchId field to given value.

### HasSwitchId

`func (o *VnicPlacementSettings) HasSwitchId() bool`

HasSwitchId returns a boolean if a field has been set.

### GetUplink

`func (o *VnicPlacementSettings) GetUplink() int64`

GetUplink returns the Uplink field if non-nil, zero value otherwise.

### GetUplinkOk

`func (o *VnicPlacementSettings) GetUplinkOk() (*int64, bool)`

GetUplinkOk returns a tuple with the Uplink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplink

`func (o *VnicPlacementSettings) SetUplink(v int64)`

SetUplink sets Uplink field to given value.

### HasUplink

`func (o *VnicPlacementSettings) HasUplink() bool`

HasUplink returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


