// Copyright 2018 ProximaX Limited. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package sdk

import (
	"errors"
	"testing"
)

const (
	mijinRoute = `{
  			"name": "mijin",
  			"description": "catapult development network"
  	}`
	mijinTestRoute = `{
  			"name": "mijinTest",
  			"description": "catapult development network"
  	}`
)

func TestNetworkService_GetNetworkType(t *testing.T) {

	serv := NewMockServerWithRouters(map[string]sRouting{pathNetwork: {resp: mijinTestRoute}})

	netType, resp, err := serv.Network.GetNetworkType(ctx)
	if err != nil {
		t.Error(err)
	} else if resp.StatusCode != 200 {
		t.Error(resp.Status)
		t.Logf("%#v", resp)
	} else if netType != MijinTest {
		t.Errorf("%d", netType)
	}

}
func TestNetworkService_GetNetworkType_MIJIN(t *testing.T) {

	serv := NewMockServerWithRouters(map[string]sRouting{pathNetwork: {resp: mijinRoute}})

	netType, resp, err := serv.Network.GetNetworkType(ctx)
	if err != nil {
		t.Error(err)
	} else if resp.StatusCode != 200 {
		t.Error(resp.Status)
		t.Logf("%#v", resp)
	} else if netType != Mijin {
		t.Errorf("%d", netType)
	}

}
func TestNetworkService_GetNetworkType_Unknow(t *testing.T) {

	serv := NewMockServerWithRouters(map[string]sRouting{pathNetwork: {
		resp: `{
 				 "name": "",
  				"description": "catapult development network"
  				}`}})

	netType, resp, err := serv.Network.GetNetworkType(ctx)
	if err == nil {
		t.Error(errors.New("Must be errror"))
	} else if resp.StatusCode != 200 {
		t.Error(resp.Status)
		t.Logf("%#v", resp)
	} else if netType != NotSupportedNet {
		t.Errorf("%d", netType)
	}

}
func TestExtractNetworkType(t *testing.T) {
	i := uint64(36888)

	nt := ExtractNetworkType(i)
	if nt != MijinTest {
		t.Errorf("wrong convert %d (%d - %d)", i, nt, MijinTest)
	}

}
