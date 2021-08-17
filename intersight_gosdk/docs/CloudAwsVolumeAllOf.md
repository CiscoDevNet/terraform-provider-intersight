# CloudAwsVolumeAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "cloud.AwsVolume"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "cloud.AwsVolume"]
**AwsBillingUnit** | Pointer to [**CloudAwsBillingUnitRelationship**](cloud.AwsBillingUnit.Relationship.md) |  | [optional] 

## Methods

### NewCloudAwsVolumeAllOf

`func NewCloudAwsVolumeAllOf(classId string, objectType string, ) *CloudAwsVolumeAllOf`

NewCloudAwsVolumeAllOf instantiates a new CloudAwsVolumeAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudAwsVolumeAllOfWithDefaults

`func NewCloudAwsVolumeAllOfWithDefaults() *CloudAwsVolumeAllOf`

NewCloudAwsVolumeAllOfWithDefaults instantiates a new CloudAwsVolumeAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CloudAwsVolumeAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CloudAwsVolumeAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CloudAwsVolumeAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CloudAwsVolumeAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CloudAwsVolumeAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CloudAwsVolumeAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAwsBillingUnit

`func (o *CloudAwsVolumeAllOf) GetAwsBillingUnit() CloudAwsBillingUnitRelationship`

GetAwsBillingUnit returns the AwsBillingUnit field if non-nil, zero value otherwise.

### GetAwsBillingUnitOk

`func (o *CloudAwsVolumeAllOf) GetAwsBillingUnitOk() (*CloudAwsBillingUnitRelationship, bool)`

GetAwsBillingUnitOk returns a tuple with the AwsBillingUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsBillingUnit

`func (o *CloudAwsVolumeAllOf) SetAwsBillingUnit(v CloudAwsBillingUnitRelationship)`

SetAwsBillingUnit sets AwsBillingUnit field to given value.

### HasAwsBillingUnit

`func (o *CloudAwsVolumeAllOf) HasAwsBillingUnit() bool`

HasAwsBillingUnit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


