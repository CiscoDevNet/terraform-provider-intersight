# CloudAvailabilityZone

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "cloud.AvailabilityZone"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "cloud.AvailabilityZone"]
**Name** | Pointer to **string** | The name of the availability zone. | [optional] [readonly] 
**ZoneId** | Pointer to **string** | The ID of the availability zone. | [optional] [readonly] 

## Methods

### NewCloudAvailabilityZone

`func NewCloudAvailabilityZone(classId string, objectType string, ) *CloudAvailabilityZone`

NewCloudAvailabilityZone instantiates a new CloudAvailabilityZone object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudAvailabilityZoneWithDefaults

`func NewCloudAvailabilityZoneWithDefaults() *CloudAvailabilityZone`

NewCloudAvailabilityZoneWithDefaults instantiates a new CloudAvailabilityZone object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CloudAvailabilityZone) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CloudAvailabilityZone) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CloudAvailabilityZone) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CloudAvailabilityZone) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CloudAvailabilityZone) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CloudAvailabilityZone) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetName

`func (o *CloudAvailabilityZone) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudAvailabilityZone) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudAvailabilityZone) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CloudAvailabilityZone) HasName() bool`

HasName returns a boolean if a field has been set.

### GetZoneId

`func (o *CloudAvailabilityZone) GetZoneId() string`

GetZoneId returns the ZoneId field if non-nil, zero value otherwise.

### GetZoneIdOk

`func (o *CloudAvailabilityZone) GetZoneIdOk() (*string, bool)`

GetZoneIdOk returns a tuple with the ZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneId

`func (o *CloudAvailabilityZone) SetZoneId(v string)`

SetZoneId sets ZoneId field to given value.

### HasZoneId

`func (o *CloudAvailabilityZone) HasZoneId() bool`

HasZoneId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


