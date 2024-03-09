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
}

```