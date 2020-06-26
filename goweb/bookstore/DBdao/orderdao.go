package DBdao

import (
	"fmt"
	"goweb/bookstore/model"
	"goweb/bookstore/utils"
)

func AddOrder(order *model.Order) error {
	sql := "insert into  orders (id ,  create_time,total_count,total_amount ,state , user_id ) values ($1,$2,$3,$4,$5,$6);"

	_, err := utils.Db.Exec(sql, order.ID, order.CreateTime, order.TotalCount, order.TotalAmount, order.State, order.Uer_id)
	if err != nil {
		fmt.Println(err)
	}
	return nil
}

func GetOrder(orderId string) ([]*model.Order, error) {

	sql := "select id , create_time,total_count,total_amount ,state , user_id from orders where id = $1;"

	rows, err := utils.Db.Query(sql, orderId)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var orderInfo []*model.Order

	for rows.Next() {
		var order model.Order
		err = rows.Scan(&order.ID, &order.CreateTime, &order.TotalCount, &order.TotalAmount, &order.State, &order.Uer_id)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		orderInfo = append(orderInfo, &order)
	}
	return orderInfo, nil
}

func UpdateOrderState(orderState int64, orderId string) error {
	sql := "update orders set state = $1 where ID = $2;"
	_, err := utils.Db.Exec(sql, orderState, orderId)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
