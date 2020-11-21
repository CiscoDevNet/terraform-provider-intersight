# AssetOrchestrationHitachiVirtualStoragePlatformOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "asset.OrchestrationHitachiVirtualStoragePlatformOptions"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "asset.OrchestrationHitachiVirtualStoragePlatformOptions"]
**OpsCenterManagementAddress** | Pointer to **string** | The DNS hostname or IP address of the Hitachi Ops Center API Configuration Manager used to manage the Hitachi Virtual Storage Platform. | [optional] 

## Methods

### NewAssetOrchestrationHitachiVirtualStoragePlatformOptions

`func NewAssetOrchestrationHitachiVirtualStoragePlatformOptions(classId string, objectType string, ) *AssetOrchestrationHitachiVirtualStoragePlatformOptions`

NewAssetOrchestrationHitachiVirtualStoragePlatformOptions instantiates a new AssetOrchestrationHitachiVirtualStoragePlatformOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetOrchestrationHitachiVirtualStoragePlatformOptionsWithDefaults

`func NewAssetOrchestrationHitachiVirtualStoragePlatformOptionsWithDefaults() *AssetOrchestrationHitachiVirtualStoragePlatformOptions`

NewAssetOrchestrationHitachiVirtualStoragePlatformOptionsWithDefaults instantiates a new AssetOrchestrationHitachiVirtualStoragePlatformOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *AssetOrchestrationHitachiVirtualStoragePlatformOptions) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *AssetOrchestrationHitachiVirtualStoragePlatformOptions) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *AssetOrchestrationHitachiVirtualStoragePlatformOptions) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *AssetOrchestrationHitachiVirtualStoragePlatformOptions) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *AssetOrchestrationHitachiVirtualStoragePlatformOptions) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *AssetOrchestrationHitachiVirtualStoragePlatformOptions) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetOpsCenterManagementAddress

`func (o *AssetOrchestrationHitachiVirtualStoragePlatformOptions) GetOpsCenterManagementAddress() string`

GetOpsCenterManagementAddress returns the OpsCenterManagementAddress field if non-nil, zero value otherwise.

### GetOpsCenterManagementAddressOk

`func (o *AssetOrchestrationHitachiVirtualStoragePlatformOptions) GetOpsCenterManagementAddressOk() (*string, bool)`

GetOpsCenterManagementAddressOk returns a tuple with the OpsCenterManagementAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpsCenterManagementAddress

`func (o *AssetOrchestrationHitachiVirtualStoragePlatformOptions) SetOpsCenterManagementAddress(v string)`

SetOpsCenterManagementAddress sets OpsCenterManagementAddress field to given value.

### HasOpsCenterManagementAddress

`func (o *AssetOrchestrationHitachiVirtualStoragePlatformOptions) HasOpsCenterManagementAddress() bool`

HasOpsCenterManagementAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


