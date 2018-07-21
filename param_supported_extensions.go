package sctp

func chunkTypeIntersect(l, r []chunkType) (c []chunkType) {
	m := make(map[chunkType]bool)

	for _, ct := range l {
		m[ct] = true
	}

	for _, ct := range r {
		if _, ok := m[ct]; ok {
			c = append(c, ct)
		}
	}
	return
}

func newEmptySupportedExtensions() *paramSupportedExtensions {
	return &paramSupportedExtensions{}
}

type paramSupportedExtensions struct {
	paramHeader
	ChunkTypes []chunkType
}

func (s *paramSupportedExtensions) marshal() ([]byte, error) {
	s.typ = supportedExt
	s.raw = make([]byte, len(s.ChunkTypes))
	for i, c := range s.ChunkTypes {
		s.raw[i] = byte(c)
	}

	return s.paramHeader.marshal()
}

func (s *paramSupportedExtensions) unmarshal(raw []byte) (param, error) {
	s.paramHeader.unmarshal(raw)

	for _, t := range s.raw {
		s.ChunkTypes = append(s.ChunkTypes, chunkType(t))
	}

	return s, nil
}