# ResourceSharedResourcesInfoHolderAllOf

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

### NewResourceSharedResourcesInfoHolderAllOf

`func NewResourceSharedResourcesInfoHolderAllOf(classId string, objectType string, ) *ResourceSharedResourcesInfoHolderAllOf`

NewResourceSharedResourcesInfoHolderAllOf instantiates a new ResourceSharedResourcesInfoHolderAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceSharedResourcesInfoHolderAllOfWithDefaults

`func NewResourceSharedResourcesInfoHolderAllOfWithDefaults() *ResourceSharedResourcesInfoHolderAllOf`

NewResourceSharedResourcesInfoHolderAllOfWithDefaults instantiates a new ResourceSharedResourcesInfoHolderAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ResourceSharedResourcesInfoHolderAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ResourceSharedResourcesInfoHolderAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ResourceSharedResourcesInfoHolderAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ResourceSharedResourcesInfoHolderAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ResourceSharedResourcesInfoHolderAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ResourceSharedResourcesInfoHolderAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAccount

`func (o *ResourceSharedResourcesInfoHolderAllOf) GetAccount() IamAccountRelationship`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *ResourceSharedResourcesInfoHolderAllOf) GetAccountOk() (*IamAccountRelationship, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *ResourceSharedResourcesInfoHolderAllOf) SetAccount(v IamAccountRelationship)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *ResourceSharedResourcesInfoHolderAllOf) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetPeerObjects

`func (o *ResourceSharedResourcesInfoHolderAllOf) GetPeerObjects() []MoBaseMoRelationship`

GetPeerObjects returns the PeerObjects field if non-nil, zero value otherwise.

### GetPeerObjectsOk

`func (o *ResourceSharedResourcesInfoHolderAllOf) GetPeerObjectsOk() (*[]MoBaseMoRelationship, bool)`

GetPeerObjectsOk returns a tuple with the PeerObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerObjects

`func (o *ResourceSharedResourcesInfoHolderAllOf) SetPeerObjects(v []MoBaseMoRelationship)`

SetPeerObjects sets PeerObjects field to given value.

### HasPeerObjects

`func (o *ResourceSharedResourcesInfoHolderAllOf) HasPeerObjects() bool`

HasPeerObjects returns a boolean if a field has been set.

### SetPeerObjectsNil

`func (o *ResourceSharedResourcesInfoHolderAllOf) SetPeerObjectsNil(b bool)`

 SetPeerObjectsNil sets the value for PeerObjects to be an explicit nil

### UnsetPeerObjects
`func (o *ResourceSharedResourcesInfoHolderAllOf) UnsetPeerObjects()`

UnsetPeerObjects ensures that no value is present for PeerObjects, not even an explicit nil
### GetSharedResource

`func (o *ResourceSharedResourcesInfoHolderAllOf) GetSharedResource() MoBaseMoRelationship`

GetSharedResource returns the SharedResource field if non-nil, zero value otherwise.

### GetSharedResourceOk

`func (o *ResourceSharedResourcesInfoHolderAllOf) GetSharedResourceOk() (*MoBaseMoRelationship, bool)`

GetSharedResourceOk returns a tuple with the SharedResource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedResource

`func (o *ResourceSharedResourcesInfoHolderAllOf) SetSharedResource(v MoBaseMoRelationship)`

SetSharedResource sets SharedResource field to given value.

### HasSharedResource

`func (o *ResourceSharedResourcesInfoHolderAllOf) HasSharedResource() bool`

HasSharedResource returns a boolean if a field has been set.

### GetSharedWithResource

`func (o *ResourceSharedResourcesInfoHolderAllOf) GetSharedWithResource() MoBaseMoRelationship`

GetSharedWithResource returns the SharedWithResource field if non-nil, zero value otherwise.

### GetSharedWithResourceOk

`func (o *ResourceSharedResourcesInfoHolderAllOf) GetSharedWithResourceOk() (*MoBaseMoRelationship, bool)`

GetSharedWithResourceOk returns a tuple with the SharedWithResource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedWithResource

`func (o *ResourceSharedResourcesInfoHolderAllOf) SetSharedWithResource(v MoBaseMoRelationship)`

SetSharedWithResource sets SharedWithResource field to given value.

### HasSharedWithResource

`func (o *ResourceSharedResourcesInfoHolderAllOf) HasSharedWithResource() bool`

HasSharedWithResource returns a boolean if a field has been set.

### GetSharingRule

`func (o *ResourceSharedResourcesInfoHolderAllOf) GetSharingRule() IamSharingRuleRelationship`

GetSharingRule returns the SharingRule field if non-nil, zero value otherwise.

### GetSharingRuleOk

`func (o *ResourceSharedResourcesInfoHolderAllOf) GetSharingRuleOk() (*IamSharingRuleRelationship, bool)`

GetSharingRuleOk returns a tuple with the SharingRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharingRule

`func (o *ResourceSharedResourcesInfoHolderAllOf) SetSharingRule(v IamSharingRuleRelationship)`

SetSharingRule sets SharingRule field to given value.

### HasSharingRule

`func (o *ResourceSharedResourcesInfoHolderAllOf) HasSharingRule() bool`

HasSharingRule returns a boolean if a field has been set.

### GetSourceObject

`func (o *ResourceSharedResourcesInfoHolderAllOf) GetSourceObject() MoBaseMoRelationship`

GetSourceObject returns the SourceObject field if non-nil, zero value otherwise.

### GetSourceObjectOk

`func (o *ResourceSharedResourcesInfoHolderAllOf) GetSourceObjectOk() (*MoBaseMoRelationship, bool)`

GetSourceObjectOk returns a tuple with the SourceObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceObject

`func (o *ResourceSharedResourcesInfoHolderAllOf) SetSourceObject(v MoBaseMoRelationship)`

SetSourceObject sets SourceObject field to given value.

### HasSourceObject

`func (o *ResourceSharedResourcesInfoHolderAllOf) HasSourceObject() bool`

HasSourceObject returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


