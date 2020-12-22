# TamAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "tam.Action"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "tam.Action"]
**AffectedObjectType** | Pointer to **string** | Type of the managed object that should be marked with an instance of the Alert (when operation type is create) or that should have an alert instance removed (when operation type is remove). | [optional] 
**AlertType** | Pointer to **string** | Alert type is used to denote the category of an Intersight alert (FieldNotice, equipment Fault etc.). * &#x60;psirt&#x60; - Respresents the psirt alert type (https://tools.cisco.com/security/center/publicationListing.x). * &#x60;fieldNotice&#x60; - Respresents the field notice alert type (https://www.cisco.com/c/en/us/support/web/tsd-products-field-notice-summary.html). | [optional] [default to "psirt"]
**Identifiers** | Pointer to [**[]TamIdentifiers**](TamIdentifiers.md) |  | [optional] 
**Name** | Pointer to **string** | Uniquely identifies a given action among the set of actions corresponding to an advisory. Primarily used to store and compare results of subsequent iterations corresponding to the action queries. | [optional] 
**OperationType** | Pointer to **string** | Operation type for the alert action. An action is used to carry out the process of \&quot;reacting\&quot; to an alert condition. For e.g.in case of a fieldNotice alert, the intention may be to create a new alert (if the condition matches and there is no existing alert) or to remove an existing alert when the alert condition has been remedied. * &#x60;create&#x60; - Create an instance of AdvisoryInstance. * &#x60;remove&#x60; - Remove an instance of AdvisoryInstance. | [optional] [default to "create"]
**Queries** | Pointer to [**[]TamQueryEntry**](TamQueryEntry.md) |  | [optional] 
**Type** | Pointer to **string** | Type of Intersight alert. An alert in Intersight could be one of several kinds (FieldNotice, PSIRT etc.). Primarily used for filtering alerts based on the type. * &#x60;restApi&#x60; - Repesents the use of REST API for carrying out alert actions. | [optional] [default to "restApi"]

## Methods

### NewTamAction

`func NewTamAction(classId string, objectType string, ) *TamAction`

NewTamAction instantiates a new TamAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTamActionWithDefaults

`func NewTamActionWithDefaults() *TamAction`

NewTamActionWithDefaults instantiates a new TamAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *TamAction) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *TamAction) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *TamAction) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *TamAction) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *TamAction) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *TamAction) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAffectedObjectType

`func (o *TamAction) GetAffectedObjectType() string`

GetAffectedObjectType returns the AffectedObjectType field if non-nil, zero value otherwise.

### GetAffectedObjectTypeOk

`func (o *TamAction) GetAffectedObjectTypeOk() (*string, bool)`

GetAffectedObjectTypeOk returns a tuple with the AffectedObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffectedObjectType

`func (o *TamAction) SetAffectedObjectType(v string)`

SetAffectedObjectType sets AffectedObjectType field to given value.

### HasAffectedObjectType

`func (o *TamAction) HasAffectedObjectType() bool`

HasAffectedObjectType returns a boolean if a field has been set.

### GetAlertType

`func (o *TamAction) GetAlertType() string`

GetAlertType returns the AlertType field if non-nil, zero value otherwise.

### GetAlertTypeOk

`func (o *TamAction) GetAlertTypeOk() (*string, bool)`

GetAlertTypeOk returns a tuple with the AlertType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertType

`func (o *TamAction) SetAlertType(v string)`

SetAlertType sets AlertType field to given value.

### HasAlertType

`func (o *TamAction) HasAlertType() bool`

HasAlertType returns a boolean if a field has been set.

### GetIdentifiers

`func (o *TamAction) GetIdentifiers() []TamIdentifiers`

GetIdentifiers returns the Identifiers field if non-nil, zero value otherwise.

### GetIdentifiersOk

`func (o *TamAction) GetIdentifiersOk() (*[]TamIdentifiers, bool)`

GetIdentifiersOk returns a tuple with the Identifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifiers

`func (o *TamAction) SetIdentifiers(v []TamIdentifiers)`

SetIdentifiers sets Identifiers field to given value.

### HasIdentifiers

`func (o *TamAction) HasIdentifiers() bool`

HasIdentifiers returns a boolean if a field has been set.

### SetIdentifiersNil

`func (o *TamAction) SetIdentifiersNil(b bool)`

 SetIdentifiersNil sets the value for Identifiers to be an explicit nil

### UnsetIdentifiers
`func (o *TamAction) UnsetIdentifiers()`

UnsetIdentifiers ensures that no value is present for Identifiers, not even an explicit nil
### GetName

`func (o *TamAction) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TamAction) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TamAction) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TamAction) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOperationType

`func (o *TamAction) GetOperationType() string`

GetOperationType returns the OperationType field if non-nil, zero value otherwise.

### GetOperationTypeOk

`func (o *TamAction) GetOperationTypeOk() (*string, bool)`

GetOperationTypeOk returns a tuple with the OperationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationType

`func (o *TamAction) SetOperationType(v string)`

SetOperationType sets OperationType field to given value.

### HasOperationType

`func (o *TamAction) HasOperationType() bool`

HasOperationType returns a boolean if a field has been set.

### GetQueries

`func (o *TamAction) GetQueries() []TamQueryEntry`

GetQueries returns the Queries field if non-nil, zero value otherwise.

### GetQueriesOk

`func (o *TamAction) GetQueriesOk() (*[]TamQueryEntry, bool)`

GetQueriesOk returns a tuple with the Queries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueries

`func (o *TamAction) SetQueries(v []TamQueryEntry)`

SetQueries sets Queries field to given value.

### HasQueries

`func (o *TamAction) HasQueries() bool`

HasQueries returns a boolean if a field has been set.

### SetQueriesNil

`func (o *TamAction) SetQueriesNil(b bool)`

 SetQueriesNil sets the value for Queries to be an explicit nil

### UnsetQueries
`func (o *TamAction) UnsetQueries()`

UnsetQueries ensures that no value is present for Queries, not even an explicit nil
### GetType

`func (o *TamAction) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TamAction) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TamAction) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *TamAction) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


