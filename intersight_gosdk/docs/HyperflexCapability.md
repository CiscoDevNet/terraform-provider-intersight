# HyperflexCapability

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hyperflex.Capability"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hyperflex.Capability"]
**EncryptionSupported** | Pointer to **bool** | Specifies if encryption is supported in the HyperFlex cluster. The HyperFlex cluster supports self encrypting drives (SED). | [optional] [readonly] 
**IscsiSupported** | Pointer to **bool** | Specifies if iSCSI is supported in the HyperFlex cluster. The HyperFlex cluster supports an iSCSI network, iSCSI initiator groups, targets, and logical unit number (LUN). | [optional] [readonly] 
**ReplicationSupported** | Pointer to **bool** | Specifies if replication is supported in the HyperFlex cluster.  The HyperFlex cluster supports 1:1 disaster recovery and N:1 backup. | [optional] [readonly] 

## Methods

### NewHyperflexCapability

`func NewHyperflexCapability(classId string, objectType string, ) *HyperflexCapability`

NewHyperflexCapability instantiates a new HyperflexCapability object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexCapabilityWithDefaults

`func NewHyperflexCapabilityWithDefaults() *HyperflexCapability`

NewHyperflexCapabilityWithDefaults instantiates a new HyperflexCapability object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HyperflexCapability) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexCapability) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexCapability) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HyperflexCapability) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexCapability) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexCapability) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetEncryptionSupported

`func (o *HyperflexCapability) GetEncryptionSupported() bool`

GetEncryptionSupported returns the EncryptionSupported field if non-nil, zero value otherwise.

### GetEncryptionSupportedOk

`func (o *HyperflexCapability) GetEncryptionSupportedOk() (*bool, bool)`

GetEncryptionSupportedOk returns a tuple with the EncryptionSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionSupported

`func (o *HyperflexCapability) SetEncryptionSupported(v bool)`

SetEncryptionSupported sets EncryptionSupported field to given value.

### HasEncryptionSupported

`func (o *HyperflexCapability) HasEncryptionSupported() bool`

HasEncryptionSupported returns a boolean if a field has been set.

### GetIscsiSupported

`func (o *HyperflexCapability) GetIscsiSupported() bool`

GetIscsiSupported returns the IscsiSupported field if non-nil, zero value otherwise.

### GetIscsiSupportedOk

`func (o *HyperflexCapability) GetIscsiSupportedOk() (*bool, bool)`

GetIscsiSupportedOk returns a tuple with the IscsiSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIscsiSupported

`func (o *HyperflexCapability) SetIscsiSupported(v bool)`

SetIscsiSupported sets IscsiSupported field to given value.

### HasIscsiSupported

`func (o *HyperflexCapability) HasIscsiSupported() bool`

HasIscsiSupported returns a boolean if a field has been set.

### GetReplicationSupported

`func (o *HyperflexCapability) GetReplicationSupported() bool`

GetReplicationSupported returns the ReplicationSupported field if non-nil, zero value otherwise.

### GetReplicationSupportedOk

`func (o *HyperflexCapability) GetReplicationSupportedOk() (*bool, bool)`

GetReplicationSupportedOk returns a tuple with the ReplicationSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationSupported

`func (o *HyperflexCapability) SetReplicationSupported(v bool)`

SetReplicationSupported sets ReplicationSupported field to given value.

### HasReplicationSupported

`func (o *HyperflexCapability) HasReplicationSupported() bool`

HasReplicationSupported returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


