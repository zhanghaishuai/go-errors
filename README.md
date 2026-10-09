[简体中文](README.md) | [English](README.en.md)

# errors

一个受 [`github.com/pkg/errors`](https://github.com/pkg/errors) 启发的 Go 错误处理包，并补充了与标准库错误处理模型保持一致的 API。

本包提供错误包装、堆栈跟踪、错误检查和错误组合等功能，同时避免在创建哨兵错误时进行不必要的堆栈捕获。

## 环境要求

* **Go 1.26 或更高版本**

本包使用了标准库的 `errors.AsType` API，因此需要 Go 1.26 或更高版本。

## 安装

```bash
go get github.com/zhanghaishuai/go-errors
```

## 功能特性

* **错误上下文：** 使用 `Wrap` 和 `Wrapf` 为错误添加上下文信息。
* **堆栈跟踪：** 在适合调试的场景中捕获并格式化堆栈跟踪信息。
* **不带堆栈跟踪的哨兵错误：** `New` 创建新错误时不会捕获堆栈，避免在初始化哨兵错误时记录不必要的堆栈信息。
* **错误检查：** 使用 `Is` 和 `As` 检查错误及其包装链。
* **错误解包：** 使用 `Unwrap` 获取被包装的错误，并在适用时使用 `Cause` 获取原始原因。
* **错误组合：** 使用 `Join` 将多个错误组合为一个错误值。
* **泛型类型检查：** 使用 `AsType` 获取指定类型的错误。

## 创建错误

使用 `New` 创建普通错误，包括哨兵错误：

```go
var ErrNotFound = errors.New("not found")
```

与创建错误时会捕获堆栈跟踪的实现不同，本包的 `New` 不会捕获堆栈。

这对于在包初始化阶段声明的哨兵错误尤其有用。此时捕获的堆栈通常反映的是初始化过程，而不是错误后来被返回或处理的位置。

## 为错误添加上下文

使用 `Wrap` 或 `Wrapf` 为错误添加上下文信息，同时保留底层错误：

```go
data, err := io.ReadAll(r)
if err != nil {
    return errors.Wrap(err, "read failed")
}
```

使用 `%+v` 格式化包装后的错误时，可以输出堆栈跟踪信息。实际输出取决于错误的创建和包装方式。

## 检查错误

使用 `Is` 检查错误是否与指定目标匹配，包括错误包装链中的错误：

```go
if errors.Is(err, ErrNotFound) {
    // 处理哨兵错误。
}
```

使用 `As` 在错误链中查找可以赋值给指定目标类型的错误：

```go
var pathErr *fs.PathError
if errors.As(err, &pathErr) {
    // 处理路径错误。
}
```

如果希望使用泛型方式进行类型检查，可以使用 `AsType`：

```go
pathErr, ok := errors.AsType[*fs.PathError](err)
if ok {
    // 处理路径错误。
}
```

## 解包错误

使用 `Unwrap` 获取错误包装链中的下一层错误：

```go
cause := errors.Unwrap(err)
```

如果代码使用原始 `pkg/errors` 的 `Cause` 约定，可以使用 `Cause` 获取实现了 `Cause() error` 方法的错误所对应的底层原因。

例如，相关接口可以定义为：

```go
type causer interface {
    Cause() error
}
```

`Cause` 和 `Unwrap` 的用途相关，但并不完全相同：

* `Unwrap` 遵循标准库的 `Unwrap() error` 约定。
* `Cause` 遵循较早的 `Cause() error` 约定。

## 组合错误

使用 `Join` 将多个错误组合为一个错误值：

```go
err := errors.Join(errRead, errClose)

if errors.Is(err, errRead) {
    // 处理读取错误。
}
```

组合后的错误可以通过 `Is` 和 `As` 进行检查，从而判断组合结果中是否包含指定的错误。

## 兼容性

本包以 `github.com/pkg/errors` 推广的错误处理 API 为基础，并结合了较新版本 Go 标准库提供的 API，包括 `Join` 和 `AsType`。

## 许可证

本项目采用 BSD-2-Clause 许可证。详情请参阅 [LICENSE](LICENSE) 文件。
