package project

import (
	"github.com/giantswarm/versionbundle"
)

func NewVersionBundle() versionbundle.Bundle {
	return versionbundle.Bundle{
		Changelogs: []versionbundle.Changelog{
			{
				Component:   "cloudconfig",
				Description: "Fix pause container image repository for China.",
				Kind:        versionbundle.KindFixed,
			},
			{
				Component:   "vault",
				Description: "Fix vault encrypter role with new nodepools iam role names.",
				Kind:        versionbundle.KindFixed,
			},
			{
				Component:   "cloudformation",
				Description: "Propagate tag name from ASG to EC2 instances.",
				Kind:        versionbundle.KindFixed,
				URLs: []string{
					"https://github.com/giantswarm/aws-operator/pull/2110",
				},
			},
			{
				Component:   "cloudformation",
				Description: "Drain nodes when deleting Node Pools.",
				Kind:        versionbundle.KindFixed,
				URLs: []string{
					"https://github.com/giantswarm/aws-operator/pull/2111",
				},
			},
		},
		Components: []versionbundle.Component{
			{
				Name:    "calico",
				Version: "3.9.1",
			},
			{
				Name:    "containerlinux",
				Version: "2191.5.0",
			},
			{
				Name:    "docker",
				Version: "18.06.1",
			},
			{
				Name:    "etcd",
				Version: "3.3.15",
			},
			{
				Name:    "kubernetes",
				Version: "1.15.5",
			},
		},
		Name:    Name(),
		Version: BundleVersion(),
	}
}
