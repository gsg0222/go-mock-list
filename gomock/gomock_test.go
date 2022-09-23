package main

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	mock_interfaces "github.com/gsg0222/go-mock-list/gomock/mock"
)
func Test(t *testing.T) {

	// この辺はおまじない
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// モックの生成
	mock := mock_interfaces.NewMockFoo(ctrl)

	// 期待される引数と呼び出された場合の戻り値を設定
	mock.EXPECT().Bar(1).Return(2, nil)

	// 実際に期待される引数を渡すと設定した戻り値が帰ってくる
	result, _ := mock.Bar(1)
	fmt.Println(result)

	// もし期待していない数値が入力されたらテストは失敗する
	// result2, _ := mock.Bar(2)

}