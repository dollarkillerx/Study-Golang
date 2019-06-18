/**
* Created by GoLand
* User: dollarkiller
* Date: 19-6-18
* Time: 下午3:08
* */
package err

func CheckErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}
