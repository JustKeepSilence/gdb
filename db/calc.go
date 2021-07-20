// +build gdbServer

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
		vm.Set("getFloatHData", gdb.getFloatHData)
		vm.Set("getIntHData", gdb.getIntHData)
		vm.Set("getStringHData", gdb.getStringHData)
		vm.Set("getBoolHData", gdb.getBoolHData)
		vm.Set("getFloatHDataWithTs", gdb.getFloatHDataWithTs)
		vm.Set("getIntHDataWithTs", gdb.getIntHDataWithTs)
		vm.Set("getStringHDataWithTs", gdb.getStringHDataWithTs)
		vm.Set("getBoolHDataWithTs", gdb.getBoolHDataWithTs)
		vm.Set("writeFloatRtData", gdb.writeFloatRtData)
		vm.Set("writeIntRtData", gdb.writeIntRtData)
		vm.Set("writeStringRtData", gdb.writeStringRtData)
		vm.Set("writeBoolRtData", gdb.writeBoolRtData)
		vm.Set("getTimeStamp", gdb.getUnixTimeStamp)
		vm.Set("getNowTime", gdb.getNowTime)
		vm.Set("console", gdb.console)
		p, runError = goja.Compile("main.js", expression, false)
		if p == nil {
			return
		}
		result, runError = vm.RunProgram(p)
	})
	if runError != nil {
		return calculationResult{}, runError
	}
	return calculationResult{result.Export()}, nil
}

func (gdb *Gdb) getCalculationItem(condition string) (calcItemsInfo, error) {
	rows, err := gdb.query("select id, description, expression, status, duration, errorMessage, createTime, updatedTime from calc_cfg where " + condition)
	if err != nil {
		return calcItemsInfo{}, err
	}
	return calcItemsInfo{rows}, nil
}

func (gdb *Gdb) updateCalculationItem(info updatedCalcInfo) (TimeRows, error) {
	st := time.Now()
	info.UpdatedTime = time.Now().Format(timeFormatString)
	sqlTemplate := template.Must(template.New("updatedCalcTemplate").Parse(`update calc_cfg set description='{{.Description}}', expression='{{.Expression}}', duration='{{.Duration}}', updatedTime='{{.UpdatedTime}}', errorMessage='' where id={{.Id}}`))
	var b bytes.Buffer
	if err := sqlTemplate.Execute(&b, info); err != nil {
		return TimeRows{}, err
	} else {
		sqlString := b.String()
		if r, err := gdb.updateItem(sqlString); err != nil {
			return TimeRows{}, err
		} else {
			return TimeRows{EffectedRows: int(r), Times: time.Since(st).Milliseconds()}, nil
		}
	}
}
