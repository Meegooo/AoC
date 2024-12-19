package matrix

import (
	"math"
)

func ForwardElim(mtx Matrix2D[float64]) int {
	if mtx.Rows+1 != mtx.Columns {
		panic("matrix is not augmented")
	}
	for k := range mtx.Rows {
		iMax := k
		vMax := mtx.Get(iMax, k)
		for i := k + 1; i < mtx.Rows; i++ {
			if vMax < math.Abs(mtx.Get(i, k)) {
				vMax = mtx.Get(i, k)
				iMax = i
			}
		}
		if mtx.Get(k, iMax) == 0 {
			return k
		}
		if iMax != k {
			mtx.SwapRows(iMax, k)
		}
		for i := k + 1; i < mtx.Rows; i++ {
			f := mtx.Get(i, k) / mtx.Get(k, k)
			for j := k + 1; j <= mtx.Rows; j++ {
				newVal := mtx.Get(i, j) - mtx.Get(k, j)*f
				mtx.Set(i, j, newVal)
			}
			mtx.Set(i, k, 0)
		}
		//mtx.Print(func(f float64) string {
		//	return fmt.Sprintf("%f ", f)
		//})
	}
	//mtx.Print(func(f float64) string {
	//	return fmt.Sprintf("%f ", f)
	//})
	return -1
}

func BackSub(mtx Matrix2D[float64]) []float64 {
	solution := make([]float64, mtx.Rows)
	for i := mtx.Rows - 1; i >= 0; i-- {
		solution[i] = mtx.Get(i, mtx.Rows)
		for j := i + 1; j < mtx.Rows; j++ {
			solution[i] -= mtx.Get(i, j) * solution[j]
		}
		solution[i] /= mtx.Get(i, i)
	}
	return solution
}
