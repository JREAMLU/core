package mysql

import

// "github.com/astaxie/beego/orm"
_ "github.com/go-sql-driver/mysql"

// "github.com/go-xorm/xorm"

/**
 *	@auther		jream.lu
 *	@intro		封装原生查询语句
 *	@return 	slice maps
 */
/*
func Select(params []interface{}, sql string) (maps []orm.Params, total int64) {
	o := orm.NewOrm()
	num, err := o.Raw(sql, params).Values(&maps)
	if err != nil {
		//log
		fmt.Println(err)
	}
	return maps, num
}
*/

/*
func Insert() {

}

func InsertMulti() {

}




func Update() {

}

func UpdateMulti() {

}

func Delete() {

}

func DeleteMulti() {

}
*/

/**
 *	@auther		jream.lu
 *	@intro		封装原生查询语句
 *	@return 	slice lists
 */
/*
func SelectList(params []interface{}, sql string) (lists []orm.ParamsList, total int64) {
	o := orm.NewOrm()
	num, err := o.Raw(sql, params).ValuesList(&lists)
	if err != nil {
		//log
		fmt.Println(err)
	}
	return lists, num
}
*/
