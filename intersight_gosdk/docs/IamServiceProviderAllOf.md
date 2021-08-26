# IamServiceProviderAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "iam.ServiceProvider"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "iam.ServiceProvider"]
**EntityId** | Pointer to **string** | Entity ID of the Intersight Service Provider. In SAML, the entity ID uniquely identifies the IdP/Service Provider. | [optional] [readonly] 
**Metadata** | Pointer to **string** | Metadata of the Intersight Service Provider. User downloads the Intersight Service Provider metadata and integrates it with their IdP for authentication. | [optional] [readonly] 
**Name** | Pointer to **string** | Name of the Intersight Service Provider. | [optional] [readonly] 
**System** | Pointer to [**IamSystemRelationship**](IamSystemRelationship.md) |  | [optional] 

## Methods

### NewIamServiceProviderAllOf

`func NewIamServiceProviderAllOf(classId string, objectType string, ) *IamServiceProviderAllOf`

NewIamServiceProviderAllOf instantiates a new IamServiceProviderAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamServiceProviderAllOfWithDefaults

`func NewIamServiceProviderAllOfWithDefaults() *IamServiceProviderAllOf`

NewIamServiceProviderAllOfWithDefaults instantiates a new IamServiceProviderAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *IamServiceProviderAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *IamServiceProviderAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *IamServiceProviderAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *IamServiceProviderAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *IamServiceProviderAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *IamServiceProviderAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetEntityId

`func (o *IamServiceProviderAllOf) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *IamServiceProviderAllOf) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *IamServiceProviderAllOf) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.

### HasEntityId

`func (o *IamServiceProviderAllOf) HasEntityId() bool`

HasEntityId returns a boolean if a field has been set.

### GetMetadata

`func (o *IamServiceProviderAllOf) GetMetadata() string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *IamServiceProviderAllOf) GetMetadataOk() (*string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *IamServiceProviderAllOf) SetMetadata(v string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *IamServiceProviderAllOf) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetName

`func (o *IamServiceProviderAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IamServiceProviderAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IamServiceProviderAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IamServiceProviderAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSystem

`func (o *IamServiceProviderAllOf) GetSystem() IamSystemRelationship`

GetSystem returns the System field if non-nil, zero value otherwise.

### GetSystemOk

`func (o *IamServiceProviderAllOf) GetSystemOk() (*IamSystemRelationship, bool)`

GetSystemOk returns a tuple with the System field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystem

`func (o *IamServiceProviderAllOf) SetSystem(v IamSystemRelationship)`

SetSystem sets System field to given value.

### HasSystem

`func (o *IamServiceProviderAllOf) HasSystem() bool`

HasSystem returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


