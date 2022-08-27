package server

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"

	orderController "orders.api/pkg/controller/order"
	orderRepository "orders.api/pkg/repository/order"
	orderService "orders.api/pkg/service/order"
)

func resolveOrderController() orderController.Controller {
	c, err := orderController.NewController(resolveOrderService())
	if err != nil {
		log.Panicf("error handled while creating instances: %v", err)
	}
	return *c
}

func resolveOrderService() orderService.Service {
	s, err := orderService.NewService(resolveOrderRepository(), nil, nil, nil)
	if err != nil {
		log.Panicf("error handled while creating instances: %v", err)
	}
	return *s
}

func resolveOrderRepository() orderRepository.Repository {
	r, err := orderRepository.NewRepository(resolveMySQLClient())
	if err != nil {
		log.Panicf("error handled while creating instances: %v", err)
	}
	return *r
}

func resolveMySQLClient() *sql.DB {
	conf := mysql.NewConfig()
	conf.ParseTime = true
	conf.Net = "tcp"
	conf.Collation = "utf8_unicode_ci"
	conf.User = "root"
	conf.Passwd = "admin"
	conf.DBName = "order_api"
	conf.Addr = "127.0.0.1"
	db, err := sql.Open("mysql", conf.FormatDSN())
	if err != nil {
		log.Panicf("error handled while creating instances: %v", err)
	}
	return db
}
