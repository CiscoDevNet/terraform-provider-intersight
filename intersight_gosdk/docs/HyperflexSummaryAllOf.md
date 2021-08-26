# HyperflexSummaryAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hyperflex.Summary"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hyperflex.Summary"]
**ActiveNodes** | Pointer to **string** | The number of nodes currently participating in the storage cluster. | [optional] [readonly] 
**Address** | Pointer to **string** | The data IP address of the HyperFlex cluster. | [optional] [readonly] 
**Boottime** | Pointer to **int64** | The time taken during last cluster startup in seconds. | [optional] [readonly] 
**ClusterAccessPolicy** | Pointer to **string** | The cluster access policy for the HyperFlex cluster. An access policy of &#39;STRICT&#39; means that the cluster becomes readonly once any fragment of data is reduced to one copy. &#39;LENIENT&#39; means that the cluster stays in read-write mode even if any fragment of data is reduced to one copy. | [optional] [readonly] 
**CompressionSavings** | Pointer to **float64** | The percentage of storage space saved using data compression. | [optional] [readonly] 
**DataReplicationCompliance** | Pointer to **string** | The compliance with the data replication factor set for the HyperFlex cluster. | [optional] [readonly] 
**DataReplicationFactor** | Pointer to **string** | The number of data copies retained by the HyperFlex cluster. | [optional] [readonly] 
**DeduplicationSavings** | Pointer to **float64** | The percentage of storage space saved using data deduplication. | [optional] [readonly] 
**Downtime** | Pointer to **string** | The amount of time the HyperFlex cluster has been offline. | [optional] [readonly] 
**FreeCapacity** | Pointer to **int64** | The amount of storage capacity currently not in use, represented in bytes. | [optional] [readonly] 
**HealingInfo** | Pointer to [**NullableHyperflexStPlatformClusterHealingInfo**](HyperflexStPlatformClusterHealingInfo.md) |  | [optional] 
**Name** | Pointer to **string** | The name of the HyperFlex cluster. | [optional] [readonly] 
**ResiliencyDetails** | Pointer to **interface{}** | The details about the resiliency health of the cluster. Includes information about the cluster healing status and the storage cluster health. | [optional] [readonly] 
**ResiliencyDetailsSize** | Pointer to **int64** | The number of elements in the resiliency details property. | [optional] [readonly] 
**ResiliencyInfo** | Pointer to [**NullableHyperflexStPlatformClusterResiliencyInfo**](HyperflexStPlatformClusterResiliencyInfo.md) |  | [optional] 
**SpaceStatus** | Pointer to **string** | The space utilization status of the HyperFlex cluster. | [optional] [readonly] 
**State** | Pointer to **string** | The operational state of the HyperFlex cluster. | [optional] [readonly] 
**TotalCapacity** | Pointer to **int64** | The total amount of storage capacity available for the HyperFlex cluster, represented in bytes. | [optional] [readonly] 
**TotalSavings** | Pointer to **float64** | The percentage of storage space saved in total. | [optional] [readonly] 
**Uptime** | Pointer to **string** | The amount of time the HyperFlex cluster has been running since last startup. | [optional] [readonly] 
**UsedCapacity** | Pointer to **int64** | The amount of storage capacity in use, represented in bytes. | [optional] [readonly] 

## Methods

### NewHyperflexSummaryAllOf

`func NewHyperflexSummaryAllOf(classId string, objectType string, ) *HyperflexSummaryAllOf`

NewHyperflexSummaryAllOf instantiates a new HyperflexSummaryAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexSummaryAllOfWithDefaults

`func NewHyperflexSummaryAllOfWithDefaults() *HyperflexSummaryAllOf`

NewHyperflexSummaryAllOfWithDefaults instantiates a new HyperflexSummaryAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HyperflexSummaryAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexSummaryAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexSummaryAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HyperflexSummaryAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexSummaryAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexSummaryAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetActiveNodes

`func (o *HyperflexSummaryAllOf) GetActiveNodes() string`

GetActiveNodes returns the ActiveNodes field if non-nil, zero value otherwise.

### GetActiveNodesOk

`func (o *HyperflexSummaryAllOf) GetActiveNodesOk() (*string, bool)`

GetActiveNodesOk returns a tuple with the ActiveNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveNodes

`func (o *HyperflexSummaryAllOf) SetActiveNodes(v string)`

SetActiveNodes sets ActiveNodes field to given value.

### HasActiveNodes

`func (o *HyperflexSummaryAllOf) HasActiveNodes() bool`

HasActiveNodes returns a boolean if a field has been set.

### GetAddress

`func (o *HyperflexSummaryAllOf) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *HyperflexSummaryAllOf) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *HyperflexSummaryAllOf) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *HyperflexSummaryAllOf) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetBoottime

`func (o *HyperflexSummaryAllOf) GetBoottime() int64`

GetBoottime returns the Boottime field if non-nil, zero value otherwise.

### GetBoottimeOk

`func (o *HyperflexSummaryAllOf) GetBoottimeOk() (*int64, bool)`

GetBoottimeOk returns a tuple with the Boottime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoottime

`func (o *HyperflexSummaryAllOf) SetBoottime(v int64)`

SetBoottime sets Boottime field to given value.

### HasBoottime

`func (o *HyperflexSummaryAllOf) HasBoottime() bool`

HasBoottime returns a boolean if a field has been set.

### GetClusterAccessPolicy

`func (o *HyperflexSummaryAllOf) GetClusterAccessPolicy() string`

GetClusterAccessPolicy returns the ClusterAccessPolicy field if non-nil, zero value otherwise.

### GetClusterAccessPolicyOk

`func (o *HyperflexSummaryAllOf) GetClusterAccessPolicyOk() (*string, bool)`

GetClusterAccessPolicyOk returns a tuple with the ClusterAccessPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterAccessPolicy

`func (o *HyperflexSummaryAllOf) SetClusterAccessPolicy(v string)`

SetClusterAccessPolicy sets ClusterAccessPolicy field to given value.

### HasClusterAccessPolicy

`func (o *HyperflexSummaryAllOf) HasClusterAccessPolicy() bool`

HasClusterAccessPolicy returns a boolean if a field has been set.

### GetCompressionSavings

`func (o *HyperflexSummaryAllOf) GetCompressionSavings() float64`

GetCompressionSavings returns the CompressionSavings field if non-nil, zero value otherwise.

### GetCompressionSavingsOk

`func (o *HyperflexSummaryAllOf) GetCompressionSavingsOk() (*float64, bool)`

GetCompressionSavingsOk returns a tuple with the CompressionSavings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompressionSavings

`func (o *HyperflexSummaryAllOf) SetCompressionSavings(v float64)`

SetCompressionSavings sets CompressionSavings field to given value.

### HasCompressionSavings

`func (o *HyperflexSummaryAllOf) HasCompressionSavings() bool`

HasCompressionSavings returns a boolean if a field has been set.

### GetDataReplicationCompliance

`func (o *HyperflexSummaryAllOf) GetDataReplicationCompliance() string`

GetDataReplicationCompliance returns the DataReplicationCompliance field if non-nil, zero value otherwise.

### GetDataReplicationComplianceOk

`func (o *HyperflexSummaryAllOf) GetDataReplicationComplianceOk() (*string, bool)`

GetDataReplicationComplianceOk returns a tuple with the DataReplicationCompliance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataReplicationCompliance

`func (o *HyperflexSummaryAllOf) SetDataReplicationCompliance(v string)`

SetDataReplicationCompliance sets DataReplicationCompliance field to given value.

### HasDataReplicationCompliance

`func (o *HyperflexSummaryAllOf) HasDataReplicationCompliance() bool`

HasDataReplicationCompliance returns a boolean if a field has been set.

### GetDataReplicationFactor

`func (o *HyperflexSummaryAllOf) GetDataReplicationFactor() string`

GetDataReplicationFactor returns the DataReplicationFactor field if non-nil, zero value otherwise.

### GetDataReplicationFactorOk

`func (o *HyperflexSummaryAllOf) GetDataReplicationFactorOk() (*string, bool)`

GetDataReplicationFactorOk returns a tuple with the DataReplicationFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataReplicationFactor

`func (o *HyperflexSummaryAllOf) SetDataReplicationFactor(v string)`

SetDataReplicationFactor sets DataReplicationFactor field to given value.

### HasDataReplicationFactor

`func (o *HyperflexSummaryAllOf) HasDataReplicationFactor() bool`

HasDataReplicationFactor returns a boolean if a field has been set.

### GetDeduplicationSavings

`func (o *HyperflexSummaryAllOf) GetDeduplicationSavings() float64`

GetDeduplicationSavings returns the DeduplicationSavings field if non-nil, zero value otherwise.

### GetDeduplicationSavingsOk

`func (o *HyperflexSummaryAllOf) GetDeduplicationSavingsOk() (*float64, bool)`

GetDeduplicationSavingsOk returns a tuple with the DeduplicationSavings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeduplicationSavings

`func (o *HyperflexSummaryAllOf) SetDeduplicationSavings(v float64)`

SetDeduplicationSavings sets DeduplicationSavings field to given value.

### HasDeduplicationSavings

`func (o *HyperflexSummaryAllOf) HasDeduplicationSavings() bool`

HasDeduplicationSavings returns a boolean if a field has been set.

### GetDowntime

`func (o *HyperflexSummaryAllOf) GetDowntime() string`

GetDowntime returns the Downtime field if non-nil, zero value otherwise.

### GetDowntimeOk

`func (o *HyperflexSummaryAllOf) GetDowntimeOk() (*string, bool)`

GetDowntimeOk returns a tuple with the Downtime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDowntime

`func (o *HyperflexSummaryAllOf) SetDowntime(v string)`

SetDowntime sets Downtime field to given value.

### HasDowntime

`func (o *HyperflexSummaryAllOf) HasDowntime() bool`

HasDowntime returns a boolean if a field has been set.

### GetFreeCapacity

`func (o *HyperflexSummaryAllOf) GetFreeCapacity() int64`

GetFreeCapacity returns the FreeCapacity field if non-nil, zero value otherwise.

### GetFreeCapacityOk

`func (o *HyperflexSummaryAllOf) GetFreeCapacityOk() (*int64, bool)`

GetFreeCapacityOk returns a tuple with the FreeCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeCapacity

`func (o *HyperflexSummaryAllOf) SetFreeCapacity(v int64)`

SetFreeCapacity sets FreeCapacity field to given value.

### HasFreeCapacity

`func (o *HyperflexSummaryAllOf) HasFreeCapacity() bool`

HasFreeCapacity returns a boolean if a field has been set.

### GetHealingInfo

`func (o *HyperflexSummaryAllOf) GetHealingInfo() HyperflexStPlatformClusterHealingInfo`

GetHealingInfo returns the HealingInfo field if non-nil, zero value otherwise.

### GetHealingInfoOk

`func (o *HyperflexSummaryAllOf) GetHealingInfoOk() (*HyperflexStPlatformClusterHealingInfo, bool)`

GetHealingInfoOk returns a tuple with the HealingInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealingInfo

`func (o *HyperflexSummaryAllOf) SetHealingInfo(v HyperflexStPlatformClusterHealingInfo)`

SetHealingInfo sets HealingInfo field to given value.

### HasHealingInfo

`func (o *HyperflexSummaryAllOf) HasHealingInfo() bool`

HasHealingInfo returns a boolean if a field has been set.

### SetHealingInfoNil

`func (o *HyperflexSummaryAllOf) SetHealingInfoNil(b bool)`

 SetHealingInfoNil sets the value for HealingInfo to be an explicit nil

### UnsetHealingInfo
`func (o *HyperflexSummaryAllOf) UnsetHealingInfo()`

UnsetHealingInfo ensures that no value is present for HealingInfo, not even an explicit nil
### GetName

`func (o *HyperflexSummaryAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HyperflexSummaryAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HyperflexSummaryAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *HyperflexSummaryAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetResiliencyDetails

`func (o *HyperflexSummaryAllOf) GetResiliencyDetails() interface{}`

GetResiliencyDetails returns the ResiliencyDetails field if non-nil, zero value otherwise.

### GetResiliencyDetailsOk

`func (o *HyperflexSummaryAllOf) GetResiliencyDetailsOk() (*interface{}, bool)`

GetResiliencyDetailsOk returns a tuple with the ResiliencyDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResiliencyDetails

`func (o *HyperflexSummaryAllOf) SetResiliencyDetails(v interface{})`

SetResiliencyDetails sets ResiliencyDetails field to given value.

### HasResiliencyDetails

`func (o *HyperflexSummaryAllOf) HasResiliencyDetails() bool`

HasResiliencyDetails returns a boolean if a field has been set.

### SetResiliencyDetailsNil

`func (o *HyperflexSummaryAllOf) SetResiliencyDetailsNil(b bool)`

 SetResiliencyDetailsNil sets the value for ResiliencyDetails to be an explicit nil

### UnsetResiliencyDetails
`func (o *HyperflexSummaryAllOf) UnsetResiliencyDetails()`

UnsetResiliencyDetails ensures that no value is present for ResiliencyDetails, not even an explicit nil
### GetResiliencyDetailsSize

`func (o *HyperflexSummaryAllOf) GetResiliencyDetailsSize() int64`

GetResiliencyDetailsSize returns the ResiliencyDetailsSize field if non-nil, zero value otherwise.

### GetResiliencyDetailsSizeOk

`func (o *HyperflexSummaryAllOf) GetResiliencyDetailsSizeOk() (*int64, bool)`

GetResiliencyDetailsSizeOk returns a tuple with the ResiliencyDetailsSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResiliencyDetailsSize

`func (o *HyperflexSummaryAllOf) SetResiliencyDetailsSize(v int64)`

SetResiliencyDetailsSize sets ResiliencyDetailsSize field to given value.

### HasResiliencyDetailsSize

`func (o *HyperflexSummaryAllOf) HasResiliencyDetailsSize() bool`

HasResiliencyDetailsSize returns a boolean if a field has been set.

### GetResiliencyInfo

`func (o *HyperflexSummaryAllOf) GetResiliencyInfo() HyperflexStPlatformClusterResiliencyInfo`

GetResiliencyInfo returns the ResiliencyInfo field if non-nil, zero value otherwise.

### GetResiliencyInfoOk

`func (o *HyperflexSummaryAllOf) GetResiliencyInfoOk() (*HyperflexStPlatformClusterResiliencyInfo, bool)`

GetResiliencyInfoOk returns a tuple with the ResiliencyInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResiliencyInfo

`func (o *HyperflexSummaryAllOf) SetResiliencyInfo(v HyperflexStPlatformClusterResiliencyInfo)`

SetResiliencyInfo sets ResiliencyInfo field to given value.

### HasResiliencyInfo

`func (o *HyperflexSummaryAllOf) HasResiliencyInfo() bool`

HasResiliencyInfo returns a boolean if a field has been set.

### SetResiliencyInfoNil

`func (o *HyperflexSummaryAllOf) SetResiliencyInfoNil(b bool)`

 SetResiliencyInfoNil sets the value for ResiliencyInfo to be an explicit nil

### UnsetResiliencyInfo
`func (o *HyperflexSummaryAllOf) UnsetResiliencyInfo()`

UnsetResiliencyInfo ensures that no value is present for ResiliencyInfo, not even an explicit nil
### GetSpaceStatus

`func (o *HyperflexSummaryAllOf) GetSpaceStatus() string`

GetSpaceStatus returns the SpaceStatus field if non-nil, zero value otherwise.

### GetSpaceStatusOk

`func (o *HyperflexSummaryAllOf) GetSpaceStatusOk() (*string, bool)`

GetSpaceStatusOk returns a tuple with the SpaceStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceStatus

`func (o *HyperflexSummaryAllOf) SetSpaceStatus(v string)`

SetSpaceStatus sets SpaceStatus field to given value.

### HasSpaceStatus

`func (o *HyperflexSummaryAllOf) HasSpaceStatus() bool`

HasSpaceStatus returns a boolean if a field has been set.

### GetState

`func (o *HyperflexSummaryAllOf) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *HyperflexSummaryAllOf) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *HyperflexSummaryAllOf) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *HyperflexSummaryAllOf) HasState() bool`

HasState returns a boolean if a field has been set.

### GetTotalCapacity

`func (o *HyperflexSummaryAllOf) GetTotalCapacity() int64`

GetTotalCapacity returns the TotalCapacity field if non-nil, zero value otherwise.

### GetTotalCapacityOk

`func (o *HyperflexSummaryAllOf) GetTotalCapacityOk() (*int64, bool)`

GetTotalCapacityOk returns a tuple with the TotalCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCapacity

`func (o *HyperflexSummaryAllOf) SetTotalCapacity(v int64)`

SetTotalCapacity sets TotalCapacity field to given value.

### HasTotalCapacity

`func (o *HyperflexSummaryAllOf) HasTotalCapacity() bool`

HasTotalCapacity returns a boolean if a field has been set.

### GetTotalSavings

`func (o *HyperflexSummaryAllOf) GetTotalSavings() float64`

GetTotalSavings returns the TotalSavings field if non-nil, zero value otherwise.

### GetTotalSavingsOk

`func (o *HyperflexSummaryAllOf) GetTotalSavingsOk() (*float64, bool)`

GetTotalSavingsOk returns a tuple with the TotalSavings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSavings

`func (o *HyperflexSummaryAllOf) SetTotalSavings(v float64)`

SetTotalSavings sets TotalSavings field to given value.

### HasTotalSavings

`func (o *HyperflexSummaryAllOf) HasTotalSavings() bool`

HasTotalSavings returns a boolean if a field has been set.

### GetUptime

`func (o *HyperflexSummaryAllOf) GetUptime() string`

GetUptime returns the Uptime field if non-nil, zero value otherwise.

### GetUptimeOk

`func (o *HyperflexSummaryAllOf) GetUptimeOk() (*string, bool)`

GetUptimeOk returns a tuple with the Uptime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUptime

`func (o *HyperflexSummaryAllOf) SetUptime(v string)`

SetUptime sets Uptime field to given value.

### HasUptime

`func (o *HyperflexSummaryAllOf) HasUptime() bool`

HasUptime returns a boolean if a field has been set.

### GetUsedCapacity

`func (o *HyperflexSummaryAllOf) GetUsedCapacity() int64`

GetUsedCapacity returns the UsedCapacity field if non-nil, zero value otherwise.

### GetUsedCapacityOk

`func (o *HyperflexSummaryAllOf) GetUsedCapacityOk() (*int64, bool)`

GetUsedCapacityOk returns a tuple with the UsedCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedCapacity

`func (o *HyperflexSummaryAllOf) SetUsedCapacity(v int64)`

SetUsedCapacity sets UsedCapacity field to given value.

### HasUsedCapacity

`func (o *HyperflexSummaryAllOf) HasUsedCapacity() bool`

HasUsedCapacity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


