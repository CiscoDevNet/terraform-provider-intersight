# CloudCloudRegionAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "cloud.CloudRegion"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "cloud.CloudRegion"]
**Name** | Pointer to **string** | The name of geographic location where your public cloud resources are located. | [optional] [readonly] 
**RegionId** | Pointer to **string** | The ID of geographic location where your public cloud resources are located. | [optional] [readonly] 

## Methods

### NewCloudCloudRegionAllOf

`func NewCloudCloudRegionAllOf(classId string, objectType string, ) *CloudCloudRegionAllOf`

NewCloudCloudRegionAllOf instantiates a new CloudCloudRegionAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudCloudRegionAllOfWithDefaults

`func NewCloudCloudRegionAllOfWithDefaults() *CloudCloudRegionAllOf`

NewCloudCloudRegionAllOfWithDefaults instantiates a new CloudCloudRegionAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CloudCloudRegionAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CloudCloudRegionAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CloudCloudRegionAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CloudCloudRegionAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CloudCloudRegionAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CloudCloudRegionAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetName

`func (o *CloudCloudRegionAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudCloudRegionAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudCloudRegionAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CloudCloudRegionAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRegionId

`func (o *CloudCloudRegionAllOf) GetRegionId() string`

GetRegionId returns the RegionId field if non-nil, zero value otherwise.

### GetRegionIdOk

`func (o *CloudCloudRegionAllOf) GetRegionIdOk() (*string, bool)`

GetRegionIdOk returns a tuple with the RegionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionId

`func (o *CloudCloudRegionAllOf) SetRegionId(v string)`

SetRegionId sets RegionId field to given value.

### HasRegionId

`func (o *CloudCloudRegionAllOf) HasRegionId() bool`

HasRegionId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


