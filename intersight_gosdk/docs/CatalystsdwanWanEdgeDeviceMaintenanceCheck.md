# CatalystsdwanWanEdgeDeviceMaintenanceCheck

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "catalystsdwan.WanEdgeDeviceMaintenanceCheck"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "catalystsdwan.WanEdgeDeviceMaintenanceCheck"]
**TargetChassis** | Pointer to **[]string** |  | [optional] 
**User** | Pointer to [**NullableIamUserRelationship**](IamUserRelationship.md) |  | [optional] 

## Methods

### NewCatalystsdwanWanEdgeDeviceMaintenanceCheck

`func NewCatalystsdwanWanEdgeDeviceMaintenanceCheck(classId string, objectType string, ) *CatalystsdwanWanEdgeDeviceMaintenanceCheck`

NewCatalystsdwanWanEdgeDeviceMaintenanceCheck instantiates a new CatalystsdwanWanEdgeDeviceMaintenanceCheck object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalystsdwanWanEdgeDeviceMaintenanceCheckWithDefaults

`func NewCatalystsdwanWanEdgeDeviceMaintenanceCheckWithDefaults() *CatalystsdwanWanEdgeDeviceMaintenanceCheck`

NewCatalystsdwanWanEdgeDeviceMaintenanceCheckWithDefaults instantiates a new CatalystsdwanWanEdgeDeviceMaintenanceCheck object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CatalystsdwanWanEdgeDeviceMaintenanceCheck) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CatalystsdwanWanEdgeDeviceMaintenanceCheck) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CatalystsdwanWanEdgeDeviceMaintenanceCheck) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CatalystsdwanWanEdgeDeviceMaintenanceCheck) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CatalystsdwanWanEdgeDeviceMaintenanceCheck) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CatalystsdwanWanEdgeDeviceMaintenanceCheck) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetTargetChassis

`func (o *CatalystsdwanWanEdgeDeviceMaintenanceCheck) GetTargetChassis() []string`

GetTargetChassis returns the TargetChassis field if non-nil, zero value otherwise.

### GetTargetChassisOk

`func (o *CatalystsdwanWanEdgeDeviceMaintenanceCheck) GetTargetChassisOk() (*[]string, bool)`

GetTargetChassisOk returns a tuple with the TargetChassis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetChassis

`func (o *CatalystsdwanWanEdgeDeviceMaintenanceCheck) SetTargetChassis(v []string)`

SetTargetChassis sets TargetChassis field to given value.

### HasTargetChassis

`func (o *CatalystsdwanWanEdgeDeviceMaintenanceCheck) HasTargetChassis() bool`

HasTargetChassis returns a boolean if a field has been set.

### SetTargetChassisNil

`func (o *CatalystsdwanWanEdgeDeviceMaintenanceCheck) SetTargetChassisNil(b bool)`

 SetTargetChassisNil sets the value for TargetChassis to be an explicit nil

### UnsetTargetChassis
`func (o *CatalystsdwanWanEdgeDeviceMaintenanceCheck) UnsetTargetChassis()`

UnsetTargetChassis ensures that no value is present for TargetChassis, not even an explicit nil
### GetUser

`func (o *CatalystsdwanWanEdgeDeviceMaintenanceCheck) GetUser() IamUserRelationship`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *CatalystsdwanWanEdgeDeviceMaintenanceCheck) GetUserOk() (*IamUserRelationship, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *CatalystsdwanWanEdgeDeviceMaintenanceCheck) SetUser(v IamUserRelationship)`

SetUser sets User field to given value.

### HasUser

`func (o *CatalystsdwanWanEdgeDeviceMaintenanceCheck) HasUser() bool`

HasUser returns a boolean if a field has been set.

### SetUserNil

`func (o *CatalystsdwanWanEdgeDeviceMaintenanceCheck) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *CatalystsdwanWanEdgeDeviceMaintenanceCheck) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


