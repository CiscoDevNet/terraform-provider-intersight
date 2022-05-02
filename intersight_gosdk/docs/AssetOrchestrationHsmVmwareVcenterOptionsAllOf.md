# AssetOrchestrationHsmVmwareVcenterOptionsAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "asset.OrchestrationHsmVmwareVcenterOptions"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "asset.OrchestrationHsmVmwareVcenterOptions"]
**HsmEnabled** | Pointer to **bool** | HsmEnabled controls whether Hardware Support Manager is enabled or not. vCenter Server version 7.0 or later. | [optional] 
**IsClientSecretSet** | Pointer to **bool** | Indicates whether the value of the &#39;clientSecret&#39; property has been set. | [optional] [readonly] [default to false]

## Methods

### NewAssetOrchestrationHsmVmwareVcenterOptionsAllOf

`func NewAssetOrchestrationHsmVmwareVcenterOptionsAllOf(classId string, objectType string, ) *AssetOrchestrationHsmVmwareVcenterOptionsAllOf`

NewAssetOrchestrationHsmVmwareVcenterOptionsAllOf instantiates a new AssetOrchestrationHsmVmwareVcenterOptionsAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetOrchestrationHsmVmwareVcenterOptionsAllOfWithDefaults

`func NewAssetOrchestrationHsmVmwareVcenterOptionsAllOfWithDefaults() *AssetOrchestrationHsmVmwareVcenterOptionsAllOf`

NewAssetOrchestrationHsmVmwareVcenterOptionsAllOfWithDefaults instantiates a new AssetOrchestrationHsmVmwareVcenterOptionsAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *AssetOrchestrationHsmVmwareVcenterOptionsAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *AssetOrchestrationHsmVmwareVcenterOptionsAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *AssetOrchestrationHsmVmwareVcenterOptionsAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *AssetOrchestrationHsmVmwareVcenterOptionsAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *AssetOrchestrationHsmVmwareVcenterOptionsAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *AssetOrchestrationHsmVmwareVcenterOptionsAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetHsmEnabled

`func (o *AssetOrchestrationHsmVmwareVcenterOptionsAllOf) GetHsmEnabled() bool`

GetHsmEnabled returns the HsmEnabled field if non-nil, zero value otherwise.

### GetHsmEnabledOk

`func (o *AssetOrchestrationHsmVmwareVcenterOptionsAllOf) GetHsmEnabledOk() (*bool, bool)`

GetHsmEnabledOk returns a tuple with the HsmEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHsmEnabled

`func (o *AssetOrchestrationHsmVmwareVcenterOptionsAllOf) SetHsmEnabled(v bool)`

SetHsmEnabled sets HsmEnabled field to given value.

### HasHsmEnabled

`func (o *AssetOrchestrationHsmVmwareVcenterOptionsAllOf) HasHsmEnabled() bool`

HasHsmEnabled returns a boolean if a field has been set.

### GetIsClientSecretSet

`func (o *AssetOrchestrationHsmVmwareVcenterOptionsAllOf) GetIsClientSecretSet() bool`

GetIsClientSecretSet returns the IsClientSecretSet field if non-nil, zero value otherwise.

### GetIsClientSecretSetOk

`func (o *AssetOrchestrationHsmVmwareVcenterOptionsAllOf) GetIsClientSecretSetOk() (*bool, bool)`

GetIsClientSecretSetOk returns a tuple with the IsClientSecretSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsClientSecretSet

`func (o *AssetOrchestrationHsmVmwareVcenterOptionsAllOf) SetIsClientSecretSet(v bool)`

SetIsClientSecretSet sets IsClientSecretSet field to given value.

### HasIsClientSecretSet

`func (o *AssetOrchestrationHsmVmwareVcenterOptionsAllOf) HasIsClientSecretSet() bool`

HasIsClientSecretSet returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


