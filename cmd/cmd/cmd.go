package parse

import (
	"fmt"
	//"time"

	"io/ioutil"

	//"net/http"

	//"github.com/forbole/bookkeeper/balancesheet"
	"github.com/forbole/bookkeeper/types"
	"github.com/forbole/bookkeeper/input"
	"github.com/forbole/bookkeeper/email"


	//coingecko "github.com/superoo7/go-gecko/v3"

	"github.com/spf13/cobra"
	
)

const(
	flagInputJsonPath = "input_json_path"
)

// ParseCmd returns the command that should be run when we want to start parsing a chain state.
func ParseCmd() *cobra.Command {
	cmd := cobra.Command{
		Use:     "bookkeeper",
		Short:   "Start parsing the blockchain data",
		RunE: Execute,
	}
	cmd.Flags().String(flagInputJsonPath,"./input.json", "The path that the input file should read from")
    return &cmd
}


func Execute(cmd *cobra.Command, arg []string)error {
	jsonPath, _ := cmd.Flags().GetString(flagInputJsonPath)

	data,err:=input.ImportJsonInput(jsonPath)
	if err!=nil{
		return err
	}
	fmt.Println(*data)

	inputfile:=[]string{"bitcoin.csv","ethereum.csv"}

	err=email.SendEmail(data.EmailDetails,inputfile)
	if err!=nil{
		return err
	}

	//coingecko
	/* httpClient := &http.Client{
		Timeout: time.Second * 10,
	}
	CG := coingecko.NewClient(httpClient)

	// Import coin info
	// Assume the dateQuantity pair is always by time asc
	coin := types.NewCoin("bitcoin", []types.DateQuantity{
		types.NewDateQuantity(5.1234,
			time.Date(2020, time.January, 28, 0, 0, 0, 0, time.UTC)),
		types.NewDateQuantity(15.5642,
			time.Date(2020, time.August, 28, 0, 0, 0, 0, time.UTC)),
	})

	eth := types.NewCoin("ethereum", []types.DateQuantity{
		types.NewDateQuantity(30.4501,
			time.Date(2020, time.January, 28, 0, 0, 0, 0, time.UTC)),
		types.NewDateQuantity(16.4564,
			time.Date(2020, time.December, 28, 0, 0, 0, 0, time.UTC)),
	})

	vsCurrency := "USD"

	// getting balance for the address
	// May need to query from graphql if we want to get exact balance
	// Like what X does?
	// For now just using the price
	balances, err := balancesheet.ParseBalanceSheet(coin, vsCurrency, CG)
	if err != nil {
		return err
	}
	
	totalBalance,err :=balancesheet.TotalValueBalanceSheet([]types.Coin{
		coin,eth,
	},vsCurrency,CG)

	ethBalance,err := balancesheet.ParseBalanceSheet(eth, vsCurrency, CG)
	if err!=nil{
		return err
	}

	// Output the .csv file contains the
	// schema
	// date, coin price, account balance
	outputcsv := balances.GetCSV()
	fmt.Println(outputcsv)
	err = ioutil.WriteFile(fmt.Sprintf("%s.csv",balances[0].Coin), []byte(outputcsv), 0777)
    if err != nil {
        return err
    }

	totalCsv:=totalBalance.GetCSV()
	fmt.Println(totalCsv)
	err= ioutil.WriteFile("totalValue.csv", []byte(totalCsv), 0777)
	if err!=nil{
		return err
	}

	err=OutputCsv(ethBalance)
	if err!=nil{
		return err
	} */
	return nil
}

func OutputCsv(b types.Balances)error{
	totalCsv:=b.GetCSV()
	fmt.Println(totalCsv)
	return ioutil.WriteFile(fmt.Sprintf("%s.csv",b[0].Coin), []byte(totalCsv), 0777)
}