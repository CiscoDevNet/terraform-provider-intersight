# AcmeFoo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "acme.Foo"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "acme.Foo"]
**Name** | Pointer to **string** | The name of the Foo object. | [optional] 
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](OrganizationOrganizationRelationship.md) |  | [optional] 

## Methods

### NewAcmeFoo

`func NewAcmeFoo(classId string, objectType string, ) *AcmeFoo`

NewAcmeFoo instantiates a new AcmeFoo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAcmeFooWithDefaults

`func NewAcmeFooWithDefaults() *AcmeFoo`

NewAcmeFooWithDefaults instantiates a new AcmeFoo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *AcmeFoo) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *AcmeFoo) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *AcmeFoo) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *AcmeFoo) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *AcmeFoo) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *AcmeFoo) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetName

`func (o *AcmeFoo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AcmeFoo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AcmeFoo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AcmeFoo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrganization

`func (o *AcmeFoo) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *AcmeFoo) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *AcmeFoo) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *AcmeFoo) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


