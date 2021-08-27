# IppoolBlockLease

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "ippool.BlockLease"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "ippool.BlockLease"]
**IpType** | Pointer to **string** | Type of the IP address requested. * &#x60;IPv4&#x60; - IP V4 address type requested. * &#x60;IPv6&#x60; - IP V6 address type requested. | [optional] [default to "IPv4"]
**AssignedToEntity** | Pointer to [**MoBaseMoRelationship**](MoBaseMoRelationship.md) |  | [optional] 
**IpLeases** | Pointer to [**[]IppoolIpLeaseRelationship**](IppoolIpLeaseRelationship.md) | An array of relationships to ippoolIpLease resources. | [optional] 
**Pool** | Pointer to [**IppoolPoolRelationship**](IppoolPoolRelationship.md) |  | [optional] 
**Universe** | Pointer to [**IppoolUniverseRelationship**](IppoolUniverseRelationship.md) |  | [optional] 
**Vrf** | Pointer to [**VrfVrfRelationship**](VrfVrfRelationship.md) |  | [optional] 

## Methods

### NewIppoolBlockLease

`func NewIppoolBlockLease(classId string, objectType string, ) *IppoolBlockLease`

NewIppoolBlockLease instantiates a new IppoolBlockLease object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIppoolBlockLeaseWithDefaults

`func NewIppoolBlockLeaseWithDefaults() *IppoolBlockLease`

NewIppoolBlockLeaseWithDefaults instantiates a new IppoolBlockLease object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *IppoolBlockLease) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *IppoolBlockLease) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *IppoolBlockLease) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *IppoolBlockLease) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *IppoolBlockLease) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *IppoolBlockLease) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetIpType

`func (o *IppoolBlockLease) GetIpType() string`

GetIpType returns the IpType field if non-nil, zero value otherwise.

### GetIpTypeOk

`func (o *IppoolBlockLease) GetIpTypeOk() (*string, bool)`

GetIpTypeOk returns a tuple with the IpType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpType

`func (o *IppoolBlockLease) SetIpType(v string)`

SetIpType sets IpType field to given value.

### HasIpType

`func (o *IppoolBlockLease) HasIpType() bool`

HasIpType returns a boolean if a field has been set.

### GetAssignedToEntity

`func (o *IppoolBlockLease) GetAssignedToEntity() MoBaseMoRelationship`

GetAssignedToEntity returns the AssignedToEntity field if non-nil, zero value otherwise.

### GetAssignedToEntityOk

`func (o *IppoolBlockLease) GetAssignedToEntityOk() (*MoBaseMoRelationship, bool)`

GetAssignedToEntityOk returns a tuple with the AssignedToEntity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedToEntity

`func (o *IppoolBlockLease) SetAssignedToEntity(v MoBaseMoRelationship)`

SetAssignedToEntity sets AssignedToEntity field to given value.

### HasAssignedToEntity

`func (o *IppoolBlockLease) HasAssignedToEntity() bool`

HasAssignedToEntity returns a boolean if a field has been set.

### GetIpLeases

`func (o *IppoolBlockLease) GetIpLeases() []IppoolIpLeaseRelationship`

GetIpLeases returns the IpLeases field if non-nil, zero value otherwise.

### GetIpLeasesOk

`func (o *IppoolBlockLease) GetIpLeasesOk() (*[]IppoolIpLeaseRelationship, bool)`

GetIpLeasesOk returns a tuple with the IpLeases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpLeases

`func (o *IppoolBlockLease) SetIpLeases(v []IppoolIpLeaseRelationship)`

SetIpLeases sets IpLeases field to given value.

### HasIpLeases

`func (o *IppoolBlockLease) HasIpLeases() bool`

HasIpLeases returns a boolean if a field has been set.

### SetIpLeasesNil

`func (o *IppoolBlockLease) SetIpLeasesNil(b bool)`

 SetIpLeasesNil sets the value for IpLeases to be an explicit nil

### UnsetIpLeases
`func (o *IppoolBlockLease) UnsetIpLeases()`

UnsetIpLeases ensures that no value is present for IpLeases, not even an explicit nil
### GetPool

`func (o *IppoolBlockLease) GetPool() IppoolPoolRelationship`

GetPool returns the Pool field if non-nil, zero value otherwise.

### GetPoolOk

`func (o *IppoolBlockLease) GetPoolOk() (*IppoolPoolRelationship, bool)`

GetPoolOk returns a tuple with the Pool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPool

`func (o *IppoolBlockLease) SetPool(v IppoolPoolRelationship)`

SetPool sets Pool field to given value.

### HasPool

`func (o *IppoolBlockLease) HasPool() bool`

HasPool returns a boolean if a field has been set.

### GetUniverse

`func (o *IppoolBlockLease) GetUniverse() IppoolUniverseRelationship`

GetUniverse returns the Universe field if non-nil, zero value otherwise.

### GetUniverseOk

`func (o *IppoolBlockLease) GetUniverseOk() (*IppoolUniverseRelationship, bool)`

GetUniverseOk returns a tuple with the Universe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniverse

`func (o *IppoolBlockLease) SetUniverse(v IppoolUniverseRelationship)`

SetUniverse sets Universe field to given value.

### HasUniverse

`func (o *IppoolBlockLease) HasUniverse() bool`

HasUniverse returns a boolean if a field has been set.

### GetVrf

`func (o *IppoolBlockLease) GetVrf() VrfVrfRelationship`

GetVrf returns the Vrf field if non-nil, zero value otherwise.

### GetVrfOk

`func (o *IppoolBlockLease) GetVrfOk() (*VrfVrfRelationship, bool)`

GetVrfOk returns a tuple with the Vrf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrf

`func (o *IppoolBlockLease) SetVrf(v VrfVrfRelationship)`

SetVrf sets Vrf field to given value.

### HasVrf

`func (o *IppoolBlockLease) HasVrf() bool`

HasVrf returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


