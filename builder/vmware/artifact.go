package vmware

import "fmt"

// Artifact is the result of running the VMware builder, namely a set
// of files associated with the resulting machine.
type Artifact struct {
	dir string
	f   []string
}

func (*Artifact) BuilderId() string {
	return BuilderId
}

func (a *Artifact) Files() []string {
	return a.f
}

func (*Artifact) Id() string {
	return "VM"
}

func (a *Artifact) String() string {
	return fmt.Sprintf("VM files in directory: %s", a.dir)
}