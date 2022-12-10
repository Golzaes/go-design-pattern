package factory

import "fmt"

type (
	DAOFactory interface {
		CreateOrderMainDAO() OrderMainDAO
		CreateOrderDetailDAO() OrderDetailDAO
	}
	OrderMainDAO   interface{ SaveOrderMain() }
	OrderDetailDAO interface{ SaveOrderDetail() }
)

type (
	TxtDAOFactory  struct{}
	JSONDAOFactory struct{}
	RDBDAOFactory  struct{}
)

type (
	TxtMainDao    struct{}
	TxtDetailDao  struct{}
	RDBMainDao    struct{}
	RDBDetailDao  struct{}
	JSONMainDao   struct{}
	JSONDetailDao struct{}
)

func (*TxtDAOFactory) CreateOrderMainDAO() OrderMainDAO      { return &TxtMainDao{} }
func (*TxtDAOFactory) CreateOrderDetailDAO() OrderDetailDAO  { return &TxtDetailDao{} }
func (*RDBDAOFactory) CreateOrderMainDAO() OrderMainDAO      { return &RDBMainDao{} }
func (*RDBDAOFactory) CreateOrderDetailDAO() OrderDetailDAO  { return &RDBDetailDao{} }
func (*JSONDAOFactory) CreateOrderMainDAO() OrderMainDAO     { return &JSONMainDao{} }
func (*JSONDAOFactory) CreateOrderDetailDAO() OrderDetailDAO { return &JSONDetailDao{} }

func (*TxtMainDao) SaveOrderMain()      { fmt.Println(`save main order message to txt`) }
func (*TxtDetailDao) SaveOrderDetail()  { fmt.Println(`save detail order message to txt`) }
func (*RDBMainDao) SaveOrderMain()      { fmt.Println(`save main order message to rdb`) }
func (*RDBDetailDao) SaveOrderDetail()  { fmt.Println(`save detail order message to rdb`) }
func (*JSONMainDao) SaveOrderMain()     { fmt.Println(`save main order message to json`) }
func (*JSONDetailDao) SaveOrderDetail() { fmt.Println(`save detail order message to json`) }
