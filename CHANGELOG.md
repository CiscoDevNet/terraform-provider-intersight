# Change Logs

## v1.0.6
Release Date: 19th Apr 2021
* Refreshes the terraform module to be in sync with the latest Cisco Intersight model (build 4240)
* Better error handling when data sources return no results 
* Default values added to ObjectType and ClassId properties

## v1.0.5
Release Date: 9th Apr 2021
* Refreshes the terraform module to be in sync with the latest Cisco Intersight model (build 4155).
* Some errors are reported better instead of a generic typecasting failure message.
* Some properties like CreateTime, Parents, Owners, Resources which were getting intentionally excluded are not included.
* Allows the provider to work from behind a proxy server.
* Adds usage example for over 200 resources.
* Documentation has been enhanced to showcase the usage examples first in the resource documentation.


## v1.0.4
Release Date: 30th Mar 2021
* Refreshes the terraform module to be in sync with the latest Cisco Intersight model (build 4136).

## v1.0.3
Release Date: 19th Mar 2021
* Refreshes the terraform module to be in sync with the latest Cisco Intersight model (build 3942).
* Provides an optional `wait_for_completion` property to disable workflow wait feature. It is set to true by default.
The following configuration triggers a server profile deployment and returns without waiting for the workflow completion.
```
resource "intersight_server_profile" "demo_profile" {
  name = "test-ci"
  action="Deploy"
  wait_for_completion=false
}
```

## v1.0.2
Release Date: 5th Mar 2021
* Refreshes the terraform module to be in sync with the latest Cisco Intersight model (build 3824).
* `terraform apply` will wait till all workflows triggered due to the operation reach an end state.
* Data sources now store the results as a list. The syntax to access a property in the `i`th index of results is `data.intersight_<model_name>.<custom_name>.results[i].<property_name>`.

## v1.0.1
Release Date: 26th Feb 2021
* Refreshes the terraform module to be in sync with the latest Cisco Intersight model (build 3714).

## v1.0.0
Release Date: 5th Feb 2021
*  Refreshes the terraform module to be in sync with the latest Cisco Intersight model (build 3562).
*  Changes `secretkeyfile` to `secretkey` in provider block. Although minor, this is a breaking change. We have updated the major version of the provider to reflect the same. The new `secretkey` property also accepts the api secret text as a string in addition to the earlier supported file path string. The existing configuration files must rename `secretkeyfile` to `secretkey` when using Intersight provider version >=1.0.0
*  Some Intersight specific properties like `object_type` and `class_id` are now internally set by the provider internally (except for a few special cases). A user need not provide these anymore. The documentation is updated to reflect the change.
*  An effort is made to make the documentation richer by adding  hand-written examples for some resources. An infrastructure is put in place for the community to contribute examples to the documentation.
*  Documentation for complex type objects is enhanced to include exact types and properties for each internal object.

## v0.1.5
Release Data: 22nd Jan 2021
* Refreshes the terraform module to be in sync with the latest Cisco Intersight model (build 3252).

## v0.1.4
Release Data: 15th Jan 2021
* Refreshes the terraform module to be in sync with the latest Cisco Intersight model (build 3181).

## v0.1.3
Release Data: 8th Jan 2021
* Refreshes the terraform module to be in sync with the latest Cisco Intersight model (build 3144).

## v0.1.2
Release Date: 22nd December 2020
* Refreshes the terraform module to be in sync with the latest Cisco Intersight model (build 3127).

## v0.1.1
Release Date: 22nd Nov 2020
* Refreshes the terraform module to be in sync with the latest Cisco Intersight model.

## v0.1.0 
Release Date: 13th October 2020
* Initial release
