# AssetOrchestrationHitachiVirtualStoragePlatformOptionsAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "asset.OrchestrationHitachiVirtualStoragePlatformOptions"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "asset.OrchestrationHitachiVirtualStoragePlatformOptions"]
**OpsCenterManagementAddress** | Pointer to **string** | The DNS hostname or IP address of the Hitachi Ops Center API Configuration Manager used to manage the Hitachi Virtual Storage Platform. | [optional] 

## Methods

### NewAssetOrchestrationHitachiVirtualStoragePlatformOptionsAllOf

`func NewAssetOrchestrationHitachiVirtualStoragePlatformOptionsAllOf(classId string, objectType string, ) *AssetOrchestrationHitachiVirtualStoragePlatformOptionsAllOf`

NewAssetOrchestrationHitachiVirtualStoragePlatformOptionsAllOf instantiates a new AssetOrchestrationHitachiVirtualStoragePlatformOptionsAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetOrchestrationHitachiVirtualStoragePlatformOptionsAllOfWithDefaults

`func NewAssetOrchestrationHitachiVirtualStoragePlatformOptionsAllOfWithDefaults() *AssetOrchestrationHitachiVirtualStoragePlatformOptionsAllOf`

NewAssetOrchestrationHitachiVirtualStoragePlatformOptionsAllOfWithDefaults instantiates a new AssetOrchestrationHitachiVirtualStoragePlatformOptionsAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *AssetOrchestrationHitachiVirtualStoragePlatformOptionsAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *AssetOrchestrationHitachiVirtualStoragePlatformOptionsAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *AssetOrchestrationHitachiVirtualStoragePlatformOptionsAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *AssetOrchestrationHitachiVirtualStoragePlatformOptionsAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *AssetOrchestrationHitachiVirtualStoragePlatformOptionsAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *AssetOrchestrationHitachiVirtualStoragePlatformOptionsAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetOpsCenterManagementAddress

`func (o *AssetOrchestrationHitachiVirtualStoragePlatformOptionsAllOf) GetOpsCenterManagementAddress() string`

GetOpsCenterManagementAddress returns the OpsCenterManagementAddress field if non-nil, zero value otherwise.

### GetOpsCenterManagementAddressOk

`func (o *AssetOrchestrationHitachiVirtualStoragePlatformOptionsAllOf) GetOpsCenterManagementAddressOk() (*string, bool)`

GetOpsCenterManagementAddressOk returns a tuple with the OpsCenterManagementAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpsCenterManagementAddress

`func (o *AssetOrchestrationHitachiVirtualStoragePlatformOptionsAllOf) SetOpsCenterManagementAddress(v string)`

SetOpsCenterManagementAddress sets OpsCenterManagementAddress field to given value.

### HasOpsCenterManagementAddress

`func (o *AssetOrchestrationHitachiVirtualStoragePlatformOptionsAllOf) HasOpsCenterManagementAddress() bool`

HasOpsCenterManagementAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


