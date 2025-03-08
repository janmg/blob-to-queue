package output

// https://github.com/fluent/fluent-logger-golang
import (
	"fmt"

	"github.com/fluent/fluent-logger-golang/fluent"
	"janmg.com/blob-to-queue/format"
)

func SendFluent(nsg format.Flatevent) {
	logger, err := fluent.New(fluent.Config{})
	if err != nil {
		fmt.Println(err)
	}
	defer logger.Close()

	tag := "myapp.access"
	error := logger.Post(tag, nsg)
	// error := logger.PostWithTime(tag, time.Now(), data)
	if error != nil {
		panic(error)
	}
}
