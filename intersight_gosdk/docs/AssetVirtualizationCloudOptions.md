# AssetVirtualizationCloudOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | [default to "asset.VirtualizationAmazonWebServiceOptions"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | [default to "asset.VirtualizationAmazonWebServiceOptions"]
**ManagedRegions** | Pointer to [**[]MoMoRef**](MoMoRef.md) |  | [optional] 
**RegionGroup** | Pointer to **string** | The name of the region group to which the managedRegions belong. Populated from the group property in cloud.Regions. | [optional] [readonly] 

## Methods

### NewAssetVirtualizationCloudOptions

`func NewAssetVirtualizationCloudOptions(classId string, objectType string, ) *AssetVirtualizationCloudOptions`

NewAssetVirtualizationCloudOptions instantiates a new AssetVirtualizationCloudOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetVirtualizationCloudOptionsWithDefaults

`func NewAssetVirtualizationCloudOptionsWithDefaults() *AssetVirtualizationCloudOptions`

NewAssetVirtualizationCloudOptionsWithDefaults instantiates a new AssetVirtualizationCloudOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *AssetVirtualizationCloudOptions) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *AssetVirtualizationCloudOptions) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *AssetVirtualizationCloudOptions) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *AssetVirtualizationCloudOptions) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *AssetVirtualizationCloudOptions) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *AssetVirtualizationCloudOptions) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetManagedRegions

`func (o *AssetVirtualizationCloudOptions) GetManagedRegions() []MoMoRef`

GetManagedRegions returns the ManagedRegions field if non-nil, zero value otherwise.

### GetManagedRegionsOk

`func (o *AssetVirtualizationCloudOptions) GetManagedRegionsOk() (*[]MoMoRef, bool)`

GetManagedRegionsOk returns a tuple with the ManagedRegions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedRegions

`func (o *AssetVirtualizationCloudOptions) SetManagedRegions(v []MoMoRef)`

SetManagedRegions sets ManagedRegions field to given value.

### HasManagedRegions

`func (o *AssetVirtualizationCloudOptions) HasManagedRegions() bool`

HasManagedRegions returns a boolean if a field has been set.

### SetManagedRegionsNil

`func (o *AssetVirtualizationCloudOptions) SetManagedRegionsNil(b bool)`

 SetManagedRegionsNil sets the value for ManagedRegions to be an explicit nil

### UnsetManagedRegions
`func (o *AssetVirtualizationCloudOptions) UnsetManagedRegions()`

UnsetManagedRegions ensures that no value is present for ManagedRegions, not even an explicit nil
### GetRegionGroup

`func (o *AssetVirtualizationCloudOptions) GetRegionGroup() string`

GetRegionGroup returns the RegionGroup field if non-nil, zero value otherwise.

### GetRegionGroupOk

`func (o *AssetVirtualizationCloudOptions) GetRegionGroupOk() (*string, bool)`

GetRegionGroupOk returns a tuple with the RegionGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionGroup

`func (o *AssetVirtualizationCloudOptions) SetRegionGroup(v string)`

SetRegionGroup sets RegionGroup field to given value.

### HasRegionGroup

`func (o *AssetVirtualizationCloudOptions) HasRegionGroup() bool`

HasRegionGroup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


