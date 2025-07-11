# ApicBridgeDomain

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "apic.BridgeDomain"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "apic.BridgeDomain"]
**Dn** | Pointer to **string** | Distinguished Name (DN) of an object in Cisco Application Policy Infrastructure Controller (APIC) GUI. | [optional] 
**Name** | Pointer to **string** | Object name in Cisco Application Policy Infrastructure Controller (APIC) GUI. | [optional] 
**Tenant** | Pointer to [**NullableApicTenantRelationship**](ApicTenantRelationship.md) |  | [optional] 

## Methods

### NewApicBridgeDomain

`func NewApicBridgeDomain(classId string, objectType string, ) *ApicBridgeDomain`

NewApicBridgeDomain instantiates a new ApicBridgeDomain object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApicBridgeDomainWithDefaults

`func NewApicBridgeDomainWithDefaults() *ApicBridgeDomain`

NewApicBridgeDomainWithDefaults instantiates a new ApicBridgeDomain object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ApicBridgeDomain) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ApicBridgeDomain) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ApicBridgeDomain) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ApicBridgeDomain) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ApicBridgeDomain) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ApicBridgeDomain) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDn

`func (o *ApicBridgeDomain) GetDn() string`

GetDn returns the Dn field if non-nil, zero value otherwise.

### GetDnOk

`func (o *ApicBridgeDomain) GetDnOk() (*string, bool)`

GetDnOk returns a tuple with the Dn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDn

`func (o *ApicBridgeDomain) SetDn(v string)`

SetDn sets Dn field to given value.

### HasDn

`func (o *ApicBridgeDomain) HasDn() bool`

HasDn returns a boolean if a field has been set.

### GetName

`func (o *ApicBridgeDomain) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApicBridgeDomain) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApicBridgeDomain) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApicBridgeDomain) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTenant

`func (o *ApicBridgeDomain) GetTenant() ApicTenantRelationship`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *ApicBridgeDomain) GetTenantOk() (*ApicTenantRelationship, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *ApicBridgeDomain) SetTenant(v ApicTenantRelationship)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *ApicBridgeDomain) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *ApicBridgeDomain) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *ApicBridgeDomain) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


