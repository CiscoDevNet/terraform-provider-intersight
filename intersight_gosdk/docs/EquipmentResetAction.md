# EquipmentResetAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "equipment.ResetAction"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "equipment.ResetAction"]
**Action** | Pointer to **string** | The reboot behavior for the Fabric Interconnect. Reboot - An action to reset the Fabric Interconnect by initiating its reboot. ForceReboot - Forces an immediate reboot of the Fabric Interconnect, overriding normal reboot validation checks. None - No reboot action should be triggered on the the Fabric Interconnect. * &#x60;None&#x60; - No action to be triggered on the Fabric Interconnect. * &#x60;Reboot&#x60; - An action to reset the Fabric Interconnect by initiating its reboot. * &#x60;ForceReboot&#x60; - An action to enforce an immediate reboot of the Fabric Interconnect regardless of existing validation checks.By default, a reset action is not allowed on an Fabric Interconnect if Domain Profile deployment, Manual Data Evacuation, or a reset action on the peer FI is already in progress. The force option allows users to override this check and perform the reset action on the FI. | [optional] [default to "None"]
**EnableFabricEvacuation** | Pointer to **bool** | The flag to enable or disable fabric evacuation before rebooting the Fabric Interconnect. | [optional] [default to true]

## Methods

### NewEquipmentResetAction

`func NewEquipmentResetAction(classId string, objectType string, ) *EquipmentResetAction`

NewEquipmentResetAction instantiates a new EquipmentResetAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEquipmentResetActionWithDefaults

`func NewEquipmentResetActionWithDefaults() *EquipmentResetAction`

NewEquipmentResetActionWithDefaults instantiates a new EquipmentResetAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *EquipmentResetAction) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *EquipmentResetAction) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *EquipmentResetAction) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *EquipmentResetAction) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *EquipmentResetAction) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *EquipmentResetAction) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAction

`func (o *EquipmentResetAction) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *EquipmentResetAction) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *EquipmentResetAction) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *EquipmentResetAction) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetEnableFabricEvacuation

`func (o *EquipmentResetAction) GetEnableFabricEvacuation() bool`

GetEnableFabricEvacuation returns the EnableFabricEvacuation field if non-nil, zero value otherwise.

### GetEnableFabricEvacuationOk

`func (o *EquipmentResetAction) GetEnableFabricEvacuationOk() (*bool, bool)`

GetEnableFabricEvacuationOk returns a tuple with the EnableFabricEvacuation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableFabricEvacuation

`func (o *EquipmentResetAction) SetEnableFabricEvacuation(v bool)`

SetEnableFabricEvacuation sets EnableFabricEvacuation field to given value.

### HasEnableFabricEvacuation

`func (o *EquipmentResetAction) HasEnableFabricEvacuation() bool`

HasEnableFabricEvacuation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


