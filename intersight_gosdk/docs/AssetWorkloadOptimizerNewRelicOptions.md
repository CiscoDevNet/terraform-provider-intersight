# AssetWorkloadOptimizerNewRelicOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "asset.WorkloadOptimizerNewRelicOptions"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "asset.WorkloadOptimizerNewRelicOptions"]
**AccountId** | Pointer to **string** | Your NewRelic account id. | [optional] 
**Region** | Pointer to **string** | The region associated with the NewRelic account. * &#x60;US&#x60; - The United States (US) region. * &#x60;EU&#x60; - The European Union (EU) region. | [optional] [default to "US"]

## Methods

### NewAssetWorkloadOptimizerNewRelicOptions

`func NewAssetWorkloadOptimizerNewRelicOptions(classId string, objectType string, ) *AssetWorkloadOptimizerNewRelicOptions`

NewAssetWorkloadOptimizerNewRelicOptions instantiates a new AssetWorkloadOptimizerNewRelicOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetWorkloadOptimizerNewRelicOptionsWithDefaults

`func NewAssetWorkloadOptimizerNewRelicOptionsWithDefaults() *AssetWorkloadOptimizerNewRelicOptions`

NewAssetWorkloadOptimizerNewRelicOptionsWithDefaults instantiates a new AssetWorkloadOptimizerNewRelicOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *AssetWorkloadOptimizerNewRelicOptions) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *AssetWorkloadOptimizerNewRelicOptions) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *AssetWorkloadOptimizerNewRelicOptions) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *AssetWorkloadOptimizerNewRelicOptions) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *AssetWorkloadOptimizerNewRelicOptions) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *AssetWorkloadOptimizerNewRelicOptions) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAccountId

`func (o *AssetWorkloadOptimizerNewRelicOptions) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *AssetWorkloadOptimizerNewRelicOptions) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *AssetWorkloadOptimizerNewRelicOptions) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *AssetWorkloadOptimizerNewRelicOptions) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetRegion

`func (o *AssetWorkloadOptimizerNewRelicOptions) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *AssetWorkloadOptimizerNewRelicOptions) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *AssetWorkloadOptimizerNewRelicOptions) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *AssetWorkloadOptimizerNewRelicOptions) HasRegion() bool`

HasRegion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


