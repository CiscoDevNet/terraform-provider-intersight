# AssetWorkloadOptimizerHypervOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "asset.WorkloadOptimizerHypervOptions"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "asset.WorkloadOptimizerHypervOptions"]
**DiscoverHostCluster** | Pointer to **bool** | When enabled, all HyperV hosts in the same cluster of specified HyperV host target will be discovered Each server must still be configured to allow remote management (for example, configuring WinRM using a GPO). | [optional] 
**FullDomainName** | Pointer to **string** | Fully qualified domain name of the HyperV target. It is used to get the name of the Hyper-V cluster node and do Active Directory authentication process through Kerberos scheme. | [optional] 

## Methods

### NewAssetWorkloadOptimizerHypervOptions

`func NewAssetWorkloadOptimizerHypervOptions(classId string, objectType string, ) *AssetWorkloadOptimizerHypervOptions`

NewAssetWorkloadOptimizerHypervOptions instantiates a new AssetWorkloadOptimizerHypervOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetWorkloadOptimizerHypervOptionsWithDefaults

`func NewAssetWorkloadOptimizerHypervOptionsWithDefaults() *AssetWorkloadOptimizerHypervOptions`

NewAssetWorkloadOptimizerHypervOptionsWithDefaults instantiates a new AssetWorkloadOptimizerHypervOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *AssetWorkloadOptimizerHypervOptions) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *AssetWorkloadOptimizerHypervOptions) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *AssetWorkloadOptimizerHypervOptions) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *AssetWorkloadOptimizerHypervOptions) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *AssetWorkloadOptimizerHypervOptions) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *AssetWorkloadOptimizerHypervOptions) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDiscoverHostCluster

`func (o *AssetWorkloadOptimizerHypervOptions) GetDiscoverHostCluster() bool`

GetDiscoverHostCluster returns the DiscoverHostCluster field if non-nil, zero value otherwise.

### GetDiscoverHostClusterOk

`func (o *AssetWorkloadOptimizerHypervOptions) GetDiscoverHostClusterOk() (*bool, bool)`

GetDiscoverHostClusterOk returns a tuple with the DiscoverHostCluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoverHostCluster

`func (o *AssetWorkloadOptimizerHypervOptions) SetDiscoverHostCluster(v bool)`

SetDiscoverHostCluster sets DiscoverHostCluster field to given value.

### HasDiscoverHostCluster

`func (o *AssetWorkloadOptimizerHypervOptions) HasDiscoverHostCluster() bool`

HasDiscoverHostCluster returns a boolean if a field has been set.

### GetFullDomainName

`func (o *AssetWorkloadOptimizerHypervOptions) GetFullDomainName() string`

GetFullDomainName returns the FullDomainName field if non-nil, zero value otherwise.

### GetFullDomainNameOk

`func (o *AssetWorkloadOptimizerHypervOptions) GetFullDomainNameOk() (*string, bool)`

GetFullDomainNameOk returns a tuple with the FullDomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullDomainName

`func (o *AssetWorkloadOptimizerHypervOptions) SetFullDomainName(v string)`

SetFullDomainName sets FullDomainName field to given value.

### HasFullDomainName

`func (o *AssetWorkloadOptimizerHypervOptions) HasFullDomainName() bool`

HasFullDomainName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


