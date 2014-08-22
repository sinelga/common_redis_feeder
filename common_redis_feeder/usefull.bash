GOPATH=$GOPATH:/home/juno/git/common_redis_feeder/common_redis_feeder go test -v


ZRANGEBYSCORE en_US:finance:news:business -inf +inf withscores  LIMIT 0 2000

ZCOUNT en_US:finance:news:business -inf +inf

 bin/tabmenufeedup -locale=en_US -themes=finance
 bin/cron_parser -locale=en_US -themes=finance
