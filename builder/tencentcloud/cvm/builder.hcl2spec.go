// Code generated by "mapstructure-to-hcl2 -type Config"; DO NOT EDIT.
package cvm

import (
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/zclconf/go-cty/cty"
)

// FlatConfig is an auto-generated flat version of Config.
// Where the contents of a field with a `mapstructure:,squash` tag are bubbled up.
type FlatConfig struct {
	PackerBuildName           *string                    `mapstructure:"packer_build_name" cty:"packer_build_name"`
	PackerBuilderType         *string                    `mapstructure:"packer_builder_type" cty:"packer_builder_type"`
	PackerDebug               *bool                      `mapstructure:"packer_debug" cty:"packer_debug"`
	PackerForce               *bool                      `mapstructure:"packer_force" cty:"packer_force"`
	PackerOnError             *string                    `mapstructure:"packer_on_error" cty:"packer_on_error"`
	PackerUserVars            map[string]string          `mapstructure:"packer_user_variables" cty:"packer_user_variables"`
	PackerSensitiveVars       []string                   `mapstructure:"packer_sensitive_variables" cty:"packer_sensitive_variables"`
	SecretId                  *string                    `mapstructure:"secret_id" required:"true" cty:"secret_id"`
	SecretKey                 *string                    `mapstructure:"secret_key" required:"true" cty:"secret_key"`
	Region                    *string                    `mapstructure:"region" required:"true" cty:"region"`
	Zone                      *string                    `mapstructure:"zone" required:"true" cty:"zone"`
	SkipValidation            *bool                      `mapstructure:"skip_region_validation" required:"false" cty:"skip_region_validation"`
	ImageName                 *string                    `mapstructure:"image_name" required:"true" cty:"image_name"`
	ImageDescription          *string                    `mapstructure:"image_description" required:"false" cty:"image_description"`
	Reboot                    *bool                      `mapstructure:"reboot" required:"false" cty:"reboot"`
	ForcePoweroff             *bool                      `mapstructure:"force_poweroff" required:"false" cty:"force_poweroff"`
	Sysprep                   *bool                      `mapstructure:"sysprep" required:"false" cty:"sysprep"`
	ImageForceDelete          *bool                      `mapstructure:"image_force_delete" cty:"image_force_delete"`
	ImageCopyRegions          []string                   `mapstructure:"image_copy_regions" required:"false" cty:"image_copy_regions"`
	ImageShareAccounts        []string                   `mapstructure:"image_share_accounts" required:"false" cty:"image_share_accounts"`
	AssociatePublicIpAddress  *bool                      `mapstructure:"associate_public_ip_address" required:"false" cty:"associate_public_ip_address"`
	SourceImageId             *string                    `mapstructure:"source_image_id" required:"true" cty:"source_image_id"`
	InstanceType              *string                    `mapstructure:"instance_type" required:"true" cty:"instance_type"`
	InstanceName              *string                    `mapstructure:"instance_name" required:"false" cty:"instance_name"`
	DiskType                  *string                    `mapstructure:"disk_type" required:"false" cty:"disk_type"`
	DiskSize                  *int64                     `mapstructure:"disk_size" required:"false" cty:"disk_size"`
	DataDisks                 []FlattencentCloudDataDisk `mapstructure:"data_disks" cty:"data_disks"`
	VpcId                     *string                    `mapstructure:"vpc_id" required:"false" cty:"vpc_id"`
	VpcName                   *string                    `mapstructure:"vpc_name" required:"false" cty:"vpc_name"`
	VpcIp                     *string                    `mapstructure:"vpc_ip" cty:"vpc_ip"`
	SubnetId                  *string                    `mapstructure:"subnet_id" required:"false" cty:"subnet_id"`
	SubnetName                *string                    `mapstructure:"subnet_name" required:"false" cty:"subnet_name"`
	CidrBlock                 *string                    `mapstructure:"cidr_block" required:"false" cty:"cidr_block"`
	SubnectCidrBlock          *string                    `mapstructure:"subnect_cidr_block" required:"false" cty:"subnect_cidr_block"`
	InternetChargeType        *string                    `mapstructure:"internet_charge_type" cty:"internet_charge_type"`
	InternetMaxBandwidthOut   *int64                     `mapstructure:"internet_max_bandwidth_out" required:"false" cty:"internet_max_bandwidth_out"`
	SecurityGroupId           *string                    `mapstructure:"security_group_id" required:"false" cty:"security_group_id"`
	SecurityGroupName         *string                    `mapstructure:"security_group_name" required:"false" cty:"security_group_name"`
	UserData                  *string                    `mapstructure:"user_data" required:"false" cty:"user_data"`
	UserDataFile              *string                    `mapstructure:"user_data_file" required:"false" cty:"user_data_file"`
	HostName                  *string                    `mapstructure:"host_name" required:"false" cty:"host_name"`
	RunTags                   map[string]string          `mapstructure:"run_tags" required:"false" cty:"run_tags"`
	Type                      *string                    `mapstructure:"communicator" cty:"communicator"`
	PauseBeforeConnect        *string                    `mapstructure:"pause_before_connecting" cty:"pause_before_connecting"`
	SSHHost                   *string                    `mapstructure:"ssh_host" cty:"ssh_host"`
	SSHPort                   *int                       `mapstructure:"ssh_port" cty:"ssh_port"`
	SSHUsername               *string                    `mapstructure:"ssh_username" cty:"ssh_username"`
	SSHPassword               *string                    `mapstructure:"ssh_password" cty:"ssh_password"`
	SSHKeyPairName            *string                    `mapstructure:"ssh_keypair_name" cty:"ssh_keypair_name"`
	SSHTemporaryKeyPairName   *string                    `mapstructure:"temporary_key_pair_name" cty:"temporary_key_pair_name"`
	SSHClearAuthorizedKeys    *bool                      `mapstructure:"ssh_clear_authorized_keys" cty:"ssh_clear_authorized_keys"`
	SSHPrivateKeyFile         *string                    `mapstructure:"ssh_private_key_file" cty:"ssh_private_key_file"`
	SSHPty                    *bool                      `mapstructure:"ssh_pty" cty:"ssh_pty"`
	SSHTimeout                *string                    `mapstructure:"ssh_timeout" cty:"ssh_timeout"`
	SSHAgentAuth              *bool                      `mapstructure:"ssh_agent_auth" cty:"ssh_agent_auth"`
	SSHDisableAgentForwarding *bool                      `mapstructure:"ssh_disable_agent_forwarding" cty:"ssh_disable_agent_forwarding"`
	SSHHandshakeAttempts      *int                       `mapstructure:"ssh_handshake_attempts" cty:"ssh_handshake_attempts"`
	SSHBastionHost            *string                    `mapstructure:"ssh_bastion_host" cty:"ssh_bastion_host"`
	SSHBastionPort            *int                       `mapstructure:"ssh_bastion_port" cty:"ssh_bastion_port"`
	SSHBastionAgentAuth       *bool                      `mapstructure:"ssh_bastion_agent_auth" cty:"ssh_bastion_agent_auth"`
	SSHBastionUsername        *string                    `mapstructure:"ssh_bastion_username" cty:"ssh_bastion_username"`
	SSHBastionPassword        *string                    `mapstructure:"ssh_bastion_password" cty:"ssh_bastion_password"`
	SSHBastionPrivateKeyFile  *string                    `mapstructure:"ssh_bastion_private_key_file" cty:"ssh_bastion_private_key_file"`
	SSHFileTransferMethod     *string                    `mapstructure:"ssh_file_transfer_method" cty:"ssh_file_transfer_method"`
	SSHProxyHost              *string                    `mapstructure:"ssh_proxy_host" cty:"ssh_proxy_host"`
	SSHProxyPort              *int                       `mapstructure:"ssh_proxy_port" cty:"ssh_proxy_port"`
	SSHProxyUsername          *string                    `mapstructure:"ssh_proxy_username" cty:"ssh_proxy_username"`
	SSHProxyPassword          *string                    `mapstructure:"ssh_proxy_password" cty:"ssh_proxy_password"`
	SSHKeepAliveInterval      *string                    `mapstructure:"ssh_keep_alive_interval" cty:"ssh_keep_alive_interval"`
	SSHReadWriteTimeout       *string                    `mapstructure:"ssh_read_write_timeout" cty:"ssh_read_write_timeout"`
	SSHRemoteTunnels          []string                   `mapstructure:"ssh_remote_tunnels" cty:"ssh_remote_tunnels"`
	SSHLocalTunnels           []string                   `mapstructure:"ssh_local_tunnels" cty:"ssh_local_tunnels"`
	SSHPublicKey              []byte                     `cty:"ssh_public_key"`
	SSHPrivateKey             []byte                     `cty:"ssh_private_key"`
	WinRMUser                 *string                    `mapstructure:"winrm_username" cty:"winrm_username"`
	WinRMPassword             *string                    `mapstructure:"winrm_password" cty:"winrm_password"`
	WinRMHost                 *string                    `mapstructure:"winrm_host" cty:"winrm_host"`
	WinRMPort                 *int                       `mapstructure:"winrm_port" cty:"winrm_port"`
	WinRMTimeout              *string                    `mapstructure:"winrm_timeout" cty:"winrm_timeout"`
	WinRMUseSSL               *bool                      `mapstructure:"winrm_use_ssl" cty:"winrm_use_ssl"`
	WinRMInsecure             *bool                      `mapstructure:"winrm_insecure" cty:"winrm_insecure"`
	WinRMUseNTLM              *bool                      `mapstructure:"winrm_use_ntlm" cty:"winrm_use_ntlm"`
	SSHPrivateIp              *bool                      `mapstructure:"ssh_private_ip" cty:"ssh_private_ip"`
}

// FlatMapstructure returns a new FlatConfig.
// FlatConfig is an auto-generated flat version of Config.
// Where the contents a fields with a `mapstructure:,squash` tag are bubbled up.
func (*Config) FlatMapstructure() interface{} { return new(FlatConfig) }

// HCL2Spec returns the hcldec.Spec of a FlatConfig.
// This spec is used by HCL to read the fields of FlatConfig.
func (*FlatConfig) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"packer_build_name":            &hcldec.AttrSpec{Name: "packer_build_name", Type: cty.String, Required: false},
		"packer_builder_type":          &hcldec.AttrSpec{Name: "packer_builder_type", Type: cty.String, Required: false},
		"packer_debug":                 &hcldec.AttrSpec{Name: "packer_debug", Type: cty.Bool, Required: false},
		"packer_force":                 &hcldec.AttrSpec{Name: "packer_force", Type: cty.Bool, Required: false},
		"packer_on_error":              &hcldec.AttrSpec{Name: "packer_on_error", Type: cty.String, Required: false},
		"packer_user_variables":        &hcldec.BlockAttrsSpec{TypeName: "packer_user_variables", ElementType: cty.String, Required: false},
		"packer_sensitive_variables":   &hcldec.AttrSpec{Name: "packer_sensitive_variables", Type: cty.List(cty.String), Required: false},
		"secret_id":                    &hcldec.AttrSpec{Name: "secret_id", Type: cty.String, Required: false},
		"secret_key":                   &hcldec.AttrSpec{Name: "secret_key", Type: cty.String, Required: false},
		"region":                       &hcldec.AttrSpec{Name: "region", Type: cty.String, Required: false},
		"zone":                         &hcldec.AttrSpec{Name: "zone", Type: cty.String, Required: false},
		"skip_region_validation":       &hcldec.AttrSpec{Name: "skip_region_validation", Type: cty.Bool, Required: false},
		"image_name":                   &hcldec.AttrSpec{Name: "image_name", Type: cty.String, Required: false},
		"image_description":            &hcldec.AttrSpec{Name: "image_description", Type: cty.String, Required: false},
		"reboot":                       &hcldec.AttrSpec{Name: "reboot", Type: cty.Bool, Required: false},
		"force_poweroff":               &hcldec.AttrSpec{Name: "force_poweroff", Type: cty.Bool, Required: false},
		"sysprep":                      &hcldec.AttrSpec{Name: "sysprep", Type: cty.Bool, Required: false},
		"image_force_delete":           &hcldec.AttrSpec{Name: "image_force_delete", Type: cty.Bool, Required: false},
		"image_copy_regions":           &hcldec.AttrSpec{Name: "image_copy_regions", Type: cty.List(cty.String), Required: false},
		"image_share_accounts":         &hcldec.AttrSpec{Name: "image_share_accounts", Type: cty.List(cty.String), Required: false},
		"associate_public_ip_address":  &hcldec.AttrSpec{Name: "associate_public_ip_address", Type: cty.Bool, Required: false},
		"source_image_id":              &hcldec.AttrSpec{Name: "source_image_id", Type: cty.String, Required: false},
		"instance_type":                &hcldec.AttrSpec{Name: "instance_type", Type: cty.String, Required: false},
		"instance_name":                &hcldec.AttrSpec{Name: "instance_name", Type: cty.String, Required: false},
		"disk_type":                    &hcldec.AttrSpec{Name: "disk_type", Type: cty.String, Required: false},
		"disk_size":                    &hcldec.AttrSpec{Name: "disk_size", Type: cty.Number, Required: false},
		"data_disks":                   &hcldec.BlockListSpec{TypeName: "data_disks", Nested: &hcldec.BlockSpec{TypeName: "data_disks", Nested: hcldec.ObjectSpec((*FlattencentCloudDataDisk)(nil).HCL2Spec())}},
		"vpc_id":                       &hcldec.AttrSpec{Name: "vpc_id", Type: cty.String, Required: false},
		"vpc_name":                     &hcldec.AttrSpec{Name: "vpc_name", Type: cty.String, Required: false},
		"vpc_ip":                       &hcldec.AttrSpec{Name: "vpc_ip", Type: cty.String, Required: false},
		"subnet_id":                    &hcldec.AttrSpec{Name: "subnet_id", Type: cty.String, Required: false},
		"subnet_name":                  &hcldec.AttrSpec{Name: "subnet_name", Type: cty.String, Required: false},
		"cidr_block":                   &hcldec.AttrSpec{Name: "cidr_block", Type: cty.String, Required: false},
		"subnect_cidr_block":           &hcldec.AttrSpec{Name: "subnect_cidr_block", Type: cty.String, Required: false},
		"internet_charge_type":         &hcldec.AttrSpec{Name: "internet_charge_type", Type: cty.String, Required: false},
		"internet_max_bandwidth_out":   &hcldec.AttrSpec{Name: "internet_max_bandwidth_out", Type: cty.Number, Required: false},
		"security_group_id":            &hcldec.AttrSpec{Name: "security_group_id", Type: cty.String, Required: false},
		"security_group_name":          &hcldec.AttrSpec{Name: "security_group_name", Type: cty.String, Required: false},
		"user_data":                    &hcldec.AttrSpec{Name: "user_data", Type: cty.String, Required: false},
		"user_data_file":               &hcldec.AttrSpec{Name: "user_data_file", Type: cty.String, Required: false},
		"host_name":                    &hcldec.AttrSpec{Name: "host_name", Type: cty.String, Required: false},
		"run_tags":                     &hcldec.BlockAttrsSpec{TypeName: "run_tags", ElementType: cty.String, Required: false},
		"communicator":                 &hcldec.AttrSpec{Name: "communicator", Type: cty.String, Required: false},
		"pause_before_connecting":      &hcldec.AttrSpec{Name: "pause_before_connecting", Type: cty.String, Required: false},
		"ssh_host":                     &hcldec.AttrSpec{Name: "ssh_host", Type: cty.String, Required: false},
		"ssh_port":                     &hcldec.AttrSpec{Name: "ssh_port", Type: cty.Number, Required: false},
		"ssh_username":                 &hcldec.AttrSpec{Name: "ssh_username", Type: cty.String, Required: false},
		"ssh_password":                 &hcldec.AttrSpec{Name: "ssh_password", Type: cty.String, Required: false},
		"ssh_keypair_name":             &hcldec.AttrSpec{Name: "ssh_keypair_name", Type: cty.String, Required: false},
		"temporary_key_pair_name":      &hcldec.AttrSpec{Name: "temporary_key_pair_name", Type: cty.String, Required: false},
		"ssh_clear_authorized_keys":    &hcldec.AttrSpec{Name: "ssh_clear_authorized_keys", Type: cty.Bool, Required: false},
		"ssh_private_key_file":         &hcldec.AttrSpec{Name: "ssh_private_key_file", Type: cty.String, Required: false},
		"ssh_pty":                      &hcldec.AttrSpec{Name: "ssh_pty", Type: cty.Bool, Required: false},
		"ssh_timeout":                  &hcldec.AttrSpec{Name: "ssh_timeout", Type: cty.String, Required: false},
		"ssh_agent_auth":               &hcldec.AttrSpec{Name: "ssh_agent_auth", Type: cty.Bool, Required: false},
		"ssh_disable_agent_forwarding": &hcldec.AttrSpec{Name: "ssh_disable_agent_forwarding", Type: cty.Bool, Required: false},
		"ssh_handshake_attempts":       &hcldec.AttrSpec{Name: "ssh_handshake_attempts", Type: cty.Number, Required: false},
		"ssh_bastion_host":             &hcldec.AttrSpec{Name: "ssh_bastion_host", Type: cty.String, Required: false},
		"ssh_bastion_port":             &hcldec.AttrSpec{Name: "ssh_bastion_port", Type: cty.Number, Required: false},
		"ssh_bastion_agent_auth":       &hcldec.AttrSpec{Name: "ssh_bastion_agent_auth", Type: cty.Bool, Required: false},
		"ssh_bastion_username":         &hcldec.AttrSpec{Name: "ssh_bastion_username", Type: cty.String, Required: false},
		"ssh_bastion_password":         &hcldec.AttrSpec{Name: "ssh_bastion_password", Type: cty.String, Required: false},
		"ssh_bastion_private_key_file": &hcldec.AttrSpec{Name: "ssh_bastion_private_key_file", Type: cty.String, Required: false},
		"ssh_file_transfer_method":     &hcldec.AttrSpec{Name: "ssh_file_transfer_method", Type: cty.String, Required: false},
		"ssh_proxy_host":               &hcldec.AttrSpec{Name: "ssh_proxy_host", Type: cty.String, Required: false},
		"ssh_proxy_port":               &hcldec.AttrSpec{Name: "ssh_proxy_port", Type: cty.Number, Required: false},
		"ssh_proxy_username":           &hcldec.AttrSpec{Name: "ssh_proxy_username", Type: cty.String, Required: false},
		"ssh_proxy_password":           &hcldec.AttrSpec{Name: "ssh_proxy_password", Type: cty.String, Required: false},
		"ssh_keep_alive_interval":      &hcldec.AttrSpec{Name: "ssh_keep_alive_interval", Type: cty.String, Required: false},
		"ssh_read_write_timeout":       &hcldec.AttrSpec{Name: "ssh_read_write_timeout", Type: cty.String, Required: false},
		"ssh_remote_tunnels":           &hcldec.AttrSpec{Name: "ssh_remote_tunnels", Type: cty.List(cty.String), Required: false},
		"ssh_local_tunnels":            &hcldec.AttrSpec{Name: "ssh_local_tunnels", Type: cty.List(cty.String), Required: false},
		"ssh_public_key":               &hcldec.AttrSpec{Name: "ssh_public_key", Type: cty.List(cty.Number), Required: false},
		"ssh_private_key":              &hcldec.AttrSpec{Name: "ssh_private_key", Type: cty.List(cty.Number), Required: false},
		"winrm_username":               &hcldec.AttrSpec{Name: "winrm_username", Type: cty.String, Required: false},
		"winrm_password":               &hcldec.AttrSpec{Name: "winrm_password", Type: cty.String, Required: false},
		"winrm_host":                   &hcldec.AttrSpec{Name: "winrm_host", Type: cty.String, Required: false},
		"winrm_port":                   &hcldec.AttrSpec{Name: "winrm_port", Type: cty.Number, Required: false},
		"winrm_timeout":                &hcldec.AttrSpec{Name: "winrm_timeout", Type: cty.String, Required: false},
		"winrm_use_ssl":                &hcldec.AttrSpec{Name: "winrm_use_ssl", Type: cty.Bool, Required: false},
		"winrm_insecure":               &hcldec.AttrSpec{Name: "winrm_insecure", Type: cty.Bool, Required: false},
		"winrm_use_ntlm":               &hcldec.AttrSpec{Name: "winrm_use_ntlm", Type: cty.Bool, Required: false},
		"ssh_private_ip":               &hcldec.AttrSpec{Name: "ssh_private_ip", Type: cty.Bool, Required: false},
	}
	return s
}
