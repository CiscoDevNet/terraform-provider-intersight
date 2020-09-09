# VnicPlacementSettingsAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | PCIe Slot where the VIC adapter is installed. Supported values are (1-15) and MLOM. | [optional] 
**PciLink** | Pointer to **int64** | The PCI Link used as transport for the virtual interface. All VIC adapters have a single PCI link except VIC 1385 which has two. | [optional] 
**SwitchId** | Pointer to **string** | The fabric port to which the vnics will be associated. * &#x60;None&#x60; - Fabric Id is not set to either A or B for the standalone case where the server is not connected to Fabric Interconnects. The value &#39;None&#39; should be used. * &#x60;A&#x60; - Fabric A of the FI cluster. * &#x60;B&#x60; - Fabric B of the FI cluster. | [optional] [default to "None"]
**Uplink** | Pointer to **int64** | Adapter port on which the virtual interface will be created. | [optional] 

## Methods

### NewVnicPlacementSettingsAllOf

`func NewVnicPlacementSettingsAllOf() *VnicPlacementSettingsAllOf`

NewVnicPlacementSettingsAllOf instantiates a new VnicPlacementSettingsAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVnicPlacementSettingsAllOfWithDefaults

`func NewVnicPlacementSettingsAllOfWithDefaults() *VnicPlacementSettingsAllOf`

NewVnicPlacementSettingsAllOfWithDefaults instantiates a new VnicPlacementSettingsAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VnicPlacementSettingsAllOf) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VnicPlacementSettingsAllOf) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VnicPlacementSettingsAllOf) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *VnicPlacementSettingsAllOf) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPciLink

`func (o *VnicPlacementSettingsAllOf) GetPciLink() int64`

GetPciLink returns the PciLink field if non-nil, zero value otherwise.

### GetPciLinkOk

`func (o *VnicPlacementSettingsAllOf) GetPciLinkOk() (*int64, bool)`

GetPciLinkOk returns a tuple with the PciLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPciLink

`func (o *VnicPlacementSettingsAllOf) SetPciLink(v int64)`

SetPciLink sets PciLink field to given value.

### HasPciLink

`func (o *VnicPlacementSettingsAllOf) HasPciLink() bool`

HasPciLink returns a boolean if a field has been set.

### GetSwitchId

`func (o *VnicPlacementSettingsAllOf) GetSwitchId() string`

GetSwitchId returns the SwitchId field if non-nil, zero value otherwise.

### GetSwitchIdOk

`func (o *VnicPlacementSettingsAllOf) GetSwitchIdOk() (*string, bool)`

GetSwitchIdOk returns a tuple with the SwitchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchId

`func (o *VnicPlacementSettingsAllOf) SetSwitchId(v string)`

SetSwitchId sets SwitchId field to given value.

### HasSwitchId

`func (o *VnicPlacementSettingsAllOf) HasSwitchId() bool`

HasSwitchId returns a boolean if a field has been set.

### GetUplink

`func (o *VnicPlacementSettingsAllOf) GetUplink() int64`

GetUplink returns the Uplink field if non-nil, zero value otherwise.

### GetUplinkOk

`func (o *VnicPlacementSettingsAllOf) GetUplinkOk() (*int64, bool)`

GetUplinkOk returns a tuple with the Uplink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplink

`func (o *VnicPlacementSettingsAllOf) SetUplink(v int64)`

SetUplink sets Uplink field to given value.

### HasUplink

`func (o *VnicPlacementSettingsAllOf) HasUplink() bool`

HasUplink returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


