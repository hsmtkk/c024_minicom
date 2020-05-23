package minicom

type Computer interface {
	Set(int, int)
	Add(int)
	Sub(int)
	Output() (int, int)
}

func New() Computer {
	return &computerImpl{}
}

type computerImpl struct {
	reg1 int
	reg2 int
}

func (com *computerImpl) Set(reg int, val int) {
	if reg == 1 {
		com.reg1 = val
	} else if reg == 2 {
		com.reg2 = val
	}
}

func (com *computerImpl) Add(n int) {
	com.reg2 = com.reg1 + n
}

func (com *computerImpl) Sub(n int) {
	com.reg2 = com.reg1 - n
}

func (com *computerImpl) Output() (int, int) {
	return com.reg1, com.reg2
}
