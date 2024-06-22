package protovalidate

import "google.golang.org/protobuf/proto"

type PathSet map[string]struct{}

// NewPathSet creates a string set.
// Duplicate keys are de-duped.
func NewPathSet(paths ...string) PathSet {
	ps := make(PathSet)
	for _, p := range paths {
		ps[p] = struct{}{}
	}
	return ps
}

func (ps PathSet) Empty() bool {
	return ps == nil || len(ps) == 0
}

func (ps PathSet) Has(e string) bool {
	_, ok := ps[e]
	return ok
}

func pathSetFromMsg[T proto.Message](msg T) PathSet {
	return nil
}