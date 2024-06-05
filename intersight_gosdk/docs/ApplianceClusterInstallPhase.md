# ApplianceClusterInstallPhase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "appliance.ClusterInstallPhase"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "appliance.ClusterInstallPhase"]
**CurrentNode** | Pointer to **int64** | Node number install is running on. | [optional] [readonly] 
**CurrentSubphase** | Pointer to **int64** | Subphase of the phase currently running. | [optional] [readonly] 
**ElapsedTime** | Pointer to **int64** | Elapsed time in seconds to complete the current install phase. | [optional] [readonly] 
**EndTime** | Pointer to **time.Time** | End date of the software install phase. | [optional] [readonly] 
**Failed** | Pointer to **bool** | Indicates if the install phase has failed or not. | [optional] [readonly] 
**Message** | Pointer to **string** | Status message set during the install phase. | [optional] [readonly] 
**Name** | Pointer to **string** | Name of the install phase. * &#x60;Backup&#x60; - Backup the currently running node. * &#x60;EnableBootstrap&#x60; - Disable echo and enable nginx on the currently running node. * &#x60;UpdateAnsibleHosts&#x60; - Update /etc/ansible/hosts on each node. * &#x60;UpdateNetworkConfig&#x60; - Update Network config for node IP change scenario. * &#x60;SyncImages&#x60; - Sync images and manifest to each node. * &#x60;ConfigureEtcd&#x60; - Configure etcd on each node. * &#x60;DeployServices&#x60; - Deploy kubernetes services from node 1. * &#x60;InstallEquinox&#x60; - Configure and install equinox on each node. * &#x60;Validate&#x60; - Validate equinox is running in primary/secondary mode. * &#x60;CheckApplianceClusterState&#x60; - Check the appliance cluster status. * &#x60;Success&#x60; - Upgrade completed successfully. * &#x60;Fail&#x60; - Indicates that the upgrade process has failed. * &#x60;Cancel&#x60; - Indicates that the upgrade was canceled by the Intersight Appliance. | [optional] [readonly] [default to "Backup"]
**PendingNodes** | Pointer to **[]int64** |  | [optional] 
**StartTime** | Pointer to **time.Time** | Start date of the software install phase. | [optional] [readonly] 
**Status** | Pointer to **string** | Status of the install phase. | [optional] [readonly] 

## Methods

### NewApplianceClusterInstallPhase

`func NewApplianceClusterInstallPhase(classId string, objectType string, ) *ApplianceClusterInstallPhase`

NewApplianceClusterInstallPhase instantiates a new ApplianceClusterInstallPhase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplianceClusterInstallPhaseWithDefaults

`func NewApplianceClusterInstallPhaseWithDefaults() *ApplianceClusterInstallPhase`

NewApplianceClusterInstallPhaseWithDefaults instantiates a new ApplianceClusterInstallPhase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ApplianceClusterInstallPhase) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ApplianceClusterInstallPhase) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ApplianceClusterInstallPhase) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ApplianceClusterInstallPhase) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ApplianceClusterInstallPhase) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ApplianceClusterInstallPhase) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCurrentNode

`func (o *ApplianceClusterInstallPhase) GetCurrentNode() int64`

GetCurrentNode returns the CurrentNode field if non-nil, zero value otherwise.

### GetCurrentNodeOk

`func (o *ApplianceClusterInstallPhase) GetCurrentNodeOk() (*int64, bool)`

GetCurrentNodeOk returns a tuple with the CurrentNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentNode

`func (o *ApplianceClusterInstallPhase) SetCurrentNode(v int64)`

SetCurrentNode sets CurrentNode field to given value.

### HasCurrentNode

`func (o *ApplianceClusterInstallPhase) HasCurrentNode() bool`

HasCurrentNode returns a boolean if a field has been set.

### GetCurrentSubphase

`func (o *ApplianceClusterInstallPhase) GetCurrentSubphase() int64`

GetCurrentSubphase returns the CurrentSubphase field if non-nil, zero value otherwise.

### GetCurrentSubphaseOk

`func (o *ApplianceClusterInstallPhase) GetCurrentSubphaseOk() (*int64, bool)`

GetCurrentSubphaseOk returns a tuple with the CurrentSubphase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentSubphase

`func (o *ApplianceClusterInstallPhase) SetCurrentSubphase(v int64)`

SetCurrentSubphase sets CurrentSubphase field to given value.

### HasCurrentSubphase

`func (o *ApplianceClusterInstallPhase) HasCurrentSubphase() bool`

HasCurrentSubphase returns a boolean if a field has been set.

### GetElapsedTime

`func (o *ApplianceClusterInstallPhase) GetElapsedTime() int64`

GetElapsedTime returns the ElapsedTime field if non-nil, zero value otherwise.

### GetElapsedTimeOk

`func (o *ApplianceClusterInstallPhase) GetElapsedTimeOk() (*int64, bool)`

GetElapsedTimeOk returns a tuple with the ElapsedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElapsedTime

`func (o *ApplianceClusterInstallPhase) SetElapsedTime(v int64)`

SetElapsedTime sets ElapsedTime field to given value.

### HasElapsedTime

`func (o *ApplianceClusterInstallPhase) HasElapsedTime() bool`

HasElapsedTime returns a boolean if a field has been set.

### GetEndTime

`func (o *ApplianceClusterInstallPhase) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *ApplianceClusterInstallPhase) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *ApplianceClusterInstallPhase) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *ApplianceClusterInstallPhase) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetFailed

`func (o *ApplianceClusterInstallPhase) GetFailed() bool`

GetFailed returns the Failed field if non-nil, zero value otherwise.

### GetFailedOk

`func (o *ApplianceClusterInstallPhase) GetFailedOk() (*bool, bool)`

GetFailedOk returns a tuple with the Failed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailed

`func (o *ApplianceClusterInstallPhase) SetFailed(v bool)`

SetFailed sets Failed field to given value.

### HasFailed

`func (o *ApplianceClusterInstallPhase) HasFailed() bool`

HasFailed returns a boolean if a field has been set.

### GetMessage

`func (o *ApplianceClusterInstallPhase) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ApplianceClusterInstallPhase) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ApplianceClusterInstallPhase) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ApplianceClusterInstallPhase) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetName

`func (o *ApplianceClusterInstallPhase) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApplianceClusterInstallPhase) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApplianceClusterInstallPhase) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApplianceClusterInstallPhase) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPendingNodes

`func (o *ApplianceClusterInstallPhase) GetPendingNodes() []int64`

GetPendingNodes returns the PendingNodes field if non-nil, zero value otherwise.

### GetPendingNodesOk

`func (o *ApplianceClusterInstallPhase) GetPendingNodesOk() (*[]int64, bool)`

GetPendingNodesOk returns a tuple with the PendingNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingNodes

`func (o *ApplianceClusterInstallPhase) SetPendingNodes(v []int64)`

SetPendingNodes sets PendingNodes field to given value.

### HasPendingNodes

`func (o *ApplianceClusterInstallPhase) HasPendingNodes() bool`

HasPendingNodes returns a boolean if a field has been set.

### SetPendingNodesNil

`func (o *ApplianceClusterInstallPhase) SetPendingNodesNil(b bool)`

 SetPendingNodesNil sets the value for PendingNodes to be an explicit nil

### UnsetPendingNodes
`func (o *ApplianceClusterInstallPhase) UnsetPendingNodes()`

UnsetPendingNodes ensures that no value is present for PendingNodes, not even an explicit nil
### GetStartTime

`func (o *ApplianceClusterInstallPhase) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *ApplianceClusterInstallPhase) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *ApplianceClusterInstallPhase) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *ApplianceClusterInstallPhase) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetStatus

`func (o *ApplianceClusterInstallPhase) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ApplianceClusterInstallPhase) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ApplianceClusterInstallPhase) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ApplianceClusterInstallPhase) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


