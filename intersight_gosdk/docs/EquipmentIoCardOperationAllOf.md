# EquipmentIoCardOperationAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdminPowerState** | Pointer to **string** | User configured power state of the iomodule. * &#x60;None&#x60; - Placeholder default value for iom power state property. * &#x60;Reboot&#x60; - IO Module reboot state property value. | [optional] [default to "None"]
**ConfigState** | Pointer to **string** | The configured state of these settings in the target chassis. The value is any one of Applied, Applying, Failed. Applied - This state denotes that the settings are applied successfully in the target chassis iomodule. Applying - This state denotes that the settings are being applied in the target chassis iomodule. Failed - This state denotes that the settings could not be applied in the target chassis iomodule. * &#x60;None&#x60; - Nil value when no action has been triggered by the user. * &#x60;Applied&#x60; - User configured settings are in applied state. * &#x60;Applying&#x60; - User settings are being applied on the target server. * &#x60;Failed&#x60; - User configured settings could not be applied. | [optional] [readonly] [default to "None"]
**DeviceRegistration** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 
**IoCard** | Pointer to [**EquipmentIoCardRelationship**](equipment.IoCard.Relationship.md) |  | [optional] 

## Methods

### NewEquipmentIoCardOperationAllOf

`func NewEquipmentIoCardOperationAllOf() *EquipmentIoCardOperationAllOf`

NewEquipmentIoCardOperationAllOf instantiates a new EquipmentIoCardOperationAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEquipmentIoCardOperationAllOfWithDefaults

`func NewEquipmentIoCardOperationAllOfWithDefaults() *EquipmentIoCardOperationAllOf`

NewEquipmentIoCardOperationAllOfWithDefaults instantiates a new EquipmentIoCardOperationAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdminPowerState

`func (o *EquipmentIoCardOperationAllOf) GetAdminPowerState() string`

GetAdminPowerState returns the AdminPowerState field if non-nil, zero value otherwise.

### GetAdminPowerStateOk

`func (o *EquipmentIoCardOperationAllOf) GetAdminPowerStateOk() (*string, bool)`

GetAdminPowerStateOk returns a tuple with the AdminPowerState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminPowerState

`func (o *EquipmentIoCardOperationAllOf) SetAdminPowerState(v string)`

SetAdminPowerState sets AdminPowerState field to given value.

### HasAdminPowerState

`func (o *EquipmentIoCardOperationAllOf) HasAdminPowerState() bool`

HasAdminPowerState returns a boolean if a field has been set.

### GetConfigState

`func (o *EquipmentIoCardOperationAllOf) GetConfigState() string`

GetConfigState returns the ConfigState field if non-nil, zero value otherwise.

### GetConfigStateOk

`func (o *EquipmentIoCardOperationAllOf) GetConfigStateOk() (*string, bool)`

GetConfigStateOk returns a tuple with the ConfigState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigState

`func (o *EquipmentIoCardOperationAllOf) SetConfigState(v string)`

SetConfigState sets ConfigState field to given value.

### HasConfigState

`func (o *EquipmentIoCardOperationAllOf) HasConfigState() bool`

HasConfigState returns a boolean if a field has been set.

### GetDeviceRegistration

`func (o *EquipmentIoCardOperationAllOf) GetDeviceRegistration() AssetDeviceRegistrationRelationship`

GetDeviceRegistration returns the DeviceRegistration field if non-nil, zero value otherwise.

### GetDeviceRegistrationOk

`func (o *EquipmentIoCardOperationAllOf) GetDeviceRegistrationOk() (*AssetDeviceRegistrationRelationship, bool)`

GetDeviceRegistrationOk returns a tuple with the DeviceRegistration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceRegistration

`func (o *EquipmentIoCardOperationAllOf) SetDeviceRegistration(v AssetDeviceRegistrationRelationship)`

SetDeviceRegistration sets DeviceRegistration field to given value.

### HasDeviceRegistration

`func (o *EquipmentIoCardOperationAllOf) HasDeviceRegistration() bool`

HasDeviceRegistration returns a boolean if a field has been set.

### GetIoCard

`func (o *EquipmentIoCardOperationAllOf) GetIoCard() EquipmentIoCardRelationship`

GetIoCard returns the IoCard field if non-nil, zero value otherwise.

### GetIoCardOk

`func (o *EquipmentIoCardOperationAllOf) GetIoCardOk() (*EquipmentIoCardRelationship, bool)`

GetIoCardOk returns a tuple with the IoCard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIoCard

`func (o *EquipmentIoCardOperationAllOf) SetIoCard(v EquipmentIoCardRelationship)`

SetIoCard sets IoCard field to given value.

### HasIoCard

`func (o *EquipmentIoCardOperationAllOf) HasIoCard() bool`

HasIoCard returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


