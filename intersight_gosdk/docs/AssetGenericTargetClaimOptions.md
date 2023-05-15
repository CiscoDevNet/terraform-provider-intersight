# AssetGenericTargetClaimOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "asset.GenericTargetClaimOptions"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "asset.GenericTargetClaimOptions"]
**TargetOptions** | Pointer to **interface{}** | The generic device connector will consume these options. | [optional] 

## Methods

### NewAssetGenericTargetClaimOptions

`func NewAssetGenericTargetClaimOptions(classId string, objectType string, ) *AssetGenericTargetClaimOptions`

NewAssetGenericTargetClaimOptions instantiates a new AssetGenericTargetClaimOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetGenericTargetClaimOptionsWithDefaults

`func NewAssetGenericTargetClaimOptionsWithDefaults() *AssetGenericTargetClaimOptions`

NewAssetGenericTargetClaimOptionsWithDefaults instantiates a new AssetGenericTargetClaimOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *AssetGenericTargetClaimOptions) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *AssetGenericTargetClaimOptions) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *AssetGenericTargetClaimOptions) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *AssetGenericTargetClaimOptions) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *AssetGenericTargetClaimOptions) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *AssetGenericTargetClaimOptions) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetTargetOptions

`func (o *AssetGenericTargetClaimOptions) GetTargetOptions() interface{}`

GetTargetOptions returns the TargetOptions field if non-nil, zero value otherwise.

### GetTargetOptionsOk

`func (o *AssetGenericTargetClaimOptions) GetTargetOptionsOk() (*interface{}, bool)`

GetTargetOptionsOk returns a tuple with the TargetOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetOptions

`func (o *AssetGenericTargetClaimOptions) SetTargetOptions(v interface{})`

SetTargetOptions sets TargetOptions field to given value.

### HasTargetOptions

`func (o *AssetGenericTargetClaimOptions) HasTargetOptions() bool`

HasTargetOptions returns a boolean if a field has been set.

### SetTargetOptionsNil

`func (o *AssetGenericTargetClaimOptions) SetTargetOptionsNil(b bool)`

 SetTargetOptionsNil sets the value for TargetOptions to be an explicit nil

### UnsetTargetOptions
`func (o *AssetGenericTargetClaimOptions) UnsetTargetOptions()`

UnsetTargetOptions ensures that no value is present for TargetOptions, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


