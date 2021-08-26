# TamAdvisoryInfoAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "tam.AdvisoryInfo"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "tam.AdvisoryInfo"]
**State** | Pointer to **string** | Current state of the advisory for the owner. Indicates if the user is interested in getting updates for the advisory. * &#x60;active&#x60; - Advisory is currently active and the user wants to receive updates for this advisory. * &#x60;acknowledged&#x60; - Advisory is seen and acknowledged by the user and she no longer wants to recieve updates. | [optional] [default to "active"]
**Account** | Pointer to [**IamAccountRelationship**](IamAccountRelationship.md) |  | [optional] 
**Advisory** | Pointer to [**TamBaseAdvisoryRelationship**](TamBaseAdvisoryRelationship.md) |  | [optional] 

## Methods

### NewTamAdvisoryInfoAllOf

`func NewTamAdvisoryInfoAllOf(classId string, objectType string, ) *TamAdvisoryInfoAllOf`

NewTamAdvisoryInfoAllOf instantiates a new TamAdvisoryInfoAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTamAdvisoryInfoAllOfWithDefaults

`func NewTamAdvisoryInfoAllOfWithDefaults() *TamAdvisoryInfoAllOf`

NewTamAdvisoryInfoAllOfWithDefaults instantiates a new TamAdvisoryInfoAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *TamAdvisoryInfoAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *TamAdvisoryInfoAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *TamAdvisoryInfoAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *TamAdvisoryInfoAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *TamAdvisoryInfoAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *TamAdvisoryInfoAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetState

`func (o *TamAdvisoryInfoAllOf) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *TamAdvisoryInfoAllOf) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *TamAdvisoryInfoAllOf) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *TamAdvisoryInfoAllOf) HasState() bool`

HasState returns a boolean if a field has been set.

### GetAccount

`func (o *TamAdvisoryInfoAllOf) GetAccount() IamAccountRelationship`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *TamAdvisoryInfoAllOf) GetAccountOk() (*IamAccountRelationship, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *TamAdvisoryInfoAllOf) SetAccount(v IamAccountRelationship)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *TamAdvisoryInfoAllOf) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetAdvisory

`func (o *TamAdvisoryInfoAllOf) GetAdvisory() TamBaseAdvisoryRelationship`

GetAdvisory returns the Advisory field if non-nil, zero value otherwise.

### GetAdvisoryOk

`func (o *TamAdvisoryInfoAllOf) GetAdvisoryOk() (*TamBaseAdvisoryRelationship, bool)`

GetAdvisoryOk returns a tuple with the Advisory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvisory

`func (o *TamAdvisoryInfoAllOf) SetAdvisory(v TamBaseAdvisoryRelationship)`

SetAdvisory sets Advisory field to given value.

### HasAdvisory

`func (o *TamAdvisoryInfoAllOf) HasAdvisory() bool`

HasAdvisory returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


