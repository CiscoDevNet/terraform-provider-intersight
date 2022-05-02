# ConnectorTargetSpecificationAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "connector.TargetSpecification"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "connector.TargetSpecification"]
**CpuLimit** | Pointer to **string** | CPU limit assigned to the docker container. It is total amount of CPU time that a container can use every 100ms. A container cannot use more than its share of CPU time during this interval. | [optional] 
**CpuRequest** | Pointer to **string** | Requested CPU value for a docker container to run in Intersight Assist. | [optional] 
**ImageTag** | Pointer to **string** | Docker image tag used to define kubernetes deployment for each target. Image tag should be the complete URL. This image can be found locally in case of Intersight Appliance or can be pulled from Intersight cloud in Intersight Assist deployment. | [optional] 
**MemoryLimit** | Pointer to **string** | Intersight Assist prevents the docker container from using more than the configured memory limit. If a Container exceeds its memory limit, it might be terminated. If it is restartable, the kubelet will restart it, as with any other type of runtime failure. | [optional] 
**MemoryRequest** | Pointer to **string** | Requested memory value for a docker container to run in Intersight Assist. | [optional] 

## Methods

### NewConnectorTargetSpecificationAllOf

`func NewConnectorTargetSpecificationAllOf(classId string, objectType string, ) *ConnectorTargetSpecificationAllOf`

NewConnectorTargetSpecificationAllOf instantiates a new ConnectorTargetSpecificationAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectorTargetSpecificationAllOfWithDefaults

`func NewConnectorTargetSpecificationAllOfWithDefaults() *ConnectorTargetSpecificationAllOf`

NewConnectorTargetSpecificationAllOfWithDefaults instantiates a new ConnectorTargetSpecificationAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ConnectorTargetSpecificationAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ConnectorTargetSpecificationAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ConnectorTargetSpecificationAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ConnectorTargetSpecificationAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ConnectorTargetSpecificationAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ConnectorTargetSpecificationAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCpuLimit

`func (o *ConnectorTargetSpecificationAllOf) GetCpuLimit() string`

GetCpuLimit returns the CpuLimit field if non-nil, zero value otherwise.

### GetCpuLimitOk

`func (o *ConnectorTargetSpecificationAllOf) GetCpuLimitOk() (*string, bool)`

GetCpuLimitOk returns a tuple with the CpuLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuLimit

`func (o *ConnectorTargetSpecificationAllOf) SetCpuLimit(v string)`

SetCpuLimit sets CpuLimit field to given value.

### HasCpuLimit

`func (o *ConnectorTargetSpecificationAllOf) HasCpuLimit() bool`

HasCpuLimit returns a boolean if a field has been set.

### GetCpuRequest

`func (o *ConnectorTargetSpecificationAllOf) GetCpuRequest() string`

GetCpuRequest returns the CpuRequest field if non-nil, zero value otherwise.

### GetCpuRequestOk

`func (o *ConnectorTargetSpecificationAllOf) GetCpuRequestOk() (*string, bool)`

GetCpuRequestOk returns a tuple with the CpuRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuRequest

`func (o *ConnectorTargetSpecificationAllOf) SetCpuRequest(v string)`

SetCpuRequest sets CpuRequest field to given value.

### HasCpuRequest

`func (o *ConnectorTargetSpecificationAllOf) HasCpuRequest() bool`

HasCpuRequest returns a boolean if a field has been set.

### GetImageTag

`func (o *ConnectorTargetSpecificationAllOf) GetImageTag() string`

GetImageTag returns the ImageTag field if non-nil, zero value otherwise.

### GetImageTagOk

`func (o *ConnectorTargetSpecificationAllOf) GetImageTagOk() (*string, bool)`

GetImageTagOk returns a tuple with the ImageTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageTag

`func (o *ConnectorTargetSpecificationAllOf) SetImageTag(v string)`

SetImageTag sets ImageTag field to given value.

### HasImageTag

`func (o *ConnectorTargetSpecificationAllOf) HasImageTag() bool`

HasImageTag returns a boolean if a field has been set.

### GetMemoryLimit

`func (o *ConnectorTargetSpecificationAllOf) GetMemoryLimit() string`

GetMemoryLimit returns the MemoryLimit field if non-nil, zero value otherwise.

### GetMemoryLimitOk

`func (o *ConnectorTargetSpecificationAllOf) GetMemoryLimitOk() (*string, bool)`

GetMemoryLimitOk returns a tuple with the MemoryLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryLimit

`func (o *ConnectorTargetSpecificationAllOf) SetMemoryLimit(v string)`

SetMemoryLimit sets MemoryLimit field to given value.

### HasMemoryLimit

`func (o *ConnectorTargetSpecificationAllOf) HasMemoryLimit() bool`

HasMemoryLimit returns a boolean if a field has been set.

### GetMemoryRequest

`func (o *ConnectorTargetSpecificationAllOf) GetMemoryRequest() string`

GetMemoryRequest returns the MemoryRequest field if non-nil, zero value otherwise.

### GetMemoryRequestOk

`func (o *ConnectorTargetSpecificationAllOf) GetMemoryRequestOk() (*string, bool)`

GetMemoryRequestOk returns a tuple with the MemoryRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryRequest

`func (o *ConnectorTargetSpecificationAllOf) SetMemoryRequest(v string)`

SetMemoryRequest sets MemoryRequest field to given value.

### HasMemoryRequest

`func (o *ConnectorTargetSpecificationAllOf) HasMemoryRequest() bool`

HasMemoryRequest returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


