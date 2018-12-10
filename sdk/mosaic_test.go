// Copyright 2018 ProximaX Limited. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package sdk

import (
	"fmt"
	"github.com/proximax-storage/proximax-utils-go/mock"
	"github.com/proximax-storage/proximax-utils-go/tests"
	"github.com/stretchr/testify/assert"
	"math/big"
	"net/http"
	"testing"
)

func init() {
	i, _ := (&big.Int{}).SetString("15358872602548358953", 10)
	testMosaicId.Id = i
}

var (
	mosaicClient  = mockServer.getTestNetClientUnsafe().Mosaic
	testMosaicId  = &MosaicId{Id: uint64DTO{3646934825, 3576016193}.toBigInt()}
	testMosaicIds = MosaicIds{MosaicIds: []*MosaicId{
		testMosaicId,
		{Id: big.NewInt(5734678065854194365)},
	}}
)

const (
	testMosaicPathID         = "d525ad41d95fcf29"
	testMosaicNamespaceEmpty = "a887d82dfeb659b0"
	testMosaicFromNamesaceId = "5B55E02EACCB7B00015DB6EC"
)

var (
	tplMosaic = `
{
  "meta": {
    "active": true,
    "index": 0,
    "id": "5B55E02EACCB7B00015DB6EC"
  },
  "mosaic": {
    "namespaceId": [
      929036875,
      2226345261
    ],
    "mosaicId": [
      3646934825,
      3576016193
    ],
    "supply": [
      3403414400,
      2095475
    ],
    "height": [
      1,
      0
    ],
    "owner": "321DE652C4D3362FC2DDF7800F6582F4A10CFEA134B81F8AB6E4BE78BBA4D18E",
    "properties": [
      [
        2,
        0
      ],
      [
        6,
        0
      ],
      [
        0,
        0
      ]
    ],
    "levy": {}
  }
}`

	mosaicCorr = &MosaicInfo{
		MosaicId: &MosaicId{
			Id: uint64DTO{3646934825, 3576016193}.toBigInt(),
		},
		MetaId: "5B55E02EACCB7B00015DB6EC",
		NamespaceId: &NamespaceId{
			Id: uint64DTO{929036875, 2226345261}.toBigInt(),
		},
		Supply: uint64DTO{3403414400, 2095475}.toBigInt(),
		Active: true,
		Height: big.NewInt(1),
		Owner: &PublicAccount{
			Address: &Address{
				Type:    mosaicClient.client.config.NetworkType,
				Address: "TBFBW6TUGLEWQIBCMTBMXXQORZKUP3WTVVPAYGJN",
			},

			PublicKey: "321DE652C4D3362FC2DDF7800F6582F4A10CFEA134B81F8AB6E4BE78BBA4D18E",
		},
		Properties: &MosaicProperties{
			Transferable: true,
			Divisibility: 6,
			Duration:     big.NewInt(0),
		},
	}

	mosaicName = &MosaicName{
		MosaicId: &MosaicId{
			Id: uint64DTO{3646934825, 3576016193}.toBigInt(),
		},
		Name: "xem",
		ParentId: &NamespaceId{
			Id: uint64DTO{929036875, 2226345261}.toBigInt(),
		},
	}
)

func TestMosaicService_GetMosaic(t *testing.T) {
	mockServer.AddRouter(&mock.Router{
		Path:     pathMosaic + testMosaicPathID,
		RespBody: tplMosaic,
	})

	mscInfo, resp, err := mosaicClient.GetMosaic(ctx, mosaicCorr.MosaicId)

	assert.Nilf(t, err, "MosaicService.GetMosaic returned error: %s", err)

	if tests.IsOkResponse(t, resp) {
		tests.ValidateStringers(t, mosaicCorr, mscInfo)
	}
}

func TestMosaicService_GetMosaics(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockServer.AddRouter(&mock.Router{
			Path:     pathMosaic,
			RespBody: "[" + tplMosaic + "]",
			ReqJsonBodyStruct: struct {
				MosaicIds []string `json:"mosaicIds"`
			}{},
		})

		mscInfoArr, resp, err := mosaicClient.GetMosaics(ctx, MosaicIds{[]*MosaicId{mosaicCorr.MosaicId}})

		assert.Nilf(t, err, "MosaicService.GetMosaics returned error: %s", err)

		if tests.IsOkResponse(t, resp) {
			for _, mscInfo := range mscInfoArr {
				tests.ValidateStringers(t, mosaicCorr, mscInfo)
			}
		}
	})

	t.Run("empty url params", func(t *testing.T) {
		_, resp, err := mosaicClient.GetMosaics(ctx, MosaicIds{})

		assert.NotNil(t, err, "MosaicService.GetMosaics returned error: %s", err)

		if resp != nil {
			assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
		}
	})
}

func TestMosaicService_GetMosaicNames(t *testing.T) {
	mockServer.AddRouter(&mock.Router{
		Path: pathMosaicNames,
		RespBody: `[
						  {
							"mosaicId": [
							  3646934825,
							  3576016193
							],
							"name": "xem",
							"parentId": [
							  929036875,
							  2226345261
							]
						  }
						]`,
		ReqJsonBodyStruct: struct {
			MosaicIds []string `json:"mosaicIds"`
		}{},
	})

	mscInfoArr, resp, err := mosaicClient.GetMosaicNames(ctx, testMosaicIds)

	assert.Nil(t, err, "MosaicService.GetMosaicNames returned error: %s", err)

	if tests.IsOkResponse(t, resp) {
		for _, mscInfo := range mscInfoArr {
			tests.ValidateStringers(t, mosaicName, mscInfo)
		}
	}
}

func TestMosaicService_GetMosaicsFromNamespace(t *testing.T) {
	t.Run("regular case", func(t *testing.T) {
		mockServer.AddRouter(&mock.Router{
			Path:     fmt.Sprintf(pathMosaicFromNamespace, mosaicNamespace),
			RespBody: "[" + tplMosaic + "]",
		})

		mscInfoArr, resp, err := mosaicClient.GetMosaicsFromNamespace(ctx, testNamespaceId, testMosaicId, pageSize)

		assert.Nil(t, err, "MosaicService.GetMosaicsFromNamespace returned error: %s", err)

		if tests.IsOkResponse(t, resp) {
			for _, mscInfo := range mscInfoArr {
				tests.ValidateStringers(t, mosaicCorr, mscInfo)
			}
		}
	})

	t.Run("no mosaic id", func(t *testing.T) {
		mockServer.AddRouter(&mock.Router{
			Path:     fmt.Sprintf(pathMosaicFromNamespace, testMosaicNamespaceEmpty),
			RespBody: "[]",
		})

		nsId, _ := (&big.Int{}).SetString("12143912612286323120", 10)

		mscInfoArr, resp, err := mosaicClient.GetMosaicsFromNamespace(ctx, &NamespaceId{Id: nsId}, nil, pageSize)

		assert.Nil(t, err, "MosaicService.GetMosaicsFromNamespace returned error: %s", err)

		if tests.IsOkResponse(t, resp) {
			assert.Equal(t, len(mscInfoArr), 0)
		}
	})
}
