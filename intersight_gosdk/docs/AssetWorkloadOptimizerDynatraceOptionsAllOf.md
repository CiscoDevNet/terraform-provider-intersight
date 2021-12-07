# AssetWorkloadOptimizerDynatraceOptionsAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "asset.WorkloadOptimizerDynatraceOptions"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "asset.WorkloadOptimizerDynatraceOptions"]
**EnvironmentId** | Pointer to **string** | Each environment monitored with Dynatrace is identified with a unique character string—the environment ID. The Dynatrace API relies heavily on environment IDs to ensure that it pulls monitoring data from and pushes relevant external events to the correct Dynatrace environments. | [optional] 

## Methods

### NewAssetWorkloadOptimizerDynatraceOptionsAllOf

`func NewAssetWorkloadOptimizerDynatraceOptionsAllOf(classId string, objectType string, ) *AssetWorkloadOptimizerDynatraceOptionsAllOf`

NewAssetWorkloadOptimizerDynatraceOptionsAllOf instantiates a new AssetWorkloadOptimizerDynatraceOptionsAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetWorkloadOptimizerDynatraceOptionsAllOfWithDefaults

`func NewAssetWorkloadOptimizerDynatraceOptionsAllOfWithDefaults() *AssetWorkloadOptimizerDynatraceOptionsAllOf`

NewAssetWorkloadOptimizerDynatraceOptionsAllOfWithDefaults instantiates a new AssetWorkloadOptimizerDynatraceOptionsAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *AssetWorkloadOptimizerDynatraceOptionsAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *AssetWorkloadOptimizerDynatraceOptionsAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *AssetWorkloadOptimizerDynatraceOptionsAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *AssetWorkloadOptimizerDynatraceOptionsAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *AssetWorkloadOptimizerDynatraceOptionsAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *AssetWorkloadOptimizerDynatraceOptionsAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetEnvironmentId

`func (o *AssetWorkloadOptimizerDynatraceOptionsAllOf) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *AssetWorkloadOptimizerDynatraceOptionsAllOf) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *AssetWorkloadOptimizerDynatraceOptionsAllOf) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.

### HasEnvironmentId

`func (o *AssetWorkloadOptimizerDynatraceOptionsAllOf) HasEnvironmentId() bool`

HasEnvironmentId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


