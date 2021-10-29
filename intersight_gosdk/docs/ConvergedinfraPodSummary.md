# ConvergedinfraPodSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "convergedinfra.PodSummary"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "convergedinfra.PodSummary"]
**ActiveNodes** | Pointer to **int64** | Number of Nodes that are powered on and have at least 1 VM associated with the pod. | [optional] [readonly] 
**VmCount** | Pointer to **int64** | Number of VMs associated with the pod. | [optional] [readonly] 

## Methods

### NewConvergedinfraPodSummary

`func NewConvergedinfraPodSummary(classId string, objectType string, ) *ConvergedinfraPodSummary`

NewConvergedinfraPodSummary instantiates a new ConvergedinfraPodSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConvergedinfraPodSummaryWithDefaults

`func NewConvergedinfraPodSummaryWithDefaults() *ConvergedinfraPodSummary`

NewConvergedinfraPodSummaryWithDefaults instantiates a new ConvergedinfraPodSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ConvergedinfraPodSummary) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ConvergedinfraPodSummary) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ConvergedinfraPodSummary) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ConvergedinfraPodSummary) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ConvergedinfraPodSummary) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ConvergedinfraPodSummary) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetActiveNodes

`func (o *ConvergedinfraPodSummary) GetActiveNodes() int64`

GetActiveNodes returns the ActiveNodes field if non-nil, zero value otherwise.

### GetActiveNodesOk

`func (o *ConvergedinfraPodSummary) GetActiveNodesOk() (*int64, bool)`

GetActiveNodesOk returns a tuple with the ActiveNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveNodes

`func (o *ConvergedinfraPodSummary) SetActiveNodes(v int64)`

SetActiveNodes sets ActiveNodes field to given value.

### HasActiveNodes

`func (o *ConvergedinfraPodSummary) HasActiveNodes() bool`

HasActiveNodes returns a boolean if a field has been set.

### GetVmCount

`func (o *ConvergedinfraPodSummary) GetVmCount() int64`

GetVmCount returns the VmCount field if non-nil, zero value otherwise.

### GetVmCountOk

`func (o *ConvergedinfraPodSummary) GetVmCountOk() (*int64, bool)`

GetVmCountOk returns a tuple with the VmCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmCount

`func (o *ConvergedinfraPodSummary) SetVmCount(v int64)`

SetVmCount sets VmCount field to given value.

### HasVmCount

`func (o *ConvergedinfraPodSummary) HasVmCount() bool`

HasVmCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


