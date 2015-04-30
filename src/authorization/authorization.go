package main

import ( "gateway"
         "fmt"
         "os"
)

func main () {
    url := "https://gateway-sb.clearent.net/rest/v2/transactions"
    apikey := "762ad38136b54570a03be604a5a5a1e5"
    transaction := []byte(`<transaction><type>AUTH</type> <amount>200.00</amount> <card>4111111111111111</card> <exp-date>1020</exp-date> </transaction>`)

    resp, err  := gateway.PostTrans(url, apikey, transaction)
    fmt.Errorf("error: %g", err)
    os.Stdout.Write(resp)
}
