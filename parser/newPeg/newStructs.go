package newPeg
//-------------------------------------------------------------------------------------------//
//All related to source
type SourceJob struct{
	Type 		src
	OperateOn	interface{}
}
//source file reader type
type SrcFile struct{
	fileName string
}
//source fake creator type
type SrcFake struct{
	number int64
}

//source branch type
type SrcBranch struct{
	identifier string
}
//---------------------------------------------------------------------------------------------//

//********************************************************************************************//
//all related to transform job and its subtypes
type TransformJob struct{
	Type 	  trns
	OperateOn interface{}
}
// do type filter job transform
type Filter func(map[string]interface{}) (bool, error)

//all aggregator transforms
type Count struct {
	Alias  string
	Filter Filter
}
//*******************************************************************************************//


//###########################################################################################//
//all related to sink job and its subtypes
type SinkJob struct{
	Type snk
	OperateOn interface{}
}

type StdOut struct{
	Args interface{}
	Header []string
}

type InTo struct{
	streamTo string
}


