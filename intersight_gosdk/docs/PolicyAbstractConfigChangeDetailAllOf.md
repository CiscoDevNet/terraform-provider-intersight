# PolicyAbstractConfigChangeDetailAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**Changes** | Pointer to **[]string** |  | [optional] 
**ConfigChangeContext** | Pointer to [**NullablePolicyConfigResultContext**](PolicyConfigResultContext.md) |  | [optional] 
**ConfigChangeFlag** | Pointer to **string** | Config change flag to differentiate Pending-changes and Config-drift. * &#x60;Pending-changes&#x60; - Config change flag represents changes are due to not deployed changes from Intersight. * &#x60;Drift-changes&#x60; - Config change flag represents changes are due to endpoint configuration changes. | [optional] [default to "Pending-changes"]
**Disruptions** | Pointer to **[]string** |  | [optional] 
**Message** | Pointer to **string** | Detailed description of the config change. | [optional] 
**ModStatus** | Pointer to **string** | Modification status of the mo in this config change. * &#x60;None&#x60; - The &#39;none&#39; operation/state.Indicates a configuration profile has been deployed, and the desired configuration matches the actual device configuration. * &#x60;Created&#x60; - The &#39;create&#39; operation/state.Indicates a configuration profile has been created and associated with a device, but the configuration specified in the profilehas not been applied yet. For example, this could happen when the user creates a server profile and has not deployed the profile yet. * &#x60;Modified&#x60; - The &#39;update&#39; operation/state.Indicates some of the desired configuration changes specified in a profile have not been been applied to the associated device.This happens when the user has made changes to a profile and has not deployed the changes yet, or when the workflow to applythe configuration changes has not completed succesfully. * &#x60;Deleted&#x60; - The &#39;delete&#39; operation/state.Indicates a configuration profile has been been disassociated from a device and the user has not undeployed these changes yet. | [optional] [default to "None"]

## Methods

### NewPolicyAbstractConfigChangeDetailAllOf

`func NewPolicyAbstractConfigChangeDetailAllOf(classId string, objectType string, ) *PolicyAbstractConfigChangeDetailAllOf`

NewPolicyAbstractConfigChangeDetailAllOf instantiates a new PolicyAbstractConfigChangeDetailAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyAbstractConfigChangeDetailAllOfWithDefaults

`func NewPolicyAbstractConfigChangeDetailAllOfWithDefaults() *PolicyAbstractConfigChangeDetailAllOf`

NewPolicyAbstractConfigChangeDetailAllOfWithDefaults instantiates a new PolicyAbstractConfigChangeDetailAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *PolicyAbstractConfigChangeDetailAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *PolicyAbstractConfigChangeDetailAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *PolicyAbstractConfigChangeDetailAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *PolicyAbstractConfigChangeDetailAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *PolicyAbstractConfigChangeDetailAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *PolicyAbstractConfigChangeDetailAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetChanges

`func (o *PolicyAbstractConfigChangeDetailAllOf) GetChanges() []string`

GetChanges returns the Changes field if non-nil, zero value otherwise.

### GetChangesOk

`func (o *PolicyAbstractConfigChangeDetailAllOf) GetChangesOk() (*[]string, bool)`

GetChangesOk returns a tuple with the Changes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChanges

`func (o *PolicyAbstractConfigChangeDetailAllOf) SetChanges(v []string)`

SetChanges sets Changes field to given value.

### HasChanges

`func (o *PolicyAbstractConfigChangeDetailAllOf) HasChanges() bool`

HasChanges returns a boolean if a field has been set.

### SetChangesNil

`func (o *PolicyAbstractConfigChangeDetailAllOf) SetChangesNil(b bool)`

 SetChangesNil sets the value for Changes to be an explicit nil

### UnsetChanges
`func (o *PolicyAbstractConfigChangeDetailAllOf) UnsetChanges()`

UnsetChanges ensures that no value is present for Changes, not even an explicit nil
### GetConfigChangeContext

`func (o *PolicyAbstractConfigChangeDetailAllOf) GetConfigChangeContext() PolicyConfigResultContext`

GetConfigChangeContext returns the ConfigChangeContext field if non-nil, zero value otherwise.

### GetConfigChangeContextOk

`func (o *PolicyAbstractConfigChangeDetailAllOf) GetConfigChangeContextOk() (*PolicyConfigResultContext, bool)`

GetConfigChangeContextOk returns a tuple with the ConfigChangeContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigChangeContext

`func (o *PolicyAbstractConfigChangeDetailAllOf) SetConfigChangeContext(v PolicyConfigResultContext)`

SetConfigChangeContext sets ConfigChangeContext field to given value.

### HasConfigChangeContext

`func (o *PolicyAbstractConfigChangeDetailAllOf) HasConfigChangeContext() bool`

HasConfigChangeContext returns a boolean if a field has been set.

### SetConfigChangeContextNil

`func (o *PolicyAbstractConfigChangeDetailAllOf) SetConfigChangeContextNil(b bool)`

 SetConfigChangeContextNil sets the value for ConfigChangeContext to be an explicit nil

### UnsetConfigChangeContext
`func (o *PolicyAbstractConfigChangeDetailAllOf) UnsetConfigChangeContext()`

UnsetConfigChangeContext ensures that no value is present for ConfigChangeContext, not even an explicit nil
### GetConfigChangeFlag

`func (o *PolicyAbstractConfigChangeDetailAllOf) GetConfigChangeFlag() string`

GetConfigChangeFlag returns the ConfigChangeFlag field if non-nil, zero value otherwise.

### GetConfigChangeFlagOk

`func (o *PolicyAbstractConfigChangeDetailAllOf) GetConfigChangeFlagOk() (*string, bool)`

GetConfigChangeFlagOk returns a tuple with the ConfigChangeFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigChangeFlag

`func (o *PolicyAbstractConfigChangeDetailAllOf) SetConfigChangeFlag(v string)`

SetConfigChangeFlag sets ConfigChangeFlag field to given value.

### HasConfigChangeFlag

`func (o *PolicyAbstractConfigChangeDetailAllOf) HasConfigChangeFlag() bool`

HasConfigChangeFlag returns a boolean if a field has been set.

### GetDisruptions

`func (o *PolicyAbstractConfigChangeDetailAllOf) GetDisruptions() []string`

GetDisruptions returns the Disruptions field if non-nil, zero value otherwise.

### GetDisruptionsOk

`func (o *PolicyAbstractConfigChangeDetailAllOf) GetDisruptionsOk() (*[]string, bool)`

GetDisruptionsOk returns a tuple with the Disruptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisruptions

`func (o *PolicyAbstractConfigChangeDetailAllOf) SetDisruptions(v []string)`

SetDisruptions sets Disruptions field to given value.

### HasDisruptions

`func (o *PolicyAbstractConfigChangeDetailAllOf) HasDisruptions() bool`

HasDisruptions returns a boolean if a field has been set.

### SetDisruptionsNil

`func (o *PolicyAbstractConfigChangeDetailAllOf) SetDisruptionsNil(b bool)`

 SetDisruptionsNil sets the value for Disruptions to be an explicit nil

### UnsetDisruptions
`func (o *PolicyAbstractConfigChangeDetailAllOf) UnsetDisruptions()`

UnsetDisruptions ensures that no value is present for Disruptions, not even an explicit nil
### GetMessage

`func (o *PolicyAbstractConfigChangeDetailAllOf) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *PolicyAbstractConfigChangeDetailAllOf) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *PolicyAbstractConfigChangeDetailAllOf) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *PolicyAbstractConfigChangeDetailAllOf) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetModStatus

`func (o *PolicyAbstractConfigChangeDetailAllOf) GetModStatus() string`

GetModStatus returns the ModStatus field if non-nil, zero value otherwise.

### GetModStatusOk

`func (o *PolicyAbstractConfigChangeDetailAllOf) GetModStatusOk() (*string, bool)`

GetModStatusOk returns a tuple with the ModStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModStatus

`func (o *PolicyAbstractConfigChangeDetailAllOf) SetModStatus(v string)`

SetModStatus sets ModStatus field to given value.

### HasModStatus

`func (o *PolicyAbstractConfigChangeDetailAllOf) HasModStatus() bool`

HasModStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


