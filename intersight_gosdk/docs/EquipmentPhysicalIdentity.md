# EquipmentPhysicalIdentity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**LastDiscoveryTriggered** | Pointer to **string** | Denotes the type of discovery that was most recently triggered on this server. * &#x60;Unknown&#x60; - The last discovery type is unknown. * &#x60;Deep&#x60; - The last discovery triggered is deep. * &#x60;Shallow&#x60; - The last discovery triggered is shallow. | [optional] [readonly] [default to "Unknown"]
**ResetToDefault** | Pointer to **bool** | Specifies whether device configurations need to be reset to default upon first-time discovery or recommission of a server. | [optional] [readonly] 
**PhysicalDeviceRegistration** | Pointer to [**NullableAssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewEquipmentPhysicalIdentity

`func NewEquipmentPhysicalIdentity(classId string, objectType string, ) *EquipmentPhysicalIdentity`

NewEquipmentPhysicalIdentity instantiates a new EquipmentPhysicalIdentity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEquipmentPhysicalIdentityWithDefaults

`func NewEquipmentPhysicalIdentityWithDefaults() *EquipmentPhysicalIdentity`

NewEquipmentPhysicalIdentityWithDefaults instantiates a new EquipmentPhysicalIdentity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *EquipmentPhysicalIdentity) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *EquipmentPhysicalIdentity) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *EquipmentPhysicalIdentity) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *EquipmentPhysicalIdentity) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *EquipmentPhysicalIdentity) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *EquipmentPhysicalIdentity) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetLastDiscoveryTriggered

`func (o *EquipmentPhysicalIdentity) GetLastDiscoveryTriggered() string`

GetLastDiscoveryTriggered returns the LastDiscoveryTriggered field if non-nil, zero value otherwise.

### GetLastDiscoveryTriggeredOk

`func (o *EquipmentPhysicalIdentity) GetLastDiscoveryTriggeredOk() (*string, bool)`

GetLastDiscoveryTriggeredOk returns a tuple with the LastDiscoveryTriggered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastDiscoveryTriggered

`func (o *EquipmentPhysicalIdentity) SetLastDiscoveryTriggered(v string)`

SetLastDiscoveryTriggered sets LastDiscoveryTriggered field to given value.

### HasLastDiscoveryTriggered

`func (o *EquipmentPhysicalIdentity) HasLastDiscoveryTriggered() bool`

HasLastDiscoveryTriggered returns a boolean if a field has been set.

### GetResetToDefault

`func (o *EquipmentPhysicalIdentity) GetResetToDefault() bool`

GetResetToDefault returns the ResetToDefault field if non-nil, zero value otherwise.

### GetResetToDefaultOk

`func (o *EquipmentPhysicalIdentity) GetResetToDefaultOk() (*bool, bool)`

GetResetToDefaultOk returns a tuple with the ResetToDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetToDefault

`func (o *EquipmentPhysicalIdentity) SetResetToDefault(v bool)`

SetResetToDefault sets ResetToDefault field to given value.

### HasResetToDefault

`func (o *EquipmentPhysicalIdentity) HasResetToDefault() bool`

HasResetToDefault returns a boolean if a field has been set.

### GetPhysicalDeviceRegistration

`func (o *EquipmentPhysicalIdentity) GetPhysicalDeviceRegistration() AssetDeviceRegistrationRelationship`

GetPhysicalDeviceRegistration returns the PhysicalDeviceRegistration field if non-nil, zero value otherwise.

### GetPhysicalDeviceRegistrationOk

`func (o *EquipmentPhysicalIdentity) GetPhysicalDeviceRegistrationOk() (*AssetDeviceRegistrationRelationship, bool)`

GetPhysicalDeviceRegistrationOk returns a tuple with the PhysicalDeviceRegistration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalDeviceRegistration

`func (o *EquipmentPhysicalIdentity) SetPhysicalDeviceRegistration(v AssetDeviceRegistrationRelationship)`

SetPhysicalDeviceRegistration sets PhysicalDeviceRegistration field to given value.

### HasPhysicalDeviceRegistration

`func (o *EquipmentPhysicalIdentity) HasPhysicalDeviceRegistration() bool`

HasPhysicalDeviceRegistration returns a boolean if a field has been set.

### SetPhysicalDeviceRegistrationNil

`func (o *EquipmentPhysicalIdentity) SetPhysicalDeviceRegistrationNil(b bool)`

 SetPhysicalDeviceRegistrationNil sets the value for PhysicalDeviceRegistration to be an explicit nil

### UnsetPhysicalDeviceRegistration
`func (o *EquipmentPhysicalIdentity) UnsetPhysicalDeviceRegistration()`

UnsetPhysicalDeviceRegistration ensures that no value is present for PhysicalDeviceRegistration, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


