# CloudAvailabilityZoneAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "cloud.AvailabilityZone"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "cloud.AvailabilityZone"]
**Name** | Pointer to **string** | The name of the availability zone. | [optional] [readonly] 
**ZoneId** | Pointer to **string** | The ID of the availability zone. | [optional] [readonly] 

## Methods

### NewCloudAvailabilityZoneAllOf

`func NewCloudAvailabilityZoneAllOf(classId string, objectType string, ) *CloudAvailabilityZoneAllOf`

NewCloudAvailabilityZoneAllOf instantiates a new CloudAvailabilityZoneAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudAvailabilityZoneAllOfWithDefaults

`func NewCloudAvailabilityZoneAllOfWithDefaults() *CloudAvailabilityZoneAllOf`

NewCloudAvailabilityZoneAllOfWithDefaults instantiates a new CloudAvailabilityZoneAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CloudAvailabilityZoneAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CloudAvailabilityZoneAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CloudAvailabilityZoneAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CloudAvailabilityZoneAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CloudAvailabilityZoneAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CloudAvailabilityZoneAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetName

`func (o *CloudAvailabilityZoneAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudAvailabilityZoneAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudAvailabilityZoneAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CloudAvailabilityZoneAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetZoneId

`func (o *CloudAvailabilityZoneAllOf) GetZoneId() string`

GetZoneId returns the ZoneId field if non-nil, zero value otherwise.

### GetZoneIdOk

`func (o *CloudAvailabilityZoneAllOf) GetZoneIdOk() (*string, bool)`

GetZoneIdOk returns a tuple with the ZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneId

`func (o *CloudAvailabilityZoneAllOf) SetZoneId(v string)`

SetZoneId sets ZoneId field to given value.

### HasZoneId

`func (o *CloudAvailabilityZoneAllOf) HasZoneId() bool`

HasZoneId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


