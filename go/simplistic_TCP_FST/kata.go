package kata

type State uint

const (
	CLOSED State = iota
	LISTEN
	SYN_SENT
	SYN_RCVD
	ESTABLISHED
	CLOSE_WAIT
	LAST_ACK
	FIN_WAIT_1
	FIN_WAIT_2
	CLOSING
	TIME_WAIT
)

const (
	APP_PASSIVE_OPEN = "APP_PASSIVE_OPEN"
	APP_ACTIVE_OPEN  = "APP_ACTIVE_OPEN"
	APP_SEND         = "APP_SEND"
	APP_CLOSE        = "APP_CLOSE"
	APP_TIMEOUT      = "APP_TIMEOUT"
	RCV_SYN          = "RCV_SYN"
	RCV_ACK          = "RCV_ACK"
	RCV_SYN_ACK      = "RCV_SYN_ACK"
	RCV_FIN          = "RCV_FIN"
	RCV_FIN_ACK      = "RCV_FIN_ACK"
)

type StateMashine struct {
	State State
}

func (p *StateMashine) Execute(cmd string) string {
	t := StateCmd{p.State, cmd}
	if v := StateMigration[t]; v == nil {
		return "ERROR"
	} else {
		return v(&p.State)
	}
}

type StateCmd struct {
	State State
	Cmd   string
}

type Migration func(state *State) string

var StateMigration = map[StateCmd]Migration{
	{CLOSED, APP_PASSIVE_OPEN}: func(state *State) string {
		*state = LISTEN
		return "LISTEN"
	},
	{CLOSED, APP_ACTIVE_OPEN}: func(state *State) string {
		*state = SYN_SENT
		return "SYN_SENT"
	},
	{LISTEN, RCV_SYN}: func(state *State) string {
		*state = SYN_RCVD
		return "SYN_RCVD"
	},
	{LISTEN, APP_SEND}: func(state *State) string {
		*state = SYN_SENT
		return "SYN_SENT"
	},
	{LISTEN, APP_CLOSE}: func(state *State) string {
		*state = CLOSED
		return "CLOSED"
	},
	{SYN_RCVD, APP_CLOSE}: func(state *State) string {
		*state = FIN_WAIT_1
		return "FIN_WAIT_1"
	},
	{SYN_RCVD, RCV_ACK}: func(state *State) string {
		*state = ESTABLISHED
		return "ESTABLISHED"
	},
	{SYN_SENT, RCV_SYN}: func(state *State) string {
		*state = SYN_RCVD
		return "SYN_RCVD"
	},
	{SYN_SENT, RCV_SYN_ACK}: func(state *State) string {
		*state = ESTABLISHED
		return "ESTABLISHED"
	},
	{SYN_SENT, APP_CLOSE}: func(state *State) string {
		*state = CLOSED
		return "CLOSED"
	},
	{ESTABLISHED, APP_CLOSE}: func(state *State) string {
		*state = FIN_WAIT_1
		return "FIN_WAIT_1"
	},
	{ESTABLISHED, RCV_FIN}: func(state *State) string {
		*state = CLOSE_WAIT
		return "CLOSE_WAIT"
	},
	{FIN_WAIT_1, RCV_FIN}: func(state *State) string {
		*state = CLOSING
		return "CLOSING"
	},
	{FIN_WAIT_1, RCV_FIN_ACK}: func(state *State) string {
		*state = TIME_WAIT
		return "TIME_WAIT"
	},
	{FIN_WAIT_1, RCV_ACK}: func(state *State) string {
		*state = FIN_WAIT_2
		return "FIN_WAIT_2"
	},
	{CLOSING, RCV_ACK}: func(state *State) string {
		*state = TIME_WAIT
		return "TIME_WAIT"
	},
	{FIN_WAIT_2, RCV_FIN}: func(state *State) string {
		*state = TIME_WAIT
		return "TIME_WAIT"
	},
	{TIME_WAIT, APP_TIMEOUT}: func(state *State) string {
		*state = CLOSED
		return "CLOSED"
	},
	{CLOSE_WAIT, APP_CLOSE}: func(state *State) string {
		*state = LAST_ACK
		return "LAST_ACK"
	},
	{LAST_ACK, RCV_ACK}: func(state *State) string {
		*state = CLOSED
		return "CLOSED"
	},
}

func TraverseTCPStates(events []string) (r string) {
	mashine := StateMashine{State: CLOSED}
	for _, v := range events {
		if r = mashine.Execute(v); r == "ERROR" {
			return
		}
	}
	return
}
