# EquipmentIoCardIdentity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "equipment.IoCardIdentity"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "equipment.IoCardIdentity"]
**IoCardMoid** | Pointer to **string** | MO Reference to equipmentIoCard MO in inventory service. | [optional] [readonly] 
**Lifecycle** | Pointer to **string** | IO Card inventory lifecycle status. * &#x60;Unknown&#x60; - Default lifecycle state of a iocard. This should be an initial state when no state is defined for a specific iocard. * &#x60;Decommissioned&#x60; - Lifecycle state is set to this value after the chassis is successfully decommissioned. * &#x60;DiscoveryInProgress&#x60; - Lifecycle state is set to this value after the successful start of the iocard connection discovery process. * &#x60;DiscoveryFailed&#x60; - Lifecycle state is set to this value after the iocard connection discovery has failed. * &#x60;DiscoveryCompleted&#x60; - Lifecycle state is set to this value after the connection discovery of the iocard is completed successfully. * &#x60;None&#x60; - Lifecycle state is set to this value before the start of connection discovery and inventory collection process for a iocard. * &#x60;InventoryCompleted&#x60; - Lifecycle state is set to this value after the chassis inventory collection process is completed for a specific iocard. * &#x60;InventoryInProgress&#x60; - Lifecycle state is set to this value after successful  start of the chassis inventory collection process for a specific iocard. * &#x60;InventoryFailed&#x60; - Lifecycle state is set to this value after the chassis inventory collection process failed for a iocard. | [optional] [readonly] [default to "Unknown"]
**Model** | Pointer to **string** | IO Card or intelligent fabric module model. | [optional] [readonly] 
**ModuleId** | Pointer to **int64** | IOM/MUX Module ID connected to the FI. | [optional] [readonly] 
**NetworkElementMoid** | Pointer to **string** | MO Reference to networkElement MO in inventory service. | [optional] [readonly] 
**Serial** | Pointer to **string** | IO Card or intelligent fabric module serial number. | [optional] [readonly] 
**SwitchId** | Pointer to **int64** | Identifier of the Switch where the IOM is connected. ID can be either 1 or 2, equivalent to A or B. | [optional] [readonly] 
**Vendor** | Pointer to **string** | IO Card or intelligent fabric module vendor. | [optional] [readonly] 

## Methods

### NewEquipmentIoCardIdentity

`func NewEquipmentIoCardIdentity(classId string, objectType string, ) *EquipmentIoCardIdentity`

NewEquipmentIoCardIdentity instantiates a new EquipmentIoCardIdentity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEquipmentIoCardIdentityWithDefaults

`func NewEquipmentIoCardIdentityWithDefaults() *EquipmentIoCardIdentity`

NewEquipmentIoCardIdentityWithDefaults instantiates a new EquipmentIoCardIdentity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *EquipmentIoCardIdentity) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *EquipmentIoCardIdentity) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *EquipmentIoCardIdentity) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *EquipmentIoCardIdentity) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *EquipmentIoCardIdentity) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *EquipmentIoCardIdentity) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetIoCardMoid

`func (o *EquipmentIoCardIdentity) GetIoCardMoid() string`

GetIoCardMoid returns the IoCardMoid field if non-nil, zero value otherwise.

### GetIoCardMoidOk

`func (o *EquipmentIoCardIdentity) GetIoCardMoidOk() (*string, bool)`

GetIoCardMoidOk returns a tuple with the IoCardMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIoCardMoid

`func (o *EquipmentIoCardIdentity) SetIoCardMoid(v string)`

SetIoCardMoid sets IoCardMoid field to given value.

### HasIoCardMoid

`func (o *EquipmentIoCardIdentity) HasIoCardMoid() bool`

HasIoCardMoid returns a boolean if a field has been set.

### GetLifecycle

`func (o *EquipmentIoCardIdentity) GetLifecycle() string`

GetLifecycle returns the Lifecycle field if non-nil, zero value otherwise.

### GetLifecycleOk

`func (o *EquipmentIoCardIdentity) GetLifecycleOk() (*string, bool)`

GetLifecycleOk returns a tuple with the Lifecycle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifecycle

`func (o *EquipmentIoCardIdentity) SetLifecycle(v string)`

SetLifecycle sets Lifecycle field to given value.

### HasLifecycle

`func (o *EquipmentIoCardIdentity) HasLifecycle() bool`

HasLifecycle returns a boolean if a field has been set.

### GetModel

`func (o *EquipmentIoCardIdentity) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *EquipmentIoCardIdentity) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *EquipmentIoCardIdentity) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *EquipmentIoCardIdentity) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetModuleId

`func (o *EquipmentIoCardIdentity) GetModuleId() int64`

GetModuleId returns the ModuleId field if non-nil, zero value otherwise.

### GetModuleIdOk

`func (o *EquipmentIoCardIdentity) GetModuleIdOk() (*int64, bool)`

GetModuleIdOk returns a tuple with the ModuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModuleId

`func (o *EquipmentIoCardIdentity) SetModuleId(v int64)`

SetModuleId sets ModuleId field to given value.

### HasModuleId

`func (o *EquipmentIoCardIdentity) HasModuleId() bool`

HasModuleId returns a boolean if a field has been set.

### GetNetworkElementMoid

`func (o *EquipmentIoCardIdentity) GetNetworkElementMoid() string`

GetNetworkElementMoid returns the NetworkElementMoid field if non-nil, zero value otherwise.

### GetNetworkElementMoidOk

`func (o *EquipmentIoCardIdentity) GetNetworkElementMoidOk() (*string, bool)`

GetNetworkElementMoidOk returns a tuple with the NetworkElementMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkElementMoid

`func (o *EquipmentIoCardIdentity) SetNetworkElementMoid(v string)`

SetNetworkElementMoid sets NetworkElementMoid field to given value.

### HasNetworkElementMoid

`func (o *EquipmentIoCardIdentity) HasNetworkElementMoid() bool`

HasNetworkElementMoid returns a boolean if a field has been set.

### GetSerial

`func (o *EquipmentIoCardIdentity) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *EquipmentIoCardIdentity) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *EquipmentIoCardIdentity) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *EquipmentIoCardIdentity) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetSwitchId

`func (o *EquipmentIoCardIdentity) GetSwitchId() int64`

GetSwitchId returns the SwitchId field if non-nil, zero value otherwise.

### GetSwitchIdOk

`func (o *EquipmentIoCardIdentity) GetSwitchIdOk() (*int64, bool)`

GetSwitchIdOk returns a tuple with the SwitchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchId

`func (o *EquipmentIoCardIdentity) SetSwitchId(v int64)`

SetSwitchId sets SwitchId field to given value.

### HasSwitchId

`func (o *EquipmentIoCardIdentity) HasSwitchId() bool`

HasSwitchId returns a boolean if a field has been set.

### GetVendor

`func (o *EquipmentIoCardIdentity) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *EquipmentIoCardIdentity) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *EquipmentIoCardIdentity) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *EquipmentIoCardIdentity) HasVendor() bool`

HasVendor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


