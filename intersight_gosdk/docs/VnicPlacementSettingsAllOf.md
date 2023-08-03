# VnicPlacementSettingsAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "vnic.PlacementSettings"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "vnic.PlacementSettings"]
**AutoPciLink** | Pointer to **bool** | Enable or disable automatic assignment of the PCI Link in a dual-link adapter. This option applies only to 13xx series VICs that support dual-link. If enabled, the system determines the placement of the vNIC/vHBA on either of the PCI Links. | [optional] [default to false]
**AutoSlotId** | Pointer to **bool** | Enable or disable automatic assignment of the VIC slot ID. If enabled and the server has only one VIC, the same VIC is chosen for all the vNICs. If enabled and the server has multiple VICs, the vNIC/vHBA are deployed on the first VIC. The Slot ID determines the first VIC. MLOM is the first Slot ID and the ID increments to 2, 3, and so on. | [optional] [default to false]
**Id** | Pointer to **string** | PCIe Slot where the VIC adapter is installed. Supported values are (1-15) and MLOM. | [optional] 
**PciLink** | Pointer to **int64** | The PCI Link used as transport for the virtual interface. PCI Link is only applicable for select Cisco UCS VIC 1300 models (UCSC-PCIE-C40Q-03, UCSB-MLOM-40G-03, UCSB-VIC-M83-8P) that support two PCI links. The value, if specified, for any other VIC model will be ignored. | [optional] [default to 0]
**PciLinkAssignmentMode** | Pointer to **string** | If the autoPciLink is disabled, the user can either choose to place the vNICs manually or based on a policy.If the autoPciLink is enabled, it will be set to None. * &#x60;Custom&#x60; - The user needs to specify the PCI Link manually. * &#x60;Load-Balanced&#x60; - The system will uniformly distribute the interfaces across the PCI Links. * &#x60;None&#x60; - Assignment is not applicable and will be set when the AutoPciLink is set to true. | [optional] [default to "Custom"]
**SwitchId** | Pointer to **string** | The fabric port to which the vNICs will be associated. * &#x60;None&#x60; - Fabric Id is not set to either A or B for the standalone case where the server is not connected to Fabric Interconnects. The value &#39;None&#39; should be used. * &#x60;A&#x60; - Fabric A of the FI cluster. * &#x60;B&#x60; - Fabric B of the FI cluster. | [optional] [default to "None"]
**Uplink** | Pointer to **int64** | Adapter port on which the virtual interface will be created. | [optional] 

## Methods

### NewVnicPlacementSettingsAllOf

`func NewVnicPlacementSettingsAllOf(classId string, objectType string, ) *VnicPlacementSettingsAllOf`

NewVnicPlacementSettingsAllOf instantiates a new VnicPlacementSettingsAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVnicPlacementSettingsAllOfWithDefaults

`func NewVnicPlacementSettingsAllOfWithDefaults() *VnicPlacementSettingsAllOf`

NewVnicPlacementSettingsAllOfWithDefaults instantiates a new VnicPlacementSettingsAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VnicPlacementSettingsAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VnicPlacementSettingsAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VnicPlacementSettingsAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VnicPlacementSettingsAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VnicPlacementSettingsAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VnicPlacementSettingsAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAutoPciLink

`func (o *VnicPlacementSettingsAllOf) GetAutoPciLink() bool`

GetAutoPciLink returns the AutoPciLink field if non-nil, zero value otherwise.

### GetAutoPciLinkOk

`func (o *VnicPlacementSettingsAllOf) GetAutoPciLinkOk() (*bool, bool)`

GetAutoPciLinkOk returns a tuple with the AutoPciLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoPciLink

`func (o *VnicPlacementSettingsAllOf) SetAutoPciLink(v bool)`

SetAutoPciLink sets AutoPciLink field to given value.

### HasAutoPciLink

`func (o *VnicPlacementSettingsAllOf) HasAutoPciLink() bool`

HasAutoPciLink returns a boolean if a field has been set.

### GetAutoSlotId

`func (o *VnicPlacementSettingsAllOf) GetAutoSlotId() bool`

GetAutoSlotId returns the AutoSlotId field if non-nil, zero value otherwise.

### GetAutoSlotIdOk

`func (o *VnicPlacementSettingsAllOf) GetAutoSlotIdOk() (*bool, bool)`

GetAutoSlotIdOk returns a tuple with the AutoSlotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoSlotId

`func (o *VnicPlacementSettingsAllOf) SetAutoSlotId(v bool)`

SetAutoSlotId sets AutoSlotId field to given value.

### HasAutoSlotId

`func (o *VnicPlacementSettingsAllOf) HasAutoSlotId() bool`

HasAutoSlotId returns a boolean if a field has been set.

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

### GetPciLinkAssignmentMode

`func (o *VnicPlacementSettingsAllOf) GetPciLinkAssignmentMode() string`

GetPciLinkAssignmentMode returns the PciLinkAssignmentMode field if non-nil, zero value otherwise.

### GetPciLinkAssignmentModeOk

`func (o *VnicPlacementSettingsAllOf) GetPciLinkAssignmentModeOk() (*string, bool)`

GetPciLinkAssignmentModeOk returns a tuple with the PciLinkAssignmentMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPciLinkAssignmentMode

`func (o *VnicPlacementSettingsAllOf) SetPciLinkAssignmentMode(v string)`

SetPciLinkAssignmentMode sets PciLinkAssignmentMode field to given value.

### HasPciLinkAssignmentMode

`func (o *VnicPlacementSettingsAllOf) HasPciLinkAssignmentMode() bool`

HasPciLinkAssignmentMode returns a boolean if a field has been set.

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


