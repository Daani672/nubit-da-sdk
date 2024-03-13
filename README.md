# nubit-da-sdk

### 用法
```go
import "github.com/nubit/nubit-da-sdk"

func main() {
	ctx:=context.Background()
	invoice:="invoicexxxxxx"
	fee:=50
	client := nubit.NewClient(&types.PaymentParams{
        XAPIKEY:"xxx"
    })
	status:=client.Payment(ctx,invoice,fee)
	fmt.Println(status)

    client2 := nubit.NewClient(&types.PaymentParams{
        Target:"127.0.0.1：:10009"
		MacaroonFile:"/path/to/macaroon"
    })
    status2:=client.Payment(ctx,invoice,fee)
    fmt.Println(status2)
}

```