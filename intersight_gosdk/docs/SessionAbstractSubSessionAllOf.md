# SessionAbstractSubSessionAllOf

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

### NewSessionAbstractSubSessionAllOf

`func NewSessionAbstractSubSessionAllOf(classId string, objectType string, ) *SessionAbstractSubSessionAllOf`

NewSessionAbstractSubSessionAllOf instantiates a new SessionAbstractSubSessionAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSessionAbstractSubSessionAllOfWithDefaults

`func NewSessionAbstractSubSessionAllOfWithDefaults() *SessionAbstractSubSessionAllOf`

NewSessionAbstractSubSessionAllOfWithDefaults instantiates a new SessionAbstractSubSessionAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *SessionAbstractSubSessionAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *SessionAbstractSubSessionAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *SessionAbstractSubSessionAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *SessionAbstractSubSessionAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *SessionAbstractSubSessionAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *SessionAbstractSubSessionAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetTargetName

`func (o *SessionAbstractSubSessionAllOf) GetTargetName() string`

GetTargetName returns the TargetName field if non-nil, zero value otherwise.

### GetTargetNameOk

`func (o *SessionAbstractSubSessionAllOf) GetTargetNameOk() (*string, bool)`

GetTargetNameOk returns a tuple with the TargetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetName

`func (o *SessionAbstractSubSessionAllOf) SetTargetName(v string)`

SetTargetName sets TargetName field to given value.

### HasTargetName

`func (o *SessionAbstractSubSessionAllOf) HasTargetName() bool`

HasTargetName returns a boolean if a field has been set.

### GetSession

`func (o *SessionAbstractSubSessionAllOf) GetSession() SessionAbstractSessionRelationship`

GetSession returns the Session field if non-nil, zero value otherwise.

### GetSessionOk

`func (o *SessionAbstractSubSessionAllOf) GetSessionOk() (*SessionAbstractSessionRelationship, bool)`

GetSessionOk returns a tuple with the Session field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSession

`func (o *SessionAbstractSubSessionAllOf) SetSession(v SessionAbstractSessionRelationship)`

SetSession sets Session field to given value.

### HasSession

`func (o *SessionAbstractSubSessionAllOf) HasSession() bool`

HasSession returns a boolean if a field has been set.

### GetTarget

`func (o *SessionAbstractSubSessionAllOf) GetTarget() MoBaseMoRelationship`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *SessionAbstractSubSessionAllOf) GetTargetOk() (*MoBaseMoRelationship, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *SessionAbstractSubSessionAllOf) SetTarget(v MoBaseMoRelationship)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *SessionAbstractSubSessionAllOf) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### GetUser

`func (o *SessionAbstractSubSessionAllOf) GetUser() IamUserRelationship`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *SessionAbstractSubSessionAllOf) GetUserOk() (*IamUserRelationship, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *SessionAbstractSubSessionAllOf) SetUser(v IamUserRelationship)`

SetUser sets User field to given value.

### HasUser

`func (o *SessionAbstractSubSessionAllOf) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


