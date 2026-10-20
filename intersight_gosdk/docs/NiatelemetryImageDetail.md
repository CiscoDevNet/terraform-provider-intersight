# NiatelemetryImageDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "niatelemetry.ImageDetail"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "niatelemetry.ImageDetail"]
**ImageName** | Pointer to **string** | Returns name of the image on controller. | [optional] 
**Name** | Pointer to **string** | Returns name of the image on controller. | [optional] 
**Version** | Pointer to **string** | Returns version of the image on controller. | [optional] 

## Methods

### NewNiatelemetryImageDetail

`func NewNiatelemetryImageDetail(classId string, objectType string, ) *NiatelemetryImageDetail`

NewNiatelemetryImageDetail instantiates a new NiatelemetryImageDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiatelemetryImageDetailWithDefaults

`func NewNiatelemetryImageDetailWithDefaults() *NiatelemetryImageDetail`

NewNiatelemetryImageDetailWithDefaults instantiates a new NiatelemetryImageDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NiatelemetryImageDetail) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NiatelemetryImageDetail) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NiatelemetryImageDetail) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NiatelemetryImageDetail) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NiatelemetryImageDetail) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NiatelemetryImageDetail) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetImageName

`func (o *NiatelemetryImageDetail) GetImageName() string`

GetImageName returns the ImageName field if non-nil, zero value otherwise.

### GetImageNameOk

`func (o *NiatelemetryImageDetail) GetImageNameOk() (*string, bool)`

GetImageNameOk returns a tuple with the ImageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageName

`func (o *NiatelemetryImageDetail) SetImageName(v string)`

SetImageName sets ImageName field to given value.

### HasImageName

`func (o *NiatelemetryImageDetail) HasImageName() bool`

HasImageName returns a boolean if a field has been set.

### GetName

`func (o *NiatelemetryImageDetail) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NiatelemetryImageDetail) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NiatelemetryImageDetail) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NiatelemetryImageDetail) HasName() bool`

HasName returns a boolean if a field has been set.

### GetVersion

`func (o *NiatelemetryImageDetail) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *NiatelemetryImageDetail) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *NiatelemetryImageDetail) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *NiatelemetryImageDetail) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


