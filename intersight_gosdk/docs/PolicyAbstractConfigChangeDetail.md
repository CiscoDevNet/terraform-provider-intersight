# PolicyAbstractConfigChangeDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Changes** | Pointer to **[]string** |  | [optional] 
**ConfigChangeContext** | Pointer to [**PolicyConfigResultContext**](policy.ConfigResultContext.md) |  | [optional] 
**ConfigChangeFlag** | Pointer to **string** | Config change flag to differentiate Pending-changes and Config-drift. * &#x60;Pending-changes&#x60; - Config change flag represents changes are due to not deployed changes from Intersight. * &#x60;Drift-changes&#x60; - Config change flag represents changes are due to endpoint configuration changes. | [optional] [default to "Pending-changes"]
**Disruptions** | Pointer to **[]string** |  | [optional] 
**Message** | Pointer to **string** | Detailed description of the config change. | [optional] 
**ModStatus** | Pointer to **string** | Modification status of the mo in this config change. * &#x60;None&#x60; - The &#39;none&#39; operation/state.Indicates a configuration profile has been deployed, and the desired configuration matches the actual device configuration. * &#x60;Created&#x60; - The &#39;create&#39; operation/state.Indicates a configuration profile has been created and associated with a device, but the configuration specified in the profilehas not been applied yet. For example, this could happen when the user creates a server profile and has not deployed the profile yet. * &#x60;Modified&#x60; - The &#39;update&#39; operation/state.Indicates some of the desired configuration changes specified in a profile have not been been applied to the associated device.This happens when the user has made changes to a profile and has not deployed the changes yet, or when the workflow to applythe configuration changes has not completed succesfully. * &#x60;Deleted&#x60; - The &#39;delete&#39; operation/state.Indicates a configuration profile has been been disassociated from a device and the user has not undeployed these changes yet. | [optional] [default to "None"]

## Methods

### NewPolicyAbstractConfigChangeDetail

`func NewPolicyAbstractConfigChangeDetail() *PolicyAbstractConfigChangeDetail`

NewPolicyAbstractConfigChangeDetail instantiates a new PolicyAbstractConfigChangeDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyAbstractConfigChangeDetailWithDefaults

`func NewPolicyAbstractConfigChangeDetailWithDefaults() *PolicyAbstractConfigChangeDetail`

NewPolicyAbstractConfigChangeDetailWithDefaults instantiates a new PolicyAbstractConfigChangeDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChanges

`func (o *PolicyAbstractConfigChangeDetail) GetChanges() []string`

GetChanges returns the Changes field if non-nil, zero value otherwise.

### GetChangesOk

`func (o *PolicyAbstractConfigChangeDetail) GetChangesOk() (*[]string, bool)`

GetChangesOk returns a tuple with the Changes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChanges

`func (o *PolicyAbstractConfigChangeDetail) SetChanges(v []string)`

SetChanges sets Changes field to given value.

### HasChanges

`func (o *PolicyAbstractConfigChangeDetail) HasChanges() bool`

HasChanges returns a boolean if a field has been set.

### GetConfigChangeContext

`func (o *PolicyAbstractConfigChangeDetail) GetConfigChangeContext() PolicyConfigResultContext`

GetConfigChangeContext returns the ConfigChangeContext field if non-nil, zero value otherwise.

### GetConfigChangeContextOk

`func (o *PolicyAbstractConfigChangeDetail) GetConfigChangeContextOk() (*PolicyConfigResultContext, bool)`

GetConfigChangeContextOk returns a tuple with the ConfigChangeContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigChangeContext

`func (o *PolicyAbstractConfigChangeDetail) SetConfigChangeContext(v PolicyConfigResultContext)`

SetConfigChangeContext sets ConfigChangeContext field to given value.

### HasConfigChangeContext

`func (o *PolicyAbstractConfigChangeDetail) HasConfigChangeContext() bool`

HasConfigChangeContext returns a boolean if a field has been set.

### GetConfigChangeFlag

`func (o *PolicyAbstractConfigChangeDetail) GetConfigChangeFlag() string`

GetConfigChangeFlag returns the ConfigChangeFlag field if non-nil, zero value otherwise.

### GetConfigChangeFlagOk

`func (o *PolicyAbstractConfigChangeDetail) GetConfigChangeFlagOk() (*string, bool)`

GetConfigChangeFlagOk returns a tuple with the ConfigChangeFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigChangeFlag

`func (o *PolicyAbstractConfigChangeDetail) SetConfigChangeFlag(v string)`

SetConfigChangeFlag sets ConfigChangeFlag field to given value.

### HasConfigChangeFlag

`func (o *PolicyAbstractConfigChangeDetail) HasConfigChangeFlag() bool`

HasConfigChangeFlag returns a boolean if a field has been set.

### GetDisruptions

`func (o *PolicyAbstractConfigChangeDetail) GetDisruptions() []string`

GetDisruptions returns the Disruptions field if non-nil, zero value otherwise.

### GetDisruptionsOk

`func (o *PolicyAbstractConfigChangeDetail) GetDisruptionsOk() (*[]string, bool)`

GetDisruptionsOk returns a tuple with the Disruptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisruptions

`func (o *PolicyAbstractConfigChangeDetail) SetDisruptions(v []string)`

SetDisruptions sets Disruptions field to given value.

### HasDisruptions

`func (o *PolicyAbstractConfigChangeDetail) HasDisruptions() bool`

HasDisruptions returns a boolean if a field has been set.

### GetMessage

`func (o *PolicyAbstractConfigChangeDetail) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *PolicyAbstractConfigChangeDetail) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *PolicyAbstractConfigChangeDetail) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *PolicyAbstractConfigChangeDetail) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetModStatus

`func (o *PolicyAbstractConfigChangeDetail) GetModStatus() string`

GetModStatus returns the ModStatus field if non-nil, zero value otherwise.

### GetModStatusOk

`func (o *PolicyAbstractConfigChangeDetail) GetModStatusOk() (*string, bool)`

GetModStatusOk returns a tuple with the ModStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModStatus

`func (o *PolicyAbstractConfigChangeDetail) SetModStatus(v string)`

SetModStatus sets ModStatus field to given value.

### HasModStatus

`func (o *PolicyAbstractConfigChangeDetail) HasModStatus() bool`

HasModStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


