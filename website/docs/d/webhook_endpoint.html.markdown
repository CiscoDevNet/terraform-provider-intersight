---
subcategory: "webhook"
layout: "intersight"
page_title: "Intersight: intersight_webhook_endpoint"
description: |-
        The webhook endpoint which an controller can use to forward events.

---

# Data Source: intersight_webhook_endpoint
The webhook endpoint which an controller can use to forward events.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_webhook_endpoint.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `auth_type`:(string) Type of authentication used by the clientApp.* `basic` - The client uses basic username/password authentication. The password is expected to be a JWT token.* `none` - No authentication method specified by the client.* `bearer-token` - The client uses a long-lived bearer token to authenticate.* `auth-code` - The client uses OAuth Authorization Grant Flow without PKCE for authentication.* `client-credentials` - The client uses OAuth Client Credentials Flow for authentication. 
* `create_time`:(string) The time when this managed object was created. 
* `credentials_action`:(string) An action to be performed on the credentials.* `none` - No action to be performed.* `regenerateCredentials` - Allows for revocation and regeneration of a token. The old token associated with the client application. will not be usable and a new token will be generated. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) The name of the Endpoint. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `url`:(string) The endpoint URL. The CREATE and UPDATE APIs can cause the change to the value. 
 
