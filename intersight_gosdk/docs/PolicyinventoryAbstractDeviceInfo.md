# PolicyinventoryAbstractDeviceInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfigState** | Pointer to **string** | Configuration state of server profile config context. | [optional] [readonly] 
**ControlAction** | Pointer to **string** | Control action of server profile config context. | [optional] [readonly] 
**ErrorState** | Pointer to **string** | Error state of server profile config context. | [optional] [readonly] 
**JobInfo** | Pointer to [**[]PolicyinventoryJobInfo**](policyinventory.JobInfo.md) |  | [optional] 
**OperState** | Pointer to **string** | Operational state of server profile config context. | [optional] [readonly] 
**ProfileMoId** | Pointer to **string** | Server profile MO ID of the server. | [optional] [readonly] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 

## Methods

### NewPolicyinventoryAbstractDeviceInfo

`func NewPolicyinventoryAbstractDeviceInfo() *PolicyinventoryAbstractDeviceInfo`

NewPolicyinventoryAbstractDeviceInfo instantiates a new PolicyinventoryAbstractDeviceInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyinventoryAbstractDeviceInfoWithDefaults

`func NewPolicyinventoryAbstractDeviceInfoWithDefaults() *PolicyinventoryAbstractDeviceInfo`

NewPolicyinventoryAbstractDeviceInfoWithDefaults instantiates a new PolicyinventoryAbstractDeviceInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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


