# IamServiceProvider

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntityId** | Pointer to **string** | Entity ID of the Intersight Service Provider. In SAML, the entity ID uniquely identifies the IdP/Service Provider. | [optional] [readonly] 
**Metadata** | Pointer to **string** | Metadata of the Intersight Service Provider. User downloads the Intersight Service Provider metadata and integrates it with their IdP for authentication. | [optional] [readonly] 
**Name** | Pointer to **string** | Name of the Intersight Service Provider. | [optional] [readonly] 
**System** | Pointer to [**IamSystemRelationship**](iam.System.Relationship.md) |  | [optional] 

## Methods

### NewIamServiceProvider

`func NewIamServiceProvider() *IamServiceProvider`

NewIamServiceProvider instantiates a new IamServiceProvider object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamServiceProviderWithDefaults

`func NewIamServiceProviderWithDefaults() *IamServiceProvider`

NewIamServiceProviderWithDefaults instantiates a new IamServiceProvider object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


