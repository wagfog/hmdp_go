package lua

const (
	Seckkill string = `
		-- 1.参数列表
		-- 1.1 优惠券id
		local voucherId = ARGV[1]
		-- 1.2 用户id
		local userID = ARGV[2]
		-- 1.3 订单号码
		local orderId = ARGV[3]
		
		-- 2.数据key
		local stockKey = 'seckill:stock:' .. voucherId
		
		-- 3.脚本业务
		-- 3.1 判断库存是否充足 get stockKey
		local orderKey = 'seckill:order:' .. voucherId
		if (tonumber(redis.call('get',stockKey)) <= 0) then
			-- 库存不足
			return 1
		end
		-- 判断用户是否下单
		if (redis.call('sismember',orderKey,userId) == 1) then
			-- 重复下单
			return 2
		end
		
		-- 扣减库存
		redis.call('incrby',stockKey,-1);
		-- 下单
		redis.call('sadd',orderKey,userId);

		-- 发送消息到队列中，XADD stream.order * k1 v1 k2 v2
		-- redis.call('xadd', 'stream.orders', '*', 'userId',userId,'voucherId',voucherId,'id',orderId)
		
		return 0
	`
)
