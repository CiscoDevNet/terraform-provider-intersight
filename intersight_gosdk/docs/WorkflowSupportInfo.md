# WorkflowSupportInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "workflow.SupportInfo"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "workflow.SupportInfo"]
**DeprecationInfo** | Pointer to [**NullableWorkflowDeprecationInfo**](WorkflowDeprecationInfo.md) |  | [optional] 
**SupportStatus** | Pointer to **string** | The support status of the versioned object. * &#x60;Supported&#x60; - The definition is a supported version and there will be no changes to the mandatory inputs or outputs. * &#x60;Beta&#x60; - The definition is a Beta version and this version can under go changes until the version is marked supported. * &#x60;Deprecated&#x60; - The version of definition is deprecated and typically there will be a higher version of the same definition that has been added. | [optional] [default to "Supported"]

## Methods

### NewWorkflowSupportInfo

`func NewWorkflowSupportInfo(classId string, objectType string, ) *WorkflowSupportInfo`

NewWorkflowSupportInfo instantiates a new WorkflowSupportInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowSupportInfoWithDefaults

`func NewWorkflowSupportInfoWithDefaults() *WorkflowSupportInfo`

NewWorkflowSupportInfoWithDefaults instantiates a new WorkflowSupportInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *WorkflowSupportInfo) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *WorkflowSupportInfo) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *WorkflowSupportInfo) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *WorkflowSupportInfo) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkflowSupportInfo) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkflowSupportInfo) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDeprecationInfo

`func (o *WorkflowSupportInfo) GetDeprecationInfo() WorkflowDeprecationInfo`

GetDeprecationInfo returns the DeprecationInfo field if non-nil, zero value otherwise.

### GetDeprecationInfoOk

`func (o *WorkflowSupportInfo) GetDeprecationInfoOk() (*WorkflowDeprecationInfo, bool)`

GetDeprecationInfoOk returns a tuple with the DeprecationInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeprecationInfo

`func (o *WorkflowSupportInfo) SetDeprecationInfo(v WorkflowDeprecationInfo)`

SetDeprecationInfo sets DeprecationInfo field to given value.

### HasDeprecationInfo

`func (o *WorkflowSupportInfo) HasDeprecationInfo() bool`

HasDeprecationInfo returns a boolean if a field has been set.

### SetDeprecationInfoNil

`func (o *WorkflowSupportInfo) SetDeprecationInfoNil(b bool)`

 SetDeprecationInfoNil sets the value for DeprecationInfo to be an explicit nil

### UnsetDeprecationInfo
`func (o *WorkflowSupportInfo) UnsetDeprecationInfo()`

UnsetDeprecationInfo ensures that no value is present for DeprecationInfo, not even an explicit nil
### GetSupportStatus

`func (o *WorkflowSupportInfo) GetSupportStatus() string`

GetSupportStatus returns the SupportStatus field if non-nil, zero value otherwise.

### GetSupportStatusOk

`func (o *WorkflowSupportInfo) GetSupportStatusOk() (*string, bool)`

GetSupportStatusOk returns a tuple with the SupportStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportStatus

`func (o *WorkflowSupportInfo) SetSupportStatus(v string)`

SetSupportStatus sets SupportStatus field to given value.

### HasSupportStatus

`func (o *WorkflowSupportInfo) HasSupportStatus() bool`

HasSupportStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


