package cron

import (
	"log"

	"github.com/everFinance/goar"
	"github.com/everFinance/goar/types"
	"github.com/liteseed/argo/transaction"
	"github.com/liteseed/bungo/internal/database/schema"
)

func (c *Cron) postBundle() {
	log.Println("posting bundle")
	o, err := c.database.GetOrders(25)
	if err != nil {
		return
	}
	if len(*o) == 0 {
		log.Println("no dataitem to post")
		return
	}

	dataItems := []transaction.DataItem{}

	for _, order := range *o {
		rawDataItem, err := c.store.Get(order.StoreID)
		if err != nil {
			log.Println("failed to fetch:", order.StoreID)
			continue
		}
		dataItem, err := transaction.DecodeDataItem(rawDataItem)
		if err != nil {
			log.Println("failed to decode:", order.StoreID)
			continue
		}
		dataItems = append(dataItems, *dataItem)
		err = c.database.UpdateStatus(order.ID, schema.Sent)
		if err != nil {
			log.Println("failed to update status:", order.ID, err)
			continue
		}
	}

	bundle, err := transaction.NewBundle(&dataItems)
	if err != nil {
		log.Println("failed to bundle:", err)
		return
	}

	wallet, err := goar.NewWalletFromPath("./data/signer.json", "https://arweave.dev/")
	if err != nil {
		log.Println("failed to upload:", err)
		return
	}

	tx, err := wallet.SendData([]byte(bundle.RawData), []types.Tag{{Name: "Bundle-Format", Value: "binary"}, {Name: "Bundle-Version", Value: "2.0.0"}})
	if err != nil {
		log.Println("failed to upload:", err)
		return
	}
	log.Println(tx.ID)
}