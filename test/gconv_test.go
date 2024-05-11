/**
 * @package test
 * @file      : gconv_test.go
 * @author    : LeiXiaoTian
 * @contact   : 1124378213@qq.com
 * @time      : 2024/5/11 18:27
 **/
package test

import (
	"fmt"
	"github.com/leiphp/unit-go-sdk/pkg/gconv"
	"testing"
)

func TestGconvString(t *testing.T) {
	t.Log("========== TestGconvString ==========")
	gconv.String(1)
	fmt.Println("result:", "")
}
