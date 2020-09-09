# HyperflexNodeProfileAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HxdpDataIp** | Pointer to **string** | IP address for storage data network (Controller VM interface). | [optional] 
**HxdpMgmtIp** | Pointer to **string** | IP address for HyperFlex management network. | [optional] 
**HypervisorDataIp** | Pointer to **string** | IP address for storage data network (Hypervisor interface). | [optional] 
**HypervisorMgmtIp** | Pointer to **string** | IP address for Hypervisor management network. | [optional] 
**AssignedServer** | Pointer to [**ComputeRackUnitRelationship**](compute.RackUnit.Relationship.md) |  | [optional] 
**ClusterProfile** | Pointer to [**HyperflexClusterProfileRelationship**](hyperflex.ClusterProfile.Relationship.md) |  | [optional] 

## Methods

### NewHyperflexNodeProfileAllOf

`func NewHyperflexNodeProfileAllOf() *HyperflexNodeProfileAllOf`

NewHyperflexNodeProfileAllOf instantiates a new HyperflexNodeProfileAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexNodeProfileAllOfWithDefaults

`func NewHyperflexNodeProfileAllOfWithDefaults() *HyperflexNodeProfileAllOf`

NewHyperflexNodeProfileAllOfWithDefaults instantiates a new HyperflexNodeProfileAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHxdpDataIp

`func (o *HyperflexNodeProfileAllOf) GetHxdpDataIp() string`

GetHxdpDataIp returns the HxdpDataIp field if non-nil, zero value otherwise.

### GetHxdpDataIpOk

`func (o *HyperflexNodeProfileAllOf) GetHxdpDataIpOk() (*string, bool)`

GetHxdpDataIpOk returns a tuple with the HxdpDataIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHxdpDataIp

`func (o *HyperflexNodeProfileAllOf) SetHxdpDataIp(v string)`

SetHxdpDataIp sets HxdpDataIp field to given value.

### HasHxdpDataIp

`func (o *HyperflexNodeProfileAllOf) HasHxdpDataIp() bool`

HasHxdpDataIp returns a boolean if a field has been set.

### GetHxdpMgmtIp

`func (o *HyperflexNodeProfileAllOf) GetHxdpMgmtIp() string`

GetHxdpMgmtIp returns the HxdpMgmtIp field if non-nil, zero value otherwise.

### GetHxdpMgmtIpOk

`func (o *HyperflexNodeProfileAllOf) GetHxdpMgmtIpOk() (*string, bool)`

GetHxdpMgmtIpOk returns a tuple with the HxdpMgmtIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHxdpMgmtIp

`func (o *HyperflexNodeProfileAllOf) SetHxdpMgmtIp(v string)`

SetHxdpMgmtIp sets HxdpMgmtIp field to given value.

### HasHxdpMgmtIp

`func (o *HyperflexNodeProfileAllOf) HasHxdpMgmtIp() bool`

HasHxdpMgmtIp returns a boolean if a field has been set.

### GetHypervisorDataIp

`func (o *HyperflexNodeProfileAllOf) GetHypervisorDataIp() string`

GetHypervisorDataIp returns the HypervisorDataIp field if non-nil, zero value otherwise.

### GetHypervisorDataIpOk

`func (o *HyperflexNodeProfileAllOf) GetHypervisorDataIpOk() (*string, bool)`

GetHypervisorDataIpOk returns a tuple with the HypervisorDataIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHypervisorDataIp

`func (o *HyperflexNodeProfileAllOf) SetHypervisorDataIp(v string)`

SetHypervisorDataIp sets HypervisorDataIp field to given value.

### HasHypervisorDataIp

`func (o *HyperflexNodeProfileAllOf) HasHypervisorDataIp() bool`

HasHypervisorDataIp returns a boolean if a field has been set.

### GetHypervisorMgmtIp

`func (o *HyperflexNodeProfileAllOf) GetHypervisorMgmtIp() string`

GetHypervisorMgmtIp returns the HypervisorMgmtIp field if non-nil, zero value otherwise.

### GetHypervisorMgmtIpOk

`func (o *HyperflexNodeProfileAllOf) GetHypervisorMgmtIpOk() (*string, bool)`

GetHypervisorMgmtIpOk returns a tuple with the HypervisorMgmtIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHypervisorMgmtIp

`func (o *HyperflexNodeProfileAllOf) SetHypervisorMgmtIp(v string)`

SetHypervisorMgmtIp sets HypervisorMgmtIp field to given value.

### HasHypervisorMgmtIp

`func (o *HyperflexNodeProfileAllOf) HasHypervisorMgmtIp() bool`

HasHypervisorMgmtIp returns a boolean if a field has been set.

### GetAssignedServer

`func (o *HyperflexNodeProfileAllOf) GetAssignedServer() ComputeRackUnitRelationship`

GetAssignedServer returns the AssignedServer field if non-nil, zero value otherwise.

### GetAssignedServerOk

`func (o *HyperflexNodeProfileAllOf) GetAssignedServerOk() (*ComputeRackUnitRelationship, bool)`

GetAssignedServerOk returns a tuple with the AssignedServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedServer

`func (o *HyperflexNodeProfileAllOf) SetAssignedServer(v ComputeRackUnitRelationship)`

SetAssignedServer sets AssignedServer field to given value.

### HasAssignedServer

`func (o *HyperflexNodeProfileAllOf) HasAssignedServer() bool`

HasAssignedServer returns a boolean if a field has been set.

### GetClusterProfile

`func (o *HyperflexNodeProfileAllOf) GetClusterProfile() HyperflexClusterProfileRelationship`

GetClusterProfile returns the ClusterProfile field if non-nil, zero value otherwise.

### GetClusterProfileOk

`func (o *HyperflexNodeProfileAllOf) GetClusterProfileOk() (*HyperflexClusterProfileRelationship, bool)`

GetClusterProfileOk returns a tuple with the ClusterProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterProfile

`func (o *HyperflexNodeProfileAllOf) SetClusterProfile(v HyperflexClusterProfileRelationship)`

SetClusterProfile sets ClusterProfile field to given value.

### HasClusterProfile

`func (o *HyperflexNodeProfileAllOf) HasClusterProfile() bool`

HasClusterProfile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


