// Code generated by bpf2go; DO NOT EDIT.
//go:build 386 || amd64 || arm || arm64 || loong64 || mips64le || mipsle || ppc64le || riscv64

package anonmapexec

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"

	"github.com/cilium/ebpf"
)

type anonmapexecBufsK struct {
	Path   [256]int8
	Source [256]int8
}

type anonmapexecBufsT struct{ Buf [32768]int8 }

// loadAnonmapexec returns the embedded CollectionSpec for anonmapexec.
func loadAnonmapexec() (*ebpf.CollectionSpec, error) {
	reader := bytes.NewReader(_AnonmapexecBytes)
	spec, err := ebpf.LoadCollectionSpecFromReader(reader)
	if err != nil {
		return nil, fmt.Errorf("can't load anonmapexec: %w", err)
	}

	return spec, err
}

// loadAnonmapexecObjects loads anonmapexec and converts it into a struct.
//
// The following types are suitable as obj argument:
//
//	*anonmapexecObjects
//	*anonmapexecPrograms
//	*anonmapexecMaps
//
// See ebpf.CollectionSpec.LoadAndAssign documentation for details.
func loadAnonmapexecObjects(obj interface{}, opts *ebpf.CollectionOptions) error {
	spec, err := loadAnonmapexec()
	if err != nil {
		return err
	}

	return spec.LoadAndAssign(obj, opts)
}

// anonmapexecSpecs contains maps and programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type anonmapexecSpecs struct {
	anonmapexecProgramSpecs
	anonmapexecMapSpecs
	anonmapexecVariableSpecs
}

// anonmapexecProgramSpecs contains programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type anonmapexecProgramSpecs struct {
	EnforceMmapFile *ebpf.ProgramSpec `ebpf:"enforce_mmap_file"`
}

// anonmapexecMapSpecs contains maps before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type anonmapexecMapSpecs struct {
	AnonMapExecPresetContainers *ebpf.MapSpec `ebpf:"anon_map_exec_preset_containers"`
	Bufk                        *ebpf.MapSpec `ebpf:"bufk"`
	Bufs                        *ebpf.MapSpec `ebpf:"bufs"`
	BufsOff                     *ebpf.MapSpec `ebpf:"bufs_off"`
	Events                      *ebpf.MapSpec `ebpf:"events"`
	KubearmorAlertThrottle      *ebpf.MapSpec `ebpf:"kubearmor_alert_throttle"`
	KubearmorConfig             *ebpf.MapSpec `ebpf:"kubearmor_config"`
	KubearmorContainers         *ebpf.MapSpec `ebpf:"kubearmor_containers"`
	KubearmorEvents             *ebpf.MapSpec `ebpf:"kubearmor_events"`
}

// anonmapexecVariableSpecs contains global variables before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type anonmapexecVariableSpecs struct {
	Unused *ebpf.VariableSpec `ebpf:"unused"`
}

// anonmapexecObjects contains all objects after they have been loaded into the kernel.
//
// It can be passed to loadAnonmapexecObjects or ebpf.CollectionSpec.LoadAndAssign.
type anonmapexecObjects struct {
	anonmapexecPrograms
	anonmapexecMaps
	anonmapexecVariables
}

func (o *anonmapexecObjects) Close() error {
	return _AnonmapexecClose(
		&o.anonmapexecPrograms,
		&o.anonmapexecMaps,
	)
}

// anonmapexecMaps contains all maps after they have been loaded into the kernel.
//
// It can be passed to loadAnonmapexecObjects or ebpf.CollectionSpec.LoadAndAssign.
type anonmapexecMaps struct {
	AnonMapExecPresetContainers *ebpf.Map `ebpf:"anon_map_exec_preset_containers"`
	Bufk                        *ebpf.Map `ebpf:"bufk"`
	Bufs                        *ebpf.Map `ebpf:"bufs"`
	BufsOff                     *ebpf.Map `ebpf:"bufs_off"`
	Events                      *ebpf.Map `ebpf:"events"`
	KubearmorAlertThrottle      *ebpf.Map `ebpf:"kubearmor_alert_throttle"`
	KubearmorConfig             *ebpf.Map `ebpf:"kubearmor_config"`
	KubearmorContainers         *ebpf.Map `ebpf:"kubearmor_containers"`
	KubearmorEvents             *ebpf.Map `ebpf:"kubearmor_events"`
}

func (m *anonmapexecMaps) Close() error {
	return _AnonmapexecClose(
		m.AnonMapExecPresetContainers,
		m.Bufk,
		m.Bufs,
		m.BufsOff,
		m.Events,
		m.KubearmorAlertThrottle,
		m.KubearmorConfig,
		m.KubearmorContainers,
		m.KubearmorEvents,
	)
}

// anonmapexecVariables contains all global variables after they have been loaded into the kernel.
//
// It can be passed to loadAnonmapexecObjects or ebpf.CollectionSpec.LoadAndAssign.
type anonmapexecVariables struct {
	Unused *ebpf.Variable `ebpf:"unused"`
}

// anonmapexecPrograms contains all programs after they have been loaded into the kernel.
//
// It can be passed to loadAnonmapexecObjects or ebpf.CollectionSpec.LoadAndAssign.
type anonmapexecPrograms struct {
	EnforceMmapFile *ebpf.Program `ebpf:"enforce_mmap_file"`
}

func (p *anonmapexecPrograms) Close() error {
	return _AnonmapexecClose(
		p.EnforceMmapFile,
	)
}

func _AnonmapexecClose(closers ...io.Closer) error {
	for _, closer := range closers {
		if err := closer.Close(); err != nil {
			return err
		}
	}
	return nil
}

// Do not access this directly.
//
//go:embed anonmapexec_bpfel.o
var _AnonmapexecBytes []byte
