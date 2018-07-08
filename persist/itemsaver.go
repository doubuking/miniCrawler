package persist

import (
	"log"
	"gopkg.in/olivere/elastic.v5"
	"context"
	"miniCrawler/engine"
	"errors"
)

func ItemSaver(index string)  (chan engine.Item,error) {
	out := make(chan engine.Item)
	client,err := elastic.NewClient(
		//Must turn off sniff in docker
		elastic.SetSniff(false))
	if err != nil{
		return nil,err
	}

	go func() {
		//fd,err := os.OpenFile("baletu.log",os.O_RDWR|os.O_CREATE|os.O_APPEND,0644)
		//if err != nil {
		//	panic(err)
		//}
		//defer fd.Close()
		itemCount := 0
		for  {
			item := <-out
			log.Printf("Item saver:got item #%d : %v",itemCount,item)
			itemCount++

			err := save(client,item,index)
			if err != nil{
				log.Print("Item Save error saving item %v : %v",item,err)

			}


			//写入文件
			//b,err := json.Marshal(item)
			//if err != nil{
			//	panic(err)
			//}
			//fd.Write(b)
			//fd.Write([]byte("\n"))
		}





	}()

	return out,nil
}

func save(client *elastic.Client,item engine.Item,index string) error {


	if item.Type == ""{
		return errors.New("Must supply Type")
	}

	indexService := client.Index().
		Index(index).
		Type(item.Type).
		BodyJson(item)

	if item.Id == ""{
		indexService = indexService.Id(item.Id)
	}

	_,err :=  indexService.
		Do(context.Background())

	if err != nil{
		return err
	}
	return nil
}
