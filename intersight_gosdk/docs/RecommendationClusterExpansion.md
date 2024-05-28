# RecommendationClusterExpansion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "recommendation.ClusterExpansion"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "recommendation.ClusterExpansion"]
**ClusterName** | Pointer to **string** | Name of the cluster for which the expansion recommendation is provided. | [optional] [readonly] 
**HardwareExpansionRequest** | Pointer to [**NullableRecommendationHardwareExpansionRequestRelationship**](RecommendationHardwareExpansionRequestRelationship.md) |  | [optional] 
**PhysicalItem** | Pointer to [**[]RecommendationPhysicalItemRelationship**](RecommendationPhysicalItemRelationship.md) | An array of relationships to recommendationPhysicalItem resources. | [optional] [readonly] 
**RegisteredDevice** | Pointer to [**NullableAssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 
**SoftwareItem** | Pointer to [**[]RecommendationSoftwareItemRelationship**](RecommendationSoftwareItemRelationship.md) | An array of relationships to recommendationSoftwareItem resources. | [optional] [readonly] 

## Methods

### NewRecommendationClusterExpansion

`func NewRecommendationClusterExpansion(classId string, objectType string, ) *RecommendationClusterExpansion`

NewRecommendationClusterExpansion instantiates a new RecommendationClusterExpansion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecommendationClusterExpansionWithDefaults

`func NewRecommendationClusterExpansionWithDefaults() *RecommendationClusterExpansion`

NewRecommendationClusterExpansionWithDefaults instantiates a new RecommendationClusterExpansion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *RecommendationClusterExpansion) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *RecommendationClusterExpansion) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *RecommendationClusterExpansion) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *RecommendationClusterExpansion) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *RecommendationClusterExpansion) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *RecommendationClusterExpansion) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetClusterName

`func (o *RecommendationClusterExpansion) GetClusterName() string`

GetClusterName returns the ClusterName field if non-nil, zero value otherwise.

### GetClusterNameOk

`func (o *RecommendationClusterExpansion) GetClusterNameOk() (*string, bool)`

GetClusterNameOk returns a tuple with the ClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterName

`func (o *RecommendationClusterExpansion) SetClusterName(v string)`

SetClusterName sets ClusterName field to given value.

### HasClusterName

`func (o *RecommendationClusterExpansion) HasClusterName() bool`

HasClusterName returns a boolean if a field has been set.

### GetHardwareExpansionRequest

`func (o *RecommendationClusterExpansion) GetHardwareExpansionRequest() RecommendationHardwareExpansionRequestRelationship`

GetHardwareExpansionRequest returns the HardwareExpansionRequest field if non-nil, zero value otherwise.

### GetHardwareExpansionRequestOk

`func (o *RecommendationClusterExpansion) GetHardwareExpansionRequestOk() (*RecommendationHardwareExpansionRequestRelationship, bool)`

GetHardwareExpansionRequestOk returns a tuple with the HardwareExpansionRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardwareExpansionRequest

`func (o *RecommendationClusterExpansion) SetHardwareExpansionRequest(v RecommendationHardwareExpansionRequestRelationship)`

SetHardwareExpansionRequest sets HardwareExpansionRequest field to given value.

### HasHardwareExpansionRequest

`func (o *RecommendationClusterExpansion) HasHardwareExpansionRequest() bool`

HasHardwareExpansionRequest returns a boolean if a field has been set.

### SetHardwareExpansionRequestNil

`func (o *RecommendationClusterExpansion) SetHardwareExpansionRequestNil(b bool)`

 SetHardwareExpansionRequestNil sets the value for HardwareExpansionRequest to be an explicit nil

### UnsetHardwareExpansionRequest
`func (o *RecommendationClusterExpansion) UnsetHardwareExpansionRequest()`

UnsetHardwareExpansionRequest ensures that no value is present for HardwareExpansionRequest, not even an explicit nil
### GetPhysicalItem

`func (o *RecommendationClusterExpansion) GetPhysicalItem() []RecommendationPhysicalItemRelationship`

GetPhysicalItem returns the PhysicalItem field if non-nil, zero value otherwise.

### GetPhysicalItemOk

`func (o *RecommendationClusterExpansion) GetPhysicalItemOk() (*[]RecommendationPhysicalItemRelationship, bool)`

GetPhysicalItemOk returns a tuple with the PhysicalItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalItem

`func (o *RecommendationClusterExpansion) SetPhysicalItem(v []RecommendationPhysicalItemRelationship)`

SetPhysicalItem sets PhysicalItem field to given value.

### HasPhysicalItem

`func (o *RecommendationClusterExpansion) HasPhysicalItem() bool`

HasPhysicalItem returns a boolean if a field has been set.

### SetPhysicalItemNil

`func (o *RecommendationClusterExpansion) SetPhysicalItemNil(b bool)`

 SetPhysicalItemNil sets the value for PhysicalItem to be an explicit nil

### UnsetPhysicalItem
`func (o *RecommendationClusterExpansion) UnsetPhysicalItem()`

UnsetPhysicalItem ensures that no value is present for PhysicalItem, not even an explicit nil
### GetRegisteredDevice

`func (o *RecommendationClusterExpansion) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *RecommendationClusterExpansion) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *RecommendationClusterExpansion) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *RecommendationClusterExpansion) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### SetRegisteredDeviceNil

`func (o *RecommendationClusterExpansion) SetRegisteredDeviceNil(b bool)`

 SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil

### UnsetRegisteredDevice
`func (o *RecommendationClusterExpansion) UnsetRegisteredDevice()`

UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil
### GetSoftwareItem

`func (o *RecommendationClusterExpansion) GetSoftwareItem() []RecommendationSoftwareItemRelationship`

GetSoftwareItem returns the SoftwareItem field if non-nil, zero value otherwise.

### GetSoftwareItemOk

`func (o *RecommendationClusterExpansion) GetSoftwareItemOk() (*[]RecommendationSoftwareItemRelationship, bool)`

GetSoftwareItemOk returns a tuple with the SoftwareItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareItem

`func (o *RecommendationClusterExpansion) SetSoftwareItem(v []RecommendationSoftwareItemRelationship)`

SetSoftwareItem sets SoftwareItem field to given value.

### HasSoftwareItem

`func (o *RecommendationClusterExpansion) HasSoftwareItem() bool`

HasSoftwareItem returns a boolean if a field has been set.

### SetSoftwareItemNil

`func (o *RecommendationClusterExpansion) SetSoftwareItemNil(b bool)`

 SetSoftwareItemNil sets the value for SoftwareItem to be an explicit nil

### UnsetSoftwareItem
`func (o *RecommendationClusterExpansion) UnsetSoftwareItem()`

UnsetSoftwareItem ensures that no value is present for SoftwareItem, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


