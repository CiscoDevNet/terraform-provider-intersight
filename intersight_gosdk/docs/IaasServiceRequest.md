# IaasServiceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "iaas.ServiceRequest"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "iaas.ServiceRequest"]
**Duration** | Pointer to **string** | Service request duration. | [optional] [readonly] 
**InitiatingUser** | Pointer to **string** | Service Request initiating user. | [optional] [readonly] 
**RequestEndTime** | Pointer to **string** | Service request end time. | [optional] [readonly] 
**RequestId** | Pointer to **string** | Service request id of an SR. | [optional] [readonly] 
**RequestStartTime** | Pointer to **string** | Service request start time. | [optional] [readonly] 
**RequestType** | Pointer to **string** | Service request type of an SR. | [optional] [readonly] 
**Status** | Pointer to **string** | UCSD service request status. | [optional] [readonly] 
**WorkflowName** | Pointer to **string** | Executed workflow name for an SR. | [optional] [readonly] 
**WorkflowSteps** | Pointer to [**[]IaasWorkflowSteps**](IaasWorkflowSteps.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewIaasServiceRequest

`func NewIaasServiceRequest(classId string, objectType string, ) *IaasServiceRequest`

NewIaasServiceRequest instantiates a new IaasServiceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIaasServiceRequestWithDefaults

`func NewIaasServiceRequestWithDefaults() *IaasServiceRequest`

NewIaasServiceRequestWithDefaults instantiates a new IaasServiceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *IaasServiceRequest) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *IaasServiceRequest) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *IaasServiceRequest) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *IaasServiceRequest) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *IaasServiceRequest) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *IaasServiceRequest) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDuration

`func (o *IaasServiceRequest) GetDuration() string`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *IaasServiceRequest) GetDurationOk() (*string, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *IaasServiceRequest) SetDuration(v string)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *IaasServiceRequest) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetInitiatingUser

`func (o *IaasServiceRequest) GetInitiatingUser() string`

GetInitiatingUser returns the InitiatingUser field if non-nil, zero value otherwise.

### GetInitiatingUserOk

`func (o *IaasServiceRequest) GetInitiatingUserOk() (*string, bool)`

GetInitiatingUserOk returns a tuple with the InitiatingUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiatingUser

`func (o *IaasServiceRequest) SetInitiatingUser(v string)`

SetInitiatingUser sets InitiatingUser field to given value.

### HasInitiatingUser

`func (o *IaasServiceRequest) HasInitiatingUser() bool`

HasInitiatingUser returns a boolean if a field has been set.

### GetRequestEndTime

`func (o *IaasServiceRequest) GetRequestEndTime() string`

GetRequestEndTime returns the RequestEndTime field if non-nil, zero value otherwise.

### GetRequestEndTimeOk

`func (o *IaasServiceRequest) GetRequestEndTimeOk() (*string, bool)`

GetRequestEndTimeOk returns a tuple with the RequestEndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestEndTime

`func (o *IaasServiceRequest) SetRequestEndTime(v string)`

SetRequestEndTime sets RequestEndTime field to given value.

### HasRequestEndTime

`func (o *IaasServiceRequest) HasRequestEndTime() bool`

HasRequestEndTime returns a boolean if a field has been set.

### GetRequestId

`func (o *IaasServiceRequest) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *IaasServiceRequest) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *IaasServiceRequest) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *IaasServiceRequest) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetRequestStartTime

`func (o *IaasServiceRequest) GetRequestStartTime() string`

GetRequestStartTime returns the RequestStartTime field if non-nil, zero value otherwise.

### GetRequestStartTimeOk

`func (o *IaasServiceRequest) GetRequestStartTimeOk() (*string, bool)`

GetRequestStartTimeOk returns a tuple with the RequestStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestStartTime

`func (o *IaasServiceRequest) SetRequestStartTime(v string)`

SetRequestStartTime sets RequestStartTime field to given value.

### HasRequestStartTime

`func (o *IaasServiceRequest) HasRequestStartTime() bool`

HasRequestStartTime returns a boolean if a field has been set.

### GetRequestType

`func (o *IaasServiceRequest) GetRequestType() string`

GetRequestType returns the RequestType field if non-nil, zero value otherwise.

### GetRequestTypeOk

`func (o *IaasServiceRequest) GetRequestTypeOk() (*string, bool)`

GetRequestTypeOk returns a tuple with the RequestType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestType

`func (o *IaasServiceRequest) SetRequestType(v string)`

SetRequestType sets RequestType field to given value.

### HasRequestType

`func (o *IaasServiceRequest) HasRequestType() bool`

HasRequestType returns a boolean if a field has been set.

### GetStatus

`func (o *IaasServiceRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *IaasServiceRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *IaasServiceRequest) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *IaasServiceRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetWorkflowName

`func (o *IaasServiceRequest) GetWorkflowName() string`

GetWorkflowName returns the WorkflowName field if non-nil, zero value otherwise.

### GetWorkflowNameOk

`func (o *IaasServiceRequest) GetWorkflowNameOk() (*string, bool)`

GetWorkflowNameOk returns a tuple with the WorkflowName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowName

`func (o *IaasServiceRequest) SetWorkflowName(v string)`

SetWorkflowName sets WorkflowName field to given value.

### HasWorkflowName

`func (o *IaasServiceRequest) HasWorkflowName() bool`

HasWorkflowName returns a boolean if a field has been set.

### GetWorkflowSteps

`func (o *IaasServiceRequest) GetWorkflowSteps() []IaasWorkflowSteps`

GetWorkflowSteps returns the WorkflowSteps field if non-nil, zero value otherwise.

### GetWorkflowStepsOk

`func (o *IaasServiceRequest) GetWorkflowStepsOk() (*[]IaasWorkflowSteps, bool)`

GetWorkflowStepsOk returns a tuple with the WorkflowSteps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowSteps

`func (o *IaasServiceRequest) SetWorkflowSteps(v []IaasWorkflowSteps)`

SetWorkflowSteps sets WorkflowSteps field to given value.

### HasWorkflowSteps

`func (o *IaasServiceRequest) HasWorkflowSteps() bool`

HasWorkflowSteps returns a boolean if a field has been set.

### SetWorkflowStepsNil

`func (o *IaasServiceRequest) SetWorkflowStepsNil(b bool)`

 SetWorkflowStepsNil sets the value for WorkflowSteps to be an explicit nil

### UnsetWorkflowSteps
`func (o *IaasServiceRequest) UnsetWorkflowSteps()`

UnsetWorkflowSteps ensures that no value is present for WorkflowSteps, not even an explicit nil
### GetRegisteredDevice

`func (o *IaasServiceRequest) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *IaasServiceRequest) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *IaasServiceRequest) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *IaasServiceRequest) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


