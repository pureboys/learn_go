package router

import (
	"demo/day14/SecKill/SecAdmin/controller/product"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/product/list", &product.ProductController{}, "*:ListProduct")
}
