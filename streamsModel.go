package redisClient

import "github.com/redis/go-redis/v9"

type XAddArgs = redis.XAddArgs

type XReadArgs = redis.XReadArgs

type XReadGroupArgs = redis.XReadGroupArgs

type XClaimArgs = redis.XClaimArgs

type XAutoClaimArgs = redis.XAutoClaimArgs

type XPendingExtArgs = redis.XPendingExtArgs

type XPending = redis.XPending

type XPendingExt = redis.XPendingExt

type XMessage = redis.XMessage

type XStream = redis.XStream

type XInfoStream = redis.XInfoStream

type XInfoStreamFull = redis.XInfoStreamFull

type XInfoGroup = redis.XInfoGroup

type XInfoConsumer = redis.XInfoConsumer
