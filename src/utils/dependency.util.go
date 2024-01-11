package utils

import (
	"log"
	"reflect"
)

/*
依存するオブジェクトがフィールドに入れられているか確認
依存注入されていない場合、該当が実行されるまでエラーとならないためアプリ起動時に確認
ポインタだとフィールド取得できないのでstructを引数にとる
*/
func CheckDependencies(anyStruct any) {
	value := reflect.ValueOf(anyStruct)
	numFields := value.NumField()
	for i := 0; i < numFields; i++ {
		if value.Field(i).IsNil() {
			log.Fatalf("依存関係を解決できませんでした: %s", value.Type())
		}
	}
}
