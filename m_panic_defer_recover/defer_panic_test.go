package m_panic_defer_recover

import "testing"

func TestDefer(t *testing.T) {

}

func TestPanic(t *testing.T) {
	TestPanic_()
}

func Test_coverPanic(t *testing.T) {
	coverPanic()
}
