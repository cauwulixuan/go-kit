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

func SDebug(args ...interface{}) {
	Slogger.Debug(args...)
}

func SDebugf(template string, args ...interface{}) {
	Slogger.Debugf(template, args...)
}

func SDebugw(msg string, keysAndValues ...interface{}) {
	Slogger.Debugw(msg, keysAndValues...)
}

func SInfo(args ...interface{}) {
	Slogger.Info(args...)
}

func SInfof(template string, args ...interface{}) {
	Slogger.Infof(template, args...)
}

func SInfow(msg string, keysAndValues ...interface{}) {
	Slogger.Infow(msg, keysAndValues...)
}

func SWarn(args ...interface{}) {
	Slogger.Warn(args...)
}

func SWarnf(template string, args ...interface{}) {
	Slogger.Warnf(template, args...)
}

func SWarnw(msg string, keysAndValues ...interface{}) {
	Slogger.Warnw(msg, keysAndValues...)
}

func SDPanic(args ...interface{}) {
	Slogger.DPanic(args...)
}

func SDPanicf(template string, args ...interface{}) {
	Slogger.DPanicf(template, args...)
}

func SDPanicw(msg string, keysAndValues ...interface{}) {
	Slogger.DPanicw(msg, keysAndValues...)
}

func SPanic(args ...interface{}) {
	Slogger.Panic(args...)
}

func SPanicf(template string, args ...interface{}) {
	Slogger.Panicf(template, args...)
}

func SPanicw(msg string, keysAndValues ...interface{}) {
	Slogger.Panicw(msg, keysAndValues...)
}

func SError(args ...interface{}) {
	Slogger.Error(args...)
}

func SErrorf(template string, args ...interface{}) {
	Slogger.Errorf(template, args...)
}

func SErrorw(msg string, keysAndValues ...interface{}) {
	Slogger.Errorw(msg, keysAndValues...)
}

func SFatal(args ...interface{}) {
	Slogger.Fatal(args...)
}

func SFatalf(template string, args ...interface{}) {
	Slogger.Fatalf(template, args...)
}

func SFatalw(msg string, keysAndValues ...interface{}) {
	Slogger.Fatalw(msg, keysAndValues...)
}
