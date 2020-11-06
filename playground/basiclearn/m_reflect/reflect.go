package m_reflect

import (
	"fmt"
	"log"
	"reflect"
)

type order struct {
	id   int
	name string
}
type employee struct {
	name   string
	age    int
	salary float32
}

func insert(t interface{}) {
	if reflect.ValueOf(t).Kind() == reflect.Struct {
		table := reflect.TypeOf(t).Name()
		sql := fmt.Sprintf("insert into %s values(", table)
		v := reflect.ValueOf(t)
		for i := 0; i < v.NumField(); i++ {
			switch v.Field(i).Kind() {
			case reflect.Int:
				if i == 0 {
					sql = fmt.Sprintf("%s%d", sql, v.Field(i).Int())
				} else {
					sql = fmt.Sprintf("%s, %d", sql, v.Field(i).Int())
				}
			case reflect.String:
				if i == 0 {
					sql = fmt.Sprintf("%s%s", sql, v.Field(i).String())
				} else {
					sql = fmt.Sprintf("%s, %s", sql, v.Field(i).String())
				}
			case reflect.Float32:
				if i == 0 {
					sql = fmt.Sprintf("%s%f", sql, v.Field(i).Float())
				} else {
					sql = fmt.Sprintf("%s, %f", sql, v.Field(i).Float())
				}
			default:
				fmt.Println("unsupported type")
			}
		}
		sql = fmt.Sprintf("%s)", sql)
		fmt.Println(sql)
		return
	}
	fmt.Println("unsupported type")
}

type User struct {
	Name string `json:"name",yaml:"Name"`
	Age  int    `json:"age,omitempty"`
	From string `json:"from"`
}

func (u *User) Hello() {
	fmt.Println("hello")
}

func (u User) Hi() {
	fmt.Println("hi")
}

func TestReflect(t interface{}) {

	objT := reflect.TypeOf(t)
	objV := reflect.ValueOf(t)

	//获取去这个类型的名称
	fmt.Println("这个类型的名称是:", objT.Name())
	//
	for i := 0; i < objT.NumField(); i++ {
		//从0开始获取结构体所包含的key
		key := objT.Field(i)
		//从0开始通过interface方法来获取key所对应的值
		value := objV.Field(i).Interface()

		fmt.Printf("第%d个字段是：%s:%v = %v \n", i+1, key.Name, key.Type, value)
	}

	for i := 0; i < objT.NumMethod(); i++ {
		m := objT.Method(i)
		fmt.Printf("第%d个方法是：%s:%v\n", i+1, m.Name, m.Type)
	}

}

func modifyByReflect() {
	i := 20
	fmt.Println("before i =", i)
	e := reflect.Indirect(reflect.ValueOf(&i))
	// 或者e := reflect.ValueOf(&i).Elem()
	fmt.Println("canset: ", e.CanSet())
	if e.CanSet() {
		e.SetInt(22)
	}
	fmt.Println("after i =", i)
}

// 反射字段值修改
func TestReflectField() {
	u := User{"tom", 27, "beijing"}
	// elem := reflect.Indirect(reflect.ValueOf(&u)) // 等效于下方
	elem := reflect.ValueOf(&u).Elem()
	t := elem.Type()

	fmt.Println("before user:", u)
	for i := 0; i < t.NumField(); i++ {
		// 反射获取字段的元信息，例如：名称、Tag 等
		ft := t.Field(i)
		fmt.Printf("field name:%s\t", ft.Name)
		fmt.Printf("Tags:%s\t", ft.Tag)                 // 获取所有tag
		fmt.Printf("Tag json:%s\n", ft.Tag.Get("json")) // 根据tag名字获取

		// 反射修改字段的值（）
		fv := elem.Field(i)
		if fv.CanSet() {
			if fv.Kind() == reflect.Int {
				fmt.Println("change age to 30")
				fv.SetInt(30)
			}
		}
	}
	fmt.Println("after user:", u)
}

func TestReflectKind() {
	u := User{"tom", 27, "beijing"}
	v := reflect.ValueOf(u)
	t := reflect.TypeOf(u)
	// 获取 Kind 类型
	k := t.Kind()
	fmt.Println(k)
	k1 := v.Kind()
	fmt.Println(k1)
	fmt.Println(k == k1) // true
}

// 反射值接收器的方法
func TestCallMethod() {
	u := User{"tom", 27, "beijing"}
	rv := reflect.Indirect(reflect.ValueOf(u))
	rt := rv.Type()
	fmt.Println(rv.NumMethod())

	for i := 0; i < rv.NumMethod(); i++ {
		mv := rv.Method(i) // 通过value
		mt := rt.Method(i) // 通过type获取
		fmt.Println(mt)

		param := []reflect.Value{reflect.ValueOf(u)}
		mt.Func.Call(param) // 同样可以调用，不过len(params)必须>0

		if mv.IsValid() {
			if v := mv.Call(nil); len(v) > 0 {
				if err, ok := v[0].Interface().(error); ok {
					log.Fatal(err)
				}
			}
		}

	}

	//通过函数名字调用
	fm := rv.MethodByName("Hi")
	if fm.IsValid() {
		if v := fm.Call(nil); len(v) > 0 {
			if err, ok := v[0].Interface().(error); ok {
				log.Fatal(err)
			}
		}
	}
}

// 反射指针接收器的方法
func TestCallMethodPtr() {
	u := User{"tom", 27, "beijing"}
	rv := reflect.Indirect(reflect.ValueOf(u))
	rt := rv.Type()

	t := reflect.PtrTo(rt)
	for i := 0; i < t.NumMethod(); i++ {
		mt := t.Method(i) // 通过value

		fmt.Println(mt)

		param := []reflect.Value{reflect.ValueOf(&u)}
		if mt.Func.IsValid() {
			if v := mt.Func.Call(param); len(v) > 0 {
				if err, ok := v[0].Interface().(error); ok {
					log.Fatal(err)
				}
			}
		}

	}
}

// 反射对象value还原
func TestBack() {
	u := User{"tom", 27, "beijing"}
	v := reflect.ValueOf(u)
	if u, ok := v.Interface().(User); ok {
		fmt.Println(u)
	}

	var num float64 = 1.2345

	pointer := reflect.ValueOf(&num)
	value := reflect.ValueOf(num)

	// 可以理解为“强制转换”，但是需要注意的时候，转换的时候，如果转换的类型不完全符合，则直接panic
	// Golang 对类型要求非常严格，类型一定要完全符合
	// 如下两个，一个是*float64，一个是float64，如果弄混，则会panic
	convertPointer := pointer.Interface().(*float64)
	convertValue := value.Interface().(float64)

	fmt.Println(convertPointer, *convertPointer)
	fmt.Println(convertValue)
}

// 位置原有类型的情况转换

func TestUnknown(input interface{}) {
	getType := reflect.TypeOf(input)
	fmt.Println("get Type is :", getType.Name())

	getValue := reflect.ValueOf(input)
	fmt.Println("get all Fields is:", getValue)

	// 获取字段
	// 1. 先获取interface的reflect.Type，然后通过NumField进行遍历
	// 2. 再通过reflect.Type的Field获取其Field
	// 3. 最后通过reflect.Value获取Field的Interface()得到对应的value
	for i := 0; i < getType.NumField(); i++ {
		field := getType.Field(i)
		value := getValue.Field(i).Interface()
		fmt.Printf("%s: %v = %v\n", field.Name, field.Type, value)
	}

	// 获取方法
	// 1. 先获取interface的reflect.Type，然后通过.NumMethod进行遍历
	for i := 0; i < getType.NumMethod(); i++ {
		m := getType.Method(i)
		fmt.Printf("%s: %v\n", m.Name, m.Type)
	}
}
