package code_based_hash

// Chaos = Code-based Hash Algorithm On Security

type chaos struct {
}

func New() *chaos {
	return &chaos{}
}

func (receiver *chaos) Write(bytes []byte) {

}

func (receiver *chaos) Sum(bytes []byte) []byte {
	return nil
}
