# Golang『bread与continue』注意点

## break

_**作用：跳出循环**_

一个 break 的作用范围为该语句出现后的最内部的结构，它可以被用于任何形式的 for 循环（计数器、条件判断等）。但在 switch 或 select 语句中，break 语句的作用结果是跳过整个代码块，执行后续的代码。

## continue

_**作用：跳过当前循环，继续执行下一次循环**_

continue 只能被用于 for 循环中。