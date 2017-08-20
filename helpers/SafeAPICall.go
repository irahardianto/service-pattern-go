package helpers

import (
	"io/ioutil"
	"net/http"

	"github.com/afex/hystrix-go/hystrix"
)

type SafeAPICall struct{}

func (helper *SafeAPICall) Get(commandName string, endpoint string, timeout int) []byte {

	responseChannel := make(chan []byte)
	hystrix.ConfigureCommand(commandName, hystrix.CommandConfig{Timeout: timeout})
	hystrix.Go(commandName, func() error {

		resp, err := http.Get(endpoint)
		if err != nil {
			// handle error
		}
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)

		responseChannel <- body
		return nil
	}, func(err error) error {
		//fmt.Println("a")
		var emptyByteArray []byte
		responseChannel <- emptyByteArray
		return nil
	})

	return <-responseChannel
}
