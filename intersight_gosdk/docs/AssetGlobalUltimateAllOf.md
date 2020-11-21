# AssetGlobalUltimateAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "asset.GlobalUltimate"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "asset.GlobalUltimate"]
**Id** | Pointer to **string** | ID of the user in BillToGlobal. | [optional] [readonly] 
**Name** | Pointer to **string** | Name of the user in BillToGlobal. | [optional] [readonly] 

## Methods

### NewAssetGlobalUltimateAllOf

`func NewAssetGlobalUltimateAllOf(classId string, objectType string, ) *AssetGlobalUltimateAllOf`

NewAssetGlobalUltimateAllOf instantiates a new AssetGlobalUltimateAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetGlobalUltimateAllOfWithDefaults

`func NewAssetGlobalUltimateAllOfWithDefaults() *AssetGlobalUltimateAllOf`

NewAssetGlobalUltimateAllOfWithDefaults instantiates a new AssetGlobalUltimateAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *AssetGlobalUltimateAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *AssetGlobalUltimateAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *AssetGlobalUltimateAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *AssetGlobalUltimateAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *AssetGlobalUltimateAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *AssetGlobalUltimateAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetId

`func (o *AssetGlobalUltimateAllOf) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AssetGlobalUltimateAllOf) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AssetGlobalUltimateAllOf) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AssetGlobalUltimateAllOf) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *AssetGlobalUltimateAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AssetGlobalUltimateAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AssetGlobalUltimateAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AssetGlobalUltimateAllOf) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


