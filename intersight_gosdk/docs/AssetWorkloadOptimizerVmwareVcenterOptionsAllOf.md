# AssetWorkloadOptimizerVmwareVcenterOptionsAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "asset.WorkloadOptimizerVmwareVcenterOptions"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "asset.WorkloadOptimizerVmwareVcenterOptions"]
**DatastoreBrowsingEnabled** | Pointer to **bool** | DatastoreBrowsingEnabled controls whether Workload Optimizer scans vCenter datastores to identify files which are not used and can be deleted to reclaim space and improve actual disk utilization. For example orphaned VMDK files. | [optional] 
**GuestMetricsEnabled** | Pointer to **bool** | Enable retrieval of advanced memory metrics. Only supported on vCenter Server version 6.5U3 or later. Guest VMs must run VMWare Tools 10.3.2 Build 10338 or later. | [optional] 

## Methods

### NewAssetWorkloadOptimizerVmwareVcenterOptionsAllOf

`func NewAssetWorkloadOptimizerVmwareVcenterOptionsAllOf(classId string, objectType string, ) *AssetWorkloadOptimizerVmwareVcenterOptionsAllOf`

NewAssetWorkloadOptimizerVmwareVcenterOptionsAllOf instantiates a new AssetWorkloadOptimizerVmwareVcenterOptionsAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetWorkloadOptimizerVmwareVcenterOptionsAllOfWithDefaults

`func NewAssetWorkloadOptimizerVmwareVcenterOptionsAllOfWithDefaults() *AssetWorkloadOptimizerVmwareVcenterOptionsAllOf`

NewAssetWorkloadOptimizerVmwareVcenterOptionsAllOfWithDefaults instantiates a new AssetWorkloadOptimizerVmwareVcenterOptionsAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *AssetWorkloadOptimizerVmwareVcenterOptionsAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *AssetWorkloadOptimizerVmwareVcenterOptionsAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *AssetWorkloadOptimizerVmwareVcenterOptionsAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *AssetWorkloadOptimizerVmwareVcenterOptionsAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *AssetWorkloadOptimizerVmwareVcenterOptionsAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *AssetWorkloadOptimizerVmwareVcenterOptionsAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDatastoreBrowsingEnabled

`func (o *AssetWorkloadOptimizerVmwareVcenterOptionsAllOf) GetDatastoreBrowsingEnabled() bool`

GetDatastoreBrowsingEnabled returns the DatastoreBrowsingEnabled field if non-nil, zero value otherwise.

### GetDatastoreBrowsingEnabledOk

`func (o *AssetWorkloadOptimizerVmwareVcenterOptionsAllOf) GetDatastoreBrowsingEnabledOk() (*bool, bool)`

GetDatastoreBrowsingEnabledOk returns a tuple with the DatastoreBrowsingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastoreBrowsingEnabled

`func (o *AssetWorkloadOptimizerVmwareVcenterOptionsAllOf) SetDatastoreBrowsingEnabled(v bool)`

SetDatastoreBrowsingEnabled sets DatastoreBrowsingEnabled field to given value.

### HasDatastoreBrowsingEnabled

`func (o *AssetWorkloadOptimizerVmwareVcenterOptionsAllOf) HasDatastoreBrowsingEnabled() bool`

HasDatastoreBrowsingEnabled returns a boolean if a field has been set.

### GetGuestMetricsEnabled

`func (o *AssetWorkloadOptimizerVmwareVcenterOptionsAllOf) GetGuestMetricsEnabled() bool`

GetGuestMetricsEnabled returns the GuestMetricsEnabled field if non-nil, zero value otherwise.

### GetGuestMetricsEnabledOk

`func (o *AssetWorkloadOptimizerVmwareVcenterOptionsAllOf) GetGuestMetricsEnabledOk() (*bool, bool)`

GetGuestMetricsEnabledOk returns a tuple with the GuestMetricsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestMetricsEnabled

`func (o *AssetWorkloadOptimizerVmwareVcenterOptionsAllOf) SetGuestMetricsEnabled(v bool)`

SetGuestMetricsEnabled sets GuestMetricsEnabled field to given value.

### HasGuestMetricsEnabled

`func (o *AssetWorkloadOptimizerVmwareVcenterOptionsAllOf) HasGuestMetricsEnabled() bool`

HasGuestMetricsEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


