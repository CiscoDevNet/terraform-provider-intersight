# ConnectorUrlAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ForceQuery** | Pointer to **bool** | Flag to append a query to the url even if rawQuery is empty. | [optional] 
**Fragment** | Pointer to **string** | The fragment identifier component of a URI allows indirect identification of a secondary resource by reference to a primary resource and additional identifying information. The identified secondary resource may be some portion or subset of the primary resource, some view on representations of the primary resource, or some other resource defined or described by those representations. A fragment identifier component is indicated by the presence of a number sign (\&quot;#\&quot;) character and terminated by the end of the URI. | [optional] 
**Host** | Pointer to **string** | The host name identifies the host that holds the resource. The host can be an IP or a hostname that is resolvable by the dns server configured on the platform. | [optional] 
**Opaque** | Pointer to **string** | A URI is opaque if, and only if, it is absolute and its scheme-specific part does not begin with a slash character (&#39;/&#39;). An opaque URI has a scheme, a scheme-specific part, and possibly a fragment; all other components are undefined. | [optional] 
**Path** | Pointer to **string** | The path identifies the specific resource in the host that the web client wants to access. Value is the decoded form of the path. e.g. &#39;/foo/bar&#39;. | [optional] 
**RawPath** | Pointer to **string** | The URI encoded form of the path property. e.g. &#39;%2Fapi%2Fv1%2F&#39;. | [optional] 
**RawQuery** | Pointer to **string** | The query component, as defined in RFC 3986, contains non-hierarchical data that, along with data in the path component, serves to identify a resource within the scope of the URI&#39;s scheme and naming authority (if any). The query component is indicated by the first question mark character and terminated by a number sign character or by the end of the URI. The rawQuery contains the URIs encoded query component, excluding the ? character. | [optional] 
**Scheme** | Pointer to **string** | The scheme identifies the protocol to be used to access the resource on the Internet. It can be HTTP (without SSL) or HTTPS (with SSL). | [optional] 

## Methods

### NewConnectorUrlAllOf

`func NewConnectorUrlAllOf() *ConnectorUrlAllOf`

NewConnectorUrlAllOf instantiates a new ConnectorUrlAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectorUrlAllOfWithDefaults

`func NewConnectorUrlAllOfWithDefaults() *ConnectorUrlAllOf`

NewConnectorUrlAllOfWithDefaults instantiates a new ConnectorUrlAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetForceQuery

`func (o *ConnectorUrlAllOf) GetForceQuery() bool`

GetForceQuery returns the ForceQuery field if non-nil, zero value otherwise.

### GetForceQueryOk

`func (o *ConnectorUrlAllOf) GetForceQueryOk() (*bool, bool)`

GetForceQueryOk returns a tuple with the ForceQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceQuery

`func (o *ConnectorUrlAllOf) SetForceQuery(v bool)`

SetForceQuery sets ForceQuery field to given value.

### HasForceQuery

`func (o *ConnectorUrlAllOf) HasForceQuery() bool`

HasForceQuery returns a boolean if a field has been set.

### GetFragment

`func (o *ConnectorUrlAllOf) GetFragment() string`

GetFragment returns the Fragment field if non-nil, zero value otherwise.

### GetFragmentOk

`func (o *ConnectorUrlAllOf) GetFragmentOk() (*string, bool)`

GetFragmentOk returns a tuple with the Fragment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFragment

`func (o *ConnectorUrlAllOf) SetFragment(v string)`

SetFragment sets Fragment field to given value.

### HasFragment

`func (o *ConnectorUrlAllOf) HasFragment() bool`

HasFragment returns a boolean if a field has been set.

### GetHost

`func (o *ConnectorUrlAllOf) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *ConnectorUrlAllOf) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *ConnectorUrlAllOf) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *ConnectorUrlAllOf) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetOpaque

`func (o *ConnectorUrlAllOf) GetOpaque() string`

GetOpaque returns the Opaque field if non-nil, zero value otherwise.

### GetOpaqueOk

`func (o *ConnectorUrlAllOf) GetOpaqueOk() (*string, bool)`

GetOpaqueOk returns a tuple with the Opaque field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpaque

`func (o *ConnectorUrlAllOf) SetOpaque(v string)`

SetOpaque sets Opaque field to given value.

### HasOpaque

`func (o *ConnectorUrlAllOf) HasOpaque() bool`

HasOpaque returns a boolean if a field has been set.

### GetPath

`func (o *ConnectorUrlAllOf) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *ConnectorUrlAllOf) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *ConnectorUrlAllOf) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *ConnectorUrlAllOf) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetRawPath

`func (o *ConnectorUrlAllOf) GetRawPath() string`

GetRawPath returns the RawPath field if non-nil, zero value otherwise.

### GetRawPathOk

`func (o *ConnectorUrlAllOf) GetRawPathOk() (*string, bool)`

GetRawPathOk returns a tuple with the RawPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawPath

`func (o *ConnectorUrlAllOf) SetRawPath(v string)`

SetRawPath sets RawPath field to given value.

### HasRawPath

`func (o *ConnectorUrlAllOf) HasRawPath() bool`

HasRawPath returns a boolean if a field has been set.

### GetRawQuery

`func (o *ConnectorUrlAllOf) GetRawQuery() string`

GetRawQuery returns the RawQuery field if non-nil, zero value otherwise.

### GetRawQueryOk

`func (o *ConnectorUrlAllOf) GetRawQueryOk() (*string, bool)`

GetRawQueryOk returns a tuple with the RawQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawQuery

`func (o *ConnectorUrlAllOf) SetRawQuery(v string)`

SetRawQuery sets RawQuery field to given value.

### HasRawQuery

`func (o *ConnectorUrlAllOf) HasRawQuery() bool`

HasRawQuery returns a boolean if a field has been set.

### GetScheme

`func (o *ConnectorUrlAllOf) GetScheme() string`

GetScheme returns the Scheme field if non-nil, zero value otherwise.

### GetSchemeOk

`func (o *ConnectorUrlAllOf) GetSchemeOk() (*string, bool)`

GetSchemeOk returns a tuple with the Scheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheme

`func (o *ConnectorUrlAllOf) SetScheme(v string)`

SetScheme sets Scheme field to given value.

### HasScheme

`func (o *ConnectorUrlAllOf) HasScheme() bool`

HasScheme returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


