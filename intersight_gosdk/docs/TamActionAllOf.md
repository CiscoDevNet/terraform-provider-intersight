# TamActionAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AffectedObjectType** | Pointer to **string** | Type of the managed object that should be marked with an instance of the Alert (when operation type is create) or that should have an alert instance removed (when operation type is remove). | [optional] 
**AlertType** | Pointer to **string** | Alert type is used to denote the category of an Intersight alert (FieldNotice, equipment Fault etc.). * &#x60;psirt&#x60; - Respresents the psirt alert type (https://tools.cisco.com/security/center/publicationListing.x). * &#x60;fieldNotice&#x60; - Respresents the field notice alert type (https://www.cisco.com/c/en/us/support/web/tsd-products-field-notice-summary.html). | [optional] [default to "psirt"]
**Identifiers** | Pointer to [**[]TamIdentifiers**](tam.Identifiers.md) |  | [optional] 
**Name** | Pointer to **string** | Uniquely identifies a given action among the set of actions corresponding to an advisory. Primarily used to store and compare results of subsequent iterations corresponding to the action queries. | [optional] 
**OperationType** | Pointer to **string** | Operation type for the alert action. An action is used to carry out the process of \&quot;reacting\&quot; to an alert condition. For e.g.in case of a fieldNotice alert, the intention may be to create a new alert (if the condition matches and there is no existing alert) or to remove an existing alert when the alert condition has been remedied. * &#x60;create&#x60; - Create an instance of AdvisoryInstance. * &#x60;remove&#x60; - Remove an instance of AdvisoryInstance. | [optional] [default to "create"]
**Queries** | Pointer to [**[]TamQueryEntry**](tam.QueryEntry.md) |  | [optional] 
**Type** | Pointer to **string** | Type of Intersight alert. An alert in Intersight could be one of several kinds (FieldNotice, PSIRT etc.). Primarily used for filtering alerts based on the type. * &#x60;restApi&#x60; - Repesents the use of REST API for carrying out alert actions. | [optional] [default to "restApi"]

## Methods

### NewTamActionAllOf

`func NewTamActionAllOf() *TamActionAllOf`

NewTamActionAllOf instantiates a new TamActionAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTamActionAllOfWithDefaults

`func NewTamActionAllOfWithDefaults() *TamActionAllOf`

NewTamActionAllOfWithDefaults instantiates a new TamActionAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAffectedObjectType

`func (o *TamActionAllOf) GetAffectedObjectType() string`

GetAffectedObjectType returns the AffectedObjectType field if non-nil, zero value otherwise.

### GetAffectedObjectTypeOk

`func (o *TamActionAllOf) GetAffectedObjectTypeOk() (*string, bool)`

GetAffectedObjectTypeOk returns a tuple with the AffectedObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffectedObjectType

`func (o *TamActionAllOf) SetAffectedObjectType(v string)`

SetAffectedObjectType sets AffectedObjectType field to given value.

### HasAffectedObjectType

`func (o *TamActionAllOf) HasAffectedObjectType() bool`

HasAffectedObjectType returns a boolean if a field has been set.

### GetAlertType

`func (o *TamActionAllOf) GetAlertType() string`

GetAlertType returns the AlertType field if non-nil, zero value otherwise.

### GetAlertTypeOk

`func (o *TamActionAllOf) GetAlertTypeOk() (*string, bool)`

GetAlertTypeOk returns a tuple with the AlertType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertType

`func (o *TamActionAllOf) SetAlertType(v string)`

SetAlertType sets AlertType field to given value.

### HasAlertType

`func (o *TamActionAllOf) HasAlertType() bool`

HasAlertType returns a boolean if a field has been set.

### GetIdentifiers

`func (o *TamActionAllOf) GetIdentifiers() []TamIdentifiers`

GetIdentifiers returns the Identifiers field if non-nil, zero value otherwise.

### GetIdentifiersOk

`func (o *TamActionAllOf) GetIdentifiersOk() (*[]TamIdentifiers, bool)`

GetIdentifiersOk returns a tuple with the Identifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifiers

`func (o *TamActionAllOf) SetIdentifiers(v []TamIdentifiers)`

SetIdentifiers sets Identifiers field to given value.

### HasIdentifiers

`func (o *TamActionAllOf) HasIdentifiers() bool`

HasIdentifiers returns a boolean if a field has been set.

### GetName

`func (o *TamActionAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TamActionAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TamActionAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TamActionAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOperationType

`func (o *TamActionAllOf) GetOperationType() string`

GetOperationType returns the OperationType field if non-nil, zero value otherwise.

### GetOperationTypeOk

`func (o *TamActionAllOf) GetOperationTypeOk() (*string, bool)`

GetOperationTypeOk returns a tuple with the OperationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationType

`func (o *TamActionAllOf) SetOperationType(v string)`

SetOperationType sets OperationType field to given value.

### HasOperationType

`func (o *TamActionAllOf) HasOperationType() bool`

HasOperationType returns a boolean if a field has been set.

### GetQueries

`func (o *TamActionAllOf) GetQueries() []TamQueryEntry`

GetQueries returns the Queries field if non-nil, zero value otherwise.

### GetQueriesOk

`func (o *TamActionAllOf) GetQueriesOk() (*[]TamQueryEntry, bool)`

GetQueriesOk returns a tuple with the Queries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueries

`func (o *TamActionAllOf) SetQueries(v []TamQueryEntry)`

SetQueries sets Queries field to given value.

### HasQueries

`func (o *TamActionAllOf) HasQueries() bool`

HasQueries returns a boolean if a field has been set.

### GetType

`func (o *TamActionAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TamActionAllOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TamActionAllOf) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *TamActionAllOf) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


