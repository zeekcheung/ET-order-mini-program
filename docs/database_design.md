# 1. 实体

- 系统使用小程序官方的接口进行登录，可以直接获取到用户信息，因此数据库中不需要再额外存储用户实体的数据。
- 在微信小程序官方所提供的接口中，用户唯一标识为 `openid`
- 其他实体的数据结构参考接口文档

# 2. 关系

## 商品关系

- 商品分类-商品（多对多关系）
  - 一个商品分类可以有多个商品；一个商品可以属于多个商品分类。
- 店铺-商品（一对多关系）
  - 一个店铺拥有多个商品；一个商品只属于一个店铺。
- 商品-商品评论（一对多关系）
  - 一个商品可以有多个商品评论；一个商品评论只属于一个商品。
- 商品-商品规格（一对多关系）
  - 一个商品可以有多个商品规格；一个商品规格只属于一个商品。

## 购物车关系

- 用户-购物车（一对一关系）
  - 一个用户只能有一个购物车；一个购物车只属于一个用户。
- 购物车-购物车列表项（一对多关系）
  - 一个购物车可以有多个购物车列表项；一个购物车列表项只属于一个购物车。
- 购物车列表项-商品（多对一关系）
  - 一个购物车列表项对应一个商品；一个商品可以对应多个购物车列表项。
- 购物车列表项-商品规格（多对一关系）
  - 一个购物车列表项对应一个商品规格；一个商品规格可以对应多个购物车列表项。

## 订单关系

- 用户-订单（一对多关系）
  - 一个用户可以下多个订单；一个订单只属于一个用户。
- 订单-订单列表项（一对多关系）
  - 一个订单可以有多个订单列表项；一个订单列表项只属于一个订单。
- 订单列表项-商品（多对一关系）
  - 一个订单列表项对应一个商品；一个商品可以对应多个订单列表项。
- 订单列表项-商品规格（多对一关系）
  - 一个订单列表项对应一个商品规格；一个商品规格可以对应多个订单列表项。
- 订单-收货地址（多对一关系）
  - 一个订单拥有一个收货地址；一个顾客地址属于多个订单。
- 用户-收货地址（一对多关系）
  - 一个用户可以有多个收货地址；一个收货地址只属于一个用户。