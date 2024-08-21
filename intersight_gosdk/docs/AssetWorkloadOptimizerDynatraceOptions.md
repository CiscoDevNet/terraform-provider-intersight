# AssetWorkloadOptimizerDynatraceOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "asset.WorkloadOptimizerDynatraceOptions"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "asset.WorkloadOptimizerDynatraceOptions"]
**CollectTagInfo** | Pointer to **bool** | Collect tag information from Dynatrace. | [optional] [default to false]
**CollectVmMetrics** | Pointer to **bool** | Overwrite Hypervisor or Cloud Provider Virtual Machine metrics with data from the target. | [optional] 
**EnvironmentId** | Pointer to **string** | Each environment monitored with Dynatrace is identified with a unique character string—the environment ID. The Dynatrace API relies heavily on environment IDs to ensure that it pulls monitoring data from and pushes relevant external events to the correct Dynatrace environments. | [optional] 

## Methods

### NewAssetWorkloadOptimizerDynatraceOptions

`func NewAssetWorkloadOptimizerDynatraceOptions(classId string, objectType string, ) *AssetWorkloadOptimizerDynatraceOptions`

NewAssetWorkloadOptimizerDynatraceOptions instantiates a new AssetWorkloadOptimizerDynatraceOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetWorkloadOptimizerDynatraceOptionsWithDefaults

`func NewAssetWorkloadOptimizerDynatraceOptionsWithDefaults() *AssetWorkloadOptimizerDynatraceOptions`

NewAssetWorkloadOptimizerDynatraceOptionsWithDefaults instantiates a new AssetWorkloadOptimizerDynatraceOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *AssetWorkloadOptimizerDynatraceOptions) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *AssetWorkloadOptimizerDynatraceOptions) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *AssetWorkloadOptimizerDynatraceOptions) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *AssetWorkloadOptimizerDynatraceOptions) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *AssetWorkloadOptimizerDynatraceOptions) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *AssetWorkloadOptimizerDynatraceOptions) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCollectTagInfo

`func (o *AssetWorkloadOptimizerDynatraceOptions) GetCollectTagInfo() bool`

GetCollectTagInfo returns the CollectTagInfo field if non-nil, zero value otherwise.

### GetCollectTagInfoOk

`func (o *AssetWorkloadOptimizerDynatraceOptions) GetCollectTagInfoOk() (*bool, bool)`

GetCollectTagInfoOk returns a tuple with the CollectTagInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectTagInfo

`func (o *AssetWorkloadOptimizerDynatraceOptions) SetCollectTagInfo(v bool)`

SetCollectTagInfo sets CollectTagInfo field to given value.

### HasCollectTagInfo

`func (o *AssetWorkloadOptimizerDynatraceOptions) HasCollectTagInfo() bool`

HasCollectTagInfo returns a boolean if a field has been set.

### GetCollectVmMetrics

`func (o *AssetWorkloadOptimizerDynatraceOptions) GetCollectVmMetrics() bool`

GetCollectVmMetrics returns the CollectVmMetrics field if non-nil, zero value otherwise.

### GetCollectVmMetricsOk

`func (o *AssetWorkloadOptimizerDynatraceOptions) GetCollectVmMetricsOk() (*bool, bool)`

GetCollectVmMetricsOk returns a tuple with the CollectVmMetrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectVmMetrics

`func (o *AssetWorkloadOptimizerDynatraceOptions) SetCollectVmMetrics(v bool)`

SetCollectVmMetrics sets CollectVmMetrics field to given value.

### HasCollectVmMetrics

`func (o *AssetWorkloadOptimizerDynatraceOptions) HasCollectVmMetrics() bool`

HasCollectVmMetrics returns a boolean if a field has been set.

### GetEnvironmentId

`func (o *AssetWorkloadOptimizerDynatraceOptions) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *AssetWorkloadOptimizerDynatraceOptions) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *AssetWorkloadOptimizerDynatraceOptions) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.

### HasEnvironmentId

`func (o *AssetWorkloadOptimizerDynatraceOptions) HasEnvironmentId() bool`

HasEnvironmentId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


