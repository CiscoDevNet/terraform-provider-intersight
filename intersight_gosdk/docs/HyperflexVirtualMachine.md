# HyperflexVirtualMachine

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hyperflex.VirtualMachine"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hyperflex.VirtualMachine"]
**RunTimeInfo** | Pointer to [**NullableHyperflexVirtualMachineRuntimeInfo**](HyperflexVirtualMachineRuntimeInfo.md) |  | [optional] 
**StatusCode** | Pointer to **string** | Virtual Machine Status Code. * &#x60;VM_ACCESSIBLE&#x60; - This virtual machine is accessible. * &#x60;VM_INACCESSIBLE&#x60; - This virtual machine is not accessible. * &#x60;VM_NOT_SUPPORTED&#x60; - This virtual machine is not supported. * &#x60;NONE&#x60; - This virtual machine does not have a status code. | [optional] [readonly] [default to "VM_ACCESSIBLE"]
**Uuid** | Pointer to **string** | Virtual machine unique UUID. | [optional] [readonly] 

## Methods

### NewHyperflexVirtualMachine

`func NewHyperflexVirtualMachine(classId string, objectType string, ) *HyperflexVirtualMachine`

NewHyperflexVirtualMachine instantiates a new HyperflexVirtualMachine object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexVirtualMachineWithDefaults

`func NewHyperflexVirtualMachineWithDefaults() *HyperflexVirtualMachine`

NewHyperflexVirtualMachineWithDefaults instantiates a new HyperflexVirtualMachine object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HyperflexVirtualMachine) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexVirtualMachine) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexVirtualMachine) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HyperflexVirtualMachine) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexVirtualMachine) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexVirtualMachine) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetRunTimeInfo

`func (o *HyperflexVirtualMachine) GetRunTimeInfo() HyperflexVirtualMachineRuntimeInfo`

GetRunTimeInfo returns the RunTimeInfo field if non-nil, zero value otherwise.

### GetRunTimeInfoOk

`func (o *HyperflexVirtualMachine) GetRunTimeInfoOk() (*HyperflexVirtualMachineRuntimeInfo, bool)`

GetRunTimeInfoOk returns a tuple with the RunTimeInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunTimeInfo

`func (o *HyperflexVirtualMachine) SetRunTimeInfo(v HyperflexVirtualMachineRuntimeInfo)`

SetRunTimeInfo sets RunTimeInfo field to given value.

### HasRunTimeInfo

`func (o *HyperflexVirtualMachine) HasRunTimeInfo() bool`

HasRunTimeInfo returns a boolean if a field has been set.

### SetRunTimeInfoNil

`func (o *HyperflexVirtualMachine) SetRunTimeInfoNil(b bool)`

 SetRunTimeInfoNil sets the value for RunTimeInfo to be an explicit nil

### UnsetRunTimeInfo
`func (o *HyperflexVirtualMachine) UnsetRunTimeInfo()`

UnsetRunTimeInfo ensures that no value is present for RunTimeInfo, not even an explicit nil
### GetStatusCode

`func (o *HyperflexVirtualMachine) GetStatusCode() string`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *HyperflexVirtualMachine) GetStatusCodeOk() (*string, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *HyperflexVirtualMachine) SetStatusCode(v string)`

SetStatusCode sets StatusCode field to given value.

### HasStatusCode

`func (o *HyperflexVirtualMachine) HasStatusCode() bool`

HasStatusCode returns a boolean if a field has been set.

### GetUuid

`func (o *HyperflexVirtualMachine) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *HyperflexVirtualMachine) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *HyperflexVirtualMachine) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *HyperflexVirtualMachine) HasUuid() bool`

HasUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


