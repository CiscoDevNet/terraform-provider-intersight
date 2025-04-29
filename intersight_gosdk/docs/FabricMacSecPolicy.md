# FabricMacSecPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "fabric.MacSecPolicy"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "fabric.MacSecPolicy"]
**CipherSuite** | Pointer to **string** | Cipher suite to be used for MACsec encryption. * &#x60;GCM-AES-XPN-256&#x60; - An extended Cipher Suite of GCM-AES-256 used in MACsec (Media Access Control Security) that incorporates Extended Packet Numbering (XPN) for enhanced security and scalability. * &#x60;GCM-AES-128&#x60; - This Cipher Suite employs the Advanced Encryption Standard (AES) with a 128-bit key in Galois/Counter Mode, offering both encryption and authentication. * &#x60;GCM-AES-256&#x60; - This Cipher Suite utilizes Advanced Encryption Standard (AES) with a 256-bit key in Galois/Counter Mode, offering a higher level of security compared to GCM-AES-128 due to the larger key size. * &#x60;GCM-AES-XPN-128&#x60; - An extended Cipher Suite of GCM-AES-128  used in MACsec (Media Access Control Security) that incorporates Extended Packet Numbering (XPN) to enhance security and scalability. | [optional] [default to "GCM-AES-XPN-256"]
**ConfidentialityOffset** | Pointer to **string** | The MACsec confidentiality offset specifies the number of bytes starting from the frame header. MACsec encrypts only the bytes after the offset in a frame. * &#x60;CONF-OFFSET-0&#x60; - A value of 0 means the entire ethernet frame is encrypted. * &#x60;CONF-OFFSET-30&#x60; - The first 30 bytes of the ethernet frame are not encrypted, and the rest of the frame is encrypted. * &#x60;CONF-OFFSET-50&#x60; - The first 50 bytes of the ethernet frame are not encrypted, and the rest of the frame is encrypted. | [optional] [default to "CONF-OFFSET-0"]
**FallbackKeyChain** | Pointer to [**NullableFabricSecKeyChain**](FabricSecKeyChain.md) |  | [optional] 
**IncludeIcvIndicator** | Pointer to **bool** | Configures inclusion of the optional integrity check value (ICV) indicator as part of the transmitted MACsec key agreement protocol data unit (PDU). | [optional] [default to false]
**KeyServerPriority** | Pointer to **int64** | The key server is selected by comparing key-server priority values during MACsec key agreement (MKA) message exchange between peer devices. Valid values range from 0 to 255. The lower the value, the higher the chance it will be selected as the key server. | [optional] [default to 16]
**MacSecEaPol** | Pointer to [**NullableFabricMacSecEaPol**](FabricMacSecEaPol.md) |  | [optional] 
**PrimaryKeyChain** | Pointer to [**NullableFabricSecKeyChain**](FabricSecKeyChain.md) |  | [optional] 
**ReplayWindowSize** | Pointer to **int64** | Defines the size of the replay protection window. It determines the number of packets that can be received out of order without being considered replay attacks. | [optional] [default to 148809600]
**SakExpiryTime** | Pointer to **int64** | Time in seconds to force secure association key (SAK) rekey. Valid range is from 60 to 2592000 seconds when configured. When not configured, the SAK rekey interval is determined based on the interface speed. | [optional] 
**SecurityPolicy** | Pointer to **string** | The security policy specifies the level of MACsec enforcement on network traffic passing through a given interface. Should secure allows unencrypted traffic to flow until the MACsec key agreement (MKA) session is secured. After the MKA session is secured, the policy switches to only allow encrypted traffic to flow. Must secure imposes only MACsec encrypted traffic to flow. Traffic will be dropped, until the MKA session is not secured. * &#x60;Should-secure&#x60; - Should secure allows unencrypted traffic to flow until the MACsec key agreement (MKA) session is secured. After the MKA session is secured, the policy switches to only allow encrypted traffic to flow. * &#x60;Must-secure&#x60; - Must secure imposes only MACsec encrypted traffic to flow. Traffic will be dropped, until the MKA session is not secured. | [optional] [default to "Should-secure"]
**Organization** | Pointer to [**NullableOrganizationOrganizationRelationship**](OrganizationOrganizationRelationship.md) |  | [optional] 

## Methods

### NewFabricMacSecPolicy

`func NewFabricMacSecPolicy(classId string, objectType string, ) *FabricMacSecPolicy`

NewFabricMacSecPolicy instantiates a new FabricMacSecPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFabricMacSecPolicyWithDefaults

`func NewFabricMacSecPolicyWithDefaults() *FabricMacSecPolicy`

NewFabricMacSecPolicyWithDefaults instantiates a new FabricMacSecPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *FabricMacSecPolicy) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *FabricMacSecPolicy) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *FabricMacSecPolicy) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *FabricMacSecPolicy) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *FabricMacSecPolicy) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *FabricMacSecPolicy) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCipherSuite

`func (o *FabricMacSecPolicy) GetCipherSuite() string`

GetCipherSuite returns the CipherSuite field if non-nil, zero value otherwise.

### GetCipherSuiteOk

`func (o *FabricMacSecPolicy) GetCipherSuiteOk() (*string, bool)`

GetCipherSuiteOk returns a tuple with the CipherSuite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCipherSuite

`func (o *FabricMacSecPolicy) SetCipherSuite(v string)`

SetCipherSuite sets CipherSuite field to given value.

### HasCipherSuite

`func (o *FabricMacSecPolicy) HasCipherSuite() bool`

HasCipherSuite returns a boolean if a field has been set.

### GetConfidentialityOffset

`func (o *FabricMacSecPolicy) GetConfidentialityOffset() string`

GetConfidentialityOffset returns the ConfidentialityOffset field if non-nil, zero value otherwise.

### GetConfidentialityOffsetOk

`func (o *FabricMacSecPolicy) GetConfidentialityOffsetOk() (*string, bool)`

GetConfidentialityOffsetOk returns a tuple with the ConfidentialityOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidentialityOffset

`func (o *FabricMacSecPolicy) SetConfidentialityOffset(v string)`

SetConfidentialityOffset sets ConfidentialityOffset field to given value.

### HasConfidentialityOffset

`func (o *FabricMacSecPolicy) HasConfidentialityOffset() bool`

HasConfidentialityOffset returns a boolean if a field has been set.

### GetFallbackKeyChain

`func (o *FabricMacSecPolicy) GetFallbackKeyChain() FabricSecKeyChain`

GetFallbackKeyChain returns the FallbackKeyChain field if non-nil, zero value otherwise.

### GetFallbackKeyChainOk

`func (o *FabricMacSecPolicy) GetFallbackKeyChainOk() (*FabricSecKeyChain, bool)`

GetFallbackKeyChainOk returns a tuple with the FallbackKeyChain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallbackKeyChain

`func (o *FabricMacSecPolicy) SetFallbackKeyChain(v FabricSecKeyChain)`

SetFallbackKeyChain sets FallbackKeyChain field to given value.

### HasFallbackKeyChain

`func (o *FabricMacSecPolicy) HasFallbackKeyChain() bool`

HasFallbackKeyChain returns a boolean if a field has been set.

### SetFallbackKeyChainNil

`func (o *FabricMacSecPolicy) SetFallbackKeyChainNil(b bool)`

 SetFallbackKeyChainNil sets the value for FallbackKeyChain to be an explicit nil

### UnsetFallbackKeyChain
`func (o *FabricMacSecPolicy) UnsetFallbackKeyChain()`

UnsetFallbackKeyChain ensures that no value is present for FallbackKeyChain, not even an explicit nil
### GetIncludeIcvIndicator

`func (o *FabricMacSecPolicy) GetIncludeIcvIndicator() bool`

GetIncludeIcvIndicator returns the IncludeIcvIndicator field if non-nil, zero value otherwise.

### GetIncludeIcvIndicatorOk

`func (o *FabricMacSecPolicy) GetIncludeIcvIndicatorOk() (*bool, bool)`

GetIncludeIcvIndicatorOk returns a tuple with the IncludeIcvIndicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeIcvIndicator

`func (o *FabricMacSecPolicy) SetIncludeIcvIndicator(v bool)`

SetIncludeIcvIndicator sets IncludeIcvIndicator field to given value.

### HasIncludeIcvIndicator

`func (o *FabricMacSecPolicy) HasIncludeIcvIndicator() bool`

HasIncludeIcvIndicator returns a boolean if a field has been set.

### GetKeyServerPriority

`func (o *FabricMacSecPolicy) GetKeyServerPriority() int64`

GetKeyServerPriority returns the KeyServerPriority field if non-nil, zero value otherwise.

### GetKeyServerPriorityOk

`func (o *FabricMacSecPolicy) GetKeyServerPriorityOk() (*int64, bool)`

GetKeyServerPriorityOk returns a tuple with the KeyServerPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyServerPriority

`func (o *FabricMacSecPolicy) SetKeyServerPriority(v int64)`

SetKeyServerPriority sets KeyServerPriority field to given value.

### HasKeyServerPriority

`func (o *FabricMacSecPolicy) HasKeyServerPriority() bool`

HasKeyServerPriority returns a boolean if a field has been set.

### GetMacSecEaPol

`func (o *FabricMacSecPolicy) GetMacSecEaPol() FabricMacSecEaPol`

GetMacSecEaPol returns the MacSecEaPol field if non-nil, zero value otherwise.

### GetMacSecEaPolOk

`func (o *FabricMacSecPolicy) GetMacSecEaPolOk() (*FabricMacSecEaPol, bool)`

GetMacSecEaPolOk returns a tuple with the MacSecEaPol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacSecEaPol

`func (o *FabricMacSecPolicy) SetMacSecEaPol(v FabricMacSecEaPol)`

SetMacSecEaPol sets MacSecEaPol field to given value.

### HasMacSecEaPol

`func (o *FabricMacSecPolicy) HasMacSecEaPol() bool`

HasMacSecEaPol returns a boolean if a field has been set.

### SetMacSecEaPolNil

`func (o *FabricMacSecPolicy) SetMacSecEaPolNil(b bool)`

 SetMacSecEaPolNil sets the value for MacSecEaPol to be an explicit nil

### UnsetMacSecEaPol
`func (o *FabricMacSecPolicy) UnsetMacSecEaPol()`

UnsetMacSecEaPol ensures that no value is present for MacSecEaPol, not even an explicit nil
### GetPrimaryKeyChain

`func (o *FabricMacSecPolicy) GetPrimaryKeyChain() FabricSecKeyChain`

GetPrimaryKeyChain returns the PrimaryKeyChain field if non-nil, zero value otherwise.

### GetPrimaryKeyChainOk

`func (o *FabricMacSecPolicy) GetPrimaryKeyChainOk() (*FabricSecKeyChain, bool)`

GetPrimaryKeyChainOk returns a tuple with the PrimaryKeyChain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryKeyChain

`func (o *FabricMacSecPolicy) SetPrimaryKeyChain(v FabricSecKeyChain)`

SetPrimaryKeyChain sets PrimaryKeyChain field to given value.

### HasPrimaryKeyChain

`func (o *FabricMacSecPolicy) HasPrimaryKeyChain() bool`

HasPrimaryKeyChain returns a boolean if a field has been set.

### SetPrimaryKeyChainNil

`func (o *FabricMacSecPolicy) SetPrimaryKeyChainNil(b bool)`

 SetPrimaryKeyChainNil sets the value for PrimaryKeyChain to be an explicit nil

### UnsetPrimaryKeyChain
`func (o *FabricMacSecPolicy) UnsetPrimaryKeyChain()`

UnsetPrimaryKeyChain ensures that no value is present for PrimaryKeyChain, not even an explicit nil
### GetReplayWindowSize

`func (o *FabricMacSecPolicy) GetReplayWindowSize() int64`

GetReplayWindowSize returns the ReplayWindowSize field if non-nil, zero value otherwise.

### GetReplayWindowSizeOk

`func (o *FabricMacSecPolicy) GetReplayWindowSizeOk() (*int64, bool)`

GetReplayWindowSizeOk returns a tuple with the ReplayWindowSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplayWindowSize

`func (o *FabricMacSecPolicy) SetReplayWindowSize(v int64)`

SetReplayWindowSize sets ReplayWindowSize field to given value.

### HasReplayWindowSize

`func (o *FabricMacSecPolicy) HasReplayWindowSize() bool`

HasReplayWindowSize returns a boolean if a field has been set.

### GetSakExpiryTime

`func (o *FabricMacSecPolicy) GetSakExpiryTime() int64`

GetSakExpiryTime returns the SakExpiryTime field if non-nil, zero value otherwise.

### GetSakExpiryTimeOk

`func (o *FabricMacSecPolicy) GetSakExpiryTimeOk() (*int64, bool)`

GetSakExpiryTimeOk returns a tuple with the SakExpiryTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSakExpiryTime

`func (o *FabricMacSecPolicy) SetSakExpiryTime(v int64)`

SetSakExpiryTime sets SakExpiryTime field to given value.

### HasSakExpiryTime

`func (o *FabricMacSecPolicy) HasSakExpiryTime() bool`

HasSakExpiryTime returns a boolean if a field has been set.

### GetSecurityPolicy

`func (o *FabricMacSecPolicy) GetSecurityPolicy() string`

GetSecurityPolicy returns the SecurityPolicy field if non-nil, zero value otherwise.

### GetSecurityPolicyOk

`func (o *FabricMacSecPolicy) GetSecurityPolicyOk() (*string, bool)`

GetSecurityPolicyOk returns a tuple with the SecurityPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityPolicy

`func (o *FabricMacSecPolicy) SetSecurityPolicy(v string)`

SetSecurityPolicy sets SecurityPolicy field to given value.

### HasSecurityPolicy

`func (o *FabricMacSecPolicy) HasSecurityPolicy() bool`

HasSecurityPolicy returns a boolean if a field has been set.

### GetOrganization

`func (o *FabricMacSecPolicy) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *FabricMacSecPolicy) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *FabricMacSecPolicy) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *FabricMacSecPolicy) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### SetOrganizationNil

`func (o *FabricMacSecPolicy) SetOrganizationNil(b bool)`

 SetOrganizationNil sets the value for Organization to be an explicit nil

### UnsetOrganization
`func (o *FabricMacSecPolicy) UnsetOrganization()`

UnsetOrganization ensures that no value is present for Organization, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


