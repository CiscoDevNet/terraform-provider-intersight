# PolicyAbstractDisruption

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | [default to "server.Disruption"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | [default to "server.Disruption"]
**Label** | Pointer to **string** | Description of the disruption. | [optional] [readonly] 
**Name** | Pointer to **string** | Disruption that happens when the policy is deployed. | [optional] [readonly] 
**Severity** | Pointer to **string** | Severity of the disruption. * &#x60;Minor&#x60; - A disruption of minor severity. * &#x60;Major&#x60; - A disruption of major severity. * &#x60;Moderate&#x60; - A disruption of moderate severity. | [optional] [readonly] [default to "Minor"]
**Type** | Pointer to **string** | Type of disruption that happens when the policy is deployed. * &#x60;Informational&#x60; - Disruptions categorized as informational do not require any user action, nor do they cause a reboot. * &#x60;ActionRequired&#x60; - User action is required to deploy the profile. | [optional] [readonly] [default to "Informational"]

## Methods

### NewPolicyAbstractDisruption

`func NewPolicyAbstractDisruption(classId string, objectType string, ) *PolicyAbstractDisruption`

NewPolicyAbstractDisruption instantiates a new PolicyAbstractDisruption object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyAbstractDisruptionWithDefaults

`func NewPolicyAbstractDisruptionWithDefaults() *PolicyAbstractDisruption`

NewPolicyAbstractDisruptionWithDefaults instantiates a new PolicyAbstractDisruption object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *PolicyAbstractDisruption) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *PolicyAbstractDisruption) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *PolicyAbstractDisruption) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *PolicyAbstractDisruption) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *PolicyAbstractDisruption) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *PolicyAbstractDisruption) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetLabel

`func (o *PolicyAbstractDisruption) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *PolicyAbstractDisruption) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *PolicyAbstractDisruption) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *PolicyAbstractDisruption) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetName

`func (o *PolicyAbstractDisruption) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PolicyAbstractDisruption) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PolicyAbstractDisruption) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PolicyAbstractDisruption) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSeverity

`func (o *PolicyAbstractDisruption) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *PolicyAbstractDisruption) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *PolicyAbstractDisruption) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *PolicyAbstractDisruption) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetType

`func (o *PolicyAbstractDisruption) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PolicyAbstractDisruption) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PolicyAbstractDisruption) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PolicyAbstractDisruption) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


