# IppoolBlockLeaseAllOf

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

### NewIppoolBlockLeaseAllOf

`func NewIppoolBlockLeaseAllOf(classId string, objectType string, ) *IppoolBlockLeaseAllOf`

NewIppoolBlockLeaseAllOf instantiates a new IppoolBlockLeaseAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIppoolBlockLeaseAllOfWithDefaults

`func NewIppoolBlockLeaseAllOfWithDefaults() *IppoolBlockLeaseAllOf`

NewIppoolBlockLeaseAllOfWithDefaults instantiates a new IppoolBlockLeaseAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *IppoolBlockLeaseAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *IppoolBlockLeaseAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *IppoolBlockLeaseAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *IppoolBlockLeaseAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *IppoolBlockLeaseAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *IppoolBlockLeaseAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetIpType

`func (o *IppoolBlockLeaseAllOf) GetIpType() string`

GetIpType returns the IpType field if non-nil, zero value otherwise.

### GetIpTypeOk

`func (o *IppoolBlockLeaseAllOf) GetIpTypeOk() (*string, bool)`

GetIpTypeOk returns a tuple with the IpType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpType

`func (o *IppoolBlockLeaseAllOf) SetIpType(v string)`

SetIpType sets IpType field to given value.

### HasIpType

`func (o *IppoolBlockLeaseAllOf) HasIpType() bool`

HasIpType returns a boolean if a field has been set.

### GetAssignedToEntity

`func (o *IppoolBlockLeaseAllOf) GetAssignedToEntity() MoBaseMoRelationship`

GetAssignedToEntity returns the AssignedToEntity field if non-nil, zero value otherwise.

### GetAssignedToEntityOk

`func (o *IppoolBlockLeaseAllOf) GetAssignedToEntityOk() (*MoBaseMoRelationship, bool)`

GetAssignedToEntityOk returns a tuple with the AssignedToEntity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedToEntity

`func (o *IppoolBlockLeaseAllOf) SetAssignedToEntity(v MoBaseMoRelationship)`

SetAssignedToEntity sets AssignedToEntity field to given value.

### HasAssignedToEntity

`func (o *IppoolBlockLeaseAllOf) HasAssignedToEntity() bool`

HasAssignedToEntity returns a boolean if a field has been set.

### GetIpLeases

`func (o *IppoolBlockLeaseAllOf) GetIpLeases() []IppoolIpLeaseRelationship`

GetIpLeases returns the IpLeases field if non-nil, zero value otherwise.

### GetIpLeasesOk

`func (o *IppoolBlockLeaseAllOf) GetIpLeasesOk() (*[]IppoolIpLeaseRelationship, bool)`

GetIpLeasesOk returns a tuple with the IpLeases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpLeases

`func (o *IppoolBlockLeaseAllOf) SetIpLeases(v []IppoolIpLeaseRelationship)`

SetIpLeases sets IpLeases field to given value.

### HasIpLeases

`func (o *IppoolBlockLeaseAllOf) HasIpLeases() bool`

HasIpLeases returns a boolean if a field has been set.

### SetIpLeasesNil

`func (o *IppoolBlockLeaseAllOf) SetIpLeasesNil(b bool)`

 SetIpLeasesNil sets the value for IpLeases to be an explicit nil

### UnsetIpLeases
`func (o *IppoolBlockLeaseAllOf) UnsetIpLeases()`

UnsetIpLeases ensures that no value is present for IpLeases, not even an explicit nil
### GetPool

`func (o *IppoolBlockLeaseAllOf) GetPool() IppoolPoolRelationship`

GetPool returns the Pool field if non-nil, zero value otherwise.

### GetPoolOk

`func (o *IppoolBlockLeaseAllOf) GetPoolOk() (*IppoolPoolRelationship, bool)`

GetPoolOk returns a tuple with the Pool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPool

`func (o *IppoolBlockLeaseAllOf) SetPool(v IppoolPoolRelationship)`

SetPool sets Pool field to given value.

### HasPool

`func (o *IppoolBlockLeaseAllOf) HasPool() bool`

HasPool returns a boolean if a field has been set.

### GetUniverse

`func (o *IppoolBlockLeaseAllOf) GetUniverse() IppoolUniverseRelationship`

GetUniverse returns the Universe field if non-nil, zero value otherwise.

### GetUniverseOk

`func (o *IppoolBlockLeaseAllOf) GetUniverseOk() (*IppoolUniverseRelationship, bool)`

GetUniverseOk returns a tuple with the Universe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniverse

`func (o *IppoolBlockLeaseAllOf) SetUniverse(v IppoolUniverseRelationship)`

SetUniverse sets Universe field to given value.

### HasUniverse

`func (o *IppoolBlockLeaseAllOf) HasUniverse() bool`

HasUniverse returns a boolean if a field has been set.

### GetVrf

`func (o *IppoolBlockLeaseAllOf) GetVrf() VrfVrfRelationship`

GetVrf returns the Vrf field if non-nil, zero value otherwise.

### GetVrfOk

`func (o *IppoolBlockLeaseAllOf) GetVrfOk() (*VrfVrfRelationship, bool)`

GetVrfOk returns a tuple with the Vrf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrf

`func (o *IppoolBlockLeaseAllOf) SetVrf(v VrfVrfRelationship)`

SetVrf sets Vrf field to given value.

### HasVrf

`func (o *IppoolBlockLeaseAllOf) HasVrf() bool`

HasVrf returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


