# CloudSecurityGroupRuleAllOf

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

### NewCloudSecurityGroupRuleAllOf

`func NewCloudSecurityGroupRuleAllOf(classId string, objectType string, ) *CloudSecurityGroupRuleAllOf`

NewCloudSecurityGroupRuleAllOf instantiates a new CloudSecurityGroupRuleAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudSecurityGroupRuleAllOfWithDefaults

`func NewCloudSecurityGroupRuleAllOfWithDefaults() *CloudSecurityGroupRuleAllOf`

NewCloudSecurityGroupRuleAllOfWithDefaults instantiates a new CloudSecurityGroupRuleAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CloudSecurityGroupRuleAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CloudSecurityGroupRuleAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CloudSecurityGroupRuleAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CloudSecurityGroupRuleAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CloudSecurityGroupRuleAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CloudSecurityGroupRuleAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAction

`func (o *CloudSecurityGroupRuleAllOf) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *CloudSecurityGroupRuleAllOf) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *CloudSecurityGroupRuleAllOf) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *CloudSecurityGroupRuleAllOf) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetDescription

`func (o *CloudSecurityGroupRuleAllOf) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CloudSecurityGroupRuleAllOf) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CloudSecurityGroupRuleAllOf) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CloudSecurityGroupRuleAllOf) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEndPort

`func (o *CloudSecurityGroupRuleAllOf) GetEndPort() int64`

GetEndPort returns the EndPort field if non-nil, zero value otherwise.

### GetEndPortOk

`func (o *CloudSecurityGroupRuleAllOf) GetEndPortOk() (*int64, bool)`

GetEndPortOk returns a tuple with the EndPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndPort

`func (o *CloudSecurityGroupRuleAllOf) SetEndPort(v int64)`

SetEndPort sets EndPort field to given value.

### HasEndPort

`func (o *CloudSecurityGroupRuleAllOf) HasEndPort() bool`

HasEndPort returns a boolean if a field has been set.

### GetEtherType

`func (o *CloudSecurityGroupRuleAllOf) GetEtherType() string`

GetEtherType returns the EtherType field if non-nil, zero value otherwise.

### GetEtherTypeOk

`func (o *CloudSecurityGroupRuleAllOf) GetEtherTypeOk() (*string, bool)`

GetEtherTypeOk returns a tuple with the EtherType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEtherType

`func (o *CloudSecurityGroupRuleAllOf) SetEtherType(v string)`

SetEtherType sets EtherType field to given value.

### HasEtherType

`func (o *CloudSecurityGroupRuleAllOf) HasEtherType() bool`

HasEtherType returns a boolean if a field has been set.

### GetIdentity

`func (o *CloudSecurityGroupRuleAllOf) GetIdentity() string`

GetIdentity returns the Identity field if non-nil, zero value otherwise.

### GetIdentityOk

`func (o *CloudSecurityGroupRuleAllOf) GetIdentityOk() (*string, bool)`

GetIdentityOk returns a tuple with the Identity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentity

`func (o *CloudSecurityGroupRuleAllOf) SetIdentity(v string)`

SetIdentity sets Identity field to given value.

### HasIdentity

`func (o *CloudSecurityGroupRuleAllOf) HasIdentity() bool`

HasIdentity returns a boolean if a field has been set.

### GetIndex

`func (o *CloudSecurityGroupRuleAllOf) GetIndex() int64`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *CloudSecurityGroupRuleAllOf) GetIndexOk() (*int64, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *CloudSecurityGroupRuleAllOf) SetIndex(v int64)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *CloudSecurityGroupRuleAllOf) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### GetName

`func (o *CloudSecurityGroupRuleAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudSecurityGroupRuleAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudSecurityGroupRuleAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CloudSecurityGroupRuleAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPortList

`func (o *CloudSecurityGroupRuleAllOf) GetPortList() []int64`

GetPortList returns the PortList field if non-nil, zero value otherwise.

### GetPortListOk

`func (o *CloudSecurityGroupRuleAllOf) GetPortListOk() (*[]int64, bool)`

GetPortListOk returns a tuple with the PortList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortList

`func (o *CloudSecurityGroupRuleAllOf) SetPortList(v []int64)`

SetPortList sets PortList field to given value.

### HasPortList

`func (o *CloudSecurityGroupRuleAllOf) HasPortList() bool`

HasPortList returns a boolean if a field has been set.

### SetPortListNil

`func (o *CloudSecurityGroupRuleAllOf) SetPortListNil(b bool)`

 SetPortListNil sets the value for PortList to be an explicit nil

### UnsetPortList
`func (o *CloudSecurityGroupRuleAllOf) UnsetPortList()`

UnsetPortList ensures that no value is present for PortList, not even an explicit nil
### GetProtocol

`func (o *CloudSecurityGroupRuleAllOf) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *CloudSecurityGroupRuleAllOf) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *CloudSecurityGroupRuleAllOf) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *CloudSecurityGroupRuleAllOf) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetSourceCidr

`func (o *CloudSecurityGroupRuleAllOf) GetSourceCidr() []string`

GetSourceCidr returns the SourceCidr field if non-nil, zero value otherwise.

### GetSourceCidrOk

`func (o *CloudSecurityGroupRuleAllOf) GetSourceCidrOk() (*[]string, bool)`

GetSourceCidrOk returns a tuple with the SourceCidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceCidr

`func (o *CloudSecurityGroupRuleAllOf) SetSourceCidr(v []string)`

SetSourceCidr sets SourceCidr field to given value.

### HasSourceCidr

`func (o *CloudSecurityGroupRuleAllOf) HasSourceCidr() bool`

HasSourceCidr returns a boolean if a field has been set.

### SetSourceCidrNil

`func (o *CloudSecurityGroupRuleAllOf) SetSourceCidrNil(b bool)`

 SetSourceCidrNil sets the value for SourceCidr to be an explicit nil

### UnsetSourceCidr
`func (o *CloudSecurityGroupRuleAllOf) UnsetSourceCidr()`

UnsetSourceCidr ensures that no value is present for SourceCidr, not even an explicit nil
### GetSourceSecurityGroup

`func (o *CloudSecurityGroupRuleAllOf) GetSourceSecurityGroup() string`

GetSourceSecurityGroup returns the SourceSecurityGroup field if non-nil, zero value otherwise.

### GetSourceSecurityGroupOk

`func (o *CloudSecurityGroupRuleAllOf) GetSourceSecurityGroupOk() (*string, bool)`

GetSourceSecurityGroupOk returns a tuple with the SourceSecurityGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceSecurityGroup

`func (o *CloudSecurityGroupRuleAllOf) SetSourceSecurityGroup(v string)`

SetSourceSecurityGroup sets SourceSecurityGroup field to given value.

### HasSourceSecurityGroup

`func (o *CloudSecurityGroupRuleAllOf) HasSourceSecurityGroup() bool`

HasSourceSecurityGroup returns a boolean if a field has been set.

### GetStartPort

`func (o *CloudSecurityGroupRuleAllOf) GetStartPort() int64`

GetStartPort returns the StartPort field if non-nil, zero value otherwise.

### GetStartPortOk

`func (o *CloudSecurityGroupRuleAllOf) GetStartPortOk() (*int64, bool)`

GetStartPortOk returns a tuple with the StartPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartPort

`func (o *CloudSecurityGroupRuleAllOf) SetStartPort(v int64)`

SetStartPort sets StartPort field to given value.

### HasStartPort

`func (o *CloudSecurityGroupRuleAllOf) HasStartPort() bool`

HasStartPort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


