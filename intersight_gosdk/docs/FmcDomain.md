# FmcDomain

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "fmc.Domain"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "fmc.Domain"]
**Name** | Pointer to **string** | A user provided name of the FMC Domain. | [optional] [readonly] 
**Uuid** | Pointer to **string** | Unique Identifier of the Domain. | [optional] [readonly] 

## Methods

### NewFmcDomain

`func NewFmcDomain(classId string, objectType string, ) *FmcDomain`

NewFmcDomain instantiates a new FmcDomain object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFmcDomainWithDefaults

`func NewFmcDomainWithDefaults() *FmcDomain`

NewFmcDomainWithDefaults instantiates a new FmcDomain object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *FmcDomain) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *FmcDomain) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *FmcDomain) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *FmcDomain) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *FmcDomain) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *FmcDomain) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetName

`func (o *FmcDomain) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FmcDomain) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FmcDomain) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FmcDomain) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUuid

`func (o *FmcDomain) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *FmcDomain) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *FmcDomain) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *FmcDomain) HasUuid() bool`

HasUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


