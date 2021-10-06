package main

import (
	"context"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"time"
)

func main() {
	cli ,err := clientv3.New(clientv3.Config{
		Endpoints: []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err !=nil{
		fmt.Printf("connect to etcd faild,err:%v\n,err")
		return
	}
	fmt.Printf("connect to etcd success!")
	defer cli.Close()
	//etcd put
	ctx ,cancel:=context.WithTimeout(context.Background(),time.Second)
	_ , err =cli.Put(ctx,"donglin","he is studying etcd")
	if err !=nil{
	fmt.Printf("put to etcd failed,err %v\n",err)
	return
	}
	cancel()
	//etcd get
	ctx ,cancel =context.WithTimeout(context.Background(),time.Second)
	resp, err := cli.Get(ctx,"donglin")
	cancel()
	if err !=nil{
		fmt.Printf("get to etcd failed,err %v\n",err)
		return
	}
	for _, ev :=range resp.Kvs{
		fmt.Printf("%s:%s\n",ev.Key,ev.Value)
	}

}
