# AssetWorkloadOptimizerVmwareVcenterOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "asset.WorkloadOptimizerVmwareVcenterOptions"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "asset.WorkloadOptimizerVmwareVcenterOptions"]
**DatastoreBrowsingEnabled** | Pointer to **bool** | DatastoreBrowsingEnabled controls whether Workload Optimizer scans vCenter datastores to identify files which are not used and can be deleted to reclaim space and improve actual disk utilization. For example orphaned VMDK files. | [optional] 

## Methods

### NewAssetWorkloadOptimizerVmwareVcenterOptions

`func NewAssetWorkloadOptimizerVmwareVcenterOptions(classId string, objectType string, ) *AssetWorkloadOptimizerVmwareVcenterOptions`

NewAssetWorkloadOptimizerVmwareVcenterOptions instantiates a new AssetWorkloadOptimizerVmwareVcenterOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetWorkloadOptimizerVmwareVcenterOptionsWithDefaults

`func NewAssetWorkloadOptimizerVmwareVcenterOptionsWithDefaults() *AssetWorkloadOptimizerVmwareVcenterOptions`

NewAssetWorkloadOptimizerVmwareVcenterOptionsWithDefaults instantiates a new AssetWorkloadOptimizerVmwareVcenterOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *AssetWorkloadOptimizerVmwareVcenterOptions) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *AssetWorkloadOptimizerVmwareVcenterOptions) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *AssetWorkloadOptimizerVmwareVcenterOptions) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *AssetWorkloadOptimizerVmwareVcenterOptions) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *AssetWorkloadOptimizerVmwareVcenterOptions) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *AssetWorkloadOptimizerVmwareVcenterOptions) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDatastoreBrowsingEnabled

`func (o *AssetWorkloadOptimizerVmwareVcenterOptions) GetDatastoreBrowsingEnabled() bool`

GetDatastoreBrowsingEnabled returns the DatastoreBrowsingEnabled field if non-nil, zero value otherwise.

### GetDatastoreBrowsingEnabledOk

`func (o *AssetWorkloadOptimizerVmwareVcenterOptions) GetDatastoreBrowsingEnabledOk() (*bool, bool)`

GetDatastoreBrowsingEnabledOk returns a tuple with the DatastoreBrowsingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastoreBrowsingEnabled

`func (o *AssetWorkloadOptimizerVmwareVcenterOptions) SetDatastoreBrowsingEnabled(v bool)`

SetDatastoreBrowsingEnabled sets DatastoreBrowsingEnabled field to given value.

### HasDatastoreBrowsingEnabled

`func (o *AssetWorkloadOptimizerVmwareVcenterOptions) HasDatastoreBrowsingEnabled() bool`

HasDatastoreBrowsingEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


