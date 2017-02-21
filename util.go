package talib



func FormatFloat64(srcData []float64, begIdx, NBElement int) (outData []float64) {
    outData = make([]float64, begIdx)
    outData = append(outData, srcData[0:NBElement]...)
    return
}

func FormatInt32(srcData []int32, begIdx, NBElement int) (outData []int32) {
    outData = make([]int, begIdx)
    outData = append(outData, srcData[0:NBElement]...)
    return
}

