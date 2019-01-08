/*
 * Tencent is pleased to support the open source community by making 蓝鲸 available.
 * Copyright (C) 2017-2018 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except 
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and 
 * limitations under the License.
 */
 
package actions_test

import (
	"bytes"
	"configcenter/src/common"
	"encoding/json"
	"github.com/stretchr/testify/require"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

func testDelete(t *testing.T, ids ...string) {
	server := CCAPITester(t)
	originData := map[string]interface{}{
		common.BKHostIDField:  strings.Join(ids, ","),
		common.BKOwnerIDField: "0",
	}
	reqbody, err := json.Marshal(originData)
	r, err := http.NewRequest("DELETE", server.URL+"/host/v1/host/batch", bytes.NewBuffer(reqbody))
	require.NoError(t, err)
	resp, err := http.DefaultClient.Do(r)
	require.NoError(t, err)
	respbody, err := ioutil.ReadAll(resp.Body)
	require.NoError(t, err)
	body := gjson.ParseBytes(respbody)
	require.Equal(t, "true", body.Get("result").String(), "response:[%s]", respbody)
}
