package main
/*
每个测试函数必须导入testing包，测试函数的基本格式（签名）如下：
	func TestName(t *testing.T){
		// ...
	}

测试函数的名字必须以Test开头，可选的后缀名必须以大写字母开头，举几个例子：\
	func TestAdd(t *testing.T){ ... }
	func TestSum(t *testing.T){ ... }
	func TestLog(t *testing.T){ ... }

其中参数t用于报告测试失败和附加的日志信息。 testing.T的拥有的方法如下：
	func (c *T) Error(args ...interface{})
	func (c *T) Errorf(format string, args ...interface{})
	func (c *T) Fail()
	func (c *T) FailNow()
	func (c *T) Failed() bool
	func (c *T) Fatal(args ...interface{})
	func (c *T) Fatalf(format string, args ...interface{})
	func (c *T) Log(args ...interface{})
	func (c *T) Logf(format string, args ...interface{})
	func (c *T) Name() string
	func (t *T) Parallel()
	func (t *T) Run(name string, f func(t *T)) bool
	func (c *T) Skip(args ...interface{})
	func (c *T) SkipNow()
	func (c *T) Skipf(format string, args ...interface{})
	func (c *T) Skipped() bool
*/
func main() {
	
}
