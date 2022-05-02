# AssetOrchestrationHsmVmwareVcenterOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "asset.OrchestrationHsmVmwareVcenterOptions"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "asset.OrchestrationHsmVmwareVcenterOptions"]
**HsmEnabled** | Pointer to **bool** | HsmEnabled controls whether Hardware Support Manager is enabled or not. vCenter Server version 7.0 or later. | [optional] 
**IsClientSecretSet** | Pointer to **bool** | Indicates whether the value of the &#39;clientSecret&#39; property has been set. | [optional] [readonly] [default to false]

## Methods

### NewAssetOrchestrationHsmVmwareVcenterOptions

`func NewAssetOrchestrationHsmVmwareVcenterOptions(classId string, objectType string, ) *AssetOrchestrationHsmVmwareVcenterOptions`

NewAssetOrchestrationHsmVmwareVcenterOptions instantiates a new AssetOrchestrationHsmVmwareVcenterOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetOrchestrationHsmVmwareVcenterOptionsWithDefaults

`func NewAssetOrchestrationHsmVmwareVcenterOptionsWithDefaults() *AssetOrchestrationHsmVmwareVcenterOptions`

NewAssetOrchestrationHsmVmwareVcenterOptionsWithDefaults instantiates a new AssetOrchestrationHsmVmwareVcenterOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *AssetOrchestrationHsmVmwareVcenterOptions) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *AssetOrchestrationHsmVmwareVcenterOptions) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *AssetOrchestrationHsmVmwareVcenterOptions) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *AssetOrchestrationHsmVmwareVcenterOptions) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *AssetOrchestrationHsmVmwareVcenterOptions) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *AssetOrchestrationHsmVmwareVcenterOptions) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetHsmEnabled

`func (o *AssetOrchestrationHsmVmwareVcenterOptions) GetHsmEnabled() bool`

GetHsmEnabled returns the HsmEnabled field if non-nil, zero value otherwise.

### GetHsmEnabledOk

`func (o *AssetOrchestrationHsmVmwareVcenterOptions) GetHsmEnabledOk() (*bool, bool)`

GetHsmEnabledOk returns a tuple with the HsmEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHsmEnabled

`func (o *AssetOrchestrationHsmVmwareVcenterOptions) SetHsmEnabled(v bool)`

SetHsmEnabled sets HsmEnabled field to given value.

### HasHsmEnabled

`func (o *AssetOrchestrationHsmVmwareVcenterOptions) HasHsmEnabled() bool`

HasHsmEnabled returns a boolean if a field has been set.

### GetIsClientSecretSet

`func (o *AssetOrchestrationHsmVmwareVcenterOptions) GetIsClientSecretSet() bool`

GetIsClientSecretSet returns the IsClientSecretSet field if non-nil, zero value otherwise.

### GetIsClientSecretSetOk

`func (o *AssetOrchestrationHsmVmwareVcenterOptions) GetIsClientSecretSetOk() (*bool, bool)`

GetIsClientSecretSetOk returns a tuple with the IsClientSecretSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsClientSecretSet

`func (o *AssetOrchestrationHsmVmwareVcenterOptions) SetIsClientSecretSet(v bool)`

SetIsClientSecretSet sets IsClientSecretSet field to given value.

### HasIsClientSecretSet

`func (o *AssetOrchestrationHsmVmwareVcenterOptions) HasIsClientSecretSet() bool`

HasIsClientSecretSet returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


