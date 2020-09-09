# ResourceMembershipHolderAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of this resource membership holder. | [optional] [readonly] 
**Account** | Pointer to [**IamAccountRelationship**](iam.Account.Relationship.md) |  | [optional] 

## Methods

### NewResourceMembershipHolderAllOf

`func NewResourceMembershipHolderAllOf() *ResourceMembershipHolderAllOf`

NewResourceMembershipHolderAllOf instantiates a new ResourceMembershipHolderAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceMembershipHolderAllOfWithDefaults

`func NewResourceMembershipHolderAllOfWithDefaults() *ResourceMembershipHolderAllOf`

NewResourceMembershipHolderAllOfWithDefaults instantiates a new ResourceMembershipHolderAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ResourceMembershipHolderAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ResourceMembershipHolderAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ResourceMembershipHolderAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ResourceMembershipHolderAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAccount

`func (o *ResourceMembershipHolderAllOf) GetAccount() IamAccountRelationship`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *ResourceMembershipHolderAllOf) GetAccountOk() (*IamAccountRelationship, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *ResourceMembershipHolderAllOf) SetAccount(v IamAccountRelationship)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *ResourceMembershipHolderAllOf) HasAccount() bool`

HasAccount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


