# MoTag

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AncestorDefinitions** | Pointer to [**[]MoMoRef**](MoMoRef.md) |  | [optional] 
**Definition** | Pointer to [**MoMoRef**](MoMoRef.md) |  | [optional] 
**Key** | Pointer to **string** | The string representation of a tag key. | [optional] 
**Propagated** | Pointer to **bool** | Propagated is a boolean flag that indicates whether the tag is propagated to the related managed objects. | [optional] [readonly] 
**SysTag** | Pointer to **bool** | Specifies whether the tag is user-defined or owned by the system. | [optional] [readonly] 
**Type** | Pointer to **string** | An enum type that defines the type of tag. Supported values are &#39;pathtag&#39; and &#39;keyvalue&#39;. * &#x60;KeyValue&#x60; - KeyValue type of tag. Key is required for these tags. Value is optional. * &#x60;PathTag&#x60; - Key contain path information. Value is not present for these tags. The path is created by using the &#39;/&#39; character as a delimiter.For example, if the tag is \&quot;A/B/C\&quot;, then \&quot;A\&quot; is the parent tag, \&quot;B\&quot; is the child tag of \&quot;A\&quot; and \&quot;C\&quot; is the child tag of \&quot;B\&quot;. | [optional] [readonly] [default to "KeyValue"]
**Value** | Pointer to **string** | The string representation of a tag value. | [optional] 

## Methods

### NewMoTag

`func NewMoTag() *MoTag`

NewMoTag instantiates a new MoTag object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMoTagWithDefaults

`func NewMoTagWithDefaults() *MoTag`

NewMoTagWithDefaults instantiates a new MoTag object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAncestorDefinitions

`func (o *MoTag) GetAncestorDefinitions() []MoMoRef`

GetAncestorDefinitions returns the AncestorDefinitions field if non-nil, zero value otherwise.

### GetAncestorDefinitionsOk

`func (o *MoTag) GetAncestorDefinitionsOk() (*[]MoMoRef, bool)`

GetAncestorDefinitionsOk returns a tuple with the AncestorDefinitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAncestorDefinitions

`func (o *MoTag) SetAncestorDefinitions(v []MoMoRef)`

SetAncestorDefinitions sets AncestorDefinitions field to given value.

### HasAncestorDefinitions

`func (o *MoTag) HasAncestorDefinitions() bool`

HasAncestorDefinitions returns a boolean if a field has been set.

### SetAncestorDefinitionsNil

`func (o *MoTag) SetAncestorDefinitionsNil(b bool)`

 SetAncestorDefinitionsNil sets the value for AncestorDefinitions to be an explicit nil

### UnsetAncestorDefinitions
`func (o *MoTag) UnsetAncestorDefinitions()`

UnsetAncestorDefinitions ensures that no value is present for AncestorDefinitions, not even an explicit nil
### GetDefinition

`func (o *MoTag) GetDefinition() MoMoRef`

GetDefinition returns the Definition field if non-nil, zero value otherwise.

### GetDefinitionOk

`func (o *MoTag) GetDefinitionOk() (*MoMoRef, bool)`

GetDefinitionOk returns a tuple with the Definition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinition

`func (o *MoTag) SetDefinition(v MoMoRef)`

SetDefinition sets Definition field to given value.

### HasDefinition

`func (o *MoTag) HasDefinition() bool`

HasDefinition returns a boolean if a field has been set.

### GetKey

`func (o *MoTag) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *MoTag) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *MoTag) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *MoTag) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetPropagated

`func (o *MoTag) GetPropagated() bool`

GetPropagated returns the Propagated field if non-nil, zero value otherwise.

### GetPropagatedOk

`func (o *MoTag) GetPropagatedOk() (*bool, bool)`

GetPropagatedOk returns a tuple with the Propagated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropagated

`func (o *MoTag) SetPropagated(v bool)`

SetPropagated sets Propagated field to given value.

### HasPropagated

`func (o *MoTag) HasPropagated() bool`

HasPropagated returns a boolean if a field has been set.

### GetSysTag

`func (o *MoTag) GetSysTag() bool`

GetSysTag returns the SysTag field if non-nil, zero value otherwise.

### GetSysTagOk

`func (o *MoTag) GetSysTagOk() (*bool, bool)`

GetSysTagOk returns a tuple with the SysTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSysTag

`func (o *MoTag) SetSysTag(v bool)`

SetSysTag sets SysTag field to given value.

### HasSysTag

`func (o *MoTag) HasSysTag() bool`

HasSysTag returns a boolean if a field has been set.

### GetType

`func (o *MoTag) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MoTag) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MoTag) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *MoTag) HasType() bool`

HasType returns a boolean if a field has been set.

### GetValue

`func (o *MoTag) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *MoTag) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *MoTag) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *MoTag) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


