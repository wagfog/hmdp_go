package utils

const (
	//Login
	LOGIN_CODE_KEY string = "login:code:"
	LOGIN_CODE_TTL int64  = 2
	LOGIN_USER_KEY string = "login:token:"
	LOGIN_USER_TTL int64  = 36000

	CACHE_NULL_TTL int64 = 2

	CACHE_SHOP_TTL int64  = 30
	CACHE_SHOP_KEY string = "cache:shop:"

	SHOP_TYPE_KEY string = "shop:type"

	LOCK_SHOP_KEY string = "lock:shop:"
	LOCK_SHOP_TTL int64  = 10

	SECKILL_STOCK_KEY string = "seckill:stock:"
	BLOG_LIKED_KEY    string = "blog:liked:"
	FEED_KEY          string = "feed:"
	SHOP_GEO_KEY      string = "shop:geo:"
	USER_SIGN_KEY     string = "sign:"
)
