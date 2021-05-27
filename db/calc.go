// +build gdbClient

/*
creatTime: 2020/12/26
creator: JustKeepSilence
github: https://github.com/JustKeepSilence
goVersion: 1.16
*/

package db

import (
	"bytes"
	"github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/eventloop"
	"text/template"
	"time"
)

// test
func (gdb *Gdb) testCalculation(expression string) (calculationResult, error) {
	loop := eventloop.NewEventLoop()
	var runError error
	var result goja.Value
	var p *goja.Program
	loop.Run(func(vm *goja.Runtime) {
		vm.Set("getRtData", gdb.getRtData)
		vm.Set("getHData", gdb.getHData)
		vm.Set("getHDataWithTs", gdb.getHDataWithTs)
		vm.Set("writeRtData", gdb.writeRtData)
		vm.Set("getTimeStamp", gdb.getUnixTimeStamp)
		vm.Set("getNowTime", gdb.getNowTime)
		vm.Set("testItemValue", gdb.testItemValue)
		p, runError = goja.Compile("main.js", expression, false)
		if p == nil {
			return
		}
		result, runError = vm.RunProgram(p)
	})
	if runError != nil {
		return calculationResult{}, runTimeError{"runTimeError:" + runError.Error()}
	}
	return calculationResult{result.Export()}, nil
}

func (gdb *Gdb) getCalculationItem(condition string) (calcItemsInfo, error) {
	rows, err := query(gdb.ItemDbPath, "select id, description, expression, status, duration, errorMessage, createTime, updatedTime from calc_cfg where "+condition)
	if err != nil {
		return calcItemsInfo{}, err
	}
	return calcItemsInfo{rows}, nil
}

func (gdb *Gdb) updateCalculationItem(info updatedCalcInfo) (Rows, error) {
	info.UpdatedTime = time.Now().Format(timeFormatString)
	sqlTemplate := template.Must(template.New("updatedCalcTemplate").Parse(`update calc_cfg set description='{{.Description}}', expression='{{.Expression}}', duration='{{.Duration}}', updatedTime='{{.UpdatedTime}}', errorMessage='' where id={{.Id}}`))
	var b bytes.Buffer
	if err := sqlTemplate.Execute(&b, info); err != nil {
		return Rows{}, err
	} else {
		sqlString := b.String()
		if r, err := updateItem(gdb.ItemDbPath, sqlString); err != nil {
			return Rows{}, err
		} else {
			return Rows{int(r)}, nil
		}
	}
}
