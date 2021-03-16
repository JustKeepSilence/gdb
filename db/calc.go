/*
creatTime: 2020/12/26
creator: JustKeepSilence
github: https://github.com/JustKeepSilence
goVersion: 1.15.3
*/

package db

import (
	"github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/eventloop"
)

// test
func (gdb *Gdb) testCalculation(expression string) (interface{}, error) {
	loop := eventloop.NewEventLoop()
	var runError error
	var result goja.Value
	var p *goja.Program
	loop.Run(func(vm *goja.Runtime) {
		vm.Set("getRtData", gdb.getRtData)           // get realTime data
		vm.Set("getHData", gdb.getHData)             // get history
		vm.Set("writeRtData", gdb.BatchWrite)        // write data
		vm.Set("getTimeStamp", gdb.getUnixTimeStamp) // get timeStamp of given time string
		vm.Set("getNowTime", gdb.getNowTime)         // get current Time
		vm.Set("getTime", gdb.getTime)               // get time
		p, runError = goja.Compile("main.js", expression, false)
		if p == nil {
			return
		}
		result, runError = vm.RunProgram(p)
	})
	if runError != nil {
		return nil, runTimeError{"runTimeError:" + runError.Error()}
	}
	return result.Export(), nil
}

func (gdb *Gdb) getCalculationItem(condition string) ([]map[string]string, error) {
	rows, err := query(gdb.ItemDbPath, "select id, description, expression, status, duration, errorMessage, createTime, updatedTime from calc_cfg where "+condition)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func (gdb *Gdb) updateCalculationItem(info updatedCalculationInfo) (Rows, error) {
	id, description, expression, duration := info.Id, info.Description, info.Expression, info.Duration
	r, err := updateItem(gdb.ItemDbPath, "update calc_cfg set description='"+description+"', expression='"+expression+"', duration='"+duration+"' where id="+id)
	if err != nil {
		return Rows{}, nil
	}
	return Rows{int(r)}, nil
}
