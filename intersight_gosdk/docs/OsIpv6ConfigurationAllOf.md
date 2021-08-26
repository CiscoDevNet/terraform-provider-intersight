# OsIpv6ConfigurationAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "os.Ipv6Configuration"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "os.Ipv6Configuration"]
**IpV6Config** | Pointer to [**NullableCommIpV6Interface**](CommIpV6Interface.md) |  | [optional] 

## Methods

### NewOsIpv6ConfigurationAllOf

`func NewOsIpv6ConfigurationAllOf(classId string, objectType string, ) *OsIpv6ConfigurationAllOf`

NewOsIpv6ConfigurationAllOf instantiates a new OsIpv6ConfigurationAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOsIpv6ConfigurationAllOfWithDefaults

`func NewOsIpv6ConfigurationAllOfWithDefaults() *OsIpv6ConfigurationAllOf`

NewOsIpv6ConfigurationAllOfWithDefaults instantiates a new OsIpv6ConfigurationAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *OsIpv6ConfigurationAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *OsIpv6ConfigurationAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *OsIpv6ConfigurationAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *OsIpv6ConfigurationAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *OsIpv6ConfigurationAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *OsIpv6ConfigurationAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetIpV6Config

`func (o *OsIpv6ConfigurationAllOf) GetIpV6Config() CommIpV6Interface`

GetIpV6Config returns the IpV6Config field if non-nil, zero value otherwise.

### GetIpV6ConfigOk

`func (o *OsIpv6ConfigurationAllOf) GetIpV6ConfigOk() (*CommIpV6Interface, bool)`

GetIpV6ConfigOk returns a tuple with the IpV6Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpV6Config

`func (o *OsIpv6ConfigurationAllOf) SetIpV6Config(v CommIpV6Interface)`

SetIpV6Config sets IpV6Config field to given value.

### HasIpV6Config

`func (o *OsIpv6ConfigurationAllOf) HasIpV6Config() bool`

HasIpV6Config returns a boolean if a field has been set.

### SetIpV6ConfigNil

`func (o *OsIpv6ConfigurationAllOf) SetIpV6ConfigNil(b bool)`

 SetIpV6ConfigNil sets the value for IpV6Config to be an explicit nil

### UnsetIpV6Config
`func (o *OsIpv6ConfigurationAllOf) UnsetIpV6Config()`

UnsetIpV6Config ensures that no value is present for IpV6Config, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


