# FirmwareUpgradeAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExcludeComponentList** | Pointer to **[]string** |  | [optional] 
**Device** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 
**Server** | Pointer to [**ComputePhysicalRelationship**](compute.Physical.Relationship.md) |  | [optional] 

## Methods

### NewFirmwareUpgradeAllOf

`func NewFirmwareUpgradeAllOf() *FirmwareUpgradeAllOf`

NewFirmwareUpgradeAllOf instantiates a new FirmwareUpgradeAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirmwareUpgradeAllOfWithDefaults

`func NewFirmwareUpgradeAllOfWithDefaults() *FirmwareUpgradeAllOf`

NewFirmwareUpgradeAllOfWithDefaults instantiates a new FirmwareUpgradeAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExcludeComponentList

`func (o *FirmwareUpgradeAllOf) GetExcludeComponentList() []string`

GetExcludeComponentList returns the ExcludeComponentList field if non-nil, zero value otherwise.

### GetExcludeComponentListOk

`func (o *FirmwareUpgradeAllOf) GetExcludeComponentListOk() (*[]string, bool)`

GetExcludeComponentListOk returns a tuple with the ExcludeComponentList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeComponentList

`func (o *FirmwareUpgradeAllOf) SetExcludeComponentList(v []string)`

SetExcludeComponentList sets ExcludeComponentList field to given value.

### HasExcludeComponentList

`func (o *FirmwareUpgradeAllOf) HasExcludeComponentList() bool`

HasExcludeComponentList returns a boolean if a field has been set.

### GetDevice

`func (o *FirmwareUpgradeAllOf) GetDevice() AssetDeviceRegistrationRelationship`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *FirmwareUpgradeAllOf) GetDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *FirmwareUpgradeAllOf) SetDevice(v AssetDeviceRegistrationRelationship)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *FirmwareUpgradeAllOf) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### GetServer

`func (o *FirmwareUpgradeAllOf) GetServer() ComputePhysicalRelationship`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *FirmwareUpgradeAllOf) GetServerOk() (*ComputePhysicalRelationship, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *FirmwareUpgradeAllOf) SetServer(v ComputePhysicalRelationship)`

SetServer sets Server field to given value.

### HasServer

`func (o *FirmwareUpgradeAllOf) HasServer() bool`

HasServer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


