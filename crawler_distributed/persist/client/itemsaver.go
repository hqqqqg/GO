// 客户端 通过RPC调用远程ItemSaver服务将数据存入Elasticsearch
package client

import (
	"google/crawler/engine"
	"google/crawler_distributed/config"
	"google/crawler_distributed/rpcsupport"
	"log"
)

func ItemSaver( //通过RPC将item发往远程存储
	host string) (chan engine.Item, error) {
	client, err := rpcsupport.NewClient(host) //RPC客户端连接到itemSaver服务
	if err != nil {
		return nil, err
	}
	out := make(chan engine.Item) //创建item通道
	go func() {
		itemCount := 0 //item计数
		for {
			item := <-out
			log.Printf("Item Saver:got item "+
				"#%d:%v", itemCount, item)
			itemCount++
			result := ""
			err := client.Call( //调用远程ItemSaverService.Save 方法
				config.ItemSaverRpc,
				item, &result)
			if err != nil {
				log.Printf("Item Saver:error"+
					"Saving item %v:%v",
					item, err)
			}
		}
	}()
	return out, nil //返回item通道给调用
}
