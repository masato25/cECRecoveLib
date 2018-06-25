package cerclib

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/signer/core"
	"strings"
)

type EcRevoerFeilds struct {
	EthAddress        string `form:"eth_address" json:"eth_address" binding:"required"`
	EthereumSignature string `form:"eth_signature" json:"eth_signature" binding:"required"`
	Msg               string `form:"msg" json:"msg"`
}

func signHash(data []byte) []byte {
	msg := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(data), data)
	return crypto.Keccak256([]byte(msg))
}

func decodeEthHash(inputs EcRevoerFeilds) (recoveredAddrHex string, err error) {
	sig := hexutil.MustDecode(inputs.EthereumSignature)
	sigMg := []byte(inputs.Msg)
	personal := core.SignerAPI{}
	recoveredAddr, err := personal.EcRecover(nil, sigMg, sig)
	if err == nil {
		recoveredAddrHex = fmt.Sprintf("0x%x", recoveredAddr)
		recoveredAddrHex = strings.ToLower(recoveredAddrHex)
	}
	return
}

func Resolve(inputs EcRevoerFeilds) (*string, error) {
	recoveredAddrHex, err := decodeEthHash(inputs)
	if err != nil {
		return nil, err
	}
	return &recoveredAddrHex, err
}
