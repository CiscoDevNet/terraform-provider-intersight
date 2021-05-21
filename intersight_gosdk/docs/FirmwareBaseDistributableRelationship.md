# FirmwareBaseDistributableRelationship

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "mo.MoRef"]
**ObjectType** | **string** | The fully-qualified name of the remote type referred by this relationship. | 
**Moid** | Pointer to **string** | The Moid of the referenced REST resource. | [optional] 
**Selector** | Pointer to **string** | An OData $filter expression which describes the REST resource to be referenced. This field may be set instead of &#39;moid&#39; by clients. 1. If &#39;moid&#39; is set this field is ignored. 1. If &#39;selector&#39; is set and &#39;moid&#39; is empty/absent from the request, Intersight determines the Moid of the resource matching the filter expression and populates it in the MoRef that is part of the object instance being inserted/updated to fulfill the REST request. An error is returned if the filter matches zero or more than one REST resource. An example filter string is: Serial eq &#39;3AA8B7T11&#39;. | [optional] 
**Link** | Pointer to **string** | A URL to an instance of the &#39;mo.MoRef&#39; class. | [optional] 

## Methods

### NewFirmwareBaseDistributableRelationship

`func NewFirmwareBaseDistributableRelationship(classId string, objectType string, ) *FirmwareBaseDistributableRelationship`

NewFirmwareBaseDistributableRelationship instantiates a new FirmwareBaseDistributableRelationship object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirmwareBaseDistributableRelationshipWithDefaults

`func NewFirmwareBaseDistributableRelationshipWithDefaults() *FirmwareBaseDistributableRelationship`

NewFirmwareBaseDistributableRelationshipWithDefaults instantiates a new FirmwareBaseDistributableRelationship object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *FirmwareBaseDistributableRelationship) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *FirmwareBaseDistributableRelationship) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *FirmwareBaseDistributableRelationship) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *FirmwareBaseDistributableRelationship) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *FirmwareBaseDistributableRelationship) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *FirmwareBaseDistributableRelationship) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetMoid

`func (o *FirmwareBaseDistributableRelationship) GetMoid() string`

GetMoid returns the Moid field if non-nil, zero value otherwise.

### GetMoidOk

`func (o *FirmwareBaseDistributableRelationship) GetMoidOk() (*string, bool)`

GetMoidOk returns a tuple with the Moid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoid

`func (o *FirmwareBaseDistributableRelationship) SetMoid(v string)`

SetMoid sets Moid field to given value.

### HasMoid

`func (o *FirmwareBaseDistributableRelationship) HasMoid() bool`

HasMoid returns a boolean if a field has been set.

### GetSelector

`func (o *FirmwareBaseDistributableRelationship) GetSelector() string`

GetSelector returns the Selector field if non-nil, zero value otherwise.

### GetSelectorOk

`func (o *FirmwareBaseDistributableRelationship) GetSelectorOk() (*string, bool)`

GetSelectorOk returns a tuple with the Selector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelector

`func (o *FirmwareBaseDistributableRelationship) SetSelector(v string)`

SetSelector sets Selector field to given value.

### HasSelector

`func (o *FirmwareBaseDistributableRelationship) HasSelector() bool`

HasSelector returns a boolean if a field has been set.

### GetLink

`func (o *FirmwareBaseDistributableRelationship) GetLink() string`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *FirmwareBaseDistributableRelationship) GetLinkOk() (*string, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *FirmwareBaseDistributableRelationship) SetLink(v string)`

SetLink sets Link field to given value.

### HasLink

`func (o *FirmwareBaseDistributableRelationship) HasLink() bool`

HasLink returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


