package queue

import (
    "github.com/redis/go-redis/v9"
    "leoferaderonugraha/go-backend-boilerplate/pkg/config"
    "context"
)

type RedisQueue struct {
    Name string
    client *redis.Client
    context context.Context
}

// TODO: implement cache and sync
func NewRedisQueue(name string) (Queue[string, []string], error) {
    url := config.Get[string]("redis_host", "localhost:6379")
    password := config.Get[string]("redis_password", "")
    db := config.Get[int]("redis_db", 0)

    options := &redis.Options{
        Addr: url,
        Password: password,
        DB: db,
    }
    client := redis.NewClient(options)
    ctx := context.Background()

    _, err := client.Ping(ctx).Result()

    if err != nil {
        return nil, err
    }

    return &RedisQueue{
        Name: name,
        client: client,
        context: ctx,
    }, nil
}

func (q *RedisQueue) Pop() (string, error) {
    result := q.client.LPop(q.context, q.Name)

    if result.Err() != nil {
        return "", result.Err()
    }

    return result.Val(), nil
}

func (q *RedisQueue) Push(item string) error {
    result := q.client.RPush(q.context, q.Name, item)

    return result.Err()
}

func (q *RedisQueue) PeekFront() (string, error) {
    result := q.client.LIndex(q.context, q.Name, 0)

    if result.Err() != nil {
        return "", result.Err()
    }

    return result.Val(), nil
}

func (q *RedisQueue) PeekBack() (string, error) {
    result := q.client.LIndex(q.context, q.Name, -1)

    if result.Err() != nil {
        return "", result.Err()
    }

    return result.Val(), nil
}

func (q *RedisQueue) Clear() error {
    result := q.client.Del(q.context, q.Name)

    return result.Err()
}

func (q *RedisQueue) Items() []string {
    result := q.client.LRange(q.context, q.Name, 0, -1).Val()

    return result
}

func (q *RedisQueue) Size() int64 {
    result := q.client.LLen(q.context, q.Name).Val()

    return result
}
