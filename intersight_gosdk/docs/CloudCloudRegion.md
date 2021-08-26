# CloudCloudRegion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "cloud.CloudRegion"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "cloud.CloudRegion"]
**Name** | Pointer to **string** | The name of geographic location where your public cloud resources are located. | [optional] [readonly] 
**RegionId** | Pointer to **string** | The ID of geographic location where your public cloud resources are located. | [optional] [readonly] 

## Methods

### NewCloudCloudRegion

`func NewCloudCloudRegion(classId string, objectType string, ) *CloudCloudRegion`

NewCloudCloudRegion instantiates a new CloudCloudRegion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudCloudRegionWithDefaults

`func NewCloudCloudRegionWithDefaults() *CloudCloudRegion`

NewCloudCloudRegionWithDefaults instantiates a new CloudCloudRegion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CloudCloudRegion) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CloudCloudRegion) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CloudCloudRegion) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CloudCloudRegion) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CloudCloudRegion) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CloudCloudRegion) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetName

`func (o *CloudCloudRegion) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudCloudRegion) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudCloudRegion) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CloudCloudRegion) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRegionId

`func (o *CloudCloudRegion) GetRegionId() string`

GetRegionId returns the RegionId field if non-nil, zero value otherwise.

### GetRegionIdOk

`func (o *CloudCloudRegion) GetRegionIdOk() (*string, bool)`

GetRegionIdOk returns a tuple with the RegionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionId

`func (o *CloudCloudRegion) SetRegionId(v string)`

SetRegionId sets RegionId field to given value.

### HasRegionId

`func (o *CloudCloudRegion) HasRegionId() bool`

HasRegionId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


