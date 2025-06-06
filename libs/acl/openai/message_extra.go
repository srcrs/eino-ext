/*Add commentMore actions
 * Copyright 2025 CloudWeGo Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package openai

import (
	"strings"

	"github.com/cloudwego/eino/compose"
	"github.com/cloudwego/eino/schema"
)

const (
	extraKeyReasoningContent = "_eino_openai_reasoning_content"
)

type reasoningContentType string

func init() {
	compose.RegisterStreamChunkConcatFunc(func(ts []reasoningContentType) (reasoningContentType, error) {
		sb := strings.Builder{}
		for _, t := range ts {
			sb.WriteString(string(t))
		}
		return reasoningContentType(sb.String()), nil
	})

	_ = compose.RegisterSerializableType[reasoningContentType]("_eino_ext_openai_reasoning_content_type")
}

func SetReasoningContent(message *schema.Message, content string) {
	if message == nil {
		return
	}
	if message.Extra == nil {
		message.Extra = make(map[string]interface{})
	}
	message.Extra[extraKeyReasoningContent] = reasoningContentType(content)
}

func GetReasoningContent(message *schema.Message) (string, bool) {
	if message == nil || message.Extra == nil {
		return "", false
	}
	result, ok := message.Extra[extraKeyReasoningContent].(reasoningContentType)
	return string(result), ok
}
