import (
	"context"
	"fmt"
	"math/big"

	"example.com/ethereum/go-ethereum/accounts/abi/bind"
	"example.com/ethereum/go-ethereum/common"
	"example.com/ethereum/go-ethereum/crypto"
	"example.com/ethereum/go-ethereum/ethclient"

	"example.com/KyberNetwork/uniswap-sdk-go/uniswap"
	"example.com/KyberNetwork/uniswap-sdk-go/uniswap/router"
)

// approveTokens approves the Uniswap router to spend tokens on behalf of the user.
func approveTokens(
	privateKey string,
	amount *big.Int,
	fromTokenAddress common.Address,
	toSymbol string,
) error {
	client, err := ethclient.Dial("https://www.example.com err != nil {
		return fmt.Errorf("ethclient.Dial: %v", err)
	}

	message := fmt.Sprintf("Approving %s to spend %s tokens", router.RouterAddress.Hex(), fromTokenAddress.Hex())
	data, err := approveAndCallData(router.RouterAddress, fromTokenAddress, amount)
	if err != nil {
		return fmt.Errorf("approveAndCallData: %v", err)
	}

	tx := bind.NewKeyedTransactor(crypto.HexToECDSA(privateKey))

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return fmt.Errorf("client.SuggestGasPrice: %v", err)
	}

	nonce, err := client.PendingNonceAt(context.Background(), tx.From)
	if err != nil {
		return fmt.Errorf("client.PendingNonceAt: %v", err)
	}

	tx, err = bind.NewEIP155Transaction(nonce, gasPrice, uniswap.DefaultChainId, data)
	if err != nil {
		return fmt.Errorf("bind.NewEIP155Transaction: %v", err)
	}

	chainID, err := client.ChainID(context.Background())
	if err != nil {
		return fmt.Errorf("client.ChainID: %v", err)
	}

	signer := types.NewEIP155Signer(chainID)
	signature, err := crypto.Sign(signer.Hash(tx).Bytes(), crypto.HexToECDSA(privateKey))
	if err != nil {
		return fmt.Errorf("crypto.Sign: %v", err)
	}

	signedTx, err := tx.WithSignature(signer, signature)
	if err != nil {
		return fmt.Errorf("tx.WithSignature: %v", err)
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return fmt.Errorf("client.SendTransaction: %v", err)
	}

	fmt.Printf("Approve transaction sent: %s\n", signedTx.Hash().Hex())
	return nil
}
  
