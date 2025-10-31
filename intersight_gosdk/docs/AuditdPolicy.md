# AuditdPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "auditd.Policy"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "auditd.Policy"]
**AdminState** | Pointer to **string** | Admin state for the AuditD feature. * &#x60;Disabled&#x60; - Admin configured Disabled State. * &#x60;Enabled&#x60; - Admin configured Enabled State. | [optional] [default to "Disabled"]
**AuditdLogLevel** | Pointer to **string** | The log level for the AuditD feature. The default value is \&quot;notifications\&quot;. * &#x60;notifications&#x60; - Generated logs are of Notification level, providing information about normal but significant system events requiring awareness. * &#x60;emergencies&#x60; - Generated logs will be of Emergency log level, indicating a critical and unstable system state. * &#x60;alerts&#x60; - Generated logs are of Alert level, indicating critical issues needing immediate attention to prevent system disruption. * &#x60;critical&#x60; - Generated logs are of Critical level, signaling severe issues that may cause system failure if not addressed immediately. * &#x60;errors&#x60; - Generated logs are of Error level, indicating significant problems that affect functionality but do not cause system failure. * &#x60;warnings&#x60; - Generated logs are of Warning level, highlighting potential issues that require attention but do not yet impact functionality. * &#x60;information&#x60; - Generated logs are of Information level, detailing routine operational messages without indicating any issues or errors. * &#x60;debugging&#x60; - Generated logs are of Debugging level, providing detailed information to help diagnose and troubleshoot system issues. | [optional] [default to "notifications"]
**Organization** | Pointer to [**NullableOrganizationOrganizationRelationship**](OrganizationOrganizationRelationship.md) |  | [optional] 
**Profiles** | Pointer to [**[]FabricBaseSwitchProfileRelationship**](FabricBaseSwitchProfileRelationship.md) | An array of relationships to fabricBaseSwitchProfile resources. | [optional] 

## Methods

### NewAuditdPolicy

`func NewAuditdPolicy(classId string, objectType string, ) *AuditdPolicy`

NewAuditdPolicy instantiates a new AuditdPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuditdPolicyWithDefaults

`func NewAuditdPolicyWithDefaults() *AuditdPolicy`

NewAuditdPolicyWithDefaults instantiates a new AuditdPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *AuditdPolicy) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *AuditdPolicy) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *AuditdPolicy) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *AuditdPolicy) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *AuditdPolicy) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *AuditdPolicy) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAdminState

`func (o *AuditdPolicy) GetAdminState() string`

GetAdminState returns the AdminState field if non-nil, zero value otherwise.

### GetAdminStateOk

`func (o *AuditdPolicy) GetAdminStateOk() (*string, bool)`

GetAdminStateOk returns a tuple with the AdminState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminState

`func (o *AuditdPolicy) SetAdminState(v string)`

SetAdminState sets AdminState field to given value.

### HasAdminState

`func (o *AuditdPolicy) HasAdminState() bool`

HasAdminState returns a boolean if a field has been set.

### GetAuditdLogLevel

`func (o *AuditdPolicy) GetAuditdLogLevel() string`

GetAuditdLogLevel returns the AuditdLogLevel field if non-nil, zero value otherwise.

### GetAuditdLogLevelOk

`func (o *AuditdPolicy) GetAuditdLogLevelOk() (*string, bool)`

GetAuditdLogLevelOk returns a tuple with the AuditdLogLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditdLogLevel

`func (o *AuditdPolicy) SetAuditdLogLevel(v string)`

SetAuditdLogLevel sets AuditdLogLevel field to given value.

### HasAuditdLogLevel

`func (o *AuditdPolicy) HasAuditdLogLevel() bool`

HasAuditdLogLevel returns a boolean if a field has been set.

### GetOrganization

`func (o *AuditdPolicy) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *AuditdPolicy) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *AuditdPolicy) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *AuditdPolicy) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### SetOrganizationNil

`func (o *AuditdPolicy) SetOrganizationNil(b bool)`

 SetOrganizationNil sets the value for Organization to be an explicit nil

### UnsetOrganization
`func (o *AuditdPolicy) UnsetOrganization()`

UnsetOrganization ensures that no value is present for Organization, not even an explicit nil
### GetProfiles

`func (o *AuditdPolicy) GetProfiles() []FabricBaseSwitchProfileRelationship`

GetProfiles returns the Profiles field if non-nil, zero value otherwise.

### GetProfilesOk

`func (o *AuditdPolicy) GetProfilesOk() (*[]FabricBaseSwitchProfileRelationship, bool)`

GetProfilesOk returns a tuple with the Profiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfiles

`func (o *AuditdPolicy) SetProfiles(v []FabricBaseSwitchProfileRelationship)`

SetProfiles sets Profiles field to given value.

### HasProfiles

`func (o *AuditdPolicy) HasProfiles() bool`

HasProfiles returns a boolean if a field has been set.

### SetProfilesNil

`func (o *AuditdPolicy) SetProfilesNil(b bool)`

 SetProfilesNil sets the value for Profiles to be an explicit nil

### UnsetProfiles
`func (o *AuditdPolicy) UnsetProfiles()`

UnsetProfiles ensures that no value is present for Profiles, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


