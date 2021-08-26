# IamSecurityHolderAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "iam.SecurityHolder"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "iam.SecurityHolder"]
**Account** | Pointer to [**IamAccountRelationship**](IamAccountRelationship.md) |  | [optional] 
**IpRulesConfiguration** | Pointer to [**IamIpAccessManagementRelationship**](IamIpAccessManagementRelationship.md) |  | [optional] 
**ResourcePermissions** | Pointer to [**[]IamResourcePermissionRelationship**](IamResourcePermissionRelationship.md) | An array of relationships to iamResourcePermission resources. | [optional] [readonly] 

## Methods

### NewIamSecurityHolderAllOf

`func NewIamSecurityHolderAllOf(classId string, objectType string, ) *IamSecurityHolderAllOf`

NewIamSecurityHolderAllOf instantiates a new IamSecurityHolderAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamSecurityHolderAllOfWithDefaults

`func NewIamSecurityHolderAllOfWithDefaults() *IamSecurityHolderAllOf`

NewIamSecurityHolderAllOfWithDefaults instantiates a new IamSecurityHolderAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *IamSecurityHolderAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *IamSecurityHolderAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *IamSecurityHolderAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *IamSecurityHolderAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *IamSecurityHolderAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *IamSecurityHolderAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAccount

`func (o *IamSecurityHolderAllOf) GetAccount() IamAccountRelationship`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *IamSecurityHolderAllOf) GetAccountOk() (*IamAccountRelationship, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *IamSecurityHolderAllOf) SetAccount(v IamAccountRelationship)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *IamSecurityHolderAllOf) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetIpRulesConfiguration

`func (o *IamSecurityHolderAllOf) GetIpRulesConfiguration() IamIpAccessManagementRelationship`

GetIpRulesConfiguration returns the IpRulesConfiguration field if non-nil, zero value otherwise.

### GetIpRulesConfigurationOk

`func (o *IamSecurityHolderAllOf) GetIpRulesConfigurationOk() (*IamIpAccessManagementRelationship, bool)`

GetIpRulesConfigurationOk returns a tuple with the IpRulesConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpRulesConfiguration

`func (o *IamSecurityHolderAllOf) SetIpRulesConfiguration(v IamIpAccessManagementRelationship)`

SetIpRulesConfiguration sets IpRulesConfiguration field to given value.

### HasIpRulesConfiguration

`func (o *IamSecurityHolderAllOf) HasIpRulesConfiguration() bool`

HasIpRulesConfiguration returns a boolean if a field has been set.

### GetResourcePermissions

`func (o *IamSecurityHolderAllOf) GetResourcePermissions() []IamResourcePermissionRelationship`

GetResourcePermissions returns the ResourcePermissions field if non-nil, zero value otherwise.

### GetResourcePermissionsOk

`func (o *IamSecurityHolderAllOf) GetResourcePermissionsOk() (*[]IamResourcePermissionRelationship, bool)`

GetResourcePermissionsOk returns a tuple with the ResourcePermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePermissions

`func (o *IamSecurityHolderAllOf) SetResourcePermissions(v []IamResourcePermissionRelationship)`

SetResourcePermissions sets ResourcePermissions field to given value.

### HasResourcePermissions

`func (o *IamSecurityHolderAllOf) HasResourcePermissions() bool`

HasResourcePermissions returns a boolean if a field has been set.

### SetResourcePermissionsNil

`func (o *IamSecurityHolderAllOf) SetResourcePermissionsNil(b bool)`

 SetResourcePermissionsNil sets the value for ResourcePermissions to be an explicit nil

### UnsetResourcePermissions
`func (o *IamSecurityHolderAllOf) UnsetResourcePermissions()`

UnsetResourcePermissions ensures that no value is present for ResourcePermissions, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


