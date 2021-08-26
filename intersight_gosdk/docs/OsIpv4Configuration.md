# OsIpv4Configuration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "os.Ipv4Configuration"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "os.Ipv4Configuration"]
**IpV4Config** | Pointer to [**NullableCommIpV4Interface**](CommIpV4Interface.md) |  | [optional] 

## Methods

### NewOsIpv4Configuration

`func NewOsIpv4Configuration(classId string, objectType string, ) *OsIpv4Configuration`

NewOsIpv4Configuration instantiates a new OsIpv4Configuration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOsIpv4ConfigurationWithDefaults

`func NewOsIpv4ConfigurationWithDefaults() *OsIpv4Configuration`

NewOsIpv4ConfigurationWithDefaults instantiates a new OsIpv4Configuration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *OsIpv4Configuration) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *OsIpv4Configuration) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *OsIpv4Configuration) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *OsIpv4Configuration) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *OsIpv4Configuration) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *OsIpv4Configuration) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetIpV4Config

`func (o *OsIpv4Configuration) GetIpV4Config() CommIpV4Interface`

GetIpV4Config returns the IpV4Config field if non-nil, zero value otherwise.

### GetIpV4ConfigOk

`func (o *OsIpv4Configuration) GetIpV4ConfigOk() (*CommIpV4Interface, bool)`

GetIpV4ConfigOk returns a tuple with the IpV4Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpV4Config

`func (o *OsIpv4Configuration) SetIpV4Config(v CommIpV4Interface)`

SetIpV4Config sets IpV4Config field to given value.

### HasIpV4Config

`func (o *OsIpv4Configuration) HasIpV4Config() bool`

HasIpV4Config returns a boolean if a field has been set.

### SetIpV4ConfigNil

`func (o *OsIpv4Configuration) SetIpV4ConfigNil(b bool)`

 SetIpV4ConfigNil sets the value for IpV4Config to be an explicit nil

### UnsetIpV4Config
`func (o *OsIpv4Configuration) UnsetIpV4Config()`

UnsetIpV4Config ensures that no value is present for IpV4Config, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


