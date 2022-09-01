# NiatelemetryCloudRoutersElementAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "niatelemetry.CloudRoutersElement"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "niatelemetry.CloudRoutersElement"]
**Name** | Pointer to **string** | Return value of name of the cloud router. | [optional] 
**Version** | Pointer to **string** | Return value of cloud router version. | [optional] 

## Methods

### NewNiatelemetryCloudRoutersElementAllOf

`func NewNiatelemetryCloudRoutersElementAllOf(classId string, objectType string, ) *NiatelemetryCloudRoutersElementAllOf`

NewNiatelemetryCloudRoutersElementAllOf instantiates a new NiatelemetryCloudRoutersElementAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiatelemetryCloudRoutersElementAllOfWithDefaults

`func NewNiatelemetryCloudRoutersElementAllOfWithDefaults() *NiatelemetryCloudRoutersElementAllOf`

NewNiatelemetryCloudRoutersElementAllOfWithDefaults instantiates a new NiatelemetryCloudRoutersElementAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NiatelemetryCloudRoutersElementAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NiatelemetryCloudRoutersElementAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NiatelemetryCloudRoutersElementAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NiatelemetryCloudRoutersElementAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NiatelemetryCloudRoutersElementAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NiatelemetryCloudRoutersElementAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetName

`func (o *NiatelemetryCloudRoutersElementAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NiatelemetryCloudRoutersElementAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NiatelemetryCloudRoutersElementAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NiatelemetryCloudRoutersElementAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetVersion

`func (o *NiatelemetryCloudRoutersElementAllOf) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *NiatelemetryCloudRoutersElementAllOf) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *NiatelemetryCloudRoutersElementAllOf) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *NiatelemetryCloudRoutersElementAllOf) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


