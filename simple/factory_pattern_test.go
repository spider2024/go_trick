package simple

import "testing"

// 测试 NewPerson 函数是否正确创建 Person 实例
func TestNewPerson(t *testing.T) {
	name := "John"
	age := 30
	p := NewPerson(name, age)
	if p.Name != name || p.Age != age {
		t.Errorf("NewPerson() error, expected name %s and age %d, got name %s and age %d", name, age, p.Name, p.Age)
	}
	p.Greet()
}

// 测试 NewPersonAbstract 函数是否正确创建 person 实例
func TestNewPersonAbstract(t *testing.T) {
	name := "Alice"
	age := 25
	p := NewPersonAbstract(name, age)
	pConcrete, ok := p.(person) // 类型断言，确保返回的是 person 类型
	if !ok {
		t.Errorf("NewPersonAbstract() error, expected type person, got %T", p)
	}
	if pConcrete.Name != name || pConcrete.Age != age {
		t.Errorf("NewPersonAbstract() error, expected name %s and age %d, got name %s and age %d", name, age, pConcrete.Name, pConcrete.Age)
	}
	p.Greet()
}
