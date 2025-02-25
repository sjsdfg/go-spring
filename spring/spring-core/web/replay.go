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

package web

import (
	"github.com/go-spring/spring-base/fastdev"
	"github.com/go-spring/spring-base/knife"
	"github.com/go-spring/spring-base/util"
)

// ReplaySessionID 流量回放模式下传递会话 ID 使用的 Header 。
const ReplaySessionID = "REPLAY-SESSION-ID"

// StartReplay 开始流量回放
func StartReplay(ctx Context) {
	session := ctx.GetHeader(ReplaySessionID)
	if session == "" {
		return
	}
	err := knife.Set(ctx.Context(), fastdev.ReplaySessionIDKey, session)
	util.Panic(err).When(err != nil)
}

// StopReplay 停止流量回放
func StopReplay(ctx Context) {

}
