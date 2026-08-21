resource "intersight_ippool_pool" "ippool_pool_tf" {
  name             = "ippool_pool_tf"
  description      = "ippool test"
  assignment_order = "sequential"
  
  ip_v4_blocks {
  class_id = "ippool.IpV4Block"
  object_type = "ippool.IpV4Block"
  from = "10.1.1.1" 
  to = "10.1.1.3"
  }
  
  ip_v4_config {    
     object_type = "ippool.IpV4Config"
     class_id = "ippool.IpV4Config"
     gateway     = "10.1.1.1"
     netmask     = "255.0.0.0"
     primary_dns = "8.8.8.8"
   }
   
   organization {
     object_type = "organization.Organization"
     moid        = data.intersight_organization_organization.default.results.0.moid
   }
}
