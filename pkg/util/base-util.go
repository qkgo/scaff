package util

type STR struct {
	Value string
}

func (self *STR) String() string {
	return self.Value
}

type NUM struct {
	Value int
}

type I64 struct {
	Value int64
}
