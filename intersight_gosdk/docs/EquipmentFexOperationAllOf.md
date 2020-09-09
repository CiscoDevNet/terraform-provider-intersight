# EquipmentFexOperationAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdminLocatorLedAction** | Pointer to **string** | Action performed on the locator LED for a FEX. * &#x60;None&#x60; - No operation action for the Locator Led of an equipment. * &#x60;TurnOn&#x60; - Turn on the Locator Led of an equipment. * &#x60;TurnOff&#x60; - Turn off the Locator Led of an equipment. | [optional] [default to "None"]
**AdminLocatorLedActionState** | Pointer to **string** | Defines status of action performed on AdminLocatorLedState. * &#x60;None&#x60; - Nil value when no action has been triggered by the user. * &#x60;Applied&#x60; - User configured settings are in applied state. * &#x60;Applying&#x60; - User settings are being applied on the target server. * &#x60;Failed&#x60; - User configured settings could not be applied. | [optional] [default to "None"]
**DeviceRegistration** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 
**Fex** | Pointer to [**EquipmentFexRelationship**](equipment.Fex.Relationship.md) |  | [optional] 

## Methods

### NewEquipmentFexOperationAllOf

`func NewEquipmentFexOperationAllOf() *EquipmentFexOperationAllOf`

NewEquipmentFexOperationAllOf instantiates a new EquipmentFexOperationAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEquipmentFexOperationAllOfWithDefaults

`func NewEquipmentFexOperationAllOfWithDefaults() *EquipmentFexOperationAllOf`

NewEquipmentFexOperationAllOfWithDefaults instantiates a new EquipmentFexOperationAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdminLocatorLedAction

`func (o *EquipmentFexOperationAllOf) GetAdminLocatorLedAction() string`

GetAdminLocatorLedAction returns the AdminLocatorLedAction field if non-nil, zero value otherwise.

### GetAdminLocatorLedActionOk

`func (o *EquipmentFexOperationAllOf) GetAdminLocatorLedActionOk() (*string, bool)`

GetAdminLocatorLedActionOk returns a tuple with the AdminLocatorLedAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminLocatorLedAction

`func (o *EquipmentFexOperationAllOf) SetAdminLocatorLedAction(v string)`

SetAdminLocatorLedAction sets AdminLocatorLedAction field to given value.

### HasAdminLocatorLedAction

`func (o *EquipmentFexOperationAllOf) HasAdminLocatorLedAction() bool`

HasAdminLocatorLedAction returns a boolean if a field has been set.

### GetAdminLocatorLedActionState

`func (o *EquipmentFexOperationAllOf) GetAdminLocatorLedActionState() string`

GetAdminLocatorLedActionState returns the AdminLocatorLedActionState field if non-nil, zero value otherwise.

### GetAdminLocatorLedActionStateOk

`func (o *EquipmentFexOperationAllOf) GetAdminLocatorLedActionStateOk() (*string, bool)`

GetAdminLocatorLedActionStateOk returns a tuple with the AdminLocatorLedActionState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminLocatorLedActionState

`func (o *EquipmentFexOperationAllOf) SetAdminLocatorLedActionState(v string)`

SetAdminLocatorLedActionState sets AdminLocatorLedActionState field to given value.

### HasAdminLocatorLedActionState

`func (o *EquipmentFexOperationAllOf) HasAdminLocatorLedActionState() bool`

HasAdminLocatorLedActionState returns a boolean if a field has been set.

### GetDeviceRegistration

`func (o *EquipmentFexOperationAllOf) GetDeviceRegistration() AssetDeviceRegistrationRelationship`

GetDeviceRegistration returns the DeviceRegistration field if non-nil, zero value otherwise.

### GetDeviceRegistrationOk

`func (o *EquipmentFexOperationAllOf) GetDeviceRegistrationOk() (*AssetDeviceRegistrationRelationship, bool)`

GetDeviceRegistrationOk returns a tuple with the DeviceRegistration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceRegistration

`func (o *EquipmentFexOperationAllOf) SetDeviceRegistration(v AssetDeviceRegistrationRelationship)`

SetDeviceRegistration sets DeviceRegistration field to given value.

### HasDeviceRegistration

`func (o *EquipmentFexOperationAllOf) HasDeviceRegistration() bool`

HasDeviceRegistration returns a boolean if a field has been set.

### GetFex

`func (o *EquipmentFexOperationAllOf) GetFex() EquipmentFexRelationship`

GetFex returns the Fex field if non-nil, zero value otherwise.

### GetFexOk

`func (o *EquipmentFexOperationAllOf) GetFexOk() (*EquipmentFexRelationship, bool)`

GetFexOk returns a tuple with the Fex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFex

`func (o *EquipmentFexOperationAllOf) SetFex(v EquipmentFexRelationship)`

SetFex sets Fex field to given value.

### HasFex

`func (o *EquipmentFexOperationAllOf) HasFex() bool`

HasFex returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


