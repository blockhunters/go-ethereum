package types


func SetSignatureValues(tx *Transaction, signer Signer, v, r, s *big.Int) (*Transaction) {
	// Copied from WithSignature func
	// https://github.com/ethereum/go-ethereum/blob/c503f98f6d5e80e079c1d8a3601d188af2a899da/core/types/transaction.go#L397

	cpy := tx.inner.copy()
	cpy.setSignatureValues(signer.ChainID(), v, r, s)
	return &Transaction{inner: cpy, time: tx.time}
}