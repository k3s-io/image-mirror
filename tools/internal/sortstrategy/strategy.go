package sortstrategy

import (
	"fmt"
	"sort"

	"github.com/Masterminds/semver/v3"
)

func Semver(t []string) ([]string, error) {
	vs := make([]*semver.Version, 0, len(t))
	for _, r := range t {
		v, err := semver.NewVersion(r)
		if err != nil {
			return nil, fmt.Errorf("error parsing version: %s", err)
		}
		vs = append(vs, v)

	}
	sort.Sort(semver.Collection(vs))

	st := make([]string, 0, len(t))
	for _, t := range vs {
		st = append(st, t.String())
	}
	return st, nil
}
