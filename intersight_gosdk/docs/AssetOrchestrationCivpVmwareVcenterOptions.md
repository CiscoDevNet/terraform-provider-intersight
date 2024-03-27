# AssetOrchestrationCivpVmwareVcenterOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "asset.OrchestrationCivpVmwareVcenterOptions"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "asset.OrchestrationCivpVmwareVcenterOptions"]
**CivpEnabled** | Pointer to **bool** | CivpEnabled controls whether VSphere Remote Plugin is enabled or not. vCenter Server version 8.0 or later. | [optional] 
**IsClientSecretSet** | Pointer to **bool** | Indicates whether the value of the &#39;clientSecret&#39; property has been set. | [optional] [readonly] [default to false]

## Methods

### NewAssetOrchestrationCivpVmwareVcenterOptions

`func NewAssetOrchestrationCivpVmwareVcenterOptions(classId string, objectType string, ) *AssetOrchestrationCivpVmwareVcenterOptions`

NewAssetOrchestrationCivpVmwareVcenterOptions instantiates a new AssetOrchestrationCivpVmwareVcenterOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetOrchestrationCivpVmwareVcenterOptionsWithDefaults

`func NewAssetOrchestrationCivpVmwareVcenterOptionsWithDefaults() *AssetOrchestrationCivpVmwareVcenterOptions`

NewAssetOrchestrationCivpVmwareVcenterOptionsWithDefaults instantiates a new AssetOrchestrationCivpVmwareVcenterOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *AssetOrchestrationCivpVmwareVcenterOptions) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *AssetOrchestrationCivpVmwareVcenterOptions) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *AssetOrchestrationCivpVmwareVcenterOptions) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *AssetOrchestrationCivpVmwareVcenterOptions) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *AssetOrchestrationCivpVmwareVcenterOptions) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *AssetOrchestrationCivpVmwareVcenterOptions) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCivpEnabled

`func (o *AssetOrchestrationCivpVmwareVcenterOptions) GetCivpEnabled() bool`

GetCivpEnabled returns the CivpEnabled field if non-nil, zero value otherwise.

### GetCivpEnabledOk

`func (o *AssetOrchestrationCivpVmwareVcenterOptions) GetCivpEnabledOk() (*bool, bool)`

GetCivpEnabledOk returns a tuple with the CivpEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCivpEnabled

`func (o *AssetOrchestrationCivpVmwareVcenterOptions) SetCivpEnabled(v bool)`

SetCivpEnabled sets CivpEnabled field to given value.

### HasCivpEnabled

`func (o *AssetOrchestrationCivpVmwareVcenterOptions) HasCivpEnabled() bool`

HasCivpEnabled returns a boolean if a field has been set.

### GetIsClientSecretSet

`func (o *AssetOrchestrationCivpVmwareVcenterOptions) GetIsClientSecretSet() bool`

GetIsClientSecretSet returns the IsClientSecretSet field if non-nil, zero value otherwise.

### GetIsClientSecretSetOk

`func (o *AssetOrchestrationCivpVmwareVcenterOptions) GetIsClientSecretSetOk() (*bool, bool)`

GetIsClientSecretSetOk returns a tuple with the IsClientSecretSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsClientSecretSet

`func (o *AssetOrchestrationCivpVmwareVcenterOptions) SetIsClientSecretSet(v bool)`

SetIsClientSecretSet sets IsClientSecretSet field to given value.

### HasIsClientSecretSet

`func (o *AssetOrchestrationCivpVmwareVcenterOptions) HasIsClientSecretSet() bool`

HasIsClientSecretSet returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


