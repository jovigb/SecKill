package setup

import (
	service "SecKill/sk_proxy/config"
	"context"
	"encoding/json"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"log"
	"testing"
	"time"
)

func TestInitEtcd(t *testing.T) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Printf("Connect etcd failed. Error : %v", err)
	}

	nTime := time.Now()
	st := nTime.AddDate(0, 0, -3).Unix()
	et := nTime.AddDate(0, 0, 3).Unix()
	var SecInfoConfArr []service.SecProductInfoConf
	SecInfoConfArr = append(
		SecInfoConfArr,
		service.SecProductInfoConf{
			ProductId:         1,
			StartTime:         st,
			EndTime:           et,
			Status:            0,
			Total:             20,
			Left:              20,
			SoldMaxLimit:      20,
			OnePersonBuyLimit: 2,
			BuyRate:           0.2,
		},
		service.SecProductInfoConf{
			ProductId:         2,
			StartTime:         0,
			EndTime:           0,
			Status:            0,
			Total:             2000,
			Left:              1000,
			SoldMaxLimit:      1,
			OnePersonBuyLimit: 1,
			BuyRate:           0.2,
		},
	)
	data, err := json.Marshal(SecInfoConfArr)
	if err != nil {
		log.Printf("Data Marshal. Error : %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	_, err = cli.Put(ctx, "sec_kill_product", string(data))
	if err != nil {
		log.Printf("Put failed. Error : %v", err)
	}
	cancel()

	ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	resp, err := cli.Get(ctx, "sec_kill_product")
	if err != nil {
		log.Printf("Get falied. Error : %v", err)
	}

	for _, ev := range resp.Kvs {
		fmt.Printf("%s : %s\n", ev.Key, ev.Value)
	}
}
