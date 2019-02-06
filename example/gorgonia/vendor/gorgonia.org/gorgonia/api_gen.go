package gorgonia

import (
	"gorgonia.org/gorgonia/internal/engine"
	"gorgonia.org/gorgonia/node"
)

// Code generated by genapi, which is a API generation tool for Gorgonia. DO NOT EDIT.

 // Abs performs a pointwise abs.
func Abs(g *Graph, a node.Node) (node.Node, error) { 
	retval := g.g.NewNode().(*engine.Node)
	g.g.AddNode(retval)
	g.g.SetWeightedEdge(g.g.NewWeightedEdge(retval, a, 1.0))
	err := g.g.ApplyOp(engine.NewAbsOperation(), retval)
	return retval, err
}

 // Sign performs a pointwise sign.
func Sign(g *Graph, a node.Node) (node.Node, error) { 
	retval := g.g.NewNode().(*engine.Node)
	g.g.AddNode(retval)
	g.g.SetWeightedEdge(g.g.NewWeightedEdge(retval, a, 1.0))
	err := g.g.ApplyOp(engine.NewSignOperation(), retval)
	return retval, err
}

 // Ceil performs a pointwise ceil.
func Ceil(g *Graph, a node.Node) (node.Node, error) { 
	retval := g.g.NewNode().(*engine.Node)
	g.g.AddNode(retval)
	g.g.SetWeightedEdge(g.g.NewWeightedEdge(retval, a, 1.0))
	err := g.g.ApplyOp(engine.NewCeilOperation(), retval)
	return retval, err
}

 // Floor performs a pointwise floor.
func Floor(g *Graph, a node.Node) (node.Node, error) { 
	retval := g.g.NewNode().(*engine.Node)
	g.g.AddNode(retval)
	g.g.SetWeightedEdge(g.g.NewWeightedEdge(retval, a, 1.0))
	err := g.g.ApplyOp(engine.NewFloorOperation(), retval)
	return retval, err
}

 // Sin performs a pointwise sin.
func Sin(g *Graph, a node.Node) (node.Node, error) { 
	retval := g.g.NewNode().(*engine.Node)
	g.g.AddNode(retval)
	g.g.SetWeightedEdge(g.g.NewWeightedEdge(retval, a, 1.0))
	err := g.g.ApplyOp(engine.NewSinOperation(), retval)
	return retval, err
}

 // Cos performs a pointwise cos.
func Cos(g *Graph, a node.Node) (node.Node, error) { 
	retval := g.g.NewNode().(*engine.Node)
	g.g.AddNode(retval)
	g.g.SetWeightedEdge(g.g.NewWeightedEdge(retval, a, 1.0))
	err := g.g.ApplyOp(engine.NewCosOperation(), retval)
	return retval, err
}

 // Exp performs a pointwise exp.
func Exp(g *Graph, a node.Node) (node.Node, error) { 
	retval := g.g.NewNode().(*engine.Node)
	g.g.AddNode(retval)
	g.g.SetWeightedEdge(g.g.NewWeightedEdge(retval, a, 1.0))
	err := g.g.ApplyOp(engine.NewExpOperation(), retval)
	return retval, err
}

 // Log performs a pointwise log.
func Log(g *Graph, a node.Node) (node.Node, error) { 
	retval := g.g.NewNode().(*engine.Node)
	g.g.AddNode(retval)
	g.g.SetWeightedEdge(g.g.NewWeightedEdge(retval, a, 1.0))
	err := g.g.ApplyOp(engine.NewLogOperation(), retval)
	return retval, err
}

 // Log2 performs a pointwise log2.
func Log2(g *Graph, a node.Node) (node.Node, error) { 
	retval := g.g.NewNode().(*engine.Node)
	g.g.AddNode(retval)
	g.g.SetWeightedEdge(g.g.NewWeightedEdge(retval, a, 1.0))
	err := g.g.ApplyOp(engine.NewLog2Operation(), retval)
	return retval, err
}

 // Neg performs a pointwise neg.
func Neg(g *Graph, a node.Node) (node.Node, error) { 
	retval := g.g.NewNode().(*engine.Node)
	g.g.AddNode(retval)
	g.g.SetWeightedEdge(g.g.NewWeightedEdge(retval, a, 1.0))
	err := g.g.ApplyOp(engine.NewNegOperation(), retval)
	return retval, err
}

 // Square performs a pointwise square.
func Square(g *Graph, a node.Node) (node.Node, error) { 
	retval := g.g.NewNode().(*engine.Node)
	g.g.AddNode(retval)
	g.g.SetWeightedEdge(g.g.NewWeightedEdge(retval, a, 1.0))
	err := g.g.ApplyOp(engine.NewSquareOperation(), retval)
	return retval, err
}

 // Sqrt performs a pointwise sqrt.
func Sqrt(g *Graph, a node.Node) (node.Node, error) { 
	retval := g.g.NewNode().(*engine.Node)
	g.g.AddNode(retval)
	g.g.SetWeightedEdge(g.g.NewWeightedEdge(retval, a, 1.0))
	err := g.g.ApplyOp(engine.NewSqrtOperation(), retval)
	return retval, err
}

 // Inverse performs a pointwise inverse.
func Inverse(g *Graph, a node.Node) (node.Node, error) { 
	retval := g.g.NewNode().(*engine.Node)
	g.g.AddNode(retval)
	g.g.SetWeightedEdge(g.g.NewWeightedEdge(retval, a, 1.0))
	err := g.g.ApplyOp(engine.NewInverseOperation(), retval)
	return retval, err
}

 // InverseSqrt performs a pointwise inversesqrt.
func InverseSqrt(g *Graph, a node.Node) (node.Node, error) { 
	retval := g.g.NewNode().(*engine.Node)
	g.g.AddNode(retval)
	g.g.SetWeightedEdge(g.g.NewWeightedEdge(retval, a, 1.0))
	err := g.g.ApplyOp(engine.NewInverseSqrtOperation(), retval)
	return retval, err
}

 // Cube performs a pointwise cube.
func Cube(g *Graph, a node.Node) (node.Node, error) { 
	retval := g.g.NewNode().(*engine.Node)
	g.g.AddNode(retval)
	g.g.SetWeightedEdge(g.g.NewWeightedEdge(retval, a, 1.0))
	err := g.g.ApplyOp(engine.NewCubeOperation(), retval)
	return retval, err
}

 // Tanh performs a pointwise tanh.
func Tanh(g *Graph, a node.Node) (node.Node, error) { 
	retval := g.g.NewNode().(*engine.Node)
	g.g.AddNode(retval)
	g.g.SetWeightedEdge(g.g.NewWeightedEdge(retval, a, 1.0))
	err := g.g.ApplyOp(engine.NewTanhOperation(), retval)
	return retval, err
}

 // Sigmoid performs a pointwise sigmoid.
func Sigmoid(g *Graph, a node.Node) (node.Node, error) { 
	retval := g.g.NewNode().(*engine.Node)
	g.g.AddNode(retval)
	g.g.SetWeightedEdge(g.g.NewWeightedEdge(retval, a, 1.0))
	err := g.g.ApplyOp(engine.NewSigmoidOperation(), retval)
	return retval, err
}

 // Log1p performs a pointwise log1p.
func Log1p(g *Graph, a node.Node) (node.Node, error) { 
	retval := g.g.NewNode().(*engine.Node)
	g.g.AddNode(retval)
	g.g.SetWeightedEdge(g.g.NewWeightedEdge(retval, a, 1.0))
	err := g.g.ApplyOp(engine.NewLog1pOperation(), retval)
	return retval, err
}

 // Expm1 performs a pointwise expm1.
func Expm1(g *Graph, a node.Node) (node.Node, error) { 
	retval := g.g.NewNode().(*engine.Node)
	g.g.AddNode(retval)
	g.g.SetWeightedEdge(g.g.NewWeightedEdge(retval, a, 1.0))
	err := g.g.ApplyOp(engine.NewExpm1Operation(), retval)
	return retval, err
}

 // Softplus performs a pointwise softplus.
func Softplus(g *Graph, a node.Node) (node.Node, error) { 
	retval := g.g.NewNode().(*engine.Node)
	g.g.AddNode(retval)
	g.g.SetWeightedEdge(g.g.NewWeightedEdge(retval, a, 1.0))
	err := g.g.ApplyOp(engine.NewSoftplusOperation(), retval)
	return retval, err
}

// Add perfoms a pointwise add operation.
func Add(g *Graph, a, b node.Node) (node.Node, error) { 
	retval := g.g.NewNode().(*engine.Node)
	g.g.AddNode(retval)
	g.g.SetWeightedEdge(g.g.NewWeightedEdge(retval, a, 1.0))
	g.g.SetWeightedEdge(g.g.NewWeightedEdge(retval, b, 2.0))
	err := g.g.ApplyOp(engine.NewAddOperation(nil,nil), retval)
	return retval, err
}

// Sub perfoms a pointwise sub operation.
func Sub(g *Graph, a, b node.Node) (node.Node, error) { 
	retval := g.g.NewNode().(*engine.Node)
	g.g.AddNode(retval)
	g.g.SetWeightedEdge(g.g.NewWeightedEdge(retval, a, 1.0))
	g.g.SetWeightedEdge(g.g.NewWeightedEdge(retval, b, 2.0))
	err := g.g.ApplyOp(engine.NewSubOperation(nil,nil), retval)
	return retval, err
}

// HadamardProd perfoms a pointwise hadamardprod operation.
func HadamardProd(g *Graph, a, b node.Node) (node.Node, error) { 
	retval := g.g.NewNode().(*engine.Node)
	g.g.AddNode(retval)
	g.g.SetWeightedEdge(g.g.NewWeightedEdge(retval, a, 1.0))
	g.g.SetWeightedEdge(g.g.NewWeightedEdge(retval, b, 2.0))
	err := g.g.ApplyOp(engine.NewHadamardProdOperation(nil,nil), retval)
	return retval, err
}

// HadamardDiv perfoms a pointwise hadamarddiv operation.
func HadamardDiv(g *Graph, a, b node.Node) (node.Node, error) { 
	retval := g.g.NewNode().(*engine.Node)
	g.g.AddNode(retval)
	g.g.SetWeightedEdge(g.g.NewWeightedEdge(retval, a, 1.0))
	g.g.SetWeightedEdge(g.g.NewWeightedEdge(retval, b, 2.0))
	err := g.g.ApplyOp(engine.NewHadamardDivOperation(nil,nil), retval)
	return retval, err
}

// Pow perfoms a pointwise pow operation.
func Pow(g *Graph, a, b node.Node) (node.Node, error) { 
	retval := g.g.NewNode().(*engine.Node)
	g.g.AddNode(retval)
	g.g.SetWeightedEdge(g.g.NewWeightedEdge(retval, a, 1.0))
	g.g.SetWeightedEdge(g.g.NewWeightedEdge(retval, b, 2.0))
	err := g.g.ApplyOp(engine.NewPowOperation(nil,nil), retval)
	return retval, err
}

// Lt perfoms a pointwise lt operation.
// retSame indicates if the data type of the return value should be the same as the input data type. It defaults to Bool otherwise.
func Lt(g *Graph, a, b node.Node, retSame bool) (node.Node, error) { 
	retval := g.g.NewNode().(*engine.Node)
	g.g.AddNode(retval)
	g.g.SetWeightedEdge(g.g.NewWeightedEdge(retval, a, 1.0))
	g.g.SetWeightedEdge(g.g.NewWeightedEdge(retval, b, 2.0))
	err := g.g.ApplyOp(engine.NewLtOperation(nil,nil,retSame), retval)
	return retval, err
}

// Gt perfoms a pointwise gt operation.
// retSame indicates if the data type of the return value should be the same as the input data type. It defaults to Bool otherwise.
func Gt(g *Graph, a, b node.Node, retSame bool) (node.Node, error) { 
	retval := g.g.NewNode().(*engine.Node)
	g.g.AddNode(retval)
	g.g.SetWeightedEdge(g.g.NewWeightedEdge(retval, a, 1.0))
	g.g.SetWeightedEdge(g.g.NewWeightedEdge(retval, b, 2.0))
	err := g.g.ApplyOp(engine.NewGtOperation(nil,nil,retSame), retval)
	return retval, err
}

// Lte perfoms a pointwise lte operation.
// retSame indicates if the data type of the return value should be the same as the input data type. It defaults to Bool otherwise.
func Lte(g *Graph, a, b node.Node, retSame bool) (node.Node, error) { 
	retval := g.g.NewNode().(*engine.Node)
	g.g.AddNode(retval)
	g.g.SetWeightedEdge(g.g.NewWeightedEdge(retval, a, 1.0))
	g.g.SetWeightedEdge(g.g.NewWeightedEdge(retval, b, 2.0))
	err := g.g.ApplyOp(engine.NewLteOperation(nil,nil,retSame), retval)
	return retval, err
}

// Gte perfoms a pointwise gte operation.
// retSame indicates if the data type of the return value should be the same as the input data type. It defaults to Bool otherwise.
func Gte(g *Graph, a, b node.Node, retSame bool) (node.Node, error) { 
	retval := g.g.NewNode().(*engine.Node)
	g.g.AddNode(retval)
	g.g.SetWeightedEdge(g.g.NewWeightedEdge(retval, a, 1.0))
	g.g.SetWeightedEdge(g.g.NewWeightedEdge(retval, b, 2.0))
	err := g.g.ApplyOp(engine.NewGteOperation(nil,nil,retSame), retval)
	return retval, err
}

// Eq perfoms a pointwise eq operation.
// retSame indicates if the data type of the return value should be the same as the input data type. It defaults to Bool otherwise.
func Eq(g *Graph, a, b node.Node, retSame bool) (node.Node, error) { 
	retval := g.g.NewNode().(*engine.Node)
	g.g.AddNode(retval)
	g.g.SetWeightedEdge(g.g.NewWeightedEdge(retval, a, 1.0))
	g.g.SetWeightedEdge(g.g.NewWeightedEdge(retval, b, 2.0))
	err := g.g.ApplyOp(engine.NewEqOperation(nil,nil,retSame), retval)
	return retval, err
}

// Ne perfoms a pointwise ne operation.
// retSame indicates if the data type of the return value should be the same as the input data type. It defaults to Bool otherwise.
func Ne(g *Graph, a, b node.Node, retSame bool) (node.Node, error) { 
	retval := g.g.NewNode().(*engine.Node)
	g.g.AddNode(retval)
	g.g.SetWeightedEdge(g.g.NewWeightedEdge(retval, a, 1.0))
	g.g.SetWeightedEdge(g.g.NewWeightedEdge(retval, b, 2.0))
	err := g.g.ApplyOp(engine.NewNeOperation(nil,nil,retSame), retval)
	return retval, err
}
