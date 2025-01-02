package game

type Timer struct {
	currentTicks int
	targretTicks int
}

func NewTimer(targretTicks int) *Timer{
	return &Timer{
		currentTicks: 0,
		targretTicks: targretTicks,
	}
}	

func (t *Timer) Update() {
		
	if t.currentTicks < t.targretTicks{
			t.currentTicks++
	}
}

func (t *Timer) IsReady() bool {
	return t.currentTicks >= t.targretTicks
}

func (t *Timer) Reset(){
	t.currentTicks = 0
}

