package lua

const (
	Seckkill string = `
	-- 1.参数列表
	-- 1.1 优惠券id
	local voucherId = KEYS[1]
	-- 1.2 用户id
	local userId = KEYS[2]
	-- 1.3 订单号码
	local orderId = KEYS[3]
	
	-- 2.数据key
	local stockKey = 'seckill:stock:' .. voucherId
	local orderKey = 'seckill:order:' .. voucherId
	
	-- 3.脚本业务
	-- 3.1 判断库存是否充足 get stockKey
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
		
		return 0
	`
)
