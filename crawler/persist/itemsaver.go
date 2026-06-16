package persist

import "log"

func ItemSaver() chan interface{} {
	out := make(chan interface{})
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Item Saver:got item "+
				"#%d:%v", itemCount, item) //要存的内容先打印
			itemCount++
		}
	}()
	return out
}
