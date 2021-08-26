# CloudNetworkAccessConfigAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "cloud.NetworkAccessConfig"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "cloud.NetworkAccessConfig"]
**ExternalIps** | Pointer to [**[]CloudNetworkAddress**](CloudNetworkAddress.md) |  | [optional] 
**PrivateDns** | Pointer to **string** | Private DNS name assigned to the network interface. | [optional] [readonly] 
**PublicDns** | Pointer to **string** | Public DNS name assigned to the network interface. | [optional] [readonly] 

## Methods

### NewCloudNetworkAccessConfigAllOf

`func NewCloudNetworkAccessConfigAllOf(classId string, objectType string, ) *CloudNetworkAccessConfigAllOf`

NewCloudNetworkAccessConfigAllOf instantiates a new CloudNetworkAccessConfigAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudNetworkAccessConfigAllOfWithDefaults

`func NewCloudNetworkAccessConfigAllOfWithDefaults() *CloudNetworkAccessConfigAllOf`

NewCloudNetworkAccessConfigAllOfWithDefaults instantiates a new CloudNetworkAccessConfigAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CloudNetworkAccessConfigAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CloudNetworkAccessConfigAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CloudNetworkAccessConfigAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CloudNetworkAccessConfigAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CloudNetworkAccessConfigAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CloudNetworkAccessConfigAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetExternalIps

`func (o *CloudNetworkAccessConfigAllOf) GetExternalIps() []CloudNetworkAddress`

GetExternalIps returns the ExternalIps field if non-nil, zero value otherwise.

### GetExternalIpsOk

`func (o *CloudNetworkAccessConfigAllOf) GetExternalIpsOk() (*[]CloudNetworkAddress, bool)`

GetExternalIpsOk returns a tuple with the ExternalIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIps

`func (o *CloudNetworkAccessConfigAllOf) SetExternalIps(v []CloudNetworkAddress)`

SetExternalIps sets ExternalIps field to given value.

### HasExternalIps

`func (o *CloudNetworkAccessConfigAllOf) HasExternalIps() bool`

HasExternalIps returns a boolean if a field has been set.

### SetExternalIpsNil

`func (o *CloudNetworkAccessConfigAllOf) SetExternalIpsNil(b bool)`

 SetExternalIpsNil sets the value for ExternalIps to be an explicit nil

### UnsetExternalIps
`func (o *CloudNetworkAccessConfigAllOf) UnsetExternalIps()`

UnsetExternalIps ensures that no value is present for ExternalIps, not even an explicit nil
### GetPrivateDns

`func (o *CloudNetworkAccessConfigAllOf) GetPrivateDns() string`

GetPrivateDns returns the PrivateDns field if non-nil, zero value otherwise.

### GetPrivateDnsOk

`func (o *CloudNetworkAccessConfigAllOf) GetPrivateDnsOk() (*string, bool)`

GetPrivateDnsOk returns a tuple with the PrivateDns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateDns

`func (o *CloudNetworkAccessConfigAllOf) SetPrivateDns(v string)`

SetPrivateDns sets PrivateDns field to given value.

### HasPrivateDns

`func (o *CloudNetworkAccessConfigAllOf) HasPrivateDns() bool`

HasPrivateDns returns a boolean if a field has been set.

### GetPublicDns

`func (o *CloudNetworkAccessConfigAllOf) GetPublicDns() string`

GetPublicDns returns the PublicDns field if non-nil, zero value otherwise.

### GetPublicDnsOk

`func (o *CloudNetworkAccessConfigAllOf) GetPublicDnsOk() (*string, bool)`

GetPublicDnsOk returns a tuple with the PublicDns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicDns

`func (o *CloudNetworkAccessConfigAllOf) SetPublicDns(v string)`

SetPublicDns sets PublicDns field to given value.

### HasPublicDns

`func (o *CloudNetworkAccessConfigAllOf) HasPublicDns() bool`

HasPublicDns returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


