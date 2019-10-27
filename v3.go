/*
-------------------------------------------------
   Author :       Zhang Fan
   date：         2019/10/24
   Description :
-------------------------------------------------
*/

package zver

import (
    "fmt"
    "github.com/zlyuancn/zerrors"
    "strings"
)

type V3 struct {
    Main  uint // 主版本
    Minor uint // 次版本
    Mini  uint // 小版本
}

// 从版本文本中解析
func (v3 *V3) Parser(version, sep string) error {
    if version == "" {
        return zerrors.New("version是空的")
    }

    vs, err := SplitToUint(version, sep)
    if err != nil {
        return zerrors.Wrap(err, "解析版本号失败")
    }

    if len(vs) == 1 {
        vs = append(vs, 0, 0)
    } else if len(vs) == 2 {
        vs = append(vs, 0)
    } else if len(vs) != 3 {
        return zerrors.New("无法解析版本号")
    }

    v3.Main = vs[0]
    v3.Minor = vs[1]
    v3.Mini = vs[2]
    return nil
}

// 从有前缀的版本文本中解析
func (v3 *V3) ParserHasPrefix(version, sep, prefix string) error {
    if strings.Index(version, prefix) != 0 {
        return zerrors.New("版本前缀不正确")
    }
    return v3.Parser(version[len(prefix):], sep)
}

// 转为默认格式的版本文本
func (v3 *V3) String() string {
    return fmt.Sprintf("v%d.%d.%d", v3.Main, v3.Minor, v3.Mini)
}

// 判断当前版本大于传入的版本
func (v3 *V3) Gt(v *V3) bool {
    if v3.Main < v.Main {
        return false
    }
    if v3.Main > v.Main {
        return true
    }

    if v3.Minor < v.Minor {
        return false
    }
    if v3.Minor > v.Minor {
        return true
    }
    return v3.Mini > v.Mini
}

// 判断当前版本大于或等于传入的版本
func (v3 *V3) Gte(v *V3) bool {
    return !v3.Lt(v)
}

// 判断当前版本小于传入的版本
func (v3 *V3) Lt(v *V3) bool {
    if v3.Main < v.Main {
        return true
    }
    if v3.Main > v.Main {
        return false
    }

    if v3.Minor < v.Minor {
        return true
    }
    if v3.Minor > v.Minor {
        return false
    }
    return v3.Mini < v.Mini
}

// 判断当前版本小于或等于传入的版本
func (v3 *V3) Lte(v *V3) bool {
    return !v3.Gt(v)
}

// 判断两个版本是否相等
func (v3 *V3) Eq(v *V3) bool {
    return v3.Main == v.Main && v3.Minor == v.Minor && v3.Mini == v.Mini
}

// 判断两个版本是否不相等
func (v3 *V3) Ne(v *V3) bool {
    return v3.Main != v.Main || v3.Minor != v.Minor || v3.Mini != v.Mini
}

// 转为版本文本
func (v3 *V3) ToText(sep string) string {
    return fmt.Sprintf("%d%s%d%s%d", v3.Main, sep, v3.Minor, sep, v3.Mini)
}

// 转为有前缀的版本文本
func (v3 *V3) ToTextHasPrefix(sep, prefix string) string {
    return fmt.Sprintf("%s%d%s%d%s%d", prefix, v3.Main, sep, v3.Minor, sep, v3.Mini)
}
