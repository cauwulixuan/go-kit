/*
 * Copyright 2022 The Inspur AIStation Group Authors.
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

// Note: the example only works with the code within the same release/branch.

package log

import "go.uber.org/zap"

func Debug(message string, fields ...zap.Field) {
	logger.Debug(message, fields...)
}

func Info(message string, fields ...zap.Field) {
	logger.Info(message, fields...)
}

func Warn(message string, fields ...zap.Field) {
	logger.Warn(message, fields...)
}

func DPanic(message string, fields ...zap.Field) {
	logger.DPanic(message, fields...)
}

func Panic(message string, fields ...zap.Field) {
	logger.Panic(message, fields...)
}

func Error(message string, fields ...zap.Field) {
	logger.Error(message, fields...)
}

func Fatal(message string, fields ...zap.Field) {
	logger.Fatal(message, fields...)
}
