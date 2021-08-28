# VirtualizationVmwareDatastoreCluster

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "virtualization.VmwareDatastoreCluster"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "virtualization.VmwareDatastoreCluster"]
**AutomationLevel** | Pointer to **string** | The global automation level for all virtual machines in this datastore cluster. | [optional] 
**FreeSpaceThreshold** | Pointer to **int64** | Minimum level of free space for each datastore that is the threshold for action. | [optional] 
**InventoryPath** | Pointer to **string** | Inventory path of the Datastore Cluster. | [optional] 
**IoLatencyThreshold** | Pointer to **int32** | Minimum I/O latency for each datastore below which I/O load balancing moves are not considered. | [optional] 
**IoLoadBalanceAutomationMode** | Pointer to **string** | Storage DRS behavior when it generates recommendations for correcting I/O load imbalance in a datastore cluster. | [optional] 
**IoLoadImbalanceThreshold** | Pointer to **int32** | Amount of imbalance that Storage DRS should tolerate. | [optional] 
**IoMetricsEnabled** | Pointer to **bool** | Is I/O Metrics enabled for this datastore cluster. | [optional] 
**MinSpaceUtilizationDifference** | Pointer to **int32** | Specify how much of an improvement DRS should look for before making a recommendation or performing a migration. | [optional] 
**PolicyEnforcementAutomationMode** | Pointer to **string** | Storage DRS behavior when it generates recommendations for correcting storage and VM policy violations in a datastore cluster. | [optional] 
**ReservablePercentThreshold** | Pointer to **int32** | Storage DRS makes storage migration recommendations if total IOPs reservation of all VMs running on a datastore is higher than the specified threshold. | [optional] 
**RuleEnforcementAutomationMode** | Pointer to **string** | Storage DRS behavior when it generates recommendations for correcting affinity rule violations in a datastore cluster. | [optional] 
**SpaceLoadBalanceAutomationMode** | Pointer to **string** | Storage DRS behavior when it generates recommendations for correcting space load imbalance in a datastore cluster. | [optional] 
**SpaceThresholdMode** | Pointer to **string** | Runtime thresholds govern when Storage DRS performs or recommends migrations. | [optional] 
**Status** | Pointer to **string** | Datastore cluster health status, as reported by the hypervisor platform. * &#x60;Unknown&#x60; - Entity status is unknown. * &#x60;Degraded&#x60; - State is degraded, and might impact normal operation of the entity. * &#x60;Critical&#x60; - Entity is in a critical state, impacting operations. * &#x60;Ok&#x60; - Entity status is in a stable state, operating normally. | [optional] [default to "Unknown"]
**StorageDrsEnabled** | Pointer to **bool** | Is Storage DRS enabled for this datastore cluster. | [optional] 
**UtilizedSpaceThreshold** | Pointer to **int32** | Minimum level of consumed space for each datastore that is the threshold for action. | [optional] 
**VmEvacuationAutomationMode** | Pointer to **string** | Storage DRS behavior when it generates recommendations for VM evacuations from datastores in a datastore cluster. | [optional] 
**Datacenter** | Pointer to [**VirtualizationVmwareDatacenterRelationship**](VirtualizationVmwareDatacenterRelationship.md) |  | [optional] 

## Methods

### NewVirtualizationVmwareDatastoreCluster

`func NewVirtualizationVmwareDatastoreCluster(classId string, objectType string, ) *VirtualizationVmwareDatastoreCluster`

NewVirtualizationVmwareDatastoreCluster instantiates a new VirtualizationVmwareDatastoreCluster object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualizationVmwareDatastoreClusterWithDefaults

`func NewVirtualizationVmwareDatastoreClusterWithDefaults() *VirtualizationVmwareDatastoreCluster`

NewVirtualizationVmwareDatastoreClusterWithDefaults instantiates a new VirtualizationVmwareDatastoreCluster object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VirtualizationVmwareDatastoreCluster) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VirtualizationVmwareDatastoreCluster) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VirtualizationVmwareDatastoreCluster) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VirtualizationVmwareDatastoreCluster) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VirtualizationVmwareDatastoreCluster) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VirtualizationVmwareDatastoreCluster) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAutomationLevel

`func (o *VirtualizationVmwareDatastoreCluster) GetAutomationLevel() string`

GetAutomationLevel returns the AutomationLevel field if non-nil, zero value otherwise.

### GetAutomationLevelOk

`func (o *VirtualizationVmwareDatastoreCluster) GetAutomationLevelOk() (*string, bool)`

GetAutomationLevelOk returns a tuple with the AutomationLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomationLevel

`func (o *VirtualizationVmwareDatastoreCluster) SetAutomationLevel(v string)`

SetAutomationLevel sets AutomationLevel field to given value.

### HasAutomationLevel

`func (o *VirtualizationVmwareDatastoreCluster) HasAutomationLevel() bool`

HasAutomationLevel returns a boolean if a field has been set.

### GetFreeSpaceThreshold

`func (o *VirtualizationVmwareDatastoreCluster) GetFreeSpaceThreshold() int64`

GetFreeSpaceThreshold returns the FreeSpaceThreshold field if non-nil, zero value otherwise.

### GetFreeSpaceThresholdOk

`func (o *VirtualizationVmwareDatastoreCluster) GetFreeSpaceThresholdOk() (*int64, bool)`

GetFreeSpaceThresholdOk returns a tuple with the FreeSpaceThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeSpaceThreshold

`func (o *VirtualizationVmwareDatastoreCluster) SetFreeSpaceThreshold(v int64)`

SetFreeSpaceThreshold sets FreeSpaceThreshold field to given value.

### HasFreeSpaceThreshold

`func (o *VirtualizationVmwareDatastoreCluster) HasFreeSpaceThreshold() bool`

HasFreeSpaceThreshold returns a boolean if a field has been set.

### GetInventoryPath

`func (o *VirtualizationVmwareDatastoreCluster) GetInventoryPath() string`

GetInventoryPath returns the InventoryPath field if non-nil, zero value otherwise.

### GetInventoryPathOk

`func (o *VirtualizationVmwareDatastoreCluster) GetInventoryPathOk() (*string, bool)`

GetInventoryPathOk returns a tuple with the InventoryPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryPath

`func (o *VirtualizationVmwareDatastoreCluster) SetInventoryPath(v string)`

SetInventoryPath sets InventoryPath field to given value.

### HasInventoryPath

`func (o *VirtualizationVmwareDatastoreCluster) HasInventoryPath() bool`

HasInventoryPath returns a boolean if a field has been set.

### GetIoLatencyThreshold

`func (o *VirtualizationVmwareDatastoreCluster) GetIoLatencyThreshold() int32`

GetIoLatencyThreshold returns the IoLatencyThreshold field if non-nil, zero value otherwise.

### GetIoLatencyThresholdOk

`func (o *VirtualizationVmwareDatastoreCluster) GetIoLatencyThresholdOk() (*int32, bool)`

GetIoLatencyThresholdOk returns a tuple with the IoLatencyThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIoLatencyThreshold

`func (o *VirtualizationVmwareDatastoreCluster) SetIoLatencyThreshold(v int32)`

SetIoLatencyThreshold sets IoLatencyThreshold field to given value.

### HasIoLatencyThreshold

`func (o *VirtualizationVmwareDatastoreCluster) HasIoLatencyThreshold() bool`

HasIoLatencyThreshold returns a boolean if a field has been set.

### GetIoLoadBalanceAutomationMode

`func (o *VirtualizationVmwareDatastoreCluster) GetIoLoadBalanceAutomationMode() string`

GetIoLoadBalanceAutomationMode returns the IoLoadBalanceAutomationMode field if non-nil, zero value otherwise.

### GetIoLoadBalanceAutomationModeOk

`func (o *VirtualizationVmwareDatastoreCluster) GetIoLoadBalanceAutomationModeOk() (*string, bool)`

GetIoLoadBalanceAutomationModeOk returns a tuple with the IoLoadBalanceAutomationMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIoLoadBalanceAutomationMode

`func (o *VirtualizationVmwareDatastoreCluster) SetIoLoadBalanceAutomationMode(v string)`

SetIoLoadBalanceAutomationMode sets IoLoadBalanceAutomationMode field to given value.

### HasIoLoadBalanceAutomationMode

`func (o *VirtualizationVmwareDatastoreCluster) HasIoLoadBalanceAutomationMode() bool`

HasIoLoadBalanceAutomationMode returns a boolean if a field has been set.

### GetIoLoadImbalanceThreshold

`func (o *VirtualizationVmwareDatastoreCluster) GetIoLoadImbalanceThreshold() int32`

GetIoLoadImbalanceThreshold returns the IoLoadImbalanceThreshold field if non-nil, zero value otherwise.

### GetIoLoadImbalanceThresholdOk

`func (o *VirtualizationVmwareDatastoreCluster) GetIoLoadImbalanceThresholdOk() (*int32, bool)`

GetIoLoadImbalanceThresholdOk returns a tuple with the IoLoadImbalanceThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIoLoadImbalanceThreshold

`func (o *VirtualizationVmwareDatastoreCluster) SetIoLoadImbalanceThreshold(v int32)`

SetIoLoadImbalanceThreshold sets IoLoadImbalanceThreshold field to given value.

### HasIoLoadImbalanceThreshold

`func (o *VirtualizationVmwareDatastoreCluster) HasIoLoadImbalanceThreshold() bool`

HasIoLoadImbalanceThreshold returns a boolean if a field has been set.

### GetIoMetricsEnabled

`func (o *VirtualizationVmwareDatastoreCluster) GetIoMetricsEnabled() bool`

GetIoMetricsEnabled returns the IoMetricsEnabled field if non-nil, zero value otherwise.

### GetIoMetricsEnabledOk

`func (o *VirtualizationVmwareDatastoreCluster) GetIoMetricsEnabledOk() (*bool, bool)`

GetIoMetricsEnabledOk returns a tuple with the IoMetricsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIoMetricsEnabled

`func (o *VirtualizationVmwareDatastoreCluster) SetIoMetricsEnabled(v bool)`

SetIoMetricsEnabled sets IoMetricsEnabled field to given value.

### HasIoMetricsEnabled

`func (o *VirtualizationVmwareDatastoreCluster) HasIoMetricsEnabled() bool`

HasIoMetricsEnabled returns a boolean if a field has been set.

### GetMinSpaceUtilizationDifference

`func (o *VirtualizationVmwareDatastoreCluster) GetMinSpaceUtilizationDifference() int32`

GetMinSpaceUtilizationDifference returns the MinSpaceUtilizationDifference field if non-nil, zero value otherwise.

### GetMinSpaceUtilizationDifferenceOk

`func (o *VirtualizationVmwareDatastoreCluster) GetMinSpaceUtilizationDifferenceOk() (*int32, bool)`

GetMinSpaceUtilizationDifferenceOk returns a tuple with the MinSpaceUtilizationDifference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinSpaceUtilizationDifference

`func (o *VirtualizationVmwareDatastoreCluster) SetMinSpaceUtilizationDifference(v int32)`

SetMinSpaceUtilizationDifference sets MinSpaceUtilizationDifference field to given value.

### HasMinSpaceUtilizationDifference

`func (o *VirtualizationVmwareDatastoreCluster) HasMinSpaceUtilizationDifference() bool`

HasMinSpaceUtilizationDifference returns a boolean if a field has been set.

### GetPolicyEnforcementAutomationMode

`func (o *VirtualizationVmwareDatastoreCluster) GetPolicyEnforcementAutomationMode() string`

GetPolicyEnforcementAutomationMode returns the PolicyEnforcementAutomationMode field if non-nil, zero value otherwise.

### GetPolicyEnforcementAutomationModeOk

`func (o *VirtualizationVmwareDatastoreCluster) GetPolicyEnforcementAutomationModeOk() (*string, bool)`

GetPolicyEnforcementAutomationModeOk returns a tuple with the PolicyEnforcementAutomationMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyEnforcementAutomationMode

`func (o *VirtualizationVmwareDatastoreCluster) SetPolicyEnforcementAutomationMode(v string)`

SetPolicyEnforcementAutomationMode sets PolicyEnforcementAutomationMode field to given value.

### HasPolicyEnforcementAutomationMode

`func (o *VirtualizationVmwareDatastoreCluster) HasPolicyEnforcementAutomationMode() bool`

HasPolicyEnforcementAutomationMode returns a boolean if a field has been set.

### GetReservablePercentThreshold

`func (o *VirtualizationVmwareDatastoreCluster) GetReservablePercentThreshold() int32`

GetReservablePercentThreshold returns the ReservablePercentThreshold field if non-nil, zero value otherwise.

### GetReservablePercentThresholdOk

`func (o *VirtualizationVmwareDatastoreCluster) GetReservablePercentThresholdOk() (*int32, bool)`

GetReservablePercentThresholdOk returns a tuple with the ReservablePercentThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservablePercentThreshold

`func (o *VirtualizationVmwareDatastoreCluster) SetReservablePercentThreshold(v int32)`

SetReservablePercentThreshold sets ReservablePercentThreshold field to given value.

### HasReservablePercentThreshold

`func (o *VirtualizationVmwareDatastoreCluster) HasReservablePercentThreshold() bool`

HasReservablePercentThreshold returns a boolean if a field has been set.

### GetRuleEnforcementAutomationMode

`func (o *VirtualizationVmwareDatastoreCluster) GetRuleEnforcementAutomationMode() string`

GetRuleEnforcementAutomationMode returns the RuleEnforcementAutomationMode field if non-nil, zero value otherwise.

### GetRuleEnforcementAutomationModeOk

`func (o *VirtualizationVmwareDatastoreCluster) GetRuleEnforcementAutomationModeOk() (*string, bool)`

GetRuleEnforcementAutomationModeOk returns a tuple with the RuleEnforcementAutomationMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleEnforcementAutomationMode

`func (o *VirtualizationVmwareDatastoreCluster) SetRuleEnforcementAutomationMode(v string)`

SetRuleEnforcementAutomationMode sets RuleEnforcementAutomationMode field to given value.

### HasRuleEnforcementAutomationMode

`func (o *VirtualizationVmwareDatastoreCluster) HasRuleEnforcementAutomationMode() bool`

HasRuleEnforcementAutomationMode returns a boolean if a field has been set.

### GetSpaceLoadBalanceAutomationMode

`func (o *VirtualizationVmwareDatastoreCluster) GetSpaceLoadBalanceAutomationMode() string`

GetSpaceLoadBalanceAutomationMode returns the SpaceLoadBalanceAutomationMode field if non-nil, zero value otherwise.

### GetSpaceLoadBalanceAutomationModeOk

`func (o *VirtualizationVmwareDatastoreCluster) GetSpaceLoadBalanceAutomationModeOk() (*string, bool)`

GetSpaceLoadBalanceAutomationModeOk returns a tuple with the SpaceLoadBalanceAutomationMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceLoadBalanceAutomationMode

`func (o *VirtualizationVmwareDatastoreCluster) SetSpaceLoadBalanceAutomationMode(v string)`

SetSpaceLoadBalanceAutomationMode sets SpaceLoadBalanceAutomationMode field to given value.

### HasSpaceLoadBalanceAutomationMode

`func (o *VirtualizationVmwareDatastoreCluster) HasSpaceLoadBalanceAutomationMode() bool`

HasSpaceLoadBalanceAutomationMode returns a boolean if a field has been set.

### GetSpaceThresholdMode

`func (o *VirtualizationVmwareDatastoreCluster) GetSpaceThresholdMode() string`

GetSpaceThresholdMode returns the SpaceThresholdMode field if non-nil, zero value otherwise.

### GetSpaceThresholdModeOk

`func (o *VirtualizationVmwareDatastoreCluster) GetSpaceThresholdModeOk() (*string, bool)`

GetSpaceThresholdModeOk returns a tuple with the SpaceThresholdMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceThresholdMode

`func (o *VirtualizationVmwareDatastoreCluster) SetSpaceThresholdMode(v string)`

SetSpaceThresholdMode sets SpaceThresholdMode field to given value.

### HasSpaceThresholdMode

`func (o *VirtualizationVmwareDatastoreCluster) HasSpaceThresholdMode() bool`

HasSpaceThresholdMode returns a boolean if a field has been set.

### GetStatus

`func (o *VirtualizationVmwareDatastoreCluster) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VirtualizationVmwareDatastoreCluster) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VirtualizationVmwareDatastoreCluster) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *VirtualizationVmwareDatastoreCluster) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStorageDrsEnabled

`func (o *VirtualizationVmwareDatastoreCluster) GetStorageDrsEnabled() bool`

GetStorageDrsEnabled returns the StorageDrsEnabled field if non-nil, zero value otherwise.

### GetStorageDrsEnabledOk

`func (o *VirtualizationVmwareDatastoreCluster) GetStorageDrsEnabledOk() (*bool, bool)`

GetStorageDrsEnabledOk returns a tuple with the StorageDrsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageDrsEnabled

`func (o *VirtualizationVmwareDatastoreCluster) SetStorageDrsEnabled(v bool)`

SetStorageDrsEnabled sets StorageDrsEnabled field to given value.

### HasStorageDrsEnabled

`func (o *VirtualizationVmwareDatastoreCluster) HasStorageDrsEnabled() bool`

HasStorageDrsEnabled returns a boolean if a field has been set.

### GetUtilizedSpaceThreshold

`func (o *VirtualizationVmwareDatastoreCluster) GetUtilizedSpaceThreshold() int32`

GetUtilizedSpaceThreshold returns the UtilizedSpaceThreshold field if non-nil, zero value otherwise.

### GetUtilizedSpaceThresholdOk

`func (o *VirtualizationVmwareDatastoreCluster) GetUtilizedSpaceThresholdOk() (*int32, bool)`

GetUtilizedSpaceThresholdOk returns a tuple with the UtilizedSpaceThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtilizedSpaceThreshold

`func (o *VirtualizationVmwareDatastoreCluster) SetUtilizedSpaceThreshold(v int32)`

SetUtilizedSpaceThreshold sets UtilizedSpaceThreshold field to given value.

### HasUtilizedSpaceThreshold

`func (o *VirtualizationVmwareDatastoreCluster) HasUtilizedSpaceThreshold() bool`

HasUtilizedSpaceThreshold returns a boolean if a field has been set.

### GetVmEvacuationAutomationMode

`func (o *VirtualizationVmwareDatastoreCluster) GetVmEvacuationAutomationMode() string`

GetVmEvacuationAutomationMode returns the VmEvacuationAutomationMode field if non-nil, zero value otherwise.

### GetVmEvacuationAutomationModeOk

`func (o *VirtualizationVmwareDatastoreCluster) GetVmEvacuationAutomationModeOk() (*string, bool)`

GetVmEvacuationAutomationModeOk returns a tuple with the VmEvacuationAutomationMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmEvacuationAutomationMode

`func (o *VirtualizationVmwareDatastoreCluster) SetVmEvacuationAutomationMode(v string)`

SetVmEvacuationAutomationMode sets VmEvacuationAutomationMode field to given value.

### HasVmEvacuationAutomationMode

`func (o *VirtualizationVmwareDatastoreCluster) HasVmEvacuationAutomationMode() bool`

HasVmEvacuationAutomationMode returns a boolean if a field has been set.

### GetDatacenter

`func (o *VirtualizationVmwareDatastoreCluster) GetDatacenter() VirtualizationVmwareDatacenterRelationship`

GetDatacenter returns the Datacenter field if non-nil, zero value otherwise.

### GetDatacenterOk

`func (o *VirtualizationVmwareDatastoreCluster) GetDatacenterOk() (*VirtualizationVmwareDatacenterRelationship, bool)`

GetDatacenterOk returns a tuple with the Datacenter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenter

`func (o *VirtualizationVmwareDatastoreCluster) SetDatacenter(v VirtualizationVmwareDatacenterRelationship)`

SetDatacenter sets Datacenter field to given value.

### HasDatacenter

`func (o *VirtualizationVmwareDatastoreCluster) HasDatacenter() bool`

HasDatacenter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


