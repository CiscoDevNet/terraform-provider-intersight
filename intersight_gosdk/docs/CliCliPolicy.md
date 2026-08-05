# CliCliPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "cli.CliPolicy"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "cli.CliPolicy"]
**FiAndUnifiedEdgeEcmcCliLaunch** | Pointer to **bool** | Enable or Disable launching FI and Unified Edge eCMC CLI within the account. | [optional] [default to false]
**Account** | Pointer to [**NullableIamAccountRelationship**](IamAccountRelationship.md) |  | [optional] 

## Methods

### NewCliCliPolicy

`func NewCliCliPolicy(classId string, objectType string, ) *CliCliPolicy`

NewCliCliPolicy instantiates a new CliCliPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCliCliPolicyWithDefaults

`func NewCliCliPolicyWithDefaults() *CliCliPolicy`

NewCliCliPolicyWithDefaults instantiates a new CliCliPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CliCliPolicy) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CliCliPolicy) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CliCliPolicy) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CliCliPolicy) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CliCliPolicy) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CliCliPolicy) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetFiAndUnifiedEdgeEcmcCliLaunch

`func (o *CliCliPolicy) GetFiAndUnifiedEdgeEcmcCliLaunch() bool`

GetFiAndUnifiedEdgeEcmcCliLaunch returns the FiAndUnifiedEdgeEcmcCliLaunch field if non-nil, zero value otherwise.

### GetFiAndUnifiedEdgeEcmcCliLaunchOk

`func (o *CliCliPolicy) GetFiAndUnifiedEdgeEcmcCliLaunchOk() (*bool, bool)`

GetFiAndUnifiedEdgeEcmcCliLaunchOk returns a tuple with the FiAndUnifiedEdgeEcmcCliLaunch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiAndUnifiedEdgeEcmcCliLaunch

`func (o *CliCliPolicy) SetFiAndUnifiedEdgeEcmcCliLaunch(v bool)`

SetFiAndUnifiedEdgeEcmcCliLaunch sets FiAndUnifiedEdgeEcmcCliLaunch field to given value.

### HasFiAndUnifiedEdgeEcmcCliLaunch

`func (o *CliCliPolicy) HasFiAndUnifiedEdgeEcmcCliLaunch() bool`

HasFiAndUnifiedEdgeEcmcCliLaunch returns a boolean if a field has been set.

### GetAccount

`func (o *CliCliPolicy) GetAccount() IamAccountRelationship`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *CliCliPolicy) GetAccountOk() (*IamAccountRelationship, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *CliCliPolicy) SetAccount(v IamAccountRelationship)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *CliCliPolicy) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### SetAccountNil

`func (o *CliCliPolicy) SetAccountNil(b bool)`

 SetAccountNil sets the value for Account to be an explicit nil

### UnsetAccount
`func (o *CliCliPolicy) UnsetAccount()`

UnsetAccount ensures that no value is present for Account, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


