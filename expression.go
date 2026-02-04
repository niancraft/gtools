package gtools

import "golang.org/x/exp/constraints"

// Ternary 三元操作符
// 根据expr条件返回a或b
func Ternary[T any](expr bool, a, b T) T {
	if expr {
		return a
	}
	return b
}

// LimitValue 限制值在指定范围内
// 如果v小于min，返回min；如果v大于max，返回max；否则返回v
func LimitValue[T constraints.Ordered](v T, min T, max T) T {
	if v < min {
		return min
	}
	if v > max {
		return max
	}
	return v
}

// LimitValueSafe 安全的限制值函数，自动处理min>max的情况
// 如果min大于max，会自动交换两者
func LimitValueSafe[T constraints.Ordered](v T, min T, max T) T {
	if min > max {
		min, max = max, min
	}

	if v < min {
		return min
	}
	if v > max {
		return max
	}
	return v
}
