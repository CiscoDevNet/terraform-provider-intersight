# CloudSecurityGroupRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "cloud.SecurityGroupRule"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "cloud.SecurityGroupRule"]
**Action** | Pointer to **string** | Action against the traffic to  the virtual machine, for example deny, permit, etc. | [optional] [readonly] 
**Description** | Pointer to **string** | Description about the security group rule. | [optional] [readonly] 
**EndPort** | Pointer to **int64** | The end of port range for the security group rule IP protocol.When all the traffic is allowed into/from network boundary of virtual machine, both start port and end port will be zero. | [optional] [readonly] 
**EtherType** | Pointer to **string** | IP version of source CIDR (Classless Inter-Domain Routing), such as IPv4 or IPv6. * &#x60;IPv4&#x60; - Internet protocol version 4. * &#x60;IPv6&#x60; - Internet protocol version 6. | [optional] [readonly] [default to "IPv4"]
**Identity** | Pointer to **string** | Identity of this security group rule, which aids in uniquely identifying the security group rule. | [optional] [readonly] 
**Index** | Pointer to **int64** | Position of security group rule in a security group. | [optional] [readonly] 
**Name** | Pointer to **string** | User-provided name to identify the security group rule. | [optional] [readonly] 
**PortList** | Pointer to **[]int64** |  | [optional] 
**Protocol** | Pointer to **string** | The IP protocol name that&#39;s open to network traffic, such as TCP, UDP, etc. * &#x60;tcp&#x60; - The TCP (Transmission Control Protocol) protocol. * &#x60;udp&#x60; - The UDP (User Datagram Protocol) protocol. * &#x60;icmp&#x60; - The ICMP (Internet Control Message Protocol) protocol. * &#x60;esp&#x60; - The ESP (Encapsulating Security Payload) protocol. * &#x60;ah&#x60; - The AH (Authentication Header) protocol. * &#x60;sctp&#x60; - The SCTP (Stream Control Transmission Protocol) protocol. * &#x60;all&#x60; - All of TCP, UDP, ICMP, ESP, AH and SCTP. * &#x60;none&#x60; - None of TCP, UDP, ICMP, ESP, AH and SCTP. | [optional] [readonly] [default to "tcp"]
**SourceCidr** | Pointer to **[]string** |  | [optional] 
**SourceSecurityGroup** | Pointer to **string** | Reference to the existing security group, where the security group rule is defined. | [optional] [readonly] 
**StartPort** | Pointer to **int64** | The start of port range for the security group rule IP protocol. | [optional] [readonly] 

## Methods

### NewCloudSecurityGroupRule

`func NewCloudSecurityGroupRule(classId string, objectType string, ) *CloudSecurityGroupRule`

NewCloudSecurityGroupRule instantiates a new CloudSecurityGroupRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudSecurityGroupRuleWithDefaults

`func NewCloudSecurityGroupRuleWithDefaults() *CloudSecurityGroupRule`

NewCloudSecurityGroupRuleWithDefaults instantiates a new CloudSecurityGroupRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CloudSecurityGroupRule) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CloudSecurityGroupRule) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CloudSecurityGroupRule) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CloudSecurityGroupRule) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CloudSecurityGroupRule) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CloudSecurityGroupRule) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAction

`func (o *CloudSecurityGroupRule) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *CloudSecurityGroupRule) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *CloudSecurityGroupRule) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *CloudSecurityGroupRule) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetDescription

`func (o *CloudSecurityGroupRule) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CloudSecurityGroupRule) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CloudSecurityGroupRule) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CloudSecurityGroupRule) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEndPort

`func (o *CloudSecurityGroupRule) GetEndPort() int64`

GetEndPort returns the EndPort field if non-nil, zero value otherwise.

### GetEndPortOk

`func (o *CloudSecurityGroupRule) GetEndPortOk() (*int64, bool)`

GetEndPortOk returns a tuple with the EndPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndPort

`func (o *CloudSecurityGroupRule) SetEndPort(v int64)`

SetEndPort sets EndPort field to given value.

### HasEndPort

`func (o *CloudSecurityGroupRule) HasEndPort() bool`

HasEndPort returns a boolean if a field has been set.

### GetEtherType

`func (o *CloudSecurityGroupRule) GetEtherType() string`

GetEtherType returns the EtherType field if non-nil, zero value otherwise.

### GetEtherTypeOk

`func (o *CloudSecurityGroupRule) GetEtherTypeOk() (*string, bool)`

GetEtherTypeOk returns a tuple with the EtherType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEtherType

`func (o *CloudSecurityGroupRule) SetEtherType(v string)`

SetEtherType sets EtherType field to given value.

### HasEtherType

`func (o *CloudSecurityGroupRule) HasEtherType() bool`

HasEtherType returns a boolean if a field has been set.

### GetIdentity

`func (o *CloudSecurityGroupRule) GetIdentity() string`

GetIdentity returns the Identity field if non-nil, zero value otherwise.

### GetIdentityOk

`func (o *CloudSecurityGroupRule) GetIdentityOk() (*string, bool)`

GetIdentityOk returns a tuple with the Identity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentity

`func (o *CloudSecurityGroupRule) SetIdentity(v string)`

SetIdentity sets Identity field to given value.

### HasIdentity

`func (o *CloudSecurityGroupRule) HasIdentity() bool`

HasIdentity returns a boolean if a field has been set.

### GetIndex

`func (o *CloudSecurityGroupRule) GetIndex() int64`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *CloudSecurityGroupRule) GetIndexOk() (*int64, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *CloudSecurityGroupRule) SetIndex(v int64)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *CloudSecurityGroupRule) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### GetName

`func (o *CloudSecurityGroupRule) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudSecurityGroupRule) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudSecurityGroupRule) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CloudSecurityGroupRule) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPortList

`func (o *CloudSecurityGroupRule) GetPortList() []int64`

GetPortList returns the PortList field if non-nil, zero value otherwise.

### GetPortListOk

`func (o *CloudSecurityGroupRule) GetPortListOk() (*[]int64, bool)`

GetPortListOk returns a tuple with the PortList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortList

`func (o *CloudSecurityGroupRule) SetPortList(v []int64)`

SetPortList sets PortList field to given value.

### HasPortList

`func (o *CloudSecurityGroupRule) HasPortList() bool`

HasPortList returns a boolean if a field has been set.

### SetPortListNil

`func (o *CloudSecurityGroupRule) SetPortListNil(b bool)`

 SetPortListNil sets the value for PortList to be an explicit nil

### UnsetPortList
`func (o *CloudSecurityGroupRule) UnsetPortList()`

UnsetPortList ensures that no value is present for PortList, not even an explicit nil
### GetProtocol

`func (o *CloudSecurityGroupRule) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *CloudSecurityGroupRule) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *CloudSecurityGroupRule) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *CloudSecurityGroupRule) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetSourceCidr

`func (o *CloudSecurityGroupRule) GetSourceCidr() []string`

GetSourceCidr returns the SourceCidr field if non-nil, zero value otherwise.

### GetSourceCidrOk

`func (o *CloudSecurityGroupRule) GetSourceCidrOk() (*[]string, bool)`

GetSourceCidrOk returns a tuple with the SourceCidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceCidr

`func (o *CloudSecurityGroupRule) SetSourceCidr(v []string)`

SetSourceCidr sets SourceCidr field to given value.

### HasSourceCidr

`func (o *CloudSecurityGroupRule) HasSourceCidr() bool`

HasSourceCidr returns a boolean if a field has been set.

### SetSourceCidrNil

`func (o *CloudSecurityGroupRule) SetSourceCidrNil(b bool)`

 SetSourceCidrNil sets the value for SourceCidr to be an explicit nil

### UnsetSourceCidr
`func (o *CloudSecurityGroupRule) UnsetSourceCidr()`

UnsetSourceCidr ensures that no value is present for SourceCidr, not even an explicit nil
### GetSourceSecurityGroup

`func (o *CloudSecurityGroupRule) GetSourceSecurityGroup() string`

GetSourceSecurityGroup returns the SourceSecurityGroup field if non-nil, zero value otherwise.

### GetSourceSecurityGroupOk

`func (o *CloudSecurityGroupRule) GetSourceSecurityGroupOk() (*string, bool)`

GetSourceSecurityGroupOk returns a tuple with the SourceSecurityGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceSecurityGroup

`func (o *CloudSecurityGroupRule) SetSourceSecurityGroup(v string)`

SetSourceSecurityGroup sets SourceSecurityGroup field to given value.

### HasSourceSecurityGroup

`func (o *CloudSecurityGroupRule) HasSourceSecurityGroup() bool`

HasSourceSecurityGroup returns a boolean if a field has been set.

### GetStartPort

`func (o *CloudSecurityGroupRule) GetStartPort() int64`

GetStartPort returns the StartPort field if non-nil, zero value otherwise.

### GetStartPortOk

`func (o *CloudSecurityGroupRule) GetStartPortOk() (*int64, bool)`

GetStartPortOk returns a tuple with the StartPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartPort

`func (o *CloudSecurityGroupRule) SetStartPort(v int64)`

SetStartPort sets StartPort field to given value.

### HasStartPort

`func (o *CloudSecurityGroupRule) HasStartPort() bool`

HasStartPort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


