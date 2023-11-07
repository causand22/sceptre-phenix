package types

import (
	"phenix/store"
	v1 "phenix/types/version/v1"
)

type Image struct {
	Metadata store.ConfigMetadata `json:"metadata" yaml:"metadata"`
	Spec     *v1.Image            `json:"spec" yaml:"spec"`
	Status   *v1.ImageStatus      `json:"status" yaml:"status"`
}

func (this *Image) GetCreatedBy() string {
	if entry, ok := this.Metadata.Annotations["created_by"]; ok {
		return entry
	}
	return ""
}
