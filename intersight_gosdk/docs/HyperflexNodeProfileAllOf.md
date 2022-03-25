# HyperflexNodeProfileAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hyperflex.NodeProfile"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hyperflex.NodeProfile"]
**HxdpDataIp** | Pointer to **string** | IP address for storage data network (Controller VM interface). | [optional] 
**HxdpMgmtIp** | Pointer to **string** | IP address for HyperFlex management network. | [optional] 
**HxdpStorageClientIp** | Pointer to **string** | IP address for storage client network (Controller VM interface). | [optional] 
**HypervisorControlIp** | Pointer to **string** | IP address for hypervisor control such as VM migration or pod management. | [optional] 
**HypervisorDataIp** | Pointer to **string** | IP address for storage data network (Hypervisor interface). | [optional] 
**HypervisorMgmtIp** | Pointer to **string** | IP address for Hypervisor management network. | [optional] 
**NodeRole** | Pointer to **string** | The role that this node performs in the HyperFlex cluster. * &#x60;Unknown&#x60; - The node role is not available. * &#x60;Storage&#x60; - The node persists data and contributes to the storage capacity of a cluster. * &#x60;Compute&#x60; - The node contributes to the compute capacity of a cluster. | [optional] [readonly] [default to "Unknown"]
**AssignedServer** | Pointer to [**ComputePhysicalRelationship**](ComputePhysicalRelationship.md) |  | [optional] 
**ClusterProfile** | Pointer to [**HyperflexClusterProfileRelationship**](HyperflexClusterProfileRelationship.md) |  | [optional] 
**Node** | Pointer to [**HyperflexNodeRelationship**](HyperflexNodeRelationship.md) |  | [optional] 

## Methods

### NewHyperflexNodeProfileAllOf

`func NewHyperflexNodeProfileAllOf(classId string, objectType string, ) *HyperflexNodeProfileAllOf`

NewHyperflexNodeProfileAllOf instantiates a new HyperflexNodeProfileAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexNodeProfileAllOfWithDefaults

`func NewHyperflexNodeProfileAllOfWithDefaults() *HyperflexNodeProfileAllOf`

NewHyperflexNodeProfileAllOfWithDefaults instantiates a new HyperflexNodeProfileAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HyperflexNodeProfileAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexNodeProfileAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexNodeProfileAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HyperflexNodeProfileAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexNodeProfileAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexNodeProfileAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


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

### GetHxdpStorageClientIp

`func (o *HyperflexNodeProfileAllOf) GetHxdpStorageClientIp() string`

GetHxdpStorageClientIp returns the HxdpStorageClientIp field if non-nil, zero value otherwise.

### GetHxdpStorageClientIpOk

`func (o *HyperflexNodeProfileAllOf) GetHxdpStorageClientIpOk() (*string, bool)`

GetHxdpStorageClientIpOk returns a tuple with the HxdpStorageClientIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHxdpStorageClientIp

`func (o *HyperflexNodeProfileAllOf) SetHxdpStorageClientIp(v string)`

SetHxdpStorageClientIp sets HxdpStorageClientIp field to given value.

### HasHxdpStorageClientIp

`func (o *HyperflexNodeProfileAllOf) HasHxdpStorageClientIp() bool`

HasHxdpStorageClientIp returns a boolean if a field has been set.

### GetHypervisorControlIp

`func (o *HyperflexNodeProfileAllOf) GetHypervisorControlIp() string`

GetHypervisorControlIp returns the HypervisorControlIp field if non-nil, zero value otherwise.

### GetHypervisorControlIpOk

`func (o *HyperflexNodeProfileAllOf) GetHypervisorControlIpOk() (*string, bool)`

GetHypervisorControlIpOk returns a tuple with the HypervisorControlIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHypervisorControlIp

`func (o *HyperflexNodeProfileAllOf) SetHypervisorControlIp(v string)`

SetHypervisorControlIp sets HypervisorControlIp field to given value.

### HasHypervisorControlIp

`func (o *HyperflexNodeProfileAllOf) HasHypervisorControlIp() bool`

HasHypervisorControlIp returns a boolean if a field has been set.

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

### GetNodeRole

`func (o *HyperflexNodeProfileAllOf) GetNodeRole() string`

GetNodeRole returns the NodeRole field if non-nil, zero value otherwise.

### GetNodeRoleOk

`func (o *HyperflexNodeProfileAllOf) GetNodeRoleOk() (*string, bool)`

GetNodeRoleOk returns a tuple with the NodeRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeRole

`func (o *HyperflexNodeProfileAllOf) SetNodeRole(v string)`

SetNodeRole sets NodeRole field to given value.

### HasNodeRole

`func (o *HyperflexNodeProfileAllOf) HasNodeRole() bool`

HasNodeRole returns a boolean if a field has been set.

### GetAssignedServer

`func (o *HyperflexNodeProfileAllOf) GetAssignedServer() ComputePhysicalRelationship`

GetAssignedServer returns the AssignedServer field if non-nil, zero value otherwise.

### GetAssignedServerOk

`func (o *HyperflexNodeProfileAllOf) GetAssignedServerOk() (*ComputePhysicalRelationship, bool)`

GetAssignedServerOk returns a tuple with the AssignedServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedServer

`func (o *HyperflexNodeProfileAllOf) SetAssignedServer(v ComputePhysicalRelationship)`

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

### GetNode

`func (o *HyperflexNodeProfileAllOf) GetNode() HyperflexNodeRelationship`

GetNode returns the Node field if non-nil, zero value otherwise.

### GetNodeOk

`func (o *HyperflexNodeProfileAllOf) GetNodeOk() (*HyperflexNodeRelationship, bool)`

GetNodeOk returns a tuple with the Node field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNode

`func (o *HyperflexNodeProfileAllOf) SetNode(v HyperflexNodeRelationship)`

SetNode sets Node field to given value.

### HasNode

`func (o *HyperflexNodeProfileAllOf) HasNode() bool`

HasNode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


