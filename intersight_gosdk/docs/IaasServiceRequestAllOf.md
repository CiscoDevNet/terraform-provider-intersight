# IaasServiceRequestAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Duration** | Pointer to **string** | Service request duration. | [optional] [readonly] 
**InitiatingUser** | Pointer to **string** | Service Request initiating user. | [optional] [readonly] 
**RequestEndTime** | Pointer to **string** | Service request end time. | [optional] [readonly] 
**RequestId** | Pointer to **string** | Service request id of an SR. | [optional] [readonly] 
**RequestStartTime** | Pointer to **string** | Service request start time. | [optional] [readonly] 
**RequestType** | Pointer to **string** | Service request type of an SR. | [optional] [readonly] 
**Status** | Pointer to **string** | UCSD service request status. | [optional] [readonly] 
**WorkflowName** | Pointer to **string** | Executed workflow name for an SR. | [optional] [readonly] 
**WorkflowSteps** | Pointer to [**[]IaasWorkflowSteps**](iaas.WorkflowSteps.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 

## Methods

### NewIaasServiceRequestAllOf

`func NewIaasServiceRequestAllOf() *IaasServiceRequestAllOf`

NewIaasServiceRequestAllOf instantiates a new IaasServiceRequestAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIaasServiceRequestAllOfWithDefaults

`func NewIaasServiceRequestAllOfWithDefaults() *IaasServiceRequestAllOf`

NewIaasServiceRequestAllOfWithDefaults instantiates a new IaasServiceRequestAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDuration

`func (o *IaasServiceRequestAllOf) GetDuration() string`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *IaasServiceRequestAllOf) GetDurationOk() (*string, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *IaasServiceRequestAllOf) SetDuration(v string)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *IaasServiceRequestAllOf) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetInitiatingUser

`func (o *IaasServiceRequestAllOf) GetInitiatingUser() string`

GetInitiatingUser returns the InitiatingUser field if non-nil, zero value otherwise.

### GetInitiatingUserOk

`func (o *IaasServiceRequestAllOf) GetInitiatingUserOk() (*string, bool)`

GetInitiatingUserOk returns a tuple with the InitiatingUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiatingUser

`func (o *IaasServiceRequestAllOf) SetInitiatingUser(v string)`

SetInitiatingUser sets InitiatingUser field to given value.

### HasInitiatingUser

`func (o *IaasServiceRequestAllOf) HasInitiatingUser() bool`

HasInitiatingUser returns a boolean if a field has been set.

### GetRequestEndTime

`func (o *IaasServiceRequestAllOf) GetRequestEndTime() string`

GetRequestEndTime returns the RequestEndTime field if non-nil, zero value otherwise.

### GetRequestEndTimeOk

`func (o *IaasServiceRequestAllOf) GetRequestEndTimeOk() (*string, bool)`

GetRequestEndTimeOk returns a tuple with the RequestEndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestEndTime

`func (o *IaasServiceRequestAllOf) SetRequestEndTime(v string)`

SetRequestEndTime sets RequestEndTime field to given value.

### HasRequestEndTime

`func (o *IaasServiceRequestAllOf) HasRequestEndTime() bool`

HasRequestEndTime returns a boolean if a field has been set.

### GetRequestId

`func (o *IaasServiceRequestAllOf) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *IaasServiceRequestAllOf) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *IaasServiceRequestAllOf) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *IaasServiceRequestAllOf) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetRequestStartTime

`func (o *IaasServiceRequestAllOf) GetRequestStartTime() string`

GetRequestStartTime returns the RequestStartTime field if non-nil, zero value otherwise.

### GetRequestStartTimeOk

`func (o *IaasServiceRequestAllOf) GetRequestStartTimeOk() (*string, bool)`

GetRequestStartTimeOk returns a tuple with the RequestStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestStartTime

`func (o *IaasServiceRequestAllOf) SetRequestStartTime(v string)`

SetRequestStartTime sets RequestStartTime field to given value.

### HasRequestStartTime

`func (o *IaasServiceRequestAllOf) HasRequestStartTime() bool`

HasRequestStartTime returns a boolean if a field has been set.

### GetRequestType

`func (o *IaasServiceRequestAllOf) GetRequestType() string`

GetRequestType returns the RequestType field if non-nil, zero value otherwise.

### GetRequestTypeOk

`func (o *IaasServiceRequestAllOf) GetRequestTypeOk() (*string, bool)`

GetRequestTypeOk returns a tuple with the RequestType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestType

`func (o *IaasServiceRequestAllOf) SetRequestType(v string)`

SetRequestType sets RequestType field to given value.

### HasRequestType

`func (o *IaasServiceRequestAllOf) HasRequestType() bool`

HasRequestType returns a boolean if a field has been set.

### GetStatus

`func (o *IaasServiceRequestAllOf) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *IaasServiceRequestAllOf) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *IaasServiceRequestAllOf) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *IaasServiceRequestAllOf) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetWorkflowName

`func (o *IaasServiceRequestAllOf) GetWorkflowName() string`

GetWorkflowName returns the WorkflowName field if non-nil, zero value otherwise.

### GetWorkflowNameOk

`func (o *IaasServiceRequestAllOf) GetWorkflowNameOk() (*string, bool)`

GetWorkflowNameOk returns a tuple with the WorkflowName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowName

`func (o *IaasServiceRequestAllOf) SetWorkflowName(v string)`

SetWorkflowName sets WorkflowName field to given value.

### HasWorkflowName

`func (o *IaasServiceRequestAllOf) HasWorkflowName() bool`

HasWorkflowName returns a boolean if a field has been set.

### GetWorkflowSteps

`func (o *IaasServiceRequestAllOf) GetWorkflowSteps() []IaasWorkflowSteps`

GetWorkflowSteps returns the WorkflowSteps field if non-nil, zero value otherwise.

### GetWorkflowStepsOk

`func (o *IaasServiceRequestAllOf) GetWorkflowStepsOk() (*[]IaasWorkflowSteps, bool)`

GetWorkflowStepsOk returns a tuple with the WorkflowSteps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowSteps

`func (o *IaasServiceRequestAllOf) SetWorkflowSteps(v []IaasWorkflowSteps)`

SetWorkflowSteps sets WorkflowSteps field to given value.

### HasWorkflowSteps

`func (o *IaasServiceRequestAllOf) HasWorkflowSteps() bool`

HasWorkflowSteps returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *IaasServiceRequestAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *IaasServiceRequestAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *IaasServiceRequestAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *IaasServiceRequestAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


