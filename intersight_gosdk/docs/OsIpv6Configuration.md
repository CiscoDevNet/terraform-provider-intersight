# OsIpv6Configuration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "os.Ipv6Configuration"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "os.Ipv6Configuration"]
**IpV6Config** | Pointer to [**NullableCommIpV6Interface**](CommIpV6Interface.md) |  | [optional] 

## Methods

### NewOsIpv6Configuration

`func NewOsIpv6Configuration(classId string, objectType string, ) *OsIpv6Configuration`

NewOsIpv6Configuration instantiates a new OsIpv6Configuration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOsIpv6ConfigurationWithDefaults

`func NewOsIpv6ConfigurationWithDefaults() *OsIpv6Configuration`

NewOsIpv6ConfigurationWithDefaults instantiates a new OsIpv6Configuration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *OsIpv6Configuration) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *OsIpv6Configuration) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *OsIpv6Configuration) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *OsIpv6Configuration) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *OsIpv6Configuration) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *OsIpv6Configuration) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetIpV6Config

`func (o *OsIpv6Configuration) GetIpV6Config() CommIpV6Interface`

GetIpV6Config returns the IpV6Config field if non-nil, zero value otherwise.

### GetIpV6ConfigOk

`func (o *OsIpv6Configuration) GetIpV6ConfigOk() (*CommIpV6Interface, bool)`

GetIpV6ConfigOk returns a tuple with the IpV6Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpV6Config

`func (o *OsIpv6Configuration) SetIpV6Config(v CommIpV6Interface)`

SetIpV6Config sets IpV6Config field to given value.

### HasIpV6Config

`func (o *OsIpv6Configuration) HasIpV6Config() bool`

HasIpV6Config returns a boolean if a field has been set.

### SetIpV6ConfigNil

`func (o *OsIpv6Configuration) SetIpV6ConfigNil(b bool)`

 SetIpV6ConfigNil sets the value for IpV6Config to be an explicit nil

### UnsetIpV6Config
`func (o *OsIpv6Configuration) UnsetIpV6Config()`

UnsetIpV6Config ensures that no value is present for IpV6Config, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


