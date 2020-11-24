package newMulti

import (

	"fmt"
	//"github.com/raralabs/canal/ext/transforms"
	"os"
	"log"
	canalSnk "github.com/raralabs/canal/ext/sinks"
	"github.com/raralabs/hike/plugins/canal/sinks"
	"github.com/raralabs/canal/core/transforms/agg"
	canalSrc "github.com/raralabs/canal/ext/sources"
	"github.com/raralabs/canal/core/pipeline"
	"github.com/raralabs/hike/plugins/canal/sources"
	"github.com/raralabs/canal/core/message"
	"github.com/raralabs/canal/core/message/content"
	"github.com/raralabs/canal/ext/transforms/aggregates"
	"github.com/raralabs/hike/parser/newPeg"
	"github.com/raralabs/canal/ext/transforms/doFn"
)


//This function returns the executorPod based on the
//types of stage. If the stages are primary the pod contains
//pipeline.Executor and other fileds are left default while
//if its a secondary stage i.e Branch (source) or InTo (sink)
//then streamLbl contains streamFrom and StreamTo respectively
func getExecutor(stg interface{}) executorPod {
	var executor executorPod
	var exec pipeline.Executor
	switch stg.(type) {
	case newPeg.SourceJob:
		currentStg := stg.(newPeg.SourceJob)
		if currentStg.Type == newPeg.BRANCHJOB||currentStg.Type== newPeg.SECSOURCE{
			switch currentStg.Type {
			case newPeg.BRANCHJOB:
				operator := currentStg.OperateOn.(newPeg.SrcBranch)
				streamFrom := operator.Identifier
				executor = executorPod{exec,false, streamFrom,"SOURCE"}

			case newPeg.SECSOURCE:
				operator := currentStg.OperateOn.(newPeg.SrcSec)
				streamFrom := operator.Sources
				executor = executorPod{exec,false, streamFrom,"SOURCE"}

			}
			}else{
			exec = getSourceExecutor(currentStg)
			executor = executorPod{exec,true,"","SOURCE"}
		}
		
	case newPeg.TransformJob:
		currentStg := stg.(newPeg.TransformJob)
		exec = getTransformExecutor(currentStg)
		executor = executorPod{exec,true,"","TRANSFORM"}

	case newPeg.SinkJob:
		currentStg := stg.(newPeg.SinkJob)
		if currentStg.Type == newPeg.INTO{
			operator := currentStg.OperateOn.(newPeg.InTo)
			streamTo := operator.StreamTo
			executor = executorPod{exec,false,streamTo,"SINK"}
		}else{
			exec = getSinkExecutor(currentStg)
			executor = executorPod{exec,true,"","SINK"}

		}
	}
	return executor
}




func getTransformExecutor(stg newPeg.TransformJob)pipeline.Executor{
	var exec pipeline.Executor
	switch stg.Type{
	case newPeg.AGGJOB:
		stgContent := stg.OperateOn.(newPeg.AggNodeJob)
		aggFuncs := stgContent.Functions
		var aggs []agg.IAggFuncTemplate
		after := func(m message.Msg, proc pipeline.IProcessorForExecutor, cntnt, pContent []content.IContent) {
			contents := m.Content()
			if contents != nil {
				if v, ok := contents.Get("eof"); ok {
					if v.Val == true {
						proc.Result(m, contents, nil)
						proc.Done()
						return
					}
				}
			}
			for i := range cntnt {
				proc.Result(m, cntnt[i], pContent[i])
			}
		}
		for _, ags := range aggFuncs {
			switch ag := ags.(type) {
			case newPeg.Count:
				cnt := aggregates.NewCount(ag.Alias, func(m map[string]interface{}) bool {
					if v, ok := m["eof"]; ok {
						if v == true {
							return false
						}
					}

					match, err := ag.Filter(m)
					if err != nil {
						log.Panic(err)
					}
					return match
				})
				aggs = append(aggs, cnt)

			case newPeg.Max:
				mx := aggregates.NewMax(ag.Alias, ag.Field, func(m map[string]interface{}) bool {
					if v, ok := m["eof"]; ok {
						if v == true {
							return false
						}
					}

					match, err := ag.Filter(m)
					if err != nil {
						log.Panicf("Max Filter Error: %v", err)
					}
					return match
				})
				aggs = append(aggs, mx)

			case newPeg.Min:
				mn := aggregates.NewMin(ag.Alias, ag.Field, func(m map[string]interface{}) bool {
					if v, ok := m["eof"]; ok {
						if v == true {
							return false
						}
					}

					match, err := ag.Filter(m)
					if err != nil {
						log.Panic(err)
					}
					return match
				})
				aggs = append(aggs, mn)

			case newPeg.Avg:
				avg := aggregates.NewAvg(ag.Alias, ag.Field, func(m map[string]interface{}) bool {
					if v, ok := m["eof"]; ok {
						if v == true {
							return false
						}
					}

					match, err := ag.Filter(m)
					if err != nil {
						log.Panic(err)
					}
					return match
				})
				aggs = append(aggs, avg)

			case newPeg.Sum:
				sum := aggregates.NewSum(ag.Alias, ag.Field, func(m map[string]interface{}) bool {
					if v, ok := m["eof"]; ok {
						if v == true {
							return false
						}
					}

					match, err := ag.Filter(m)
					if err != nil {
						log.Panic(err)
					}
					return match
				})
				aggs = append(aggs, sum)

			case newPeg.Variance:
				variance := aggregates.NewVariance(ag.Alias, ag.Field, func(m map[string]interface{}) bool {
					if v, ok := m["eof"]; ok {
						if v == true {
							return false
						}
					}

					match, err := ag.Filter(m)
					if err != nil {
						log.Panic(err)
					}
					return match
				})
				aggs = append(aggs, variance)

			case newPeg.DistinctCount:
				variance := aggregates.NewDCount(ag.Alias, ag.Field, func(m map[string]interface{}) bool {
					if v, ok := m["eof"]; ok {
						if v == true {
							return false
						}
					}

					match, err := ag.Filter(m)
					if err != nil {
						log.Panic(err)
					}
					return match
				})
				aggs = append(aggs, variance)

			case newPeg.Mode:
				variance := aggregates.NewMode(ag.Alias, ag.Field, func(m map[string]interface{}) bool {
					if v, ok := m["eof"]; ok {
						if v == true {
							return false
						}
					}

					match, err := ag.Filter(m)
					if err != nil {
						log.Panic(err)
					}
					return match
				})
				aggs = append(aggs, variance)

			case newPeg.Quantile:
				quantile := aggregates.NewQuantile(ag.Alias, ag.Field, ag.Qth,
					func(m map[string]interface{}) bool {
						if v, ok := m["eof"]; ok {
							if v == true {
								return false
							}
						}

						match, err := ag.Filter(m)
						if err != nil {
							log.Panic(err)
						}
						return match
					})
				aggs = append(aggs, quantile)

			case newPeg.Covariance:
				cov := aggregates.NewCovariance(ag.Alias, ag.Field1, ag.Field2,
					func(m map[string]interface{}) bool {
						if v, ok := m["eof"]; ok {
							if v == true {
								return false
							}
						}

						match, err := ag.Filter(m)
						if err != nil {
							log.Panic(err)
						}
						return match
					})
				aggs = append(aggs, cov)

			case newPeg.Correlation:
				corr := aggregates.NewCorrelation(ag.Alias, ag.Field1, ag.Field2,
					func(m map[string]interface{}) bool {
						if v, ok := m["eof"]; ok {
							if v == true {
								return false
							}
						}

						match, err := ag.Filter(m)
						if err != nil {
							log.Panic(err)
						}
						return match
					})
				aggs = append(aggs, corr)

			case newPeg.PValue:
				pval := aggregates.NewPValue(ag.Alias, ag.Field1, ag.Field2,
					func(m map[string]interface{}) bool {
						if v, ok := m["eof"]; ok {
							if v == true {
								return false
							}
						}

						match, err := ag.Filter(m)
						if err != nil {
							log.Panic(err)
						}
						return match
					})
				aggs = append(aggs, pval)
			}
		}

		aggregator := agg.NewAggregator(aggs, after, stgContent.GroupBy...)

		exec = aggregator.Function()
		exec.SetName("agg")

	case newPeg.DOJOB:
		stgContent := stg.OperateOn.(newPeg.DoNodeJob)
		doneFunc := func(m message.Msg) bool {
			cont := m.Content()

			if cont != nil {
				if v, ok := cont.Get("eof"); ok {
					if v.Val == true {
						return true
					}
				}
			}
			return false
		}

		switch doJob := stgContent.Function.(type) {
		case newPeg.Filter:
			exec = doFn.FilterFunction(doJob, doneFunc)
			exec.SetName("filter")

		case newPeg.Select:
			fields := doJob.Fields
			exec = doFn.SelectFunction(fields, doneFunc)
			exec.SetName("select")

		case newPeg.Pick:
			dsc := doJob.Desc
			//if !strIn(availablePicks, dsc) {
			//	log.Panicf("Can't pick by: %s", dsc)
			//}
			exec = doFn.PickFunction(dsc, doJob.Num, doneFunc)
			exec.SetName("pick")

		case newPeg.Sort:
			fld := doJob.Field
			exec = doFn.SortFunction(fld, doneFunc)
			exec.SetName("sort")

		case newPeg.Batch:
			exec = doFn.BatchAgg(doneFunc)
			exec.SetName("batch")

		case newPeg.Enrich:
			if expr := doJob.Expr; expr != nil {
				exec = doFn.EnrichFunction(doJob.Field, expr, doneFunc)
				exec.SetName("enrich")
			}
		}
	case newPeg.JOINJOB:
		stgContent := stg.OperateOn.(newPeg.JoinNodeJob)
		attributes := stgContent.Attributes
		selFields := attributes.SelectFields
		condition := attributes.JoinCondition
		leftFields := condition.LeftFields
		rightFields := condition.RightFields
		switch stgContent.Type{
		case "INNER":
			fmt.Println("goodgame",selFields,leftFields,rightFields)
			//exec,_:=transforms.NewJoinProcessor()
			//exec,_ := canalTrns.NewJoinProcessor("inner",leftFields,rightFields,selFields,"","S1","S2")
			//NewJoinProcessor(name string,fields1,fields2,selectFields []string,
			//	joinType joinUsingHshMap.JoinType,firstTableName,secondTableName string)
			//(*joinProcessor,error) {
			//exec.SetName("join")
		}

	}
	return exec
}

//generates all the sink execcutors based on the types
//input is the stg of type SinkJob and output is the
//respective executor.
func getSinkExecutor(stg newPeg.SinkJob)pipeline.Executor{
	var exec pipeline.Executor
	switch stg.Type{
	case newPeg.STDOUT:
		s :=stg.OperateOn.(newPeg.StdOut)
		exec = sinks.NewPrettyPrinter(os.Stdout, 10, s.Header...)
		exec.SetName("stdout")

	case newPeg.BLACKHOLE:
		exec = canalSnk.NewBlackholeSink()
		exec.SetName("blackhole")

	default:
		fmt.Println("plot disabled for now")
		}
	return exec
}


//generates all the source execcutors based on the types
//input is the stg of type SourceJob and output is the
//respective executor.
func getSourceExecutor(stg newPeg.SourceJob)pipeline.Executor{
	var exec pipeline.Executor
	switch stg.Type {
	case newPeg.FILEJOB:

		params := stg.OperateOn.(newPeg.SrcFile)
		fileName := params.FileName
		f, err := os.Open(fileName)
		if err != nil {
			log.Panic(err)
		}
		log.Printf("[INFO] %s file opened as source.", fileName)
		exec = sources.NewCsvReader(f, -1)
		exec.SetName("file")

	case newPeg.FAKEJOB:
		params := stg.OperateOn.(newPeg.SrcFake)
		numData := params.Number
		log.Printf("[INFO] Generating %d fake data", numData)
		exec = canalSrc.NewFaker(numData, nil)
		exec.SetName("fake")
	}
	return exec
}
