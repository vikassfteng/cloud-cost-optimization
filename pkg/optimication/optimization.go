package optimization

import (
	"log"

	"github.com/whipstik/go-glpk"
)

// OptimizeCosts runs the linear programming optimization algorithm
func OptimizeCosts() {
	// Create a new LP problem
	lp := glpk.NewProblem()
	defer lp.Free()

	// Set objective function direction (minimization)
	lp.SetObjective(glpk.ObjectiveMinimize)

	// Define the number of variables and constraints
	lp.AddCols(2) // e.g., x1 and x2
	lp.AddRows(2) // e.g., constraints

	// Set variable names and bounds
	lp.SetColName(1, "x1")
	lp.SetColName(2, "x2")
	lp.SetColBnds(1, glpk.BoundLow, 0, glpk.INF)
	lp.SetColBnds(2, glpk.BoundLow, 0, glpk.INF)

	// Set coefficients for the objective function
	lp.SetObjCoef(1, 10) // Cost of x1
	lp.SetObjCoef(2, 15) // Cost of x2

	// Define constraints
	lp.SetRowName(1, "constraint1")
	lp.SetRowName(2, "constraint2")
	lp.SetRowBnds(1, glpk.BoundUp, 0, 100) // e.g., a1*x1 + a2*x2 <= 100
	lp.SetRowBnds(2, glpk.BoundUp, 0, 200) // e.g., b1*x1 + b2*x2 <= 200

	// Set coefficients for constraints
	lp.SetMatCoef(1, 1, 5)  // Coefficient of x1 in constraint1
	lp.SetMatCoef(1, 2, 10) // Coefficient of x2 in constraint1
	lp.SetMatCoef(2, 1, 20) // Coefficient of x1 in constraint2
	lp.SetMatCoef(2, 2, 30) // Coefficient of x2 in constraint2

	// Solve the problem
	lp.Solve(glpk.Minimize)

	// Retrieve and display results
	x1 := lp.GetColPrim(1)
	x2 := lp.GetColPrim(2)
	log.Printf("Optimal values: x1 = %f, x2 = %f", x1, x2)
}
