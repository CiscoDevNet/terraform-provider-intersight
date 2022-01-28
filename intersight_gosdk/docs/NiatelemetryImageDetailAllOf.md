# NiatelemetryImageDetailAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "niatelemetry.ImageDetail"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "niatelemetry.ImageDetail"]
**ImageName** | Pointer to **string** | Returns name of the image on controller. | [optional] 
**Name** | Pointer to **string** | Returns name of the image on controller. | [optional] 
**Version** | Pointer to **string** | Returns version of the image on controller. | [optional] 

## Methods

### NewNiatelemetryImageDetailAllOf

`func NewNiatelemetryImageDetailAllOf(classId string, objectType string, ) *NiatelemetryImageDetailAllOf`

NewNiatelemetryImageDetailAllOf instantiates a new NiatelemetryImageDetailAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiatelemetryImageDetailAllOfWithDefaults

`func NewNiatelemetryImageDetailAllOfWithDefaults() *NiatelemetryImageDetailAllOf`

NewNiatelemetryImageDetailAllOfWithDefaults instantiates a new NiatelemetryImageDetailAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NiatelemetryImageDetailAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NiatelemetryImageDetailAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NiatelemetryImageDetailAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NiatelemetryImageDetailAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NiatelemetryImageDetailAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NiatelemetryImageDetailAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetImageName

`func (o *NiatelemetryImageDetailAllOf) GetImageName() string`

GetImageName returns the ImageName field if non-nil, zero value otherwise.

### GetImageNameOk

`func (o *NiatelemetryImageDetailAllOf) GetImageNameOk() (*string, bool)`

GetImageNameOk returns a tuple with the ImageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageName

`func (o *NiatelemetryImageDetailAllOf) SetImageName(v string)`

SetImageName sets ImageName field to given value.

### HasImageName

`func (o *NiatelemetryImageDetailAllOf) HasImageName() bool`

HasImageName returns a boolean if a field has been set.

### GetName

`func (o *NiatelemetryImageDetailAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NiatelemetryImageDetailAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NiatelemetryImageDetailAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NiatelemetryImageDetailAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetVersion

`func (o *NiatelemetryImageDetailAllOf) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *NiatelemetryImageDetailAllOf) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *NiatelemetryImageDetailAllOf) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *NiatelemetryImageDetailAllOf) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


