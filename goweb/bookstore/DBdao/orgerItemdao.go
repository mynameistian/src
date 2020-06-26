package DBdao

import (
	"fmt"
	"goweb/bookstore/model"
	"goweb/bookstore/utils"
)

func AddOrderItem(orderItem *model.OrderItem) error {
	sql := "insert into order_items (count , amount, title, author,price,img_path,order_id) values($1,$2,$3,$4,$5,$6,$7);"
	_, err := utils.Db.Exec(sql, orderItem.Count, orderItem.Amount, orderItem.Title, orderItem.Author, orderItem.Price, orderItem.ImgPath, orderItem.OrderId)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return err
}

func GetOrderItem(orderId string) ([]*model.OrderItem, error) {
	sql := "select id , count , amount, title, author,price,img_path,order_id from order_items where order_id = $1;"
	rows, err := utils.Db.Query(sql, orderId)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var orderItems []*model.OrderItem

	for rows.Next() {
		var orderItem model.OrderItem
		err := rows.Scan(&orderItem.ID, &orderItem.Count, &orderItem.Amount, &orderItem.Title, &orderItem.Author, &orderItem.Price, &orderItem.ImgPath, &orderItem.OrderId)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		orderItems = append(orderItems, &orderItem)
	}
	return orderItems, err
}
