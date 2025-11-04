package verify

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"golang.org/x/time/rate"
)

type EtherscanGenericResp struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Result  string `json:"result"`
}

type EtherscanContractCreationResp struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Result  []struct {
		ContractCreator string `json:"contractCreator"`
		TxHash          string `json:"txHash"`
	} `json:"result"`
}

type EtherscanClient struct {
	apiKey      string
	chainID     uint64
	url         string
	rateLimiter *rate.Limiter
}

func getAPIEndpoint(l1ChainID uint64) (string, error) {
	switch l1ChainID {
	case 1:
		return "https://api.etherscan.io/v2/api", nil // eth-mainnet
	case 11155111:
		return "https://api-sepolia.etherscan.io/v2/api", nil // eth-sepolia
	case 84532:
		return "https://api-sepolia.basescan.org/v2/api", nil // base-sepolia
	default:
		return "", fmt.Errorf("unsupported L1 chain ID: %d", l1ChainID)
	}
}

func NewEtherscanClient(apiKey string, chainID uint64, url string, rateLimiter *rate.Limiter) *EtherscanClient {
	return &EtherscanClient{
		apiKey:      apiKey,
		chainID:     chainID,
		url:         url,
		rateLimiter: rateLimiter,
	}
}

// APIChecker implementation for EtherscanClient (V2 API)
func (c *EtherscanClient) CanCheck() bool {
	return c.apiKey != ""
}

func (c *EtherscanClient) GetDefaultURL(chainID uint64) (string, error) {
	return getAPIEndpoint(chainID)
}

func (c *EtherscanClient) GetChainArg(chainID uint64) (string, error) {
	return getChainName(chainID)
}

func (c *EtherscanClient) CheckStatus(ctx context.Context, address common.Address) (*VerificationStatus, error) {
	verified, err := c.isVerified(address)
	if err != nil {
		return nil, err
	}
	return &VerificationStatus{
		IsVerified:          verified,
		IsFullyVerified:     verified,
		IsPartiallyVerified: false,
	}, nil
}

func getChainName(chainID uint64) (string, error) {
	switch chainID {
	case 1:
		return "mainnet", nil
	case 11155111:
		return "sepolia", nil
	default:
		return "", fmt.Errorf("unsupported chain ID: %d", chainID)
	}
}

// sendRateLimitedRequest is a helper function which waits for a rate limit token
// before sending a request
func (c *EtherscanClient) sendRateLimitedRequest(req *http.Request) (*http.Response, error) {
	if err := c.rateLimiter.Wait(context.Background()); err != nil {
		return nil, fmt.Errorf("rate limiter error: %w", err)
	}
	return http.DefaultClient.Do(req)
}

// getContractCreation returns the txHash of the contract creation tx
// (useful for extracting constructor args)
func (c *EtherscanClient) getContractCreation(address common.Address) (common.Hash, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s?chainid=%d&module=contract&action=getcontractcreation&contractaddresses=%s&apikey=%s",
		c.url, c.chainID, address.Hex(), c.apiKey), nil)
	if err != nil {
		return common.Hash{}, fmt.Errorf("failed to create contract creation request: %w", err)
	}

	resp, err := c.sendRateLimitedRequest(req)
	if err != nil {
		return common.Hash{}, fmt.Errorf("failed to send contract creation request: %w", err)
	}
	defer resp.Body.Close()

	var creationResp EtherscanContractCreationResp
	if err := json.NewDecoder(resp.Body).Decode(&creationResp); err != nil {
		return common.Hash{}, fmt.Errorf("failed to decode contract creation response: %w", err)
	}
	if creationResp.Status != "1" {
		return common.Hash{}, fmt.Errorf("contract creation query failed: %s", creationResp.Message)
	}

	txHash := common.HexToHash(creationResp.Result[0].TxHash)
	return txHash, nil
}

func (c *EtherscanClient) isVerified(address common.Address) (bool, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s?chainid=%d&module=contract&action=getabi&address=%s&apikey=%s",
		c.url, c.chainID, address.Hex(), c.apiKey), nil)
	if err != nil {
		return false, err
	}

	resp, err := c.sendRateLimitedRequest(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	var result EtherscanGenericResp
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	return result.Status == "1", nil
}

func (c *EtherscanClient) pollVerificationStatus(reqId string) error {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s?chainid=%d&apikey=%s&module=contract&action=checkverifystatus&guid=%s",
		c.url, c.chainID, c.apiKey, reqId), nil)
	if err != nil {
		return fmt.Errorf("failed to create checkverifystatus request: %w", err)
	}

	for i := 0; i < 10; i++ { // Try 10 times with increasing delays
		resp, err := c.sendRateLimitedRequest(req)
		if err != nil {
			return fmt.Errorf("failed to send checkverifystatus request: %w", err)
		}
		defer resp.Body.Close()

		var result EtherscanGenericResp
		if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
			return fmt.Errorf("failed to decode checkverifystatus response: %w", err)
		}

		if result.Status == "1" {
			return nil
		}
		if result.Result == "Already Verified" {
			return nil
		}
		if result.Result != "Pending in queue" {
			return fmt.Errorf("verification failed: %s, %s", result.Result, result.Message)
		}
		time.Sleep(time.Duration(i+2) * time.Second)
	}
	return fmt.Errorf("verification timed out")
}
