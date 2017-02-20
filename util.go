package talib



func FormatFloat64(srcData []float64, begIdx, NBElement int) (outData []float64) {
    outData = make([]float64, begIdx)
    outData = append(outData, srcData[0:NBElement]...)
    return
}

func FormatInt(srcData []int, begIdx, NBElement int) (outData []int) {
    outData = make([]int, begIdx)
    outData = append(outData, srcData[0:NBElement]...)
    return
}

