package simple

type IStrategy interface {
	do(int, int) int
}

type add struct {
}

func (*add) do(a, b int) int {
	return a + b
}

type reduce struct {
}

func (*reduce) do(a, b int) int {
	return a - b
}

type Operator struct {
	strategy IStrategy
}

func (op *Operator) setStrategy(strategy IStrategy) {
	op.strategy = strategy
}

func (op *Operator) calculate(a, b int) int {
	return op.strategy.do(a, b)
}
