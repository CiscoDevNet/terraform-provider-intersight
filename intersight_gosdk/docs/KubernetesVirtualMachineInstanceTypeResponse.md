# KubernetesVirtualMachineInstanceTypeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjectType** | **string** | A discriminator value to disambiguate the schema of a HTTP GET response body. | 
**Count** | Pointer to **int32** | The total number of &#39;kubernetes.VirtualMachineInstanceType&#39; resources matching the request, accross all pages. The &#39;Count&#39; attribute is included when the HTTP GET request includes the &#39;$inlinecount&#39; parameter. | [optional] 
**Results** | Pointer to [**[]MoTagKeySummary**](MoTagKeySummary.md) |  | [optional] 

## Methods

### NewKubernetesVirtualMachineInstanceTypeResponse

`func NewKubernetesVirtualMachineInstanceTypeResponse(objectType string, ) *KubernetesVirtualMachineInstanceTypeResponse`

NewKubernetesVirtualMachineInstanceTypeResponse instantiates a new KubernetesVirtualMachineInstanceTypeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesVirtualMachineInstanceTypeResponseWithDefaults

`func NewKubernetesVirtualMachineInstanceTypeResponseWithDefaults() *KubernetesVirtualMachineInstanceTypeResponse`

NewKubernetesVirtualMachineInstanceTypeResponseWithDefaults instantiates a new KubernetesVirtualMachineInstanceTypeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjectType

`func (o *KubernetesVirtualMachineInstanceTypeResponse) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *KubernetesVirtualMachineInstanceTypeResponse) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *KubernetesVirtualMachineInstanceTypeResponse) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCount

`func (o *KubernetesVirtualMachineInstanceTypeResponse) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *KubernetesVirtualMachineInstanceTypeResponse) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *KubernetesVirtualMachineInstanceTypeResponse) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *KubernetesVirtualMachineInstanceTypeResponse) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetResults

`func (o *KubernetesVirtualMachineInstanceTypeResponse) GetResults() []MoTagKeySummary`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *KubernetesVirtualMachineInstanceTypeResponse) GetResultsOk() (*[]MoTagKeySummary, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *KubernetesVirtualMachineInstanceTypeResponse) SetResults(v []MoTagKeySummary)`

SetResults sets Results field to given value.

### HasResults

`func (o *KubernetesVirtualMachineInstanceTypeResponse) HasResults() bool`

HasResults returns a boolean if a field has been set.

### SetResultsNil

`func (o *KubernetesVirtualMachineInstanceTypeResponse) SetResultsNil(b bool)`

 SetResultsNil sets the value for Results to be an explicit nil

### UnsetResults
`func (o *KubernetesVirtualMachineInstanceTypeResponse) UnsetResults()`

UnsetResults ensures that no value is present for Results, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


