package templatepool

// type TestRunnersServiceParams struct {
// 	message domain.Message
// 	context domain.UserContext
// 	state   domain.UserState
// 	input   RunnerInput
// 	flows   map[string]StepHash
// 	expect  Expect
// }

// type Expect struct {
// 	state    domain.State
// 	messages []domain.Message
// }

// func Start(message domain.Message, context domain.UserContext, messages []domain.Message) []domain.Component {
// 	return []domain.Component{
// 		domain.NewText("Hello"),
// 		domain.NewGoTo("begin", "end"),
// 	}
// }

// func End(message domain.Message, context domain.UserContext, messages []domain.Message) []domain.Component {
// 	return []domain.Component{
// 		domain.NewText("World"),
// 		domain.NewHold("begin", "start"),
// 	}
// }

// func StopTest(message domain.Message, context domain.UserContext, messages []domain.Message) []domain.Component {
// 	return []domain.Component{
// 		domain.NewText("Hello"),
// 		domain.NewHold("begin", "end"),
// 	}
// }

// func MockTestParams(start Step, end Step, expectState domain.State) TestRunnersServiceParams {
// 	message := domain.NewTextMessage("Hi")
// 	context := domain.NewUserContext()
// 	userState := domain.NewUserState("1", "begin", "start")
// 	input := NewRunnerInput(message, context, userState)
// 	flows := map[string]StepHash{
// 		"begin": {
// 			"start": start,
// 			"end":   end,
// 		},
// 	}
// 	return TestRunnersServiceParams{
// 		message: message,
// 		context: context,
// 		state:   userState,
// 		input:   input,
// 		flows:   flows,
// 		expect: Expect{
// 			state: expectState,
// 			messages: []domain.Message{
// 				domain.NewTextMessage("Hello"),
// 				domain.NewTextMessage("World"),
// 			},
// 		},
// 	}
// }
// func TestRunnerService(t *testing.T) {
// 	var tests = []TestRunnersServiceParams{
// 		MockTestParams(Start, End, domain.NewState("begin", "start")),
// 		MockTestParams(Stop, End, domain.NewState("begin", "end")),
// 	}
// 	for _, test := range tests {
// 		Describe(test, t)
// 	}
// }

// func Describe(i TestRunnersServiceParams, t *testing.T) {
// 	runner := NewRunnerService(i.flows)
// 	botMessages, newState := runner.Run(i.input)
// 	if !cmp.Equal(botMessages, i.expect.messages) {
// 		expect := spew.Sdump(i.expect.messages)
// 		result := spew.Sdump(botMessages)
// 		t.Errorf("Result message incorrect, got: %v, want: %v.", result, expect)
// 	}
// 	if !cmp.Equal(newState, i.expect.state) {
// 		expect := spew.Sdump(i.expect.state)
// 		result := spew.Sdump(newState)
// 		t.Errorf("Result state incorrect, got: %v, want: %v.", result, expect)
// 	}
// }
