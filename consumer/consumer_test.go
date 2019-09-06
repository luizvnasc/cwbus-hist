package consumer

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"testing"
	"time"
)

func print(res *http.Response) error {
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	f, err := os.Create("out.json")
	defer f.Close()
	if err != nil {
		return err
	}
	f.WriteString(string(body))
	return nil
}

func TestConsumer(t *testing.T) {
	cherr := make(chan error)
	timeout := 5 * time.Second
	consumer := New("http://transporteservico.urbs.curitiba.pr.gov.br/getLinhas.php?c=a1150", print, cherr)
	go consumer.Run()
	select {
	case err := <-cherr:
		t.Errorf("expected print a response body, get error %v", err)
	case <-time.After(timeout):
		result, _ := os.Open("out.json")
		defer result.Close()
		var str string
		fmt.Fscan(result, &str)
		if len(str) == 0 {
			t.Errorf("No result")
		}
	}
}
