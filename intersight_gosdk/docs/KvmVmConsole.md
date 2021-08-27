# KvmVmConsole

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "kvm.VmConsole"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "kvm.VmConsole"]
**KvmMuxUrl** | Pointer to **string** | The URL of the KVM MUX to connect to. | [optional] [readonly] 
**VirtualMachine** | Pointer to [**HyperflexHxapVirtualMachineRelationship**](HyperflexHxapVirtualMachineRelationship.md) |  | [optional] 

## Methods

### NewKvmVmConsole

`func NewKvmVmConsole(classId string, objectType string, ) *KvmVmConsole`

NewKvmVmConsole instantiates a new KvmVmConsole object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKvmVmConsoleWithDefaults

`func NewKvmVmConsoleWithDefaults() *KvmVmConsole`

NewKvmVmConsoleWithDefaults instantiates a new KvmVmConsole object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *KvmVmConsole) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *KvmVmConsole) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *KvmVmConsole) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *KvmVmConsole) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *KvmVmConsole) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *KvmVmConsole) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetKvmMuxUrl

`func (o *KvmVmConsole) GetKvmMuxUrl() string`

GetKvmMuxUrl returns the KvmMuxUrl field if non-nil, zero value otherwise.

### GetKvmMuxUrlOk

`func (o *KvmVmConsole) GetKvmMuxUrlOk() (*string, bool)`

GetKvmMuxUrlOk returns a tuple with the KvmMuxUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKvmMuxUrl

`func (o *KvmVmConsole) SetKvmMuxUrl(v string)`

SetKvmMuxUrl sets KvmMuxUrl field to given value.

### HasKvmMuxUrl

`func (o *KvmVmConsole) HasKvmMuxUrl() bool`

HasKvmMuxUrl returns a boolean if a field has been set.

### GetVirtualMachine

`func (o *KvmVmConsole) GetVirtualMachine() HyperflexHxapVirtualMachineRelationship`

GetVirtualMachine returns the VirtualMachine field if non-nil, zero value otherwise.

### GetVirtualMachineOk

`func (o *KvmVmConsole) GetVirtualMachineOk() (*HyperflexHxapVirtualMachineRelationship, bool)`

GetVirtualMachineOk returns a tuple with the VirtualMachine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualMachine

`func (o *KvmVmConsole) SetVirtualMachine(v HyperflexHxapVirtualMachineRelationship)`

SetVirtualMachine sets VirtualMachine field to given value.

### HasVirtualMachine

`func (o *KvmVmConsole) HasVirtualMachine() bool`

HasVirtualMachine returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


