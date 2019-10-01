package simulation

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/FactomProject/factomd/common/primitives"
	"io/ioutil"
	"net/http"
)

func V2Request(req *primitives.JSON2Request, port int) (*primitives.JSON2Response, error) {
j, err := json.Marshal(req)
if err != nil {
return nil, err
}

portStr := fmt.Sprintf("%d", port)
resp, err := http.Post(
"http://localhost:"+portStr+"/v2",
"application/json",
bytes.NewBuffer(j))
if err != nil {
return nil, err
}
defer resp.Body.Close()

body, err := ioutil.ReadAll(resp.Body)
if err != nil {
return nil, err
}
r := primitives.NewJSON2Response()
if err := json.Unmarshal(body, r); err != nil {
return nil, err
}
return nil, nil
}

