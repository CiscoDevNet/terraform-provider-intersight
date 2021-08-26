# CloudImageReferenceAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "cloud.ImageReference"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "cloud.ImageReference"]
**ImageId** | Pointer to **string** | Identity of the image used in deployment of virtual machine. | [optional] [readonly] 
**Name** | Pointer to **string** | Name of the image used in deployment of virtual machine. | [optional] [readonly] 

## Methods

### NewCloudImageReferenceAllOf

`func NewCloudImageReferenceAllOf(classId string, objectType string, ) *CloudImageReferenceAllOf`

NewCloudImageReferenceAllOf instantiates a new CloudImageReferenceAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudImageReferenceAllOfWithDefaults

`func NewCloudImageReferenceAllOfWithDefaults() *CloudImageReferenceAllOf`

NewCloudImageReferenceAllOfWithDefaults instantiates a new CloudImageReferenceAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CloudImageReferenceAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CloudImageReferenceAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CloudImageReferenceAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CloudImageReferenceAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CloudImageReferenceAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CloudImageReferenceAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetImageId

`func (o *CloudImageReferenceAllOf) GetImageId() string`

GetImageId returns the ImageId field if non-nil, zero value otherwise.

### GetImageIdOk

`func (o *CloudImageReferenceAllOf) GetImageIdOk() (*string, bool)`

GetImageIdOk returns a tuple with the ImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageId

`func (o *CloudImageReferenceAllOf) SetImageId(v string)`

SetImageId sets ImageId field to given value.

### HasImageId

`func (o *CloudImageReferenceAllOf) HasImageId() bool`

HasImageId returns a boolean if a field has been set.

### GetName

`func (o *CloudImageReferenceAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudImageReferenceAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudImageReferenceAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CloudImageReferenceAllOf) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


