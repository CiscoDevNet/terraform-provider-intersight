# AssetOrchestrationCivpVmwareVcenterOptionsAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "asset.OrchestrationCivpVmwareVcenterOptions"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "asset.OrchestrationCivpVmwareVcenterOptions"]
**CivpEnabled** | Pointer to **bool** | CivpEnabled controls whether VSphere Remote Plugin is enabled or not. vCenter Server version 8.0 or later. | [optional] 
**IsClientSecretSet** | Pointer to **bool** | Indicates whether the value of the &#39;clientSecret&#39; property has been set. | [optional] [readonly] [default to false]

## Methods

### NewAssetOrchestrationCivpVmwareVcenterOptionsAllOf

`func NewAssetOrchestrationCivpVmwareVcenterOptionsAllOf(classId string, objectType string, ) *AssetOrchestrationCivpVmwareVcenterOptionsAllOf`

NewAssetOrchestrationCivpVmwareVcenterOptionsAllOf instantiates a new AssetOrchestrationCivpVmwareVcenterOptionsAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetOrchestrationCivpVmwareVcenterOptionsAllOfWithDefaults

`func NewAssetOrchestrationCivpVmwareVcenterOptionsAllOfWithDefaults() *AssetOrchestrationCivpVmwareVcenterOptionsAllOf`

NewAssetOrchestrationCivpVmwareVcenterOptionsAllOfWithDefaults instantiates a new AssetOrchestrationCivpVmwareVcenterOptionsAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *AssetOrchestrationCivpVmwareVcenterOptionsAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *AssetOrchestrationCivpVmwareVcenterOptionsAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *AssetOrchestrationCivpVmwareVcenterOptionsAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *AssetOrchestrationCivpVmwareVcenterOptionsAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *AssetOrchestrationCivpVmwareVcenterOptionsAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *AssetOrchestrationCivpVmwareVcenterOptionsAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCivpEnabled

`func (o *AssetOrchestrationCivpVmwareVcenterOptionsAllOf) GetCivpEnabled() bool`

GetCivpEnabled returns the CivpEnabled field if non-nil, zero value otherwise.

### GetCivpEnabledOk

`func (o *AssetOrchestrationCivpVmwareVcenterOptionsAllOf) GetCivpEnabledOk() (*bool, bool)`

GetCivpEnabledOk returns a tuple with the CivpEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCivpEnabled

`func (o *AssetOrchestrationCivpVmwareVcenterOptionsAllOf) SetCivpEnabled(v bool)`

SetCivpEnabled sets CivpEnabled field to given value.

### HasCivpEnabled

`func (o *AssetOrchestrationCivpVmwareVcenterOptionsAllOf) HasCivpEnabled() bool`

HasCivpEnabled returns a boolean if a field has been set.

### GetIsClientSecretSet

`func (o *AssetOrchestrationCivpVmwareVcenterOptionsAllOf) GetIsClientSecretSet() bool`

GetIsClientSecretSet returns the IsClientSecretSet field if non-nil, zero value otherwise.

### GetIsClientSecretSetOk

`func (o *AssetOrchestrationCivpVmwareVcenterOptionsAllOf) GetIsClientSecretSetOk() (*bool, bool)`

GetIsClientSecretSetOk returns a tuple with the IsClientSecretSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsClientSecretSet

`func (o *AssetOrchestrationCivpVmwareVcenterOptionsAllOf) SetIsClientSecretSet(v bool)`

SetIsClientSecretSet sets IsClientSecretSet field to given value.

### HasIsClientSecretSet

`func (o *AssetOrchestrationCivpVmwareVcenterOptionsAllOf) HasIsClientSecretSet() bool`

HasIsClientSecretSet returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


