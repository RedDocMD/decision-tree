package ida3

import (
	"RedDocMD/decision_tree/parser"
	"RedDocMD/decision_tree/utils"
	"errors"
	"fmt"
	"log"
	"math"

	"github.com/goccy/go-graphviz"
	"github.com/goccy/go-graphviz/cgraph"
)

func entropyGain(rows []parser.Row, attribute string, data *parser.InputData) (float64, error) {
	originalEntropy := utils.Entropy(rows)
	partitions, err := utils.AttributePartition(data, rows, attribute)
	if err != nil {
		return 0.0, errors.New("Invalid atttribute")
	}
	newTotalEntropy := 0.0
	for _, partitionedRows := range partitions {
		newEntropy := utils.Entropy(partitionedRows)
		newTotalEntropy += float64(len(partitionedRows)) / float64(len(rows)) * newEntropy
	}
	return originalEntropy - newTotalEntropy, nil
}

func containsString(arr []string, elem string) bool {
	for _, val := range arr {
		if elem == val {
			return true
		}
	}
	return false
}

func attributeForMaxEntropyGain(rows []parser.Row, data *parser.InputData, omitList []string) string {
	maxGain := 0.0
	bestAttribute := ""
	for _, attribute := range data.AttributeNames {
		if !containsString(omitList, attribute) {
			gain, _ := entropyGain(rows, attribute, data)
			if gain >= maxGain {
				maxGain = gain
				bestAttribute = attribute
			}
		}
	}
	return bestAttribute
}

// DecisionTree represents the decision tree generated by IDA3 algorithm
type DecisionTree struct {
	Attribute                 string
	Parent                    *DecisionTree
	ParentVariant             string
	Children                  map[string]*DecisionTree // Key values are variant names
	IsLeaf                    bool
	LeafAnswer                bool
	previousAttributes        []string
	LeafOccurrenceProbability float64
	GraphvizNode              *cgraph.Node
}

func newDecisionTree(attribute string, parentAttribute *DecisionTree, parentVariant string) *DecisionTree {
	tree := new(DecisionTree)
	tree.Attribute = attribute
	tree.Parent = parentAttribute
	tree.ParentVariant = parentVariant
	tree.Children = make(map[string]*DecisionTree)
	tree.IsLeaf = false
	tree.previousAttributes = make([]string, 0)
	tree.LeafOccurrenceProbability = -1.0
	return tree
}

const eps float64 = 1e-6

func ida3Internal(rows []parser.Row, data *parser.InputData, tree *DecisionTree) {
	baseEntropy := utils.Entropy(rows)
	bestAttribute := attributeForMaxEntropyGain(rows, data, tree.previousAttributes)
	if math.Abs(baseEntropy) <= eps || bestAttribute == "" {
		tree.IsLeaf = true
		result, cnt := mostCommonResult(rows)
		tree.LeafAnswer = result
		tree.LeafOccurrenceProbability = float64(cnt) / float64(len(rows))
	} else {
		partitionedRows, _ := utils.AttributePartition(data, rows, bestAttribute)
		tree.Attribute = bestAttribute
		for _, variant := range data.AttributesMap[bestAttribute].Values {
			if len(partitionedRows[variant]) > 0 {
				subTree := newDecisionTree("", tree, variant)
				subTree.previousAttributes = append(subTree.previousAttributes, tree.previousAttributes...)
				subTree.previousAttributes = append(subTree.previousAttributes, bestAttribute)
				ida3Internal(partitionedRows[variant], data, subTree)
				tree.Children[variant] = subTree
			}
		}
	}
}

func mostCommonResult(rows []parser.Row) (bool, int) {
	trueCount := 0
	falseCount := 0
	for _, row := range rows {
		if row.Result {
			trueCount++
		} else {
			falseCount++
		}
	}
	return trueCount >= falseCount, max(trueCount, falseCount)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// IDA3 performs the IDA3 algorithm on the data passed and returns a decision tree
func IDA3(data *parser.InputData) *DecisionTree {
	decisionTree := newDecisionTree("", nil, "")
	ida3Internal(data.Rows, data, decisionTree)
	return decisionTree
}

func (tree *DecisionTree) String() string {
	out := ""
	depth := len(tree.previousAttributes)
	if tree.IsLeaf {
		for i := 0; i < depth; i++ {
			out += " "
		}
		out += fmt.Sprintf("%t %.3f\n", tree.LeafAnswer, tree.LeafOccurrenceProbability)
	} else {
		for variant, subTree := range tree.Children {
			for i := 0; i < depth; i++ {
				out += " "
			}
			out += tree.Attribute + " " + variant
			out += "\n"
			out += subTree.String()
		}
	}
	return out
}

// ToGraphvizPNG generates a SVG with a Graphviz dot graph
func (tree *DecisionTree) ToGraphvizPNG(filename string) {
	g := graphviz.New()
	graph, err := g.Graph()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := graph.Close(); err != nil {
			log.Fatal(err)
		}
		g.Close()
	}()
	stack := make([]*DecisionTree, 0)
	stack = append(stack, tree)
	counter := make(map[string]int)
	for len(stack) > 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		var name string
		if top.IsLeaf {
			name = fmt.Sprintf("%t", top.LeafAnswer)
		} else {
			name = top.Attribute
		}
		limit := counter[top.Attribute]
		counter[top.Attribute]++
		for i := 0; i < limit; i++ {
			// Hack to avoid naming nodes the same
			name += " "
		}
		node, _ := graph.CreateNode(name)
		top.GraphvizNode = node
		parent := top.Parent
		if parent != nil {
			edge, _ := graph.CreateEdge(top.ParentVariant, parent.GraphvizNode, node)
			edge.SetLabel(top.ParentVariant)
		}
		if !top.IsLeaf {
			for _, subTree := range top.Children {
				stack = append(stack, subTree)
			}
		}
	}
	if err := g.RenderFilename(graph, graphviz.SVG, filename); err != nil {
		log.Fatal("Failed to render ", filename, "\n", err)
	}
}
