package multi

//func TestBuild(t *testing.T) {
//	cmds := []string{
//		"fake(5) | select(age, last_name) into s1",
//		"s1 | stdout()",
//		"s1 | filter(age > 20) | stdout()",
//		"s1 | map twice_age = 2 * age | stdout()",
//		"s1 | map half_age = age / 2 | stdout()",
//	}
//	absTree := Build(cmds...)
//	stream.Plot(absTree)
//
//	p := stream.Build(absTree)
//	c, cancel := context.WithTimeout(context.Background(), 1000*time.Second)
//
//	p.Start(c, cancel)
//}
//
//func TestEngine(t *testing.T) {
//
//	src := &node{
//		id:       1,
//		executor: sources.NewFaker(10, nil),
//	}
//
//	proc1 := &node{
//		id: 2,
//		executor: doFn.FilterFunction(func(m map[string]interface{}) (bool, error) {
//			return true, nil
//		}, func(msg message.Msg) bool {
//			return false
//		}),
//	}
//
//	expr, err := govaluate.NewEvaluableExpression("age * 2")
//	if err != nil {
//		log.Println(err)
//		return
//	}
//
//	proc2 := &node{
//		id: 4,
//		executor: doFn.EnrichFunction("twice_age", expr, func(msg message.Msg) bool {
//			return false
//		}),
//	}
//
//	snk1 := &node{
//		id:       3,
//		executor: sinks.NewPrettyPrinter(os.Stdout, 10),
//	}
//
//	snk2 := &node{
//		id:       5,
//		executor: sinks.NewPrettyPrinter(os.Stdout, 5),
//	}
//
//	src.toNodes = []at.Node{proc1, proc2}
//	proc1.toNodes = []at.Node{snk1}
//	proc2.toNodes = []at.Node{snk2}
//
//	tr := &tree{sources: []at.Node{src}}
//
//	p := stream.Build(tr)
//	c, cancel := context.WithTimeout(context.Background(), 1000*time.Second)
//
//	p.Start(c, cancel)
//}
