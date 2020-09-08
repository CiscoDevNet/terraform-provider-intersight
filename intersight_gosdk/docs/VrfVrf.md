# VrfVrf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Description to help identify or describe this VRF. | [optional] 
**Name** | Pointer to **string** | Name of the Virtual Routing and Forwarding Instance. | [optional] 
**Account** | Pointer to [**IamAccountRelationship**](iam.Account.Relationship.md) |  | [optional] 

## Methods

### NewVrfVrf

`func NewVrfVrf() *VrfVrf`

NewVrfVrf instantiates a new VrfVrf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVrfVrfWithDefaults

`func NewVrfVrfWithDefaults() *VrfVrf`

NewVrfVrfWithDefaults instantiates a new VrfVrf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *VrfVrf) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VrfVrf) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VrfVrf) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VrfVrf) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *VrfVrf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VrfVrf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VrfVrf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VrfVrf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAccount

`func (o *VrfVrf) GetAccount() IamAccountRelationship`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *VrfVrf) GetAccountOk() (*IamAccountRelationship, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *VrfVrf) SetAccount(v IamAccountRelationship)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *VrfVrf) HasAccount() bool`

HasAccount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


