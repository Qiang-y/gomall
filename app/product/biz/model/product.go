package model

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v9"
	"github.com/redis/go-redis/v9"
	"golang.org/x/sync/errgroup"
	"gorm.io/gorm"
	"time"
)

type Product struct {
	Base
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Picture     string  `json:"picture"`
	Price       float32 `json:"price"`
	Quantity    uint32  `json:"quantity"`

	Categories []Category `json:"categories" gorm:"many2many:product_category"`
}

func (Product) TableName() string {
	return "product"
}

type ProductQuery struct {
	ctx context.Context
	db  *gorm.DB
}

func (p ProductQuery) GetById(productId int) (product Product, err error) {
	err = p.db.WithContext(p.ctx).Model(&Product{}).First(&product, productId).Error
	return
}

func (p ProductQuery) SearchProducts(q string) (products []*Product, err error) {
	err = p.db.WithContext(p.ctx).Model(&Product{}).Find(&products, "name like ? or description like ?",
		"%"+q+"%", "%"+q+"%").Error
	return
}

func (p ProductQuery) ReduceQuantity(productId int, quantity uint32) error {
	err := p.db.WithContext(p.ctx).Model(&Product{}).
		Where("id = ? AND quantity >= ?", productId, quantity).
		Update("quantity", gorm.Expr("quantity - ?", quantity)).Error
	return err
}

func NewProductQuery(ctx context.Context, db *gorm.DB) *ProductQuery {
	return &ProductQuery{
		ctx: ctx,
		db:  db,
	}
}

// CachedProductQuetry 带Cache的product查询�?
type CachedProductQuetry struct {
	productQuery ProductQuery
	cacheClient  *redis.Client
	prefix       string
}

func (c CachedProductQuetry) GetById(productId int) (product Product, err error) {
	cachedKey := fmt.Sprintf("%s_%s_%d", c.prefix, "product_by_id", productId)
	cachedResult := c.cacheClient.Get(c.productQuery.ctx, cachedKey)

	// 使用闭包构建错误链，中间发生错误即返�?
	err = func() error {
		if err := cachedResult.Err(); err != nil {
			return err
		}
		cachedResultByte, err := cachedResult.Bytes()
		if err != nil {
			return err
		}

		err = json.Unmarshal(cachedResultByte, &product)
		if err != nil {
			return err
		}
		return nil
	}()

	// 如果闭包任何一步出问题返回则从数据库获取数�?
	if err != nil {
		product, err = c.productQuery.GetById(productId)
		if err != nil {
			return Product{}, err
		}

		// 将数据缓存进redis
		encoded, err := json.Marshal(product)
		if err != nil {
			return product, nil
		}
		_ = c.cacheClient.Set(c.productQuery.ctx, cachedKey, encoded, time.Hour)
	}
	return
}

func (c CachedProductQuetry) SearchProducts(q string) (products []*Product, err error) {
	return c.productQuery.SearchProducts(q)
}

func NewCachedProductQuery(ctx context.Context, db *gorm.DB, cacheClient *redis.Client) *CachedProductQuetry {
	return &CachedProductQuetry{
		productQuery: *NewProductQuery(ctx, db),
		cacheClient:  cacheClient,
		prefix:       "shop",
	}
}

type ReduceProduct struct {
	ID       int
	Quantity uint32
}

// ProductMutation 数据库读写分离，用来进行写操�?
type ProductMutation struct {
	//ctx context.Context
	//db  *gorm.DB
	productQuery ProductQuery
	cacheClient  *redis.Client
	lockClient   *redsync.Redsync
	lockPrefix   string
	cachePrefix  string
}

func (pm *ProductMutation) ReduceQuantity(reduceList []ReduceProduct) (bool, error) {
	ctx, cancel := context.WithTimeout(pm.productQuery.ctx, 10*time.Second)
	defer cancel()

	g, ctx := errgroup.WithContext(ctx)

	// 预检�?
	for _, item := range reduceList {
		product, err := pm.productQuery.GetById(item.ID)
		if err != nil {
			return false, err
		}
		if product.Quantity < item.Quantity {
			return false, fmt.Errorf("prodcut: %s have no enough quantity, only have: %d", product.Name, product.Quantity)
		}

		// 删除redis缓存
		cachedKey := fmt.Sprintf("%s_%s_%d", pm.cachePrefix, "product_by_id", item.ID)
		_ = pm.cacheClient.Del(pm.productQuery.ctx, cachedKey)
	}
	// 并发处理
	for _, item := range reduceList {
		item := item // 储存本地副本
		g.Go(func() error {
			mutexName := fmt.Sprintf("%s-%s-%d", pm.lockPrefix, "product", item.ID)
			mutex := pm.lockClient.NewMutex(mutexName)
			// 带超时的�?
			if err := mutex.LockContext(ctx); err != nil {
				return err
			}
			defer mutex.UnlockContext(ctx)

			err := pm.productQuery.ReduceQuantity(item.ID, item.Quantity)
			return err
		})
	}

	if err := g.Wait(); err != nil {
		return false, err
	}
	return true, nil
}

func NewProductMutation(ctx context.Context, db *gorm.DB, redisClient *redis.Client) *ProductMutation {
	return &ProductMutation{
		productQuery: *NewProductQuery(ctx, db),
		cacheClient:  redisClient,
		lockClient:   redsync.New(goredis.NewPool(redisClient)),
		lockPrefix:   "shopLock",
		cachePrefix:  "shop",
	}
}
