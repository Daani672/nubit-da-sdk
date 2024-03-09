# nubit-da-sdk

### 用法
```go
import "github.com/nubit/nubit-da-sdk"

func main() {
	ctx:=context.Background()
	invoice:="invoicexxxxxx"
	fee:=50
	ctx = metadata.AppendToOutgoingContext(ctx, "address", "1JSAbzibK4PLQp4U876o6wmCPV39Ry3wdp")
	client := nubit.NewClient(&types.PaymentParams{
        XAPIKEY:"xxx"
    })
	status:=client.Payment(ctx,invoice,fee)
	fmt.Println(status)
}

```