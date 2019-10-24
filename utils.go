/*
-------------------------------------------------
   Author :       Zhang Fan
   date：         2019/10/24
   Description :
-------------------------------------------------
*/

package zver

import (
    "errors"
    "strconv"
    "strings"
)

// 将文本切割并转换为uint类型
func SplitToUint(text, sep string) ([]uint, error) {
    if text == "" {
        return nil, nil
    }

    vs := strings.Split(text, sep)
    outs := make([]uint, len(vs))
    for i, v := range vs {
        n, err := strconv.Atoi(v)
        if err != nil {
            return nil, errors.New("某个值无法转换为数字")
        }
        if n < 0 {
            return nil, errors.New("所有值必须>=0")
        }
        outs[i] = uint(n)
    }
    return outs, nil
}
