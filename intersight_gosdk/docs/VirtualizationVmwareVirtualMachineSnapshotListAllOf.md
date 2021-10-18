# VirtualizationVmwareVirtualMachineSnapshotListAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** | The total number of &#39;virtualization.VmwareVirtualMachineSnapshot&#39; resources matching the request, accross all pages. The &#39;Count&#39; attribute is included when the HTTP GET request includes the &#39;$inlinecount&#39; parameter. | [optional] 
**Results** | Pointer to [**[]VirtualizationVmwareVirtualMachineSnapshot**](VirtualizationVmwareVirtualMachineSnapshot.md) | The array of &#39;virtualization.VmwareVirtualMachineSnapshot&#39; resources matching the request. | [optional] 

## Methods

### NewVirtualizationVmwareVirtualMachineSnapshotListAllOf

`func NewVirtualizationVmwareVirtualMachineSnapshotListAllOf() *VirtualizationVmwareVirtualMachineSnapshotListAllOf`

NewVirtualizationVmwareVirtualMachineSnapshotListAllOf instantiates a new VirtualizationVmwareVirtualMachineSnapshotListAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualizationVmwareVirtualMachineSnapshotListAllOfWithDefaults

`func NewVirtualizationVmwareVirtualMachineSnapshotListAllOfWithDefaults() *VirtualizationVmwareVirtualMachineSnapshotListAllOf`

NewVirtualizationVmwareVirtualMachineSnapshotListAllOfWithDefaults instantiates a new VirtualizationVmwareVirtualMachineSnapshotListAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *VirtualizationVmwareVirtualMachineSnapshotListAllOf) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *VirtualizationVmwareVirtualMachineSnapshotListAllOf) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *VirtualizationVmwareVirtualMachineSnapshotListAllOf) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *VirtualizationVmwareVirtualMachineSnapshotListAllOf) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetResults

`func (o *VirtualizationVmwareVirtualMachineSnapshotListAllOf) GetResults() []VirtualizationVmwareVirtualMachineSnapshot`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *VirtualizationVmwareVirtualMachineSnapshotListAllOf) GetResultsOk() (*[]VirtualizationVmwareVirtualMachineSnapshot, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *VirtualizationVmwareVirtualMachineSnapshotListAllOf) SetResults(v []VirtualizationVmwareVirtualMachineSnapshot)`

SetResults sets Results field to given value.

### HasResults

`func (o *VirtualizationVmwareVirtualMachineSnapshotListAllOf) HasResults() bool`

HasResults returns a boolean if a field has been set.

### SetResultsNil

`func (o *VirtualizationVmwareVirtualMachineSnapshotListAllOf) SetResultsNil(b bool)`

 SetResultsNil sets the value for Results to be an explicit nil

### UnsetResults
`func (o *VirtualizationVmwareVirtualMachineSnapshotListAllOf) UnsetResults()`

UnsetResults ensures that no value is present for Results, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


