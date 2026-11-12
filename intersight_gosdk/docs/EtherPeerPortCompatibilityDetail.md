# EtherPeerPortCompatibilityDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "ether.PeerPortCompatibilityDetail"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "ether.PeerPortCompatibilityDetail"]
**Active** | Pointer to **bool** | Indicates whether this compatibility issue is currently active. | [optional] [readonly] 
**IssueType** | Pointer to **string** | The type of compatibility issue detected. * &#x60;&#x60; - Default value for compatibility issue type. * &#x60;Mode&#x60; - The mode of this host port does not match the mode of the connected peer device port. * &#x60;Native VLAN&#x60; - The native VLAN configuration of the port does not match the connected peer device port. * &#x60;Access VLAN&#x60; - The access VLAN configuration of the port does not match the connected peer device port. * &#x60;Allowed VLANs&#x60; - The allowed VLAN configuration of the port does not match the connected peer device port. * &#x60;Auto Negotiation&#x60; - The auto-negotiation configuration of the port does not match the connected peer device port. * &#x60;MTU&#x60; - The MTU configuration of the port does not match the connected peer device port. * &#x60;FEC&#x60; - The FEC configuration of the port does not match the connected peer device port. * &#x60;Port Speed&#x60; - The Port speed configuration of the port does not match the connected peer device port. * &#x60;Port Channel ID&#x60; - The Port Channel ID configuration of the port does not match the connected peer device port. | [optional] [readonly] [default to ""]
**LocalValue** | Pointer to **string** | The value configured on the local host port. | [optional] [readonly] 
**LocalValueModTime** | Pointer to **string** | The last modification time of the local value on the host port. | [optional] [readonly] 
**PeerValue** | Pointer to **string** | The value configured on the peer Catalyst SD-WAN device port. | [optional] [readonly] 
**PeerValueModTime** | Pointer to **string** | The last modification time of the peer value on the Catalyst SD-WAN device port. | [optional] [readonly] 

## Methods

### NewEtherPeerPortCompatibilityDetail

`func NewEtherPeerPortCompatibilityDetail(classId string, objectType string, ) *EtherPeerPortCompatibilityDetail`

NewEtherPeerPortCompatibilityDetail instantiates a new EtherPeerPortCompatibilityDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEtherPeerPortCompatibilityDetailWithDefaults

`func NewEtherPeerPortCompatibilityDetailWithDefaults() *EtherPeerPortCompatibilityDetail`

NewEtherPeerPortCompatibilityDetailWithDefaults instantiates a new EtherPeerPortCompatibilityDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *EtherPeerPortCompatibilityDetail) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *EtherPeerPortCompatibilityDetail) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *EtherPeerPortCompatibilityDetail) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *EtherPeerPortCompatibilityDetail) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *EtherPeerPortCompatibilityDetail) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *EtherPeerPortCompatibilityDetail) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetActive

`func (o *EtherPeerPortCompatibilityDetail) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *EtherPeerPortCompatibilityDetail) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *EtherPeerPortCompatibilityDetail) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *EtherPeerPortCompatibilityDetail) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetIssueType

`func (o *EtherPeerPortCompatibilityDetail) GetIssueType() string`

GetIssueType returns the IssueType field if non-nil, zero value otherwise.

### GetIssueTypeOk

`func (o *EtherPeerPortCompatibilityDetail) GetIssueTypeOk() (*string, bool)`

GetIssueTypeOk returns a tuple with the IssueType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueType

`func (o *EtherPeerPortCompatibilityDetail) SetIssueType(v string)`

SetIssueType sets IssueType field to given value.

### HasIssueType

`func (o *EtherPeerPortCompatibilityDetail) HasIssueType() bool`

HasIssueType returns a boolean if a field has been set.

### GetLocalValue

`func (o *EtherPeerPortCompatibilityDetail) GetLocalValue() string`

GetLocalValue returns the LocalValue field if non-nil, zero value otherwise.

### GetLocalValueOk

`func (o *EtherPeerPortCompatibilityDetail) GetLocalValueOk() (*string, bool)`

GetLocalValueOk returns a tuple with the LocalValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalValue

`func (o *EtherPeerPortCompatibilityDetail) SetLocalValue(v string)`

SetLocalValue sets LocalValue field to given value.

### HasLocalValue

`func (o *EtherPeerPortCompatibilityDetail) HasLocalValue() bool`

HasLocalValue returns a boolean if a field has been set.

### GetLocalValueModTime

`func (o *EtherPeerPortCompatibilityDetail) GetLocalValueModTime() string`

GetLocalValueModTime returns the LocalValueModTime field if non-nil, zero value otherwise.

### GetLocalValueModTimeOk

`func (o *EtherPeerPortCompatibilityDetail) GetLocalValueModTimeOk() (*string, bool)`

GetLocalValueModTimeOk returns a tuple with the LocalValueModTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalValueModTime

`func (o *EtherPeerPortCompatibilityDetail) SetLocalValueModTime(v string)`

SetLocalValueModTime sets LocalValueModTime field to given value.

### HasLocalValueModTime

`func (o *EtherPeerPortCompatibilityDetail) HasLocalValueModTime() bool`

HasLocalValueModTime returns a boolean if a field has been set.

### GetPeerValue

`func (o *EtherPeerPortCompatibilityDetail) GetPeerValue() string`

GetPeerValue returns the PeerValue field if non-nil, zero value otherwise.

### GetPeerValueOk

`func (o *EtherPeerPortCompatibilityDetail) GetPeerValueOk() (*string, bool)`

GetPeerValueOk returns a tuple with the PeerValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerValue

`func (o *EtherPeerPortCompatibilityDetail) SetPeerValue(v string)`

SetPeerValue sets PeerValue field to given value.

### HasPeerValue

`func (o *EtherPeerPortCompatibilityDetail) HasPeerValue() bool`

HasPeerValue returns a boolean if a field has been set.

### GetPeerValueModTime

`func (o *EtherPeerPortCompatibilityDetail) GetPeerValueModTime() string`

GetPeerValueModTime returns the PeerValueModTime field if non-nil, zero value otherwise.

### GetPeerValueModTimeOk

`func (o *EtherPeerPortCompatibilityDetail) GetPeerValueModTimeOk() (*string, bool)`

GetPeerValueModTimeOk returns a tuple with the PeerValueModTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerValueModTime

`func (o *EtherPeerPortCompatibilityDetail) SetPeerValueModTime(v string)`

SetPeerValueModTime sets PeerValueModTime field to given value.

### HasPeerValueModTime

`func (o *EtherPeerPortCompatibilityDetail) HasPeerValueModTime() bool`

HasPeerValueModTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


