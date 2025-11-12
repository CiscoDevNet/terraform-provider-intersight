# WorkloadAssignedResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "workload.AssignedResource"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "workload.AssignedResource"]
**BluePrint** | Pointer to [**MoMoRef**](MoMoRef.md) |  | [optional] 
**BluePrintRefName** | Pointer to **string** | The reference name for the blueprint as given in the workload definition. | [optional] [readonly] 
**Lease** | Pointer to [**MoMoRef**](MoMoRef.md) |  | [optional] 
**Resource** | Pointer to [**MoMoRef**](MoMoRef.md) |  | [optional] 
**ServiceItemInstance** | Pointer to [**MoMoRef**](MoMoRef.md) |  | [optional] 

## Methods

### NewWorkloadAssignedResource

`func NewWorkloadAssignedResource(classId string, objectType string, ) *WorkloadAssignedResource`

NewWorkloadAssignedResource instantiates a new WorkloadAssignedResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkloadAssignedResourceWithDefaults

`func NewWorkloadAssignedResourceWithDefaults() *WorkloadAssignedResource`

NewWorkloadAssignedResourceWithDefaults instantiates a new WorkloadAssignedResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *WorkloadAssignedResource) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *WorkloadAssignedResource) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *WorkloadAssignedResource) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *WorkloadAssignedResource) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkloadAssignedResource) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkloadAssignedResource) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetBluePrint

`func (o *WorkloadAssignedResource) GetBluePrint() MoMoRef`

GetBluePrint returns the BluePrint field if non-nil, zero value otherwise.

### GetBluePrintOk

`func (o *WorkloadAssignedResource) GetBluePrintOk() (*MoMoRef, bool)`

GetBluePrintOk returns a tuple with the BluePrint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBluePrint

`func (o *WorkloadAssignedResource) SetBluePrint(v MoMoRef)`

SetBluePrint sets BluePrint field to given value.

### HasBluePrint

`func (o *WorkloadAssignedResource) HasBluePrint() bool`

HasBluePrint returns a boolean if a field has been set.

### GetBluePrintRefName

`func (o *WorkloadAssignedResource) GetBluePrintRefName() string`

GetBluePrintRefName returns the BluePrintRefName field if non-nil, zero value otherwise.

### GetBluePrintRefNameOk

`func (o *WorkloadAssignedResource) GetBluePrintRefNameOk() (*string, bool)`

GetBluePrintRefNameOk returns a tuple with the BluePrintRefName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBluePrintRefName

`func (o *WorkloadAssignedResource) SetBluePrintRefName(v string)`

SetBluePrintRefName sets BluePrintRefName field to given value.

### HasBluePrintRefName

`func (o *WorkloadAssignedResource) HasBluePrintRefName() bool`

HasBluePrintRefName returns a boolean if a field has been set.

### GetLease

`func (o *WorkloadAssignedResource) GetLease() MoMoRef`

GetLease returns the Lease field if non-nil, zero value otherwise.

### GetLeaseOk

`func (o *WorkloadAssignedResource) GetLeaseOk() (*MoMoRef, bool)`

GetLeaseOk returns a tuple with the Lease field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLease

`func (o *WorkloadAssignedResource) SetLease(v MoMoRef)`

SetLease sets Lease field to given value.

### HasLease

`func (o *WorkloadAssignedResource) HasLease() bool`

HasLease returns a boolean if a field has been set.

### GetResource

`func (o *WorkloadAssignedResource) GetResource() MoMoRef`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *WorkloadAssignedResource) GetResourceOk() (*MoMoRef, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *WorkloadAssignedResource) SetResource(v MoMoRef)`

SetResource sets Resource field to given value.

### HasResource

`func (o *WorkloadAssignedResource) HasResource() bool`

HasResource returns a boolean if a field has been set.

### GetServiceItemInstance

`func (o *WorkloadAssignedResource) GetServiceItemInstance() MoMoRef`

GetServiceItemInstance returns the ServiceItemInstance field if non-nil, zero value otherwise.

### GetServiceItemInstanceOk

`func (o *WorkloadAssignedResource) GetServiceItemInstanceOk() (*MoMoRef, bool)`

GetServiceItemInstanceOk returns a tuple with the ServiceItemInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceItemInstance

`func (o *WorkloadAssignedResource) SetServiceItemInstance(v MoMoRef)`

SetServiceItemInstance sets ServiceItemInstance field to given value.

### HasServiceItemInstance

`func (o *WorkloadAssignedResource) HasServiceItemInstance() bool`

HasServiceItemInstance returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


