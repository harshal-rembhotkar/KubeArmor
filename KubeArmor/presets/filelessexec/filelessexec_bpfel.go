// Code generated by bpf2go; DO NOT EDIT.
//go:build 386 || amd64 || arm || arm64 || loong64 || mips64le || mipsle || ppc64le || riscv64

package filelessexec

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"

	"github.com/cilium/ebpf"
)

type filelessexecBufsK struct {
	Path   [256]int8
	Source [256]int8
}

type filelessexecBufsT struct{ Buf [32768]int8 }

// loadFilelessexec returns the embedded CollectionSpec for filelessexec.
func loadFilelessexec() (*ebpf.CollectionSpec, error) {
	reader := bytes.NewReader(_FilelessexecBytes)
	spec, err := ebpf.LoadCollectionSpecFromReader(reader)
	if err != nil {
		return nil, fmt.Errorf("can't load filelessexec: %w", err)
	}

	return spec, err
}

// loadFilelessexecObjects loads filelessexec and converts it into a struct.
//
// The following types are suitable as obj argument:
//
//	*filelessexecObjects
//	*filelessexecPrograms
//	*filelessexecMaps
//
// See ebpf.CollectionSpec.LoadAndAssign documentation for details.
func loadFilelessexecObjects(obj interface{}, opts *ebpf.CollectionOptions) error {
	spec, err := loadFilelessexec()
	if err != nil {
		return err
	}

	return spec.LoadAndAssign(obj, opts)
}

// filelessexecSpecs contains maps and programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type filelessexecSpecs struct {
	filelessexecProgramSpecs
	filelessexecMapSpecs
	filelessexecVariableSpecs
}

// filelessexecProgramSpecs contains programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type filelessexecProgramSpecs struct {
	EnforceBprmCheckSecurity *ebpf.ProgramSpec `ebpf:"enforce_bprm_check_security"`
}

// filelessexecMapSpecs contains maps before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type filelessexecMapSpecs struct {
	Bufk                         *ebpf.MapSpec `ebpf:"bufk"`
	Bufs                         *ebpf.MapSpec `ebpf:"bufs"`
	BufsOff                      *ebpf.MapSpec `ebpf:"bufs_off"`
	Events                       *ebpf.MapSpec `ebpf:"events"`
	FilelessExecPresetContainers *ebpf.MapSpec `ebpf:"fileless_exec_preset_containers"`
	KubearmorAlertThrottle       *ebpf.MapSpec `ebpf:"kubearmor_alert_throttle"`
	KubearmorConfig              *ebpf.MapSpec `ebpf:"kubearmor_config"`
	KubearmorContainers          *ebpf.MapSpec `ebpf:"kubearmor_containers"`
	KubearmorEvents              *ebpf.MapSpec `ebpf:"kubearmor_events"`
}

// filelessexecVariableSpecs contains global variables before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type filelessexecVariableSpecs struct {
	Unused *ebpf.VariableSpec `ebpf:"unused"`
}

// filelessexecObjects contains all objects after they have been loaded into the kernel.
//
// It can be passed to loadFilelessexecObjects or ebpf.CollectionSpec.LoadAndAssign.
type filelessexecObjects struct {
	filelessexecPrograms
	filelessexecMaps
	filelessexecVariables
}

func (o *filelessexecObjects) Close() error {
	return _FilelessexecClose(
		&o.filelessexecPrograms,
		&o.filelessexecMaps,
	)
}

// filelessexecMaps contains all maps after they have been loaded into the kernel.
//
// It can be passed to loadFilelessexecObjects or ebpf.CollectionSpec.LoadAndAssign.
type filelessexecMaps struct {
	Bufk                         *ebpf.Map `ebpf:"bufk"`
	Bufs                         *ebpf.Map `ebpf:"bufs"`
	BufsOff                      *ebpf.Map `ebpf:"bufs_off"`
	Events                       *ebpf.Map `ebpf:"events"`
	FilelessExecPresetContainers *ebpf.Map `ebpf:"fileless_exec_preset_containers"`
	KubearmorAlertThrottle       *ebpf.Map `ebpf:"kubearmor_alert_throttle"`
	KubearmorConfig              *ebpf.Map `ebpf:"kubearmor_config"`
	KubearmorContainers          *ebpf.Map `ebpf:"kubearmor_containers"`
	KubearmorEvents              *ebpf.Map `ebpf:"kubearmor_events"`
}

func (m *filelessexecMaps) Close() error {
	return _FilelessexecClose(
		m.Bufk,
		m.Bufs,
		m.BufsOff,
		m.Events,
		m.FilelessExecPresetContainers,
		m.KubearmorAlertThrottle,
		m.KubearmorConfig,
		m.KubearmorContainers,
		m.KubearmorEvents,
	)
}

// filelessexecVariables contains all global variables after they have been loaded into the kernel.
//
// It can be passed to loadFilelessexecObjects or ebpf.CollectionSpec.LoadAndAssign.
type filelessexecVariables struct {
	Unused *ebpf.Variable `ebpf:"unused"`
}

// filelessexecPrograms contains all programs after they have been loaded into the kernel.
//
// It can be passed to loadFilelessexecObjects or ebpf.CollectionSpec.LoadAndAssign.
type filelessexecPrograms struct {
	EnforceBprmCheckSecurity *ebpf.Program `ebpf:"enforce_bprm_check_security"`
}

func (p *filelessexecPrograms) Close() error {
	return _FilelessexecClose(
		p.EnforceBprmCheckSecurity,
	)
}

func _FilelessexecClose(closers ...io.Closer) error {
	for _, closer := range closers {
		if err := closer.Close(); err != nil {
			return err
		}
	}
	return nil
}

// Do not access this directly.
//
//go:embed filelessexec_bpfel.o
var _FilelessexecBytes []byte
