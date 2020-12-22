# KubernetesVirtualMachineInstanceTypeList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** | The total number of &#39;kubernetes.VirtualMachineInstanceType&#39; resources matching the request, accross all pages. The &#39;Count&#39; attribute is included when the HTTP GET request includes the &#39;$inlinecount&#39; parameter. | [optional] 
**Results** | Pointer to [**[]KubernetesVirtualMachineInstanceType**](KubernetesVirtualMachineInstanceType.md) | The array of &#39;kubernetes.VirtualMachineInstanceType&#39; resources matching the request. | [optional] 

## Methods

### NewKubernetesVirtualMachineInstanceTypeList

`func NewKubernetesVirtualMachineInstanceTypeList() *KubernetesVirtualMachineInstanceTypeList`

NewKubernetesVirtualMachineInstanceTypeList instantiates a new KubernetesVirtualMachineInstanceTypeList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesVirtualMachineInstanceTypeListWithDefaults

`func NewKubernetesVirtualMachineInstanceTypeListWithDefaults() *KubernetesVirtualMachineInstanceTypeList`

NewKubernetesVirtualMachineInstanceTypeListWithDefaults instantiates a new KubernetesVirtualMachineInstanceTypeList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *KubernetesVirtualMachineInstanceTypeList) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *KubernetesVirtualMachineInstanceTypeList) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *KubernetesVirtualMachineInstanceTypeList) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *KubernetesVirtualMachineInstanceTypeList) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetResults

`func (o *KubernetesVirtualMachineInstanceTypeList) GetResults() []KubernetesVirtualMachineInstanceType`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *KubernetesVirtualMachineInstanceTypeList) GetResultsOk() (*[]KubernetesVirtualMachineInstanceType, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *KubernetesVirtualMachineInstanceTypeList) SetResults(v []KubernetesVirtualMachineInstanceType)`

SetResults sets Results field to given value.

### HasResults

`func (o *KubernetesVirtualMachineInstanceTypeList) HasResults() bool`

HasResults returns a boolean if a field has been set.

### SetResultsNil

`func (o *KubernetesVirtualMachineInstanceTypeList) SetResultsNil(b bool)`

 SetResultsNil sets the value for Results to be an explicit nil

### UnsetResults
`func (o *KubernetesVirtualMachineInstanceTypeList) UnsetResults()`

UnsetResults ensures that no value is present for Results, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


