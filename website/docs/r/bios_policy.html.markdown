---
subcategory: "bios"
layout: "intersight"
page_title: "Intersight: intersight_bios_policy"
description: |-
  Policy for setting BIOS tokens on the endpoint.
---

# Resource: intersight_bios_policy
Policy for setting BIOS tokens on the endpoint.
## Usage Example
### Resource Creation

```hcl
resource "intersight_bios_policy" "bios_policy1" {
  name                                  = "TEST_BIOS_POLICY"
  description                           = "Bios policy"
  acs_control_gpu1state                 = "disabled"
  acs_control_gpu2state                 = "disabled"
  acs_control_gpu3state                 = "disabled"
  acs_control_gpu4state                 = "disabled"
  acs_control_gpu5state                 = "disabled"
  acs_control_gpu6state                 = "disabled"
  acs_control_gpu7state                 = "disabled"
  acs_control_gpu8state                 = "disabled"
  acs_control_slot11state               = "disabled"
  acs_control_slot12state               = "disabled"
  acs_control_slot13state               = "disabled"
  acs_control_slot14state               = "disabled"
  cdn_support                           = "disabled"
  lom_port0state                        = "disabled"
  lom_port1state                        = "disabled"
  lom_port2state                        = "disabled"
  lom_port3state                        = "disabled"
  lom_ports_all_state                   = "disabled"
  pci_option_ro_ms                      = "disabled"
  pci_rom_clp                           = "disabled"
  slot10link_speed                      = "Auto"
  slot10state                           = "disabled"
  slot11link_speed                      = "Auto"
  slot11state                           = "disabled"
  slot12link_speed                      = "Auto"
  slot12state                           = "disabled"
  slot13state                           = "disabled"
  slot14state                           = "disabled"
  slot1link_speed                       = "Auto"
  slot1state                            = "disabled"
  slot2link_speed                       = "Auto"
  slot2state                            = "disabled"
  slot3link_speed                       = "GEN3"
  slot3state                            = "disabled"
  slot4link_speed                       = "Auto"
  slot4state                            = "disabled"
  slot5link_speed                       = "GEN2"
  slot5state                            = "disabled"
  slot6link_speed                       = "Auto"
  slot6state                            = "disabled"
  slot7link_speed                       = "GEN1"
  slot7state                            = "disabled"
  slot8link_speed                       = "Auto"
  slot8state                            = "disabled"
  slot9link_speed                       = "Auto"
  slot9state                            = "disabled"
  slot_flom_link_speed                  = "Auto"
  slot_front_nvme1link_speed            = "Auto"
  slot_front_nvme2link_speed            = "Auto"
  slot_front_slot5link_speed            = "Auto"
  slot_front_slot6link_speed            = "Auto"
  slot_gpu1state                        = "disabled"
  slot_gpu2state                        = "disabled"
  slot_gpu3state                        = "disabled"
  slot_gpu4state                        = "disabled"
  slot_gpu5state                        = "disabled"
  slot_gpu6state                        = "disabled"
  slot_gpu7state                        = "disabled"
  slot_gpu8state                        = "disabled"
  slot_hba_link_speed                   = "Auto"
  slot_hba_state                        = "disabled"
  slot_lom1link                         = "disabled"
  slot_lom2link                         = "disabled"
  slot_mezz_state                       = "disabled"
  slot_mlom_link_speed                  = "Auto"
  slot_mlom_state                       = "disabled"
  slot_mraid_link_speed                 = "Auto"
  slot_mraid_state                      = "disabled"
  slot_n10state                         = "disabled"
  slot_n11state                         = "disabled"
  slot_n12state                         = "disabled"
  slot_n13state                         = "disabled"
  slot_n14state                         = "disabled"
  slot_n15state                         = "disabled"
  slot_n16state                         = "disabled"
  slot_n17state                         = "disabled"
  slot_n18state                         = "disabled"
  slot_n19state                         = "disabled"
  slot_n1state                          = "disabled"
  slot_n20state                         = "disabled"
  slot_n21state                         = "disabled"
  slot_n22state                         = "disabled"
  slot_n23state                         = "disabled"
  slot_n24state                         = "disabled"
  slot_n2state                          = "disabled"
  slot_n3state                          = "disabled"
  slot_n4state                          = "disabled"
  slot_n5state                          = "disabled"
  slot_n6state                          = "disabled"
  slot_n7state                          = "disabled"
  slot_n8state                          = "disabled"
  slot_n9state                          = "disabled"
  slot_raid_link_speed                  = "Auto"
  slot_raid_state                       = "disabled"
  slot_rear_nvme1link_speed             = "Auto"
  slot_rear_nvme1state                  = "disabled"
  slot_rear_nvme2link_speed             = "Auto"
  slot_rear_nvme2state                  = "disabled"
  slot_rear_nvme3state                  = "disabled"
  slot_rear_nvme4state                  = "disabled"
  slot_rear_nvme5state                  = "disabled"
  slot_rear_nvme6state                  = "disabled"
  slot_rear_nvme7state                  = "disabled"
  slot_rear_nvme8state                  = "disabled"
  slot_riser1link_speed                 = "Auto"
  slot_riser1slot1link_speed            = "Auto"
  slot_riser1slot2link_speed            = "Auto"
  slot_riser1slot3link_speed            = "Auto"
  slot_riser2link_speed                 = "Auto"
  slot_riser2slot4link_speed            = "Auto"
  slot_riser2slot5link_speed            = "Auto"
  slot_riser2slot6link_speed            = "Auto"
  slot_sas_state                        = "disabled"
  slot_ssd_slot1link_speed              = "Auto"
  slot_ssd_slot2link_speed              = "Auto"
  adjacent_cache_line_prefetch          = "disabled"
  altitude                              = "300-m"
  auto_cc_state                         = "disabled"
  autonumous_cstate_enable              = "disabled"
  boot_performance_mode                 = "Max Performance"
  cbs_cmn_cpu_gen_downcore_ctrl         = "FOUR (2 + 2)"
  channel_inter_leave                   = "auto"
  closed_loop_therm_throtl              = "disabled"
  cmci_enable                           = "disabled"
  config_tdp                            = "disabled"
  core_multi_processing                 = "2"
  cpu_energy_performance                = "performance"
  cpu_frequency_floor                   = "disabled"
  cpu_performance                       = "custom"
  cpu_power_management                  = "custom"
  demand_scrub                          = "disabled"
  direct_cache_access                   = "disabled"
  dram_clock_throttling                 = "Auto"
  energy_efficient_turbo                = "disabled"
  eng_perf_tuning                       = "OS"
  enhanced_intel_speed_step_tech        = "disabled"
  epp_profile                           = "Power"
  execute_disable_bit                   = "disabled"
  extended_apic                         = "disabled"
  hardware_prefetch                     = "disabled"
  hwpm_enable                           = "HWPM Native Mode"
  imc_interleave                        = "1-way Interleave"
  intel_hyper_threading_tech            = "disabled"
  intel_speed_select                    = "Base"
  intel_turbo_boost_tech                = "disabled"
  intel_virtualization_technology       = "disabled"
  ioh_error_enable                      = "Yes"
  ip_prefetch                           = "disabled"
  kti_prefetch                          = "disabled"
  llc_prefetch                          = "disabled"
  memory_inter_leave                    = "disabled"
  package_cstate_limit                  = "Auto"
  patrol_scrub                          = "disabled"
  patrol_scrub_duration                 = "platform-default"
  pc_ie_ssd_hot_plug_support            = "disabled"
  processor_c1e                         = "disabled"
  processor_c3report                    = "disabled"
  processor_c6report                    = "disabled"
  processor_cstate                      = "disabled"
  pstate_coord_type                     = "HW ALL"
  pwr_perf_tuning                       = "bios"
  rank_inter_leave                      = "auto"
  single_pctl_enable                    = "Yes"
  smt_mode                              = "Auto"
  snc                                   = "disabled"
  streamer_prefetch                     = "disabled"
  svm_mode                              = "disabled"
  work_load_config                      = "Balanced"
  xpt_prefetch                          = "Auto"
  all_usb_devices                       = "disabled"
  legacy_usb_support                    = "disabled"
  make_device_non_bootable              = "disabled"
  pch_usb30mode                         = "disabled"
  usb_emul6064                          = "disabled"
  usb_port_front                        = "disabled"
  usb_port_internal                     = "disabled"
  usb_port_kvm                          = "disabled"
  usb_port_rear                         = "disabled"
  usb_port_sd_card                      = "disabled"
  usb_port_vmedia                       = "disabled"
  usb_xhci_support                      = "disabled"
  aspm_support                          = "Auto"
  ioh_resource                          = "IOH0 24k IOH1 40k"
  memory_mapped_io_above4gb             = "disabled"
  mmcfg_base                            = "2 GB"
  onboard10gbit_lom                     = "disabled"
  onboard_gbit_lom                      = "disabled"
  sr_iov                                = "disabled"
  vga_priority                          = "Onboard"
  assert_nmi_on_perr                    = "disabled"
  assert_nmi_on_serr                    = "disabled"
  baud_rate                             = "115200"
  cdn_enable                            = "disabled"
  cisco_adaptive_mem_training           = "disabled"
  cisco_debug_level                     = "Minimum"
  cisco_oprom_launch_optimization       = "disabled"
  console_redirection                   = "disabled"
  flow_control                          = "rts-cts"
  frb2enable                            = "disabled"
  legacy_os_redirection                 = "disabled"
  os_boot_watchdog_timer                = "disabled"
  os_boot_watchdog_timer_policy         = "power-off"
  os_boot_watchdog_timer_timeout        = "15-minutes"
  out_of_band_mgmt_port                 = "disabled"
  putty_key_pad                         = "LINUX"
  redirection_after_post                = "Always Enable"
  terminal_type                         = "vt100"
  ucsm_boot_order_rule                  = "Loose"
  bme_dma_mitigation                    = "disabled"
  cbs_cmn_gnb_nb_iommu                  = "disabled"
  cbs_cmn_mem_ctrl_bank_group_swap_ddr4 = "disabled"
  cbs_cmn_mem_map_bank_interleave_ddr4  = "Auto"
  cbs_df_cmn_mem_intlv                  = "Auto"
  cbs_df_cmn_mem_intlv_size             = "Auto"
  dcpmm_firmware_downgrade              = "disabled"
  smee                                  = "disabled"
  boot_option_num_retry                 = "13"
  boot_option_re_cool_down              = "90"
  boot_option_retry                     = "disabled"
  ipv6pxe                               = "disabled"
  onboard_scu_storage_support           = "disabled"
  onboard_scu_storage_sw_stack          = "Intel RSTe"
  pop_support                           = "disabled"
  psata                                 = "AHCI"
  sata_mode_select                      = "AHCI"
  vmd_enable                            = "disabled"
  cbs_cmn_cpu_cpb                       = "Auto"
  cbs_cmn_cpu_global_cstate_ctrl        = "Auto"
  cbs_cmn_cpu_l1stream_hw_prefetcher    = "Auto"
  cbs_cmn_cpu_l2stream_hw_prefetcher    = "Auto"
  cbs_cmn_determinism_slider            = "Auto"
  cbs_cmnc_tdp_ctl                      = "Auto"
  cke_low_policy                        = "auto"
  dram_refresh_rate                     = "2x"
  lv_ddr_mode                           = "auto"
  mirroring_mode                        = "inter-socket"
  numa_optimized                        = "disabled"
  select_memory_ras_configuration       = "mirror-mode-1lm"
  sparing_mode                          = "dimm-sparing"
  intel_vt_for_directed_io              = "disabled"
  intel_vtd_coherency_support           = "disabled"
  intel_vtd_interrupt_remapping         = "disabled"
  intel_vtd_pass_through_dma_support    = "disabled"
  intel_vtdats_support                  = "disabled"
  post_error_pause                      = "disabled"
  tpm_support                           = "disabled"
  qpi_link_frequency                    = "7.2-gt/s"
  qpi_snoop_mode                        = "auto"
  serial_port_aenable                   = "disabled"
  tpm_control                           = "disabled"
  txt_support                           = "disabled"
}
```
## Argument Reference
The following arguments are supported:
* `account_moid`:(string)(Computed) The Account ID for this managed object. 
* `acs_control_gpu1state`:(string) BIOS Token for setting ACS Control GPU-1 configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `enabled` - Enables the BIOS setting.* `disabled` - Disables the BIOS setting. 
* `acs_control_gpu2state`:(string) BIOS Token for setting ACS Control GPU-2 configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `enabled` - Enables the BIOS setting.* `disabled` - Disables the BIOS setting. 
* `acs_control_gpu3state`:(string) BIOS Token for setting ACS Control GPU-3 configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `enabled` - Enables the BIOS setting.* `disabled` - Disables the BIOS setting. 
* `acs_control_gpu4state`:(string) BIOS Token for setting ACS Control GPU-4 configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `enabled` - Enables the BIOS setting.* `disabled` - Disables the BIOS setting. 
* `acs_control_gpu5state`:(string) BIOS Token for setting ACS Control GPU-5 configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `enabled` - Enables the BIOS setting.* `disabled` - Disables the BIOS setting. 
* `acs_control_gpu6state`:(string) BIOS Token for setting ACS Control GPU-6 configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `enabled` - Enables the BIOS setting.* `disabled` - Disables the BIOS setting. 
* `acs_control_gpu7state`:(string) BIOS Token for setting ACS Control GPU-7 configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `enabled` - Enables the BIOS setting.* `disabled` - Disables the BIOS setting. 
* `acs_control_gpu8state`:(string) BIOS Token for setting ACS Control GPU-8 configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `enabled` - Enables the BIOS setting.* `disabled` - Disables the BIOS setting. 
* `acs_control_slot11state`:(string) BIOS Token for setting ACS Control Slot 11 configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `enabled` - Enables the BIOS setting.* `disabled` - Disables the BIOS setting. 
* `acs_control_slot12state`:(string) BIOS Token for setting ACS Control Slot 12 configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `enabled` - Enables the BIOS setting.* `disabled` - Disables the BIOS setting. 
* `acs_control_slot13state`:(string) BIOS Token for setting ACS Control Slot 13 configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `enabled` - Enables the BIOS setting.* `disabled` - Disables the BIOS setting. 
* `acs_control_slot14state`:(string) BIOS Token for setting ACS Control Slot 14 configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `enabled` - Enables the BIOS setting.* `disabled` - Disables the BIOS setting. 
* `adjacent_cache_line_prefetch`:(string) BIOS Token for setting Adjacent Cache Line Prefetcher configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `enabled` - Enables the BIOS setting.* `disabled` - Disables the BIOS setting. 
* `advanced_mem_test`:(string) BIOS Token for setting Advanced Memory Test configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `enabled` - Enables the BIOS setting.* `disabled` - Disables the BIOS setting. 
* `all_usb_devices`:(string) BIOS Token for setting All USB Devices configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `enabled` - Enables the BIOS setting.* `disabled` - Disables the BIOS setting. 
* `altitude`:(string) BIOS Token for setting Altitude configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `1500-m` - Value - 1500-m for configuring Altitude token.* `300-m` - Value - 300-m for configuring Altitude token.* `3000-m` - Value - 3000-m for configuring Altitude token.* `900-m` - Value - 900-m for configuring Altitude token.* `auto` - Value - auto for configuring Altitude token. 
* `ancestors`:(Array)(Computed) An array of relationships to moBaseMo resources. 
This complex property has following sub-properties:
  + `moid`:(string)(Computed) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string)(Computed) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `aspm_support`:(string) BIOS Token for setting ASPM Support configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `Auto` - Value - Auto for configuring AspmSupport token.* `Disabled` - Value - Disabled for configuring AspmSupport token.* `Force L0s` - Value - Force L0s for configuring AspmSupport token.* `L1 Only` - Value - L1 Only for configuring AspmSupport token. 
* `assert_nmi_on_perr`:(string) BIOS Token for setting Assert NMI on PERR configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `enabled` - Enables the BIOS setting.* `disabled` - Disables the BIOS setting. 
* `assert_nmi_on_serr`:(string) BIOS Token for setting Assert NMI on SERR configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `enabled` - Enables the BIOS setting.* `disabled` - Disables the BIOS setting. 
* `auto_cc_state`:(string) BIOS Token for setting Autonomous Core C-state configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `enabled` - Enables the BIOS setting.* `disabled` - Disables the BIOS setting. 
* `autonumous_cstate_enable`:(string) BIOS Token for setting CPU Autonomous Cstate configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `enabled` - Enables the BIOS setting.* `disabled` - Disables the BIOS setting. 
* `baud_rate`:(string) BIOS Token for setting Baud Rate configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `115200` - Value - 115200 for configuring BaudRate token.* `19200` - Value - 19200 for configuring BaudRate token.* `38400` - Value - 38400 for configuring BaudRate token.* `57600` - Value - 57600 for configuring BaudRate token.* `9600` - Value - 9600 for configuring BaudRate token. 
* `bme_dma_mitigation`:(string) BIOS Token for setting BME DMA Mitigation configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `enabled` - Enables the BIOS setting.* `disabled` - Disables the BIOS setting. 
* `boot_option_num_retry`:(string) BIOS Token for setting Number of Retries configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `13` - Value - 13 for configuring BootOptionNumRetry token.* `5` - Value - 5 for configuring BootOptionNumRetry token.* `Infinite` - Value - Infinite for configuring BootOptionNumRetry token. 
* `boot_option_re_cool_down`:(string) BIOS Token for setting Cool Down Time  (sec) configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `15` - Value - 15 for configuring BootOptionReCoolDown token.* `45` - Value - 45 for configuring BootOptionReCoolDown token.* `90` - Value - 90 for configuring BootOptionReCoolDown token. 
* `boot_option_retry`:(string) BIOS Token for setting Boot Option Retry configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `enabled` - Enables the BIOS setting.* `disabled` - Disables the BIOS setting. 
* `boot_performance_mode`:(string) BIOS Token for setting Boot Performance Mode configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `Max Efficient` - Value - Max Efficient for configuring BootPerformanceMode token.* `Max Performance` - Value - Max Performance for configuring BootPerformanceMode token.* `Set by Intel NM` - Value - Set by Intel NM for configuring BootPerformanceMode token. 
* `cbs_cmn_cpu_cpb`:(string) BIOS Token for setting Core Performance Boost configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `Auto` - Value - Auto for configuring CbsCmnCpuCpb token.* `disabled` - Value - disabled for configuring CbsCmnCpuCpb token. 
* `cbs_cmn_cpu_gen_downcore_ctrl`:(string) BIOS Token for setting Downcore Control configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `Auto` - Value - Auto for configuring CbsCmnCpuGenDowncoreCtrl token.* `FOUR (2 + 2)` - Value - FOUR (2 + 2) for configuring CbsCmnCpuGenDowncoreCtrl token.* `FOUR (4 + 0)` - Value - FOUR (4 + 0) for configuring CbsCmnCpuGenDowncoreCtrl token.* `SIX (3 + 3)` - Value - SIX (3 + 3) for configuring CbsCmnCpuGenDowncoreCtrl token.* `THREE (3 + 0)` - Value - THREE (3 + 0) for configuring CbsCmnCpuGenDowncoreCtrl token.* `TWO (1 + 1)` - Value - TWO (1 + 1) for configuring CbsCmnCpuGenDowncoreCtrl token.* `TWO (2 + 0)` - Value - TWO (2 + 0) for configuring CbsCmnCpuGenDowncoreCtrl token. 
* `cbs_cmn_cpu_global_cstate_ctrl`:(string) BIOS Token for setting Global C-state Control configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `Auto` - Value - Auto for configuring CbsCmnCpuGlobalCstateCtrl token.* `disabled` - Value - disabled for configuring CbsCmnCpuGlobalCstateCtrl token.* `enabled` - Value - enabled for configuring CbsCmnCpuGlobalCstateCtrl token. 
* `cbs_cmn_cpu_l1stream_hw_prefetcher`:(string) BIOS Token for setting L1 Stream HW Prefetcher configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `Auto` - Value - Auto for configuring CbsCmnCpuL1streamHwPrefetcher token.* `disabled` - Value - disabled for configuring CbsCmnCpuL1streamHwPrefetcher token.* `enabled` - Value - enabled for configuring CbsCmnCpuL1streamHwPrefetcher token. 
* `cbs_cmn_cpu_l2stream_hw_prefetcher`:(string) BIOS Token for setting L2 Stream HW Prefetcher configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `Auto` - Value - Auto for configuring CbsCmnCpuL2streamHwPrefetcher token.* `disabled` - Value - disabled for configuring CbsCmnCpuL2streamHwPrefetcher token.* `enabled` - Value - enabled for configuring CbsCmnCpuL2streamHwPrefetcher token. 
* `cbs_cmn_determinism_slider`:(string) BIOS Token for setting Determinism Slider configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `Auto` - Value - Auto for configuring CbsCmnDeterminismSlider token.* `Performance` - Value - Performance for configuring CbsCmnDeterminismSlider token.* `Power` - Value - Power for configuring CbsCmnDeterminismSlider token. 
* `cbs_cmn_gnb_nb_iommu`:(string) BIOS Token for setting IOMMU configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `Auto` - Value - Auto for configuring CbsCmnGnbNbIommu token.* `disabled` - Value - disabled for configuring CbsCmnGnbNbIommu token.* `enabled` - Value - enabled for configuring CbsCmnGnbNbIommu token. 
* `cbs_cmn_mem_ctrl_bank_group_swap_ddr4`:(string) BIOS Token for setting Bank Group Swap configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `Auto` - Value - Auto for configuring CbsCmnMemCtrlBankGroupSwapDdr4 token.* `disabled` - Value - disabled for configuring CbsCmnMemCtrlBankGroupSwapDdr4 token.* `enabled` - Value - enabled for configuring CbsCmnMemCtrlBankGroupSwapDdr4 token. 
* `cbs_cmn_mem_map_bank_interleave_ddr4`:(string) BIOS Token for setting Chipset Interleave configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `Auto` - Value - Auto for configuring CbsCmnMemMapBankInterleaveDdr4 token.* `disabled` - Value - disabled for configuring CbsCmnMemMapBankInterleaveDdr4 token. 
* `cbs_cmnc_tdp_ctl`:(string) BIOS Token for setting cTDP Control configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `Auto` - Value - Auto for configuring CbsCmncTdpCtl token.* `Manual` - Value - Manual for configuring CbsCmncTdpCtl token. 
* `cbs_df_cmn_mem_intlv`:(string) BIOS Token for setting AMD Memory Interleaving configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `Auto` - Value - Auto for configuring CbsDfCmnMemIntlv token.* `Channel` - Value - Channel for configuring CbsDfCmnMemIntlv token.* `Die` - Value - Die for configuring CbsDfCmnMemIntlv token.* `None` - Value - None for configuring CbsDfCmnMemIntlv token.* `Socket` - Value - Socket for configuring CbsDfCmnMemIntlv token. 
* `cbs_df_cmn_mem_intlv_size`:(string) BIOS Token for setting AMD Memory Interleaving Size configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `1 KB` - Value - 1 KiB for configuring CbsDfCmnMemIntlvSize token.* `2 KB` - Value - 2 KiB for configuring CbsDfCmnMemIntlvSize token.* `256 Bytes` - Value - 256 Bytes for configuring CbsDfCmnMemIntlvSize token.* `512 Bytes` - Value - 512 Bytes for configuring CbsDfCmnMemIntlvSize token.* `Auto` - Value - Auto for configuring CbsDfCmnMemIntlvSize token. 
* `cdn_enable`:(string) BIOS Token for setting Consistent Device Naming configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `enabled` - Enables the BIOS setting.* `disabled` - Disables the BIOS setting. 
* `cdn_support`:(string) BIOS Token for setting CDN Support for LOM configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `disabled` - Value - disabled for configuring CdnSupport token.* `enabled` - Value - enabled for configuring CdnSupport token.* `LOMs Only` - Value - LOMs Only for configuring CdnSupport token. 
* `channel_inter_leave`:(string) BIOS Token for setting Channel Interleaving configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `1-way` - Value - 1-way for configuring ChannelInterLeave token.* `2-way` - Value - 2-way for configuring ChannelInterLeave token.* `3-way` - Value - 3-way for configuring ChannelInterLeave token.* `4-way` - Value - 4-way for configuring ChannelInterLeave token.* `auto` - Value - auto for configuring ChannelInterLeave token. 
* `cisco_adaptive_mem_training`:(string) BIOS Token for setting Adaptive Memory Training configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `enabled` - Enables the BIOS setting.* `disabled` - Disables the BIOS setting. 
* `cisco_debug_level`:(string) BIOS Token for setting BIOS Techlog Level configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `Maximum` - Value - Maximum for configuring CiscoDebugLevel token.* `Minimum` - Value - Minimum for configuring CiscoDebugLevel token.* `Normal` - Value - Normal for configuring CiscoDebugLevel token. 
* `cisco_oprom_launch_optimization`:(string) BIOS Token for setting OptionROM Launch Optimization configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `enabled` - Enables the BIOS setting.* `disabled` - Disables the BIOS setting. 
* `cke_low_policy`:(string) BIOS Token for setting CKE Low Policy configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `auto` - Value - auto for configuring CkeLowPolicy token.* `disabled` - Value - disabled for configuring CkeLowPolicy token.* `fast` - Value - fast for configuring CkeLowPolicy token.* `slow` - Value - slow for configuring CkeLowPolicy token. 
* `closed_loop_therm_throtl`:(string) BIOS Token for setting Closed Loop Therm Throt configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `enabled` - Enables the BIOS setting.* `disabled` - Disables the BIOS setting. 
* `cmci_enable`:(string) BIOS Token for setting Processor CMCI configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `enabled` - Enables the BIOS setting.* `disabled` - Disables the BIOS setting. 
* `config_tdp`:(string) BIOS Token for setting Config TDP configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `enabled` - Enables the BIOS setting.* `disabled` - Disables the BIOS setting. 
* `config_tdp_level`:(string) BIOS Token for setting Configurable TDP Level configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `Level 1` - Value - Level 1 for configuring ConfigTdpLevel token.* `Level 2` - Value - Level 2 for configuring ConfigTdpLevel token.* `Normal` - Value - Normal for configuring ConfigTdpLevel token. 
* `console_redirection`:(string) BIOS Token for setting Console Redirection configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `com-0` - Value - com-0 for configuring ConsoleRedirection token.* `com-1` - Value - com-1 for configuring ConsoleRedirection token.* `disabled` - Value - disabled for configuring ConsoleRedirection token.* `enabled` - Value - enabled for configuring ConsoleRedirection token.* `serial-port-a` - Value - serial-port-a for configuring ConsoleRedirection token. 
* `core_multi_processing`:(string) BIOS Token for setting Core MultiProcessing configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `1` - Value - 1 for configuring CoreMultiProcessing token.* `10` - Value - 10 for configuring CoreMultiProcessing token.* `11` - Value - 11 for configuring CoreMultiProcessing token.* `12` - Value - 12 for configuring CoreMultiProcessing token.* `13` - Value - 13 for configuring CoreMultiProcessing token.* `14` - Value - 14 for configuring CoreMultiProcessing token.* `15` - Value - 15 for configuring CoreMultiProcessing token.* `16` - Value - 16 for configuring CoreMultiProcessing token.* `17` - Value - 17 for configuring CoreMultiProcessing token.* `18` - Value - 18 for configuring CoreMultiProcessing token.* `19` - Value - 19 for configuring CoreMultiProcessing token.* `2` - Value - 2 for configuring CoreMultiProcessing token.* `20` - Value - 20 for configuring CoreMultiProcessing token.* `21` - Value - 21 for configuring CoreMultiProcessing token.* `22` - Value - 22 for configuring CoreMultiProcessing token.* `23` - Value - 23 for configuring CoreMultiProcessing token.* `24` - Value - 24 for configuring CoreMultiProcessing token.* `25` - Value - 25 for configuring CoreMultiProcessing token.* `26` - Value - 26 for configuring CoreMultiProcessing token.* `27` - Value - 27 for configuring CoreMultiProcessing token.* `28` - Value - 28 for configuring CoreMultiProcessing token.* `3` - Value - 3 for configuring CoreMultiProcessing token.* `4` - Value - 4 for configuring CoreMultiProcessing token.* `5` - Value - 5 for configuring CoreMultiProcessing token.* `6` - Value - 6 for configuring CoreMultiProcessing token.* `7` - Value - 7 for configuring CoreMultiProcessing token.* `8` - Value - 8 for configuring CoreMultiProcessing token.* `9` - Value - 9 for configuring CoreMultiProcessing token.* `all` - Value - all for configuring CoreMultiProcessing token. 
* `cpu_energy_performance`:(string) BIOS Token for setting Energy Performance configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `balanced-energy` - Value - balanced-energy for configuring CpuEnergyPerformance token.* `balanced-performance` - Value - balanced-performance for configuring CpuEnergyPerformance token.* `balanced-power` - Value - balanced-power for configuring CpuEnergyPerformance token.* `energy-efficient` - Value - energy-efficient for configuring CpuEnergyPerformance token.* `performance` - Value - performance for configuring CpuEnergyPerformance token.* `power` - Value - power for configuring CpuEnergyPerformance token. 
* `cpu_frequency_floor`:(string) BIOS Token for setting Frequency Floor Override configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `enabled` - Enables the BIOS setting.* `disabled` - Disables the BIOS setting. 
* `cpu_performance`:(string) BIOS Token for setting CPU Performance configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `custom` - Value - custom for configuring CpuPerformance token.* `enterprise` - Value - enterprise for configuring CpuPerformance token.* `high-throughput` - Value - high-throughput for configuring CpuPerformance token.* `hpc` - Value - hpc for configuring CpuPerformance token. 
* `cpu_power_management`:(string) BIOS Token for setting Power Technology configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `custom` - Value - custom for configuring CpuPowerManagement token.* `disabled` - Value - disabled for configuring CpuPowerManagement token.* `energy-efficient` - Value - energy-efficient for configuring CpuPowerManagement token.* `performance` - Value - performance for configuring CpuPowerManagement token. 
* `cr_qos`:(string) BIOS Token for setting CR QoS configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `Disabled` - Value - Disabled for configuring CrQos token.* `Recipe 1` - Value - Recipe 1 for configuring CrQos token.* `Recipe 2` - Value - Recipe 2 for configuring CrQos token.* `Recipe 3` - Value - Recipe 3 for configuring CrQos token. 
* `create_time`:(string)(Computed) The time when this managed object was created. 
* `crfastgo_config`:(string) BIOS Token for setting CR FastGo Config configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `Auto` - Value - Auto for configuring CrfastgoConfig token.* `Default` - Value - Default for configuring CrfastgoConfig token.* `Option 1` - Value - Option 1 for configuring CrfastgoConfig token.* `Option 2` - Value - Option 2 for configuring CrfastgoConfig token.* `Option 3` - Value - Option 3 for configuring CrfastgoConfig token.* `Option 4` - Value - Option 4 for configuring CrfastgoConfig token.* `Option 5` - Value - Option 5 for configuring CrfastgoConfig token. 
* `dcpmm_firmware_downgrade`:(string) BIOS Token for setting DCPMM Firmware Downgrade configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `enabled` - Enables the BIOS setting.* `disabled` - Disables the BIOS setting. 
* `demand_scrub`:(string) BIOS Token for setting Demand Scrub configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `enabled` - Enables the BIOS setting.* `disabled` - Disables the BIOS setting. 
* `description`:(string) Description of the policy. 
* `direct_cache_access`:(string) BIOS Token for setting Direct Cache Access Support configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `auto` - Value - auto for configuring DirectCacheAccess token.* `disabled` - Value - disabled for configuring DirectCacheAccess token.* `enabled` - Value - enabled for configuring DirectCacheAccess token. 
* `domain_group_moid`:(string)(Computed) The DomainGroup ID for this managed object. 
* `dram_clock_throttling`:(string) BIOS Token for setting DRAM Clock Throttling configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `Auto` - Value - Auto for configuring DramClockThrottling token.* `Balanced` - Value - Balanced for configuring DramClockThrottling token.* `Energy Efficient` - Value - Energy Efficient for configuring DramClockThrottling token.* `Performance` - Value - Performance for configuring DramClockThrottling token. 
* `dram_refresh_rate`:(string) BIOS Token for setting DRAM Refresh Rate configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `1x` - Value - 1x for configuring DramRefreshRate token.* `2x` - Value - 2x for configuring DramRefreshRate token.* `3x` - Value - 3x for configuring DramRefreshRate token.* `4x` - Value - 4x for configuring DramRefreshRate token.* `Auto` - Value - Auto for configuring DramRefreshRate token. 
* `dram_sw_thermal_throttling`:(string) BIOS Token for setting DRAM SW Thermal Throttling configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `disabled` - Value - disabled for configuring DramSwThermalThrottling token.* `enabled` - Value - enabled for configuring DramSwThermalThrottling token. 
* `enable_clock_spread_spec`:(string) BIOS Token for setting External SSC Enable configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `enabled` - Enables the BIOS setting.* `disabled` - Disables the BIOS setting. 
* `energy_efficient_turbo`:(string) BIOS Token for setting Energy Efficient Turbo configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `enabled` - Enables the BIOS setting.* `disabled` - Disables the BIOS setting. 
* `eng_perf_tuning`:(string) BIOS Token for setting Energy Performance Tuning configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `BIOS` - Value - BIOS for configuring EngPerfTuning token.* `OS` - Value - OS for configuring EngPerfTuning token. 
* `enhanced_intel_speed_step_tech`:(string) BIOS Token for setting Enhanced Intel Speedstep (R) Technology configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `enabled` - Enables the BIOS setting.* `disabled` - Disables the BIOS setting. 
* `epp_enable`:(string) BIOS Token for setting Processor EPP Enable configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `enabled` - Enables the BIOS setting.* `disabled` - Disables the BIOS setting. 
* `epp_profile`:(string) BIOS Token for setting EPP Profile configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `Balanced Performance` - Value - Balanced Performance for configuring EppProfile token.* `Balanced Power` - Value - Balanced Power for configuring EppProfile token.* `Performance` - Value - Performance for configuring EppProfile token.* `Power` - Value - Power for configuring EppProfile token. 
* `execute_disable_bit`:(string) BIOS Token for setting Execute Disable Bit configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `enabled` - Enables the BIOS setting.* `disabled` - Disables the BIOS setting. 
* `extended_apic`:(string) BIOS Token for setting Local X2 Apic configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `disabled` - Value - disabled for configuring ExtendedApic token.* `enabled` - Value - enabled for configuring ExtendedApic token.* `X2APIC` - Value - X2APIC for configuring ExtendedApic token.* `XAPIC` - Value - XAPIC for configuring ExtendedApic token. 
* `flow_control`:(string) BIOS Token for setting Flow Control configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `none` - Value - none for configuring FlowControl token.* `rts-cts` - Value - rts-cts for configuring FlowControl token. 
* `frb2enable`:(string) BIOS Token for setting FRB-2 Timer configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `enabled` - Enables the BIOS setting.* `disabled` - Disables the BIOS setting. 
* `hardware_prefetch`:(string) BIOS Token for setting Hardware Prefetcher configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `enabled` - Enables the BIOS setting.* `disabled` - Disables the BIOS setting. 
* `hwpm_enable`:(string) BIOS Token for setting CPU Hardware Power Management configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `Disabled` - Value - Disabled for configuring HwpmEnable token.* `HWPM Native Mode` - Value - HWPM Native Mode for configuring HwpmEnable token.* `HWPM OOB Mode` - Value - HWPM OOB Mode for configuring HwpmEnable token.* `NATIVE MODE` - Value - NATIVE MODE for configuring HwpmEnable token.* `Native Mode with no Legacy` - Value - Native Mode with no Legacy for configuring HwpmEnable token.* `OOB MODE` - Value - OOB MODE for configuring HwpmEnable token. 
* `imc_interleave`:(string) BIOS Token for setting IMC Interleaving configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `1-way Interleave` - Value - 1-way Interleave for configuring ImcInterleave token.* `2-way Interleave` - Value - 2-way Interleave for configuring ImcInterleave token.* `Auto` - Value - Auto for configuring ImcInterleave token. 
* `intel_hyper_threading_tech`:(string) BIOS Token for setting Intel HyperThreading Tech configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `enabled` - Enables the BIOS setting.* `disabled` - Disables the BIOS setting. 
* `intel_speed_select`:(string) BIOS Token for setting Intel Speed Select configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `Base` - Value - Base for configuring IntelSpeedSelect token.* `Config 1` - Value - Config 1 for configuring IntelSpeedSelect token.* `Config 2` - Value - Config 2 for configuring IntelSpeedSelect token. 
* `intel_turbo_boost_tech`:(string) BIOS Token for setting Intel Turbo Boost Tech configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `enabled` - Enables the BIOS setting.* `disabled` - Disables the BIOS setting. 
* `intel_virtualization_technology`:(string) BIOS Token for setting Intel (R) VT configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `enabled` - Enables the BIOS setting.* `disabled` - Disables the BIOS setting. 
* `intel_vt_for_directed_io`:(string) BIOS Token for setting Intel VT for Directed IO configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `enabled` - Enables the BIOS setting.* `disabled` - Disables the BIOS setting. 
* `intel_vtd_coherency_support`:(string) BIOS Token for setting Intel (R) VT-d Coherency Support configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `enabled` - Enables the BIOS setting.* `disabled` - Disables the BIOS setting. 
* `intel_vtd_interrupt_remapping`:(string) BIOS Token for setting Intel (R) VT-d Interrupt Remapping configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `enabled` - Enables the BIOS setting.* `disabled` - Disables the BIOS setting. 
* `intel_vtd_pass_through_dma_support`:(string) BIOS Token for setting Intel (R) VT-d PassThrough DMA Support configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `enabled` - Enables the BIOS setting.* `disabled` - Disables the BIOS setting. 
* `intel_vtdats_support`:(string) BIOS Token for setting Intel VTD ATS Support configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `enabled` - Enables the BIOS setting.* `disabled` - Disables the BIOS setting. 
* `ioh_error_enable`:(string) BIOS Token for setting IIO Error Enable configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `No` - Value - No for configuring IohErrorEnable token.* `Yes` - Value - Yes for configuring IohErrorEnable token. 
* `ioh_resource`:(string) BIOS Token for setting IOH Resource Allocation configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `IOH0 24k IOH1 40k` - Value - IOH0 24k IOH1 40k for configuring IohResource token.* `IOH0 32k IOH1 32k` - Value - IOH0 32k IOH1 32k for configuring IohResource token.* `IOH0 40k IOH1 24k` - Value - IOH0 40k IOH1 24k for configuring IohResource token.* `IOH0 48k IOH1 16k` - Value - IOH0 48k IOH1 16k for configuring IohResource token.* `IOH0 56k IOH1 8k` - Value - IOH0 56k IOH1 8k for configuring IohResource token. 
* `ip_prefetch`:(string) BIOS Token for setting DCU IP Prefetcher configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `enabled` - Enables the BIOS setting.* `disabled` - Disables the BIOS setting. 
* `ipv4http`:(string) BIOS Token for setting IPV4 HTTP Support configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `enabled` - Enables the BIOS setting.* `disabled` - Disables the BIOS setting. 
* `ipv4pxe`:(string) BIOS Token for setting IPv4 PXE Support configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `enabled` - Enables the BIOS setting.* `disabled` - Disables the BIOS setting. 
* `ipv6http`:(string) BIOS Token for setting IPV6 HTTP Support configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `enabled` - Enables the BIOS setting.* `disabled` - Disables the BIOS setting. 
* `ipv6pxe`:(string) BIOS Token for setting IPV6 PXE Support configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `enabled` - Enables the BIOS setting.* `disabled` - Disables the BIOS setting. 
* `kti_prefetch`:(string) BIOS Token for setting KTI Prefetch configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `enabled` - Enables the BIOS setting.* `disabled` - Disables the BIOS setting. 
* `legacy_os_redirection`:(string) BIOS Token for setting Legacy OS Redirection configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `enabled` - Enables the BIOS setting.* `disabled` - Disables the BIOS setting. 
* `legacy_usb_support`:(string) BIOS Token for setting Legacy USB Support configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `auto` - Value - auto for configuring LegacyUsbSupport token.* `disabled` - Value - disabled for configuring LegacyUsbSupport token.* `enabled` - Value - enabled for configuring LegacyUsbSupport token. 
* `llc_prefetch`:(string) BIOS Token for setting LLC Prefetch configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `enabled` - Enables the BIOS setting.* `disabled` - Disables the BIOS setting. 
* `lom_port0state`:(string) BIOS Token for setting LOM Port 0 OptionROM configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `disabled` - Value - disabled for configuring LomPort0state token.* `enabled` - Value - enabled for configuring LomPort0state token.* `Legacy Only` - Value - Legacy Only for configuring LomPort0state token.* `UEFI Only` - Value - UEFI Only for configuring LomPort0state token. 
* `lom_port1state`:(string) BIOS Token for setting LOM Port 1 OptionROM configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `disabled` - Value - disabled for configuring LomPort1state token.* `enabled` - Value - enabled for configuring LomPort1state token.* `Legacy Only` - Value - Legacy Only for configuring LomPort1state token.* `UEFI Only` - Value - UEFI Only for configuring LomPort1state token. 
* `lom_port2state`:(string) BIOS Token for setting LOM Port 2 OptionROM configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `disabled` - Value - disabled for configuring LomPort2state token.* `enabled` - Value - enabled for configuring LomPort2state token.* `Legacy Only` - Value - Legacy Only for configuring LomPort2state token.* `UEFI Only` - Value - UEFI Only for configuring LomPort2state token. 
* `lom_port3state`:(string) BIOS Token for setting LOM Port 3 OptionROM configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `disabled` - Value - disabled for configuring LomPort3state token.* `enabled` - Value - enabled for configuring LomPort3state token.* `Legacy Only` - Value - Legacy Only for configuring LomPort3state token.* `UEFI Only` - Value - UEFI Only for configuring LomPort3state token. 
* `lom_ports_all_state`:(string) BIOS Token for setting All Onboard LOM Ports configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `enabled` - Enables the BIOS setting.* `disabled` - Disables the BIOS setting. 
* `lv_ddr_mode`:(string) BIOS Token for setting Low Voltage DDR Mode configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `auto` - Value - auto for configuring LvDdrMode token.* `performance-mode` - Value - performance-mode for configuring LvDdrMode token.* `power-saving-mode` - Value - power-saving-mode for configuring LvDdrMode token. 
* `make_device_non_bootable`:(string) BIOS Token for setting Make Device Non Bootable configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `enabled` - Enables the BIOS setting.* `disabled` - Disables the BIOS setting. 
* `memory_inter_leave`:(string) BIOS Token for setting Intel Memory Interleaving configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `1 Way Node Interleave` - Value - 1 Way Node Interleave for configuring MemoryInterLeave token.* `2 Way Node Interleave` - Value - 2 Way Node Interleave for configuring MemoryInterLeave token.* `4 Way Node Interleave` - Value - 4 Way Node Interleave for configuring MemoryInterLeave token.* `8 Way Node Interleave` - Value - 8 Way Node Interleave for configuring MemoryInterLeave token.* `disabled` - Value - disabled for configuring MemoryInterLeave token.* `enabled` - Value - enabled for configuring MemoryInterLeave token. 
* `memory_mapped_io_above4gb`:(string) BIOS Token for setting Memory Mapped IO above 4GiB configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `enabled` - Enables the BIOS setting.* `disabled` - Disables the BIOS setting. 
* `memory_refresh_rate`:(string) BIOS Token for setting Memory Refresh Rate configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `1x Refresh` - Value - 1x Refresh for configuring MemoryRefreshRate token.* `2x Refresh` - Value - 2x Refresh for configuring MemoryRefreshRate token. 
* `memory_size_limit`:(string) BIOS Token for setting Memory Size Limit in GiB configuration (0-65535). 
* `memory_thermal_throttling`:(string) BIOS Token for setting Memory Thermal Throttling Mode configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `CLTT with PECI` - Value - CLTT with PECI for configuring MemoryThermalThrottling token.* `Disabled` - Value - Disabled for configuring MemoryThermalThrottling token. 
* `mirroring_mode`:(string) BIOS Token for setting Mirroring Mode configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `inter-socket` - Value - inter-socket for configuring MirroringMode token.* `intra-socket` - Value - intra-socket for configuring MirroringMode token. 
* `mmcfg_base`:(string) BIOS Token for setting MMCFG BASE configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `1 GB` - Value - 1 GiB for configuring MmcfgBase token.* `2 GB` - Value - 2 GiB for configuring MmcfgBase token.* `2.5 GB` - Value - 2.5 GiB for configuring MmcfgBase token.* `3 GB` - Value - 3 GiB for configuring MmcfgBase token.* `Auto` - Value - Auto for configuring MmcfgBase token. 
* `mod_time`:(string)(Computed) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete policy. 
* `network_stack`:(string) BIOS Token for setting Network Stack configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `enabled` - Enables the BIOS setting.* `disabled` - Disables the BIOS setting. 
* `numa_optimized`:(string) BIOS Token for setting NUMA Optimized configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `enabled` - Enables the BIOS setting.* `disabled` - Disables the BIOS setting. 
* `nvmdimm_perform_config`:(string) BIOS Token for setting NVM Performance Setting configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `BW Optimized` - Value - BW Optimized for configuring NvmdimmPerformConfig token.* `Balanced Profile` - Value - Balanced Profile for configuring NvmdimmPerformConfig token.* `Latency Optimized` - Value - Latency Optimized for configuring NvmdimmPerformConfig token. 
* `onboard10gbit_lom`:(string) BIOS Token for setting Onboard 10Gbit LOM configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `enabled` - Enables the BIOS setting.* `disabled` - Disables the BIOS setting. 
* `onboard_gbit_lom`:(string) BIOS Token for setting Onboard Gbit LOM configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `enabled` - Enables the BIOS setting.* `disabled` - Disables the BIOS setting. 
* `onboard_scu_storage_support`:(string) BIOS Token for setting Onboard SCU Storage Support configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `enabled` - Enables the BIOS setting.* `disabled` - Disables the BIOS setting. 
* `onboard_scu_storage_sw_stack`:(string) BIOS Token for setting Onboard SCU Storage SW Stack configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `Intel RSTe` - Value - Intel RSTe for configuring OnboardScuStorageSwStack token.* `LSI SW RAID` - Value - LSI SW RAID for configuring OnboardScuStorageSwStack token. 
* `organization`:(HashMap) - A reference to a organizationOrganization resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string)(Computed) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string)(Computed) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `os_boot_watchdog_timer`:(string) BIOS Token for setting OS Boot Watchdog Timer configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `enabled` - Enables the BIOS setting.* `disabled` - Disables the BIOS setting. 
* `os_boot_watchdog_timer_policy`:(string) BIOS Token for setting OS Boot Watchdog Timer Policy configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `do-nothing` - Value - do-nothing for configuring OsBootWatchdogTimerPolicy token.* `power-off` - Value - power-off for configuring OsBootWatchdogTimerPolicy token.* `reset` - Value - reset for configuring OsBootWatchdogTimerPolicy token. 
* `os_boot_watchdog_timer_timeout`:(string) BIOS Token for setting OS Boot Watchdog Timer Timeout configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `10-minutes` - Value - 10-minutes for configuring OsBootWatchdogTimerTimeout token.* `15-minutes` - Value - 15-minutes for configuring OsBootWatchdogTimerTimeout token.* `20-minutes` - Value - 20-minutes for configuring OsBootWatchdogTimerTimeout token.* `5-minutes` - Value - 5-minutes for configuring OsBootWatchdogTimerTimeout token. 
* `out_of_band_mgmt_port`:(string) BIOS Token for setting Out-of-Band Mgmt Port configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `enabled` - Enables the BIOS setting.* `disabled` - Disables the BIOS setting. 
* `owners`:
                (Array of schema.TypeString) -(Computed)
* `package_cstate_limit`:(string) BIOS Token for setting Package C State Limit configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `Auto` - Value - Auto for configuring PackageCstateLimit token.* `C0 C1 State` - Value - C0 C1 State for configuring PackageCstateLimit token.* `C0/C1` - Value - C0/C1 for configuring PackageCstateLimit token.* `C2` - Value - C2 for configuring PackageCstateLimit token.* `C6 Non Retention` - Value - C6 Non Retention for configuring PackageCstateLimit token.* `C6 Retention` - Value - C6 Retention for configuring PackageCstateLimit token.* `No Limit` - Value - No Limit for configuring PackageCstateLimit token. 
* `panic_high_watermark`:(string) BIOS Token for setting Panic and High Watermark configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `High` - Value - High for configuring PanicHighWatermark token.* `Low` - Value - Low for configuring PanicHighWatermark token. 
* `parent`:(HashMap) -(Computed) A reference to a moBaseMo resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string)(Computed) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string)(Computed) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `partial_mirror_mode_config`:(string) BIOS Token for setting Partial Memory Mirror Mode configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `disabled` - Value - disabled for configuring PartialMirrorModeConfig token.* `Percentage` - Value - Percentage for configuring PartialMirrorModeConfig token.* `Value in GB` - Value - Value in GiB for configuring PartialMirrorModeConfig token. 
* `partial_mirror_percent`:(string) BIOS Token for setting Partial Mirror Percentage configuration (0.00-50.00). 
* `partial_mirror_value1`:(string) BIOS Token for setting Partial Mirror1 Size in GiB configuration (0-65535). 
* `partial_mirror_value2`:(string) BIOS Token for setting Partial Mirror2 Size in GiB configuration (0-65535). 
* `partial_mirror_value3`:(string) BIOS Token for setting Partial Mirror3 Size in GiB configuration (0-65535). 
* `partial_mirror_value4`:(string) BIOS Token for setting Partial Mirror4 Size in GiB configuration (0-65535). 
* `patrol_scrub`:(string) BIOS Token for setting Patrol Scrub configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `enabled` - Enables the BIOS setting.* `disabled` - Disables the BIOS setting. 
* `patrol_scrub_duration`:(string) BIOS Token for setting Patrol Scrub Interval configuration (5-23). 
* `pc_ie_ras_support`:(string) BIOS Token for setting PCIe RAS Support configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `enabled` - Enables the BIOS setting.* `disabled` - Disables the BIOS setting. 
* `pc_ie_ssd_hot_plug_support`:(string) BIOS Token for setting NVMe SSD Hot-Plug Support configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `enabled` - Enables the BIOS setting.* `disabled` - Disables the BIOS setting. 
* `pch_usb30mode`:(string) BIOS Token for setting xHCI Mode configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `enabled` - Enables the BIOS setting.* `disabled` - Disables the BIOS setting. 
* `pci_option_ro_ms`:(string) BIOS Token for setting All PCIe Slots OptionROM configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `disabled` - Value - disabled for configuring PciOptionRoMs token.* `enabled` - Value - enabled for configuring PciOptionRoMs token.* `Legacy Only` - Value - Legacy Only for configuring PciOptionRoMs token.* `UEFI Only` - Value - UEFI Only for configuring PciOptionRoMs token. 
* `pci_rom_clp`:(string) BIOS Token for setting PCI ROM CLP configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `enabled` - Enables the BIOS setting.* `disabled` - Disables the BIOS setting. 
* `pcie_ari_support`:(string) BIOS Token for setting PCIe ARI Support configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `Auto` - Value - Auto for configuring PcieAriSupport token.* `disabled` - Value - disabled for configuring PcieAriSupport token.* `enabled` - Value - enabled for configuring PcieAriSupport token. 
* `pcie_pll_ssc`:(string) BIOS Token for setting PCIe PLL SSC configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `Auto` - Value - Auto for configuring PciePllSsc token.* `Disabled` - Value - Disabled for configuring PciePllSsc token.* `ZeroPointFive` - Value - ZeroPointFive for configuring PciePllSsc token. 
* `pcie_slot_mstorraid_option_rom`:(string) BIOS Token for setting PCIe Slot MSTOR RAID OptionROM configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `enabled` - Enables the BIOS setting.* `disabled` - Disables the BIOS setting. 
* `pcie_slot_nvme1link_speed`:(string) BIOS Token for setting NVME-1 Link Speed configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `Auto` - Value - Auto for configuring PcieSlotNvme1linkSpeed token.* `Disabled` - Value - Disabled for configuring PcieSlotNvme1linkSpeed token.* `GEN1` - Value - GEN1 for configuring PcieSlotNvme1linkSpeed token.* `GEN2` - Value - GEN2 for configuring PcieSlotNvme1linkSpeed token.* `GEN3` - Value - GEN3 for configuring PcieSlotNvme1linkSpeed token. 
* `pcie_slot_nvme1option_rom`:(string) BIOS Token for setting NVME-1 OptionROM configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `enabled` - Enables the BIOS setting.* `disabled` - Disables the BIOS setting. 
* `pcie_slot_nvme2link_speed`:(string) BIOS Token for setting NVME-2 Link Speed configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `Auto` - Value - Auto for configuring PcieSlotNvme2linkSpeed token.* `Disabled` - Value - Disabled for configuring PcieSlotNvme2linkSpeed token.* `GEN1` - Value - GEN1 for configuring PcieSlotNvme2linkSpeed token.* `GEN2` - Value - GEN2 for configuring PcieSlotNvme2linkSpeed token.* `GEN3` - Value - GEN3 for configuring PcieSlotNvme2linkSpeed token. 
* `pcie_slot_nvme2option_rom`:(string) BIOS Token for setting NVME-2 OptionROM configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `enabled` - Enables the BIOS setting.* `disabled` - Disables the BIOS setting. 
* `pcie_slot_nvme3link_speed`:(string) BIOS Token for setting NVME-3 Link Speed configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `Auto` - Value - Auto for configuring PcieSlotNvme3linkSpeed token.* `Disabled` - Value - Disabled for configuring PcieSlotNvme3linkSpeed token.* `GEN1` - Value - GEN1 for configuring PcieSlotNvme3linkSpeed token.* `GEN2` - Value - GEN2 for configuring PcieSlotNvme3linkSpeed token.* `GEN3` - Value - GEN3 for configuring PcieSlotNvme3linkSpeed token. 
* `pcie_slot_nvme3option_rom`:(string) BIOS Token for setting NVME-3 OptionROM configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `enabled` - Enables the BIOS setting.* `disabled` - Disables the BIOS setting. 
* `pcie_slot_nvme4link_speed`:(string) BIOS Token for setting NVME-4 Link Speed configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `Auto` - Value - Auto for configuring PcieSlotNvme4linkSpeed token.* `Disabled` - Value - Disabled for configuring PcieSlotNvme4linkSpeed token.* `GEN1` - Value - GEN1 for configuring PcieSlotNvme4linkSpeed token.* `GEN2` - Value - GEN2 for configuring PcieSlotNvme4linkSpeed token.* `GEN3` - Value - GEN3 for configuring PcieSlotNvme4linkSpeed token. 
* `pcie_slot_nvme4option_rom`:(string) BIOS Token for setting NVME-4 OptionROM configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `enabled` - Enables the BIOS setting.* `disabled` - Disables the BIOS setting. 
* `pcie_slot_nvme5link_speed`:(string) BIOS Token for setting NVME-5 Link Speed configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `Auto` - Value - Auto for configuring PcieSlotNvme5linkSpeed token.* `Disabled` - Value - Disabled for configuring PcieSlotNvme5linkSpeed token.* `GEN1` - Value - GEN1 for configuring PcieSlotNvme5linkSpeed token.* `GEN2` - Value - GEN2 for configuring PcieSlotNvme5linkSpeed token.* `GEN3` - Value - GEN3 for configuring PcieSlotNvme5linkSpeed token. 
* `pcie_slot_nvme5option_rom`:(string) BIOS Token for setting NVME-5 OptionROM configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `enabled` - Enables the BIOS setting.* `disabled` - Disables the BIOS setting. 
* `pcie_slot_nvme6link_speed`:(string) BIOS Token for setting NVME-6 Link Speed configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `Auto` - Value - Auto for configuring PcieSlotNvme6linkSpeed token.* `Disabled` - Value - Disabled for configuring PcieSlotNvme6linkSpeed token.* `GEN1` - Value - GEN1 for configuring PcieSlotNvme6linkSpeed token.* `GEN2` - Value - GEN2 for configuring PcieSlotNvme6linkSpeed token.* `GEN3` - Value - GEN3 for configuring PcieSlotNvme6linkSpeed token. 
* `pcie_slot_nvme6option_rom`:(string) BIOS Token for setting NVME-6 OptionROM configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `enabled` - Enables the BIOS setting.* `disabled` - Disables the BIOS setting. 
* `permission_resources`:(Array)(Computed) An array of relationships to moBaseMo resources. 
This complex property has following sub-properties:
  + `moid`:(string)(Computed) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string)(Computed) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `pop_support`:(string) BIOS Token for setting Power ON Password configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `enabled` - Enables the BIOS setting.* `disabled` - Disables the BIOS setting. 
* `post_error_pause`:(string) BIOS Token for setting POST Error Pause configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `enabled` - Enables the BIOS setting.* `disabled` - Disables the BIOS setting. 
* `processor_c1e`:(string) BIOS Token for setting Processor C1E configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `enabled` - Enables the BIOS setting.* `disabled` - Disables the BIOS setting. 
* `processor_c3report`:(string) BIOS Token for setting Processor C3 Report configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `enabled` - Enables the BIOS setting.* `disabled` - Disables the BIOS setting. 
* `processor_c6report`:(string) BIOS Token for setting Processor C6 Report configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `enabled` - Enables the BIOS setting.* `disabled` - Disables the BIOS setting. 
* `processor_cstate`:(string) BIOS Token for setting CPU C State configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `enabled` - Enables the BIOS setting.* `disabled` - Disables the BIOS setting. 
* `profiles`:(Array) An array of relationships to policyAbstractConfigProfile resources. 
This complex property has following sub-properties:
  + `moid`:(string)(Computed) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string)(Computed) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `psata`:(string) BIOS Token for setting P-SATA Mode configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `AHCI` - Value - AHCI for configuring Psata token.* `Disabled` - Value - Disabled for configuring Psata token.* `LSI SW RAID` - Value - LSI SW RAID for configuring Psata token. 
* `pstate_coord_type`:(string) BIOS Token for setting P-STATE Coordination configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `HW ALL` - Value - HW ALL for configuring PstateCoordType token.* `SW ALL` - Value - SW ALL for configuring PstateCoordType token.* `SW ANY` - Value - SW ANY for configuring PstateCoordType token. 
* `putty_key_pad`:(string) BIOS Token for setting Putty KeyPad configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `ESCN` - Value - ESCN for configuring PuttyKeyPad token.* `LINUX` - Value - LINUX for configuring PuttyKeyPad token.* `SCO` - Value - SCO for configuring PuttyKeyPad token.* `VT100` - Value - VT100 for configuring PuttyKeyPad token.* `VT400` - Value - VT400 for configuring PuttyKeyPad token.* `XTERMR6` - Value - XTERMR6 for configuring PuttyKeyPad token. 
* `pwr_perf_tuning`:(string) BIOS Token for setting Power Performance Tuning configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `bios` - Value - BIOS for configuring PwrPerfTuning token.* `os` - Value - os for configuring PwrPerfTuning token. 
* `qpi_link_frequency`:(string) BIOS Token for setting QPI Link Frequency Select configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `6.4-gt/s` - Value - 6.4-gt/s for configuring QpiLinkFrequency token.* `7.2-gt/s` - Value - 7.2-gt/s for configuring QpiLinkFrequency token.* `8.0-gt/s` - Value - 8.0-gt/s for configuring QpiLinkFrequency token.* `9.6-gt/s` - Value - 9.6-gt/s for configuring QpiLinkFrequency token.* `auto` - Value - auto for configuring QpiLinkFrequency token. 
* `qpi_link_speed`:(string) BIOS Token for setting UPI Link Frequency Select configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `10.4GT/s` - Value - 10.4GT/s for configuring QpiLinkSpeed token.* `9.6GT/s` - Value - 9.6GT/s for configuring QpiLinkSpeed token.* `Auto` - Value - Auto for configuring QpiLinkSpeed token. 
* `qpi_snoop_mode`:(string) BIOS Token for setting QPI Snoop Mode configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `auto` - Value - auto for configuring QpiSnoopMode token.* `cluster-on-die` - Value - cluster-on-die for configuring QpiSnoopMode token.* `early-snoop` - Value - early-snoop for configuring QpiSnoopMode token.* `home-directory-snoop` - Value - home-directory-snoop for configuring QpiSnoopMode token.* `home-directory-snoop-with-osb` - Value - home-directory-snoop-with-osb for configuring QpiSnoopMode token.* `home-snoop` - Value - home-snoop for configuring QpiSnoopMode token. 
* `rank_inter_leave`:(string) BIOS Token for setting Rank Interleaving configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `1-way` - Value - 1-way for configuring RankInterLeave token.* `2-way` - Value - 2-way for configuring RankInterLeave token.* `4-way` - Value - 4-way for configuring RankInterLeave token.* `8-way` - Value - 8-way for configuring RankInterLeave token.* `auto` - Value - auto for configuring RankInterLeave token. 
* `redirection_after_post`:(string) BIOS Token for setting Redirection After BIOS POST configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `Always Enable` - Value - Always Enable for configuring RedirectionAfterPost token.* `Bootloader` - Value - Bootloader for configuring RedirectionAfterPost token. 
* `sata_mode_select`:(string) BIOS Token for setting SATA Mode configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `AHCI` - Value - AHCI for configuring SataModeSelect token.* `Disabled` - Value - Disabled for configuring SataModeSelect token.* `LSI SW RAID` - Value - LSI SW RAID for configuring SataModeSelect token. 
* `select_memory_ras_configuration`:(string) BIOS Token for setting Memory RAS Configuration configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `adddc-sparing` - Value - adddc-sparing for configuring SelectMemoryRasConfiguration token.* `lockstep` - Value - lockstep for configuring SelectMemoryRasConfiguration token.* `maximum-performance` - Value - maximum-performance for configuring SelectMemoryRasConfiguration token.* `mirror-mode-1lm` - Value - mirror-mode-1lm for configuring SelectMemoryRasConfiguration token.* `mirroring` - Value - mirroring for configuring SelectMemoryRasConfiguration token.* `partial-mirror-mode-1lm` - Value - partial-mirror-mode-1lm for configuring SelectMemoryRasConfiguration token.* `sparing` - Value - sparing for configuring SelectMemoryRasConfiguration token. 
* `select_ppr_type`:(string) BIOS Token for setting PPR Type configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `disabled` - Value - disabled for configuring SelectPprType token.* `Hard PPR` - Value - Hard PPR for configuring SelectPprType token. 
* `serial_port_aenable`:(string) BIOS Token for setting Serial A Enable configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `enabled` - Enables the BIOS setting.* `disabled` - Disables the BIOS setting. 
* `sev`:(string) BIOS Token for setting Secured Encrypted Virtualization configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `253 ASIDs` - Value - 253 ASIDs for configuring Sev token.* `509 ASIDs` - Value - 509 ASIDs for configuring Sev token.* `Auto` - Value - Auto for configuring Sev token. 
* `shared_scope`:(string)(Computed) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `single_pctl_enable`:(string) BIOS Token for setting Single PCTL configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `No` - Value - No for configuring SinglePctlEnable token.* `Yes` - Value - Yes for configuring SinglePctlEnable token. 
* `slot10link_speed`:(string) BIOS Token for setting PCIe Slot:10 Link Speed configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `Auto` - Value - Auto for configuring Slot10linkSpeed token.* `Disabled` - Value - Disabled for configuring Slot10linkSpeed token.* `GEN1` - Value - GEN1 for configuring Slot10linkSpeed token.* `GEN2` - Value - GEN2 for configuring Slot10linkSpeed token.* `GEN3` - Value - GEN3 for configuring Slot10linkSpeed token. 
* `slot10state`:(string) BIOS Token for setting Slot 10 State configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `disabled` - Value - disabled for configuring Slot10state token.* `enabled` - Value - enabled for configuring Slot10state token.* `Legacy Only` - Value - Legacy Only for configuring Slot10state token.* `UEFI Only` - Value - UEFI Only for configuring Slot10state token. 
* `slot11link_speed`:(string) BIOS Token for setting PCIe Slot:11 Link Speed configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `Auto` - Value - Auto for configuring Slot11linkSpeed token.* `Disabled` - Value - Disabled for configuring Slot11linkSpeed token.* `GEN1` - Value - GEN1 for configuring Slot11linkSpeed token.* `GEN2` - Value - GEN2 for configuring Slot11linkSpeed token.* `GEN3` - Value - GEN3 for configuring Slot11linkSpeed token. 
* `slot11state`:(string) BIOS Token for setting Slot 11 State configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `enabled` - Enables the BIOS setting.* `disabled` - Disables the BIOS setting. 
* `slot12link_speed`:(string) BIOS Token for setting PCIe Slot:12 Link Speed configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `Auto` - Value - Auto for configuring Slot12linkSpeed token.* `Disabled` - Value - Disabled for configuring Slot12linkSpeed token.* `GEN1` - Value - GEN1 for configuring Slot12linkSpeed token.* `GEN2` - Value - GEN2 for configuring Slot12linkSpeed token.* `GEN3` - Value - GEN3 for configuring Slot12linkSpeed token. 
* `slot12state`:(string) BIOS Token for setting Slot 12 State configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `enabled` - Enables the BIOS setting.* `disabled` - Disables the BIOS setting. 
* `slot13state`:(string) BIOS Token for setting Slot 13 State configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `enabled` - Enables the BIOS setting.* `disabled` - Disables the BIOS setting. 
* `slot14state`:(string) BIOS Token for setting Slot 14 State configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `enabled` - Enables the BIOS setting.* `disabled` - Disables the BIOS setting. 
* `slot1link_speed`:(string) BIOS Token for setting PCIe Slot:1 Link Speed configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `Auto` - Value - Auto for configuring Slot1linkSpeed token.* `Disabled` - Value - Disabled for configuring Slot1linkSpeed token.* `GEN1` - Value - GEN1 for configuring Slot1linkSpeed token.* `GEN2` - Value - GEN2 for configuring Slot1linkSpeed token.* `GEN3` - Value - GEN3 for configuring Slot1linkSpeed token. 
* `slot1state`:(string) BIOS Token for setting Slot 1 State configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `disabled` - Value - disabled for configuring Slot1state token.* `enabled` - Value - enabled for configuring Slot1state token.* `Legacy Only` - Value - Legacy Only for configuring Slot1state token.* `UEFI Only` - Value - UEFI Only for configuring Slot1state token. 
* `slot2link_speed`:(string) BIOS Token for setting PCIe Slot:2 Link Speed configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `Auto` - Value - Auto for configuring Slot2linkSpeed token.* `Disabled` - Value - Disabled for configuring Slot2linkSpeed token.* `GEN1` - Value - GEN1 for configuring Slot2linkSpeed token.* `GEN2` - Value - GEN2 for configuring Slot2linkSpeed token.* `GEN3` - Value - GEN3 for configuring Slot2linkSpeed token. 
* `slot2state`:(string) BIOS Token for setting Slot 2 State configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `disabled` - Value - disabled for configuring Slot2state token.* `enabled` - Value - enabled for configuring Slot2state token.* `Legacy Only` - Value - Legacy Only for configuring Slot2state token.* `UEFI Only` - Value - UEFI Only for configuring Slot2state token. 
* `slot3link_speed`:(string) BIOS Token for setting PCIe Slot:3 Link Speed configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `Auto` - Value - Auto for configuring Slot3linkSpeed token.* `Disabled` - Value - Disabled for configuring Slot3linkSpeed token.* `GEN1` - Value - GEN1 for configuring Slot3linkSpeed token.* `GEN2` - Value - GEN2 for configuring Slot3linkSpeed token.* `GEN3` - Value - GEN3 for configuring Slot3linkSpeed token. 
* `slot3state`:(string) BIOS Token for setting Slot 3 State configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `disabled` - Value - disabled for configuring Slot3state token.* `enabled` - Value - enabled for configuring Slot3state token.* `Legacy Only` - Value - Legacy Only for configuring Slot3state token.* `UEFI Only` - Value - UEFI Only for configuring Slot3state token. 
* `slot4link_speed`:(string) BIOS Token for setting PCIe Slot:4 Link Speed configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `Auto` - Value - Auto for configuring Slot4linkSpeed token.* `Disabled` - Value - Disabled for configuring Slot4linkSpeed token.* `GEN1` - Value - GEN1 for configuring Slot4linkSpeed token.* `GEN2` - Value - GEN2 for configuring Slot4linkSpeed token.* `GEN3` - Value - GEN3 for configuring Slot4linkSpeed token. 
* `slot4state`:(string) BIOS Token for setting Slot 4 State configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `disabled` - Value - disabled for configuring Slot4state token.* `enabled` - Value - enabled for configuring Slot4state token.* `Legacy Only` - Value - Legacy Only for configuring Slot4state token.* `UEFI Only` - Value - UEFI Only for configuring Slot4state token. 
* `slot5link_speed`:(string) BIOS Token for setting PCIe Slot:5 Link Speed configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `Auto` - Value - Auto for configuring Slot5linkSpeed token.* `Disabled` - Value - Disabled for configuring Slot5linkSpeed token.* `GEN1` - Value - GEN1 for configuring Slot5linkSpeed token.* `GEN2` - Value - GEN2 for configuring Slot5linkSpeed token.* `GEN3` - Value - GEN3 for configuring Slot5linkSpeed token. 
* `slot5state`:(string) BIOS Token for setting Slot 5 State configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `disabled` - Value - disabled for configuring Slot5state token.* `enabled` - Value - enabled for configuring Slot5state token.* `Legacy Only` - Value - Legacy Only for configuring Slot5state token.* `UEFI Only` - Value - UEFI Only for configuring Slot5state token. 
* `slot6link_speed`:(string) BIOS Token for setting PCIe Slot:6 Link Speed configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `Auto` - Value - Auto for configuring Slot6linkSpeed token.* `Disabled` - Value - Disabled for configuring Slot6linkSpeed token.* `GEN1` - Value - GEN1 for configuring Slot6linkSpeed token.* `GEN2` - Value - GEN2 for configuring Slot6linkSpeed token.* `GEN3` - Value - GEN3 for configuring Slot6linkSpeed token. 
* `slot6state`:(string) BIOS Token for setting Slot 6 State configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `disabled` - Value - disabled for configuring Slot6state token.* `enabled` - Value - enabled for configuring Slot6state token.* `Legacy Only` - Value - Legacy Only for configuring Slot6state token.* `UEFI Only` - Value - UEFI Only for configuring Slot6state token. 
* `slot7link_speed`:(string) BIOS Token for setting PCIe Slot:7 Link Speed configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `Auto` - Value - Auto for configuring Slot7linkSpeed token.* `Disabled` - Value - Disabled for configuring Slot7linkSpeed token.* `GEN1` - Value - GEN1 for configuring Slot7linkSpeed token.* `GEN2` - Value - GEN2 for configuring Slot7linkSpeed token.* `GEN3` - Value - GEN3 for configuring Slot7linkSpeed token. 
* `slot7state`:(string) BIOS Token for setting Slot 7 State configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `disabled` - Value - disabled for configuring Slot7state token.* `enabled` - Value - enabled for configuring Slot7state token.* `Legacy Only` - Value - Legacy Only for configuring Slot7state token.* `UEFI Only` - Value - UEFI Only for configuring Slot7state token. 
* `slot8link_speed`:(string) BIOS Token for setting PCIe Slot:8 Link Speed configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `Auto` - Value - Auto for configuring Slot8linkSpeed token.* `Disabled` - Value - Disabled for configuring Slot8linkSpeed token.* `GEN1` - Value - GEN1 for configuring Slot8linkSpeed token.* `GEN2` - Value - GEN2 for configuring Slot8linkSpeed token.* `GEN3` - Value - GEN3 for configuring Slot8linkSpeed token. 
* `slot8state`:(string) BIOS Token for setting Slot 8 State configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `disabled` - Value - disabled for configuring Slot8state token.* `enabled` - Value - enabled for configuring Slot8state token.* `Legacy Only` - Value - Legacy Only for configuring Slot8state token.* `UEFI Only` - Value - UEFI Only for configuring Slot8state token. 
* `slot9link_speed`:(string) BIOS Token for setting PCIe Slot:9 Link Speed configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `Auto` - Value - Auto for configuring Slot9linkSpeed token.* `Disabled` - Value - Disabled for configuring Slot9linkSpeed token.* `GEN1` - Value - GEN1 for configuring Slot9linkSpeed token.* `GEN2` - Value - GEN2 for configuring Slot9linkSpeed token.* `GEN3` - Value - GEN3 for configuring Slot9linkSpeed token. 
* `slot9state`:(string) BIOS Token for setting Slot 9 State configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `disabled` - Value - disabled for configuring Slot9state token.* `enabled` - Value - enabled for configuring Slot9state token.* `Legacy Only` - Value - Legacy Only for configuring Slot9state token.* `UEFI Only` - Value - UEFI Only for configuring Slot9state token. 
* `slot_flom_link_speed`:(string) BIOS Token for setting PCIe Slot:FLOM Link Speed configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `Auto` - Value - Auto for configuring SlotFlomLinkSpeed token.* `Disabled` - Value - Disabled for configuring SlotFlomLinkSpeed token.* `GEN1` - Value - GEN1 for configuring SlotFlomLinkSpeed token.* `GEN2` - Value - GEN2 for configuring SlotFlomLinkSpeed token.* `GEN3` - Value - GEN3 for configuring SlotFlomLinkSpeed token. 
* `slot_front_nvme1link_speed`:(string) BIOS Token for setting PCIe Slot:Front NVME1 Link Speed configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `Auto` - Value - Auto for configuring SlotFrontNvme1linkSpeed token.* `Disabled` - Value - Disabled for configuring SlotFrontNvme1linkSpeed token.* `GEN1` - Value - GEN1 for configuring SlotFrontNvme1linkSpeed token.* `GEN2` - Value - GEN2 for configuring SlotFrontNvme1linkSpeed token.* `GEN3` - Value - GEN3 for configuring SlotFrontNvme1linkSpeed token. 
* `slot_front_nvme2link_speed`:(string) BIOS Token for setting PCIe Slot:Front NVME2 Link Speed configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `Auto` - Value - Auto for configuring SlotFrontNvme2linkSpeed token.* `Disabled` - Value - Disabled for configuring SlotFrontNvme2linkSpeed token.* `GEN1` - Value - GEN1 for configuring SlotFrontNvme2linkSpeed token.* `GEN2` - Value - GEN2 for configuring SlotFrontNvme2linkSpeed token.* `GEN3` - Value - GEN3 for configuring SlotFrontNvme2linkSpeed token. 
* `slot_front_slot5link_speed`:(string) BIOS Token for setting PCIe Slot:Front1 Link Speed configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `Auto` - Value - Auto for configuring SlotFrontSlot5linkSpeed token.* `Disabled` - Value - Disabled for configuring SlotFrontSlot5linkSpeed token.* `GEN1` - Value - GEN1 for configuring SlotFrontSlot5linkSpeed token.* `GEN2` - Value - GEN2 for configuring SlotFrontSlot5linkSpeed token.* `GEN3` - Value - GEN3 for configuring SlotFrontSlot5linkSpeed token. 
* `slot_front_slot6link_speed`:(string) BIOS Token for setting PCIe Slot:Front2 Link Speed configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `Auto` - Value - Auto for configuring SlotFrontSlot6linkSpeed token.* `Disabled` - Value - Disabled for configuring SlotFrontSlot6linkSpeed token.* `GEN1` - Value - GEN1 for configuring SlotFrontSlot6linkSpeed token.* `GEN2` - Value - GEN2 for configuring SlotFrontSlot6linkSpeed token.* `GEN3` - Value - GEN3 for configuring SlotFrontSlot6linkSpeed token. 
* `slot_gpu1state`:(string) BIOS Token for setting GPU1 OptionROM configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `enabled` - Enables the BIOS setting.* `disabled` - Disables the BIOS setting. 
* `slot_gpu2state`:(string) BIOS Token for setting GPU2 OptionROM configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `enabled` - Enables the BIOS setting.* `disabled` - Disables the BIOS setting. 
* `slot_gpu3state`:(string) BIOS Token for setting GPU3 OptionROM configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `enabled` - Enables the BIOS setting.* `disabled` - Disables the BIOS setting. 
* `slot_gpu4state`:(string) BIOS Token for setting GPU4 OptionROM configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `enabled` - Enables the BIOS setting.* `disabled` - Disables the BIOS setting. 
* `slot_gpu5state`:(string) BIOS Token for setting GPU5 OptionROM configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `enabled` - Enables the BIOS setting.* `disabled` - Disables the BIOS setting. 
* `slot_gpu6state`:(string) BIOS Token for setting GPU6 OptionROM configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `enabled` - Enables the BIOS setting.* `disabled` - Disables the BIOS setting. 
* `slot_gpu7state`:(string) BIOS Token for setting GPU7 OptionROM configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `enabled` - Enables the BIOS setting.* `disabled` - Disables the BIOS setting. 
* `slot_gpu8state`:(string) BIOS Token for setting GPU8 OptionROM configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `enabled` - Enables the BIOS setting.* `disabled` - Disables the BIOS setting. 
* `slot_hba_link_speed`:(string) BIOS Token for setting PCIe Slot:HBA Link Speed configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `Auto` - Value - Auto for configuring SlotHbaLinkSpeed token.* `Disabled` - Value - Disabled for configuring SlotHbaLinkSpeed token.* `GEN1` - Value - GEN1 for configuring SlotHbaLinkSpeed token.* `GEN2` - Value - GEN2 for configuring SlotHbaLinkSpeed token.* `GEN3` - Value - GEN3 for configuring SlotHbaLinkSpeed token. 
* `slot_hba_state`:(string) BIOS Token for setting PCIe Slot:HBA OptionROM configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `disabled` - Value - disabled for configuring SlotHbaState token.* `enabled` - Value - enabled for configuring SlotHbaState token.* `Legacy Only` - Value - Legacy Only for configuring SlotHbaState token.* `UEFI Only` - Value - UEFI Only for configuring SlotHbaState token. 
* `slot_lom1link`:(string) BIOS Token for setting PCIe LOM:1 Link configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `enabled` - Enables the BIOS setting.* `disabled` - Disables the BIOS setting. 
* `slot_lom2link`:(string) BIOS Token for setting PCIe LOM:2 Link configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `enabled` - Enables the BIOS setting.* `disabled` - Disables the BIOS setting. 
* `slot_mezz_state`:(string) BIOS Token for setting Slot Mezz State configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `disabled` - Value - disabled for configuring SlotMezzState token.* `enabled` - Value - enabled for configuring SlotMezzState token.* `Legacy Only` - Value - Legacy Only for configuring SlotMezzState token.* `UEFI Only` - Value - UEFI Only for configuring SlotMezzState token. 
* `slot_mlom_link_speed`:(string) BIOS Token for setting PCIe Slot:MLOM Link Speed configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `Auto` - Value - Auto for configuring SlotMlomLinkSpeed token.* `Disabled` - Value - Disabled for configuring SlotMlomLinkSpeed token.* `GEN1` - Value - GEN1 for configuring SlotMlomLinkSpeed token.* `GEN2` - Value - GEN2 for configuring SlotMlomLinkSpeed token.* `GEN3` - Value - GEN3 for configuring SlotMlomLinkSpeed token. 
* `slot_mlom_state`:(string) BIOS Token for setting PCIe Slot MLOM OptionROM configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `disabled` - Value - disabled for configuring SlotMlomState token.* `enabled` - Value - enabled for configuring SlotMlomState token.* `Legacy Only` - Value - Legacy Only for configuring SlotMlomState token.* `UEFI Only` - Value - UEFI Only for configuring SlotMlomState token. 
* `slot_mraid_link_speed`:(string) BIOS Token for setting MRAID Link Speed configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `Auto` - Value - Auto for configuring SlotMraidLinkSpeed token.* `Disabled` - Value - Disabled for configuring SlotMraidLinkSpeed token.* `GEN1` - Value - GEN1 for configuring SlotMraidLinkSpeed token.* `GEN2` - Value - GEN2 for configuring SlotMraidLinkSpeed token.* `GEN3` - Value - GEN3 for configuring SlotMraidLinkSpeed token. 
* `slot_mraid_state`:(string) BIOS Token for setting PCIe Slot MRAID OptionROM configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `enabled` - Enables the BIOS setting.* `disabled` - Disables the BIOS setting. 
* `slot_n10state`:(string) BIOS Token for setting PCIe Slot N10 OptionROM configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `enabled` - Enables the BIOS setting.* `disabled` - Disables the BIOS setting. 
* `slot_n11state`:(string) BIOS Token for setting PCIe Slot N11 OptionROM configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `enabled` - Enables the BIOS setting.* `disabled` - Disables the BIOS setting. 
* `slot_n12state`:(string) BIOS Token for setting PCIe Slot N12 OptionROM configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `enabled` - Enables the BIOS setting.* `disabled` - Disables the BIOS setting. 
* `slot_n13state`:(string) BIOS Token for setting PCIe Slot N13 OptionROM configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `enabled` - Enables the BIOS setting.* `disabled` - Disables the BIOS setting. 
* `slot_n14state`:(string) BIOS Token for setting PCIe Slot N14 OptionROM configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `enabled` - Enables the BIOS setting.* `disabled` - Disables the BIOS setting. 
* `slot_n15state`:(string) BIOS Token for setting PCIe Slot N15 OptionROM configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `enabled` - Enables the BIOS setting.* `disabled` - Disables the BIOS setting. 
* `slot_n16state`:(string) BIOS Token for setting PCIe Slot N16 OptionROM configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `enabled` - Enables the BIOS setting.* `disabled` - Disables the BIOS setting. 
* `slot_n17state`:(string) BIOS Token for setting PCIe Slot N17 OptionROM configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `enabled` - Enables the BIOS setting.* `disabled` - Disables the BIOS setting. 
* `slot_n18state`:(string) BIOS Token for setting PCIe Slot N18 OptionROM configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `enabled` - Enables the BIOS setting.* `disabled` - Disables the BIOS setting. 
* `slot_n19state`:(string) BIOS Token for setting PCIe Slot N19 OptionROM configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `enabled` - Enables the BIOS setting.* `disabled` - Disables the BIOS setting. 
* `slot_n1state`:(string) BIOS Token for setting PCIe Slot N1 OptionROM configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `disabled` - Value - disabled for configuring SlotN1state token.* `enabled` - Value - enabled for configuring SlotN1state token.* `Legacy Only` - Value - Legacy Only for configuring SlotN1state token.* `UEFI Only` - Value - UEFI Only for configuring SlotN1state token. 
* `slot_n20state`:(string) BIOS Token for setting PCIe Slot N20 OptionROM configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `enabled` - Enables the BIOS setting.* `disabled` - Disables the BIOS setting. 
* `slot_n21state`:(string) BIOS Token for setting PCIe Slot N21 OptionROM configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `enabled` - Enables the BIOS setting.* `disabled` - Disables the BIOS setting. 
* `slot_n22state`:(string) BIOS Token for setting PCIe Slot N22 OptionROM configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `enabled` - Enables the BIOS setting.* `disabled` - Disables the BIOS setting. 
* `slot_n23state`:(string) BIOS Token for setting PCIe Slot N23 OptionROM configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `enabled` - Enables the BIOS setting.* `disabled` - Disables the BIOS setting. 
* `slot_n24state`:(string) BIOS Token for setting PCIe Slot N24 OptionROM configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `enabled` - Enables the BIOS setting.* `disabled` - Disables the BIOS setting. 
* `slot_n2state`:(string) BIOS Token for setting PCIe Slot N2 OptionROM configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `disabled` - Value - disabled for configuring SlotN2state token.* `enabled` - Value - enabled for configuring SlotN2state token.* `Legacy Only` - Value - Legacy Only for configuring SlotN2state token.* `UEFI Only` - Value - UEFI Only for configuring SlotN2state token. 
* `slot_n3state`:(string) BIOS Token for setting PCIe Slot N3 OptionROM configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `enabled` - Enables the BIOS setting.* `disabled` - Disables the BIOS setting. 
* `slot_n4state`:(string) BIOS Token for setting PCIe Slot N4 OptionROM configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `enabled` - Enables the BIOS setting.* `disabled` - Disables the BIOS setting. 
* `slot_n5state`:(string) BIOS Token for setting PCIe Slot N5 OptionROM configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `enabled` - Enables the BIOS setting.* `disabled` - Disables the BIOS setting. 
* `slot_n6state`:(string) BIOS Token for setting PCIe Slot N6 OptionROM configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `enabled` - Enables the BIOS setting.* `disabled` - Disables the BIOS setting. 
* `slot_n7state`:(string) BIOS Token for setting PCIe Slot N7 OptionROM configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `enabled` - Enables the BIOS setting.* `disabled` - Disables the BIOS setting. 
* `slot_n8state`:(string) BIOS Token for setting PCIe Slot N8 OptionROM configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `enabled` - Enables the BIOS setting.* `disabled` - Disables the BIOS setting. 
* `slot_n9state`:(string) BIOS Token for setting PCIe Slot N9 OptionROM configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `enabled` - Enables the BIOS setting.* `disabled` - Disables the BIOS setting. 
* `slot_raid_link_speed`:(string) BIOS Token for setting RAID Link Speed configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `Auto` - Value - Auto for configuring SlotRaidLinkSpeed token.* `Disabled` - Value - Disabled for configuring SlotRaidLinkSpeed token.* `GEN1` - Value - GEN1 for configuring SlotRaidLinkSpeed token.* `GEN2` - Value - GEN2 for configuring SlotRaidLinkSpeed token.* `GEN3` - Value - GEN3 for configuring SlotRaidLinkSpeed token. 
* `slot_raid_state`:(string) BIOS Token for setting PCIe Slot RAID OptionROM configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `enabled` - Enables the BIOS setting.* `disabled` - Disables the BIOS setting. 
* `slot_rear_nvme1link_speed`:(string) BIOS Token for setting PCIe Slot:Rear NVME1 Link Speed configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `Auto` - Value - Auto for configuring SlotRearNvme1linkSpeed token.* `Disabled` - Value - Disabled for configuring SlotRearNvme1linkSpeed token.* `GEN1` - Value - GEN1 for configuring SlotRearNvme1linkSpeed token.* `GEN2` - Value - GEN2 for configuring SlotRearNvme1linkSpeed token.* `GEN3` - Value - GEN3 for configuring SlotRearNvme1linkSpeed token. 
* `slot_rear_nvme1state`:(string) BIOS Token for setting PCIe Slot:Rear NVME 1 OptionROM configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `enabled` - Enables the BIOS setting.* `disabled` - Disables the BIOS setting. 
* `slot_rear_nvme2link_speed`:(string) BIOS Token for setting PCIe Slot:Rear NVME2 Link Speed configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `Auto` - Value - Auto for configuring SlotRearNvme2linkSpeed token.* `Disabled` - Value - Disabled for configuring SlotRearNvme2linkSpeed token.* `GEN1` - Value - GEN1 for configuring SlotRearNvme2linkSpeed token.* `GEN2` - Value - GEN2 for configuring SlotRearNvme2linkSpeed token.* `GEN3` - Value - GEN3 for configuring SlotRearNvme2linkSpeed token. 
* `slot_rear_nvme2state`:(string) BIOS Token for setting PCIe Slot:Rear NVME 2 OptionROM configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `enabled` - Enables the BIOS setting.* `disabled` - Disables the BIOS setting. 
* `slot_rear_nvme3state`:(string) BIOS Token for setting PCIe Slot:Rear NVME 3 OptionROM configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `enabled` - Enables the BIOS setting.* `disabled` - Disables the BIOS setting. 
* `slot_rear_nvme4state`:(string) BIOS Token for setting PCIe Slot:Rear NVME 4 OptionROM configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `enabled` - Enables the BIOS setting.* `disabled` - Disables the BIOS setting. 
* `slot_rear_nvme5state`:(string) BIOS Token for setting PCIe Slot:Rear NVME 5 OptionROM configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `enabled` - Enables the BIOS setting.* `disabled` - Disables the BIOS setting. 
* `slot_rear_nvme6state`:(string) BIOS Token for setting PCIe Slot:Rear NVME 6 OptionROM configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `enabled` - Enables the BIOS setting.* `disabled` - Disables the BIOS setting. 
* `slot_rear_nvme7state`:(string) BIOS Token for setting PCIe Slot:Rear NVME 7 OptionROM configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `enabled` - Enables the BIOS setting.* `disabled` - Disables the BIOS setting. 
* `slot_rear_nvme8state`:(string) BIOS Token for setting PCIe Slot:Rear NVME 8 OptionROM configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `enabled` - Enables the BIOS setting.* `disabled` - Disables the BIOS setting. 
* `slot_riser1link_speed`:(string) BIOS Token for setting PCIe Slot:Riser1 Link Speed configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `Auto` - Value - Auto for configuring SlotRiser1linkSpeed token.* `Disabled` - Value - Disabled for configuring SlotRiser1linkSpeed token.* `GEN1` - Value - GEN1 for configuring SlotRiser1linkSpeed token.* `GEN2` - Value - GEN2 for configuring SlotRiser1linkSpeed token.* `GEN3` - Value - GEN3 for configuring SlotRiser1linkSpeed token. 
* `slot_riser1slot1link_speed`:(string) BIOS Token for setting PCIe Slot:Riser1 Slot1 Link Speed configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `Auto` - Value - Auto for configuring SlotRiser1slot1linkSpeed token.* `Disabled` - Value - Disabled for configuring SlotRiser1slot1linkSpeed token.* `GEN1` - Value - GEN1 for configuring SlotRiser1slot1linkSpeed token.* `GEN2` - Value - GEN2 for configuring SlotRiser1slot1linkSpeed token.* `GEN3` - Value - GEN3 for configuring SlotRiser1slot1linkSpeed token. 
* `slot_riser1slot2link_speed`:(string) BIOS Token for setting PCIe Slot:Riser1 Slot2 Link Speed configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `Auto` - Value - Auto for configuring SlotRiser1slot2linkSpeed token.* `Disabled` - Value - Disabled for configuring SlotRiser1slot2linkSpeed token.* `GEN1` - Value - GEN1 for configuring SlotRiser1slot2linkSpeed token.* `GEN2` - Value - GEN2 for configuring SlotRiser1slot2linkSpeed token.* `GEN3` - Value - GEN3 for configuring SlotRiser1slot2linkSpeed token. 
* `slot_riser1slot3link_speed`:(string) BIOS Token for setting PCIe Slot:Riser1 Slot3 Link Speed configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `Auto` - Value - Auto for configuring SlotRiser1slot3linkSpeed token.* `Disabled` - Value - Disabled for configuring SlotRiser1slot3linkSpeed token.* `GEN1` - Value - GEN1 for configuring SlotRiser1slot3linkSpeed token.* `GEN2` - Value - GEN2 for configuring SlotRiser1slot3linkSpeed token.* `GEN3` - Value - GEN3 for configuring SlotRiser1slot3linkSpeed token. 
* `slot_riser2link_speed`:(string) BIOS Token for setting PCIe Slot:Riser2 Link Speed configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `Auto` - Value - Auto for configuring SlotRiser2linkSpeed token.* `Disabled` - Value - Disabled for configuring SlotRiser2linkSpeed token.* `GEN1` - Value - GEN1 for configuring SlotRiser2linkSpeed token.* `GEN2` - Value - GEN2 for configuring SlotRiser2linkSpeed token.* `GEN3` - Value - GEN3 for configuring SlotRiser2linkSpeed token. 
* `slot_riser2slot4link_speed`:(string) BIOS Token for setting PCIe Slot:Riser2 Slot4 Link Speed configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `Auto` - Value - Auto for configuring SlotRiser2slot4linkSpeed token.* `Disabled` - Value - Disabled for configuring SlotRiser2slot4linkSpeed token.* `GEN1` - Value - GEN1 for configuring SlotRiser2slot4linkSpeed token.* `GEN2` - Value - GEN2 for configuring SlotRiser2slot4linkSpeed token.* `GEN3` - Value - GEN3 for configuring SlotRiser2slot4linkSpeed token. 
* `slot_riser2slot5link_speed`:(string) BIOS Token for setting PCIe Slot:Riser2 Slot5 Link Speed configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `Auto` - Value - Auto for configuring SlotRiser2slot5linkSpeed token.* `Disabled` - Value - Disabled for configuring SlotRiser2slot5linkSpeed token.* `GEN1` - Value - GEN1 for configuring SlotRiser2slot5linkSpeed token.* `GEN2` - Value - GEN2 for configuring SlotRiser2slot5linkSpeed token.* `GEN3` - Value - GEN3 for configuring SlotRiser2slot5linkSpeed token. 
* `slot_riser2slot6link_speed`:(string) BIOS Token for setting PCIe Slot:Riser2 Slot6 Link Speed configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `Auto` - Value - Auto for configuring SlotRiser2slot6linkSpeed token.* `Disabled` - Value - Disabled for configuring SlotRiser2slot6linkSpeed token.* `GEN1` - Value - GEN1 for configuring SlotRiser2slot6linkSpeed token.* `GEN2` - Value - GEN2 for configuring SlotRiser2slot6linkSpeed token.* `GEN3` - Value - GEN3 for configuring SlotRiser2slot6linkSpeed token. 
* `slot_sas_state`:(string) BIOS Token for setting PCIe Slot:SAS OptionROM configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `disabled` - Value - disabled for configuring SlotSasState token.* `enabled` - Value - enabled for configuring SlotSasState token.* `Legacy Only` - Value - Legacy Only for configuring SlotSasState token.* `UEFI Only` - Value - UEFI Only for configuring SlotSasState token. 
* `slot_ssd_slot1link_speed`:(string) BIOS Token for setting PCIe Slot:FrontSSD1 Link Speed configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `Auto` - Value - Auto for configuring SlotSsdSlot1linkSpeed token.* `Disabled` - Value - Disabled for configuring SlotSsdSlot1linkSpeed token.* `GEN1` - Value - GEN1 for configuring SlotSsdSlot1linkSpeed token.* `GEN2` - Value - GEN2 for configuring SlotSsdSlot1linkSpeed token.* `GEN3` - Value - GEN3 for configuring SlotSsdSlot1linkSpeed token. 
* `slot_ssd_slot2link_speed`:(string) BIOS Token for setting PCIe Slot:FrontSSD2 Link Speed configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `Auto` - Value - Auto for configuring SlotSsdSlot2linkSpeed token.* `Disabled` - Value - Disabled for configuring SlotSsdSlot2linkSpeed token.* `GEN1` - Value - GEN1 for configuring SlotSsdSlot2linkSpeed token.* `GEN2` - Value - GEN2 for configuring SlotSsdSlot2linkSpeed token.* `GEN3` - Value - GEN3 for configuring SlotSsdSlot2linkSpeed token. 
* `smee`:(string) BIOS Token for setting SMEE configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `enabled` - Enables the BIOS setting.* `disabled` - Disables the BIOS setting. 
* `smt_mode`:(string) BIOS Token for setting SMT Mode configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `Auto` - Value - Auto for configuring SmtMode token.* `Off` - Value - Off for configuring SmtMode token. 
* `snc`:(string) BIOS Token for setting Sub Numa Clustering configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `Auto` - Value - Auto for configuring Snc token.* `disabled` - Value - disabled for configuring Snc token.* `enabled` - Value - enabled for configuring Snc token. 
* `snoopy_mode_for2lm`:(string) BIOS Token for setting Snoopy Mode for 2LM configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `enabled` - Enables the BIOS setting.* `disabled` - Disables the BIOS setting. 
* `snoopy_mode_for_ad`:(string) BIOS Token for setting Snoopy Mode for AD configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `enabled` - Enables the BIOS setting.* `disabled` - Disables the BIOS setting. 
* `sparing_mode`:(string) BIOS Token for setting Sparing Mode configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `dimm-sparing` - Value - dimm-sparing for configuring SparingMode token.* `rank-sparing` - Value - rank-sparing for configuring SparingMode token. 
* `sr_iov`:(string) BIOS Token for setting SR-IOV Support configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `enabled` - Enables the BIOS setting.* `disabled` - Disables the BIOS setting. 
* `streamer_prefetch`:(string) BIOS Token for setting DCU Streamer Prefetch configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `enabled` - Enables the BIOS setting.* `disabled` - Disables the BIOS setting. 
* `svm_mode`:(string) BIOS Token for setting SVM Mode configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `enabled` - Enables the BIOS setting.* `disabled` - Disables the BIOS setting. 
* `tags`:(Array)
This complex property has following sub-properties:
  + `key`:(string) The string representation of a tag key. 
  + `value`:(string) The string representation of a tag value. 
* `terminal_type`:(string) BIOS Token for setting Terminal Type configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `pc-ansi` - Value - pc-ansi for configuring TerminalType token.* `vt-utf8` - Value - vt-utf8 for configuring TerminalType token.* `vt100` - Value - vt100 for configuring TerminalType token.* `vt100-plus` - Value - vt100-plus for configuring TerminalType token. 
* `tpm_control`:(string) BIOS Token for setting Trusted Platform Module State configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `enabled` - Enables the BIOS setting.* `disabled` - Disables the BIOS setting. 
* `tpm_support`:(string) BIOS Token for setting TPM Support configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `enabled` - Enables the BIOS setting.* `disabled` - Disables the BIOS setting. 
* `tsme`:(string) BIOS Token for setting Transparent Secure Memory Encryption configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `Auto` - Value - Auto for configuring Tsme token.* `disabled` - Value - disabled for configuring Tsme token.* `enabled` - Value - enabled for configuring Tsme token. 
* `txt_support`:(string) BIOS Token for setting Intel Trusted Execution Technology Support configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `enabled` - Enables the BIOS setting.* `disabled` - Disables the BIOS setting. 
* `ucsm_boot_order_rule`:(string) BIOS Token for setting Boot Order Rules configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `Loose` - Value - Loose for configuring UcsmBootOrderRule token.* `Strict` - Value - Strict for configuring UcsmBootOrderRule token. 
* `ufs_disable`:(string) BIOS Token for setting Uncore Frequency Scaling configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `enabled` - Enables the BIOS setting.* `disabled` - Disables the BIOS setting. 
* `usb_emul6064`:(string) BIOS Token for setting Port 60/64 Emulation configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `enabled` - Enables the BIOS setting.* `disabled` - Disables the BIOS setting. 
* `usb_port_front`:(string) BIOS Token for setting USB Port Front configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `enabled` - Enables the BIOS setting.* `disabled` - Disables the BIOS setting. 
* `usb_port_internal`:(string) BIOS Token for setting USB Port Internal configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `enabled` - Enables the BIOS setting.* `disabled` - Disables the BIOS setting. 
* `usb_port_kvm`:(string) BIOS Token for setting USB Port KVM configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `enabled` - Enables the BIOS setting.* `disabled` - Disables the BIOS setting. 
* `usb_port_rear`:(string) BIOS Token for setting USB Port Rear configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `enabled` - Enables the BIOS setting.* `disabled` - Disables the BIOS setting. 
* `usb_port_sd_card`:(string) BIOS Token for setting USB Port SD Card configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `enabled` - Enables the BIOS setting.* `disabled` - Disables the BIOS setting. 
* `usb_port_vmedia`:(string) BIOS Token for setting USB Port VMedia configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `enabled` - Enables the BIOS setting.* `disabled` - Disables the BIOS setting. 
* `usb_xhci_support`:(string) BIOS Token for setting XHCI Legacy Support configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `enabled` - Enables the BIOS setting.* `disabled` - Disables the BIOS setting. 
* `version_context`:(HashMap) -(Computed) The versioning info for this managed object. 
This complex property has following sub-properties:
  + `interested_mos`:(Array)
This complex property has following sub-properties:
    + `moid`:(string)(Computed) The Moid of the referenced REST resource. 
    + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
    + `selector`:(string)(Computed) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `ref_mo`:(HashMap) -(Computed) A reference to the original Managed Object. 
This complex property has following sub-properties:
    + `moid`:(string)(Computed) The Moid of the referenced REST resource. 
    + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
    + `selector`:(string)(Computed) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
  + `timestamp`:(string)(Computed) The time this versioned Managed Object was created. 
  + `nr_version`:(string)(Computed) The version of the Managed Object, e.g. an incrementing number or a hash id. 
  + `version_type`:(string)(Computed) Specifies type of version. Currently the only supported value is \ Configured\ that is used to keep track of snapshots of policies and profiles that are intendedto be configured to target endpoints.* `Modified` - Version created every time an object is modified.* `Configured` - Version created every time an object is configured to the service profile.* `Deployed` - Version created for objects related to a service profile when it is deployed. 
* `vga_priority`:(string) BIOS Token for setting VGA Priority configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `Offboard` - Value - Offboard for configuring VgaPriority token.* `Onboard` - Value - Onboard for configuring VgaPriority token.* `Onboard VGA Disabled` - Value - Onboard VGA Disabled for configuring VgaPriority token. 
* `vmd_enable`:(string) BIOS Token for setting VMD Enablement configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `enabled` - Enables the BIOS setting.* `disabled` - Disables the BIOS setting. 
* `work_load_config`:(string) BIOS Token for setting Workload Configuration configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `Balanced` - Value - Balanced for configuring WorkLoadConfig token.* `I/O Sensitive` - Value - I/O Sensitive for configuring WorkLoadConfig token.* `NUMA` - Value - NUMA for configuring WorkLoadConfig token.* `UMA` - Value - UMA for configuring WorkLoadConfig token. 
* `xpt_prefetch`:(string) BIOS Token for setting XPT Prefetch configuration.* `platform-default` - Default value used by the platform for the BIOS setting.* `Auto` - Value - Auto for configuring XptPrefetch token.* `disabled` - Value - disabled for configuring XptPrefetch token.* `enabled` - Value - enabled for configuring XptPrefetch token. 


## Import
`intersight_bios_policy` can be imported using the Moid of the object, e.g.
```
$ terraform import intersight_bios_policy.example 1234567890987654321abcde
``` 