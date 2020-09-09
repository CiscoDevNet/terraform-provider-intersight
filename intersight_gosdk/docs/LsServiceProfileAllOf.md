# LsServiceProfileAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssignState** | Pointer to **string** | Assignment state of the service profile. | [optional] [readonly] 
**AssocState** | Pointer to **string** | Association state of the service profile. | [optional] [readonly] 
**AssociatedServer** | Pointer to **string** | Server to which the UCS Manager service profile is associated to. | [optional] [readonly] 
**ConfigState** | Pointer to **string** | Configuration state of the service profile. | [optional] [readonly] 
**Name** | Pointer to **string** | Name of the UCS Manager service profile. | [optional] [readonly] 
**OperState** | Pointer to **string** | Operational state of the service profile. | [optional] [readonly] 
**InventoryDeviceInfo** | Pointer to [**InventoryDeviceInfoRelationship**](inventory.DeviceInfo.Relationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 

## Methods

### NewLsServiceProfileAllOf

`func NewLsServiceProfileAllOf() *LsServiceProfileAllOf`

NewLsServiceProfileAllOf instantiates a new LsServiceProfileAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLsServiceProfileAllOfWithDefaults

`func NewLsServiceProfileAllOfWithDefaults() *LsServiceProfileAllOf`

NewLsServiceProfileAllOfWithDefaults instantiates a new LsServiceProfileAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssignState

`func (o *LsServiceProfileAllOf) GetAssignState() string`

GetAssignState returns the AssignState field if non-nil, zero value otherwise.

### GetAssignStateOk

`func (o *LsServiceProfileAllOf) GetAssignStateOk() (*string, bool)`

GetAssignStateOk returns a tuple with the AssignState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignState

`func (o *LsServiceProfileAllOf) SetAssignState(v string)`

SetAssignState sets AssignState field to given value.

### HasAssignState

`func (o *LsServiceProfileAllOf) HasAssignState() bool`

HasAssignState returns a boolean if a field has been set.

### GetAssocState

`func (o *LsServiceProfileAllOf) GetAssocState() string`

GetAssocState returns the AssocState field if non-nil, zero value otherwise.

### GetAssocStateOk

`func (o *LsServiceProfileAllOf) GetAssocStateOk() (*string, bool)`

GetAssocStateOk returns a tuple with the AssocState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssocState

`func (o *LsServiceProfileAllOf) SetAssocState(v string)`

SetAssocState sets AssocState field to given value.

### HasAssocState

`func (o *LsServiceProfileAllOf) HasAssocState() bool`

HasAssocState returns a boolean if a field has been set.

### GetAssociatedServer

`func (o *LsServiceProfileAllOf) GetAssociatedServer() string`

GetAssociatedServer returns the AssociatedServer field if non-nil, zero value otherwise.

### GetAssociatedServerOk

`func (o *LsServiceProfileAllOf) GetAssociatedServerOk() (*string, bool)`

GetAssociatedServerOk returns a tuple with the AssociatedServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociatedServer

`func (o *LsServiceProfileAllOf) SetAssociatedServer(v string)`

SetAssociatedServer sets AssociatedServer field to given value.

### HasAssociatedServer

`func (o *LsServiceProfileAllOf) HasAssociatedServer() bool`

HasAssociatedServer returns a boolean if a field has been set.

### GetConfigState

`func (o *LsServiceProfileAllOf) GetConfigState() string`

GetConfigState returns the ConfigState field if non-nil, zero value otherwise.

### GetConfigStateOk

`func (o *LsServiceProfileAllOf) GetConfigStateOk() (*string, bool)`

GetConfigStateOk returns a tuple with the ConfigState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigState

`func (o *LsServiceProfileAllOf) SetConfigState(v string)`

SetConfigState sets ConfigState field to given value.

### HasConfigState

`func (o *LsServiceProfileAllOf) HasConfigState() bool`

HasConfigState returns a boolean if a field has been set.

### GetName

`func (o *LsServiceProfileAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LsServiceProfileAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LsServiceProfileAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *LsServiceProfileAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOperState

`func (o *LsServiceProfileAllOf) GetOperState() string`

GetOperState returns the OperState field if non-nil, zero value otherwise.

### GetOperStateOk

`func (o *LsServiceProfileAllOf) GetOperStateOk() (*string, bool)`

GetOperStateOk returns a tuple with the OperState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperState

`func (o *LsServiceProfileAllOf) SetOperState(v string)`

SetOperState sets OperState field to given value.

### HasOperState

`func (o *LsServiceProfileAllOf) HasOperState() bool`

HasOperState returns a boolean if a field has been set.

### GetInventoryDeviceInfo

`func (o *LsServiceProfileAllOf) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship`

GetInventoryDeviceInfo returns the InventoryDeviceInfo field if non-nil, zero value otherwise.

### GetInventoryDeviceInfoOk

`func (o *LsServiceProfileAllOf) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool)`

GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDeviceInfo

`func (o *LsServiceProfileAllOf) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship)`

SetInventoryDeviceInfo sets InventoryDeviceInfo field to given value.

### HasInventoryDeviceInfo

`func (o *LsServiceProfileAllOf) HasInventoryDeviceInfo() bool`

HasInventoryDeviceInfo returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *LsServiceProfileAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *LsServiceProfileAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *LsServiceProfileAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *LsServiceProfileAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


