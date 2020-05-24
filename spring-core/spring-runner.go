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

package SpringCore

import (
	"errors"
	"reflect"
)

// Runner 执行器
type Runner struct {
	runnable
	ctx *defaultSpringContext
}

// newRunner Runner 的构造函数
func newRunner(ctx *defaultSpringContext, fn interface{}, tags []string) *Runner {

	fnType := reflect.TypeOf(fn)
	if fnType.Kind() != reflect.Func {
		panic(errors.New("fn must be a func"))
	}

	return &Runner{
		ctx: ctx,
		runnable: runnable{
			fn:        fn,
			stringArg: newFnStringBindingArg(fnType, false, tags),
		},
	}
}

// Options 设置 Option 模式函数的参数绑定
func (r *Runner) Options(options ...*optionArg) *Runner {
	r.optionArg = &fnOptionBindingArg{options}
	return r
}

// When 参数为 true 时执行器运行
func (r *Runner) When(ok bool) error {
	if ok {
		return r.run(newDefaultBeanAssembly(r.ctx))
	}
	return nil
}

// On Condition 判断结果为 true 时执行器运行
func (r *Runner) On(cond Condition) error {
	if cond.Matches(r.ctx) {
		return r.run(newDefaultBeanAssembly(r.ctx))
	}
	return nil
}
