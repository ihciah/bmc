package ipmi

// ConfidentialityAlgorithm is the identifier of encryption algorithms used in
// the RMCP+ session establishment process. Packets with the encryption bit set
// in the session header are encrypted as per the specification for this
// algorithm.
type ConfidentialityAlgorithm uint8

var (
	// ConfidentialityAlgorithms contains all supported confidentiality
	// algorithms in the specification. This can be used for registering metric
	// labels.
	ConfidentialityAlgorithms = []ConfidentialityAlgorithm{
		ConfidentialityAlgorithmNone,
		ConfidentialityAlgorithmAESCBC128,
		ConfidentialityAlgorithmXRC4128,
		ConfidentialityAlgorithmXRC440,
	}
)

const (
	ConfidentialityAlgorithmNone ConfidentialityAlgorithm = 0x00

	// ConfidentialityAlgorithmAESCBC128 specifies the use of AES-128-CBC (the
	// naming is to be consistent with the spec) for encrypted packets. The
	// confidentiality header in the IPMI payload is a 16-byte IV, randomly
	// generated for each message. The confidentiality trailer consists of a pad
	// of length between 0 and 15 to get the data to encrypt to be a multiple of
	// the algorithm block size (16), followed by the number of these bytes
	// added. The pad bytes start at 0x01, and increment each byte;
	// implementations must validate this. Support for this algorithm is
	// mandatory.
	ConfidentialityAlgorithmAESCBC128 ConfidentialityAlgorithm = 0x01

	ConfidentialityAlgorithmXRC4128 ConfidentialityAlgorithm = 0x02
	ConfidentialityAlgorithmXRC440  ConfidentialityAlgorithm = 0x03
)

func (c ConfidentialityAlgorithm) String() string {
	switch c {
	case ConfidentialityAlgorithmNone:
		return "None"
	case ConfidentialityAlgorithmAESCBC128:
		return "AES-CBC-128"
	case ConfidentialityAlgorithmXRC4128:
		return "xRC4-128"
	case ConfidentialityAlgorithmXRC440:
		return "xRC4-40"
	}
	if 0x30 <= c && c <= 0x3f {
		return "OEM"
	}
	return "Unknown"
}
