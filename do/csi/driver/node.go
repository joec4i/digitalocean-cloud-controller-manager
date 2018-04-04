// Code generated by protoc-gen-go. DO NOT EDIT.

// NOTE: THIS IS NOT GENERATED. We have to add this line to prevent golint
// checking this file. This is needed because some methods end with xxxId, but
// golint wants them to be xxxID. But we're not able to change it as the
// official CSI spec is that way and we have to implement the interface
// exactly.

package driver

import (
	"context"

	csi "github.com/container-storage-interface/spec/lib/go/csi/v0"
)

// NodeStageVolume
func (d *Driver) NodeStageVolume(context.Context, *csi.NodeStageVolumeRequest) (*csi.NodeStageVolumeResponse, error) {
	panic("not implemented")
}

// NodeUnstageVolume ...
func (d *Driver) NodeUnstageVolume(context.Context, *csi.NodeUnstageVolumeRequest) (*csi.NodeUnstageVolumeResponse, error) {
	panic("not implemented")
}

// NodePublishVolume ...
func (d *Driver) NodePublishVolume(context.Context, *csi.NodePublishVolumeRequest) (*csi.NodePublishVolumeResponse, error) {
	panic("not implemented")
}

// NodeUnpublishVolume ...
func (d *Driver) NodeUnpublishVolume(context.Context, *csi.NodeUnpublishVolumeRequest) (*csi.NodeUnpublishVolumeResponse, error) {
	panic("not implemented")
}

// NodeGetId ...
func (d *Driver) NodeGetId(context.Context, *csi.NodeGetIdRequest) (*csi.NodeGetIdResponse, error) {
	panic("not implemented")
}

// NodeGetCapabilities ...
func (d *Driver) NodeGetCapabilities(context.Context, *csi.NodeGetCapabilitiesRequest) (*csi.NodeGetCapabilitiesResponse, error) {
	panic("not implemented")
}
