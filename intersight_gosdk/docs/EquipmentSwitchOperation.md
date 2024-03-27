# EquipmentSwitchOperation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "equipment.SwitchOperation"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "equipment.SwitchOperation"]
**AdminEvacState** | Pointer to **string** | Sets evacuation state of the switch. When evacuation is enabled, data traffic flowing through this switch will be suspended for all the servers. Fabric evacuation can be enabled during any maintenance activity on the switch in order to gracefully failover data flows to the peer switch. * &#x60;Disabled&#x60; - Admin configured Disabled State. * &#x60;Enabled&#x60; - Admin configured Enabled State. | [optional] [default to "Disabled"]
**AdminLocatorLedAction** | Pointer to **string** | Action performed on the locator LED of the switch. * &#x60;None&#x60; - No operation action for the Locator Led of an equipment. * &#x60;TurnOn&#x60; - Turn on the Locator Led of an equipment. * &#x60;TurnOff&#x60; - Turn off the Locator Led of an equipment. | [optional] [default to "None"]
**AdminLocatorLedActionState** | Pointer to **string** | Defines status of action performed on AdminLocatorLedState. * &#x60;None&#x60; - Nil value when no action has been triggered by the user. * &#x60;Applied&#x60; - User configured settings are in applied state. * &#x60;Applying&#x60; - User settings are being applied on the target server. * &#x60;Failed&#x60; - User configured settings could not be applied. | [optional] [default to "None"]
**ConfigEvacState** | Pointer to **string** | Captures the status of evacuation on this switch. * &#x60;None&#x60; - Nil value when no action has been triggered by the user. * &#x60;Applied&#x60; - User configured settings are in applied state. * &#x60;Applying&#x60; - User settings are being applied on the target server. * &#x60;Failed&#x60; - User configured settings could not be applied. | [optional] [default to "None"]
**ForceEvac** | Pointer to **bool** | Evacuation is blocked by the system if it can cause a traffic outage in the domain. Select \&quot;Force Evacuation\&quot; only if system rejects the operation and you want to override that. | [optional] 
**DeviceRegistration** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 
**NetworkElement** | Pointer to [**NetworkElementRelationship**](NetworkElementRelationship.md) |  | [optional] 

## Methods

### NewEquipmentSwitchOperation

`func NewEquipmentSwitchOperation(classId string, objectType string, ) *EquipmentSwitchOperation`

NewEquipmentSwitchOperation instantiates a new EquipmentSwitchOperation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEquipmentSwitchOperationWithDefaults

`func NewEquipmentSwitchOperationWithDefaults() *EquipmentSwitchOperation`

NewEquipmentSwitchOperationWithDefaults instantiates a new EquipmentSwitchOperation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *EquipmentSwitchOperation) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *EquipmentSwitchOperation) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *EquipmentSwitchOperation) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *EquipmentSwitchOperation) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *EquipmentSwitchOperation) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *EquipmentSwitchOperation) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAdminEvacState

`func (o *EquipmentSwitchOperation) GetAdminEvacState() string`

GetAdminEvacState returns the AdminEvacState field if non-nil, zero value otherwise.

### GetAdminEvacStateOk

`func (o *EquipmentSwitchOperation) GetAdminEvacStateOk() (*string, bool)`

GetAdminEvacStateOk returns a tuple with the AdminEvacState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminEvacState

`func (o *EquipmentSwitchOperation) SetAdminEvacState(v string)`

SetAdminEvacState sets AdminEvacState field to given value.

### HasAdminEvacState

`func (o *EquipmentSwitchOperation) HasAdminEvacState() bool`

HasAdminEvacState returns a boolean if a field has been set.

### GetAdminLocatorLedAction

`func (o *EquipmentSwitchOperation) GetAdminLocatorLedAction() string`

GetAdminLocatorLedAction returns the AdminLocatorLedAction field if non-nil, zero value otherwise.

### GetAdminLocatorLedActionOk

`func (o *EquipmentSwitchOperation) GetAdminLocatorLedActionOk() (*string, bool)`

GetAdminLocatorLedActionOk returns a tuple with the AdminLocatorLedAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminLocatorLedAction

`func (o *EquipmentSwitchOperation) SetAdminLocatorLedAction(v string)`

SetAdminLocatorLedAction sets AdminLocatorLedAction field to given value.

### HasAdminLocatorLedAction

`func (o *EquipmentSwitchOperation) HasAdminLocatorLedAction() bool`

HasAdminLocatorLedAction returns a boolean if a field has been set.

### GetAdminLocatorLedActionState

`func (o *EquipmentSwitchOperation) GetAdminLocatorLedActionState() string`

GetAdminLocatorLedActionState returns the AdminLocatorLedActionState field if non-nil, zero value otherwise.

### GetAdminLocatorLedActionStateOk

`func (o *EquipmentSwitchOperation) GetAdminLocatorLedActionStateOk() (*string, bool)`

GetAdminLocatorLedActionStateOk returns a tuple with the AdminLocatorLedActionState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminLocatorLedActionState

`func (o *EquipmentSwitchOperation) SetAdminLocatorLedActionState(v string)`

SetAdminLocatorLedActionState sets AdminLocatorLedActionState field to given value.

### HasAdminLocatorLedActionState

`func (o *EquipmentSwitchOperation) HasAdminLocatorLedActionState() bool`

HasAdminLocatorLedActionState returns a boolean if a field has been set.

### GetConfigEvacState

`func (o *EquipmentSwitchOperation) GetConfigEvacState() string`

GetConfigEvacState returns the ConfigEvacState field if non-nil, zero value otherwise.

### GetConfigEvacStateOk

`func (o *EquipmentSwitchOperation) GetConfigEvacStateOk() (*string, bool)`

GetConfigEvacStateOk returns a tuple with the ConfigEvacState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigEvacState

`func (o *EquipmentSwitchOperation) SetConfigEvacState(v string)`

SetConfigEvacState sets ConfigEvacState field to given value.

### HasConfigEvacState

`func (o *EquipmentSwitchOperation) HasConfigEvacState() bool`

HasConfigEvacState returns a boolean if a field has been set.

### GetForceEvac

`func (o *EquipmentSwitchOperation) GetForceEvac() bool`

GetForceEvac returns the ForceEvac field if non-nil, zero value otherwise.

### GetForceEvacOk

`func (o *EquipmentSwitchOperation) GetForceEvacOk() (*bool, bool)`

GetForceEvacOk returns a tuple with the ForceEvac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceEvac

`func (o *EquipmentSwitchOperation) SetForceEvac(v bool)`

SetForceEvac sets ForceEvac field to given value.

### HasForceEvac

`func (o *EquipmentSwitchOperation) HasForceEvac() bool`

HasForceEvac returns a boolean if a field has been set.

### GetDeviceRegistration

`func (o *EquipmentSwitchOperation) GetDeviceRegistration() AssetDeviceRegistrationRelationship`

GetDeviceRegistration returns the DeviceRegistration field if non-nil, zero value otherwise.

### GetDeviceRegistrationOk

`func (o *EquipmentSwitchOperation) GetDeviceRegistrationOk() (*AssetDeviceRegistrationRelationship, bool)`

GetDeviceRegistrationOk returns a tuple with the DeviceRegistration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceRegistration

`func (o *EquipmentSwitchOperation) SetDeviceRegistration(v AssetDeviceRegistrationRelationship)`

SetDeviceRegistration sets DeviceRegistration field to given value.

### HasDeviceRegistration

`func (o *EquipmentSwitchOperation) HasDeviceRegistration() bool`

HasDeviceRegistration returns a boolean if a field has been set.

### GetNetworkElement

`func (o *EquipmentSwitchOperation) GetNetworkElement() NetworkElementRelationship`

GetNetworkElement returns the NetworkElement field if non-nil, zero value otherwise.

### GetNetworkElementOk

`func (o *EquipmentSwitchOperation) GetNetworkElementOk() (*NetworkElementRelationship, bool)`

GetNetworkElementOk returns a tuple with the NetworkElement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkElement

`func (o *EquipmentSwitchOperation) SetNetworkElement(v NetworkElementRelationship)`

SetNetworkElement sets NetworkElement field to given value.

### HasNetworkElement

`func (o *EquipmentSwitchOperation) HasNetworkElement() bool`

HasNetworkElement returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


