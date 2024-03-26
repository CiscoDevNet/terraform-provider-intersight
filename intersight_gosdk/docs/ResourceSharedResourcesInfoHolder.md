# ResourceSharedResourcesInfoHolder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "resource.SharedResourcesInfoHolder"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "resource.SharedResourcesInfoHolder"]
**Account** | Pointer to [**IamAccountRelationship**](IamAccountRelationship.md) |  | [optional] 
**PeerObjects** | Pointer to [**[]MoBaseMoRelationship**](MoBaseMoRelationship.md) | An array of relationships to moBaseMo resources. | [optional] [readonly] 
**SharedResource** | Pointer to [**MoBaseMoRelationship**](MoBaseMoRelationship.md) |  | [optional] 
**SharedWithResource** | Pointer to [**MoBaseMoRelationship**](MoBaseMoRelationship.md) |  | [optional] 
**SharingRule** | Pointer to [**IamSharingRuleRelationship**](IamSharingRuleRelationship.md) |  | [optional] 
**SourceObject** | Pointer to [**MoBaseMoRelationship**](MoBaseMoRelationship.md) |  | [optional] 

## Methods

### NewResourceSharedResourcesInfoHolder

`func NewResourceSharedResourcesInfoHolder(classId string, objectType string, ) *ResourceSharedResourcesInfoHolder`

NewResourceSharedResourcesInfoHolder instantiates a new ResourceSharedResourcesInfoHolder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceSharedResourcesInfoHolderWithDefaults

`func NewResourceSharedResourcesInfoHolderWithDefaults() *ResourceSharedResourcesInfoHolder`

NewResourceSharedResourcesInfoHolderWithDefaults instantiates a new ResourceSharedResourcesInfoHolder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ResourceSharedResourcesInfoHolder) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ResourceSharedResourcesInfoHolder) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ResourceSharedResourcesInfoHolder) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ResourceSharedResourcesInfoHolder) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ResourceSharedResourcesInfoHolder) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ResourceSharedResourcesInfoHolder) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAccount

`func (o *ResourceSharedResourcesInfoHolder) GetAccount() IamAccountRelationship`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *ResourceSharedResourcesInfoHolder) GetAccountOk() (*IamAccountRelationship, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *ResourceSharedResourcesInfoHolder) SetAccount(v IamAccountRelationship)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *ResourceSharedResourcesInfoHolder) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetPeerObjects

`func (o *ResourceSharedResourcesInfoHolder) GetPeerObjects() []MoBaseMoRelationship`

GetPeerObjects returns the PeerObjects field if non-nil, zero value otherwise.

### GetPeerObjectsOk

`func (o *ResourceSharedResourcesInfoHolder) GetPeerObjectsOk() (*[]MoBaseMoRelationship, bool)`

GetPeerObjectsOk returns a tuple with the PeerObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerObjects

`func (o *ResourceSharedResourcesInfoHolder) SetPeerObjects(v []MoBaseMoRelationship)`

SetPeerObjects sets PeerObjects field to given value.

### HasPeerObjects

`func (o *ResourceSharedResourcesInfoHolder) HasPeerObjects() bool`

HasPeerObjects returns a boolean if a field has been set.

### SetPeerObjectsNil

`func (o *ResourceSharedResourcesInfoHolder) SetPeerObjectsNil(b bool)`

 SetPeerObjectsNil sets the value for PeerObjects to be an explicit nil

### UnsetPeerObjects
`func (o *ResourceSharedResourcesInfoHolder) UnsetPeerObjects()`

UnsetPeerObjects ensures that no value is present for PeerObjects, not even an explicit nil
### GetSharedResource

`func (o *ResourceSharedResourcesInfoHolder) GetSharedResource() MoBaseMoRelationship`

GetSharedResource returns the SharedResource field if non-nil, zero value otherwise.

### GetSharedResourceOk

`func (o *ResourceSharedResourcesInfoHolder) GetSharedResourceOk() (*MoBaseMoRelationship, bool)`

GetSharedResourceOk returns a tuple with the SharedResource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedResource

`func (o *ResourceSharedResourcesInfoHolder) SetSharedResource(v MoBaseMoRelationship)`

SetSharedResource sets SharedResource field to given value.

### HasSharedResource

`func (o *ResourceSharedResourcesInfoHolder) HasSharedResource() bool`

HasSharedResource returns a boolean if a field has been set.

### GetSharedWithResource

`func (o *ResourceSharedResourcesInfoHolder) GetSharedWithResource() MoBaseMoRelationship`

GetSharedWithResource returns the SharedWithResource field if non-nil, zero value otherwise.

### GetSharedWithResourceOk

`func (o *ResourceSharedResourcesInfoHolder) GetSharedWithResourceOk() (*MoBaseMoRelationship, bool)`

GetSharedWithResourceOk returns a tuple with the SharedWithResource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedWithResource

`func (o *ResourceSharedResourcesInfoHolder) SetSharedWithResource(v MoBaseMoRelationship)`

SetSharedWithResource sets SharedWithResource field to given value.

### HasSharedWithResource

`func (o *ResourceSharedResourcesInfoHolder) HasSharedWithResource() bool`

HasSharedWithResource returns a boolean if a field has been set.

### GetSharingRule

`func (o *ResourceSharedResourcesInfoHolder) GetSharingRule() IamSharingRuleRelationship`

GetSharingRule returns the SharingRule field if non-nil, zero value otherwise.

### GetSharingRuleOk

`func (o *ResourceSharedResourcesInfoHolder) GetSharingRuleOk() (*IamSharingRuleRelationship, bool)`

GetSharingRuleOk returns a tuple with the SharingRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharingRule

`func (o *ResourceSharedResourcesInfoHolder) SetSharingRule(v IamSharingRuleRelationship)`

SetSharingRule sets SharingRule field to given value.

### HasSharingRule

`func (o *ResourceSharedResourcesInfoHolder) HasSharingRule() bool`

HasSharingRule returns a boolean if a field has been set.

### GetSourceObject

`func (o *ResourceSharedResourcesInfoHolder) GetSourceObject() MoBaseMoRelationship`

GetSourceObject returns the SourceObject field if non-nil, zero value otherwise.

### GetSourceObjectOk

`func (o *ResourceSharedResourcesInfoHolder) GetSourceObjectOk() (*MoBaseMoRelationship, bool)`

GetSourceObjectOk returns a tuple with the SourceObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceObject

`func (o *ResourceSharedResourcesInfoHolder) SetSourceObject(v MoBaseMoRelationship)`

SetSourceObject sets SourceObject field to given value.

### HasSourceObject

`func (o *ResourceSharedResourcesInfoHolder) HasSourceObject() bool`

HasSourceObject returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


