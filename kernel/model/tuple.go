package model

type Tuple[T1 any] struct {
	T1 T1 `json:"t1"`
}

type Tuple2[T1, T2 any] struct {
	T1 T1 `json:"t1"`
	T2 T2 `json:"t2"`
}
