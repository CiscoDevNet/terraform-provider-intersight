# HyperflexSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveNodes** | Pointer to **string** |  | [optional] [readonly] 
**Address** | Pointer to **string** |  | [optional] [readonly] 
**Boottime** | Pointer to **int64** |  | [optional] [readonly] 
**ClusterAccessPolicy** | Pointer to **string** |  | [optional] [readonly] 
**CompressionSavings** | Pointer to **float64** |  | [optional] [readonly] 
**DataReplicationCompliance** | Pointer to **string** |  | [optional] [readonly] 
**DataReplicationFactor** | Pointer to **string** |  | [optional] [readonly] 
**DeduplicationSavings** | Pointer to **float64** |  | [optional] [readonly] 
**Downtime** | Pointer to **string** |  | [optional] [readonly] 
**FreeCapacity** | Pointer to **int64** |  | [optional] [readonly] 
**HealingInfo** | Pointer to [**HyperflexStPlatformClusterHealingInfo**](hyperflex.StPlatformClusterHealingInfo.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] [readonly] 
**ResiliencyDetails** | Pointer to **map[string]interface{}** |  | [optional] [readonly] 
**ResiliencyDetailsSize** | Pointer to **int64** |  | [optional] [readonly] 
**ResiliencyInfo** | Pointer to [**HyperflexStPlatformClusterResiliencyInfo**](hyperflex.StPlatformClusterResiliencyInfo.md) |  | [optional] 
**SpaceStatus** | Pointer to **string** |  | [optional] [readonly] 
**State** | Pointer to **string** |  | [optional] [readonly] 
**TotalCapacity** | Pointer to **int64** |  | [optional] [readonly] 
**TotalSavings** | Pointer to **float64** |  | [optional] [readonly] 
**Uptime** | Pointer to **string** |  | [optional] [readonly] 
**UsedCapacity** | Pointer to **int64** |  | [optional] [readonly] 

## Methods

### NewHyperflexSummary

`func NewHyperflexSummary() *HyperflexSummary`

NewHyperflexSummary instantiates a new HyperflexSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexSummaryWithDefaults

`func NewHyperflexSummaryWithDefaults() *HyperflexSummary`

NewHyperflexSummaryWithDefaults instantiates a new HyperflexSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActiveNodes

`func (o *HyperflexSummary) GetActiveNodes() string`

GetActiveNodes returns the ActiveNodes field if non-nil, zero value otherwise.

### GetActiveNodesOk

`func (o *HyperflexSummary) GetActiveNodesOk() (*string, bool)`

GetActiveNodesOk returns a tuple with the ActiveNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveNodes

`func (o *HyperflexSummary) SetActiveNodes(v string)`

SetActiveNodes sets ActiveNodes field to given value.

### HasActiveNodes

`func (o *HyperflexSummary) HasActiveNodes() bool`

HasActiveNodes returns a boolean if a field has been set.

### GetAddress

`func (o *HyperflexSummary) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *HyperflexSummary) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *HyperflexSummary) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *HyperflexSummary) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetBoottime

`func (o *HyperflexSummary) GetBoottime() int64`

GetBoottime returns the Boottime field if non-nil, zero value otherwise.

### GetBoottimeOk

`func (o *HyperflexSummary) GetBoottimeOk() (*int64, bool)`

GetBoottimeOk returns a tuple with the Boottime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoottime

`func (o *HyperflexSummary) SetBoottime(v int64)`

SetBoottime sets Boottime field to given value.

### HasBoottime

`func (o *HyperflexSummary) HasBoottime() bool`

HasBoottime returns a boolean if a field has been set.

### GetClusterAccessPolicy

`func (o *HyperflexSummary) GetClusterAccessPolicy() string`

GetClusterAccessPolicy returns the ClusterAccessPolicy field if non-nil, zero value otherwise.

### GetClusterAccessPolicyOk

`func (o *HyperflexSummary) GetClusterAccessPolicyOk() (*string, bool)`

GetClusterAccessPolicyOk returns a tuple with the ClusterAccessPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterAccessPolicy

`func (o *HyperflexSummary) SetClusterAccessPolicy(v string)`

SetClusterAccessPolicy sets ClusterAccessPolicy field to given value.

### HasClusterAccessPolicy

`func (o *HyperflexSummary) HasClusterAccessPolicy() bool`

HasClusterAccessPolicy returns a boolean if a field has been set.

### GetCompressionSavings

`func (o *HyperflexSummary) GetCompressionSavings() float64`

GetCompressionSavings returns the CompressionSavings field if non-nil, zero value otherwise.

### GetCompressionSavingsOk

`func (o *HyperflexSummary) GetCompressionSavingsOk() (*float64, bool)`

GetCompressionSavingsOk returns a tuple with the CompressionSavings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompressionSavings

`func (o *HyperflexSummary) SetCompressionSavings(v float64)`

SetCompressionSavings sets CompressionSavings field to given value.

### HasCompressionSavings

`func (o *HyperflexSummary) HasCompressionSavings() bool`

HasCompressionSavings returns a boolean if a field has been set.

### GetDataReplicationCompliance

`func (o *HyperflexSummary) GetDataReplicationCompliance() string`

GetDataReplicationCompliance returns the DataReplicationCompliance field if non-nil, zero value otherwise.

### GetDataReplicationComplianceOk

`func (o *HyperflexSummary) GetDataReplicationComplianceOk() (*string, bool)`

GetDataReplicationComplianceOk returns a tuple with the DataReplicationCompliance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataReplicationCompliance

`func (o *HyperflexSummary) SetDataReplicationCompliance(v string)`

SetDataReplicationCompliance sets DataReplicationCompliance field to given value.

### HasDataReplicationCompliance

`func (o *HyperflexSummary) HasDataReplicationCompliance() bool`

HasDataReplicationCompliance returns a boolean if a field has been set.

### GetDataReplicationFactor

`func (o *HyperflexSummary) GetDataReplicationFactor() string`

GetDataReplicationFactor returns the DataReplicationFactor field if non-nil, zero value otherwise.

### GetDataReplicationFactorOk

`func (o *HyperflexSummary) GetDataReplicationFactorOk() (*string, bool)`

GetDataReplicationFactorOk returns a tuple with the DataReplicationFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataReplicationFactor

`func (o *HyperflexSummary) SetDataReplicationFactor(v string)`

SetDataReplicationFactor sets DataReplicationFactor field to given value.

### HasDataReplicationFactor

`func (o *HyperflexSummary) HasDataReplicationFactor() bool`

HasDataReplicationFactor returns a boolean if a field has been set.

### GetDeduplicationSavings

`func (o *HyperflexSummary) GetDeduplicationSavings() float64`

GetDeduplicationSavings returns the DeduplicationSavings field if non-nil, zero value otherwise.

### GetDeduplicationSavingsOk

`func (o *HyperflexSummary) GetDeduplicationSavingsOk() (*float64, bool)`

GetDeduplicationSavingsOk returns a tuple with the DeduplicationSavings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeduplicationSavings

`func (o *HyperflexSummary) SetDeduplicationSavings(v float64)`

SetDeduplicationSavings sets DeduplicationSavings field to given value.

### HasDeduplicationSavings

`func (o *HyperflexSummary) HasDeduplicationSavings() bool`

HasDeduplicationSavings returns a boolean if a field has been set.

### GetDowntime

`func (o *HyperflexSummary) GetDowntime() string`

GetDowntime returns the Downtime field if non-nil, zero value otherwise.

### GetDowntimeOk

`func (o *HyperflexSummary) GetDowntimeOk() (*string, bool)`

GetDowntimeOk returns a tuple with the Downtime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDowntime

`func (o *HyperflexSummary) SetDowntime(v string)`

SetDowntime sets Downtime field to given value.

### HasDowntime

`func (o *HyperflexSummary) HasDowntime() bool`

HasDowntime returns a boolean if a field has been set.

### GetFreeCapacity

`func (o *HyperflexSummary) GetFreeCapacity() int64`

GetFreeCapacity returns the FreeCapacity field if non-nil, zero value otherwise.

### GetFreeCapacityOk

`func (o *HyperflexSummary) GetFreeCapacityOk() (*int64, bool)`

GetFreeCapacityOk returns a tuple with the FreeCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeCapacity

`func (o *HyperflexSummary) SetFreeCapacity(v int64)`

SetFreeCapacity sets FreeCapacity field to given value.

### HasFreeCapacity

`func (o *HyperflexSummary) HasFreeCapacity() bool`

HasFreeCapacity returns a boolean if a field has been set.

### GetHealingInfo

`func (o *HyperflexSummary) GetHealingInfo() HyperflexStPlatformClusterHealingInfo`

GetHealingInfo returns the HealingInfo field if non-nil, zero value otherwise.

### GetHealingInfoOk

`func (o *HyperflexSummary) GetHealingInfoOk() (*HyperflexStPlatformClusterHealingInfo, bool)`

GetHealingInfoOk returns a tuple with the HealingInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealingInfo

`func (o *HyperflexSummary) SetHealingInfo(v HyperflexStPlatformClusterHealingInfo)`

SetHealingInfo sets HealingInfo field to given value.

### HasHealingInfo

`func (o *HyperflexSummary) HasHealingInfo() bool`

HasHealingInfo returns a boolean if a field has been set.

### GetName

`func (o *HyperflexSummary) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HyperflexSummary) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HyperflexSummary) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *HyperflexSummary) HasName() bool`

HasName returns a boolean if a field has been set.

### GetResiliencyDetails

`func (o *HyperflexSummary) GetResiliencyDetails() map[string]interface{}`

GetResiliencyDetails returns the ResiliencyDetails field if non-nil, zero value otherwise.

### GetResiliencyDetailsOk

`func (o *HyperflexSummary) GetResiliencyDetailsOk() (*map[string]interface{}, bool)`

GetResiliencyDetailsOk returns a tuple with the ResiliencyDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResiliencyDetails

`func (o *HyperflexSummary) SetResiliencyDetails(v map[string]interface{})`

SetResiliencyDetails sets ResiliencyDetails field to given value.

### HasResiliencyDetails

`func (o *HyperflexSummary) HasResiliencyDetails() bool`

HasResiliencyDetails returns a boolean if a field has been set.

### GetResiliencyDetailsSize

`func (o *HyperflexSummary) GetResiliencyDetailsSize() int64`

GetResiliencyDetailsSize returns the ResiliencyDetailsSize field if non-nil, zero value otherwise.

### GetResiliencyDetailsSizeOk

`func (o *HyperflexSummary) GetResiliencyDetailsSizeOk() (*int64, bool)`

GetResiliencyDetailsSizeOk returns a tuple with the ResiliencyDetailsSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResiliencyDetailsSize

`func (o *HyperflexSummary) SetResiliencyDetailsSize(v int64)`

SetResiliencyDetailsSize sets ResiliencyDetailsSize field to given value.

### HasResiliencyDetailsSize

`func (o *HyperflexSummary) HasResiliencyDetailsSize() bool`

HasResiliencyDetailsSize returns a boolean if a field has been set.

### GetResiliencyInfo

`func (o *HyperflexSummary) GetResiliencyInfo() HyperflexStPlatformClusterResiliencyInfo`

GetResiliencyInfo returns the ResiliencyInfo field if non-nil, zero value otherwise.

### GetResiliencyInfoOk

`func (o *HyperflexSummary) GetResiliencyInfoOk() (*HyperflexStPlatformClusterResiliencyInfo, bool)`

GetResiliencyInfoOk returns a tuple with the ResiliencyInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResiliencyInfo

`func (o *HyperflexSummary) SetResiliencyInfo(v HyperflexStPlatformClusterResiliencyInfo)`

SetResiliencyInfo sets ResiliencyInfo field to given value.

### HasResiliencyInfo

`func (o *HyperflexSummary) HasResiliencyInfo() bool`

HasResiliencyInfo returns a boolean if a field has been set.

### GetSpaceStatus

`func (o *HyperflexSummary) GetSpaceStatus() string`

GetSpaceStatus returns the SpaceStatus field if non-nil, zero value otherwise.

### GetSpaceStatusOk

`func (o *HyperflexSummary) GetSpaceStatusOk() (*string, bool)`

GetSpaceStatusOk returns a tuple with the SpaceStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceStatus

`func (o *HyperflexSummary) SetSpaceStatus(v string)`

SetSpaceStatus sets SpaceStatus field to given value.

### HasSpaceStatus

`func (o *HyperflexSummary) HasSpaceStatus() bool`

HasSpaceStatus returns a boolean if a field has been set.

### GetState

`func (o *HyperflexSummary) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *HyperflexSummary) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *HyperflexSummary) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *HyperflexSummary) HasState() bool`

HasState returns a boolean if a field has been set.

### GetTotalCapacity

`func (o *HyperflexSummary) GetTotalCapacity() int64`

GetTotalCapacity returns the TotalCapacity field if non-nil, zero value otherwise.

### GetTotalCapacityOk

`func (o *HyperflexSummary) GetTotalCapacityOk() (*int64, bool)`

GetTotalCapacityOk returns a tuple with the TotalCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCapacity

`func (o *HyperflexSummary) SetTotalCapacity(v int64)`

SetTotalCapacity sets TotalCapacity field to given value.

### HasTotalCapacity

`func (o *HyperflexSummary) HasTotalCapacity() bool`

HasTotalCapacity returns a boolean if a field has been set.

### GetTotalSavings

`func (o *HyperflexSummary) GetTotalSavings() float64`

GetTotalSavings returns the TotalSavings field if non-nil, zero value otherwise.

### GetTotalSavingsOk

`func (o *HyperflexSummary) GetTotalSavingsOk() (*float64, bool)`

GetTotalSavingsOk returns a tuple with the TotalSavings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSavings

`func (o *HyperflexSummary) SetTotalSavings(v float64)`

SetTotalSavings sets TotalSavings field to given value.

### HasTotalSavings

`func (o *HyperflexSummary) HasTotalSavings() bool`

HasTotalSavings returns a boolean if a field has been set.

### GetUptime

`func (o *HyperflexSummary) GetUptime() string`

GetUptime returns the Uptime field if non-nil, zero value otherwise.

### GetUptimeOk

`func (o *HyperflexSummary) GetUptimeOk() (*string, bool)`

GetUptimeOk returns a tuple with the Uptime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUptime

`func (o *HyperflexSummary) SetUptime(v string)`

SetUptime sets Uptime field to given value.

### HasUptime

`func (o *HyperflexSummary) HasUptime() bool`

HasUptime returns a boolean if a field has been set.

### GetUsedCapacity

`func (o *HyperflexSummary) GetUsedCapacity() int64`

GetUsedCapacity returns the UsedCapacity field if non-nil, zero value otherwise.

### GetUsedCapacityOk

`func (o *HyperflexSummary) GetUsedCapacityOk() (*int64, bool)`

GetUsedCapacityOk returns a tuple with the UsedCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedCapacity

`func (o *HyperflexSummary) SetUsedCapacity(v int64)`

SetUsedCapacity sets UsedCapacity field to given value.

### HasUsedCapacity

`func (o *HyperflexSummary) HasUsedCapacity() bool`

HasUsedCapacity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


