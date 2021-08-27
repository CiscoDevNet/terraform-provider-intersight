# PolicyinventoryAbstractDeviceInfoAllOf

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

### NewPolicyinventoryAbstractDeviceInfoAllOf

`func NewPolicyinventoryAbstractDeviceInfoAllOf(classId string, objectType string, ) *PolicyinventoryAbstractDeviceInfoAllOf`

NewPolicyinventoryAbstractDeviceInfoAllOf instantiates a new PolicyinventoryAbstractDeviceInfoAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyinventoryAbstractDeviceInfoAllOfWithDefaults

`func NewPolicyinventoryAbstractDeviceInfoAllOfWithDefaults() *PolicyinventoryAbstractDeviceInfoAllOf`

NewPolicyinventoryAbstractDeviceInfoAllOfWithDefaults instantiates a new PolicyinventoryAbstractDeviceInfoAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *PolicyinventoryAbstractDeviceInfoAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *PolicyinventoryAbstractDeviceInfoAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *PolicyinventoryAbstractDeviceInfoAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *PolicyinventoryAbstractDeviceInfoAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *PolicyinventoryAbstractDeviceInfoAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *PolicyinventoryAbstractDeviceInfoAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetConfigState

`func (o *PolicyinventoryAbstractDeviceInfoAllOf) GetConfigState() string`

GetConfigState returns the ConfigState field if non-nil, zero value otherwise.

### GetConfigStateOk

`func (o *PolicyinventoryAbstractDeviceInfoAllOf) GetConfigStateOk() (*string, bool)`

GetConfigStateOk returns a tuple with the ConfigState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigState

`func (o *PolicyinventoryAbstractDeviceInfoAllOf) SetConfigState(v string)`

SetConfigState sets ConfigState field to given value.

### HasConfigState

`func (o *PolicyinventoryAbstractDeviceInfoAllOf) HasConfigState() bool`

HasConfigState returns a boolean if a field has been set.

### GetControlAction

`func (o *PolicyinventoryAbstractDeviceInfoAllOf) GetControlAction() string`

GetControlAction returns the ControlAction field if non-nil, zero value otherwise.

### GetControlActionOk

`func (o *PolicyinventoryAbstractDeviceInfoAllOf) GetControlActionOk() (*string, bool)`

GetControlActionOk returns a tuple with the ControlAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlAction

`func (o *PolicyinventoryAbstractDeviceInfoAllOf) SetControlAction(v string)`

SetControlAction sets ControlAction field to given value.

### HasControlAction

`func (o *PolicyinventoryAbstractDeviceInfoAllOf) HasControlAction() bool`

HasControlAction returns a boolean if a field has been set.

### GetErrorState

`func (o *PolicyinventoryAbstractDeviceInfoAllOf) GetErrorState() string`

GetErrorState returns the ErrorState field if non-nil, zero value otherwise.

### GetErrorStateOk

`func (o *PolicyinventoryAbstractDeviceInfoAllOf) GetErrorStateOk() (*string, bool)`

GetErrorStateOk returns a tuple with the ErrorState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorState

`func (o *PolicyinventoryAbstractDeviceInfoAllOf) SetErrorState(v string)`

SetErrorState sets ErrorState field to given value.

### HasErrorState

`func (o *PolicyinventoryAbstractDeviceInfoAllOf) HasErrorState() bool`

HasErrorState returns a boolean if a field has been set.

### GetJobInfo

`func (o *PolicyinventoryAbstractDeviceInfoAllOf) GetJobInfo() []PolicyinventoryJobInfo`

GetJobInfo returns the JobInfo field if non-nil, zero value otherwise.

### GetJobInfoOk

`func (o *PolicyinventoryAbstractDeviceInfoAllOf) GetJobInfoOk() (*[]PolicyinventoryJobInfo, bool)`

GetJobInfoOk returns a tuple with the JobInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobInfo

`func (o *PolicyinventoryAbstractDeviceInfoAllOf) SetJobInfo(v []PolicyinventoryJobInfo)`

SetJobInfo sets JobInfo field to given value.

### HasJobInfo

`func (o *PolicyinventoryAbstractDeviceInfoAllOf) HasJobInfo() bool`

HasJobInfo returns a boolean if a field has been set.

### SetJobInfoNil

`func (o *PolicyinventoryAbstractDeviceInfoAllOf) SetJobInfoNil(b bool)`

 SetJobInfoNil sets the value for JobInfo to be an explicit nil

### UnsetJobInfo
`func (o *PolicyinventoryAbstractDeviceInfoAllOf) UnsetJobInfo()`

UnsetJobInfo ensures that no value is present for JobInfo, not even an explicit nil
### GetOperState

`func (o *PolicyinventoryAbstractDeviceInfoAllOf) GetOperState() string`

GetOperState returns the OperState field if non-nil, zero value otherwise.

### GetOperStateOk

`func (o *PolicyinventoryAbstractDeviceInfoAllOf) GetOperStateOk() (*string, bool)`

GetOperStateOk returns a tuple with the OperState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperState

`func (o *PolicyinventoryAbstractDeviceInfoAllOf) SetOperState(v string)`

SetOperState sets OperState field to given value.

### HasOperState

`func (o *PolicyinventoryAbstractDeviceInfoAllOf) HasOperState() bool`

HasOperState returns a boolean if a field has been set.

### GetProfileMoId

`func (o *PolicyinventoryAbstractDeviceInfoAllOf) GetProfileMoId() string`

GetProfileMoId returns the ProfileMoId field if non-nil, zero value otherwise.

### GetProfileMoIdOk

`func (o *PolicyinventoryAbstractDeviceInfoAllOf) GetProfileMoIdOk() (*string, bool)`

GetProfileMoIdOk returns a tuple with the ProfileMoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileMoId

`func (o *PolicyinventoryAbstractDeviceInfoAllOf) SetProfileMoId(v string)`

SetProfileMoId sets ProfileMoId field to given value.

### HasProfileMoId

`func (o *PolicyinventoryAbstractDeviceInfoAllOf) HasProfileMoId() bool`

HasProfileMoId returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *PolicyinventoryAbstractDeviceInfoAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *PolicyinventoryAbstractDeviceInfoAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *PolicyinventoryAbstractDeviceInfoAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *PolicyinventoryAbstractDeviceInfoAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


