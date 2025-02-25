/*
 * Copyright 2012-2019 the original author or authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package mock

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-spring/spring-base/log"
	"github.com/go-spring/spring-core/gs"
	"github.com/jinzhu/gorm"
)

func init() {

	gs.Provide(mockDB(func(mock sqlmock.Sqlmock) {
		mock.ExpectQuery("SELECT ENGINE FROM `ENGINES`").WillReturnRows(
			mock.NewRows([]string{"ENGINE"}).AddRow("sql-mock"),
		)
	})).Name("mockDB")

	//gs.Provide(mockRedis(func(mock *redismock.ClientMock) {
	//	mock.On("Set", "key", "ok", time.Second*10).Return(redis.NewStatusResult("", nil))
	//	mock.On("Get", "key").Return(redis.NewStringResult("ok", nil))
	//})).Name("mockRedis")
}

// mockDB 创建 gorm.DB Mock 客户端
func mockDB(fn func(sqlmock.Sqlmock)) func() (*gorm.DB, error) {
	return func() (*gorm.DB, error) {
		log.Info("create sqlmock db")
		db, mock, err := sqlmock.New()
		if err != nil {
			return nil, err
		}
		fn(mock)
		return gorm.Open("mysql", db)
	}
}

//// mockRedis 创建 Redis Mock 客户端
//func mockRedis(fn func(*redismock.ClientMock)) func() redis.Cmdable {
//	return func() redis.Cmdable {
//		mock := redismock.NewMock()
//		fn(mock)
//		return mock
//	}
//}
