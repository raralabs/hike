package newPeg
//-------------------------------------------------------------------------------------------//
//All related to source
type SourceJob struct{
	Type 		src
	OperateOn	interface{}
}
//source file reader type
type SrcFile struct{
	FileName string
}
//source fake creator type
type SrcFake struct{
	Number int64
}

//source branch type
type SrcBranch struct{
	Identifier string
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
type Aggregator struct{
	Type aggregator
}
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
	StreamTo string
}



