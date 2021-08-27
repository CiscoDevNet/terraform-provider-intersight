# EquipmentPhysicalIdentity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**PhysicalDeviceRegistration** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


