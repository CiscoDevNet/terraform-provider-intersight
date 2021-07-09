# SessionAbstractSubSession

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**TargetName** | Pointer to **string** | Name of target on which session is initiated. | [optional] [readonly] 
**Session** | Pointer to [**SessionAbstractSessionRelationship**](session.AbstractSession.Relationship.md) |  | [optional] 
**Target** | Pointer to [**MoBaseMoRelationship**](mo.BaseMo.Relationship.md) |  | [optional] 
**User** | Pointer to [**IamUserRelationship**](iam.User.Relationship.md) |  | [optional] 

## Methods

### NewSessionAbstractSubSession

`func NewSessionAbstractSubSession(classId string, objectType string, ) *SessionAbstractSubSession`

NewSessionAbstractSubSession instantiates a new SessionAbstractSubSession object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSessionAbstractSubSessionWithDefaults

`func NewSessionAbstractSubSessionWithDefaults() *SessionAbstractSubSession`

NewSessionAbstractSubSessionWithDefaults instantiates a new SessionAbstractSubSession object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *SessionAbstractSubSession) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *SessionAbstractSubSession) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *SessionAbstractSubSession) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *SessionAbstractSubSession) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *SessionAbstractSubSession) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *SessionAbstractSubSession) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetTargetName

`func (o *SessionAbstractSubSession) GetTargetName() string`

GetTargetName returns the TargetName field if non-nil, zero value otherwise.

### GetTargetNameOk

`func (o *SessionAbstractSubSession) GetTargetNameOk() (*string, bool)`

GetTargetNameOk returns a tuple with the TargetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetName

`func (o *SessionAbstractSubSession) SetTargetName(v string)`

SetTargetName sets TargetName field to given value.

### HasTargetName

`func (o *SessionAbstractSubSession) HasTargetName() bool`

HasTargetName returns a boolean if a field has been set.

### GetSession

`func (o *SessionAbstractSubSession) GetSession() SessionAbstractSessionRelationship`

GetSession returns the Session field if non-nil, zero value otherwise.

### GetSessionOk

`func (o *SessionAbstractSubSession) GetSessionOk() (*SessionAbstractSessionRelationship, bool)`

GetSessionOk returns a tuple with the Session field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSession

`func (o *SessionAbstractSubSession) SetSession(v SessionAbstractSessionRelationship)`

SetSession sets Session field to given value.

### HasSession

`func (o *SessionAbstractSubSession) HasSession() bool`

HasSession returns a boolean if a field has been set.

### GetTarget

`func (o *SessionAbstractSubSession) GetTarget() MoBaseMoRelationship`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *SessionAbstractSubSession) GetTargetOk() (*MoBaseMoRelationship, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *SessionAbstractSubSession) SetTarget(v MoBaseMoRelationship)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *SessionAbstractSubSession) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### GetUser

`func (o *SessionAbstractSubSession) GetUser() IamUserRelationship`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *SessionAbstractSubSession) GetUserOk() (*IamUserRelationship, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *SessionAbstractSubSession) SetUser(v IamUserRelationship)`

SetUser sets User field to given value.

### HasUser

`func (o *SessionAbstractSubSession) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


