# ApicApplicationEndpointGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "apic.ApplicationEndpointGroup"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "apic.ApplicationEndpointGroup"]
**Dn** | Pointer to **string** | Distinguished Name (DN) of an object in Cisco Application Policy Infrastructure Controller (APIC) GUI. | [optional] 
**Name** | Pointer to **string** | Object name in Cisco Application Policy Infrastructure Controller (APIC) GUI. | [optional] 
**Application** | Pointer to [**NullableApicApplicationRelationship**](ApicApplicationRelationship.md) |  | [optional] 

## Methods

### NewApicApplicationEndpointGroup

`func NewApicApplicationEndpointGroup(classId string, objectType string, ) *ApicApplicationEndpointGroup`

NewApicApplicationEndpointGroup instantiates a new ApicApplicationEndpointGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApicApplicationEndpointGroupWithDefaults

`func NewApicApplicationEndpointGroupWithDefaults() *ApicApplicationEndpointGroup`

NewApicApplicationEndpointGroupWithDefaults instantiates a new ApicApplicationEndpointGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ApicApplicationEndpointGroup) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ApicApplicationEndpointGroup) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ApicApplicationEndpointGroup) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ApicApplicationEndpointGroup) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ApicApplicationEndpointGroup) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ApicApplicationEndpointGroup) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDn

`func (o *ApicApplicationEndpointGroup) GetDn() string`

GetDn returns the Dn field if non-nil, zero value otherwise.

### GetDnOk

`func (o *ApicApplicationEndpointGroup) GetDnOk() (*string, bool)`

GetDnOk returns a tuple with the Dn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDn

`func (o *ApicApplicationEndpointGroup) SetDn(v string)`

SetDn sets Dn field to given value.

### HasDn

`func (o *ApicApplicationEndpointGroup) HasDn() bool`

HasDn returns a boolean if a field has been set.

### GetName

`func (o *ApicApplicationEndpointGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApicApplicationEndpointGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApicApplicationEndpointGroup) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApicApplicationEndpointGroup) HasName() bool`

HasName returns a boolean if a field has been set.

### GetApplication

`func (o *ApicApplicationEndpointGroup) GetApplication() ApicApplicationRelationship`

GetApplication returns the Application field if non-nil, zero value otherwise.

### GetApplicationOk

`func (o *ApicApplicationEndpointGroup) GetApplicationOk() (*ApicApplicationRelationship, bool)`

GetApplicationOk returns a tuple with the Application field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplication

`func (o *ApicApplicationEndpointGroup) SetApplication(v ApicApplicationRelationship)`

SetApplication sets Application field to given value.

### HasApplication

`func (o *ApicApplicationEndpointGroup) HasApplication() bool`

HasApplication returns a boolean if a field has been set.

### SetApplicationNil

`func (o *ApicApplicationEndpointGroup) SetApplicationNil(b bool)`

 SetApplicationNil sets the value for Application to be an explicit nil

### UnsetApplication
`func (o *ApicApplicationEndpointGroup) UnsetApplication()`

UnsetApplication ensures that no value is present for Application, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


