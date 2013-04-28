package pckg

import (
	"composer/misc"
	"composer/net"
	"encoding/json"
	"fmt"
	"github.com/mcuadros/go-version"
	"sort"
)

type Pckg struct {
	Name      string
	Versions  map[string]*Version
	requested bool
}

var cache = make(map[string]*Pckg, 0)

func NewPckg(name string) *Pckg {
	if pckg, ok := cache[name]; ok {
		return pckg
	}

	cache[name] = new(Pckg)
	cache[name].SetName(name)

	return cache[name]
}

func (self *Pckg) GetVersion(versionConstraint string) (*Version, error) {
	self.request()

	constraint := version.NewConstrainGroupFromString(versionConstraint)
	for _, version := range sortVersions(self.Versions) {
		if constraint.Match(version.Version) {
			misc.GetOutput().Debug("Version matched %s = %s", versionConstraint, version.Version)
			return version, nil
		} else {
			misc.GetOutput().Debug("Version un-matched %s = %s", versionConstraint, version.Version)

		}
	}

	return nil, fmt.Errorf("cannot find given version %s at %s, options:", versionConstraint, self.Name, self.Versions)
}

func (self *Pckg) GetName() string {
	return self.Name
}

func (self *Pckg) SetName(name string) {
	self.Name = name
}

func (self *Pckg) Print() {
	self.request()

	fmt.Printf("Package: %s\n", self.Name)
	for _, version := range self.Versions {
		fmt.Printf("\t%-20s (%-20s)\n", version.Name, version.Version)
	}
}

func (self *Pckg) request() bool {
	if self.requested {
		return true
	}

	packagist := net.Packagist{}
	versions := packagist.GetRawVersion(self.Name)

	for number, raw := range versions {
		version := Version{}
		err := json.Unmarshal(raw, &version)
		if err != nil {
			fmt.Printf("Version: %s Error:%s Data:%s\n", number, err)
		}

		self.addVersion(number, &version)

	}

	self.requested = true
	return true
}

func (self *Pckg) addVersion(number string, version *Version) {
	if self.Versions == nil {
		self.Versions = make(map[string]*Version)
	}

	self.Versions[number] = version
}

type VersionSorter struct {
	Normalized []string
	Versions   []*Version
}

func sortVersions(m map[string]*Version) []*Version {
	vs := &VersionSorter{
		Normalized: make([]string, 0, len(m)),
		Versions:   make([]*Version, 0, len(m)),
	}

	for _, v := range m {
		normalized := version.Normalize(v.Version)

		vs.Normalized = append(vs.Normalized, normalized)
		vs.Versions = append(vs.Versions, v)
	}

	vs.Sort()
	return vs.Versions
}

func (vs *VersionSorter) Sort() {
	sort.Sort(vs)
}

func (vs *VersionSorter) Len() int {
	return len(vs.Versions)
}

func (vs *VersionSorter) Less(i, j int) bool {
	return vs.Normalized[i] < vs.Normalized[j]
}

func (vs *VersionSorter) Swap(i, j int) {
	vs.Versions[i], vs.Versions[j] = vs.Versions[j], vs.Versions[i]
	vs.Normalized[i], vs.Normalized[j] = vs.Normalized[j], vs.Normalized[i]
}
