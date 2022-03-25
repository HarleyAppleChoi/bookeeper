package types

type BigChainType struct{
	ChainType string `json:"chain_type"`
	Details []IndividualChain `json:"details"`
}

type IndividualChain struct{
	ChainName string `json:"chain_name"`
	Validators []ValidatorDetail `json:"validators"`
	FundHoldingAccount []string `json:"fund_holding_account"`
	GrpcEndpoint string `json:"grpc_endpoint"`
	RpcEndpoint string `json:"rpc_endpoint"`
	LcdEndpoint string `json:"lcd_endpoint"`
}

type ValidatorDetail struct{
	ValidatorAddress string `json:"validator_address"`
	SelfDelegationAddress string `json:"self_delegation_address"`
}

type EmailAccount struct{
	Name string `json:"name"`
	Password string `json:"password"`
	Address string `json:"address"`
	Host string `json:"host"`
}

type EmailDetails struct{
	From EmailAccount `json:"from"`
	To []string `json:"to"`
	Subject string `json:"subject"`
	Details string `json:"details"`
}

type Data struct{
	Chains []BigChainType `json:"chains"`
	EmailDetails EmailDetails `json:"email_details"`
}