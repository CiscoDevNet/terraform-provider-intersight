# PolicyinventoryAbstractDeviceInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | [default to "inventory.DeviceInfo"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | [default to "inventory.DeviceInfo"]
**ConfigState** | Pointer to **string** | Configuration state of server profile config context. | [optional] [readonly] 
**ControlAction** | Pointer to **string** | Control action of server profile config context. | [optional] [readonly] 
**ErrorState** | Pointer to **string** | Error state of server profile config context. | [optional] [readonly] 
**JobInfo** | Pointer to [**[]PolicyinventoryJobInfo**](PolicyinventoryJobInfo.md) |  | [optional] 
**OperState** | Pointer to **string** | Operational state of server profile config context. | [optional] [readonly] 
**ProfileMoId** | Pointer to **string** | Server profile MO ID of the server. | [optional] [readonly] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewPolicyinventoryAbstractDeviceInfo

`func NewPolicyinventoryAbstractDeviceInfo(classId string, objectType string, ) *PolicyinventoryAbstractDeviceInfo`

NewPolicyinventoryAbstractDeviceInfo instantiates a new PolicyinventoryAbstractDeviceInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyinventoryAbstractDeviceInfoWithDefaults

`func NewPolicyinventoryAbstractDeviceInfoWithDefaults() *PolicyinventoryAbstractDeviceInfo`

NewPolicyinventoryAbstractDeviceInfoWithDefaults instantiates a new PolicyinventoryAbstractDeviceInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *PolicyinventoryAbstractDeviceInfo) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *PolicyinventoryAbstractDeviceInfo) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *PolicyinventoryAbstractDeviceInfo) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *PolicyinventoryAbstractDeviceInfo) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *PolicyinventoryAbstractDeviceInfo) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *PolicyinventoryAbstractDeviceInfo) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetConfigState

`func (o *PolicyinventoryAbstractDeviceInfo) GetConfigState() string`

GetConfigState returns the ConfigState field if non-nil, zero value otherwise.

### GetConfigStateOk

`func (o *PolicyinventoryAbstractDeviceInfo) GetConfigStateOk() (*string, bool)`

GetConfigStateOk returns a tuple with the ConfigState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigState

`func (o *PolicyinventoryAbstractDeviceInfo) SetConfigState(v string)`

SetConfigState sets ConfigState field to given value.

### HasConfigState

`func (o *PolicyinventoryAbstractDeviceInfo) HasConfigState() bool`

HasConfigState returns a boolean if a field has been set.

### GetControlAction

`func (o *PolicyinventoryAbstractDeviceInfo) GetControlAction() string`

GetControlAction returns the ControlAction field if non-nil, zero value otherwise.

### GetControlActionOk

`func (o *PolicyinventoryAbstractDeviceInfo) GetControlActionOk() (*string, bool)`

GetControlActionOk returns a tuple with the ControlAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlAction

`func (o *PolicyinventoryAbstractDeviceInfo) SetControlAction(v string)`

SetControlAction sets ControlAction field to given value.

### HasControlAction

`func (o *PolicyinventoryAbstractDeviceInfo) HasControlAction() bool`

HasControlAction returns a boolean if a field has been set.

### GetErrorState

`func (o *PolicyinventoryAbstractDeviceInfo) GetErrorState() string`

GetErrorState returns the ErrorState field if non-nil, zero value otherwise.

### GetErrorStateOk

`func (o *PolicyinventoryAbstractDeviceInfo) GetErrorStateOk() (*string, bool)`

GetErrorStateOk returns a tuple with the ErrorState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorState

`func (o *PolicyinventoryAbstractDeviceInfo) SetErrorState(v string)`

SetErrorState sets ErrorState field to given value.

### HasErrorState

`func (o *PolicyinventoryAbstractDeviceInfo) HasErrorState() bool`

HasErrorState returns a boolean if a field has been set.

### GetJobInfo

`func (o *PolicyinventoryAbstractDeviceInfo) GetJobInfo() []PolicyinventoryJobInfo`

GetJobInfo returns the JobInfo field if non-nil, zero value otherwise.

### GetJobInfoOk

`func (o *PolicyinventoryAbstractDeviceInfo) GetJobInfoOk() (*[]PolicyinventoryJobInfo, bool)`

GetJobInfoOk returns a tuple with the JobInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobInfo

`func (o *PolicyinventoryAbstractDeviceInfo) SetJobInfo(v []PolicyinventoryJobInfo)`

SetJobInfo sets JobInfo field to given value.

### HasJobInfo

`func (o *PolicyinventoryAbstractDeviceInfo) HasJobInfo() bool`

HasJobInfo returns a boolean if a field has been set.

### SetJobInfoNil

`func (o *PolicyinventoryAbstractDeviceInfo) SetJobInfoNil(b bool)`

 SetJobInfoNil sets the value for JobInfo to be an explicit nil

### UnsetJobInfo
`func (o *PolicyinventoryAbstractDeviceInfo) UnsetJobInfo()`

UnsetJobInfo ensures that no value is present for JobInfo, not even an explicit nil
### GetOperState

`func (o *PolicyinventoryAbstractDeviceInfo) GetOperState() string`

GetOperState returns the OperState field if non-nil, zero value otherwise.

### GetOperStateOk

`func (o *PolicyinventoryAbstractDeviceInfo) GetOperStateOk() (*string, bool)`

GetOperStateOk returns a tuple with the OperState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperState

`func (o *PolicyinventoryAbstractDeviceInfo) SetOperState(v string)`

SetOperState sets OperState field to given value.

### HasOperState

`func (o *PolicyinventoryAbstractDeviceInfo) HasOperState() bool`

HasOperState returns a boolean if a field has been set.

### GetProfileMoId

`func (o *PolicyinventoryAbstractDeviceInfo) GetProfileMoId() string`

GetProfileMoId returns the ProfileMoId field if non-nil, zero value otherwise.

### GetProfileMoIdOk

`func (o *PolicyinventoryAbstractDeviceInfo) GetProfileMoIdOk() (*string, bool)`

GetProfileMoIdOk returns a tuple with the ProfileMoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileMoId

`func (o *PolicyinventoryAbstractDeviceInfo) SetProfileMoId(v string)`

SetProfileMoId sets ProfileMoId field to given value.

### HasProfileMoId

`func (o *PolicyinventoryAbstractDeviceInfo) HasProfileMoId() bool`

HasProfileMoId returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *PolicyinventoryAbstractDeviceInfo) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *PolicyinventoryAbstractDeviceInfo) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *PolicyinventoryAbstractDeviceInfo) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *PolicyinventoryAbstractDeviceInfo) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


