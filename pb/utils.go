package pb

func ToBurndownSparseMatrix(matrix [][]int64, name string) *BurndownSparseMatrix {
  r := BurndownSparseMatrix{
	  Name: name,
	  NumberOfRows: int32(len(matrix)),
	  NumberOfColumns: int32(len(matrix[len(matrix)-1])),
	  Row: make([]*BurndownSparseMatrixRow, len(matrix)),
  }
  for i, status := range matrix {
	  nnz := make([]uint32, 0, len(status))
	  changed := false
	  for j := range status {
		  v := status[len(status) - 1 - j]
		  if v < 0 {
			  v = 0
		  }
		  if !changed {
			  changed = v != 0
		  }
		  if changed {
			  nnz = append(nnz, uint32(v))
		  }
	  }
	  r.Row[i] = &BurndownSparseMatrixRow{
		  Column: make([]uint32, len(nnz)),
	  }
	  for j := range nnz {
		  r.Row[i].Column[j] = nnz[len(nnz) - 1 - j]
	  }
	}
	return &r
}

func DenseToCompressedSparseRowMatrix(matrix [][]int64) *CompressedSparseRowMatrix {
	r := CompressedSparseRowMatrix{
		NumberOfRows: int32(len(matrix)),
		NumberOfColumns: int32(len(matrix[0])),
		Data: make([]int64, 0),
		Indices: make([]int32, 0),
		Indptr: make([]int64, 1),
	}
	r.Indptr[0] = 0
	for _, row := range matrix {
		for x, col := range row {
			nnz := 0
			if col != 0 {
				r.Data = append(r.Data, col)
				r.Indices = append(r.Indices, int32(x))
				nnz += 1
			}
			r.Indptr = append(r.Indptr, r.Indptr[len(r.Indptr) - 1] + int64(nnz))
		}
	}
	return &r
}

func MapToCompressedSparseRowMatrix(matrix []map[int]int64) *CompressedSparseRowMatrix {
	r := CompressedSparseRowMatrix{
		NumberOfRows: int32(len(matrix)),
		NumberOfColumns: int32(len(matrix)),
		Data: make([]int64, 0),
		Indices: make([]int32, 0),
		Indptr: make([]int64, 1),
	}
	r.Indptr[0] = 0
	for _, row := range matrix {
		for x, col := range row {
			r.Data = append(r.Data, col)
			r.Indices = append(r.Indices, int32(x))
			r.Indptr = append(r.Indptr, r.Indptr[len(r.Indptr) - 1] + int64(len(row)))
		}
	}
	return &r
}