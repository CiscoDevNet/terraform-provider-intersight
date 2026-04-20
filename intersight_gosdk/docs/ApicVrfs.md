# ApicVrfs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "apic.Vrfs"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "apic.Vrfs"]
**Dn** | Pointer to **string** | Distinguished name generated from URL parameters. | [optional] 
**Name** | Pointer to **string** | VRF name generated from URL parameters. | [optional] 
**Tenant** | Pointer to [**NullableApicTenantRelationship**](ApicTenantRelationship.md) |  | [optional] 

## Methods

### NewApicVrfs

`func NewApicVrfs(classId string, objectType string, ) *ApicVrfs`

NewApicVrfs instantiates a new ApicVrfs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApicVrfsWithDefaults

`func NewApicVrfsWithDefaults() *ApicVrfs`

NewApicVrfsWithDefaults instantiates a new ApicVrfs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ApicVrfs) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ApicVrfs) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ApicVrfs) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ApicVrfs) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ApicVrfs) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ApicVrfs) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDn

`func (o *ApicVrfs) GetDn() string`

GetDn returns the Dn field if non-nil, zero value otherwise.

### GetDnOk

`func (o *ApicVrfs) GetDnOk() (*string, bool)`

GetDnOk returns a tuple with the Dn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDn

`func (o *ApicVrfs) SetDn(v string)`

SetDn sets Dn field to given value.

### HasDn

`func (o *ApicVrfs) HasDn() bool`

HasDn returns a boolean if a field has been set.

### GetName

`func (o *ApicVrfs) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApicVrfs) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApicVrfs) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApicVrfs) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTenant

`func (o *ApicVrfs) GetTenant() ApicTenantRelationship`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *ApicVrfs) GetTenantOk() (*ApicTenantRelationship, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *ApicVrfs) SetTenant(v ApicTenantRelationship)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *ApicVrfs) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *ApicVrfs) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *ApicVrfs) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


