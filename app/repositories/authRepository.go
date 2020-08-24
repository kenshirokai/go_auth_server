package repositories

import (
	"github.com/gomodule/redigo/redis"
	"github.com/kenshirokai/go_app_server/utils"
)

type IAuthRepository interface {
	FindByCode(code string) (info utils.AuthFlowInfo, err error)
	SetCodeValues(code string, dto utils.AuthFlowInfo) (err error)
}

//認証認可のフローで一時的にキャッシュが必要になった場合に使用する
type AuthRepository struct {
	pool *redis.Pool
}

func NewAuthRepository(pool *redis.Pool) AuthRepository {
	return AuthRepository{
		pool: pool,
	}
}

func (repo AuthRepository) FindByCode(code string) (info utils.AuthFlowInfo, err error) {
	conn := repo.pool.Get()
	defer conn.Close()
	result, err := redis.Values(conn.Do("HGETALL", code))
	if err != nil {
		return
	}
	err = redis.ScanStruct(result, &info)
	return
}

func (repo AuthRepository) SetCodeValues(code string, dto utils.AuthFlowInfo) (err error) {
	conn := repo.pool.Get()
	defer conn.Close()
	_, err = conn.Do("HSET", code, "state", dto.State)
	_, err = conn.Do("HSET", code, "redirect_uri", dto.RedirectURI)
	return
}
