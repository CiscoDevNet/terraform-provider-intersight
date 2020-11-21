# VirtualizationVmwareVmCpuSocketInfoAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "virtualization.VmwareVmCpuSocketInfo"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "virtualization.VmwareVmCpuSocketInfo"]
**CoresPerSocket** | Pointer to **int64** | The number of core per CPU socket (value may be empty). | [optional] 
**NumCpus** | Pointer to **int64** | Number of CPUs allocated to this VM. | [optional] 
**NumSockets** | Pointer to **int64** | The number of CPU sockets allocated. | [optional] 

## Methods

### NewVirtualizationVmwareVmCpuSocketInfoAllOf

`func NewVirtualizationVmwareVmCpuSocketInfoAllOf(classId string, objectType string, ) *VirtualizationVmwareVmCpuSocketInfoAllOf`

NewVirtualizationVmwareVmCpuSocketInfoAllOf instantiates a new VirtualizationVmwareVmCpuSocketInfoAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualizationVmwareVmCpuSocketInfoAllOfWithDefaults

`func NewVirtualizationVmwareVmCpuSocketInfoAllOfWithDefaults() *VirtualizationVmwareVmCpuSocketInfoAllOf`

NewVirtualizationVmwareVmCpuSocketInfoAllOfWithDefaults instantiates a new VirtualizationVmwareVmCpuSocketInfoAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VirtualizationVmwareVmCpuSocketInfoAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VirtualizationVmwareVmCpuSocketInfoAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VirtualizationVmwareVmCpuSocketInfoAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VirtualizationVmwareVmCpuSocketInfoAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VirtualizationVmwareVmCpuSocketInfoAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VirtualizationVmwareVmCpuSocketInfoAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCoresPerSocket

`func (o *VirtualizationVmwareVmCpuSocketInfoAllOf) GetCoresPerSocket() int64`

GetCoresPerSocket returns the CoresPerSocket field if non-nil, zero value otherwise.

### GetCoresPerSocketOk

`func (o *VirtualizationVmwareVmCpuSocketInfoAllOf) GetCoresPerSocketOk() (*int64, bool)`

GetCoresPerSocketOk returns a tuple with the CoresPerSocket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoresPerSocket

`func (o *VirtualizationVmwareVmCpuSocketInfoAllOf) SetCoresPerSocket(v int64)`

SetCoresPerSocket sets CoresPerSocket field to given value.

### HasCoresPerSocket

`func (o *VirtualizationVmwareVmCpuSocketInfoAllOf) HasCoresPerSocket() bool`

HasCoresPerSocket returns a boolean if a field has been set.

### GetNumCpus

`func (o *VirtualizationVmwareVmCpuSocketInfoAllOf) GetNumCpus() int64`

GetNumCpus returns the NumCpus field if non-nil, zero value otherwise.

### GetNumCpusOk

`func (o *VirtualizationVmwareVmCpuSocketInfoAllOf) GetNumCpusOk() (*int64, bool)`

GetNumCpusOk returns a tuple with the NumCpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumCpus

`func (o *VirtualizationVmwareVmCpuSocketInfoAllOf) SetNumCpus(v int64)`

SetNumCpus sets NumCpus field to given value.

### HasNumCpus

`func (o *VirtualizationVmwareVmCpuSocketInfoAllOf) HasNumCpus() bool`

HasNumCpus returns a boolean if a field has been set.

### GetNumSockets

`func (o *VirtualizationVmwareVmCpuSocketInfoAllOf) GetNumSockets() int64`

GetNumSockets returns the NumSockets field if non-nil, zero value otherwise.

### GetNumSocketsOk

`func (o *VirtualizationVmwareVmCpuSocketInfoAllOf) GetNumSocketsOk() (*int64, bool)`

GetNumSocketsOk returns a tuple with the NumSockets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumSockets

`func (o *VirtualizationVmwareVmCpuSocketInfoAllOf) SetNumSockets(v int64)`

SetNumSockets sets NumSockets field to given value.

### HasNumSockets

`func (o *VirtualizationVmwareVmCpuSocketInfoAllOf) HasNumSockets() bool`

HasNumSockets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


