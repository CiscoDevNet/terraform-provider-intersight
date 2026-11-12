# IamServiceProvider

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "iam.ServiceProvider"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "iam.ServiceProvider"]
**CertExpiryTime** | Pointer to **time.Time** | Expiry time of the certificate used in SP metadata. Tracks SP certificate expiries and rotates the certificate before it expires to avoid any authentication failure. | [optional] [readonly] 
**EntityId** | Pointer to **string** | Entity ID of the Intersight Service Provider. In SAML, the entity ID uniquely identifies the IdP/Service Provider. | [optional] [readonly] 
**Metadata** | Pointer to **string** | Metadata of the Intersight Service Provider. User downloads the Intersight Service Provider metadata and integrates it with their IdP for authentication. | [optional] [readonly] 
**Name** | Pointer to **string** | Name of the Intersight Service Provider. | [optional] [readonly] 
**SpMetadataExpiryState** | Pointer to **string** | The expiration state of the certificates in the SP metadata rotation during the grace period of SP SAML metadata rotation. * &#x60;expiration-within-info-period&#x60; - The expiry date is beyond 270 days. * &#x60;expiration-within-270-day-info-period&#x60; - The expiry date is within 270 days but beyond 180 days. * &#x60;expiration-within-180-day-info-period&#x60; - The expiry date is within 180 days but beyond 90 days. * &#x60;expiration-within-90-day-warning-period&#x60; - The expiry date is within 90 days but beyond 60 days. * &#x60;expiration-within-60-day-warning-period&#x60; - The expiry date is within 60 days but beyond 30 days. * &#x60;expiration-within-30-day-critical-period&#x60; - The expiry date is within 30 days but beyond 7 days. * &#x60;expiration-within-critical-period&#x60; - The expiry date is within 7 days. * &#x60;expired&#x60; - The expiry date has passed. * &#x60;not-applicable&#x60; - The SP certificate is not used because SAML assertion encryption and AuthNRequest signing are both disabled. | [optional] [readonly] [default to "expiration-within-info-period"]
**SpMetadataType** | Pointer to **string** | Defines which SP metadata is defined in Intersight. The SP metadata could be legacy metadata or new metadata. * &#x60;Sha2&#x60; - The new SAML SP metadata with SHA-2 hashed certificates. * &#x60;Legacy&#x60; - The legacy SAML SP metadata (expiring certificates). | [optional] [readonly] [default to "Sha2"]
**System** | Pointer to [**NullableIamSystemRelationship**](IamSystemRelationship.md) |  | [optional] 

## Methods

### NewIamServiceProvider

`func NewIamServiceProvider(classId string, objectType string, ) *IamServiceProvider`

NewIamServiceProvider instantiates a new IamServiceProvider object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamServiceProviderWithDefaults

`func NewIamServiceProviderWithDefaults() *IamServiceProvider`

NewIamServiceProviderWithDefaults instantiates a new IamServiceProvider object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *IamServiceProvider) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *IamServiceProvider) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *IamServiceProvider) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *IamServiceProvider) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *IamServiceProvider) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *IamServiceProvider) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCertExpiryTime

`func (o *IamServiceProvider) GetCertExpiryTime() time.Time`

GetCertExpiryTime returns the CertExpiryTime field if non-nil, zero value otherwise.

### GetCertExpiryTimeOk

`func (o *IamServiceProvider) GetCertExpiryTimeOk() (*time.Time, bool)`

GetCertExpiryTimeOk returns a tuple with the CertExpiryTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertExpiryTime

`func (o *IamServiceProvider) SetCertExpiryTime(v time.Time)`

SetCertExpiryTime sets CertExpiryTime field to given value.

### HasCertExpiryTime

`func (o *IamServiceProvider) HasCertExpiryTime() bool`

HasCertExpiryTime returns a boolean if a field has been set.

### GetEntityId

`func (o *IamServiceProvider) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *IamServiceProvider) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *IamServiceProvider) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.

### HasEntityId

`func (o *IamServiceProvider) HasEntityId() bool`

HasEntityId returns a boolean if a field has been set.

### GetMetadata

`func (o *IamServiceProvider) GetMetadata() string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *IamServiceProvider) GetMetadataOk() (*string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *IamServiceProvider) SetMetadata(v string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *IamServiceProvider) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetName

`func (o *IamServiceProvider) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IamServiceProvider) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IamServiceProvider) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IamServiceProvider) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSpMetadataExpiryState

`func (o *IamServiceProvider) GetSpMetadataExpiryState() string`

GetSpMetadataExpiryState returns the SpMetadataExpiryState field if non-nil, zero value otherwise.

### GetSpMetadataExpiryStateOk

`func (o *IamServiceProvider) GetSpMetadataExpiryStateOk() (*string, bool)`

GetSpMetadataExpiryStateOk returns a tuple with the SpMetadataExpiryState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpMetadataExpiryState

`func (o *IamServiceProvider) SetSpMetadataExpiryState(v string)`

SetSpMetadataExpiryState sets SpMetadataExpiryState field to given value.

### HasSpMetadataExpiryState

`func (o *IamServiceProvider) HasSpMetadataExpiryState() bool`

HasSpMetadataExpiryState returns a boolean if a field has been set.

### GetSpMetadataType

`func (o *IamServiceProvider) GetSpMetadataType() string`

GetSpMetadataType returns the SpMetadataType field if non-nil, zero value otherwise.

### GetSpMetadataTypeOk

`func (o *IamServiceProvider) GetSpMetadataTypeOk() (*string, bool)`

GetSpMetadataTypeOk returns a tuple with the SpMetadataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpMetadataType

`func (o *IamServiceProvider) SetSpMetadataType(v string)`

SetSpMetadataType sets SpMetadataType field to given value.

### HasSpMetadataType

`func (o *IamServiceProvider) HasSpMetadataType() bool`

HasSpMetadataType returns a boolean if a field has been set.

### GetSystem

`func (o *IamServiceProvider) GetSystem() IamSystemRelationship`

GetSystem returns the System field if non-nil, zero value otherwise.

### GetSystemOk

`func (o *IamServiceProvider) GetSystemOk() (*IamSystemRelationship, bool)`

GetSystemOk returns a tuple with the System field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystem

`func (o *IamServiceProvider) SetSystem(v IamSystemRelationship)`

SetSystem sets System field to given value.

### HasSystem

`func (o *IamServiceProvider) HasSystem() bool`

HasSystem returns a boolean if a field has been set.

### SetSystemNil

`func (o *IamServiceProvider) SetSystemNil(b bool)`

 SetSystemNil sets the value for System to be an explicit nil

### UnsetSystem
`func (o *IamServiceProvider) UnsetSystem()`

UnsetSystem ensures that no value is present for System, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


