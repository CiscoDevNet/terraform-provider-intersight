# VnicIscsiAdapterPolicyAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "vnic.IscsiAdapterPolicy"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "vnic.IscsiAdapterPolicy"]
**ConnectionTimeOut** | Pointer to **int64** | The number of seconds to wait until Cisco UCS assumes that the initial login has failed and the iSCSI adapter is unavailable. | [optional] 
**DhcpTimeout** | Pointer to **int64** | The number of seconds to wait before the initiator assumes that the DHCP server is unavailable. | [optional] 
**LunBusyRetryCount** | Pointer to **int64** | The number of times to retry the connection in case of a failure during iSCSI LUN discovery. | [optional] 
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](OrganizationOrganizationRelationship.md) |  | [optional] 

## Methods

### NewVnicIscsiAdapterPolicyAllOf

`func NewVnicIscsiAdapterPolicyAllOf(classId string, objectType string, ) *VnicIscsiAdapterPolicyAllOf`

NewVnicIscsiAdapterPolicyAllOf instantiates a new VnicIscsiAdapterPolicyAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVnicIscsiAdapterPolicyAllOfWithDefaults

`func NewVnicIscsiAdapterPolicyAllOfWithDefaults() *VnicIscsiAdapterPolicyAllOf`

NewVnicIscsiAdapterPolicyAllOfWithDefaults instantiates a new VnicIscsiAdapterPolicyAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VnicIscsiAdapterPolicyAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VnicIscsiAdapterPolicyAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VnicIscsiAdapterPolicyAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VnicIscsiAdapterPolicyAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VnicIscsiAdapterPolicyAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VnicIscsiAdapterPolicyAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetConnectionTimeOut

`func (o *VnicIscsiAdapterPolicyAllOf) GetConnectionTimeOut() int64`

GetConnectionTimeOut returns the ConnectionTimeOut field if non-nil, zero value otherwise.

### GetConnectionTimeOutOk

`func (o *VnicIscsiAdapterPolicyAllOf) GetConnectionTimeOutOk() (*int64, bool)`

GetConnectionTimeOutOk returns a tuple with the ConnectionTimeOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionTimeOut

`func (o *VnicIscsiAdapterPolicyAllOf) SetConnectionTimeOut(v int64)`

SetConnectionTimeOut sets ConnectionTimeOut field to given value.

### HasConnectionTimeOut

`func (o *VnicIscsiAdapterPolicyAllOf) HasConnectionTimeOut() bool`

HasConnectionTimeOut returns a boolean if a field has been set.

### GetDhcpTimeout

`func (o *VnicIscsiAdapterPolicyAllOf) GetDhcpTimeout() int64`

GetDhcpTimeout returns the DhcpTimeout field if non-nil, zero value otherwise.

### GetDhcpTimeoutOk

`func (o *VnicIscsiAdapterPolicyAllOf) GetDhcpTimeoutOk() (*int64, bool)`

GetDhcpTimeoutOk returns a tuple with the DhcpTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpTimeout

`func (o *VnicIscsiAdapterPolicyAllOf) SetDhcpTimeout(v int64)`

SetDhcpTimeout sets DhcpTimeout field to given value.

### HasDhcpTimeout

`func (o *VnicIscsiAdapterPolicyAllOf) HasDhcpTimeout() bool`

HasDhcpTimeout returns a boolean if a field has been set.

### GetLunBusyRetryCount

`func (o *VnicIscsiAdapterPolicyAllOf) GetLunBusyRetryCount() int64`

GetLunBusyRetryCount returns the LunBusyRetryCount field if non-nil, zero value otherwise.

### GetLunBusyRetryCountOk

`func (o *VnicIscsiAdapterPolicyAllOf) GetLunBusyRetryCountOk() (*int64, bool)`

GetLunBusyRetryCountOk returns a tuple with the LunBusyRetryCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLunBusyRetryCount

`func (o *VnicIscsiAdapterPolicyAllOf) SetLunBusyRetryCount(v int64)`

SetLunBusyRetryCount sets LunBusyRetryCount field to given value.

### HasLunBusyRetryCount

`func (o *VnicIscsiAdapterPolicyAllOf) HasLunBusyRetryCount() bool`

HasLunBusyRetryCount returns a boolean if a field has been set.

### GetOrganization

`func (o *VnicIscsiAdapterPolicyAllOf) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *VnicIscsiAdapterPolicyAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *VnicIscsiAdapterPolicyAllOf) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *VnicIscsiAdapterPolicyAllOf) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


