# CloudAwsBillingUnitRelationship

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "mo.MoRef"]
**ObjectType** | **string** | The fully-qualified name of the remote type referred by this relationship. | 
**Moid** | Pointer to **string** | The Moid of the referenced REST resource. | [optional] 
**Selector** | Pointer to **string** | An OData $filter expression which describes the REST resource to be referenced. This field may be set instead of &#39;moid&#39; by clients. 1. If &#39;moid&#39; is set this field is ignored. 1. If &#39;selector&#39; is set and &#39;moid&#39; is empty/absent from the request, Intersight determines the Moid of the resource matching the filter expression and populates it in the MoRef that is part of the object instance being inserted/updated to fulfill the REST request. An error is returned if the filter matches zero or more than one REST resource. An example filter string is: Serial eq &#39;3AA8B7T11&#39;. | [optional] 
**Link** | Pointer to **string** | A URL to an instance of the &#39;mo.MoRef&#39; class. | [optional] 

## Methods

### NewCloudAwsBillingUnitRelationship

`func NewCloudAwsBillingUnitRelationship(classId string, objectType string, ) *CloudAwsBillingUnitRelationship`

NewCloudAwsBillingUnitRelationship instantiates a new CloudAwsBillingUnitRelationship object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudAwsBillingUnitRelationshipWithDefaults

`func NewCloudAwsBillingUnitRelationshipWithDefaults() *CloudAwsBillingUnitRelationship`

NewCloudAwsBillingUnitRelationshipWithDefaults instantiates a new CloudAwsBillingUnitRelationship object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CloudAwsBillingUnitRelationship) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CloudAwsBillingUnitRelationship) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CloudAwsBillingUnitRelationship) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CloudAwsBillingUnitRelationship) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CloudAwsBillingUnitRelationship) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CloudAwsBillingUnitRelationship) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetMoid

`func (o *CloudAwsBillingUnitRelationship) GetMoid() string`

GetMoid returns the Moid field if non-nil, zero value otherwise.

### GetMoidOk

`func (o *CloudAwsBillingUnitRelationship) GetMoidOk() (*string, bool)`

GetMoidOk returns a tuple with the Moid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoid

`func (o *CloudAwsBillingUnitRelationship) SetMoid(v string)`

SetMoid sets Moid field to given value.

### HasMoid

`func (o *CloudAwsBillingUnitRelationship) HasMoid() bool`

HasMoid returns a boolean if a field has been set.

### GetSelector

`func (o *CloudAwsBillingUnitRelationship) GetSelector() string`

GetSelector returns the Selector field if non-nil, zero value otherwise.

### GetSelectorOk

`func (o *CloudAwsBillingUnitRelationship) GetSelectorOk() (*string, bool)`

GetSelectorOk returns a tuple with the Selector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelector

`func (o *CloudAwsBillingUnitRelationship) SetSelector(v string)`

SetSelector sets Selector field to given value.

### HasSelector

`func (o *CloudAwsBillingUnitRelationship) HasSelector() bool`

HasSelector returns a boolean if a field has been set.

### GetLink

`func (o *CloudAwsBillingUnitRelationship) GetLink() string`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *CloudAwsBillingUnitRelationship) GetLinkOk() (*string, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *CloudAwsBillingUnitRelationship) SetLink(v string)`

SetLink sets Link field to given value.

### HasLink

`func (o *CloudAwsBillingUnitRelationship) HasLink() bool`

HasLink returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


