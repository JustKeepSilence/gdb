// +build gdbServer

/*
createTime: 2021/5/29
creator: JustKeepSilence
github: https://github.com/JustKeepSilence
goVersion: 1.16
*/

package db

import (
	"fmt"
	. "github.com/ahmetb/go-linq/v3"
	"github.com/casbin/casbin/v2/model"
	"github.com/casbin/casbin/v2/persist"
	"strings"
)

// implement the interface of adapter

func (gdb *Gdb) LoadPolicy(model model.Model) error {
	if rows, err := gdb.query("select routeRoles from route_cfg where 1=1"); err != nil {
		return err
	} else {
		for _, row := range rows {
			routeRoles := []string{}
			if err := json.Unmarshal(convertStringToByte(row["routeRoles"]), &routeRoles); err != nil {
				return err
			} else {
				for _, routeRole := range routeRoles {
					persist.LoadPolicyLine(routeRole, model) // routeRole is like p, admin, /testTom, GET
				}
			}
		}
	}
	return nil
}

func (gdb *Gdb) SavePolicy(_ model.Model) error {
	return nil
}

func (gdb *Gdb) AddPolicy(_ string, _ string, actions []string) error {
	// "added_user", "data1", "read"
	name, route, action := actions[0], strings.Title(actions[1]), actions[2]
	m := gdb.e.GetModel()
	for _, ast := range m["p"] {
		if From(ast.Policy).IndexOf(func(item interface{}) bool {
			r := item.([]string)
			return r[0] == name && r[1] == route && r[2] == action
		}) > 0 {
			return fmt.Errorf("cannot add the same routing permissions")
		} else {
			// not existed
			// add policy to db
			if row, err := gdb.query("select routeRoles from route_cfg where userName='" + name + "'"); err != nil {
				return err
			} else {
				role := row[0]["routeRoles"]
				r := "p," + name + "," + route + "," + action
				var value string
				if role == "[]" {
					value = `["` + r + `"]` // if role is []
				} else {
					value = strings.Replace(role, "]", `,"`+r+`"]`, -1)
				}
				if _, err := gdb.updateItem("update route_cfg set routeRoles='" + value + "' where userName ='" + name + "'"); err != nil {
					return err
				} else {
					ast.Policy = append(ast.Policy, []string{name, route, action}) // add policy to model
				}
			}
		}
	}
	return nil
}

func (gdb *Gdb) RemovePolicy(_ string, _ string, actions []string) error {
	// "alice", "data1", "read"
	name, route, action := actions[0], strings.Title(actions[1]), actions[2]
	_ = gdb.e.LoadPolicy()
	m := gdb.e.GetModel()
	for _, ast := range m["p"] {
		if index := From(ast.Policy).IndexOf(func(item interface{}) bool {
			r := item.([]string)
			return r[0] == name && r[1] == route && r[2] == action
		}); index < 0 {
			return fmt.Errorf("routing permissions not existed")
		} else {
			us := []string{}
			for i, p := range ast.Policy {
				if i != index && p[0] == name {
					us = append(us, "p,"+strings.Join(p, ","))
				}
			}
			nr, _ := json.Marshal(us)
			if _, err := gdb.updateItem("update route_cfg set routeRoles='" + string(nr) + "' where userName='" + name + "'"); err != nil {
				return err
			} else {
				_ = gdb.e.LoadPolicy()
			}
		}
	}
	return nil
}

func (gdb *Gdb) RemoveFilteredPolicy(_ string, _ string, _ int, _ ...string) error {
	return nil
}
