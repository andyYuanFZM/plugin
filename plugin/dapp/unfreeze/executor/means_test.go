package executor

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/33cn/chain33/types"
	pty "github.com/33cn/plugin/plugin/dapp/unfreeze/types"
	"fmt"
	"time"
)

func TestCalcFrozen(t *testing.T) {
	types.SetTitleOnlyForTest("chain33")
	m, err := newMeans("LeftProportion", 15000000)
	assert.Nil(t, err)
	assert.NotNil(t, m)

	cases := []struct {
		start         int64
		now           int64
		period        int64
		total         int64
		tenThousandth int64
		expect        int64
	}{
		{10000, 10001, 10, 10000, 2, 9998},
		{10000, 10011, 10, 10000, 2, 9996},
		{10000, 10001, 10, 1e17, 2, 9998 * 1e13},
		{10000, 10011, 10, 1e17, 2, 9998 * 9998 * 1e9},
	}

	for _, c := range cases {
		c := c
		t.Run("test LeftProportion", func(t *testing.T) {
			create := pty.UnfreezeCreate{
				StartTime:   c.start,
				AssetExec:   "coins",
				AssetSymbol: "bty",
				TotalCount:  c.total,
				Beneficiary: "x",
				Means:       "LeftProportion",
				MeansOpt: &pty.UnfreezeCreate_LeftProportion{
					LeftProportion: &pty.LeftProportion{
						Period:        c.period,
						TenThousandth: c.tenThousandth,
					},
				},
			}
			u := &pty.Unfreeze{
				TotalCount: c.total,
				Means:      "LeftProportion",
				StartTime:  c.start,
				MeansOpt: &pty.Unfreeze_LeftProportion{
					LeftProportion: &pty.LeftProportion{
						Period:        c.period,
						TenThousandth: c.tenThousandth,
					},
				},
			}
			u, err = m.setOpt(u, &create)
			assert.Nil(t, err)

			f, err := m.calcFrozen(u, c.now)
			assert.Nil(t, err)

			assert.Equal(t, c.expect, f)

		})
	}
}

func TestLeftV1(t *testing.T) {
	cases := []struct {
		start         int64
		now           int64
		period        int64
		total         int64
		tenThousandth int64
		expect        int64
	}{
		{10000, 10001, 10, 10000, 2, 9998},
		{10000, 10011, 10, 10000, 2, 9996},
		{10000, 10001, 10, 1e17, 2, 9998 * 1e13},
		{10000, 10011, 10, 1e17, 2, 9998 * 9998 * 1e9},
	}

	for _, c := range cases {
		c := c
		t.Run("test LeftProportionV1", func(t *testing.T) {
			create := pty.UnfreezeCreate{
				StartTime:   c.start,
				AssetExec:   "coins",
				AssetSymbol: "bty",
				TotalCount:  c.total,
				Beneficiary: "x",
				Means:       pty.LeftProportionX,
				MeansOpt: &pty.UnfreezeCreate_LeftProportion{
					LeftProportion: &pty.LeftProportion{
						Period:        c.period,
						TenThousandth: c.tenThousandth,
					},
				},
			}
			u := &pty.Unfreeze{
				TotalCount: c.total,
				Means:      pty.LeftProportionX,
				StartTime:  c.start,
				MeansOpt: &pty.Unfreeze_LeftProportion{
					LeftProportion: &pty.LeftProportion{
						Period:        c.period,
						TenThousandth: c.tenThousandth,
					},
				},
			}
			m := leftProportion{}
			u, err := m.setOpt(u, &create)
			assert.Nil(t, err)

			f, err := m.calcFrozen(u, c.now)
			assert.Nil(t, err)

			assert.Equal(t, c.expect, f)

		})
	}
}

func TestFixV1(t *testing.T) {
	cases := []struct {
		start  int64
		now    int64
		period int64
		total  int64
		amount int64
		expect int64
	}{
		{10000, 10001, 10, 10000, 2, 9998},
		{10000, 10011, 10, 10000, 2, 9996},
		{10000, 10001, 10, 1e17, 2, 1e17 - 2},
		{10000, 10011, 10, 1e17, 2, 1e17 - 4},
	}

	for _, c := range cases {
		c := c
		t.Run("test FixAmountV1", func(t *testing.T) {
			create := pty.UnfreezeCreate{
				StartTime:   c.start,
				AssetExec:   "coins",
				AssetSymbol: "bty",
				TotalCount:  c.total,
				Beneficiary: "x",
				Means:       pty.FixAmountX,
				MeansOpt: &pty.UnfreezeCreate_FixAmount{
					FixAmount: &pty.FixAmount{
						Period: c.period,
						Amount: c.amount,
					},
				},
			}
			u := &pty.Unfreeze{
				TotalCount: c.total,
				Means:      pty.FixAmountX,
				StartTime:  c.start,
				MeansOpt: &pty.Unfreeze_FixAmount{
					FixAmount: &pty.FixAmount{
						Period: c.period,
						Amount: c.amount,
					},
				},
			}
			m := fixAmount{}
			u, err := m.setOpt(u, &create)
			assert.Nil(t, err)

			f, err := m.calcFrozen(u, c.now)
			assert.Nil(t, err)

			assert.Equal(t, c.expect, f)

		})
	}
}

// 查询可提币量， 和当前时间有关， 如对应节点时间不对， 查询结果也不对
func TestLeftV2(t *testing.T) {
	cases := []struct {
		start         int64
		now           int64
		period        int64
		total         int64
		tenThousandth int64
		expect        int64
	}{
		{1561607389, 1561607389 + 500000, 67200, 11111130, 1, 11102244},
		{1561607389, -156107389 + 500000, 67200, 11111130, 1, 11111130},
	}

	for _, c := range cases {
		c := c
		t.Run("test LeftProportionV2", func(t *testing.T) {
			create := pty.UnfreezeCreate{
				StartTime:   c.start,
				AssetExec:   "coins",
				AssetSymbol: "bty",
				TotalCount:  c.total,
				Beneficiary: "x",
				Means:       pty.LeftProportionX,
				MeansOpt: &pty.UnfreezeCreate_LeftProportion{
					LeftProportion: &pty.LeftProportion{
						Period:        c.period,
						TenThousandth: c.tenThousandth,
					},
				},
			}
			u := &pty.Unfreeze{
				TotalCount: c.total,
				Means:      pty.LeftProportionX,
				StartTime:  c.start,
				MeansOpt: &pty.Unfreeze_LeftProportion{
					LeftProportion: &pty.LeftProportion{
						Period:        c.period,
						TenThousandth: c.tenThousandth,
					},
				},
			}
			m := leftProportion{}
			u, err := m.setOpt(u, &create)
			assert.Nil(t, err)

			f, err := m.calcFrozen(u, c.now)
			assert.Nil(t, err)

			assert.Equal(t, c.expect, f)

		})
	}
}

func TestDecreaseAmount(t *testing.T) {
	amount := int64(400000)
	rate := int64(1000)

	start := int64(1546272000)


	t.Run("test TestDecreaseAmount", func(t *testing.T) {
		create := pty.UnfreezeCreate{
			StartTime:   start,
			AssetExec:   "coins",
			AssetSymbol: "bty",
			TotalCount:  int64(800000000),
			Beneficiary: "x",
			Means:       pty.DecreaseAmountX,
			MeansOpt: &pty.UnfreezeCreate_DecreaseAmount{
				DecreaseAmount: &pty.DecreaseAmount{
					Period:        int64(86400),
					TenThousandth: int64(1000),
					FristDecreaseAmount:int64(400000),
					DecreasePeriod:int64(200*86400),
					DecreaseNums:int64(34),
				},
			},
		}
		u := &pty.Unfreeze{
			TotalCount: int64(800000000),
			Means:      pty.DecreaseAmountX,
			StartTime:  start,
			MeansOpt: &pty.Unfreeze_DecreaseAmount{
				DecreaseAmount: &pty.DecreaseAmount{
					Period:        int64(86400),
					TenThousandth: int64(1000),
					FristDecreaseAmount:int64(400000),
					DecreasePeriod:int64(200*86400),
					DecreaseNums:int64(34),
				},
			},
		}
		m := decreaseAmount{}
		u, err := m.setOpt(u, &create)
		if err != nil {
			panic(err)
		}
		days := int64(0)
		//assert.Nil(t, err)
		for i:= int64(1);i <= int64(36); i ++ {
			tt := getDecreasePeriodAmount(int64(i-1),amount,rate)

			fmt.Println("amount:",tt)
			for n:= int64(1);n <= 200;n++ {
				days ++
				f, err := m.calcFrozen(u, start + (days-1)*86400)
				//assert.Nil(t, err)
				if err != nil {
					panic(err)
				}
				fmt.Println("递减次数:",i-1,"解锁序列","f:",f,time.Unix((start + (days-1)*86400),0).Format("2006-01-02"))
			}
		}

		//assert.Equal(t, c.expect, f)

	})
}