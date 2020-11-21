# AssetWorkloadOptimizerAmazonWebServicesBillingOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "asset.WorkloadOptimizerAmazonWebServicesBillingOptions"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "asset.WorkloadOptimizerAmazonWebServicesBillingOptions"]
**CostAndUsageReportBucket** | Pointer to **string** | Name of S3 bucket that contains the Amazon web service Cost and Usage report to get cloud service spend. | [optional] 
**CostAndUsageReportPath** | Pointer to **string** | Report path prefix for the Amazon web service cost and usage report to get cloud service spend. | [optional] 
**CostAndUsageReportRegion** | Pointer to **string** | Region for S3 bucket that contains the Amazon web service Cost and Usage report to get cloud service spend. | [optional] 

## Methods

### NewAssetWorkloadOptimizerAmazonWebServicesBillingOptions

`func NewAssetWorkloadOptimizerAmazonWebServicesBillingOptions(classId string, objectType string, ) *AssetWorkloadOptimizerAmazonWebServicesBillingOptions`

NewAssetWorkloadOptimizerAmazonWebServicesBillingOptions instantiates a new AssetWorkloadOptimizerAmazonWebServicesBillingOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetWorkloadOptimizerAmazonWebServicesBillingOptionsWithDefaults

`func NewAssetWorkloadOptimizerAmazonWebServicesBillingOptionsWithDefaults() *AssetWorkloadOptimizerAmazonWebServicesBillingOptions`

NewAssetWorkloadOptimizerAmazonWebServicesBillingOptionsWithDefaults instantiates a new AssetWorkloadOptimizerAmazonWebServicesBillingOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *AssetWorkloadOptimizerAmazonWebServicesBillingOptions) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *AssetWorkloadOptimizerAmazonWebServicesBillingOptions) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *AssetWorkloadOptimizerAmazonWebServicesBillingOptions) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *AssetWorkloadOptimizerAmazonWebServicesBillingOptions) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *AssetWorkloadOptimizerAmazonWebServicesBillingOptions) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *AssetWorkloadOptimizerAmazonWebServicesBillingOptions) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCostAndUsageReportBucket

`func (o *AssetWorkloadOptimizerAmazonWebServicesBillingOptions) GetCostAndUsageReportBucket() string`

GetCostAndUsageReportBucket returns the CostAndUsageReportBucket field if non-nil, zero value otherwise.

### GetCostAndUsageReportBucketOk

`func (o *AssetWorkloadOptimizerAmazonWebServicesBillingOptions) GetCostAndUsageReportBucketOk() (*string, bool)`

GetCostAndUsageReportBucketOk returns a tuple with the CostAndUsageReportBucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostAndUsageReportBucket

`func (o *AssetWorkloadOptimizerAmazonWebServicesBillingOptions) SetCostAndUsageReportBucket(v string)`

SetCostAndUsageReportBucket sets CostAndUsageReportBucket field to given value.

### HasCostAndUsageReportBucket

`func (o *AssetWorkloadOptimizerAmazonWebServicesBillingOptions) HasCostAndUsageReportBucket() bool`

HasCostAndUsageReportBucket returns a boolean if a field has been set.

### GetCostAndUsageReportPath

`func (o *AssetWorkloadOptimizerAmazonWebServicesBillingOptions) GetCostAndUsageReportPath() string`

GetCostAndUsageReportPath returns the CostAndUsageReportPath field if non-nil, zero value otherwise.

### GetCostAndUsageReportPathOk

`func (o *AssetWorkloadOptimizerAmazonWebServicesBillingOptions) GetCostAndUsageReportPathOk() (*string, bool)`

GetCostAndUsageReportPathOk returns a tuple with the CostAndUsageReportPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostAndUsageReportPath

`func (o *AssetWorkloadOptimizerAmazonWebServicesBillingOptions) SetCostAndUsageReportPath(v string)`

SetCostAndUsageReportPath sets CostAndUsageReportPath field to given value.

### HasCostAndUsageReportPath

`func (o *AssetWorkloadOptimizerAmazonWebServicesBillingOptions) HasCostAndUsageReportPath() bool`

HasCostAndUsageReportPath returns a boolean if a field has been set.

### GetCostAndUsageReportRegion

`func (o *AssetWorkloadOptimizerAmazonWebServicesBillingOptions) GetCostAndUsageReportRegion() string`

GetCostAndUsageReportRegion returns the CostAndUsageReportRegion field if non-nil, zero value otherwise.

### GetCostAndUsageReportRegionOk

`func (o *AssetWorkloadOptimizerAmazonWebServicesBillingOptions) GetCostAndUsageReportRegionOk() (*string, bool)`

GetCostAndUsageReportRegionOk returns a tuple with the CostAndUsageReportRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostAndUsageReportRegion

`func (o *AssetWorkloadOptimizerAmazonWebServicesBillingOptions) SetCostAndUsageReportRegion(v string)`

SetCostAndUsageReportRegion sets CostAndUsageReportRegion field to given value.

### HasCostAndUsageReportRegion

`func (o *AssetWorkloadOptimizerAmazonWebServicesBillingOptions) HasCostAndUsageReportRegion() bool`

HasCostAndUsageReportRegion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


