# ApicApplication

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "apic.Application"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "apic.Application"]
**Dn** | Pointer to **string** | Distinguished Name (DN) of an object in Cisco Application Policy Infrastructure Controller (APIC) GUI. | [optional] 
**Name** | Pointer to **string** | Application name of an object in Cisco Application Policy Infrastructure Controller (APIC) GUI. | [optional] 
**Tenant** | Pointer to [**NullableApicTenantRelationship**](ApicTenantRelationship.md) |  | [optional] 

## Methods

### NewApicApplication

`func NewApicApplication(classId string, objectType string, ) *ApicApplication`

NewApicApplication instantiates a new ApicApplication object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApicApplicationWithDefaults

`func NewApicApplicationWithDefaults() *ApicApplication`

NewApicApplicationWithDefaults instantiates a new ApicApplication object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ApicApplication) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ApicApplication) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ApicApplication) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ApicApplication) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ApicApplication) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ApicApplication) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDn

`func (o *ApicApplication) GetDn() string`

GetDn returns the Dn field if non-nil, zero value otherwise.

### GetDnOk

`func (o *ApicApplication) GetDnOk() (*string, bool)`

GetDnOk returns a tuple with the Dn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDn

`func (o *ApicApplication) SetDn(v string)`

SetDn sets Dn field to given value.

### HasDn

`func (o *ApicApplication) HasDn() bool`

HasDn returns a boolean if a field has been set.

### GetName

`func (o *ApicApplication) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApicApplication) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApicApplication) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApicApplication) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTenant

`func (o *ApicApplication) GetTenant() ApicTenantRelationship`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *ApicApplication) GetTenantOk() (*ApicTenantRelationship, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *ApicApplication) SetTenant(v ApicTenantRelationship)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *ApicApplication) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *ApicApplication) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *ApicApplication) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


